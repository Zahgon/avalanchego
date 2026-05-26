// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package linked

type keyValue[K, V any] struct {
	key   K
	value V
}

// Hashmap provides an ordered O(1) mapping from keys to values.
//
// Entries are tracked by insertion order.
type Hashmap[K comparable, V any] struct {
	entryMap  map[K]*ListElement[keyValue[K, V]]
	entryList *List[keyValue[K, V]]
	freeList  []*ListElement[keyValue[K, V]]
}

func NewHashmap[K comparable, V any]() *Hashmap[K, V] { _ = "STUB: not implemented"; return nil }

func NewHashmapWithSize[K comparable, V any](initialSize int) *Hashmap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (lh *Hashmap[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (lh *Hashmap[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (lh *Hashmap[K, V]) Delete(key K) bool { _ = "STUB: not implemented"; return false }

func (lh *Hashmap[K, V]) Clear() { _ = "STUB: not implemented"; return }

// remove assumes that [e] is currently in the Hashmap.
func (lh *Hashmap[K, V]) remove(e *ListElement[keyValue[K, V]]) { _ = "STUB: not implemented"; return }

// Free the key value pair

func (lh *Hashmap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (lh *Hashmap[K, V]) Oldest() (K, V, bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

func (lh *Hashmap[K, V]) Newest() (K, V, bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

func (lh *Hashmap[K, V]) NewIterator() *Iterator[K, V] { _ = "STUB: not implemented"; return nil }

// Iterates over the keys and values in a LinkedHashmap from oldest to newest.
// Assumes the underlying LinkedHashmap is not modified while the iterator is in
// use, except to delete elements that have already been iterated over.
type Iterator[K comparable, V any] struct {
	lh                     *Hashmap[K, V]
	key                    K
	value                  V
	next                   *ListElement[keyValue[K, V]]
	initialized, exhausted bool
}

func (it *Iterator[K, V]) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator has been exhausted, there is no next value.
	return false
}

// If the iterator was not yet initialized, do it now.

// It's important to ensure that [it.next] is not nil
// by not deleting elements that have not yet been iterated
// over from [it.lh]

// Next time, return next element

func (it *Iterator[K, V]) Key() K { _ = "STUB: not implemented"; return *new(K) }

func (it *Iterator[K, V]) Value() V { _ = "STUB: not implemented"; return *new(V) }
