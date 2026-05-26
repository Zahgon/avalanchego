// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handler

import (
	"github.com/prometheus/client_golang/prometheus"
)

const opLabel = "op"

var opLabels = []string{opLabel}

type messageQueueMetrics struct {
	count             *prometheus.GaugeVec
	nodesWithMessages prometheus.Gauge
	numExcessiveCPU   prometheus.Counter
}

func (m *messageQueueMetrics) initialize(
	metricsNamespace string,
	metricsRegisterer prometheus.Registerer,
) error {
	_ = "STUB: not implemented"
	return nil
}
