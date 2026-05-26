// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package iterator

var _ Iterator[any] = (*slice[any])(nil)

// ToSlice returns a slice that contains all of the elements from [it] in order.
// [it] will be released before returning.
func ToSlice[T any](it Iterator[T]) []T { _ = "STUB: not implemented"; return nil }

type slice[T any] struct {
	index    int
	elements []T
}

// FromSlice returns an iterator that contains [elements] in order. Doesn't sort
// by anything.
func FromSlice[T any](elements ...T) Iterator[T] { _ = "STUB: not implemented"; return nil }

func (i *slice[_]) Next() bool { _ = "STUB: not implemented"; return false }

func (i *slice[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (*slice[_]) Release() { _ = "STUB: not implemented"; return }
