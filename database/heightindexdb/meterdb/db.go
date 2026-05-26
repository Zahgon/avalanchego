// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package meterdb

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
)

const methodLabel = "method"

var (
	_ database.HeightIndex = (*Database)(nil)

	methodLabels = []string{methodLabel}
	putLabel     = prometheus.Labels{
		methodLabel: "put",
	}
	getLabel = prometheus.Labels{
		methodLabel: "get",
	}
	hasLabel = prometheus.Labels{
		methodLabel: "has",
	}
	syncLabel = prometheus.Labels{
		methodLabel: "sync",
	}
	closeLabel = prometheus.Labels{
		methodLabel: "close",
	}
)

// Database tracks the amount of time each operation takes and how many bytes
// are read/written to the underlying height index database.
type Database struct {
	heightDB database.HeightIndex

	calls    *prometheus.CounterVec
	duration *prometheus.GaugeVec
	size     *prometheus.CounterVec
}

func New(
	reg prometheus.Registerer,
	namespace string,
	db database.HeightIndex,
) (*Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Database) Put(height uint64, block []byte) error { _ = "STUB: not implemented"; return nil }

func (db *Database) Get(height uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *Database) Has(height uint64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db *Database) Sync(startHeight, endHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }
