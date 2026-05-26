// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package corruptabledb

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var (
	_ database.Database = (*Database)(nil)
	_ database.Batch    = (*batch)(nil)
)

// CorruptableDB is a wrapper around Database
// it prevents any future calls in case of a corruption occurs
type Database struct {
	database.Database

	log logging.Logger
	// initialError stores the error other than "not found" or "closed" while
	// performing a db operation. If not nil, Has, Get, Put, Delete and batch
	// writes will fail with initialError.
	errorLock    sync.RWMutex
	initialError error
}

// New returns a new prefixed database
func New(db database.Database, log logging.Logger) *Database { _ = "STUB: not implemented"; return nil }

// Has returns if the key is set in the database
func (db *Database) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Get returns the value the key maps to in the database
func (db *Database) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Put sets the value of the provided key to the provided value
func (db *Database) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

// Delete removes the key from the database
func (db *Database) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Compact(start []byte, limit []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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

func (db *Database) corrupted() error { _ = "STUB: not implemented"; return nil }

func (db *Database) handleError(err error) error { _ = "STUB: not implemented"; return nil }

// If we get an error other than "not found" or "closed", disallow future
// database operations to avoid possible corruption

// Set the initial error to the first unexpected error. Don't call
// corrupted() here since it would deadlock.

// batch is a wrapper around the batch to contain sizes.
type batch struct {
	database.Batch
	db *Database
}

// Write flushes any accumulated data to disk.
func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

type iterator struct {
	database.Iterator
	db *Database
}

func (it *iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }
