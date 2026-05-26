// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package proxytime measures the passage of time based on a proxy unit and
// associated unit rate.
package proxytime

import (
	"time"
)

//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

// A Duration is a type parameter for use as the unit of passage of [Time].
type Duration interface {
	~uint64
}

// Time represents an instant in time, its passage measured by an arbitrary unit
// of duration. It is not thread safe nor is the zero value valid.
//
//nolint:revive // struct-tag: canoto allows unexported fields
type Time[D Duration] struct {
	seconds uint64 `canoto:"uint,1"`
	// invariant: fraction < hertz
	fraction D `canoto:"uint,2"`
	hertz    D `canoto:"uint,3"`

	canotoData canotoData_Time `canoto:"nocopy"`
}

// IMPORTANT: keep [Time.Clone] next to the struct definition to make it easier
// to check that all fields are copied.

// Clone returns a copy of the time.
func (tm *Time[D]) Clone() *Time[D] { _ = "STUB: not implemented"; return nil }

// New returns a new [Time], set from a Unix timestamp. The passage of `hertz`
// units is equivalent to a tick of 1 second.
func New[D Duration](unixSeconds uint64, hertz D) *Time[D] { _ = "STUB: not implemented"; return nil }

// Of converts the time at nanosecond resolution. [time.Time.Unix] MUST be
// non-negative.
func Of[D Duration](t time.Time) *Time[D] { _ = "STUB: not implemented"; return nil }

//#nosec G115 -- Known to be non-negative

// Unix returns tm as a Unix timestamp.
func (tm *Time[D]) Unix() uint64 {
	_ = "STUB: not implemented"

	// A FractionalSecond represents a sub-second duration of time. The numerator is
	// equivalent to a value passed to [Time.Tick] when [Time.Rate] is the
	// denominator.
	return 0
}

type FractionalSecond[D Duration] struct {
	Numerator, Denominator D
}

// Fraction returns the fractional-second component of the time, denominated in
// [Time.Rate].
func (tm *Time[D]) Fraction() FractionalSecond[D] { _ = "STUB: not implemented"; return nil }

// Compare returns
//
//	-1 if f < g
//	 0 if f == g
//	+1 if f > g.
func (f FractionalSecond[D]) Compare(g FractionalSecond[D]) int {
	_ = "STUB: not implemented"
	// Cross-multiplication of non-negative components maintains order.
	return 0
}

// Rate returns the proxy duration required for the passage of one second.
func (tm *Time[D]) Rate() D {
	_ = "STUB: not implemented"

	// Tick advances the time by `d`.
	return *new(D)
}

func (tm *Time[D]) Tick(d D) { _ = "STUB: not implemented"; return }

// FastForwardTo sets the time to the specified Unix timestamp and fraction of a
// second, if it is in the future, returning the integer and fraction number of
// seconds by which the time was advanced. Both the incoming argument and the
// returned fraction are denominated in [Time.Rate]. Behaviour is undefined if
// `to + toFrac / tm.Rate()` overflows 64 bits, which is impossible in practice.
func (tm *Time[D]) FastForwardTo(to uint64, toFrac D) (uint64, FractionalSecond[D]) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tm *Time[D]) isFuture(sec uint64, num D) bool { _ = "STUB: not implemented"; return false }

// ConvertMilliseconds returns `ms` as a number of seconds and a fraction of a
// second, denominated in `rate` and rounded down if `rate` is not a multiple of
// 1000.
func ConvertMilliseconds[D Duration](rate D, ms uint64) (sec uint64, _ FractionalSecond[D]) {
	_ = "STUB: not implemented"
	return 0, nil
}

// overflow is impossible as ms < 1000

// SetRate changes the unit rate at which time passes. The requisite integer
// division may result in rounding up of the fractional-second component of
// time. Rounding up instead of down achieves monotonicity of the clock.
func (tm *Time[D]) SetRate(hertz D) { _ = "STUB: not implemented"; return }

// Although the > case is technically impossible, breaking the invariant
// documented in [Time.scaleFraction] would be catastrophic, so we
// defensively protect against it.

// Scale returns `val`, scaled from the existing [Time.Rate] to the newly
// specified one. See [intmath.MulDivCeil] for error cases.
func (tm *Time[D]) Scale(val, newRate D) (D, error) { _ = "STUB: not implemented"; return *new(D), nil }

// scaleFraction is a special case of [Time.Scale], as if [Time.fraction] was
// passed as the value for scaling. The invariant that [Time.fraction] is less
// than [Time.hertz] makes overflow impossible as the scaled fraction will
// always be less than `newRate`.
func (tm *Time[D]) scaleFraction(newRate D) D { _ = "STUB: not implemented"; return *new(D) }

// A broken invariant MUST be detected in tests, hence not just dropping
// the error.

// Sub returns a new [Time], `s` seconds earlier, without underflow protection.
func (tm *Time[D]) Sub(s uint64) *Time[D] { _ = "STUB: not implemented"; return nil }

// Compare returns
//
//	-1 if tm is before u
//	 0 if tm and u represent the same instant
//	+1 if tm is after u.
//
// The two instants MAY have a different [Time.Rate].
func (tm *Time[D]) Compare(u *Time[D]) int { _ = "STUB: not implemented"; return 0 }

// CompareUnix is equivalent to [Time.Compare] against a zero-fractional-second
// instant in time. Note that it does NOT only compare the seconds and that if
// `tm` has the same [Time.Unix] as `sec` but non-zero [Time.Fraction] then
// CompareUnix will return 1.
func (tm *Time[D]) CompareUnix(sec uint64) int { _ = "STUB: not implemented"; return 0 }

// AsTime converts the proxy time to a standard [time.Time] in UTC. AsTime is
// analogous to setting a rate of 1e9 (nanosecond), which might result in
// rounding up. The second-range limitations documented on [time.Unix] also
// apply to AsTime.
func (tm *Time[D]) AsTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// keeps gosec linter happy

// String returns the time as a human-readable string. It is not intended for
// parsing and its format MAY change.
func (tm *Time[D]) String() string { _ = "STUB: not implemented"; return "" }
