// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"cmp"
)

// TODO can we handle sorting where the Compare function relies on a codec?

type Sortable[T any] interface {
	Compare(T) int
}

// Sorts the elements of [s].
func Sort[T Sortable[T]](s []T) { _ = "STUB: not implemented"; return }

// Sorts the elements of [s] based on their hashes.
func SortByHash[T ~[]byte](s []T) { _ = "STUB: not implemented"; return }

// Returns true iff the elements in [s] are sorted.
func IsSortedBytes[T ~[]byte](s []T) bool { _ = "STUB: not implemented"; return false }

// Returns true iff the elements in [s] are unique and sorted.
func IsSortedAndUnique[T Sortable[T]](s []T) bool { _ = "STUB: not implemented"; return false }

// Returns true iff the elements in [s] are unique and sorted.
func IsSortedAndUniqueOrdered[T cmp.Ordered](s []T) bool { _ = "STUB: not implemented"; return false }

// Returns true iff the elements in [s] are unique and sorted
// based by their hashes.
func IsSortedAndUniqueByHash[T ~[]byte](s []T) bool { _ = "STUB: not implemented"; return false }
