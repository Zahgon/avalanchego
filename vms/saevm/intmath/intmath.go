// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package intmath provides special-case integer arithmetic.
package intmath

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// BoundedSubtract returns `max(a-b,floor)` without underflow.
func BoundedSubtract[T constraints.Unsigned](a, b, floor T) T {
	_ = "STUB: not implemented"
	// If `floor + b` overflows then it's impossible for `a` to ever be large
	// enough for the subtraction to not be bounded.
	return *new(T)
}

// BoundedAdd returns `min(a+b,ceil)` without overflow.
func BoundedAdd[T constraints.Unsigned](a, b, ceil T) T { _ = "STUB: not implemented"; return *new(T) }

// BoundedMultiply returns `min(a*b,ceil)` without overflow.
func BoundedMultiply[T constraints.Unsigned](a, b, ceil T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// ErrOverflow is returned if a return value would have overflowed its type.
var ErrOverflow = errors.New("overflow")

// MulDiv returns the quotient and remainder of `(a*b)/den` without overflow in
// the event that `a*b>=2^64`. However, if the quotient were to overflow then
// [ErrOverflow] is returned.
func MulDiv[T ~uint64](a, b, den T) (quo, rem T, err error) {
	_ = "STUB: not implemented"
	return *new(T), *new(T), nil
}

// MulDivCeil is equivalent to [MulDiv] except that it returns the rounded-up
// quotient and the complement of the remainder, i.e. the amount that would have
// had to be added to `a*b` to result in the same quotient exactly.
func MulDivCeil[T ~uint64](a, b, den T) (quo T, extra T, err error) {
	_ = "STUB: not implemented"
	return *new(T), *new(T), nil
}

func mulDiv[T ~uint64](a, b, den T, ceil bool) (quo, rem T, err error) {
	_ = "STUB: not implemented"
	return *new(T), *new(T), nil
}

// If both `a` and `b` are MaxUint64 then `hi==MaxUint64-1` so this
// can't overflow because `carry ∈ {0,1}`.

// CeilDiv returns `ceil(num/den)`, i.e. the rounded-up quotient.
func CeilDiv[T ~uint64](num, den T) T { _ = "STUB: not implemented"; return *new(T) }

// [bits.Div64] panics if the denominator is zero (expected behaviour) or if
// `den <= hi`. The latter is impossible because `hi` is a carry bit (i.e.
// can only be 0 or 1) and even if `num==MaxUint64` then `den` would have to
// be `>=2` for `hi` to be non-zero.
