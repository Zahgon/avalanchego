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
	"github.com/ava-labs/avalanchego/snow/networking/benchlist"
	"github.com/ava-labs/avalanchego/utils/timer"
)

func NewManager(
	timeoutConfig *timer.AdaptiveTimeoutConfig,
	benchlistMgr benchlist.Manager,
	requestReg prometheus.Registerer,
	responseReg prometheus.Registerer,
) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Manager manages timeouts for requests sent to peers.
type Manager struct {
	tm           timer.AdaptiveTimeoutManager
	benchlistMgr benchlist.Manager
	metrics      *timeoutMetrics
	stopOnce     sync.Once
}

// Dispatch starts the manager. Must be called before any other method.
// Should be called in a goroutine.
func (m *Manager) Dispatch() {
	_ = "STUB: not implemented"

	// TimeoutDuration returns the current timeout duration.
	return
}

func (m *Manager) TimeoutDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// IsBenched returns true if messages to [nodeID] regarding [chainID]
// should not be sent over the network and should immediately fail.
func (m *Manager) IsBenched(chainID ids.ID, nodeID ids.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

// RegisterChain registers the existence of the given chain.
// Must be called before any method calls that use the
// ID of the chain.
func (m *Manager) RegisterChain(ctx *snow.ConsensusContext) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRequest notes that we expect a response of type [op] from
// [nodeID] regarding chain [chainID]. If we don't receive a response in
// time, [timeoutHandler] is executed.
func (m *Manager) RegisterRequest(
	nodeID ids.NodeID,
	chainID ids.ID,
	measureLatency bool,
	requestID ids.RequestID,
	timeoutHandler func(),
) {
	_ = "STUB: not implemented"
	return
}

// If the request timed out and wasn't an AppRequest, tell the
// benchlist manager.

// RegisterResponse registers that [nodeID] sent us a response of type [op]
// for the given chain. The response corresponds to the given
// [requestID] we sent them. [latency] is the time between us
// sending them the request and receiving their response.
func (m *Manager) RegisterResponse(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID ids.RequestID,
	op message.Op,
	latency time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

// RemoveRequest marks that we no longer expect a response to this request.
// Does not modify the timeout.
func (m *Manager) RemoveRequest(requestID ids.RequestID) { _ = "STUB: not implemented"; return }

// Stop stops the manager.
func (m *Manager) Stop() { _ = "STUB: not implemented"; return }
