// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import "sync"

var (
	_ FIFOCache[int, int] = (*BufferFIFOCache[int, int])(nil)
	_ FIFOCache[int, int] = (*NoOpFIFOCache[int, int])(nil)
)

// FIFOCache evicts the oldest element added to it after [limit] items are
// added.
type FIFOCache[K comparable, V any] interface {
	Put(K, V)
	Get(K) (V, bool)
}

// NewFIFOCache creates a new First-In-First-Out cache of size [limit].
//
// If a [limit] of 0 is passed as an argument, a no-op cache is returned that
// does nothing.
func NewFIFOCache[K comparable, V any](limit int) FIFOCache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

type BufferFIFOCache[K comparable, V any] struct {
	l sync.RWMutex

	buffer *BoundedBuffer[K]
	m      map[K]V
}

func (f *BufferFIFOCache[K, V]) Put(key K, val V) { _ = "STUB: not implemented"; return }

// Insert will remove the oldest [K] if we are at the [limit] -
// remove always returns nil, so it is safe to ignore this error.

func (f *BufferFIFOCache[K, V]) Get(key K) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// remove is used as the callback in [BoundedBuffer]. It is assumed that the
// [WriteLock] is held when this is accessed.
func (f *BufferFIFOCache[K, V]) remove(key K) error { _ = "STUB: not implemented"; return nil }

type NoOpFIFOCache[K comparable, V any] struct{}

func (*NoOpFIFOCache[K, V]) Put(K, V)        { _ = "STUB: not implemented"; return }
func (*NoOpFIFOCache[K, V]) Get(K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }
