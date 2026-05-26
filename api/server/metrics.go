// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	numProcessing *prometheus.GaugeVec
	numCalls      *prometheus.CounterVec
	totalDuration *prometheus.GaugeVec
}

func newMetrics(registerer prometheus.Registerer) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metrics) wrapHandler(chainName string, handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
