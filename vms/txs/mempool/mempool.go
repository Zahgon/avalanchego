// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mempool

import (
	"context"
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/setmap"
	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	// MaxTxSize is the maximum number of bytes a transaction can use to be
	// allowed into the mempool.
	MaxTxSize = 64 * units.KiB

	// droppedTxIDsCacheSize is the maximum number of dropped txIDs to cache
	droppedTxIDsCacheSize = 64

	// maxMempoolSize is the maximum number of bytes allowed in the mempool
	maxMempoolSize = 64 * units.MiB
)

var (
	ErrDuplicateTx          = errors.New("duplicate tx")
	ErrTxTooLarge           = errors.New("tx too large")
	ErrMempoolFull          = errors.New("mempool is full")
	ErrConflictsWithOtherTx = errors.New("tx conflicts with other tx")
)

type Tx interface {
	InputIDs() set.Set[ids.ID]
	ID() ids.ID
	Size() int
}

type Metrics interface {
	Update(numTxs, bytesAvailable int)
}

type Mempool[T Tx] interface {
	Add(tx T) error
	Get(txID ids.ID) (T, bool)
	// Remove [txs] and any conflicts of [txs] from the mempool.
	Remove(txs ...T)

	// Peek returns the oldest tx in the mempool.
	Peek() (tx T, exists bool)

	// Iterate iterates over the txs until f returns false
	Iterate(f func(tx T) bool)

	// Note: dropped txs are added to droppedTxIDs but are not evicted from
	// unissued decision/staker txs. This allows previously dropped txs to be
	// possibly reissued.
	MarkDropped(txID ids.ID, reason error)
	GetDropReason(txID ids.ID) error

	// Len returns the number of txs in the mempool.
	Len() int

	// WaitForEvent waits until there is at least one tx in the mempool.
	WaitForEvent(ctx context.Context) (common.Message, error)
}

type mempool[T Tx] struct {
	lock           sync.RWMutex
	cond           *lock.Cond
	unissuedTxs    *linked.Hashmap[ids.ID, T]
	consumedUTXOs  *setmap.SetMap[ids.ID, ids.ID] // TxID -> Consumed UTXOs
	bytesAvailable int
	droppedTxIDs   *lru.Cache[ids.ID, error] // TxID -> Verification error

	metrics Metrics
}

func New[T Tx](
	metrics Metrics,
) *mempool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (m *mempool[T]) updateMetrics() { _ = "STUB: not implemented"; return }

func (m *mempool[T]) Add(tx T) error { _ = "STUB: not implemented"; return nil }

// Mark these UTXOs as consumed in the mempool

// An added tx must not be marked as dropped.

func (m *mempool[T]) Get(txID ids.ID) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (m *mempool[T]) Remove(txs ...T) { _ = "STUB: not implemented"; return }

// If the transaction is in the mempool, remove it.

// If the transaction isn't in the mempool, remove any conflicts it has.

func (m *mempool[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (m *mempool[T]) Iterate(f func(T) bool) { _ = "STUB: not implemented"; return }

func (m *mempool[_]) MarkDropped(txID ids.ID, reason error) { _ = "STUB: not implemented"; return }

func (m *mempool[_]) GetDropReason(txID ids.ID) error { _ = "STUB: not implemented"; return nil }

func (m *mempool[_]) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *mempool[_]) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}
