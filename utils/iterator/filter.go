// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package iterator

var _ Iterator[any] = (*filtered[any])(nil)

type filtered[T any] struct {
	it     Iterator[T]
	filter func(T) bool
}

// Filter returns an iterator that skips the elements in [it] that return true
// from [filter].
func Filter[T any](it Iterator[T], filter func(T) bool) Iterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Deduplicate returns an iterator that skips the elements that have already
// been returned from [it].
func Deduplicate[T comparable](it Iterator[T]) Iterator[T] { _ = "STUB: not implemented"; return nil }

func (i *filtered[_]) Next() bool { _ = "STUB: not implemented"; return false }

func (i *filtered[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (i *filtered[_]) Release() { _ = "STUB: not implemented"; return }
