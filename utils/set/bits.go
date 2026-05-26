// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package set

import (
	"math/big"
)

// Bits is a bit-set backed by a big.Int
// Holds values ranging from [0, INT_MAX] (arch-dependent)
// Trying to use negative values will result in a panic.
// This implementation is NOT thread-safe.
type Bits struct {
	bits *big.Int
}

// NewBits returns a new instance of Bits with [bits] set to 1.
//
// Invariants:
// 1. Negative bits will cause a panic.
// 2. Duplicate bits are allowed but will cause a no-op.
func NewBits(bits ...int) Bits { _ = "STUB: not implemented"; return *new(Bits) }

// Add sets the [i]'th bit to 1
func (b Bits) Add(i int) { _ = "STUB: not implemented"; return }

// Union performs the set union with another set.
// This adds all elements in [other] to [b]
func (b Bits) Union(other Bits) { _ = "STUB: not implemented"; return }

// Intersection performs the set intersection with another set
// This sets [b] to include only elements in both [b] and [other]
func (b Bits) Intersection(other Bits) { _ = "STUB: not implemented"; return }

// Difference removes all the elements in [other] from this set
func (b Bits) Difference(other Bits) { _ = "STUB: not implemented"; return }

// Remove sets the [i]'th bit to 0
func (b Bits) Remove(i int) { _ = "STUB: not implemented"; return }

// Clear empties out the bitset
func (b Bits) Clear() { _ = "STUB: not implemented"; return }

// Contains returns true if the [i]'th bit is 1, and false otherwise
func (b Bits) Contains(i int) bool { _ = "STUB: not implemented"; return false }

// BitLen returns the bit length of this bitset
func (b Bits) BitLen() int { _ = "STUB: not implemented"; return 0 }

// Len returns the amount of 1's in the bitset
//
// This is typically referred to as the "Hamming Weight"
// of a set of bits.
func (b Bits) Len() int { _ = "STUB: not implemented"; return 0 }

// Returns the byte representation of this bitset
func (b Bits) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Inverse of Bits.Bytes()
func BitsFromBytes(bytes []byte) Bits { _ = "STUB: not implemented"; return *new(Bits) }

// String returns the hex representation of this bitset
func (b Bits) String() string { _ = "STUB: not implemented"; return "" }
