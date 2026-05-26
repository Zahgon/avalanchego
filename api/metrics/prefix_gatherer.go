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

	errOverlappingNamespaces = errors.New("prefix could create overlapping namespaces")
)

// NewPrefixGatherer returns a new MultiGatherer that merges metrics by adding a
// prefix to their names.
func NewPrefixGatherer() MultiGatherer { _ = "STUB: not implemented"; return *new(MultiGatherer) }

type prefixGatherer struct {
	multiGatherer
}

func (g *prefixGatherer) Register(prefix string, gatherer prometheus.Gatherer) error {
	_ = "STUB: not implemented"
	return nil
}

type prefixedGatherer struct {
	prefix   string
	gatherer prometheus.Gatherer
}

func (g *prefixedGatherer) Gather() ([]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	// Gather returns partially filled metrics in the case of an error. So, it
	// is expected to still return the metrics in the case an error is returned.
	return nil, nil
}

// eitherIsPrefix returns true if either [a] is a prefix of [b] or [b] is a
// prefix of [a].
//
// This function accounts for the usage of the namespace boundary, so "hello" is
// not considered a prefix of "helloworld". However, "hello" is considered a
// prefix of "hello_world".
func eitherIsPrefix(a, b string) bool { _ = "STUB: not implemented"; return false }

// a is a prefix of b
// a is empty
// a is equal to b
// a ends at a namespace boundary of b
