// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/metric"
)

type vertexMetrics struct {
	parse,
	parseErr,
	verify,
	verifyErr,
	accept,
	reject metric.Averager
}

func (m *vertexMetrics) Initialize(reg prometheus.Registerer) error {
	_ = "STUB: not implemented"
	return nil
}
