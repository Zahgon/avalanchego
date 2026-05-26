// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	ioType    = "type"
	readType  = "read"
	writeType = "write"

	lookupType                = "type"
	valueNodeCacheType        = "valueNodeCache"
	intermediateNodeCacheType = "intermediateNodeCache"
	viewChangesValueType      = "viewChangesValue"
	viewChangesNodeType       = "viewChangesNode"

	lookupResult = "result"
	hitResult    = "hit"
	missResult   = "miss"
)

var (
	_ metrics = (*prometheusMetrics)(nil)
	_ metrics = (*mockMetrics)(nil)

	ioLabels     = []string{ioType}
	ioReadLabels = prometheus.Labels{
		ioType: readType,
	}
	ioWriteLabels = prometheus.Labels{
		ioType: writeType,
	}

	lookupLabels            = []string{lookupType, lookupResult}
	valueNodeCacheHitLabels = prometheus.Labels{
		lookupType:   valueNodeCacheType,
		lookupResult: hitResult,
	}
	valueNodeCacheMissLabels = prometheus.Labels{
		lookupType:   valueNodeCacheType,
		lookupResult: missResult,
	}
	intermediateNodeCacheHitLabels = prometheus.Labels{
		lookupType:   intermediateNodeCacheType,
		lookupResult: hitResult,
	}
	intermediateNodeCacheMissLabels = prometheus.Labels{
		lookupType:   intermediateNodeCacheType,
		lookupResult: missResult,
	}
	viewChangesValueHitLabels = prometheus.Labels{
		lookupType:   viewChangesValueType,
		lookupResult: hitResult,
	}
	viewChangesValueMissLabels = prometheus.Labels{
		lookupType:   viewChangesValueType,
		lookupResult: missResult,
	}
	viewChangesNodeHitLabels = prometheus.Labels{
		lookupType:   viewChangesNodeType,
		lookupResult: hitResult,
	}
	viewChangesNodeMissLabels = prometheus.Labels{
		lookupType:   viewChangesNodeType,
		lookupResult: missResult,
	}
)

type metrics interface {
	HashCalculated()
	DatabaseNodeRead()
	DatabaseNodeWrite()
	ValueNodeCacheHit()
	ValueNodeCacheMiss()
	IntermediateNodeCacheHit()
	IntermediateNodeCacheMiss()
	ViewChangesValueHit()
	ViewChangesValueMiss()
	ViewChangesNodeHit()
	ViewChangesNodeMiss()
}

type prometheusMetrics struct {
	hashes prometheus.Counter
	io     *prometheus.CounterVec
	lookup *prometheus.CounterVec
}

func newMetrics(prefix string, reg prometheus.Registerer) (metrics, error) {
	_ = "STUB: not implemented"
	// TODO: Should we instead return an error if reg is nil?
	return *new(metrics), nil
}

func (m *prometheusMetrics) HashCalculated() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) DatabaseNodeRead() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) DatabaseNodeWrite() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) ValueNodeCacheHit() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) ValueNodeCacheMiss() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) IntermediateNodeCacheHit() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) IntermediateNodeCacheMiss() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) ViewChangesValueHit() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) ViewChangesValueMiss() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) ViewChangesNodeHit() { _ = "STUB: not implemented"; return }

func (m *prometheusMetrics) ViewChangesNodeMiss() { _ = "STUB: not implemented"; return }

type mockMetrics struct {
	lock                      sync.Mutex
	hashCount                 int64
	nodeReadCount             int64
	nodeWriteCount            int64
	valueNodeCacheHit         int64
	valueNodeCacheMiss        int64
	intermediateNodeCacheHit  int64
	intermediateNodeCacheMiss int64
	viewChangesValueHit       int64
	viewChangesValueMiss      int64
	viewChangesNodeHit        int64
	viewChangesNodeMiss       int64
}

func (m *mockMetrics) HashCalculated() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) DatabaseNodeRead() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) DatabaseNodeWrite() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) ValueNodeCacheHit() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) ValueNodeCacheMiss() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) IntermediateNodeCacheHit() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) IntermediateNodeCacheMiss() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) ViewChangesValueHit() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) ViewChangesValueMiss() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) ViewChangesNodeHit() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) ViewChangesNodeMiss() { _ = "STUB: not implemented"; return }
