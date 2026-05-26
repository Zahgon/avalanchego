// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cache

var _ Cacher[struct{}, struct{}] = (*Empty[struct{}, struct{}])(nil)

// Empty is a cache that doesn't store anything.
type Empty[K any, V any] struct{}

func (*Empty[K, V]) Put(K, V) { _ = "STUB: not implemented"; return }

func (*Empty[K, V]) Get(K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (*Empty[K, _]) Evict(K) { _ = "STUB: not implemented"; return }

func (*Empty[_, _]) Flush() { _ = "STUB: not implemented"; return }

func (*Empty[_, _]) Len() int { _ = "STUB: not implemented"; return 0 }

func (*Empty[_, _]) PortionFilled() float64 { _ = "STUB: not implemented"; return 0 }
