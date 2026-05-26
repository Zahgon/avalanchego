// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package set

import (
	"encoding/json"
)

// The minimum capacity of a set
const minSetSize = 16

var _ json.Marshaler = (*Set[int])(nil)

// Set is a set of elements.
type Set[T comparable] map[T]struct{}

// Of returns a Set initialized with [elts]
func Of[T comparable](elts ...T) Set[T] { _ = "STUB: not implemented"; return nil }

// Return a new set with initial capacity [size].
// More or less than [size] elements can be added to this set.
// Using NewSet() rather than Set[T]{} is just an optimization that can
// be used if you know how many elements will be put in this set.
func NewSet[T comparable](size int) Set[T] { _ = "STUB: not implemented"; return nil }

func (s *Set[T]) resize(size int) { _ = "STUB: not implemented"; return }

// Add all the elements to this set.
// If the element is already in the set, nothing happens.
func (s *Set[T]) Add(elts ...T) { _ = "STUB: not implemented"; return }

// Union adds all the elements from the provided set to this set.
func (s *Set[T]) Union(set Set[T]) { _ = "STUB: not implemented"; return }

// Difference removes all the elements in [set] from [s].
func (s *Set[T]) Difference(set Set[T]) { _ = "STUB: not implemented"; return }

// Contains returns true iff the set contains this element.
func (s *Set[T]) Contains(elt T) bool { _ = "STUB: not implemented"; return false }

// Overlaps returns true if the intersection of the set is non-empty
func (s *Set[T]) Overlaps(big Set[T]) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of elements in this set.
func (s Set[_]) Len() int {
	_ = "STUB: not implemented"

	// Remove all the given elements from this set.
	// If an element isn't in the set, it's ignored.
	return 0
}

func (s *Set[T]) Remove(elts ...T) { _ = "STUB: not implemented"; return }

// Clear empties this set
func (s *Set[_]) Clear() {
	_ = "STUB: not implemented"

	// List converts this set into a list
	return
}

func (s Set[T]) List() []T { _ = "STUB: not implemented"; return nil }

// Equals returns true if the sets contain the same elements
func (s Set[T]) Equals(other Set[T]) bool { _ = "STUB: not implemented"; return false }

// Removes and returns an element.
// If the set is empty, does nothing and returns false.
func (s *Set[T]) Pop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s *Set[T]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s Set[_]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Sort for determinism

// Build the JSON

// Returns a random element. If the set is empty, returns false
func (s *Set[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Intersect returns the set intersection of s1 and s2
func Intersect[T comparable](s1 Set[T], s2 Set[T]) Set[T] { _ = "STUB: not implemented"; return nil }
