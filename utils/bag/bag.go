// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bag

import (
	"github.com/ava-labs/avalanchego/utils/set"
)

const minBagSize = 16

// Bag is a multiset.
type Bag[T comparable] struct {
	counts map[T]int
	size   int

	threshold    int
	metThreshold set.Set[T]
}

// Of returns a Bag initialized with [elts]
func Of[T comparable](elts ...T) Bag[T] { _ = "STUB: not implemented"; return nil }

func (b *Bag[T]) init() {
	if b.counts == nil {
		b.counts = make(map[T]int, minBagSize)
	}
}

// SetThreshold sets the number of times an element must be added to be contained in
// the threshold set.
func (b *Bag[_]) SetThreshold(threshold int) { _ = "STUB: not implemented"; return }

// Add increases the number of times each element has been seen by one.
func (b *Bag[T]) Add(elts ...T) { _ = "STUB: not implemented"; return }

// AddCount increases the number of times the element has been seen by [count].
// If [count] <= 0 this is a no-op.
func (b *Bag[T]) AddCount(elt T, count int) { _ = "STUB: not implemented"; return }

// Count returns the number of [elt] in the bag.
func (b *Bag[T]) Count(elt T) int { _ = "STUB: not implemented"; return 0 }

// Len returns the number of elements in the bag.
func (b *Bag[_]) Len() int {
	_ = "STUB: not implemented"

	// List returns a list of unique elements that have been added.
	// The returned list doesn't have duplicates.
	return 0
}

func (b *Bag[T]) List() []T { _ = "STUB: not implemented"; return nil }

// Equals returns true if the bags contain the same elements
func (b *Bag[T]) Equals(other Bag[T]) bool { _ = "STUB: not implemented"; return false }

// Mode returns the most common element in the bag and the count of that element.
// If there's a tie, any of the tied element may be returned.
func (b *Bag[T]) Mode() (T, int) { _ = "STUB: not implemented"; return *new(T), 0 }

// Threshold returns the elements that have been seen at least threshold times.
func (b *Bag[T]) Threshold() set.Set[T] { _ = "STUB: not implemented"; return nil }

// Returns a bag with the elements of this bag that return true for [filterFunc],
// along with their counts.
// For example, if X is in this bag with count 5, and filterFunc(X) returns true,
// then the returned bag contains X with count 5.
func (b *Bag[T]) Filter(filterFunc func(T) bool) Bag[T] { _ = "STUB: not implemented"; return nil }

// Returns:
// 1. A bag containing the elements of this bag that return false for [splitFunc].
// 2. A bag containing the elements of this bag that return true for [splitFunc].
// Counts are preserved in the returned bags.
// For example, if X is in this bag with count 5, and splitFunc(X) is false,
// then the first returned bag has X in it with count 5.
func (b *Bag[T]) Split(splitFunc func(T) bool) [2]Bag[T] { _ = "STUB: not implemented"; return nil }

// Remove all instances of [elt] from the bag.
func (b *Bag[T]) Remove(elt T) { _ = "STUB: not implemented"; return }

func (b *Bag[T]) PrefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *Bag[_]) String() string { _ = "STUB: not implemented"; return "" }

func (b *Bag[T]) Clone() Bag[T] { _ = "STUB: not implemented"; return nil }
