// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package maybe

// Maybe T = Some T | Nothing.
// A data wrapper that allows values to be something [Some T] or nothing [Nothing].
// Invariant: If [hasValue] is false, then [value] is the zero value of type T.
// Maybe is used to wrap types:
// * That can't be represented by nil.
// * That use nil as a valid value instead of an indicator of a missing value.
// For more info see https://en.wikipedia.org/wiki/Option_type
type Maybe[T any] struct {
	hasValue bool
	// If [hasValue] is false, [value] is the zero value of type T.
	value T
}

// Some returns a new Maybe[T] with the value val.
// If m.IsNothing(), returns the zero value of type T.
func Some[T any](val T) Maybe[T] { _ = "STUB: not implemented"; return nil }

// Nothing returns a new Maybe[T] with no value.
func Nothing[T any]() Maybe[T] {
	_ = "STUB: not implemented"

	// IsNothing returns false iff [m] has a value.
	return nil
}

func (m Maybe[T]) IsNothing() bool {
	_ = "STUB: not implemented"

	// HasValue returns true iff [m] has a value.
	return false
}

func (m Maybe[T]) HasValue() bool {
	_ = "STUB: not implemented"

	// Value returns the value of [m].
	return false
}

func (m Maybe[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (m Maybe[T]) String() string { _ = "STUB: not implemented"; return "" }

// Bind returns Nothing iff [m] is Nothing.
// Otherwise applies [f] to the value of [m] and returns the result as a Some.
func Bind[T, U any](m Maybe[T], f func(T) U) Maybe[U] { _ = "STUB: not implemented"; return nil }

// Equal returns true if both m1 and m2 are nothing or have the same value according to [equalFunc].
func Equal[T any](m1 Maybe[T], m2 Maybe[T], equalFunc func(T, T) bool) bool {
	_ = "STUB: not implemented"
	return false
}
