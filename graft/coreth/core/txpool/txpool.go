// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package txpool

import (
	"errors"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

var (
	// ErrOverdraft is returned if a transaction would cause the senders balance to go negative
	// thus invalidating a potential large number of transactions.
	ErrOverdraft = errors.New("transaction would cause overdraft")
)

// TxStatus is the current status of a transaction as seen by the pool.
type TxStatus uint

const (
	TxStatusUnknown TxStatus = iota
	TxStatusQueued
	TxStatusPending
)

var (
	// reservationsGaugeName is the prefix of a per-subpool address reservation
	// metric.
	//
	// This is mostly a sanity metric to ensure there's no bug that would make
	// some subpool hog all the reservations due to mis-accounting.
	reservationsGaugeName = "txpool/reservations"
)

// BlockChain defines the minimal set of methods needed to back a tx pool with
// a chain. Exists to allow mocking the live chain out of tests.
type BlockChain interface {
	// CurrentBlock returns the current head of the chain.
	CurrentBlock() *types.Header

	// SubscribeChainHeadEvent subscribes to new blocks being added to the chain.
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
}

// TxPool is an aggregator for various transaction specific pools, collectively
// tracking all the transactions deemed interesting by the node. Transactions
// enter the pool when they are received from the network or submitted locally.
// They exit the pool when they are included in the blockchain or evicted due to
// resource constraints.
type TxPool struct {
	subpools []SubPool // List of subpools for specialized transaction handling

	reservations map[common.Address]SubPool // Map with the account to pool reservations
	reserveLock  sync.Mutex                 // Lock protecting the account reservations

	subs event.SubscriptionScope // Subscription scope to unsubscribe all on shutdown
	quit chan chan error         // Quit channel to tear down the head updater
	term chan struct{}           // Termination channel to detect a closed pool

	sync chan chan error // Testing / simulator channel to block until internal reset is done

	gasTip    atomic.Pointer[big.Int] // Remember last value set so it can be retrieved
	minFee    atomic.Pointer[big.Int] // Remember last value set so it can be retrieved (in tests)
	reorgFeed event.Feed
}

// New creates a new transaction pool to gather, sort and filter inbound
// transactions from the network.
func New(gasTip uint64, chain BlockChain, subpools []SubPool) (*TxPool, error) {
	_ = "STUB: not implemented"
	// Retrieve the current head so that all subpools and this main coordinator
	// pool will have the same starting state, even if the chain moves forward
	// during initialization.
	return nil, nil
}

// Subscribe to chain head events to trigger subpool resets

// reserver is a method to create an address reservation callback to exclusively
// assign/deassign addresses to/from subpools. This can ensure that at any point
// in time, only a single subpool is able to manage an account, avoiding cross
// subpool eviction issues and nonce conflicts.
func (p *TxPool) reserver(id int, subpool SubPool) AddressReserver {
	_ = "STUB: not implemented"
	return *new(AddressReserver)
}

// Double reservations are forbidden even from the same pool to
// avoid subtle bugs in the long term.

// Ignore fault to give the pool a chance to recover while the bug gets fixed

// Ensure subpools only attempt to unreserve their own owned addresses,
// otherwise flag as a programming error.

// Close terminates the transaction pool and all its subpools.
func (p *TxPool) Close() error { _ = "STUB: not implemented"; return nil }

// Terminate the reset loop and wait for it to finish

// Terminate each subpool

// Unsubscribe anyone still listening for tx events

// loop is the transaction pool's main event loop, waiting for and reacting to
// outside blockchain events as well as for various reporting and transaction
// eviction events.
func (p *TxPool) loop(head *types.Header, newHeadCh <-chan core.ChainHeadEvent) {
	_ = "STUB: not implemented"
	// Close the termination marker when the pool stops
	return
}

// Track the previous and current head to feed to an idle reset

// Consume chain head events and start resets when none is running

// Allow 1 reset to run concurrently

// Whether a forced reset was requested, only used in simulator mode
// Channel waiting on a forced reset, only used in simulator mode

// Notify the live reset waiter to not block if the txpool is closed.

// Something interesting might have happened, run a reset if there is
// one needed but none is running. The resetter will run on its own
// goroutine to allow chain head events to be consumed contiguously.

// Try to inject a busy marker and start a reset if successful

// Busy marker injected, start a new subpool reset

// If the reset operation was explicitly requested, consider it
// being fulfilled and drop the request marker. If it was not,
// this is a noop.

// Reset already running, wait until it finishes.
//
// Note, this will not drop any forced reset request. If a forced
// reset was requested, but we were busy, then when the currently
// running reset finishes, a new one will be spun up.

// Wait for the next chain head event or a previous reset finish

// Chain moved forward, store the head for later consumption

// Previous reset finished, update the old head and allow a new reset

// If someone is waiting for a reset to finish, notify them, unless
// the forced op is still pending. In that case, wait another round
// of resets.

// Termination requested, break out on the next loop round

// Transaction pool is running inside a simulator, and we are about
// to create a new block. Request a forced sync operation to ensure
// that any running reset operation finishes to make block imports
// deterministic. On top of that, run a new reset operation to make
// transaction insertions deterministic instead of being stuck in a
// queue waiting for a reset.

// Notify the closer of termination (no error possible for now)

// GasTip returns the current gas tip enforced by the transaction pool.
func (p *TxPool) GasTip() *big.Int { _ = "STUB: not implemented"; return nil }

// SetGasTip updates the minimum gas tip required by the transaction pool for a
// new transaction, and drops all transactions below this threshold.
func (p *TxPool) SetGasTip(tip *big.Int) { _ = "STUB: not implemented"; return }

// MinFee returns the current minimum fee enforced by the transaction pool.
func (p *TxPool) MinFee() *big.Int { _ = "STUB: not implemented"; return nil }

// SetMinFee updates the minimum fee required by the transaction pool for a
// new transaction, and drops all transactions below this threshold.
func (p *TxPool) SetMinFee(fee *big.Int) { _ = "STUB: not implemented"; return }

// Has returns an indicator whether the pool has a transaction cached with the
// given hash.
func (p *TxPool) Has(hash common.Hash) bool { _ = "STUB: not implemented"; return false }

// HasLocal returns an indicator whether the pool has a local transaction cached
// with the given hash.
func (p *TxPool) HasLocal(hash common.Hash) bool { _ = "STUB: not implemented"; return false }

// Get returns a transaction if it is contained in the pool, or nil otherwise.
func (p *TxPool) Get(hash common.Hash) *types.Transaction { _ = "STUB: not implemented"; return nil }

// Add enqueues a batch of transactions into the pool if they are valid. Due
// to the large transaction churn, add may postpone fully integrating the tx
// to a later point to batch multiple ones together.
func (p *TxPool) Add(txs []*types.Transaction, local bool, sync bool) []error {
	_ = "STUB: not implemented"
	// Split the input transactions between the subpools. It shouldn't really
	// happen that we receive merged batches, but better graceful than strange
	// errors.
	//
	// We also need to track how the transactions were split across the subpools,
	// so we can piece back the returned errors into the original order.
	return nil
}

// Mark this transaction belonging to no-subpool

// Try to find a subpool that accepts the transaction

// Add the transactions split apart to the individual subpools and piece
// back the errors into the original sort order.

// If the transaction was rejected by all subpools, mark it unsupported

// Find which subpool handled it and pull in the corresponding error

func (p *TxPool) AddRemotesSync(txs []*types.Transaction) []error {
	_ = "STUB: not implemented"
	return nil
}

// Pending retrieves all currently processable transactions, grouped by origin
// account and sorted by nonce. The returned transaction set is a copy and can be
// freely modified by calling code.
//
// The transactions can also be pre-filtered by the dynamic fee components to
// reduce allocations and load on downstream subsystems.
func (p *TxPool) Pending(filter PendingFilter) map[common.Address][]*LazyTransaction {
	_ = "STUB: not implemented"
	return nil
}

// PendingSize returns the number of pending txs in the tx pool.
//
// The filter parameter can be used to do an extra filtering on the pending
// transactions.
func (p *TxPool) PendingSize(filter PendingFilter) int { _ = "STUB: not implemented"; return 0 }

// IteratePending iterates over [pool.pending] until [f] returns false.
// The caller must not modify [tx].
func (p *TxPool) IteratePending(f func(tx *types.Transaction) bool) {
	_ = "STUB: not implemented"
	return
}

// SubscribeTransactions registers a subscription for new transaction events,
// supporting feeding only newly seen or also resurrected transactions.
func (p *TxPool) SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeNewReorgEvent registers a subscription of NewReorgEvent and
// starts sending event to the given channel.
func (p *TxPool) SubscribeNewReorgEvent(ch chan<- core.NewTxPoolReorgEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// Nonce returns the next nonce of an account, with all transactions executable
// by the pool already applied on top.
func (p *TxPool) Nonce(addr common.Address) uint64 {
	_ = "STUB: not implemented"
	// Since (for now) accounts are unique to subpools, only one pool will have
	// (at max) a non-state nonce. To avoid stateful lookups, just return the
	// highest nonce for now.
	return 0
}

// Stats retrieves the current pool stats, namely the number of pending and the
// number of queued (non-executable) transactions.
func (p *TxPool) Stats() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Content retrieves the data content of the transaction pool, returning all the
// pending as well as queued transactions, grouped by account and sorted by nonce.
func (p *TxPool) Content() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContentFrom retrieves the data content of the transaction pool, returning the
// pending as well as queued transactions of this address, grouped by nonce.
func (p *TxPool) ContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Locals retrieves the accounts currently considered local by the pool.
func (p *TxPool) Locals() []common.Address {
	_ = "STUB: not implemented"
	// Retrieve the locals from each subpool and deduplicate them
	return nil
}

// Flatten and return the deduplicated local set

// Status returns the known status (unknown/pending/queued) of a transaction
// identified by its hash.
func (p *TxPool) Status(hash common.Hash) TxStatus {
	_ = "STUB: not implemented"
	return *new(TxStatus)
}

// Sync is a helper method for unit tests or simulator runs where the chain events
// are arriving in quick succession, without any time in between them to run the
// internal background reset operations. This method will run an explicit reset
// operation to ensure the pool stabilises, thus avoiding flakey behavior.
//
// Note, do not use this in production / live code. In live code, the pool is
// meant to reset on a separate thread to avoid DoS vectors.
func (p *TxPool) Sync() error { _ = "STUB: not implemented"; return nil }
