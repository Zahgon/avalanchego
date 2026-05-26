// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	dto "github.com/prometheus/client_model/go"
)

// "metric name" -> "metric value"
type NodeMetrics map[string]*dto.MetricFamily

// URI -> "metric name" -> "metric value"
type NodesMetrics map[string]NodeMetrics

// GetNodeMetrics retrieves the specified metrics the provided node URI.
func GetNodeMetrics(ctx context.Context, nodeURI string) (NodeMetrics, error) {
	_ = "STUB: not implemented"
	return *new(NodeMetrics), nil
}

// GetNodesMetrics retrieves the specified metrics for the provided node URIs.
func GetNodesMetrics(ctx context.Context, nodeURIs []string) (NodesMetrics, error) {
	_ = "STUB: not implemented"
	return *new(NodesMetrics), nil
}

// GetMetricValue returns the value of the specified metric which has the
// required labels.
//
// If multiple metrics match the provided labels, the first metric found is
// returned.
//
// Only Counter and Gauge metrics are supported.
func GetMetricValue(metrics NodeMetrics, name string, labels prometheus.Labels) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func labelsMatch(metric *dto.Metric, labels prometheus.Labels) bool {
	_ = "STUB: not implemented"
	return false
}
