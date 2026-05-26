// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	numFetchedVts, numAcceptedVts,
	numFetchedTxs, numAcceptedTxs prometheus.Counter
}

func (m *metrics) Initialize(registerer prometheus.Registerer) error {
	_ = "STUB: not implemented"
	return nil
}
