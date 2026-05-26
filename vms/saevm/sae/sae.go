// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package sae implements the [Streaming Asynchronous Execution] (SAE) virtual
// machine to be compatible with Avalanche consensus.
//
// [Streaming Asynchronous Execution]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/194-streaming-asynchronous-execution
package sae

import (
	"math/big"
	"sync"
	"time"

	"github.com/holiman/uint256"
)

func unix(t time.Time) uint64 { _ = "STUB: not implemented"; return 0 }

//#nosec G115 -- Guaranteed to be positive

// uint256FromBig is a wrapper around [uint256.FromBig] with extra checks, for
// nil input and for overflow.
func uint256FromBig(b *big.Int) (*uint256.Int, error) { _ = "STUB: not implemented"; return nil, nil }

type syncMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex

	onStore  func(V)
	onDelete func(V)
}

// newSyncMap creates a concurrent-safe map, which automatically performs
// `onStore` and `onDelete` if [syncMap.Store] and [syncMap.Delete] are called,
// respectively. If either function is nil, or the key to be deleted doesn't
// exist, no operation will be performed.
func newSyncMap[K comparable, V any](onStore func(V), onDelete func(V)) *syncMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *syncMap[K, V]) Load(k K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m *syncMap[K, V]) Store(k K, v V) { _ = "STUB: not implemented"; return }

func (m *syncMap[K, V]) Delete(k K) { _ = "STUB: not implemented"; return }
