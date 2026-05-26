// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package lru

import (
	"sync"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/utils/linked"
)

var _ cache.Cacher[struct{}, any] = (*SizedCache[struct{}, any])(nil)

// sizedElement is used to store the element with its size, so we don't
// calculate the size multiple times.
//
// This ensures that any inconsistencies returned by the size function can not
// corrupt the cache.
type sizedElement[V any] struct {
	value V
	size  int
}

// SizedCache is a key value store with bounded size. If the size is attempted
// to be exceeded, then elements are removed from the cache until the bound is
// honored, based on evicting the least recently used value.
type SizedCache[K comparable, V any] struct {
	lock        sync.Mutex
	elements    *linked.Hashmap[K, *sizedElement[V]]
	maxSize     int
	currentSize int
	size        func(K, V) int
}

func NewSizedCache[K comparable, V any](maxSize int, size func(K, V) int) *SizedCache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *SizedCache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (c *SizedCache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *SizedCache[K, V]) Evict(key K) { _ = "STUB: not implemented"; return }

func (c *SizedCache[K, V]) Flush() { _ = "STUB: not implemented"; return }

func (c *SizedCache[_, _]) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *SizedCache[_, _]) PortionFilled() float64 { _ = "STUB: not implemented"; return 0 }

func (c *SizedCache[K, V]) put(key K, value V) { _ = "STUB: not implemented"; return }

// Remove elements until the size of elements in the cache <= [c.maxSize].

func (c *SizedCache[K, V]) get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Mark [k] as MRU.

func (c *SizedCache[K, _]) evict(key K) { _ = "STUB: not implemented"; return }

func (c *SizedCache[K, V]) flush() { _ = "STUB: not implemented"; return }

func (c *SizedCache[_, _]) len() int { _ = "STUB: not implemented"; return 0 }

func (c *SizedCache[_, _]) portionFilled() float64 { _ = "STUB: not implemented"; return 0 }
