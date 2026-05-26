// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build !prod && !nocmpopts

// Package cmputils provides [cmp] options and utilities for their creation.
package cmputils

import (
	"github.com/google/go-cmp/cmp"
)

// IfIn returns a filtered equivalent of `opt` such that it is only evaluated if
// the [cmp.Path] includes at least one `T`. This is typically used for struct
// fields (and sub-fields).
func IfIn[T any](opt cmp.Option) cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

func pathIncludes[T any](p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// WithNilCheck returns a function that returns:
//
//	   true if both a and b are nil
//	  false if exactly one of a or b is nil
//	fn(a,b) if neither a nor b are nil
func WithNilCheck[T any](fn func(*T, *T) bool) func(*T, *T) bool {
	_ = "STUB: not implemented"
	return nil
}

// ComparerWithNilCheck is a convenience wrapper, returning a [cmp.Comparer]
// after wrapping `fn` in [WithNilCheck].
func ComparerWithNilCheck[T any](fn func(*T, *T) bool) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}
