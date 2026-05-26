// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"context"

	dto "github.com/prometheus/client_model/go"
)

// Client for requesting metrics from a remote AvalancheGo instance
type Client struct {
	uri string
}

// NewClient returns a new Metrics API Client
func NewClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

// GetMetrics returns the metrics from the connected node. The metrics are
// returned as a map of metric family name to the metric family.
func (c *Client) GetMetrics(ctx context.Context) (map[string]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:bodyclose // body is closed via rpc.CleanlyCloseBody

// Return an error for any non successful status code
