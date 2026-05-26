// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bag

import (
	"github.com/ava-labs/avalanchego/utils/set"
)

// Maps a key to a bitset.
type UniqueBag[T comparable] map[T]set.Bits64

func (b *UniqueBag[T]) init() {
	if *b == nil {
		*b = make(map[T]set.Bits64, minBagSize)
	}
}

// Adds [n] to the bitset associated with each key in [keys].
func (b *UniqueBag[T]) Add(n uint, keys ...T) { _ = "STUB: not implemented"; return }

// Unions [set] with the bitset associated with [key].
func (b *UniqueBag[T]) UnionSet(key T, set set.Bits64) { _ = "STUB: not implemented"; return }

// Removes each element of [set] from the bitset associated with [key].
func (b *UniqueBag[T]) DifferenceSet(key T, set set.Bits64) { _ = "STUB: not implemented"; return }

// For each key/bitset pair in [diff], removes each element of the bitset
// from the bitset associated with the key in [b].
// Keys in [diff] that are not in [b] are ignored.
// Bitset elements in [diff] that are not in the bitset associated with
// the key in [b] are ignored.
func (b *UniqueBag[T]) Difference(diff *UniqueBag[T]) { _ = "STUB: not implemented"; return }

// Returns the bitset associated with [key].
func (b *UniqueBag[T]) GetSet(key T) set.Bits64 {
	_ = "STUB: not implemented"

	// Removes the bitset associated with [key].
	return *new(set.Bits64)
}

func (b *UniqueBag[T]) RemoveSet(key T) {
	_ = "STUB: not implemented"

	// Returns the keys.
	return
}

func (b *UniqueBag[T]) List() []T { _ = "STUB: not implemented"; return nil }

// Returns a bag with the given [threshold] where each key is
// in the bag once for each element in the key's bitset.
func (b *UniqueBag[T]) Bag(threshold int) Bag[T] { _ = "STUB: not implemented"; return nil }

func (b *UniqueBag[T]) PrefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *UniqueBag[_]) String() string { _ = "STUB: not implemented"; return "" }

// Removes all key --> bitset pairs.
func (b *UniqueBag[_]) Clear() { _ = "STUB: not implemented"; return }
