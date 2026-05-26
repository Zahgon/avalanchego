// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package heap

// NewSet returns a heap without duplicates ordered by its values
func NewSet[T comparable](less func(a, b T) bool) Set[T] { _ = "STUB: not implemented"; return nil }

type Set[T comparable] struct {
	set Map[T, T]
}

// Push returns if the entry was added
func (s Set[T]) Push(t T) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) Pop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s Set[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s Set[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Set[T]) Remove(t T) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) Fix(t T) { _ = "STUB: not implemented"; return }

func (s Set[T]) Contains(t T) bool { _ = "STUB: not implemented"; return false }
