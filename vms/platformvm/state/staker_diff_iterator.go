// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/utils/heap"
	"github.com/ava-labs/avalanchego/utils/iterator"
)

var (
	_ StakerDiffIterator         = (*stakerDiffIterator)(nil)
	_ iterator.Iterator[*Staker] = (*mutableStakerIterator)(nil)
)

// StakerDiffIterator is an iterator that iterates over the events that will be
// performed on the current staker set.
//
// There are two event types affecting current staker set, removal of an
// existing staker and addition of a new staker from the pending set.
//
// The ordering of operations is:
//   - Staker operations are performed in order of their [NextTime].
//   - If operations have the same [NextTime], stakers are first added to the
//     current staker set, then removed.
//   - Further ties are broken by *Staker.Less(), returning the lesser staker
//     first.
type StakerDiffIterator interface {
	Next() bool
	// Returns:
	// - The staker that is changing
	// - True if the staker is being added to the current staker set, false if
	//   the staker is being removed from the current staker set
	Value() (*Staker, bool)
	Release()
}

type stakerDiffIterator struct {
	currentIteratorExhausted bool
	currentIterator          *mutableStakerIterator

	pendingIteratorExhausted bool
	pendingIterator          iterator.Iterator[*Staker]

	modifiedStaker *Staker
	isAdded        bool
}

func NewStakerDiffIterator(currentIterator, pendingIterator iterator.Iterator[*Staker]) StakerDiffIterator {
	_ = "STUB: not implemented"
	return *new(StakerDiffIterator)
}

func (it *stakerDiffIterator) Next() bool { _ = "STUB: not implemented"; return false }

// If the next operations share the same time, we default to adding the
// staker to the current staker set. This means that we default to
// advancing the pending iterator.

func (it *stakerDiffIterator) Value() (*Staker, bool) { _ = "STUB: not implemented"; return nil, false }

func (it *stakerDiffIterator) Release() { _ = "STUB: not implemented"; return }

func (it *stakerDiffIterator) advanceCurrent() { _ = "STUB: not implemented"; return }

func (it *stakerDiffIterator) advancePending() { _ = "STUB: not implemented"; return }

type mutableStakerIterator struct {
	iteratorExhausted bool
	iterator          iterator.Iterator[*Staker]
	heap              heap.Queue[*Staker]
}

func newMutableStakerIterator(iterator iterator.Iterator[*Staker]) *mutableStakerIterator {
	_ = "STUB: not implemented"
	return nil
}

// Add should not be called until after Next has been called at least once.
func (it *mutableStakerIterator) Add(staker *Staker) { _ = "STUB: not implemented"; return }

func (it *mutableStakerIterator) Next() bool {
	_ = "STUB: not implemented"
	// The only time the heap should be empty - is when the iterator is
	// exhausted or uninitialized.
	return false
}

// If the iterator is exhausted, the only elements left to iterate over are
// in the heap.

// If the heap doesn't contain the next staker to return, we need to move
// the next element from the iterator into the heap.

func (it *mutableStakerIterator) Value() *Staker { _ = "STUB: not implemented"; return nil }

func (it *mutableStakerIterator) Release() { _ = "STUB: not implemented"; return }
