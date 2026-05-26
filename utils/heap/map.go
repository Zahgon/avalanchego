// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package heap

import (
	"container/heap"
)

var _ heap.Interface = (*indexedQueue[int, int])(nil)

func MapValues[K comparable, V any](m Map[K, V]) []V { _ = "STUB: not implemented"; return nil }

// NewMap returns a heap without duplicates ordered by its values
func NewMap[K comparable, V any](less func(a, b V) bool) Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

type Map[K comparable, V any] struct {
	queue *indexedQueue[K, V]
}

// Push returns the evicted previous value if present
func (m *Map[K, V]) Push(k K, v V) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m *Map[K, V]) Pop() (K, V, bool) { _ = "STUB: not implemented"; return *new(K), *new(V), false }

func (m *Map[K, V]) Peek() (K, V, bool) { _ = "STUB: not implemented"; return *new(K), *new(V), false }

func (m *Map[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *Map[K, V]) Remove(k K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m *Map[K, V]) Contains(k K) bool { _ = "STUB: not implemented"; return false }

func (m *Map[K, V]) Get(k K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m *Map[K, V]) Fix(k K) { _ = "STUB: not implemented"; return }

type indexedQueue[K comparable, V any] struct {
	queue[entry[K, V]]
	index map[K]int
}

func (h *indexedQueue[K, V]) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (h *indexedQueue[K, V]) Push(x any) { _ = "STUB: not implemented"; return }

func (h *indexedQueue[K, V]) Pop() any { _ = "STUB: not implemented"; return *new(any) }

type entry[K any, V any] struct {
	k K
	v V
}
