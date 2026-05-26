// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package gastime measures time based on the consumption of gas.
package gastime

import (
	"math"
	"time"

	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/proxytime"
)

//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

// Time represents an instant in time, its passage measured in [gas.Gas]
// consumption. It is not thread safe nor is the zero value valid.
//
// In addition to the passage of time, it also tracks excess consumption above a
// target, as described in [ACP-194] as a "continuous" version of [ACP-176].
//
// Copying a Time, either directly or by dereferencing a pointer, can result in
// undefined behaviour; use [Time.Clone] instead.
//
// [ACP-176]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/176-dynamic-evm-gas-limit-and-price-discovery-updates
// [ACP-194]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/194-streaming-asynchronous-execution
//
//nolint:tagliatelle,revive // tagliatelle: can't handle embedded field; struct-tag: canoto allows unexported fields
type Time struct {
	*proxytime.Time[gas.Gas] `canoto:"pointer,1"`
	target                   gas.Gas        `canoto:"uint,2"`
	excess                   gas.Gas        `canoto:"uint,3"`
	config                   GasPriceConfig `canoto:"value,4"`

	canotoData canotoData_Time `canoto:"nocopy"`
}

// New returns a new [Time], derived from a [time.Time]. The consumption of
// `target` * [TargetToRate] units of [gas.Gas] is equivalent to a tick of 1
// second.
func New(at time.Time, target, startingExcess gas.Gas, c GasPriceConfig) (*Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(StephenButtolph): startingExcess is pretty difficult for a caller to
// meaningfully provide. We should instead take in startingPrice.

// MinTarget is the minimum allowable [Time.Target] to avoid division by zero.
// Values below this are silently clamped.
const MinTarget = gas.Gas(1)

// TargetToRate is the ratio between [Time.Target] and [proxytime.Time.Rate].
const TargetToRate = 2

// MaxTarget is the maximum allowable [Time.Target] to avoid overflows of the
// associated [proxytime.Time.Rate]. Values above this are silently clamped.
const MaxTarget = gas.Gas(math.MaxUint64 / TargetToRate)

func rateOf(target gas.Gas) gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }
func clampTarget(t gas.Gas) gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// SafeRateOfTarget returns the corresponding rate for the given gas target,
// after clamping it to the allowable range.
//
// The argument is clamped to the range [[MinTarget], [MaxTarget]] and
// multiplied by [TargetToRate].
func SafeRateOfTarget(target gas.Gas) gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// Clone returns a deep copy of the time.
func (tm *Time) Clone() *Time { _ = "STUB: not implemented"; return nil }

// Target returns the `T` parameter of ACP-176.
func (tm *Time) Target() gas.Gas {
	_ = "STUB: not implemented"

	// Excess returns the `x` variable of ACP-176.
	return *new(gas.Gas)
}

func (tm *Time) Excess() gas.Gas {
	_ = "STUB: not implemented"

	// Price returns the price of a unit of gas, i.e. the "base fee", determined by
	// [gas.CalculatePrice].
	return *new(gas.Gas)
}

func (tm *Time) Price() gas.Price { _ = "STUB: not implemented"; return *new(gas.Price) }

// When minPrice can't be represented by e^(x/k), p may be too low.

// excessScalingFactor returns the K variable of ACP-103/176, i.e.
// [GasPriceConfig.TargetToExcessScaling] * T, capped at [math.MaxUint64].
//
// TODO(StephenButtolph): Rather than capping this at MaxUint64, we should move
// the evaluation of T * K into the exponential calculation. This would allow us
// to never round any values during calculation of extreme inputs.
func (tm *Time) excessScalingFactor() gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// BaseFee is equivalent to [Time.Price], returning the result as a uint256 for
// compatibility with geth/libevm objects.
func (tm *Time) BaseFee() *uint256.Int { _ = "STUB: not implemented"; return nil }

// Tick is equivalent to [proxytime.Time.Tick] except that it also updates the
// gas excess.
func (tm *Time) Tick(g gas.Gas) {
	_ = "STUB: not implemented"

	// static pricing keeps excess at its minimum
	return
}

//nolint:revive // unexported-naming: mathematical convention
// overflow is impossible as (R-T)/R < 1

// FastForwardTo is equivalent to [proxytime.Time.FastForwardTo] except that it
// may also update the gas excess.
func (tm *Time) FastForwardTo(to uint64, toFrac gas.Gas) { _ = "STUB: not implemented"; return }

//nolint:revive // unexported-naming: mathematical convention

// Excess is reduced by the amount of gas skipped (g), multiplied by T/R.
// However, to avoid overflow, the implementation needs to be a bit more
// complicated. The reduction in excess can be calculated as follows (math
// notation, not code, and ignoring the bounding at zero):
//
// s := seconds fast-forwarded (`sec`)
// f := `frac.Numerator`
// x := excess
//
// dx = -g·T/R
// = -(sR + f)·T/R
// = -sR·T/R - fT/R
// = -sT - fT/R
//
// Note that this is equivalent to the ACP reduction of T·dt because dt is
// equal to s + f/R since `frac.Denominator == R` is a documented invariant.
// Therefore dx = -(s + f/R)·T, but we separate the terms differently for
// our implementation.

// -sT
// sT <= x; division is safe because T > 0

// sT > x

// -fT/R
// overflow is impossible as T/R < 1
