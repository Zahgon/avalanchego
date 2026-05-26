// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	reg *prometheus.Registry
	// Summary of the quantiles of Individual Issuance Tx Times
	IssuanceTxTimes prometheus.Summary
	// Summary of the quantiles of Individual Confirmation Tx Times
	ConfirmationTxTimes prometheus.Summary
	// Summary of the quantiles of Individual Issuance To Confirmation Tx Times
	IssuanceToConfirmationTxTimes prometheus.Summary
}

func NewDefaultMetrics() *Metrics { _ = "STUB: not implemented"; return nil }

// NewMetrics creates and returns a Metrics and registers it with a Collector
func NewMetrics(reg *prometheus.Registry) *Metrics { _ = "STUB: not implemented"; return nil }

type MetricsServer struct {
	cancel context.CancelFunc
	stopCh chan struct{}
}

func (m *Metrics) Serve(ctx context.Context, metricsPort string, metricsEndpoint string) *MetricsServer {
	_ = "STUB: not implemented"
	return nil
}

// Create a prometheus server to expose individual tx metrics

// Start up go routine to listen for SIGINT notifications to gracefully shut down server

// Blocks until signal is received

// Start metrics server

func (ms *MetricsServer) Shutdown() { _ = "STUB: not implemented"; return }

func (m *Metrics) Print(outputFile string) error { _ = "STUB: not implemented"; return nil }

// Printout to stdout
