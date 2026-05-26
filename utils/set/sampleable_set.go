// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package set

import (
	"encoding/json"
)

var _ json.Marshaler = (*Set[int])(nil)

// SampleableSet is a set of elements that supports sampling.
type SampleableSet[T comparable] struct {
	// indices maps the element in the set to the index that it appears in
	// elements.
	indices  map[T]int
	elements []T
}

// OfSampleable returns a Set initialized with [elts]
func OfSampleable[T comparable](elts ...T) SampleableSet[T] { _ = "STUB: not implemented"; return nil }

// Return a new sampleable set with initial capacity [size].
// More or less than [size] elements can be added to this set.
// Using NewSampleableSet() rather than SampleableSet[T]{} is just an
// optimization that can be used if you know how many elements will be put in
// this set.
func NewSampleableSet[T comparable](size int) SampleableSet[T] {
	_ = "STUB: not implemented"
	return nil
}

// Add all the elements to this set.
// If the element is already in the set, nothing happens.
func (s *SampleableSet[T]) Add(elements ...T) { _ = "STUB: not implemented"; return }

// Union adds all the elements from the provided set to this set.
func (s *SampleableSet[T]) Union(set SampleableSet[T]) { _ = "STUB: not implemented"; return }

// Difference removes all the elements in [set] from [s].
func (s *SampleableSet[T]) Difference(set SampleableSet[T]) { _ = "STUB: not implemented"; return }

// Contains returns true iff the set contains this element.
func (s SampleableSet[T]) Contains(e T) bool { _ = "STUB: not implemented"; return false }

// Overlaps returns true if the intersection of the set is non-empty
func (s SampleableSet[T]) Overlaps(big SampleableSet[T]) bool {
	_ = "STUB: not implemented"
	return false
}

// Len returns the number of elements in this set.
func (s SampleableSet[_]) Len() int { _ = "STUB: not implemented"; return 0 }

// Remove all the given elements from this set.
// If an element isn't in the set, it's ignored.
func (s *SampleableSet[T]) Remove(elements ...T) { _ = "STUB: not implemented"; return }

// Clear empties this set
func (s *SampleableSet[T]) Clear() { _ = "STUB: not implemented"; return }

// List converts this set into a list
func (s SampleableSet[T]) List() []T { _ = "STUB: not implemented"; return nil }

// Equals returns true if the sets contain the same elements
func (s SampleableSet[T]) Equals(other SampleableSet[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func (s SampleableSet[T]) Sample(numToSample int) []T { _ = "STUB: not implemented"; return nil }

func (s *SampleableSet[T]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SampleableSet[_]) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sort for determinism

// Build the JSON

func (s *SampleableSet[T]) resize(size int) { _ = "STUB: not implemented"; return }

func (s *SampleableSet[T]) add(e T) { _ = "STUB: not implemented"; return }

func (s *SampleableSet[T]) remove(e T) { _ = "STUB: not implemented"; return }
