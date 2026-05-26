// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pebbledb

import (
	"context"
	"errors"
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	Name = "pebbledb"

	// pebbleByteOverHead is the number of bytes of constant overhead that
	// should be added to a batch size per operation.
	pebbleByteOverHead = 8

	defaultCacheSize = 512 * units.MiB
)

var (
	_ database.Database = (*Database)(nil)

	errInvalidOperation = errors.New("invalid operation")

	DefaultConfig = Config{
		CacheSize:                   defaultCacheSize,
		BytesPerSync:                512 * units.KiB,
		WALBytesPerSync:             0, // Default to no background syncing.
		MemTableStopWritesThreshold: 8,
		MemTableSize:                defaultCacheSize / 4,
		MaxOpenFiles:                4096,
		MaxConcurrentCompactions:    1,
		Sync:                        true,
	}
)

type Database struct {
	lock          sync.RWMutex
	pebbleDB      *pebble.DB
	closed        bool
	openIterators set.Set[*iter]
	writeOptions  *pebble.WriteOptions
}

type Config struct {
	CacheSize                   int64  `json:"cacheSize"`
	BytesPerSync                int    `json:"bytesPerSync"`
	WALBytesPerSync             int    `json:"walBytesPerSync"` // 0 means no background syncing
	MemTableStopWritesThreshold int    `json:"memTableStopWritesThreshold"`
	MemTableSize                uint64 `json:"memTableSize"`
	MaxOpenFiles                int    `json:"maxOpenFiles"`
	MaxConcurrentCompactions    int    `json:"maxConcurrentCompactions"`
	Sync                        bool   `json:"sync"`
}

// TODO: Add metrics
func New(file string, configBytes []byte, log logging.Logger, _ prometheus.Registerer) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

// Disable seek compaction

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) HealthCheck(_ context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Database) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db *Database) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *Database) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Compact(start []byte, end []byte) error { _ = "STUB: not implemented"; return nil }

// The database.Database spec treats a nil [limit] as a key after all
// keys but pebble treats a nil [limit] as a key before all keys in
// Compact. Use the greatest key in the database as the [limit] to get
// the desired behavior.

// The database is empty.

// pebble requires [start] < [end]

/* parallelize */

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

// Converts a pebble-specific error to its Avalanche equivalent, if applicable.
func updateError(err error) error { _ = "STUB: not implemented"; return nil }

func keyRange(start, prefix []byte) *pebble.IterOptions { _ = "STUB: not implemented"; return nil }

// Returns an upper bound that stops after all keys with the given [prefix].
// Assumes the Database uses bytes.Compare for key comparison and not a custom
// comparer.
func prefixToUpperBound(prefix []byte) []byte { _ = "STUB: not implemented"; return nil }
