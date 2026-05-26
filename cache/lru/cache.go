// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package lru

import (
	"sync"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/utils/linked"
)

var _ cache.Cacher[struct{}, struct{}] = (*Cache[struct{}, struct{}])(nil)

// Cache is a key value store with bounded size. If the size is attempted to be
// exceeded, then an element is removed from the cache before the insertion is
// done, based on evicting the least recently used value.
type Cache[K comparable, V any] struct {
	lock     sync.Mutex
	elements *linked.Hashmap[K, V]
	size     int

	// onEvict is called with the key and value of an entry before eviction.
	onEvict func(K, V)
}

// NewCache creates a new LRU cache with the given size.
func NewCache[K comparable, V any](size int) *Cache[K, V] { _ = "STUB: not implemented"; return nil }

// NewCacheWithOnEvict creates a new LRU cache with the given size and eviction callback.
func NewCacheWithOnEvict[K comparable, V any](size int, onEvict func(K, V)) *Cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Mark [k] as MRU.

func (c *Cache[K, _]) Evict(key K) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) evict(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Cache[_, _]) Flush() { _ = "STUB: not implemented"; return }

func (c *Cache[_, _]) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Cache[_, _]) PortionFilled() float64 { _ = "STUB: not implemented"; return 0 }
