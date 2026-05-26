// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package meterdb

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
)

const methodLabel = "method"

var (
	_ database.Database = (*Database)(nil)
	_ database.Batch    = (*batch)(nil)
	_ database.Iterator = (*iterator)(nil)

	methodLabels = []string{methodLabel}
	hasLabel     = prometheus.Labels{
		methodLabel: "has",
	}
	getLabel = prometheus.Labels{
		methodLabel: "get",
	}
	putLabel = prometheus.Labels{
		methodLabel: "put",
	}
	deleteLabel = prometheus.Labels{
		methodLabel: "delete",
	}
	newBatchLabel = prometheus.Labels{
		methodLabel: "new_batch",
	}
	newIteratorLabel = prometheus.Labels{
		methodLabel: "new_iterator",
	}
	compactLabel = prometheus.Labels{
		methodLabel: "compact",
	}
	closeLabel = prometheus.Labels{
		methodLabel: "close",
	}
	healthCheckLabel = prometheus.Labels{
		methodLabel: "health_check",
	}
	batchPutLabel = prometheus.Labels{
		methodLabel: "batch_put",
	}
	batchDeleteLabel = prometheus.Labels{
		methodLabel: "batch_delete",
	}
	batchSizeLabel = prometheus.Labels{
		methodLabel: "batch_size",
	}
	batchWriteLabel = prometheus.Labels{
		methodLabel: "batch_write",
	}
	batchResetLabel = prometheus.Labels{
		methodLabel: "batch_reset",
	}
	batchReplayLabel = prometheus.Labels{
		methodLabel: "batch_replay",
	}
	batchInnerLabel = prometheus.Labels{
		methodLabel: "batch_inner",
	}
	iteratorNextLabel = prometheus.Labels{
		methodLabel: "iterator_next",
	}
	iteratorErrorLabel = prometheus.Labels{
		methodLabel: "iterator_error",
	}
	iteratorKeyLabel = prometheus.Labels{
		methodLabel: "iterator_key",
	}
	iteratorValueLabel = prometheus.Labels{
		methodLabel: "iterator_value",
	}
	iteratorReleaseLabel = prometheus.Labels{
		methodLabel: "iterator_release",
	}
)

// Database tracks the amount of time each operation takes and how many bytes
// are read/written to the underlying database instance.
type Database struct {
	db database.Database

	calls    *prometheus.CounterVec
	duration *prometheus.GaugeVec
	size     *prometheus.CounterVec
}

// New returns a new database with added metrics
func New(
	reg prometheus.Registerer,
	db database.Database,
) (*Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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

func (db *Database) NewIteratorWithStartAndPrefix(
	start,
	prefix []byte,
) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *Database) Compact(start, limit []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

func (db *Database) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type batch struct {
	batch database.Batch
	db    *Database
}

func (b *batch) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b *batch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (b *batch) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

func (b *batch) Reset() { _ = "STUB: not implemented"; return }

func (b *batch) Replay(w database.KeyValueWriterDeleter) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }

type iterator struct {
	iterator database.Iterator
	db       *Database
}

func (it *iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (it *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (it *iterator) Release() { _ = "STUB: not implemented"; return }
