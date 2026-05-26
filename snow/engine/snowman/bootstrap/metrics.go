// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	numFetched, numAccepted prometheus.Counter
}

func newMetrics(registerer prometheus.Registerer) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
