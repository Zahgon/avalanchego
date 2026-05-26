// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	dto "github.com/prometheus/client_model/go"
)

// MultiGatherer extends the Gatherer interface by allowing additional gatherers
// to be registered.
type MultiGatherer interface {
	prometheus.Gatherer

	// Register adds the outputs of [gatherer] to the results of future calls to
	// Gather with the provided [name] added to the metrics.
	Register(name string, gatherer prometheus.Gatherer) error

	// Deregister removes the outputs of a gatherer with [name] from the results
	// of future calls to Gather. Returns true if a gatherer with [name] was
	// found.
	Deregister(name string) bool
}

type multiGatherer struct {
	lock      sync.RWMutex
	names     []string
	gatherers prometheus.Gatherers
}

func (g *multiGatherer) Gather() ([]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *multiGatherer) register(name string, gatherer prometheus.Gatherer) {
	_ = "STUB: not implemented"
	return
}

func (g *multiGatherer) Deregister(name string) bool { _ = "STUB: not implemented"; return false }

func MakeAndRegister(gatherer MultiGatherer, name string) (*prometheus.Registry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
