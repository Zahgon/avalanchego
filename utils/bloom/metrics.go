// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bloom

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics is a collection of commonly useful metrics when using a long-lived
// bloom filter.
type Metrics struct {
	Count      prometheus.Gauge
	NumHashes  prometheus.Gauge
	NumEntries prometheus.Gauge
	MaxCount   prometheus.Gauge
	ResetCount prometheus.Counter
}

func NewMetrics(
	namespace string,
	registerer prometheus.Registerer,
) (*Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset the metrics to align with the provided bloom filter and max count.
func (m *Metrics) Reset(newFilter *Filter, maxCount int) { _ = "STUB: not implemented"; return }
