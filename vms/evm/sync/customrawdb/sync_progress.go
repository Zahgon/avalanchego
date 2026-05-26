// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customrawdb

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/utils/wrappers"
)

var (
	// syncRootKey indicates the root of the main account trie currently being synced.
	syncRootKey = []byte("sync_root")
	// syncStorageTriesPrefix + trie root + account hash indicates a storage trie must be fetched for the account.
	syncStorageTriesPrefix = []byte("sync_storage")
	// syncSegmentsPrefix + trie root + 32-byte start key indicates the trie at root has a segment starting at the specified key.
	syncSegmentsPrefix = []byte("sync_segments")
	// CodeToFetchPrefix + code hash -> empty value tracks the outstanding code hashes we need to fetch.
	CodeToFetchPrefix = []byte("CP")

	// === State sync progress key lengths ===
	syncStorageTriesKeyLength = len(syncStorageTriesPrefix) + 2*common.HashLength
	syncSegmentsKeyLength     = len(syncSegmentsPrefix) + 2*common.HashLength
	codeToFetchKeyLength      = len(CodeToFetchPrefix) + common.HashLength

	// === State sync metadata ===
	syncPerformedPrefix = []byte("sync_performed")
	// syncPerformedKeyLength is the length of the key for the sync performed metadata key,
	// and is equal to [syncPerformedPrefix] + block number as uint64.
	syncPerformedKeyLength = len(syncPerformedPrefix) + wrappers.LongLen
)

// ReadSyncRoot reads the root corresponding to the main trie of an in-progress
// sync and returns common.Hash{} if no in-progress sync was found.
func ReadSyncRoot(db ethdb.KeyValueReader) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// WriteSyncRoot writes root as the root of the main trie of the in-progress sync.
func WriteSyncRoot(db ethdb.KeyValueWriter, root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteCodeToFetch adds a marker that we need to fetch the code for `hash`.
func WriteCodeToFetch(db ethdb.KeyValueWriter, codeHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCodeToFetch removes the marker that the code corresponding to `hash` needs to be fetched.
func DeleteCodeToFetch(db ethdb.KeyValueWriter, codeHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// NewCodeToFetchIterator returns a KeyLength iterator over all code
// hashes that are pending syncing. It is the caller's responsibility to
// parse the key and call Release on the returned iterator.
func NewCodeToFetchIterator(db ethdb.Iteratee) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

func codeToFetchKey(codeHash common.Hash) []byte { _ = "STUB: not implemented"; return nil }

// NewSyncSegmentsIterator returns a KeyLength iterator over all trie segments
// added for root. It is the caller's responsibility to parse the key and call
// Release on the returned iterator.
func NewSyncSegmentsIterator(db ethdb.Iteratee, root common.Hash) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

// WriteSyncSegment adds a trie segment for root at the given start position.
func WriteSyncSegment(db ethdb.KeyValueWriter, root common.Hash, start common.Hash) error {
	_ = "STUB: not implemented"
	// packs root and account into a key for storage in db.
	return nil
}

// ClearSyncSegments removes segment markers for root from db
func ClearSyncSegments(db ethdb.KeyValueStore, root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// ClearAllSyncSegments removes all segment markers from db
func ClearAllSyncSegments(db ethdb.KeyValueStore) error { _ = "STUB: not implemented"; return nil }

// ParseSyncSegmentKey returns the root and start position for a trie segment
// key returned from NewSyncSegmentsIterator.
func ParseSyncSegmentKey(keyBytes []byte) (common.Hash, []byte) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// skip prefix

// NewSyncStorageTriesIterator returns a KeyLength iterator over all storage tries
// added for syncing (beginning at seek). It is the caller's responsibility to parse
// the key and call Release on the returned iterator.
func NewSyncStorageTriesIterator(db ethdb.Iteratee, seek []byte) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

// WriteSyncStorageTrie adds a storage trie for account (with the given root) to be synced.
func WriteSyncStorageTrie(db ethdb.KeyValueWriter, root common.Hash, account common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// ClearSyncStorageTrie removes all storage trie accounts (with the given root) from db.
// Intended for use when the trie with root has completed syncing.
func ClearSyncStorageTrie(db ethdb.KeyValueStore, root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// ClearAllSyncStorageTries removes all storage tries added for syncing from db
func ClearAllSyncStorageTries(db ethdb.KeyValueStore) error { _ = "STUB: not implemented"; return nil }

// ParseSyncStorageTrieKey returns the root and account for a storage trie
// key returned from NewSyncStorageTriesIterator. It assumes the key has the
// `syncStorageTriesPrefix` followed by a 32-byte root and 32-byte account hash,
// and panics if the key is shorter than len(syncStorageTriesPrefix)+2*common.HashLength.
func ParseSyncStorageTrieKey(keyBytes []byte) (common.Hash, common.Hash) {
	_ = "STUB: not implemented"
	return *new(common.Hash), *new(common.Hash)
}

// skip prefix

// WriteSyncPerformed logs an entry in `db` indicating the VM state synced to `blockNumber`.
func WriteSyncPerformed(db ethdb.KeyValueWriter, blockNumber uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// GetLatestSyncPerformed returns the latest block number state synced performed to.
func GetLatestSyncPerformed(db ethdb.Iteratee) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// newSyncPerformedIterator returns an iterator over all block numbers the VM
// has state synced to.
func newSyncPerformedIterator(db ethdb.Iteratee) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

// parseSyncPerformedKey returns the block number from keys returned by
// NewSyncPerformedIterator. It panics if the key is shorter than `syncPerformedKeyLength`.
func parseSyncPerformedKey(key []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// clearPrefix removes all keys in db that begin with prefix and match an
// expected key length. `keyLen` must include the length of the prefix.
func clearPrefix(db ethdb.KeyValueStore, prefix []byte, keyLen int) error {
	_ = "STUB: not implemented"
	return nil
}
