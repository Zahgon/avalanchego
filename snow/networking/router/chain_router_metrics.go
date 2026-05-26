// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"github.com/prometheus/client_golang/prometheus"
)

// routerMetrics about router messages
type routerMetrics struct {
	outstandingRequests   prometheus.Gauge
	longestRunningRequest prometheus.Gauge
	droppedRequests       prometheus.Counter
}

func newRouterMetrics(registerer prometheus.Registerer) (*routerMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
