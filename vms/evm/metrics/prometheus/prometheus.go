// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package prometheus

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	dto "github.com/prometheus/client_model/go"
)

var (
	_ prometheus.Gatherer = (*Gatherer)(nil)

	errMetricSkip             = errors.New("metric skipped")
	errMetricTypeNotSupported = errors.New("metric type is not supported")
	quantiles                 = []float64{.5, .75, .95, .99, .999, .9999}
	pvShortPercent            = []float64{50, 95, 99}
	helpText                  = ""
)

// Gatherer implements the [prometheus.Gatherer] interface by gathering all
// metrics from a [Registry].
type Gatherer struct {
	registry Registry
}

// Gather gathers metrics from the registry and converts them to
// a slice of metric families.
func (g *Gatherer) Gather() ([]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	// Gather and pre-sort the metrics to avoid random listings
	return nil, nil
}

// NewGatherer returns a [Gatherer] using the given registry.
func NewGatherer(registry Registry) *Gatherer { _ = "STUB: not implemented"; return nil }

func metricFamily(registry Registry, name string) (mf *dto.MetricFamily, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
