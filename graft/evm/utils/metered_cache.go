// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/ava-labs/libevm/metrics"
)

// MeteredCache wraps *fastcache.Cache and periodically pulls stats from it.
type MeteredCache struct {
	*fastcache.Cache
	namespace string

	// stats to be surfaced
	entriesCount metrics.Gauge
	bytesSize    metrics.Gauge
	collisions   metrics.Gauge
	gets         metrics.Gauge
	sets         metrics.Gauge
	misses       metrics.Gauge
	statsTime    metrics.Gauge

	// count all operations to decide when to update stats
	ops             uint64
	updateFrequency uint64
}

// NewMeteredCache returns a new MeteredCache that will update stats to the
// provided namespace once per each [updateFrequency] operations.
// Note: if [updateFrequency] is passed as 0, it will be treated as 1.
func NewMeteredCache(size int, namespace string, updateFrequency uint64) *MeteredCache {
	_ = "STUB: not implemented"
	return nil
}

// avoid division by zero

// only register stats if a namespace is provided.

// updateStatsIfNeeded updates metrics from fastcache
func (mc *MeteredCache) updateStatsIfNeeded() { _ = "STUB: not implemented"; return }

// cumulative metric

func (mc *MeteredCache) Del(k []byte) { _ = "STUB: not implemented"; return }

func (mc *MeteredCache) Get(dst, k []byte) []byte { _ = "STUB: not implemented"; return nil }

func (mc *MeteredCache) GetBig(dst, k []byte) []byte { _ = "STUB: not implemented"; return nil }

func (mc *MeteredCache) Has(k []byte) bool { _ = "STUB: not implemented"; return false }

func (mc *MeteredCache) HasGet(dst, k []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (mc *MeteredCache) Set(k, v []byte) { _ = "STUB: not implemented"; return }

func (mc *MeteredCache) SetBig(k, v []byte) { _ = "STUB: not implemented"; return }
