// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package setmap

import (
	"github.com/ava-labs/avalanchego/utils/set"
)

type Entry[K any, V comparable] struct {
	Key K
	Set set.Set[V]
}

// SetMap is a map to a set where all sets are non-overlapping.
type SetMap[K, V comparable] struct {
	keyToSet   map[K]set.Set[V]
	valueToKey map[V]K
}

// New creates a new empty setmap.
func New[K, V comparable]() *SetMap[K, V] { _ = "STUB: not implemented"; return nil }

// Put the new entry into the map. Removes and returns:
// * The existing entry for [key].
// * Existing entries where the set overlaps with the [set].
func (m *SetMap[K, V]) Put(key K, set set.Set[V]) []Entry[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// GetKey that maps to the provided value.
func (m *SetMap[K, V]) GetKey(val V) (K, bool) { _ = "STUB: not implemented"; return *new(K), false }

// GetSet that is mapped to by the provided key.
func (m *SetMap[K, V]) GetSet(key K) (set.Set[V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// HasKey returns true if [key] is in the map.
func (m *SetMap[K, _]) HasKey(key K) bool { _ = "STUB: not implemented"; return false }

// HasValue returns true if [val] is in a set in the map.
func (m *SetMap[_, V]) HasValue(val V) bool { _ = "STUB: not implemented"; return false }

// HasOverlap returns true if [set] overlaps with any of the sets in the map.
func (m *SetMap[_, V]) HasOverlap(set set.Set[V]) bool { _ = "STUB: not implemented"; return false }

// DeleteKey removes [key] from the map and returns the set it mapped to.
func (m *SetMap[K, V]) DeleteKey(key K) (set.Set[V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// DeleteValue removes and returns the entry that contained [val].
func (m *SetMap[K, V]) DeleteValue(val V) (K, set.Set[V], bool) {
	_ = "STUB: not implemented"
	return *new(K), nil, false
}

// DeleteOverlapping removes and returns all the entries where the set overlaps
// with [set].
func (m *SetMap[K, V]) DeleteOverlapping(set set.Set[V]) []Entry[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Len return the number of sets in the map.
func (m *SetMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// LenValues return the total number of values across all sets in the map.
func (m *SetMap[K, V]) LenValues() int { _ = "STUB: not implemented"; return 0 }
