// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package heap

import (
	"container/heap"
)

var _ heap.Interface = (*queue[int])(nil)

// NewQueue returns an empty heap. See QueueOf for more.
func NewQueue[T any](less func(a, b T) bool) Queue[T] { _ = "STUB: not implemented"; return nil }

// QueueOf returns a heap containing entries ordered by less.
func QueueOf[T any](less func(a, b T) bool, entries ...T) Queue[T] {
	_ = "STUB: not implemented"
	return nil
}

type Queue[T any] struct {
	queue *queue[T]
}

func (q *Queue[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *Queue[T]) Push(t T) { _ = "STUB: not implemented"; return }

func (q *Queue[T]) Pop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *Queue[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *Queue[T]) Fix(i int) { _ = "STUB: not implemented"; return }

type queue[T any] struct {
	entries []T
	less    func(a, b T) bool
}

func (q *queue[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *queue[T]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (q *queue[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (q *queue[T]) Push(e any) { _ = "STUB: not implemented"; return }

func (q *queue[T]) Pop() any { _ = "STUB: not implemented"; return *new(any) }
