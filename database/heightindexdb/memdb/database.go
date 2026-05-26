// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package memdb

import (
	"sync"

	"github.com/ava-labs/avalanchego/database"
)

var _ database.HeightIndex = (*Database)(nil)

// Database is an in-memory implementation of database.HeightIndex
type Database struct {
	mu     sync.RWMutex
	data   map[uint64][]byte
	closed bool
}

func (db *Database) Put(height uint64, data []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Get(height uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *Database) Has(height uint64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db *Database) Sync(_, _ uint64) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }
