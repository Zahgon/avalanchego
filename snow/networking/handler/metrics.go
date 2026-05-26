// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handler

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	expired             *prometheus.CounterVec // op
	messages            *prometheus.CounterVec // op
	lockingTime         prometheus.Gauge
	messageHandlingTime *prometheus.GaugeVec // op
}

func newMetrics(reg prometheus.Registerer) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
