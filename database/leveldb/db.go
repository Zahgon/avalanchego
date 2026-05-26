// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package leveldb

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	// Name is the name of this database for database switches
	Name = "leveldb"

	// DefaultBlockCacheSize is the number of bytes to use for block caching in
	// leveldb.
	DefaultBlockCacheSize = 12 * opt.MiB

	// DefaultWriteBufferSize is the number of bytes to use for buffers in
	// leveldb.
	DefaultWriteBufferSize = 12 * opt.MiB

	// DefaultHandleCap is the number of files descriptors to cap levelDB to
	// use.
	DefaultHandleCap = 1024

	// DefaultBitsPerKey is the number of bits to add to the bloom filter per
	// key.
	DefaultBitsPerKey = 10

	// DefaultMaxManifestFileSize is the default maximum size of a manifest
	// file.
	//
	// This avoids https://github.com/syndtr/goleveldb/issues/413.
	DefaultMaxManifestFileSize = math.MaxInt64

	// DefaultMetricUpdateFrequency is the frequency to poll the LevelDB
	// metrics.
	DefaultMetricUpdateFrequency = 10 * time.Second

	// levelDBByteOverhead is the number of bytes of constant overhead that
	// should be added to a batch size per operation.
	levelDBByteOverhead = 8
)

var (
	_ database.Database = (*Database)(nil)
	_ database.Batch    = (*batch)(nil)
	_ database.Iterator = (*iter)(nil)

	ErrInvalidConfig = errors.New("invalid config")
	ErrCouldNotOpen  = errors.New("could not open")
)

// Database is a persistent key-value store. Apart from basic data storage
// functionality it also supports batch writes and iterating over the keyspace
// in binary-alphabetical order.
type Database struct {
	*leveldb.DB
	// metrics is only initialized and used when [MetricUpdateFrequency] is >= 0
	// in the config
	metrics   metrics
	closed    utils.Atomic[bool]
	closeOnce sync.Once
	// closeCh is closed when Close() is called.
	closeCh chan struct{}
	// closeWg is used to wait for all goroutines created by New() to exit.
	// This avoids racy behavior when Close() is called at the same time as
	// Stats(). See: https://github.com/syndtr/goleveldb/issues/418
	closeWg sync.WaitGroup
}

type config struct {
	// BlockCacheCapacity defines the capacity of the 'sorted table' block caching.
	// Use -1 for zero, this has same effect as specifying NoCacher to BlockCacher.
	//
	// The default value is 12MiB.
	BlockCacheCapacity int `json:"blockCacheCapacity"`
	// BlockSize is the minimum uncompressed size in bytes of each 'sorted table'
	// block.
	//
	// The default value is 4KiB.
	BlockSize int `json:"blockSize"`
	// CompactionExpandLimitFactor limits compaction size after expanded.
	// This will be multiplied by table size limit at compaction target level.
	//
	// The default value is 25.
	CompactionExpandLimitFactor int `json:"compactionExpandLimitFactor"`
	// CompactionGPOverlapsFactor limits overlaps in grandparent (Level + 2)
	// that a single 'sorted table' generates.  This will be multiplied by
	// table size limit at grandparent level.
	//
	// The default value is 10.
	CompactionGPOverlapsFactor int `json:"compactionGPOverlapsFactor"`
	// CompactionL0Trigger defines number of 'sorted table' at level-0 that will
	// trigger compaction.
	//
	// The default value is 4.
	CompactionL0Trigger int `json:"compactionL0Trigger"`
	// CompactionSourceLimitFactor limits compaction source size. This doesn't apply to
	// level-0.
	// This will be multiplied by table size limit at compaction target level.
	//
	// The default value is 1.
	CompactionSourceLimitFactor int `json:"compactionSourceLimitFactor"`
	// CompactionTableSize limits size of 'sorted table' that compaction generates.
	// The limits for each level will be calculated as:
	//   CompactionTableSize * (CompactionTableSizeMultiplier ^ Level)
	// The multiplier for each level can also fine-tuned using CompactionTableSizeMultiplierPerLevel.
	//
	// The default value is 2MiB.
	CompactionTableSize int `json:"compactionTableSize"`
	// CompactionTableSizeMultiplier defines multiplier for CompactionTableSize.
	//
	// The default value is 1.
	CompactionTableSizeMultiplier float64 `json:"compactionTableSizeMultiplier"`
	// CompactionTableSizeMultiplierPerLevel defines per-level multiplier for
	// CompactionTableSize.
	// Use zero to skip a level.
	//
	// The default value is nil.
	CompactionTableSizeMultiplierPerLevel []float64 `json:"compactionTableSizeMultiplierPerLevel"`
	// CompactionTotalSize limits total size of 'sorted table' for each level.
	// The limits for each level will be calculated as:
	//   CompactionTotalSize * (CompactionTotalSizeMultiplier ^ Level)
	// The multiplier for each level can also fine-tuned using
	// CompactionTotalSizeMultiplierPerLevel.
	//
	// The default value is 10MiB.
	CompactionTotalSize int `json:"compactionTotalSize"`
	// CompactionTotalSizeMultiplier defines multiplier for CompactionTotalSize.
	//
	// The default value is 10.
	CompactionTotalSizeMultiplier float64 `json:"compactionTotalSizeMultiplier"`
	// DisableSeeksCompaction allows disabling 'seeks triggered compaction'.
	// The purpose of 'seeks triggered compaction' is to optimize database so
	// that 'level seeks' can be minimized, however this might generate many
	// small compaction which may not preferable.
	//
	// The default is true.
	DisableSeeksCompaction bool `json:"disableSeeksCompaction"`
	// OpenFilesCacheCapacity defines the capacity of the open files caching.
	// Use -1 for zero, this has same effect as specifying NoCacher to OpenFilesCacher.
	//
	// The default value is 1024.
	OpenFilesCacheCapacity int `json:"openFilesCacheCapacity"`
	// WriteBuffer defines maximum size of a 'memdb' before flushed to
	// 'sorted table'. 'memdb' is an in-memory DB backed by an on-disk
	// unsorted journal.
	//
	// LevelDB may held up to two 'memdb' at the same time.
	//
	// The default value is 6MiB.
	WriteBuffer      int `json:"writeBuffer"`
	FilterBitsPerKey int `json:"filterBitsPerKey"`

	// MaxManifestFileSize is the maximum size limit of the MANIFEST-****** file.
	// When the MANIFEST-****** file grows beyond this size, LevelDB will create
	// a new MANIFEST file.
	//
	// The default value is infinity.
	MaxManifestFileSize int64 `json:"maxManifestFileSize"`

	// MetricUpdateFrequency is the frequency to poll LevelDB metrics.
	// If <= 0, LevelDB metrics aren't polled.
	MetricUpdateFrequency time.Duration `json:"metricUpdateFrequency"`
}

// New returns a wrapped LevelDB object.
func New(file string, configBytes []byte, log logging.Logger, reg prometheus.Registerer) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

// Open the db and recover any potential corruptions

// Drop any close error to report the original error

// Has returns if the key is set in the database
func (db *Database) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Get returns the value the key maps to in the database
func (db *Database) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Put sets the value of the provided key to the provided value
func (db *Database) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

// Delete removes the key from the database
func (db *Database) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// NewBatch creates a write/delete-only buffer that is atomically committed to
// the database when write is called
func (db *Database) NewBatch() database.Batch {
	_ = "STUB: not implemented"
	return *

	// NewIterator creates a lexicographically ordered iterator over the database
	new(database.Batch)
}

func (db *Database) NewIterator() database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// NewIteratorWithStart creates a lexicographically ordered iterator over the
// database starting at the provided key
func (db *Database) NewIteratorWithStart(start []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// NewIteratorWithPrefix creates a lexicographically ordered iterator over the
// database ignoring keys that do not start with the provided prefix
func (db *Database) NewIteratorWithPrefix(prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// NewIteratorWithStartAndPrefix creates a lexicographically ordered iterator
// over the database starting at start and ignoring keys that do not start with
// the provided prefix
func (db *Database) NewIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// This comment is basically copy pasted from the underlying levelDB library:

// Compact the underlying DB for the given key range.
// Specifically, deleted and overwritten versions are discarded,
// and the data is rearranged to reduce the cost of operations
// needed to access the data. This operation should typically only
// be invoked by users who understand the underlying implementation.
//
// A nil start is treated as a key before all keys in the DB.
// And a nil limit is treated as a key after all keys in the DB.
// Therefore if both are nil then it will compact entire DB.
func (db *Database) Compact(start []byte, limit []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batch is a wrapper around a levelDB batch to contain sizes.
type batch struct {
	leveldb.Batch
	db   *Database
	size int
}

// Put the value into the batch for later writing
func (b *batch) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

// Delete the key during writing
func (b *batch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// Size retrieves the amount of data queued up for writing.
func (b *batch) Size() int {
	_ = "STUB: not implemented"

	// Write flushes any accumulated data to disk.
	return 0
}

func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

// Reset resets the batch for reuse.
func (b *batch) Reset() { _ = "STUB: not implemented"; return }

// Replay the batch contents.
func (b *batch) Replay(w database.KeyValueWriterDeleter) error {
	_ = "STUB: not implemented"
	return nil
}

// Never actually returns an error, because Replay just returns nil

// Inner returns itself
func (b *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }

type replayer struct {
	writerDeleter database.KeyValueWriterDeleter
	err           error
}

func (r *replayer) Put(key, value []byte) { _ = "STUB: not implemented"; return }

func (r *replayer) Delete(key []byte) { _ = "STUB: not implemented"; return }

type iter struct {
	db *Database
	iterator.Iterator

	key, val []byte
	err      error
}

func (it *iter) Next() bool {
	_ = "STUB: not implemented"
	// Short-circuit and set an error if the underlying database has been closed.
	return false
}

func (it *iter) Error() error { _ = "STUB: not implemented"; return nil }

func (it *iter) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iter) Value() []byte { _ = "STUB: not implemented"; return nil }

func updateError(err error) error { _ = "STUB: not implemented"; return nil }
