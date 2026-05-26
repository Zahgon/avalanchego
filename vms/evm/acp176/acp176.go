// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// ACP176 implements the fee logic specified here:
// https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/176-dynamic-evm-gas-limit-and-price-discovery-updates/README.md
package acp176

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/avalanchego/vms/components/gas"
)

const (
	MinTargetPerSecond  = 1_000_000                                 // P
	TargetConversion    = MaxTargetChangeRate * MaxTargetExcessDiff // D
	MaxTargetExcessDiff = 1 << 15                                   // Q
	MinGasPrice         = 1                                         // M

	TimeToFillCapacity            = 5    // in seconds
	TargetToMax                   = 2    // multiplier to convert from target per second to max per second
	TargetToPriceUpdateConversion = 87   // 87 ~= 60 / ln(2) which makes the price double at most every ~60 seconds
	MaxTargetChangeRate           = 1024 // Controls the rate that the target can change per block.

	TargetToMaxCapacity = TargetToMax * TimeToFillCapacity
	MinMaxPerSecond     = MinTargetPerSecond * TargetToMax
	MinMaxCapacity      = MinMaxPerSecond * TimeToFillCapacity

	StateSize = 3 * wrappers.LongLen

	maxTargetExcess = 1_024_950_627 // TargetConversion * ln(MaxUint64 / MinTargetPerSecond) + 1
)

var ErrStateInsufficientLength = errors.New("insufficient length for fee state")

// State represents the current state of the gas pricing and constraints.
type State struct {
	Gas          gas.State
	TargetExcess gas.Gas // q
}

// ParseState returns the state from the provided bytes. It is the inverse of
// [State.Bytes]. This function allows for additional bytes to be padded at the
// end of the provided bytes.
func ParseState(bytes []byte) (State, error) { _ = "STUB: not implemented"; return *new(State), nil }

// Target returns the target gas consumed per second, `T`.
//
// Target = MinTargetPerSecond * e^(TargetExcess / TargetConversion)
func (s *State) Target() gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// MaxCapacity returns the maximum possible accrued gas capacity, `C`.
func (s *State) MaxCapacity() gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// GasPrice returns the current required fee per gas.
//
// GasPrice = MinGasPrice * e^(Excess / (Target() * TargetToPriceUpdateConversion))
func (s *State) GasPrice() gas.Price { _ = "STUB: not implemented"; return *new(gas.Price) }

// K

// AdvanceSeconds increases the gas capacity and decreases the gas excess based on
// the elapsed seconds.
// This is used in Fortuna.
func (s *State) AdvanceSeconds(seconds uint64) { _ = "STUB: not implemented"; return }

// R
// C

// AdvanceMilliseconds increases the gas capacity and decreases the gas excess based on
// the elapsed milliseconds.
// This is used in Granite.
func (s *State) AdvanceMilliseconds(milliseconds uint64) { _ = "STUB: not implemented"; return }

// R - this can't overflow since 1000 > TargetToMax.
// rate used for calculating maxCapacity
// C

// ConsumeGas decreases the gas capacity and increases the gas excess by
// gasUsed + extraGasUsed. If the gas capacity is insufficient, an error is
// returned.
func (s *State) ConsumeGas(
	gasUsed uint64,
	extraGasUsed *big.Int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTargetExcess updates the targetExcess to be as close as possible to the
// desiredTargetExcess without exceeding the maximum targetExcess change.
func (s *State) UpdateTargetExcess(desiredTargetExcess gas.Gas) { _ = "STUB: not implemented"; return }

// Ensure the gas capacity does not exceed the maximum capacity.
// C

// Bytes returns the binary representation of the state.
func (s *State) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// DesiredTargetExcess calculates the optimal desiredTargetExcess given the
// desired target.
func DesiredTargetExcess(desiredTarget gas.Gas) gas.Gas {
	_ = "STUB: not implemented"
	// This could be solved directly by calculating D * ln(desiredTarget / P)
	// using floating point math. However, it introduces inaccuracies. So, we
	// use a binary search to find the closest integer solution.
	return *new(gas.Gas)
}

// targetExcess calculates the optimal new targetExcess for a block proposer to
// include given the current and desired excess values.
func targetExcess(excess, desired gas.Gas) gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// scaleExcess scales the excess during gas target modifications to keep the
// price constant.
func scaleExcess(
	excess,
	newTargetPerSecond,
	previousTargetPerSecond gas.Gas,
) gas.Gas {
	_ = "STUB: not implemented"
	return *new(gas.Gas)
}

// mulWithUpperBound multiplies two numbers and returns the result. If the
// result overflows, it returns [math.MaxUint64].
func mulWithUpperBound(a, b gas.Gas) gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }
