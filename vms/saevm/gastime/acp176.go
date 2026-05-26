// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gastime

import (
	"time"

	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/vms/components/gas"
)

// BeforeBlock is intended to be called before processing a block with the
// provided time. The gastime is advanced to be no earlier than the block time.
func (tm *Time) BeforeBlock(t time.Time) { _ = "STUB: not implemented"; return }

// FastForwardToTime is equivalent to [Time.FastForwardTo] except that it
// accepts a [time.Time].
func (tm *Time) FastForwardToTime(t time.Time) { _ = "STUB: not implemented"; return }

//#nosec G115 -- ns is in [0, time.Second)

// [time.Time.Nanosecond] is documented as only returning values in the
// range [0, time.Second). So either Nanosecond returned an incorrect
// value, or [intmath.MulDivCeil] incorrectly returned an error.
// Regardless, this failure MUST be detected in tests, hence not just
// dropping the error.

//#nosec G115 -- known non-negative.

// AfterBlock is intended to be called after processing a block, with the
// target and gas configuration provided.
func (tm *Time) AfterBlock(used gas.Gas, target gas.Gas, c GasPriceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// scaleExcess returns oldX * newT * newScale / (oldT * oldScale) rounded up and
// capped to [math.MaxUint64].
func scaleExcess(oldX, newT, newScale, oldT, oldScale gas.Gas) gas.Gas {
	_ = "STUB: not implemented"
	return *new(gas.Gas)
}

// Overflow can't occur, the maximum possible intermediate value is:
// MaxUint64^3 + MaxUint64^2.

// round up by adding oldK - 1

func mulAsUint256[T ~uint64](a, b T) uint256.Int {
	_ = "STUB: not implemented"
	return *new(uint256.Int)
}

// enforceMinExcess bounds excess to be no less than excessForPrice(minPrice, k).
func (tm *Time) enforceMinExcess() { _ = "STUB: not implemented"; return }

// Avoid the binary search in [excessForPrice] when the current excess
// already yields a price that satisfies the minimum.

// excessForPrice returns an integer approximation of ln(p) * k.
//
// If [calculatePrice] can produce p, excessForPrice returns the minimum excess to
// produce p. Otherwise, it returns the maximum excess to produce a number < p,
// which may happen due to overflow or integer approximation.
func excessForPrice(p gas.Price, k gas.Gas) gas.Gas {
	_ = "STUB: not implemented"
	return *

	// Binary search for the minimum x where calculatePrice(x, k) >= p.
	//
	// calculatePrice(0, k) == 1 and p > 1, so lo > 0.
	new(gas.Gas)
}

// If [calculatePrice] can't generate p due to integer approximation, honor
// the lower price expectation.

// calculatePrice returns an integer approximation of e^(x/k).
func calculatePrice(x, k gas.Gas) gas.Price { _ = "STUB: not implemented"; return *new(gas.Price) }
