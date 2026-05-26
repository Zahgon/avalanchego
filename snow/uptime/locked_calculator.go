// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package uptime

import (
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
)

var (
	errStillBootstrapping = errors.New("still bootstrapping")

	_ LockedCalculator = (*lockedCalculator)(nil)
)

type LockedCalculator interface {
	Calculator

	SetCalculator(isBootstrapped *utils.Atomic[bool], lock sync.Locker, newC Calculator)
}

type lockedCalculator struct {
	lock           sync.RWMutex
	isBootstrapped *utils.Atomic[bool]
	calculatorLock sync.Locker
	c              Calculator
}

func NewLockedCalculator() LockedCalculator {
	_ = "STUB: not implemented"
	return *new(LockedCalculator)
}

func (c *lockedCalculator) CalculateUptime(nodeID ids.NodeID) (time.Duration, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

func (c *lockedCalculator) CalculateUptimePercent(nodeID ids.NodeID) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *lockedCalculator) CalculateUptimePercentFrom(nodeID ids.NodeID, startTime time.Time) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *lockedCalculator) SetCalculator(isBootstrapped *utils.Atomic[bool], lock sync.Locker, newC Calculator) {
	_ = "STUB: not implemented"
	return
}
