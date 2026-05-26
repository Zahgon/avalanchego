// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bimap

import (
	"encoding/json"
	"errors"
)

var (
	_ json.Marshaler   = (*BiMap[int, int])(nil)
	_ json.Unmarshaler = (*BiMap[int, int])(nil)

	nullBytes       = []byte("null")
	errNotBijective = errors.New("map not bijective")
)

type Entry[K, V any] struct {
	Key   K
	Value V
}

// BiMap is a bi-directional map.
type BiMap[K, V comparable] struct {
	keyToValue map[K]V
	valueToKey map[V]K
}

// New creates a new empty bimap.
func New[K, V comparable]() *BiMap[K, V] { _ = "STUB: not implemented"; return nil }

// Put the key value pair into the map. If either [key] or [val] was previously
// in the map, the previous entries will be removed and returned.
//
// Note: Unlike normal maps, it's possible that Put removes 0, 1, or 2 existing
// entries to ensure that mappings are one-to-one.
func (m *BiMap[K, V]) Put(key K, val V) []Entry[K, V] { _ = "STUB: not implemented"; return nil }

// GetKey that maps to the provided value.
func (m *BiMap[K, V]) GetKey(val V) (K, bool) { _ = "STUB: not implemented"; return *new(K), false }

// GetValue that is mapped to the provided key.
func (m *BiMap[K, V]) GetValue(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// HasKey returns true if [key] is in the map.
func (m *BiMap[K, _]) HasKey(key K) bool { _ = "STUB: not implemented"; return false }

// HasValue returns true if [val] is in the map.
func (m *BiMap[_, V]) HasValue(val V) bool { _ = "STUB: not implemented"; return false }

// DeleteKey removes [key] from the map and returns the value it mapped to.
func (m *BiMap[K, V]) DeleteKey(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// DeleteValue removes [val] from the map and returns the key that mapped to it.
func (m *BiMap[K, V]) DeleteValue(val V) (K, bool) {
	_ = "STUB: not implemented"
	return *new(K), false
}

// Keys returns the keys of the map. The keys will be in an indeterminate order.
func (m *BiMap[K, _]) Keys() []K { _ = "STUB: not implemented"; return nil }

// Values returns the values of the map. The values will be in an indeterminate
// order.
func (m *BiMap[_, V]) Values() []V { _ = "STUB: not implemented"; return nil }

// Len return the number of entries in this map.
func (m *BiMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *BiMap[K, V]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *BiMap[K, V]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
