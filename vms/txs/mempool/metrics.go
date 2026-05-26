// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mempool

import (
	"github.com/prometheus/client_golang/prometheus"
)

var _ Metrics = (*metrics)(nil)

type metrics struct {
	numTxs               prometheus.Gauge
	bytesAvailableMetric prometheus.Gauge
}

func NewMetrics(namespace string, registerer prometheus.Registerer) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metrics) Update(numTxs, bytesAvailable int) { _ = "STUB: not implemented"; return }
