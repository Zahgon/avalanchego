// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txpool

import (
	"sync"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

const discardedTxsCacheSize = 50

// Txs stores the transactions inside of the mempool.
//
// Transactions in the mempool can be in 1 of 4 statuses:
//
//   - Pending: Pending transactions are eligible for the block builder to
//     attempt to include in the next block being built.
//   - Current: Current transactions are included inside of a block currently
//     being built.
//   - Issued: Issued transactions were included inside of a block built by this
//     node.
//   - Discarded: Discarded transactions were previously in the the mempool, but
//     were then deemed to be invalid. To prevent additional future work, these
//     transactions may be assumed to be invalid in the future.
type Txs struct {
	ctx     *snow.Context
	metrics *metrics
	// maxSize is the maximum number of transactions allowed to be kept in
	// mempool.
	maxSize int
	// pending is a channel of length one, which the mempool uses to awake the
	// block builder when a new transaction is added.
	pending chan struct{}

	lock sync.RWMutex

	// pendingTxs is the collection of transactions available to be included
	// into a block, sorted by gasPrice.
	pendingTxs *txHeap
	// currentTxs is the set of transactions that have been included into a
	// block that is currently being built.
	currentTxs map[ids.ID]*atomic.Tx
	// issuedTxs is the set of transactions that have been included into a
	// block that was previously built.
	issuedTxs map[ids.ID]*atomic.Tx

	// utxoSpenders maps utxoIDs to the Pending, Current, or Issued transaction
	// consuming them in the mempool. This map does not store references to
	// Discarded transactions.
	utxoSpenders map[ids.ID]*atomic.Tx

	// discardedTxs is an LRU Cache of transactions that have been discarded
	// after failing verification.
	discardedTxs *lru.Cache[ids.ID, *atomic.Tx]
}

func NewTxs(ctx *snow.Context, maxSize int) *Txs { _ = "STUB: not implemented"; return nil }

// PendingLen returns the number of pending transactions.
func (t *Txs) PendingLen() int { _ = "STUB: not implemented"; return 0 }

// Iterate applies f to all Pending transactions. If f returns false, the
// iteration stops early.
func (t *Txs) Iterate(f func(tx *atomic.Tx) bool) { _ = "STUB: not implemented"; return }

// NextTx returns the highest paying Pending transaction from the mempool and
// marks it as Current.
func (t *Txs) NextTx() (*atomic.Tx, bool) { _ = "STUB: not implemented"; return nil, false }

// GetPendingTx returns the transaction if it is Pending.
func (t *Txs) GetPendingTx(txID ids.ID) (*atomic.Tx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetTx returns the transaction along with if it is Discarded.
func (t *Txs) GetTx(txID ids.ID) (tx *atomic.Tx, discarded bool, found bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

// Has returns true if the mempool contains the transaction in either the
// Pending, Current, or Issued state.
func (t *Txs) Has(txID ids.ID) bool { _ = "STUB: not implemented"; return false }

// IssueCurrentTxs marks all Current transactions as Issued.
func (t *Txs) IssueCurrentTxs() { _ = "STUB: not implemented"; return }

// CancelCurrentTx attempts to mark the Current transaction as Pending.
//
// This should be called after [Txs.NextTx] returned the transaction and it
// couldn't be included in the block, but should not be discarded. For example,
// CancelCurrentTx should be called if including the transaction will put the
// block above the atomic tx gas limit.
func (t *Txs) CancelCurrentTx(txID ids.ID) { _ = "STUB: not implemented"; return }

// CancelCurrentTxs attempts to mark all Current transactions as Pending.
//
// This should be called after building a block failed due to an error unrelated
// to the transactions.
func (t *Txs) CancelCurrentTxs() { _ = "STUB: not implemented"; return }

// cancelTx attempts to mark the Current transaction as Pending. If the tx can
// not be marked as Pending, it will be Discarded.
//
// Assumes the lock is held.
func (t *Txs) cancelTx(tx *atomic.Tx) { _ = "STUB: not implemented"; return }

// Should never error to calculate the gas price of a transaction already in
// the mempool

// DiscardCurrentTx marks the Current transaction as Discarded.
//
// This should be called after [Txs.NextTx] returned the transaction and it
// failed verification. For example, DiscardCurrentTx should be called if
// including the transaction would produce a conflict with an ancestor block.
func (t *Txs) DiscardCurrentTx(txID ids.ID) { _ = "STUB: not implemented"; return }

// DiscardCurrentTxs marks all Current transactions as Discarded.
//
// This should be called after building a block failed due to an error related
// to the transactions.
func (t *Txs) DiscardCurrentTxs() { _ = "STUB: not implemented"; return }

// discardCurrentTx marks the Current transaction as Discarded.
//
// Assumes the lock is held.
func (t *Txs) discardCurrentTx(tx *atomic.Tx) { _ = "STUB: not implemented"; return }

// removeTx removes the transaction from the mempool. If discard is set, the
// transaction will be kept as Discarded, otherwise it will no longer have any
// status.
//
// Note: removeTx will delete all UTXO entries from utxoSpenders. This means
// that when replacing a conflicting tx, removeTx must be called for all
// conflicts before overwriting the utxoSpenders map.
//
// Assumes lock is held.
func (t *Txs) removeTx(tx *atomic.Tx, discard bool) { _ = "STUB: not implemented"; return }

// removeSpenders deletes the entries for all input UTXOs of the transaction
// from the utxoSpenders map.
//
// Assumes the lock is held.
func (t *Txs) removeSpenders(tx *atomic.Tx) { _ = "STUB: not implemented"; return }

// RemoveTx removes the transaction from the mempool, including removal of the
// Discarded status.
func (t *Txs) RemoveTx(tx *atomic.Tx) { _ = "STUB: not implemented"; return }

// SubscribePendingTxs returns a channel that signals when there is a
// transaction added to the mempool.
func (t *Txs) SubscribePendingTxs() <-chan struct{} { _ = "STUB: not implemented"; return nil }
