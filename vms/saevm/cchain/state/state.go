// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package state manages the on-disk state for the C-Chain's cross-chain
// transactions.
package state

import (
	"errors"
	"sync/atomic"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"

	chainsatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

// These prefixes and keys are byte-compatible with the indices written by
// [state.AtomicTrie] and [state.AtomicRepository] so that SAE does not require
// a database migration.
var (
	triePrefix = []byte("atomicTrieDB") // See state.atomicTrieStoragePrefix

	commitPrefix  = prefixdb.MakePrefix([]byte("atomicTrieMetaDB"))                          // See state.atomicTrieMetaDBPrefix
	lastHeightKey = prefixdb.PrefixKey(commitPrefix, []byte("atomicTrieLastCommittedBlock")) // See state.lastCommittedKey

	txPrefix = prefixdb.MakePrefix([]byte("atomicTxDB")) // See state.atomicTxIDDBPrefix
)

// State holds the accepted transactions and the atomic-request trie.
//
// When applying operations, shared memory is updated atomically with the state.
//
// [State.Apply] and [State.Close] MUST NOT be called concurrently with
// themselves or each other. All other methods are safe to call concurrently.
//
// [State.Close] MUST be called when finished with the state to release
// resources.
type State struct {
	snowCtx *snow.Context
	db      database.Database
	trieDB  *triedb.Database

	currentRoot common.Hash

	// currentHeight is atomic to allow [State.CurrentHeight] to be called
	// concurrently with [State.Apply].
	currentHeight atomic.Uint64
}

// New initializes the state with db.
//
// TODO(#5375): Coreth's commitInterval must be reduced to 1 prior to
// transitioning to SAE. Otherwise, the atomic trie may not contain operations
// for recent blocks.
func New(snowCtx *snow.Context, db database.Database) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Coreth previously wrapped db in a versiondb before using
// [prefixdb.New]. To maintain byte compatibility with the existing
// trie, we must use [prefixdb.NewNested] rather than [prefixdb.New] and
// not compress the prefix.

// This trie is append only, so we only need to cache the
// leading edge.

// readLast returns the last trie root and height. If none have been written, it
// returns the empty root for the genesis block.
func readLast(db database.KeyValueReader) (common.Hash, uint64, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), 0, nil
}

func readRoot(db database.KeyValueReader, height uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func rootKey(height uint64) []byte { _ = "STUB: not implemented"; return nil }

// Apply persists the txs accepted at height. It applies their atomic
// operations to the trie, indexes the txs by ID, and applies the atomic
// operations to shared memory.
//
// Apply is a noop when height is not higher than [State.CurrentHeight].
func (s *State) Apply(height uint64, txs []*tx.Tx) error { _ = "STUB: not implemented"; return nil }

// During restarts, it is expected for SAE to reprocess already-applied
// heights. Shared memory is not safe to apply multiple times for the
// same height, so we MUST skip these duplicate applications.

// Committing the batch atomically with shared memory prevents duplicate
// shared memory operations in the event of a crash.
//
// TODO(StephenButtolph): Skip applying shared memory operations for bonus
// blocks.

// atomicRequests groups the atomic requests from txs by chainID.
func atomicRequests(txs []*tx.Tx) (map[ids.ID]*chainsatomic.Requests, error) {
	_ = "STUB: not implemented"
	// To produce a byte-identical trie, txs must be merged in txID order.
	// This matches the order they were originally read from the tx index when
	// the trie was first built. Without sorting, the PutRequests and
	// RemoveRequests within a chain's atomic.Requests could be appended in a
	// different order, changing the trie value.
	return nil, nil
}

var errCleanTrieAfterUpdates = errors.New("clean trie after updates")

// applyTrie writes the per-chain ops into the trie rooted at oldRoot, flushes
// the resulting trie to disk, and returns the new root.
func applyTrie(trieDB *triedb.Database, oldRoot common.Hash, height uint64, ops map[ids.ID]*chainsatomic.Requests) (common.Hash, error) {
	_ = "STUB: not implemented"
	// Most blocks don't have atomic requests, so we avoid any unnecessary trie
	// operations in that case.
	return *new(common.Hash), nil
}

// Since each map entry corresponds to a different entry in the trie, the
// trie root is order-independent.

// [hashdb.Database.Update] would attempt to RLP-decode collected leaves as
// [types.StateAccount] to reference storage subtries.

// TODO(StephenButtolph): There isn't a good reason for us to actually write the
// full tx to the database. We could just write the height, and then expect the
// caller to fetch the block to get the tx. Paying that extra block fetch on the
// read side is almost certainly worth avoiding the duplicate write, since this
// index is only read from the API.
func writeTx(db database.KeyValueWriter, height uint64, t *tx.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func txKey(id ids.ID) []byte { _ = "STUB: not implemented"; return nil }

// GetTx returns the tx with the given ID along with the block height it was
// accepted at.
func (s *State) GetTx(txID ids.ID) (*tx.Tx, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// GetRoot returns the atomic trie root at height.
func (s *State) GetRoot(height uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// CurrentHeight returns the highest height successfully applied via
// [State.Apply].
func (s *State) CurrentHeight() uint64 { _ = "STUB: not implemented"; return 0 }

// Close closes the state, releasing any resources.
func (s *State) Close() error { _ = "STUB: not implemented"; return nil }
