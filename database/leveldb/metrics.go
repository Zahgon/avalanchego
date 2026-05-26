// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package leveldb

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/syndtr/goleveldb/leveldb"
)

var levelLabels = []string{"level"}

type metrics struct {
	// total number of writes that have been delayed due to compaction
	writesDelayedCount prometheus.Counter
	// total amount of time (in ns) that writes that have been delayed due to
	// compaction
	writesDelayedDuration prometheus.Gauge
	// set to 1 if there is currently at least one write that is being delayed
	// due to compaction
	writeIsDelayed prometheus.Gauge

	// number of currently alive snapshots
	aliveSnapshots prometheus.Gauge
	// number of currently alive iterators
	aliveIterators prometheus.Gauge

	// total amount of data written
	ioWrite prometheus.Counter
	// total amount of data read
	ioRead prometheus.Counter

	// total number of bytes of cached data blocks
	blockCacheSize prometheus.Gauge
	// current number of open tables
	openTables prometheus.Gauge

	// number of tables per level
	levelTableCount *prometheus.GaugeVec
	// size of each level
	levelSize *prometheus.GaugeVec
	// amount of time spent compacting each level
	levelDuration *prometheus.GaugeVec
	// amount of bytes read while compacting each level
	levelReads *prometheus.CounterVec
	// amount of bytes written while compacting each level
	levelWrites *prometheus.CounterVec

	// total number memory compactions performed
	memCompactions prometheus.Counter
	// total number of level 0 compactions performed
	level0Compactions prometheus.Counter
	// total number of non-level 0 compactions performed
	nonLevel0Compactions prometheus.Counter
	// total number of seek compactions performed
	seekCompactions prometheus.Counter

	priorStats, currentStats *leveldb.DBStats
}

func newMetrics(reg prometheus.Registerer) (metrics, error) {
	_ = "STUB: not implemented"
	return *new(metrics), nil
}

func (db *Database) updateMetrics() error { _ = "STUB: not implemented"; return nil }

// Retrieve the database stats

// update the priorStats to update the counters correctly next time this
// method is called

// update currentStats to a pre-allocated stats struct. This avoids
// performing memory allocations for each update
