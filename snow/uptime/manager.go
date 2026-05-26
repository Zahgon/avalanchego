// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package uptime

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

var _ Manager = (*manager)(nil)

var (
	errAlreadyStartedTracking = errors.New("already started tracking")
	errNotStartedTracking     = errors.New("not started tracking")
)

type Manager interface {
	Tracker
	Calculator
}

type Tracker interface {
	StartTracking(nodeIDs []ids.NodeID) error
	StopTracking(nodeIDs []ids.NodeID) error
	StartedTracking() bool

	Connect(nodeID ids.NodeID) error
	Disconnect(nodeID ids.NodeID) error
}

type Calculator interface {
	CalculateUptime(nodeID ids.NodeID) (time.Duration, time.Time, error)
	CalculateUptimePercent(nodeID ids.NodeID) (float64, error)
	// CalculateUptimePercentFrom expects [startTime] to be truncated (floored) to the nearest second
	CalculateUptimePercentFrom(nodeID ids.NodeID, startTime time.Time) (float64, error)
}

type manager struct {
	// Used to get time. Useful for faking time during tests.
	clock *mockable.Clock

	state       State
	connections map[ids.NodeID]time.Time // nodeID  -> connected at
	// Whether we have started tracking the uptime of the nodes
	// This is used to avoid setting the uptime before we have started tracking
	startedTracking bool
}

func NewManager(state State, clk *mockable.Clock) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (m *manager) StartTracking(nodeIDs []ids.NodeID) error { _ = "STUB: not implemented"; return nil }

func (m *manager) StopTracking(nodeIDs []ids.NodeID) error { _ = "STUB: not implemented"; return nil }

func (m *manager) StartedTracking() bool { _ = "STUB: not implemented"; return false }

func (m *manager) Connect(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

func (m *manager) IsConnected(nodeID ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (m *manager) Disconnect(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

func (m *manager) CalculateUptime(nodeID ids.NodeID) (time.Duration, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

// If we are in a weird reality where time has gone backwards, make sure
// that we don't double count or delete any uptime.

// If we haven't started tracking, then we assume that the node has been
// online since their last update.

// If we are tracking and they aren't connected, they have been offline
// since their last update.

// The time the peer connected needs to be adjusted to ensure no time period
// is double counted.

// If we are in a weird reality where time has gone backwards, make sure
// that we don't double count or delete any uptime.

// Increase the uptimes by the amount of time this node has been running
// since the last time it's uptime was written to disk.

func (m *manager) CalculateUptimePercent(nodeID ids.NodeID) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *manager) CalculateUptimePercentFrom(nodeID ids.NodeID, startTime time.Time) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// updateUptime updates the uptime of the node on the state by the amount of
// time that the node has been connected.
func (m *manager) updateUptime(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

// We don't track the uptimes of non-validators.
