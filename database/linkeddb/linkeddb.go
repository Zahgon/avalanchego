// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package linkeddb

import (
	"sync"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
)

const (
	defaultCacheSize = 1024
)

var (
	headKey = []byte{0x01}

	_ LinkedDB          = (*linkedDB)(nil)
	_ database.Iterator = (*iterator)(nil)
)

// LinkedDB provides a key value interface while allowing iteration.
type LinkedDB interface {
	database.KeyValueReaderWriterDeleter

	IsEmpty() (bool, error)
	HeadKey() ([]byte, error)
	Head() (key []byte, value []byte, err error)

	NewIterator() database.Iterator
	NewIteratorWithStart(start []byte) database.Iterator
}

type linkedDB struct {
	// lock ensure that this data structure handles its thread safety correctly.
	lock sync.RWMutex

	cacheLock sync.Mutex
	// these variables provide caching for the head key.
	headKeyIsSynced, headKeyExists, headKeyIsUpdated, updatedHeadKeyExists bool
	headKey, updatedHeadKey                                                []byte
	// these variables provide caching for the nodes.
	nodeCache    cache.Cacher[string, *node] // key -> *node
	updatedNodes map[string]*node

	// db is the underlying database that this list is stored in.
	db database.Database
	// batch writes to [db] atomically.
	batch database.Batch
}

type node struct {
	Value       []byte `serialize:"true"`
	HasNext     bool   `serialize:"true"`
	Next        []byte `serialize:"true"`
	HasPrevious bool   `serialize:"true"`
	Previous    []byte `serialize:"true"`
}

func New(db database.Database, cacheSize int) LinkedDB {
	_ = "STUB: not implemented"
	return *new(LinkedDB)
}

func NewDefault(db database.Database) LinkedDB { _ = "STUB: not implemented"; return *new(LinkedDB) }

func (ldb *linkedDB) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (ldb *linkedDB) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ldb *linkedDB) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

// If the key already has a node in the list, update that node.

// The key isn't currently in the list, so we should add it as the head.
// Note we will copy the key so it's safe to store references to it.

// The list currently has a head, so we need to update the old head.

func (ldb *linkedDB) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// We're trying to delete this node.

// We aren't modifying the head.

// We aren't modifying the tail.

// This is the only node, so we don't have a head anymore.

// The next node will be the new head.

func (ldb *linkedDB) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (ldb *linkedDB) HeadKey() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ldb *linkedDB) Head() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// This iterator does not guarantee that keys are returned in lexicographic
// order.
func (ldb *linkedDB) NewIterator() database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// NewIteratorWithStart returns an iterator that starts at [start].
// This iterator does not guarantee that keys are returned in lexicographic
// order.
// If [start] is not in the list, starts iterating from the list head.
func (ldb *linkedDB) NewIteratorWithStart(start []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// If the start key isn't present, start from the head

func (ldb *linkedDB) getHeadKey() ([]byte, error) {
	_ = "STUB: not implemented"
	// If the ldb read lock is held, then there needs to be additional
	// synchronization here to avoid racy behavior.
	return nil, nil
}

func (ldb *linkedDB) putHeadKey(key []byte) error { _ = "STUB: not implemented"; return nil }

func (ldb *linkedDB) deleteHeadKey() error { _ = "STUB: not implemented"; return nil }

func (ldb *linkedDB) getNode(key []byte) (node, error) {
	_ = "STUB: not implemented"
	// If the ldb read lock is held, then there needs to be additional
	// synchronization here to avoid racy behavior.
	return *new(node), nil
}

func (ldb *linkedDB) putNode(key []byte, n node) error { _ = "STUB: not implemented"; return nil }

func (ldb *linkedDB) deleteNode(key []byte) error { _ = "STUB: not implemented"; return nil }

func (ldb *linkedDB) resetBatch() { _ = "STUB: not implemented"; return }

func (ldb *linkedDB) writeBatch() error { _ = "STUB: not implemented"; return nil }

type iterator struct {
	ldb                    *linkedDB
	initialized, exhausted bool
	key, value, nextKey    []byte
	err                    error
}

func (it *iterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator has been exhausted, there is no next value.
	return false
}

// If the iterator was not yet initialized, do it now.

func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (it *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (*iterator) Release() { _ = "STUB: not implemented"; return }

func nodeKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }
