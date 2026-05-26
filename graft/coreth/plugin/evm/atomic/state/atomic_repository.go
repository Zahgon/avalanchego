// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	repoCommitSizeCap = 10 * units.MiB
)

var (
	atomicTxIDDBPrefix         = []byte("atomicTxDB")
	atomicHeightTxDBPrefix     = []byte("atomicHeightTxDB")
	atomicRepoMetadataDBPrefix = []byte("atomicRepoMetadataDB")
	atomicTrieStoragePrefix    = []byte("atomicTrieDB")
	atomicTrieMetaDBPrefix     = []byte("atomicTrieMetaDB")

	appliedSharedMemoryCursorKey = []byte("atomicTrieLastAppliedToSharedMemory")
	maxIndexedHeightKey          = []byte("maxIndexedAtomicTxHeight")
	// Historically used to track the completion of a migration
	// bonusBlocksRepairedKey     = []byte("bonusBlocksRepaired")
)

// AtomicRepository manages the database interactions for atomic operations.
type AtomicRepository struct {
	// [acceptedAtomicTxDB] maintains an index of [txID] => [height]+[atomic tx] for all accepted atomic txs.
	acceptedAtomicTxDB database.Database

	// [acceptedAtomicTxByHeightDB] maintains an index of [height] => [atomic txs] for all accepted block heights.
	acceptedAtomicTxByHeightDB database.Database

	// [atomicRepoMetadataDB] maintains a single key-value pair which tracks the height up to which the atomic repository
	// has indexed.
	atomicRepoMetadataDB database.Database

	metadataDB database.Database // Underlying database containing the atomic trie metadata

	atomicTrieStorage database.Database // Raw database storage for atomic trie (unwrapped)

	// [db] is used to commit to the underlying versiondb.
	db *versiondb.Database

	// Use this codec for serializing
	codec codec.Manager
}

func NewAtomicTxRepository(
	db *versiondb.Database, codec codec.Manager, lastAcceptedHeight uint64,
) (*AtomicRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initializeHeightIndex initializes the atomic repository and takes care of any required migration from the previous database
// format which did not have a height -> txs index.
func (a *AtomicRepository) initializeHeightIndex(lastAcceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// [lastTxID] will be initialized to the last transaction that we indexed
// if we are part way through a migration.

// unexpected value in the database

// partially initialized

// already initialized

// unexpected value in the database

// Iterate from [lastTxID] to complete the re-index -> generating an index
// from height to a slice of transactions accepted at that height

// Keep track of the size of the currently pending writes

// iter.Value() consists of [height packed as uint64] + [tx serialized as packed []byte]

// Get the tx iter is pointing to, len(txs) == 1 is expected here.

// Check if there are already transactions at [height], to ensure that we
// add [txs] to the already indexed transactions at [height] instead of
// overwriting them.

// call commitFn to write to underlying DB if we have reached
// [commitSizeCap]

// Periodically log progress

// Updated the value stored [maxIndexedHeightKey] to be the lastAcceptedHeight

// GetIndexHeight returns the last height that was indexed by the atomic repository
func (a *AtomicRepository) GetIndexHeight() (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetByTxID queries [acceptedAtomicTxDB] for the [txID], parses a [*atomic.Tx] object
// if an entry is found, and returns it with the block height the atomic tx it
// represents was accepted on, along with an optional error.
func (a *AtomicRepository) GetByTxID(txID ids.ID) (*atomic.Tx, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// value is stored as [height]+[tx bytes], decompose with a packer.

// GetByHeight returns all atomic txs processed on block at [height].
// Returns [database.ErrNotFound] if there are no atomic transactions indexed at [height].
// Note: if [height] is below the last accepted height, then this means that there were
// no atomic transactions in the block accepted at [height].
// If [height] is greater than the last accepted height, then this will always return
// [database.ErrNotFound]
func (a *AtomicRepository) GetByHeight(height uint64) ([]*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AtomicRepository) getByHeightBytes(heightBytes []byte) ([]*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Write updates indexes maintained on atomic txs, so they can be queried
// by txID or height. This method must be called only once per height,
// and [txs] must include all atomic txs for the block accepted at the
// corresponding height.
func (a *AtomicRepository) Write(height uint64, txs []*atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteBonus is similar to Write, except the [txID] => [height] is not
// overwritten if already exists.
func (a *AtomicRepository) WriteBonus(height uint64, txs []*atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AtomicRepository) write(height uint64, txs []*atomic.Tx, bonus bool) error {
	_ = "STUB: not implemented"

	// txs should be stored in order of txID to ensure consistency
	// with txs initialized from the txID index.
	return nil
}

// Skip adding an entry to the height index if [txs] is empty.

// avoid overwriting existing value if [bonus] is true

// no existing value to overwrite, proceed as normal

// unexpected error

// Update the index height regardless of if any atomic transactions
// were present at [height].

// indexTxByID writes [tx] into the [acceptedAtomicTxDB] stored as
// [height] + [tx bytes]
func (a *AtomicRepository) indexTxByID(heightBytes []byte, tx *atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// map txID => [height]+[tx bytes]

// indexTxsAtHeight adds [height] -> [txs] to the [acceptedAtomicTxByHeightDB]
func (a *AtomicRepository) indexTxsAtHeight(heightBytes []byte, txs []*atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// appendTxToHeightIndex retrieves the transactions stored at [heightBytes] and appends
// [tx] to the slice of transactions stored there.
// This function is used while initializing the atomic repository to re-index the atomic transactions
// by txID into the height -> txs index.
func (a *AtomicRepository) appendTxToHeightIndex(heightBytes []byte, tx *atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// Iterate over the existing transactions to ensure we do not add a
// duplicate to the index.

// IterateByHeight returns an iterator beginning at [height].
// Note [height] must be greater than 0 since we assume there are no
// atomic txs in genesis.
func (a *AtomicRepository) IterateByHeight(height uint64) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}
