// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

// Returns a new instance of a T.
func Zero[T any]() (_ T) {
	_ = "STUB: not implemented"

	// PointerTo returns a pointer to the provided value.
	return *new(T)
}

func PointerTo[T any](v T) *T { _ = "STUB: not implemented"; return nil }
