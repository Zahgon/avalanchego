// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/heap"
)

// A priority queue of syncWorkItems.
// Note that work item ranges never overlap.
// Supports range merging and priority updating.
// Not safe for concurrent use.
type workHeap struct {
	// Max heap of items by priority.
	// i.e. heap.Pop returns highest priority item.
	innerHeap heap.Set[*workItem]
	// The heap items sorted by range start.
	// A Nothing start is considered to be the smallest.
	sortedItems *btree.BTreeG[*workItem]
	closed      bool
}

func newWorkHeap() *workHeap { _ = "STUB: not implemented"; return nil }

// [a] is Nothing, so if [b] is Nothing, they're equal.
// Otherwise, [b] is greater.

// [a] has a value and [b] doesn't so [a] is greater.

// [a] and [b] both contain values. Compare the values.

// Marks the heap as closed.
func (wh *workHeap) Close() {
	_ = "STUB: not implemented"

	// Adds a new [item] into the heap. Will not merge items, unlike MergeInsert.
	return
}

func (wh *workHeap) Insert(item *workItem) { _ = "STUB: not implemented"; return }

// Pops and returns a work item from the heap.
// Returns nil if no work is available or the heap is closed.
func (wh *workHeap) GetWork() *workItem { _ = "STUB: not implemented"; return nil }

// Insert the item into the heap, merging it with existing items
// that share a boundary and root ID.
// e.g. if the heap contains a work item with range
// [0,10] and then [10,20] is inserted, we will merge the two
// into a single work item with range [0,20].
// e.g. if the heap contains work items [0,10] and [20,30],
// and we add [10,20], we will merge them into [0,30].
func (wh *workHeap) MergeInsert(item *workItem) { _ = "STUB: not implemented"; return }

// Find the item with the greatest start range which is less than [item.start].
// Note that the iterator function will run at most once, since it always returns false.

// [beforeItem.start, beforeItem.end] and [item.start, item.end] are
// merged into [beforeItem.start, item.end]

// Find the item with the smallest start range which is greater than [item.start].
// Note that the iterator function will run at most once, since it always returns false.

// [item.start, item.end] and [afterItem.start, afterItem.end] are merged into
// [item.start, afterItem.end].

// if the new item should be merged with both the item before and the item after,
// we can combine the before item with the after item

// combine the two ranges

// remove the second range since it is now covered by the first

// update the priority

// nothing was merged, so add new item to the heap

// We didn't merge [item] with an existing one; put it in the heap.

// Deletes [item] from the heap.
func (wh *workHeap) remove(item *workItem) { _ = "STUB: not implemented"; return }

func (wh *workHeap) Len() int { _ = "STUB: not implemented"; return 0 }

// KeyspacePercent returns the approximate percentage of work in the heap
// for a given [ids.ID] root, relative to the entire keyspace.
// Keys are truncated to 8 bytes when calculating the progress.
func (wh *workHeap) KeyspacePercent(root ids.ID) float64 { _ = "STUB: not implemented"; return 0 }

// Determine the start value (0x0 if no value)

// Determine the end value (max if no value)
