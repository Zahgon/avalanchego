// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils"
)

// Holds intermediate nodes. That is, those without values.
// Changes to this database aren't written to [baseDB] until
// they're evicted from the [nodeCache] or Flush is called.
type intermediateNodeDB struct {
	bufferPool *utils.BytesPool

	// The underlying storage.
	// Keys written to [baseDB] are prefixed with [intermediateNodePrefix].
	baseDB database.Database

	// The write buffer contains nodes that have been changed but have not been written to disk.
	// Note that a call to Put may cause a node to be evicted
	// from the cache, which will call [OnEviction].
	// A non-nil error returned from Put is considered fatal.
	// Keys in [nodeCache] aren't prefixed with [intermediateNodePrefix].
	writeBuffer onEvictCache[Key, *node]

	// If a value is nil, the corresponding key isn't in the trie.
	nodeCache cache.Cacher[Key, *node]

	// the number of bytes to evict during an eviction batch
	evictionBatchSize int
	metrics           metrics
	tokenSize         int
	hasher            Hasher
}

func newIntermediateNodeDB(
	db database.Database,
	bufferPool *utils.BytesPool,
	metrics metrics,
	cacheSize int,
	writeBufferSize int,
	evictionBatchSize int,
	tokenSize int,
	hasher Hasher,
) *intermediateNodeDB {
	_ = "STUB: not implemented"
	return nil
}

// A non-nil error is considered fatal and closes [db.baseDB].
func (db *intermediateNodeDB) onEviction(key Key, n *node) error {
	_ = "STUB: not implemented"
	return nil
}

// Evict the oldest [evictionBatchSize] nodes from the cache
// and write them to disk. We write a batch of them, rather than
// just [n], so that we don't immediately evict and write another
// node, because each time this method is called we do a disk write.
// Evicts a total number of bytes, rather than a number of nodes

// The cache is empty.

func (db *intermediateNodeDB) addToBatch(b database.KeyValueWriterDeleter, key Key, n *node) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *intermediateNodeDB) Get(key Key) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// constructDBKey returns a key that can be used in [db.baseDB].
// We need to be able to differentiate between two keys of equal
// byte length but different bit length, so we add padding to differentiate.
// Additionally, we add a prefix indicating it is part of the intermediateNodeDB.
func (db *intermediateNodeDB) constructDBKey(key Key) *[]byte {
	_ = "STUB: not implemented"
	return nil

	// For tokens of size byte, no padding is needed since byte
	// length == token length
}

// add prefix
// add key
// add padding

func (db *intermediateNodeDB) Put(key Key, n *node) error { _ = "STUB: not implemented"; return nil }

func (db *intermediateNodeDB) Flush() error { _ = "STUB: not implemented"; return nil }

func (db *intermediateNodeDB) Delete(key Key) error { _ = "STUB: not implemented"; return nil }

func (db *intermediateNodeDB) Clear() error { _ = "STUB: not implemented"; return nil }

// Reset the buffer. Note we don't flush because that would cause us to
// persist intermediate nodes we're about to delete.
