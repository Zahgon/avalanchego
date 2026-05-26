// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	_ SyncMetrics = (*mockMetrics)(nil)
	_ SyncMetrics = (*metrics)(nil)
)

type SyncMetrics interface {
	RequestFailed()
	RequestMade()
	RequestSucceeded()
}

type mockMetrics struct {
	lock              sync.Mutex
	requestsFailed    int
	requestsMade      int
	requestsSucceeded int
}

func (m *mockMetrics) RequestFailed() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) RequestMade() { _ = "STUB: not implemented"; return }

func (m *mockMetrics) RequestSucceeded() { _ = "STUB: not implemented"; return }

type metrics struct {
	requestsFailed    prometheus.Counter
	requestsMade      prometheus.Counter
	requestsSucceeded prometheus.Counter
}

func NewMetrics(namespace string, reg prometheus.Registerer) (SyncMetrics, error) {
	_ = "STUB: not implemented"
	return *new(SyncMetrics), nil
}

func (m *metrics) RequestFailed() { _ = "STUB: not implemented"; return }

func (m *metrics) RequestMade() { _ = "STUB: not implemented"; return }

func (m *metrics) RequestSucceeded() { _ = "STUB: not implemented"; return }
