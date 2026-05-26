// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	localhostAddr      = "127.0.0.1"
	defaultMetricsPort = 0
)

// PrometheusServer is a HTTP server that serves Prometheus metrics from the provided
// gahterer.
// Listens on localhost with a dynamic port and serves metrics at /ext/metrics.
type PrometheusServer struct {
	gatherer prometheus.Gatherer
	server   http.Server
	errChan  chan error
}

// NewPrometheusServer creates and starts a Prometheus server with the provided gatherer
// listening on 127.0.0.1:0 and serving /ext/metrics.
func NewPrometheusServer(gatherer prometheus.Gatherer) (*PrometheusServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewPrometheusServerWithPort creates and starts a Prometheus server with the provided gatherer
// listening on 127.0.0.1:port and serving /ext/metrics.
func NewPrometheusServerWithPort(gatherer prometheus.Gatherer, port uint64) (*PrometheusServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// start the Prometheus server on address.
func (s *PrometheusServer) start(address string) error { _ = "STUB: not implemented"; return nil }

// Stop gracefully shuts down the Prometheus server.
// Waits for the server to shut down and returns any error that occurred during shutdown.
func (s *PrometheusServer) Stop() error { _ = "STUB: not implemented"; return nil }

// Address returns the address the server is listening on.
// If the server has not started, the address will be empty.
func (s *PrometheusServer) Address() string { _ = "STUB: not implemented"; return "" }
