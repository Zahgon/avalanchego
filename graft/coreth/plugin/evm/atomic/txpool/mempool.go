// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txpool

import (
	"errors"

	"github.com/holiman/uint256"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/utils/bloom"
)

var (
	_ gossip.SystemSet[*atomic.Tx] = (*Mempool)(nil)

	ErrAlreadyKnown    = errors.New("already known")
	ErrConflict        = errors.New("conflict present")
	ErrInsufficientFee = errors.New("insufficient fee")
	ErrMempoolFull     = errors.New("mempool full")

	// If the transaction is already in the mempool, marking it as Discarded,
	// could unexpectedly cause the transaction to have multiple statuses.
	//
	// If the mempool is full, that is not the transaction's fault, so we should
	// not prevent adding the transaction to the mempool later.
	errsNotToDiscard = []error{
		ErrAlreadyKnown,
		ErrMempoolFull,
	}
)

// Mempool is a simple mempool for atomic transactions
type Mempool struct {
	*Txs
	// bloom is a bloom filter containing the txs in the mempool
	bloom  *gossip.BloomFilter
	verify func(tx *atomic.Tx) error
}

func NewMempool(
	txs *Txs,
	registerer prometheus.Registerer,
	verify func(tx *atomic.Tx) error,
) (*Mempool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add attempts to add tx to the mempool as a Remote transaction. It is assumed
// the snow context lock is not held.
func (m *Mempool) Add(tx *atomic.Tx) error { _ = "STUB: not implemented"; return nil }

// AddRemoteTx attempts to add tx to the mempool as a Remote transaction.
//
// Remote transactions are checked for recent verification failures prior to
// performing verification. If a Remote transaction failed verification recently
// it will not be added to the mempool.
func (m *Mempool) AddRemoteTx(tx *atomic.Tx) error { _ = "STUB: not implemented"; return nil }

// Unlike local txs, invalid remote txs are recorded as discarded so that
// they won't be requested again

// AddLocalTx attempts to add tx to the mempool as a Local transaction.
//
// Local transactions are not checked for recent verification failures prior to
// performing verification. Even if a Local transaction failed verification
// recently, the mempool will attempt to re-verify it.
func (m *Mempool) AddLocalTx(tx *atomic.Tx) error { _ = "STUB: not implemented"; return nil }

// ForceAddTx forcibly adds a tx to the mempool and bypasses all verification.
func (m *Mempool) ForceAddTx(tx *atomic.Tx) error { _ = "STUB: not implemented"; return nil }

// checkConflictTx checks for any transactions in the mempool that spend the
// same input UTXOs as the provided transaction. If any conflicts are present,
// it returns the highest gas price of any conflicting transaction, the ID of
// the corresponding tx and the full list of conflicting transactions.
func (m *Mempool) checkConflictTx(tx *atomic.Tx) (uint256.Int, ids.ID, []*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return *new(uint256.Int), *new(ids.ID), nil, nil
}

// We don't know the length ahead of time (# of conflicts), so we don't want to preallocate

// Get current gas price of the existing tx in the mempool

// Should never error to calculate the gas price of a transaction
// already in the mempool

// Assumes the lock is held.
func (m *Mempool) length() int { _ = "STUB: not implemented"; return 0 }

// addTx attempts to add tx to the mempool.
//
// Unless local, discarded transactions can not be added.
// If force, conflict checks are skipped.
//
// Assumes lock is held.
func (m *Mempool) addTx(tx *atomic.Tx, local bool, force bool) error {
	_ = "STUB: not implemented"
	// If the tx has already been included in the mempool, there's no need to
	// add it again.
	return nil
}

// If the transaction does not have a higher fee than all of its
// conflicts, we refuse to add it to the mempool. (Prefer items already
// in the mempool).

// Remove all conflicting transactions from the mempool.

// If adding this transaction would exceed the mempool's size, check if
// there is a lower priced transaction that can be evicted from the mempool.
//
// Recall that the size of the mempool is denominated in transactions, so to
// add a new transaction we only need to evict at most one other
// transaction.

// This could occur if we have used our entire size allowance on
// transactions that are either Current or Issued.

// Get the transaction with the lowest gasPrice

// If the new tx doesn't have a higher fee than the transaction it
// would replace, discard new transaction. (Prefer items already
// in the mempool).

// If the transaction was recently discarded, log the event and evict from
// discarded transactions so it's not in two places within the mempool.
// We allow the transaction to be re-issued since it may have been invalid
// due to an atomic UTXO not being present yet.

// Mark the transaction as Pending.

// Current transactions must be added to the bloom filter as well
// because they could be added back into the pending set without going
// through addTx again.

// When adding a transaction to the mempool, we make sure that there is an
// item in Pending to signal the VM to produce a block.

func (m *Mempool) BloomFilter() (*bloom.Filter, ids.ID) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID)
}
