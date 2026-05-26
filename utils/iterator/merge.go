// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package iterator

import (
	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/utils/heap"
)

var _ Iterator[any] = (*merged[any])(nil)

type merged[T any] struct {
	initialized bool
	// heap only contains iterators that have been initialized and are not
	// exhausted.
	heap heap.Queue[Iterator[T]]
}

// Merge returns an iterator that returns all of the elements of [iterators] in
// order.
func Merge[T any](less btree.LessFunc[T], iterators ...Iterator[T]) Iterator[T] {
	_ = "STUB: not implemented"
	// Filter out iterators that are already exhausted.
	return nil
}

func (it *merged[_]) Next() bool { _ = "STUB: not implemented"; return false }

// Note that on the first call to Next() (i.e. here) we don't call
// Next() on the current iterator. This is because we already called
// Next() on each iterator in Merge.

// Update the heap root.

// Calling Next() above modifies [current] so we fix the heap.

// The old root is exhausted. Remove it from the heap.

func (it *merged[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (it *merged[_]) Release() { _ = "STUB: not implemented"; return }
