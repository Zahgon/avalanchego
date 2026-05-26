// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package buffer

import (
	"sync"
)

var _ BlockingDeque[int] = (*UnboundedBlockingDeque[int])(nil)

type BlockingDeque[T any] interface {
	Deque[T]

	// Close and empty the deque.
	Close()
}

// Returns a new unbounded deque with the given initial size.
// Note that the returned deque is always empty -- [initSize] is just
// a hint to prevent unnecessary resizing.
func NewUnboundedBlockingDeque[T any](initSize int) *UnboundedBlockingDeque[T] {
	_ = "STUB: not implemented"
	return nil
}

// UnboundedBlockingDeque is a thread-safe blocking deque with unbounded growth.
type UnboundedBlockingDeque[T any] struct {
	lock   sync.RWMutex
	cond   *sync.Cond
	closed bool

	Deque[T]
}

// If the deque is closed returns false.
func (q *UnboundedBlockingDeque[T]) PushRight(elt T) bool { _ = "STUB: not implemented"; return false }

// Add the item to the queue

// Signal a waiting thread

// If the deque is closed returns false.
func (q *UnboundedBlockingDeque[T]) PopRight() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (q *UnboundedBlockingDeque[T]) PeekRight() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// If the deque is closed returns false.
func (q *UnboundedBlockingDeque[T]) PushLeft(elt T) bool { _ = "STUB: not implemented"; return false }

// Add the item to the queue

// Signal a waiting thread

// If the deque is closed returns false.
func (q *UnboundedBlockingDeque[T]) PopLeft() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (q *UnboundedBlockingDeque[T]) PeekLeft() (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (q *UnboundedBlockingDeque[T]) Index(i int) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (q *UnboundedBlockingDeque[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *UnboundedBlockingDeque[T]) List() []T { _ = "STUB: not implemented"; return nil }

func (q *UnboundedBlockingDeque[T]) Close() { _ = "STUB: not implemented"; return }

// Mark the queue as closed
