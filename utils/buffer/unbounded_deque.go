// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package buffer

const defaultInitSize = 32

// An unbounded deque (double-ended queue).
// See https://en.wikipedia.org/wiki/Double-ended_queue
// Not safe for concurrent access.
type Deque[T any] interface {
	// Place an element at the leftmost end of the deque.
	// Returns true if the element was placed in the deque.
	PushLeft(T) bool
	// Place an element at the rightmost end of the deque.
	// Returns true if the element was placed in the deque.
	PushRight(T) bool
	// Remove and return the leftmost element of the deque.
	// Returns false if the deque is empty.
	PopLeft() (T, bool)
	// Remove and return the rightmost element of the deque.
	// Returns false if the deque is empty.
	PopRight() (T, bool)
	// Return the leftmost element of the deque without removing it.
	// Returns false if the deque is empty.
	PeekLeft() (T, bool)
	// Return the rightmost element of the deque without removing it.
	// Returns false if the deque is empty.
	PeekRight() (T, bool)
	// Returns the element at the given index.
	// Returns false if the index is out of bounds.
	// The leftmost element is at index 0.
	Index(int) (T, bool)
	// Returns the number of elements in the deque.
	Len() int
	// Returns the elements in the deque from left to right.
	List() []T
}

// Returns a new unbounded deque with the given initial slice size.
// Note that the returned deque is always empty -- [initSize] is just
// a hint to prevent unnecessary resizing.
func NewUnboundedDeque[T any](initSize int) Deque[T] { _ = "STUB: not implemented"; return nil }

// Note that [initSize] must be >= 2 to satisfy invariants (1) and (2).

// Invariants after each function call and before the first call:
// (1) The next element pushed left will be placed at data[left]
// (2) The next element pushed right will be placed at data[right]
// (3) There are [size] elements in the deque.
type unboundedSliceDeque[T any] struct {
	size, left, right int
	data              []T
}

func (b *unboundedSliceDeque[T]) PushRight(elt T) bool {
	_ = "STUB: not implemented"
	// Invariant (2) says it's safe to place the element without resizing.
	return false
}

func (b *unboundedSliceDeque[T]) PushLeft(elt T) bool {
	_ = "STUB: not implemented"
	// Invariant (1) says it's safe to place the element without resizing.
	return false
}

// Wrap around

func (b *unboundedSliceDeque[T]) PopLeft() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Zero out to prevent memory leak.

func (b *unboundedSliceDeque[T]) PeekLeft() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (b *unboundedSliceDeque[T]) PopRight() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Zero out to prevent memory leak.

// Wrap around

func (b *unboundedSliceDeque[T]) PeekRight() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (b *unboundedSliceDeque[T]) Index(idx int) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (b *unboundedSliceDeque[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *unboundedSliceDeque[T]) List() []T { _ = "STUB: not implemented"; return nil }

// We copied all of the elements from the leftmost element index
// to the end of the underlying slice, but we still haven't copied
// all of the elements, so wrap around and copy the rest.

func (b *unboundedSliceDeque[T]) leftmostEltIdx() int { _ = "STUB: not implemented"; return 0 }

// Wrap around case

// Normal case

func (b *unboundedSliceDeque[T]) rightmostEltIdx() int { _ = "STUB: not implemented"; return 0 }

// Wrap around case

// Normal case

func (b *unboundedSliceDeque[T]) resize() { _ = "STUB: not implemented"; return }
