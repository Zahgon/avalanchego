// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package fee

import (
	"github.com/ava-labs/avalanchego/vms/components/gas"
)

// Config contains all the static parameters of the dynamic fee mechanism.
type Config struct {
	Capacity                 gas.Gas   `json:"capacity"`
	Target                   gas.Gas   `json:"target"`
	MinPrice                 gas.Price `json:"minPrice"`
	ExcessConversionConstant gas.Gas   `json:"excessConversionConstant"`
}

// State represents the current on-chain values used in the dynamic fee
// mechanism.
type State struct {
	Current gas.Gas `json:"current"`
	Excess  gas.Gas `json:"excess"`
}

// AdvanceTime adds (s.Current - target) * seconds to Excess.
//
// If Excess would underflow, it is set to 0.
// If Excess would overflow, it is set to MaxUint64.
func (s State) AdvanceTime(target gas.Gas, seconds uint64) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// CostOf calculates how much to charge based on the dynamic fee mechanism for
// seconds.
//
// This implements the ACP-77 cost over time formula:
func (s State) CostOf(c Config, seconds uint64) uint64 {
	_ = "STUB: not implemented"
	// If the current and target are the same, the price is constant.
	return 0
}

// Advancing the time is going to either hold excess constant,
// monotonically increase it, or monotonically decrease it. If it is
// equal to 0 after performing one of these operations, it is guaranteed
// to always remain 0.

// SecondsRemaining calculates the maximum number of seconds that a validator
// can pay fees before their fundsRemaining would be exhausted based on the
// dynamic fee mechanism. The result is capped at maxSeconds.
func (s State) SecondsRemaining(c Config, maxSeconds uint64, fundsRemaining uint64) uint64 {
	_ = "STUB: not implemented"
	// Because this function can divide by prices, we need to sanity check the
	// parameters to avoid division by 0.
	return 0
}

// If the current and target are the same, the price is constant.

// Advancing the time is going to either hold excess constant,
// monotonically increase it, or monotonically decrease it. If it is
// equal to 0 after performing one of these operations, it is guaranteed
// to always remain 0.

// This is technically unreachable, but makes the code more
// clearly correct.
