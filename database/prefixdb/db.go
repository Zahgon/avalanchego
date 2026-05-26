// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package prefixdb

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils"
)

var (
	_ database.Database = (*Database)(nil)
	_ database.Batch    = (*batch)(nil)
	_ database.Iterator = (*iterator)(nil)
)

// Database partitions a database into a sub-database by prefixing all keys with
// a unique value.
type Database struct {
	// All keys in this db begin with this byte slice
	dbPrefix []byte
	// Lexically one greater than dbPrefix, defining the end of this db's key range
	dbLimit    []byte
	bufferPool *utils.BytesPool

	// lock needs to be held during Close to guarantee db will not be set to nil
	// concurrently with another operation. All other operations can hold RLock.
	lock sync.RWMutex
	// The underlying storage
	db     database.Database
	closed bool
}

func newDB(prefix []byte, db database.Database) *Database { _ = "STUB: not implemented"; return nil }

func incrementByteSlice(orig []byte) []byte { _ = "STUB: not implemented"; return nil }

// New returns a new prefixed database
func New(prefix []byte, db database.Database) *Database { _ = "STUB: not implemented"; return nil }

// NewNested returns a new prefixed database without attempting to compress
// prefixes.
func NewNested(prefix []byte, db database.Database) *Database {
	_ = "STUB: not implemented"
	return nil
}

func MakePrefix(prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

func JoinPrefixes(firstPrefix, secondPrefix []byte) []byte { _ = "STUB: not implemented"; return nil }

func PrefixKey(prefix, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (db *Database) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db *Database) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *Database) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) NewBatch() database.Batch {
	_ = "STUB: not implemented"
	return *new(database.Batch)
}

func (db *Database) NewIterator() database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *Database) NewIteratorWithStart(start []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *Database) NewIteratorWithPrefix(prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// Assumes it is safe to modify the arguments to db.db.NewIteratorWithStartAndPrefix after it returns.
// It is safe to modify [start] and [prefix] after this method returns.
func (db *Database) NewIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *Database) Compact(start, limit []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) isClosed() bool { _ = "STUB: not implemented"; return false }

func (db *Database) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a copy of [key], prepended with this db's prefix.
// The returned slice should be put back in the pool when it's done being used.
func (db *Database) prefix(key []byte) *[]byte { _ = "STUB: not implemented"; return nil }

// Batch of database operations
type batch struct {
	database.Batch
	db *Database

	// Each key is prepended with the database's prefix.
	// Each byte slice underlying a key should be returned to the pool
	// when this batch is reset.
	ops []batchOp
}

type batchOp struct {
	Key    *[]byte
	Value  []byte
	Delete bool
}

func (b *batch) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b *batch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// Write flushes any accumulated data to the memory database.
func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

// Reset resets the batch for reuse.
func (b *batch) Reset() {
	_ = "STUB: not implemented"
	// Return the byte buffers underneath each key back to the pool.
	// Don't return the byte buffers underneath each value back to the pool
	// because we assume in batch.Replay that it's not safe to modify the
	// value argument to w.Put.
	return
}

// Clear b.writes

// Replay the batch contents.
func (b *batch) Replay(w database.KeyValueWriterDeleter) error {
	_ = "STUB: not implemented"
	return nil
}

type iterator struct {
	database.Iterator
	db *Database

	key, val []byte
	err      error
}

// Next calls the inner iterators Next() function and strips the keys prefix
func (it *iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Value() []byte {
	_ = "STUB: not implemented"

	// Error returns [database.ErrClosed] if the underlying db was closed
	// otherwise it returns the normal iterator error.
	return nil
}

func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }
