// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txheap

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/heap"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

type Heap interface {
	Add(tx *txs.Tx)
	Get(txID ids.ID) *txs.Tx
	List() []*txs.Tx
	Remove(txID ids.ID) *txs.Tx
	Peek() *txs.Tx
	RemoveTop() *txs.Tx
	Len() int
}

type txHeap struct {
	heap       heap.Map[ids.ID, *txs.Tx]
	currentAge int
}

func (h *txHeap) Add(tx *txs.Tx) { _ = "STUB: not implemented"; return }

func (h *txHeap) Get(txID ids.ID) *txs.Tx { _ = "STUB: not implemented"; return nil }

func (h *txHeap) List() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (h *txHeap) Remove(txID ids.ID) *txs.Tx { _ = "STUB: not implemented"; return nil }

func (h *txHeap) Peek() *txs.Tx { _ = "STUB: not implemented"; return nil }

func (h *txHeap) RemoveTop() *txs.Tx { _ = "STUB: not implemented"; return nil }

func (h *txHeap) Len() int { _ = "STUB: not implemented"; return 0 }
