// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package iterator

var _ Iterator[any] = Empty[any]{}

// Empty is an iterator with no elements.
type Empty[T any] struct{}

func (Empty[_]) Next() bool { _ = "STUB: not implemented"; return false }

func (Empty[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (Empty[_]) Release() { _ = "STUB: not implemented"; return }
