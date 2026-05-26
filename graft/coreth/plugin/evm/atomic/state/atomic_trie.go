// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/utils/wrappers"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
	avalanchedatabase "github.com/ava-labs/avalanchego/database"
)

const (
	TrieKeyLength = wrappers.LongLen + common.HashLength

	atomicTrieMemoryCap = 64 * units.MiB
)

var lastCommittedKey = []byte("atomicTrieLastCommittedBlock")

// AtomicTrie is a trie that is used to store atomic operations.
type AtomicTrie struct {
	commitInterval      uint64                     // commit interval, same as commitHeightInterval by default
	metadataDB          avalanchedatabase.Database // Underlying database containing the atomic trie metadata
	trieDB              *triedb.Database           // Trie database
	lastCommittedRoot   common.Hash                // trie root of the most recent commit
	lastCommittedHeight uint64                     // index height of the most recent commit
	lastAcceptedRoot    common.Hash                // most recent trie root passed to accept trie or the root of the atomic trie on intialization.
	codec               codec.Manager
	memoryCap           common.StorageSize
}

// newAtomicTrie returns a new instance of a atomicTrie with a configurable commitHeightInterval, used in testing.
// Initializes the trie before returning it.
func newAtomicTrie(
	atomicTrieStorage avalanchedatabase.Database, metadataDB avalanchedatabase.Database,
	codec codec.Manager, lastAcceptedHeight uint64, commitHeightInterval uint64,
) (*AtomicTrie, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize to EmptyRootHash if there is no committed root.

// If the last committed height is above the last accepted height, then we fall back to
// the last commit below the last accepted height.

// Allocate 64MB of memory for clean cache

// Initialize lastAcceptedRoot to the last committed root.
// If there were further blocks processed (ahead of the commit interval),
// AtomicBackend will call InsertTrie/AcceptTrie on atomic ops
// for those blocks.

// lastCommittedRootIfExists returns the last committed trie root and height if it exists
// else returns empty common.Hash{} and 0
// returns error only if there are issues with the underlying data store
// or if values present in the database are not as expected
func lastCommittedRootIfExists(db avalanchedatabase.Database) (common.Hash, uint64, error) {
	_ = "STUB: not implemented"
	// read the last committed entry if it exists and set the root hash
	return *new(common.Hash), 0, nil
}

// nearestCommitheight returns the nearest multiple of commitInterval less than or equal to blockNumber
func nearestCommitHeight(blockNumber uint64, commitInterval uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (a *AtomicTrie) OpenTrie(root common.Hash) (*trie.Trie, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Commit calls Commit on the underlying trieDB and updates metadata pointers.
func (a *AtomicTrie) Commit(height uint64, root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AtomicTrie) UpdateTrie(trie *trie.Trie, height uint64, atomicOps map[ids.ID]*avalancheatomic.Requests) error {
	_ = "STUB: not implemented"
	return nil
}

// highly unlikely but possible if atomic.Element
// has a change that is unsupported by the codec

// key is [height]+[blockchainID]

// LastCommitted returns the last committed trie hash and last committed height
func (a *AtomicTrie) LastCommitted() (common.Hash, uint64) {
	_ = "STUB: not implemented"
	return *new(common.Hash), 0
}

// updateLastCommitted adds [height] -> [root] to the index and marks it as the last committed
// root/height pair.
func (a *AtomicTrie) updateLastCommitted(root common.Hash, height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// now save the trie hash against the height it was committed at

// update lastCommittedKey with the current height

// Iterator returns an [atomicTrieIterator] that iterates the trie from the given
// atomic trie root, starting at the specified [cursor].
func (a *AtomicTrie) Iterator(root common.Hash, cursor []byte) (*atomicTrieIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AtomicTrie) TrieDB() *triedb.Database {
	_ = "STUB: not implemented"

	// Root returns hash if it exists at specified height
	// if trie was not committed at provided height, it returns
	// common.Hash{} instead
	return nil
}

func (a *AtomicTrie) Root(height uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// getRoot is a helper function to return the committed atomic trie root hash at [height]
// from [metadataDB].
func getRoot(metadataDB avalanchedatabase.Database, height uint64) (common.Hash, error) {
	_ = "STUB: not implemented"

	// if root is queried at height == 0, return the empty root hash
	// this may occur if peers ask for the most recent state summary
	// and number of accepted blocks is less than the commit interval.
	return *new(common.Hash), nil
}

func (a *AtomicTrie) LastAcceptedRoot() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (a *AtomicTrie) InsertTrie(nodes *trienode.NodeSet, root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// The use of [Cap] in [insertTrie] prevents exceeding the configured memory
// limit (and OOM) in case there is a large backlog of processing (unaccepted) blocks.

// AcceptTrie commits the triedb at [root] if needed and returns true if a commit
// was performed.
func (a *AtomicTrie) AcceptTrie(height uint64, root common.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Because we do not accept the trie at every height, we may need to
		// populate roots at prior commit heights that were skipped.
		nil
}

// Commit this root if we have reached the [commitInterval].

// The following dereferences, if any, the previously inserted root.
// This one can be dereferenced whether it has been:
// - committed, in which case the dereference is a no-op
// - not committted, in which case the current root we are inserting contains
//   references to all the relevant data from the previous root, so the previous
//   root can be dereferenced.

func (a *AtomicTrie) RejectTrie(root common.Hash) error { _ = "STUB: not implemented"; return nil }
