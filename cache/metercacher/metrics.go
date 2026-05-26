// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metercacher

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	resultLabel = "result"
	hitResult   = "hit"
	missResult  = "miss"
)

var (
	resultLabels = []string{resultLabel}
	hitLabels    = prometheus.Labels{
		resultLabel: hitResult,
	}
	missLabels = prometheus.Labels{
		resultLabel: missResult,
	}
)

type metrics struct {
	getCount *prometheus.CounterVec
	getTime  *prometheus.GaugeVec

	putCount prometheus.Counter
	putTime  prometheus.Gauge

	len           prometheus.Gauge
	portionFilled prometheus.Gauge
}

func newMetrics(
	namespace string,
	reg prometheus.Registerer,
) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
