// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package snapshot

import (
	"github.com/ava-labs/libevm/common"
	"golang.org/x/exp/slices"
)

// weightedIterator is a iterator with an assigned weight. It is used to prioritise
// which account or storage slot is the correct one if multiple iterators find the
// same one (modified in multiple consecutive blocks).
type weightedIterator struct {
	it       Iterator
	priority int
}

func (it *weightedIterator) Cmp(other *weightedIterator) int {
	_ = "STUB: not implemented"
	// Order the iterators primarily by the account hashes
	return 0
}

// Same account/storage-slot in multiple layers, split by priority

// fastIterator is a more optimized multi-layer iterator which maintains a
// direct mapping of all iterators leading down to the bottom layer.
type fastIterator struct {
	tree *Tree       // Snapshot tree to reinitialize stale sub-iterators with
	root common.Hash // Root hash to reinitialize stale sub-iterators through

	curAccount []byte
	curSlot    []byte

	iterators []*weightedIterator
	initiated bool
	account   bool
	fail      error
}

// newFastIterator creates a new hierarchical account or storage iterator with one
// element per diff layer. The returned combo iterator can be used to walk over
// the entire snapshot diff stack simultaneously.
func newFastIterator(tree *Tree, root common.Hash, account common.Hash, seek common.Hash, accountIterator bool, holdsTreeLock bool) (*fastIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the whole storage is destructed in this layer, don't
// bother deeper layer anymore. But we should still keep
// the iterator for this layer, since the iterator can contain
// some valid slots which belongs to the re-created account.

// init walks over all the iterators and resolves any clashes between them, after
// which it prepares the stack for step-by-step iteration.
func (fi *fastIterator) init() {
	// Track which account hashes are iterators positioned on
	var positioned = make(map[common.Hash]int)

	// Position all iterators and track how many remain live
	for i := 0; i < len(fi.iterators); i++ {
		// Retrieve the first element and if it clashes with a previous iterator,
		// advance either the current one or the old one. Repeat until nothing is
		// clashing any more.
		it := fi.iterators[i]
		for {
			// If the iterator is exhausted, drop it off the end
			if !it.it.Next() {
				it.it.Release()
				last := len(fi.iterators) - 1

				fi.iterators[i] = fi.iterators[last]
				fi.iterators[last] = nil
				fi.iterators = fi.iterators[:last]

				i--
				break
			}
			// The iterator is still alive, check for collisions with previous ones
			hash := it.it.Hash()
			if other, exist := positioned[hash]; !exist {
				positioned[hash] = i
				break
			} else {
				// Iterators collide, one needs to be progressed, use priority to
				// determine which.
				//
				// This whole else-block can be avoided, if we instead
				// do an initial priority-sort of the iterators. If we do that,
				// then we'll only wind up here if a lower-priority (preferred) iterator
				// has the same value, and then we will always just continue.
				// However, it costs an extra sort, so it's probably not better
				if fi.iterators[other].priority < it.priority {
					// The 'it' should be progressed
					continue
				} else {
					// The 'other' should be progressed, swap them
					it = fi.iterators[other]
					fi.iterators[other], fi.iterators[i] = fi.iterators[i], fi.iterators[other]
					continue
				}
			}
		}
	}
	// Re-sort the entire list
	slices.SortFunc(fi.iterators, func(a, b *weightedIterator) int { return a.Cmp(b) })
	fi.initiated = false
}

// Next steps the iterator forward one element, returning false if exhausted.
func (fi *fastIterator) Next() bool { _ = "STUB: not implemented"; return false }

// Don't forward first time -- we had to 'Next' once in order to
// do the sorting already

// Implicit else: we've hit a nil-account or nil-slot, and need to
// fall through to the loop below to land on something non-nil

// If an account or a slot is deleted in one of the layers, the key will
// still be there, but the actual value will be nil. However, the iterator
// should not export nil-values (but instead simply omit the key), so we
// need to loop here until we either
//  - get a non-nil value,
//  - hit an error,
//  - or exhaust the iterator

// exhausted

// error

// non-nil value found

// next handles the next operation internally and should be invoked when we know
// that two elements in the list may have the same value.
//
// For example, if the iterated hashes become [2,3,5,5,8,9,10], then we should
// invoke next(3), which will call Next on elem 3 (the second '5') and will
// cascade along the list, applying the same operation if needed.
func (fi *fastIterator) next(idx int) bool {
	_ = "STUB: not implemented"
	// If this particular iterator got exhausted, remove it and return true (the
	// next one is surely not exhausted yet, otherwise it would have been removed
	// already).
	return false
}

// If there's no one left to cascade into, return

// We next-ed the iterator at 'idx', now we may have to re-sort that element

// It is still in correct place

// So still in correct place, but we need to iterate on the next

// At this point, the iterator is in the wrong location, but the remaining
// list is sorted. Find out where to move the item.

// The iterator always advances forward, so anything before the old slot
// is known to be behind us, so just skip them altogether. This actually
// is an important clause since the sort order got invalidated.

// Can always place an elem last

// The elem we're placing it next to has the same value,
// so whichever winds up on n+1 will need further iteration

// move advances an iterator to another position in the list.
func (fi *fastIterator) move(index, newpos int) { _ = "STUB: not implemented"; return }

// Error returns any failure that occurred during iteration, which might have
// caused a premature iteration exit (e.g. snapshot stack becoming stale).
func (fi *fastIterator) Error() error {
	_ = "STUB: not implemented"

	// Hash returns the current key
	return nil
}

func (fi *fastIterator) Hash() common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

// Account returns the current account blob.
// Note the returned account is not a copy, please don't modify it.
func (fi *fastIterator) Account() []byte { _ = "STUB: not implemented"; return nil }

// Slot returns the current storage slot.
// Note the returned slot is not a copy, please don't modify it.
func (fi *fastIterator) Slot() []byte {
	_ = "STUB: not implemented"

	// Release iterates over all the remaining live layer iterators and releases each
	// of them individually.
	return nil
}

func (fi *fastIterator) Release() { _ = "STUB: not implemented"; return }

// Debug is a convenience helper during testing
func (fi *fastIterator) Debug() { _ = "STUB: not implemented"; return }

// newFastAccountIterator creates a new hierarchical account iterator with one
// element per diff layer. The returned combo iterator can be used to walk over
// the entire snapshot diff stack simultaneously.
func newFastAccountIterator(tree *Tree, root common.Hash, seek common.Hash, holdsTreeLock bool) (AccountIterator, error) {
	_ = "STUB: not implemented"
	return *new(AccountIterator), nil
}

// newFastStorageIterator creates a new hierarchical storage iterator with one
// element per diff layer. The returned combo iterator can be used to walk over
// the entire snapshot diff stack simultaneously.
func newFastStorageIterator(tree *Tree, root common.Hash, account common.Hash, seek common.Hash, holdsTreeLock bool) (StorageIterator, error) {
	_ = "STUB: not implemented"
	return *new(StorageIterator), nil
}
