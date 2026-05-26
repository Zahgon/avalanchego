// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	txsIssuedCounter      prometheus.Counter
	txIssuanceLatency     prometheus.Histogram
	txConfirmationLatency prometheus.Histogram
	txTotalLatency        prometheus.Histogram
}

func newMetrics(namespace string, registry *prometheus.Registry) (metrics, error) {
	_ = "STUB: not implemented"
	return *new(metrics), nil
}

func (m metrics) issue(d time.Duration) { _ = "STUB: not implemented"; return }

func (m metrics) accept(confirmationDuration time.Duration, totalDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}
