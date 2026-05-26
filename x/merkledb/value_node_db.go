// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"errors"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils"
)

var (
	_ database.Iterator = (*iterator)(nil)

	errNodeMissingValue = errors.New("valueNodeDB contains node without a value")
)

type valueNodeDB struct {
	bufferPool *utils.BytesPool

	// The underlying storage.
	// Keys written to [baseDB] are prefixed with [valueNodePrefix].
	baseDB database.Database

	// If a value is nil, the corresponding key isn't in the trie.
	// Paths in [nodeCache] aren't prefixed with [valueNodePrefix].
	nodeCache cache.Cacher[Key, *node]
	metrics   metrics

	hasher Hasher

	closed utils.Atomic[bool]
}

func newValueNodeDB(
	db database.Database,
	bufferPool *utils.BytesPool,
	metrics metrics,
	cacheSize int,
	hasher Hasher,
) *valueNodeDB {
	_ = "STUB: not implemented"
	return nil
}

func (db *valueNodeDB) Write(batch database.KeyValueWriterDeleter, key Key, n *node) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *valueNodeDB) newIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *valueNodeDB) Close() { _ = "STUB: not implemented"; return }

func (db *valueNodeDB) Get(key Key) (*node, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *valueNodeDB) Clear() error { _ = "STUB: not implemented"; return nil }

type iterator struct {
	db       *valueNodeDB
	nodeIter database.Iterator
	key      []byte
	value    []byte
	err      error
}

func (i *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (i *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (i *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (i *iterator) Next() bool { _ = "STUB: not implemented"; return false }

// We are discarding the other bytes from the node, so we avoid copying
// the value here.

func (i *iterator) Release() { _ = "STUB: not implemented"; return }
