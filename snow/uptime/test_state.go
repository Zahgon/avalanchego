// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package uptime

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
)

var _ State = (*TestState)(nil)

type uptime struct {
	upDuration  time.Duration
	lastUpdated time.Time
	startTime   time.Time
}

type TestState struct {
	dbReadError  error
	dbWriteError error
	nodes        map[ids.NodeID]*uptime
}

func NewTestState() *TestState { _ = "STUB: not implemented"; return nil }

func (s *TestState) AddNode(nodeID ids.NodeID, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (s *TestState) GetUptime(nodeID ids.NodeID) (time.Duration, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

func (s *TestState) SetUptime(nodeID ids.NodeID, upDuration time.Duration, lastUpdated time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TestState) GetStartTime(nodeID ids.NodeID) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
