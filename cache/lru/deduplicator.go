// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package lru

import (
	"sync"

	"github.com/ava-labs/avalanchego/utils/linked"
)

// Evictable allows the object to be notified when it is evicted
//
// Deprecated: Remove this once the vertex state no longer uses it.
type Evictable[K comparable] interface {
	Key() K
	Evict()
}

// Deduplicator is an LRU cache that notifies the objects when they are evicted.
//
// Deprecated: Remove this once the vertex state no longer uses it.
type Deduplicator[K comparable, V Evictable[K]] struct {
	lock      sync.Mutex
	entryMap  map[K]*linked.ListElement[V]
	entryList *linked.List[V]
	size      int
}

// Deprecated: Remove this once the vertex state no longer uses it.
func NewDeduplicator[K comparable, V Evictable[K]](size int) *Deduplicator[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Deduplicate returns either the provided value, or a previously provided value
// with the same ID that hasn't yet been evicted
func (d *Deduplicator[_, V]) Deduplicate(value V) V { _ = "STUB: not implemented"; return *new(V) }

// Flush removes all entries from the cache
func (d *Deduplicator[_, _]) Flush() { _ = "STUB: not implemented"; return }
