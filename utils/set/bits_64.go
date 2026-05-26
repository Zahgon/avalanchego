// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package set

// Bits64 is a set that can contain uints in the range [0, 64). All functions
// are O(1). The zero value is the empty set.
type Bits64 uint64

// Add [i] to the set of ints
func (b *Bits64) Add(i uint) {
	_ = "STUB: not implemented"

	// Union adds all the elements in [s] to this set
	return
}

func (b *Bits64) Union(s Bits64) {
	_ = "STUB: not implemented"

	// Intersection takes the intersection of [s] with this set
	return
}

func (b *Bits64) Intersection(s Bits64) {
	_ = "STUB: not implemented"

	// Difference removes all the elements in [s] from this set
	return
}

func (b *Bits64) Difference(s Bits64) {
	_ = "STUB: not implemented"

	// Remove [i] from the set of ints
	return
}

func (b *Bits64) Remove(i uint) {
	_ = "STUB: not implemented"

	// Clear removes all elements from this set
	return
}

func (b *Bits64) Clear() {
	_ = "STUB: not implemented"

	// Contains returns true if [i] was previously added to this set
	return
}

func (b Bits64) Contains(i uint) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of elements in this set
func (b Bits64) Len() int { _ = "STUB: not implemented"; return 0 }

func (b Bits64) String() string { _ = "STUB: not implemented"; return "" }
