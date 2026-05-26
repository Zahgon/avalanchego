// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txpool

import (
	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
)

type txEntry struct {
	id       ids.ID
	gasPrice uint256.Int
	tx       *atomic.Tx
	index    int
}

// internalTxHeap is used to track pending atomic transactions by gasPrice
type internalTxHeap struct {
	isMinHeap bool
	items     []*txEntry
	lookup    map[ids.ID]*txEntry
}

func newInternalTxHeap(items int, isMinHeap bool) *internalTxHeap {
	_ = "STUB: not implemented"
	return nil
}

func (th internalTxHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (th internalTxHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (th internalTxHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (th *internalTxHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (th *internalTxHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// avoid memory leak

func (th *internalTxHeap) Get(id ids.ID) (*txEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (th *internalTxHeap) Has(id ids.ID) bool { _ = "STUB: not implemented"; return false }

type txHeap struct {
	maxHeap *internalTxHeap
	minHeap *internalTxHeap
}

func newTxHeap(maxSize int) *txHeap { _ = "STUB: not implemented"; return nil }

func (th *txHeap) Push(tx *atomic.Tx, gasPrice uint256.Int) { _ = "STUB: not implemented"; return }

// Assumes there is non-zero items
func (th *txHeap) PeekMax() (*atomic.Tx, uint256.Int) {
	_ = "STUB: not implemented"
	return nil, *new(uint256.Int)
}

// Assumes there is non-zero items
func (th *txHeap) PeekMin() (*atomic.Tx, uint256.Int) {
	_ = "STUB: not implemented"
	return nil, *new(uint256.Int)
}

// Assumes there is non-zero items
func (th *txHeap) PopMax() *atomic.Tx { _ = "STUB: not implemented"; return nil }

// Assumes there is non-zero items
func (th *txHeap) PopMin() *atomic.Tx { _ = "STUB: not implemented"; return nil }

func (th *txHeap) Remove(id ids.ID) *atomic.Tx { _ = "STUB: not implemented"; return nil }

// This should never happen, as that would mean the heaps are out of
// sync.

func (th *txHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (th *txHeap) Get(id ids.ID) (*atomic.Tx, bool) { _ = "STUB: not implemented"; return nil, false }

func (th *txHeap) Has(id ids.ID) bool { _ = "STUB: not implemented"; return false }
