// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"time"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

const (
	sharedMemoryApplyBatchSize = 10_000 // specifies the number of atomic operations to batch progress updates
	progressLogFrequency       = 30 * time.Second
)

// AtomicBackend provides an interface to the atomic trie and shared memory.
// the AtomicTrie, AtomicRepository, and the VM's shared memory.
type AtomicBackend struct {
	codec        codec.Manager
	bonusBlocks  map[uint64]ids.ID // Map of height to blockID for blocks to skip indexing
	sharedMemory avalancheatomic.SharedMemory

	repo       *AtomicRepository
	atomicTrie *AtomicTrie

	lastAcceptedHash common.Hash
	verifiedRoots    map[common.Hash]*atomicState
}

// NewAtomicBackend creates an AtomicBackend from the specified dependencies
func NewAtomicBackend(
	sharedMemory avalancheatomic.SharedMemory,
	bonusBlocks map[uint64]ids.ID, repo *AtomicRepository,
	lastAcceptedHeight uint64, lastAcceptedHash common.Hash, commitInterval uint64,
) (*AtomicBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We call ApplyToSharedMemory here to ensure that if the node was shut down in the middle
// of applying atomic operations from state sync, we finish the operation to ensure we never
// return an atomic trie that is out of sync with shared memory.
// In normal operation, the cursor is not set, such that this call will be a no-op.

// initializes the atomic trie using the atomic repository height index.
// Iterating from the last committed height to the last height indexed
// in the atomic repository, making a single commit at the
// most recent height divisible by the commitInterval.
// Subsequent updates to this trie are made using the Index call as blocks are accepted.
// Note: this method assumes no atomic txs are applied at genesis.
func (a *AtomicBackend) initialize(lastAcceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil

	// track the last committed height and last committed root
}

// iterate by height, from [lastCommittedHeight+1] to [lastAcceptedBlockNumber]

// open the atomic trie at the last committed root

// Get the height and transactions for this iteration (from the key and value, respectively)
// iterate over the transactions, indexing them if the height is < commit height
// otherwise, add the atomic operations from the transaction to the uncommittedOpsMap

// combine atomic operations from all transactions at this block height

// Note: The atomic trie canonically contains the duplicate operations
// from any bonus blocks.

// Trie must be re-opened after committing (not safe for re-use after commit)

// check if there are accepted blocks after the last block with accepted atomic txs.

// ApplyToSharedMemory applies the atomic operations that have been indexed into the trie
// but not yet applied to shared memory for heights less than or equal to [lastAcceptedBlock].
// This executes operations in the range [cursorHeight+1, lastAcceptedBlock].
// The cursor is initially set by  MarkApplyToSharedMemoryCursor to signal to the atomic trie
// the range of operations that were added to the trie without being executed on shared memory.
func (a *AtomicBackend) ApplyToSharedMemory(lastAcceptedBlock uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// value of sharedMemoryCursor is either a uint64 signifying the
// height iteration should begin at or is a uint64+blockchainID
// specifying the last atomic operation that was applied to shared memory.
// To avoid applying the same operation twice, we call [it.Next()] in the
// latter case.

// If [height] is a bonus block, do not apply the atomic operations to shared memory

// Update the cursor to the key of the atomic operation being executed on shared memory.
// If the node shuts down in the middle of this function call, ApplyToSharedMemory will
// resume operation starting at the key immediately following [it.Key()].

// calling [sharedMemory.Apply] updates the last applied pointer atomically with the shared memory operation.

// MarkApplyToSharedMemoryCursor marks the atomic trie as containing atomic ops that
// have not been executed on shared memory starting at [previousLastAcceptedHeight+1].
// This is used when state sync syncs the atomic trie, such that the atomic operations
// from [previousLastAcceptedHeight+1] to the [lastAcceptedHeight] set by state sync
// will not have been executed on shared memory.
func (a *AtomicBackend) MarkApplyToSharedMemoryCursor(previousLastAcceptedHeight uint64) error {
	_ = "STUB: not implemented"
	// Set the cursor to [previousLastAcceptedHeight+1] so that we begin the iteration at the
	// first item that has not been applied to shared memory.
	return nil
}

func (a *AtomicBackend) GetVerifiedAtomicState(blockHash common.Hash) (*atomicState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAtomicRootAt returns the atomic trie root for a block that is either:
// - the last accepted block
// - a block that has been verified but not accepted or rejected yet.
// If [blockHash] is neither of the above, an error is returned.
func (a *AtomicBackend) getAtomicRootAt(blockHash common.Hash) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// SetLastAccepted is used after state-sync to update the last accepted block hash.
func (a *AtomicBackend) SetLastAccepted(lastAcceptedHash common.Hash) {
	_ = "STUB: not implemented"
	return
}

// InsertTxs calculates the root of the atomic trie that would
// result from applying [txs] to the atomic trie, starting at the state
// corresponding to previously verified block [parentHash].
// If [blockHash] is provided, the modified atomic trie is pinned in memory
// and it's the caller's responsibility to call either Accept or Reject on
// the AtomicState which can be retreived from GetVerifiedAtomicState to commit the
// changes or abort them and free memory.
func (a *AtomicBackend) InsertTxs(blockHash common.Hash, blockHeight uint64, parentHash common.Hash, txs []*atomic.Tx) (common.Hash, error) {
	_ = "STUB: not implemented"
	// access the atomic trie at the parent block
	return *new(common.Hash), nil
}

// Insert the operations into the atomic trie
//
// Note: The atomic trie canonically contains the duplicate operations from
// any bonus blocks.

// If block hash is not provided, we do not pin the atomic state in memory and can return early

// get the new root and pin the atomic trie changes in memory.

// track this block so further blocks can be inserted on top
// of this block

// IsBonus returns true if the block for atomicState is a bonus block
func (a *AtomicBackend) IsBonus(blockHeight uint64, blockHash common.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *AtomicBackend) AtomicTrie() *AtomicTrie { _ = "STUB: not implemented"; return nil }

// mergeAtomicOps merges atomic requests represented by [txs]
// to the [output] map, depending on whether [chainID] is present in the map.
func mergeAtomicOps(txs []*atomic.Tx) (map[ids.ID]*avalancheatomic.Requests, error) {
	_ = "STUB: not implemented"

	// txs should be stored in order of txID to ensure consistency
	// with txs initialized from the txID index.
	return nil, nil
}

// mergeAtomicOps merges atomic ops for [chainID] represented by [requests]
// to the [output] map provided.
func mergeAtomicOpsToMap(output map[ids.ID]*avalancheatomic.Requests, chainID ids.ID, requests *avalancheatomic.Requests) {
	_ = "STUB: not implemented"
	return
}

// AddBonusBlock adds a bonus block to the atomic backend
func (a *AtomicBackend) AddBonusBlock(height uint64, blockID ids.ID) {
	_ = "STUB: not implemented"
	return
}
