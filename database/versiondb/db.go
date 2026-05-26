// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package versiondb

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/database"
)

var (
	_ database.Database = (*Database)(nil)
	_ Commitable        = (*Database)(nil)
	_ database.Batch    = (*batch)(nil)
	_ database.Iterator = (*iterator)(nil)
)

// Commitable defines the interface that specifies that something may be
// committed.
type Commitable interface {
	// Commit writes all the queued operations to the underlying data structure.
	Commit() error
}

// Database implements the Database interface by living on top of another
// database, writing changes to the underlying database only when commit is
// called.
type Database struct {
	lock  sync.RWMutex
	mem   map[string]valueDelete
	db    database.Database
	batch database.Batch
}

type valueDelete struct {
	value  []byte
	delete bool
}

// New returns a new versioned database
func New(db database.Database) *Database { _ = "STUB: not implemented"; return nil }

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

func (db *Database) NewIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// Keys need to be in sorted order

func (db *Database) Compact(start, limit []byte) error { _ = "STUB: not implemented"; return nil }

// SetDatabase changes the underlying database to the specified database
func (db *Database) SetDatabase(newDB database.Database) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDatabase returns the underlying database
func (db *Database) GetDatabase() database.Database {
	_ = "STUB: not implemented"
	return *new(database.Database)
}

// Commit writes all the operations of this database to the underlying database
func (db *Database) Commit() error { _ = "STUB: not implemented"; return nil }

// Abort all changes to the underlying database
func (db *Database) Abort() { _ = "STUB: not implemented"; return }

func (db *Database) abort() {
	_ = "STUB: not implemented"

	// CommitBatch returns a batch that contains all uncommitted puts/deletes.
	// Calling Write() on the returned batch causes the puts/deletes to be
	// written to the underlying database. The returned batch should be written before
	// future calls to this DB unless the batch will never be written.
	return
}

func (db *Database) CommitBatch() (database.Batch, error) {
	_ = "STUB: not implemented"
	return *new(database.Batch), nil
}

// Put all of the puts/deletes in memory into db.batch
// and return the batch
func (db *Database) commitBatch() (database.Batch, error) {
	_ = "STUB: not implemented"
	return *new(database.Batch), nil
}

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) isClosed() bool { _ = "STUB: not implemented"; return false }

func (db *Database) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type batch struct {
	database.BatchOps

	db *Database
}

func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

func (b *batch) Inner() database.Batch {
	_ = "STUB: not implemented"

	// iterator walks over both the in memory database and the underlying database
	// at the same time.
	return *new(database.Batch)
}

type iterator struct {
	db *Database
	database.Iterator

	key, value []byte
	err        error

	keys   []string
	values []valueDelete

	initialized, exhausted bool
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted. We must pay careful attention to set the proper values
// based on if the in memory db or the underlying db should be read next
func (it *iterator) Next() bool {
	_ = "STUB: not implemented"
	// Short-circuit and set an error if the underlying database has been closed.
	return false
}

func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (it *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Release() { _ = "STUB: not implemented"; return }
