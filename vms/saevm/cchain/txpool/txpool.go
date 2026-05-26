// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package txpool implements an in-memory pool of cross-chain transactions
// awaiting inclusion in a block.
package txpool

import (
	"context"
	"errors"
	"iter"
	"sync"

	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/libevm"

	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/heap"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/setmap"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
)

// Backend that the [Txpool] depends on for current chain state.
type Backend interface {
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	LastExecutedState() (libevm.StateReader, error)
}

// Txpool is an in-memory pool of cross-chain transactions awaiting inclusion
// in a block.
//
// Transactions are admitted only after passing verification against the most
// recently executed state.
//
// Transactions are removed after they are included in an executed block or are
// replaced by a higher paying transaction.
type Txpool struct {
	*Pending

	snowCtx *snow.Context
	sub     event.Subscription
	maxSize int
	wg      sync.WaitGroup

	// stateLock is ordered before [Pending.lock]. Acquiring stateLock with
	// [Pending.lock] held will deadlock.
	stateLock sync.RWMutex
	state     libevm.StateReader
}

// New constructs a [Txpool] that wraps the provided [Pending].
//
// maxSize is the maximum number of transactions the pool will hold; once
// reached, [Txpool.Add] evicts the lowest-fee transaction in favor of a
// strictly higher-fee incoming transaction.
//
// [Txpool.Close] MUST be called during shutdown to release allocated resources.
func New(
	snowCtx *snow.Context,
	chainConfig *params.ChainConfig,
	pending *Pending,
	chain Backend,
	maxSize int,
) (*Txpool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// executed is unbuffered to guarantee that the pool never holds a reference
// to state older than the last-settled state. SAE does not guarantee that
// such a state exists on disk anymore.

// state must be populated after [Backend.SubscribeChainHeadEvent] is called
// to ensure we do not miss an update.

func (p *Txpool) updateState(
	chainConfig *params.ChainConfig,
	chain Backend,
	executed <-chan core.ChainHeadEvent,
) {
	_ = "STUB: not implemented"
	return
}

var (
	// ErrAlreadyKnown is returned by [Txpool.Add] when the transaction is
	// already in the pool.
	ErrAlreadyKnown = errors.New("transaction already in pool")

	errSanityCheck       = errors.New("sanity check")
	errVerifyCredentials = errors.New("credential verification")
	errVerifyState       = errors.New("state verification")
	errInsufficientFee   = errors.New("insufficient fee")
)

// Add validates tx and inserts it into the pool.
//
// If tx conflicts with a transaction already in the pool, the lower-fee
// transaction is evicted. If the pool is at capacity, the lowest-fee
// transaction is evicted in favor of a higher-fee incoming transaction.
//
// Returns [ErrAlreadyKnown] if tx is already in the pool.
func (p *Txpool) Add(tx *tx.Tx) error { _ = "STUB: not implemented"; return nil }

// TODO:(StephenButtolph): Should we enforce a maximum gas amount here?

// We must verify the tx against a state that is at least as high as the
// last block processed by the pool subscription.
//
// Verifying against an older state risks admitting a tx that would never
// be evicted.

// Close releases all allocated resources.
func (p *Txpool) Close() { _ = "STUB: not implemented"; return }

// inputUTXOs returns the union of all UTXO IDs consumed by transactions in b,
// covering both EVM-native account+nonce inputs and cross-chain inputs.
func inputUTXOs(b *types.Block, c *params.ChainConfig) (set.Set[ids.ID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errNonceMismatch     = errors.New("nonce mismatch")
	errInsufficientFunds = errors.New("insufficient funds")
)

// verifyOp verifies that op's debits are valid against state.
func verifyOp(state libevm.StateReader, op hook.Op) error { _ = "STUB: not implemented"; return nil }

// Pending stores transactions that are eligible for inclusion in a future
// block, indexed for fast conflict lookup and ordered by gas price.
type Pending struct {
	lock sync.RWMutex
	cond *lock.Cond

	// txs is the collection of transactions available to be included into a
	// block, ordered as a min-heap by gas price for eviction.
	txs heap.Map[ids.ID, *txData]
	// utxos maps a txID to the set of utxoIDs it consumes.
	utxos *setmap.SetMap[ids.ID, ids.ID]
}

// NewPending constructs an empty set of [Pending] transactions.
func NewPending() *Pending { _ = "STUB: not implemented"; return nil }

// txs is a min-heap

// Iter returns an iterator over the pool's transactions in decreasing gas
// price order.
func (p *Pending) Iter() iter.Seq[*tx.Tx] {
	_ = "STUB: not implemented"

	// TODO:(StephenButtolph): Iteration shouldn't copy the pool.
	return nil
}

// Len returns the number of transactions currently in the pool.
func (p *Pending) Len() int { _ = "STUB: not implemented"; return 0 }

// Has reports whether txID is in the pool.
func (p *Pending) Has(txID ids.ID) bool { _ = "STUB: not implemented"; return false }

// AwaitTxs blocks until at least one transaction is in the pool or ctx is
// cancelled.
func (p *Pending) AwaitTxs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *Pending) removeConflicts(utxos set.Set[ids.ID]) { _ = "STUB: not implemented"; return }

// add inserts t into the pool. It assumes there are no existing conflicts.
func (p *Pending) add(t *txData) { _ = "STUB: not implemented"; return }

// txData contains the values from [tx.Tx] that the pool uses for ordering and
// conflict detection.
type txData struct {
	id     ids.ID
	tx     *tx.Tx
	inputs set.Set[ids.ID]
	op     hook.Op
}

var errAsOp = errors.New("as op")

func newTxData(tx *tx.Tx, avaxAssetID ids.ID) (*txData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
