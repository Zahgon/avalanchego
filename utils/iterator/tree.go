// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package iterator

import (
	"sync"

	"github.com/google/btree"
)

var _ Iterator[any] = (*tree[any])(nil)

type tree[T any] struct {
	current     T
	next        chan T
	releaseOnce sync.Once
	release     chan struct{}
	wg          sync.WaitGroup
}

// FromTree returns a new iterator of the stakers in [tree] in ascending order.
// Note that it isn't safe to modify [tree] while iterating over it.
func FromTree[T any](btree *btree.BTreeG[T]) Iterator[T] { _ = "STUB: not implemented"; return nil }

func (i *tree[_]) Next() bool { _ = "STUB: not implemented"; return false }

func (i *tree[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (i *tree[_]) Release() { _ = "STUB: not implemented"; return }
