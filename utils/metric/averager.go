// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metric

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/wrappers"
)

var ErrFailedRegistering = errors.New("failed registering metric")

type Averager interface {
	Observe(float64)
}

type averager struct {
	count prometheus.Counter
	sum   prometheus.Gauge
}

func NewAverager(name, desc string, reg prometheus.Registerer) (Averager, error) {
	_ = "STUB: not implemented"
	return *new(Averager), nil
}

func NewAveragerWithErrs(name, desc string, reg prometheus.Registerer, errs *wrappers.Errs) Averager {
	_ = "STUB: not implemented"
	return *new(Averager)
}

func (a *averager) Observe(v float64) { _ = "STUB: not implemented"; return }
