// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"encoding/json"
	"sync"
)

var (
	_ json.Marshaler   = (*Atomic[struct{}])(nil)
	_ json.Unmarshaler = (*Atomic[struct{}])(nil)
)

type Atomic[T any] struct {
	lock  sync.RWMutex
	value T
}

func NewAtomic[T any](value T) *Atomic[T] { _ = "STUB: not implemented"; return nil }

func (a *Atomic[T]) Get() T { _ = "STUB: not implemented"; return *new(T) }

func (a *Atomic[T]) Set(value T) { _ = "STUB: not implemented"; return }

func (a *Atomic[T]) Swap(value T) T { _ = "STUB: not implemented"; return *new(T) }

func (a *Atomic[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Atomic[T]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
