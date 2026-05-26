// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mempool

import (
	"context"
	"errors"
	"sync"

	"github.com/google/btree"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/setmap"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	ErrNotEnoughGas = errors.New("not enough gas")
	errNoGasUsed    = errors.New("no gas used")
	errAVAXMinted   = errors.New("AVAX minted")
)

type meteredTx struct {
	*txs.Tx
	// gasPrice is the amount of AVAX burned per unit of gas used by this tx
	gasPrice float64
	gasUsed  gas.Gas
}

type Mempool struct {
	weights     gas.Dimensions
	avaxAssetID ids.ID

	lock               sync.RWMutex
	cond               *lock.Cond
	tree               *btree.BTreeG[meteredTx]
	txs                map[ids.ID]meteredTx
	consumedUTXOs      *setmap.SetMap[ids.ID, ids.ID]
	droppedTxIDs       *lru.Cache[ids.ID, error]
	gasAvailable       gas.Gas
	numTxsMetric       prometheus.Gauge
	gasAvailableMetric prometheus.Gauge
}

func New(
	namespace string,
	weights gas.Dimensions,
	gasCapacity gas.Gas,
	avaxAssetID ids.ID,
	registerer prometheus.Registerer,
) (*Mempool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Break ties with txID

// Add adds `tx` to the mempool and clears its dropped status.
func (m *Mempool) Add(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

// Try to evict lower gas priced txs if we do not have enough remaining gas
// capacity

// Try to evict transactions until there is enough capacity for the tx.
// Returns if we are able to fit this tx.
func (m *Mempool) allocateSpace(txToAdd meteredTx) bool {
	_ = "STUB: not implemented"
	// We have enough space for this tx
	return false
}

// Try to evict lower priced txs to make room for the new tx

// We do not have enough space for this tx

func (m *Mempool) meter(tx *txs.Tx) (meteredTx, error) {
	_ = "STUB: not implemented"
	return *new(meteredTx), nil
}

// The caller should verify txs but perform overflow checks anyway

// The caller should verify txs but perform overflow checks anyway

// The caller should verify txs but perform this check anyway

func (m *Mempool) updateMetrics() { _ = "STUB: not implemented"; return }

// Get returns the tx corresponding to `txID` and if it was present
func (m *Mempool) Get(txID ids.ID) (*txs.Tx, bool) { _ = "STUB: not implemented"; return nil, false }

// Remove removes `txID` from the mempool
func (m *Mempool) Remove(txID ids.ID) { _ = "STUB: not implemented"; return }

func (m *Mempool) remove(txID ids.ID) { _ = "STUB: not implemented"; return }

// RemoveConflicts removes all txs conflicting with `utxos`
func (m *Mempool) RemoveConflicts(utxos set.Set[ids.ID]) { _ = "STUB: not implemented"; return }

// Peek returns a tx in the mempool and if it was present
func (m *Mempool) Peek() (*txs.Tx, bool) { _ = "STUB: not implemented"; return nil, false }

// Iterate calls `f` over each tx in the mempool
func (m *Mempool) Iterate(f func(tx *txs.Tx) bool) { _ = "STUB: not implemented"; return }

// MarkDropped marks `txID` as dropped
func (m *Mempool) MarkDropped(txID ids.ID, reason error) { _ = "STUB: not implemented"; return }

// GetDropReason returns why `txID` was dropped
func (m *Mempool) GetDropReason(txID ids.ID) error { _ = "STUB: not implemented"; return nil }

// Len returns the number of txs in the mempool
func (m *Mempool) Len() int { _ = "STUB: not implemented"; return 0 }

// WaitForEvent blocks until the mempool has txs that are ready to build into
// a block.
func (m *Mempool) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

// TODO block until the mempool has a gas price greater than or equal to the
// chain's minimum gas price
