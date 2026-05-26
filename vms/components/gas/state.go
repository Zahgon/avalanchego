// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gas

import (
	"errors"
)

var ErrInsufficientCapacity = errors.New("insufficient capacity")

type State struct {
	Capacity Gas `serialize:"true" json:"capacity"`
	Excess   Gas `serialize:"true" json:"excess"`
}

// AdvanceTime adds capacityRate to capacity and subtracts targetRate
// from excess over the provided duration.
//
// The units chosen for time must be consistent with the units chosen for
// capacityRate and targetRate.
//
// Capacity is capped at maxCapacity.
// Excess to be removed is capped at excess.
func (s State) AdvanceTime(
	maxCapacity Gas,
	capacityRate Gas,
	targetRate Gas,
	duration uint64,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// ConsumeGas removes gas from capacity and adds gas to excess.
//
// If the capacity is insufficient, an error is returned.
// If the excess would overflow, it is capped at MaxUint64.
func (s State) ConsumeGas(gas Gas) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

//nolint:nilerr // excess is capped at MaxUint64
