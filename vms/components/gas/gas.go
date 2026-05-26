// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gas

import (
	"math"

	"github.com/holiman/uint256"
)

var maxUint64 = new(uint256.Int).SetUint64(math.MaxUint64)

type (
	Gas   uint64
	Price uint64
)

// Cost converts the gas to nAVAX based on the price.
//
// If overflow would occur, an error is returned.
func (g Gas) Cost(price Price) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// AddOverTime returns g + gasRate * duration.
//
// If overflow would occur, MaxUint64 is returned.
func (g Gas) AddOverTime(gasRate Gas, duration uint64) Gas {
	_ = "STUB: not implemented"
	return *new(Gas)
}

// SubOverTime returns g - gasRate * duration.
//
// If underflow would occur, 0 is returned.
func (g Gas) SubOverTime(gasRate Gas, duration uint64) Gas {
	_ = "STUB: not implemented"
	return *new(Gas)
}

// CalculatePrice returns the gas price given the minimum gas price, the
// excess gas, and the excess conversion constant.
//
// It is defined as an approximation of:
//
//	minPrice * e^(excess / excessConversionConstant)
//
// This implements the EIP-4844 fake exponential formula:
//
//	def fake_exponential(factor: int, numerator: int, denominator: int) -> int:
//		i = 1
//		output = 0
//		numerator_accum = factor * denominator
//		while numerator_accum > 0:
//			output += numerator_accum
//			numerator_accum = (numerator_accum * numerator) // (denominator * i)
//			i += 1
//		return output // denominator
//
// This implementation is optimized with the knowledge that any value greater
// than MaxUint64 gets returned as MaxUint64. This means that every intermediate
// value is guaranteed to be at most MaxUint193. So, we can safely use
// uint256.Int.
//
// This function does not perform any memory allocations.
//
//nolint:dupword // The python is copied from the EIP-4844 specification
func CalculatePrice(
	minPrice Price,
	excess Gas,
	excessConversionConstant Gas,
) Price {
	_ = "STUB: not implemented"
	return *new(Price)
}

// range is [0, MaxUint64]
// range is [0, MaxUint64]

// range is [0, MaxUint64]
// range is [0, MaxUint128]

// range is [0, MaxUint128]

// range is [0, MaxUint192+MaxUint128]

// maxOutput < MaxUint128 so numeratorAccum < MaxUint128.
// range is [0, MaxUint192]
