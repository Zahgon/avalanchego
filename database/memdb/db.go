// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package memdb

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/database"
)

const (
	// Name is the name of this database for database switches
	Name = "memdb"

	// DefaultSize is the default initial size of the memory database
	DefaultSize = 1024
)

var (
	_ database.Database = (*Database)(nil)
	_ database.Batch    = (*batch)(nil)
	_ database.Iterator = (*iterator)(nil)
)

// Database is an ephemeral key-value store that implements the Database
// interface.
type Database struct {
	lock sync.RWMutex
	db   map[string][]byte
}

// New returns a map with the Database interface methods implemented.
func New() *Database { _ = "STUB: not implemented"; return nil }

// NewWithSize returns a map pre-allocated to the provided size with the
// Database interface methods implemented.
func NewWithSize(size int) *Database { _ = "STUB: not implemented"; return nil }

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) isClosed() bool { _ = "STUB: not implemented"; return false }

func (db *Database) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db *Database) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *Database) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

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

func (db *Database) Compact(_, _ []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type batch struct {
	database.BatchOps

	db *Database
}

func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

func (b *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }

type iterator struct {
	db          *Database
	initialized bool
	keys        []string
	values      [][]byte
	err         error
}

func (it *iterator) Next() bool {
	_ = "STUB: not implemented"
	// Short-circuit and set an error if the underlying database has been closed.
	return false
}

// If the iterator was not yet initialized, do it now

// Iterator already initialize, advance it

func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (it *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Release() { _ = "STUB: not implemented"; return }
