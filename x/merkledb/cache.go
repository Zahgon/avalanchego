// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/utils/linked"
)

var errEmptyCacheTooLarge = errors.New("cache is empty yet still too large")

// A cache that calls [onEviction] on the evicted element.
type onEvictCache[K comparable, V any] struct {
	lock        sync.RWMutex
	maxSize     int
	currentSize int
	fifo        *linked.Hashmap[K, V]
	size        func(K, V) int
	// Must not call any method that grabs [c.lock]
	// because this would cause a deadlock.
	onEviction func(K, V) error
}

// [size] must always return a positive number.
func newOnEvictCache[K comparable, V any](
	maxSize int,
	size func(K, V) int,
	onEviction func(K, V) error,
) onEvictCache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Get an element from this cache.
func (c *onEvictCache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Put an element into this cache. If this causes an element
// to be evicted, calls [c.onEviction] on the evicted element
// and returns the error from [c.onEviction]. Otherwise, returns nil.
func (c *onEvictCache[K, V]) Put(key K, value V) error { _ = "STUB: not implemented"; return nil }

// Mark as MRU

// Flush removes all elements from the cache.
//
// Returns the first non-nil error returned by [c.onEviction], if any.
//
// If [c.onEviction] errors, it will still be called for any subsequent elements
// and the cache will still be emptied.
func (c *onEvictCache[K, V]) Flush() error { _ = "STUB: not implemented"; return nil }

// removeOldest returns and removes the oldest element from this cache.
//
// Assumes [c.lock] is held.
func (c *onEvictCache[K, V]) removeOldest() (K, V, bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

// resize removes the oldest elements from the cache until the cache is not
// larger than the provided target.
//
// Assumes [c.lock] is held.
func (c *onEvictCache[K, V]) resize(target int) error {
	_ = "STUB: not implemented"
	// Note that we can't use [c.fifo]'s iterator because [c.onEviction]
	// modifies [c.fifo], which violates the iterator's invariant.
	return nil
}

// This should really never happen unless the size of an entry
// changed or the target size is negative.
