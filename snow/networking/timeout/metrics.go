// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timeout

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/snow"
)

const (
	chainLabel = "chain"
	opLabel    = "op"
)

var opLabels = []string{chainLabel, opLabel}

type timeoutMetrics struct {
	messages         *prometheus.CounterVec // chain + op
	messageLatencies *prometheus.GaugeVec   // chain + op

	lock           sync.RWMutex
	chainIDToAlias map[ids.ID]string
}

func newTimeoutMetrics(reg prometheus.Registerer) (*timeoutMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *timeoutMetrics) RegisterChain(ctx *snow.ConsensusContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Record that a response of type [op] took [latency]
func (m *timeoutMetrics) Observe(chainID ids.ID, op message.Op, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}
