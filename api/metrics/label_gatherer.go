// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	dto "github.com/prometheus/client_model/go"
)

var (
	_ MultiGatherer = (*prefixGatherer)(nil)

	errDuplicateGatherer = errors.New("attempt to register duplicate gatherer")
)

// NewLabelGatherer returns a new MultiGatherer that merges metrics by adding a
// new label.
func NewLabelGatherer(labelName string) MultiGatherer {
	_ = "STUB: not implemented"
	return *new(MultiGatherer)
}

type labelGatherer struct {
	multiGatherer

	labelName string
}

func (g *labelGatherer) Register(labelValue string, gatherer prometheus.Gatherer) error {
	_ = "STUB: not implemented"
	return nil
}

type labeledGatherer struct {
	labelName  string
	labelValue string
	gatherer   prometheus.Gatherer
}

func (g *labeledGatherer) Gather() ([]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	// Gather returns partially filled metrics in the case of an error. So, it
	// is expected to still return the metrics in the case an error is returned.
	return nil, nil
}
