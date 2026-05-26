// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pebbledb

import (
	"github.com/cockroachdb/pebble"

	"github.com/ava-labs/avalanchego/database"
)

var _ database.Batch = (*batch)(nil)

// Not safe for concurrent use.
type batch struct {
	batch *pebble.Batch
	db    *Database
	size  int

	// True iff [batch] has been written to the database
	// since the last time [Reset] was called.
	written bool
}

func (db *Database) NewBatch() database.Batch {
	_ = "STUB: not implemented"
	return *new(database.Batch)
}

func (b *batch) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b *batch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (b *batch) Size() int {
	_ = "STUB: not implemented"

	// Assumes [b.db.lock] is not held.
	return 0
}

func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

// Committing to a closed database makes pebble panic
// so make sure [b.db] isn't closed.

// pebble doesn't support writing a batch twice so we have to clone the
// batch before writing it.

func (b *batch) Reset() { _ = "STUB: not implemented"; return }

func (b *batch) Replay(w database.KeyValueWriterDeleter) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }
