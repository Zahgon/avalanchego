// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package resource

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	numCPUCycles       *prometheus.GaugeVec
	numDiskReads       *prometheus.GaugeVec
	numDiskReadBytes   *prometheus.GaugeVec
	numDiskWrites      *prometheus.GaugeVec
	numDiskWritesBytes *prometheus.GaugeVec
}

func newMetrics(registerer prometheus.Registerer) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
