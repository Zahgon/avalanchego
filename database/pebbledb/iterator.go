// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pebbledb

import (
	"errors"
	"sync"

	"github.com/cockroachdb/pebble"

	"github.com/ava-labs/avalanchego/database"
)

var (
	_ database.Iterator = (*iter)(nil)

	errCouldNotGetValue = errors.New("could not get iterator value")
)

type iter struct {
	// [lock] ensures that only one goroutine can access [iter] at a time.
	// Note that [Database.Close] calls [iter.Release] so we need [lock] to ensure
	// that the user and [Database.Close] don't execute [iter.Release] concurrently.
	// Invariant: [Database.lock] is never grabbed while holding [lock].
	lock sync.Mutex

	db   *Database
	iter *pebble.Iterator

	initialized bool
	closed      bool
	err         error

	hasNext bool
	nextKey []byte
	nextVal []byte
}

// Must not be called with [db.lock] held.
func (it *iter) Next() bool { _ = "STUB: not implemented"; return false }

func (it *iter) Error() error { _ = "STUB: not implemented"; return nil }

func (it *iter) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iter) Value() []byte { _ = "STUB: not implemented"; return nil }

func (it *iter) Release() { _ = "STUB: not implemented"; return }

// Assumes [it.lock] and [it.db.lock] are held.
func (it *iter) release() { _ = "STUB: not implemented"; return }

// Cloning these values ensures that calling it.Key() or it.Value() after
// releasing the iterator will not segfault.

// Remove the iterator from the list of open iterators.
