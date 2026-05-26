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
// Copyright 2014 The go-ethereum Authors
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

// Package legacypool implements the normal EVM execution transaction pool.
package legacypool

import (
	"errors"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core/txpool"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/metrics"
	"github.com/holiman/uint256"

	// Force libevm metrics of the same name to be registered first.
	_ "github.com/ava-labs/libevm/core/txpool/legacypool"
)

const (
	// txSlotSize is used to calculate how many data slots a single transaction
	// takes up based on its size. The slots are used as DoS protection, ensuring
	// that validating a new transaction remains a constant operation (in reality
	// O(maxslots), where max slots are 4 currently).
	txSlotSize = 32 * 1024

	// txMaxSize is the maximum size a single transaction can have. This field has
	// non-trivial consequences: larger transactions are significantly harder and
	// more expensive to propagate; larger transactions also take more resources
	// to validate whether they fit into the pool or not.
	//
	// Note: the max contract size is 24KB
	txMaxSize = 4 * txSlotSize // 128KB
)

var (
	// ErrTxPoolOverflow is returned if the transaction pool is full and can't accept
	// another remote transaction.
	ErrTxPoolOverflow = errors.New("txpool is full")
)

var (
	evictionInterval      = time.Minute      // Time interval to check for evictable transactions
	statsReportInterval   = 8 * time.Second  // Time interval to report transaction pool stats
	baseFeeUpdateInterval = 10 * time.Second // Time interval at which to schedule a base fee update for the tx pool after SubnetEVM is enabled
)

// ====== If resolving merge conflicts ======
//
// All calls to metrics.NewRegistered*() for metrics also defined in libevm/core/txpool/legacypool
// have been replaced with metrics.GetOrRegister*() to get metrics already registered in
// libevm/core/txpool/legacypool or register them here otherwise. These replacements ensure the
// same metrics are shared between the two packages.
var (
	// Metrics for the pending pool
	pendingDiscardMeter   = metrics.GetOrRegisterMeter("txpool/pending/discard", nil)
	pendingReplaceMeter   = metrics.GetOrRegisterMeter("txpool/pending/replace", nil)
	pendingRateLimitMeter = metrics.GetOrRegisterMeter("txpool/pending/ratelimit", nil) // Dropped due to rate limiting
	pendingNofundsMeter   = metrics.GetOrRegisterMeter("txpool/pending/nofunds", nil)   // Dropped due to out-of-funds

	// Metrics for the queued pool
	queuedDiscardMeter   = metrics.GetOrRegisterMeter("txpool/queued/discard", nil)
	queuedReplaceMeter   = metrics.GetOrRegisterMeter("txpool/queued/replace", nil)
	queuedRateLimitMeter = metrics.GetOrRegisterMeter("txpool/queued/ratelimit", nil) // Dropped due to rate limiting
	queuedNofundsMeter   = metrics.GetOrRegisterMeter("txpool/queued/nofunds", nil)   // Dropped due to out-of-funds
	queuedEvictionMeter  = metrics.GetOrRegisterMeter("txpool/queued/eviction", nil)  // Dropped due to lifetime

	// General tx metrics
	knownTxMeter       = metrics.GetOrRegisterMeter("txpool/known", nil)
	validTxMeter       = metrics.GetOrRegisterMeter("txpool/valid", nil)
	invalidTxMeter     = metrics.GetOrRegisterMeter("txpool/invalid", nil)
	underpricedTxMeter = metrics.GetOrRegisterMeter("txpool/underpriced", nil)
	overflowedTxMeter  = metrics.GetOrRegisterMeter("txpool/overflowed", nil)

	// throttleTxMeter counts how many transactions are rejected due to too-many-changes between
	// txpool reorgs.
	throttleTxMeter = metrics.GetOrRegisterMeter("txpool/throttle", nil)
	// reorgDurationTimer measures how long time a txpool reorg takes.
	reorgDurationTimer = metrics.GetOrRegisterTimer("txpool/reorgtime", nil)
	// dropBetweenReorgHistogram counts how many drops we experience between two reorg runs. It is expected
	// that this number is pretty low, since txpool reorgs happen very frequently.
	dropBetweenReorgHistogram = metrics.GetOrRegisterHistogram("txpool/dropbetweenreorg", nil, metrics.NewExpDecaySample(1028, 0.015))

	pendingGauge = metrics.GetOrRegisterGauge("txpool/pending", nil)
	queuedGauge  = metrics.GetOrRegisterGauge("txpool/queued", nil)
	localGauge   = metrics.GetOrRegisterGauge("txpool/local", nil)
	slotsGauge   = metrics.GetOrRegisterGauge("txpool/slots", nil)

	reheapTimer = metrics.GetOrRegisterTimer("txpool/reheap", nil)
)

// BlockChain defines the minimal set of methods needed to back a tx pool with
// a chain. Exists to allow mocking the live chain out of tests.
type BlockChain interface {
	// Config retrieves the chain's fork configuration.
	Config() *params.ChainConfig

	// CurrentBlock returns the current head of the chain.
	CurrentBlock() *types.Header

	// GetBlock retrieves a specific block, used during pool resets.
	GetBlock(hash common.Hash, number uint64) *types.Block

	// StateAt returns a state database for a given root hash (generally the head).
	StateAt(root common.Hash) (*state.StateDB, error)

	SenderCacher() *core.TxSenderCacher
	GetFeeConfigAt(parent *types.Header) (commontype.FeeConfig, *big.Int, error)
}

// Config are the configuration parameters of the transaction pool.
type Config struct {
	Locals    []common.Address // Addresses that should be treated by default as local
	NoLocals  bool             // Whether local transaction handling should be disabled
	Journal   string           // Journal of local transactions to survive node restarts
	Rejournal time.Duration    // Time interval to regenerate the local transaction journal

	PriceLimit uint64 // Minimum gas price to enforce for acceptance into the pool
	PriceBump  uint64 // Minimum price bump percentage to replace an already existing transaction (nonce)

	AccountSlots uint64 // Number of executable transaction slots guaranteed per account
	GlobalSlots  uint64 // Maximum number of executable transaction slots for all accounts
	AccountQueue uint64 // Maximum number of non-executable transaction slots permitted per account
	GlobalQueue  uint64 // Maximum number of non-executable transaction slots for all accounts

	Lifetime time.Duration // Maximum amount of time non-executable transaction are queued
}

// DefaultConfig contains the default configurations for the transaction pool.
var DefaultConfig = Config{
	// If we re-enable txpool journaling, we should also add the saved local
	// transactions to the p2p gossip on startup.
	Journal:   "",
	Rejournal: time.Hour,

	PriceLimit: 1,
	PriceBump:  10,

	AccountSlots: 16,
	GlobalSlots:  4096 + 1024, // urgent + floating queue capacity with 4:1 ratio
	AccountQueue: 64,
	GlobalQueue:  1024,

	Lifetime: 10 * time.Minute,
}

// sanitize checks the provided user configurations and changes anything that's
// unreasonable or unworkable.
func (config *Config) sanitize() Config { _ = "STUB: not implemented"; return *new(Config) }

// LegacyPool contains all currently known transactions. Transactions
// enter the pool when they are received from the network or submitted
// locally. They exit the pool when they are included in the blockchain.
//
// The pool separates processable transactions (which can be applied to the
// current state) and future transactions. Transactions move between those
// two states over time as they are received and processed.
type LegacyPool struct {
	config      Config
	chainconfig *params.ChainConfig
	chain       BlockChain
	gasTip      atomic.Pointer[uint256.Int]
	minimumFee  *big.Int
	txFeed      event.Feed
	signer      types.Signer
	mu          sync.RWMutex

	// closed when the transaction pool is stopped. Any goroutine can listen
	// to this to be notified if it should shut down.
	generalShutdownChan chan struct{}

	currentHead   atomic.Pointer[types.Header] // Current head of the blockchain
	currentState  *state.StateDB               // Current state in the blockchain head
	pendingNonces *noncer                      // Pending state tracking virtual nonces

	locals  *accountSet // Set of local transaction to exempt from eviction rules
	journal *journal    // Journal of local transaction to back up to disk

	reserve txpool.AddressReserver       // Address reserver to ensure exclusivity across subpools
	pending map[common.Address]*list     // All currently processable transactions
	queue   map[common.Address]*list     // Queued but non-processable transactions
	beats   map[common.Address]time.Time // Last heartbeat from each known account
	all     *lookup                      // All transactions to allow lookups
	priced  *pricedList                  // All transactions sorted by price

	reqResetCh      chan *txpoolResetRequest
	reqPromoteCh    chan *accountSet
	queueTxEventCh  chan *types.Transaction
	reorgDoneCh     chan chan struct{}
	reorgShutdownCh chan struct{}  // requests shutdown of scheduleReorgLoop
	wg              sync.WaitGroup // tracks loop, scheduleReorgLoop
	initDoneCh      chan struct{}  // is closed once the pool is initialized (for tests)

	changesSinceReorg int // A counter for how many drops we've performed in-between reorg.
}

type txpoolResetRequest struct {
	oldHead, newHead *types.Header
}

// New creates a new transaction pool to gather, sort and filter inbound
// transactions from the network.
func New(config Config, chain BlockChain) *LegacyPool {
	_ = "STUB: not implemented"
	// Sanitize the input to ensure no vulnerable gas prices are set
	return nil
}

// Create the transaction pool with its initial settings

// Filter returns whether the given transaction can be consumed by the legacy
// pool, specifically, whether it is a Legacy, AccessList or Dynamic transaction.
func (pool *LegacyPool) Filter(tx *types.Transaction) bool { _ = "STUB: not implemented"; return false }

// Init sets the gas price needed to keep a transaction in the pool and the chain
// head to allow balance / nonce checks. The transaction journal will be loaded
// from disk and filtered based on the provided starting settings. The internal
// goroutines will be spun up and the pool deemed operational afterwards.
func (pool *LegacyPool) Init(gasTip uint64, head *types.Header, reserve txpool.AddressReserver) error {
	_ = "STUB: not implemented"
	// Set the address reserver to request exclusive access to pooled accounts
	return nil
}

// Set the basic pool parameters

// Initialize the state with head block, or fallback to empty one in
// case the head state is not available (might occur when node is not
// fully synced).

// Start the reorg loop early, so it can handle requests generated during
// journal loading.

// If local transactions and journaling is enabled, load from disk

// loop is the transaction pool's main event loop, waiting for and reacting to
// outside blockchain events as well as for various reporting and transaction
// eviction events.
func (pool *LegacyPool) loop() { _ = "STUB: not implemented"; return }

// Start the stats reporting and transaction eviction tickers

// Notify tests that the init phase is done

// Handle pool shutdown

// Handle stats reporting ticks

// Handle inactive account transaction eviction

// Skip local transactions from the eviction mechanism

// Any non-locals old enough should be removed

// Handle local transaction journal rotation

// Close terminates the transaction pool.
func (pool *LegacyPool) Close() error { _ = "STUB: not implemented"; return nil }

// Terminate the pool reorger and return

// remove all references to state to allow GC to reclaim memory

// Reset implements txpool.SubPool, allowing the legacy pool's internal state to be
// kept in sync with the main transaction pool's internal state.
func (pool *LegacyPool) Reset(oldHead, newHead *types.Header) { _ = "STUB: not implemented"; return }

// SubscribeTransactions registers a subscription for new transaction events,
// supporting feeding only newly seen or also resurrected transactions.
func (pool *LegacyPool) SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription {
	_ = "STUB: not implemented"
	// The legacy pool has a very messed up internal shuffling, so it's kind of
	// hard to separate newly discovered transaction from resurrected ones. This
	// is because the new txs are added to the queue, resurrected ones too and
	// reorgs run lazily, so separating the two would need a marker.
	return *new(event.Subscription)
}

// SetGasTip updates the minimum gas tip required by the transaction pool for a
// new transaction, and drops all transactions below this threshold.
func (pool *LegacyPool) SetGasTip(tip *big.Int) { _ = "STUB: not implemented"; return }

// If the min miner fee increased, remove transactions below the new threshold

// pool.priced is sorted by GasFeeCap, so we have to iterate through pool.all instead

func (pool *LegacyPool) SetMinFee(minFee *big.Int) { _ = "STUB: not implemented"; return }

// Nonce returns the next nonce of an account, with all transactions executable
// by the pool already applied on top.
func (pool *LegacyPool) Nonce(addr common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

// Stats retrieves the current pool stats, namely the number of pending and the
// number of queued (non-executable) transactions.
func (pool *LegacyPool) Stats() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// stats retrieves the current pool stats, namely the number of pending and the
// number of queued (non-executable) transactions.
func (pool *LegacyPool) stats() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Content retrieves the data content of the transaction pool, returning all the
// pending as well as queued transactions, grouped by account and sorted by nonce.
func (pool *LegacyPool) Content() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContentFrom retrieves the data content of the transaction pool, returning the
// pending as well as queued transactions of this address, grouped by nonce.
func (pool *LegacyPool) ContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pending retrieves all currently processable transactions, grouped by origin
// account and sorted by nonce.
//
// The transactions can also be pre-filtered by the dynamic fee components to
// reduce allocations and load on downstream subsystems.
func (pool *LegacyPool) Pending(filter txpool.PendingFilter) map[common.Address][]*txpool.LazyTransaction {
	_ = "STUB: not implemented"
	// If only blob transactions are requested, this pool is unsuitable as it
	// contains none, don't even bother.
	return nil
}

// Convert the new uint256.Int types to the old big.Int ones used by the legacy pool

// If the miner requests tip enforcement, cap the lists now

// IteratePending iterates over [pool.pending] until [f] returns false.
// The caller must not modify [tx]. Returns false if iteration was interrupted.
func (pool *LegacyPool) IteratePending(f func(tx *types.Transaction) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Locals retrieves the accounts currently considered local by the pool.
func (pool *LegacyPool) Locals() []common.Address { _ = "STUB: not implemented"; return nil }

// local retrieves all currently known local transactions, grouped by origin
// account and sorted by nonce. The returned transaction set is a copy and can be
// freely modified by calling code.
func (pool *LegacyPool) local() map[common.Address]types.Transactions {
	_ = "STUB: not implemented"
	return nil
}

// validateTxBasics checks whether a transaction is valid according to the consensus
// rules, but does not check state-dependent validation such as sufficient balance.
// This check is meant as an early check which only needs to be performed once,
// and does not require the pool mutex to be held.
func (pool *LegacyPool) validateTxBasics(tx *types.Transaction, local bool) error {
	_ = "STUB: not implemented"
	return nil
}

// validateTx checks whether a transaction is valid according to the consensus
// rules and adheres to some heuristic limits of the local node (price and size).
func (pool *LegacyPool) validateTx(tx *types.Transaction, local bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Pool allows arbitrary arrival order, don't invalidate nonce gaps

// add validates a transaction and inserts it into the non-executable queue for later
// pending promotion and execution. If the transaction is a replacement for an already
// pending or queued one, it overwrites the previous transaction if its price is higher.
//
// If a newly added transaction is marked as local, its sending account will be
// added to the allowlist, preventing any associated transaction from being dropped
// out of the pool due to pricing constraints.
func (pool *LegacyPool) add(tx *types.Transaction, local bool) (replaced bool, err error) {
	_ = "STUB: not implemented"
	// If the transaction is already known, discard it
	return false, nil
}

// Make the local flag. If it's from local source or it's from the network but
// the sender is marked as local previously, treat it as the local transaction.

// If the transaction fails basic validation, discard it

// already validated by this point

// If the address is not yet known, request exclusivity to track the account
// only by this subpool until all transactions are evicted

// If the transaction is rejected by some post-validation check, remove
// the lock on the reservation set.
//
// Note, `err` here is the named error return, which will be initialized
// by a return statement before running deferred methods. Take care with
// removing or subscoping err as it will break this clause.

// If the transaction pool is full, discard underpriced transactions

// If the new transaction is underpriced, don't accept it

// We're about to replace a transaction. The reorg does a more thorough
// analysis of what to remove and how, but it runs async. We don't want to
// do too many replacements between reorg-runs, so we cap the number of
// replacements to 25% of the slots

// New transaction is better than our worse ones, make room for it.
// If it's a local transaction, forcibly discard all available transactions.
// Otherwise if we can't make enough room for new one, abort the operation.

// Special case, we still can't make the room for the new remote one.

// If the new transaction is a future transaction it should never churn pending transactions

// Add all transactions back to the priced queue

// Kick out the underpriced remote transactions.

// Don't unreserve the sender of the tx being added if last from the acc

// Try to replace an existing transaction in the pending pool

// Nonce already pending, check if required price bump is met

// New transaction is better, replace old one

// Successful promotion, bump the heartbeat

// New transaction isn't replacing a pending one, push into queue

// Mark local addresses and journal local transactions

// Migrate the remotes if it's marked as local first time.

// isGapped reports whether the given transaction is immediately executable.
func (pool *LegacyPool) isGapped(from common.Address, tx *types.Transaction) bool {
	_ = "STUB: not implemented"
	// Short circuit if transaction falls within the scope of the pending list
	// or matches the next pending nonce which can be promoted as an executable
	// transaction afterwards. Note, the tx staleness is already checked in
	// 'validateTx' function previously.
	return false
}

// The transaction has a nonce gap with pending list, it's only considered
// as executable if transactions in queue can fill up the nonce gap.

// txs in queue can't fill up the nonce gap

// enqueueTx inserts a new transaction into the non-executable transaction queue.
//
// Note, this method assumes the pool lock is held!
func (pool *LegacyPool) enqueueTx(hash common.Hash, tx *types.Transaction, local bool, addAll bool) (bool, error) {
	_ = "STUB: not implemented"
	// Try to insert the transaction into the future queue
	return false, nil
}

// already validated

// An older transaction was better, discard this

// Discard any previous transaction and mark this

// Nothing was replaced, bump the queued counter

// If the transaction isn't in lookup set but it's expected to be there,
// show the error log.

// If we never record the heartbeat, do it right now.

// journalTx adds the specified transaction to the local disk journal if it is
// deemed to have been sent from a local account.
func (pool *LegacyPool) journalTx(from common.Address, tx *types.Transaction) {
	_ = "STUB: not implemented"
	// Only journal if it's enabled and the transaction is local
	return
}

// promoteTx adds a transaction to the pending (processable) list of transactions
// and returns whether it was inserted or an older was better.
//
// Note, this method assumes the pool lock is held!
func (pool *LegacyPool) promoteTx(addr common.Address, hash common.Hash, tx *types.Transaction) bool {
	_ = "STUB: not implemented"
	// Try to insert the transaction into the pending queue
	return false
}

// An older transaction was better, discard this

// Otherwise discard any previous transaction and mark this

// Nothing was replaced, bump the pending counter

// Set the potentially new pending nonce and notify any subsystems of the new tx

// Successful promotion, bump the heartbeat

// addLocals enqueues a batch of transactions into the pool if they are valid, marking the
// senders as local ones, ensuring they go around the local pricing constraints.
//
// This method is used to add transactions from the RPC API and performs synchronous pool
// reorganization and event propagation.
func (pool *LegacyPool) addLocals(txs []*types.Transaction) []error {
	_ = "STUB: not implemented"
	return nil
}

// addLocal enqueues a single local transaction into the pool if it is valid. This is
// a convenience wrapper around addLocals.
func (pool *LegacyPool) addLocal(tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// addRemotes enqueues a batch of transactions into the pool if they are valid. If the
// senders are not among the locally tracked ones, full pricing constraints will apply.
//
// This method is used to add transactions from the p2p network and does not wait for pool
// reorganization and internal event propagation.
func (pool *LegacyPool) addRemotes(txs []*types.Transaction) []error {
	_ = "STUB: not implemented"
	return nil
}

// addRemote enqueues a single transaction into the pool if it is valid. This is a convenience
// wrapper around addRemotes.
func (pool *LegacyPool) addRemote(tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// addRemotesSync is like addRemotes, but waits for pool reorganization. Tests use this method.
func (pool *LegacyPool) addRemotesSync(txs []*types.Transaction) []error {
	_ = "STUB: not implemented"
	return nil
}

// This is like addRemotes with a single transaction, but waits for pool reorganization. Tests use this method.
func (pool *LegacyPool) addRemoteSync(tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// Add enqueues a batch of transactions into the pool if they are valid. Depending
// on the local flag, full pricing constraints will or will not be applied.
//
// If sync is set, the method will block until all internal maintenance related
// to the add is finished. Only use this during tests for determinism!
func (pool *LegacyPool) Add(txs []*types.Transaction, local, sync bool) []error {
	_ = "STUB: not implemented"
	// Do not treat as local if local transactions have been disabled
	return nil
}

// Filter out known ones without obtaining the pool lock or recovering signatures

// If the transaction is known, pre-set the error slot

// Exclude transactions with basic errors, e.g invalid signatures and
// insufficient intrinsic gas as soon as possible and cache senders
// in transactions before obtaining lock

// Accumulate all unknown transactions for deeper processing

// Process all the new transaction and merge any errors into the original slice

// Reorg the pool internals if needed and return

// addTxsLocked attempts to queue a batch of transactions if they are valid.
// The transaction pool lock must be held.
func (pool *LegacyPool) addTxsLocked(txs []*types.Transaction, local bool) ([]error, *accountSet) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Status returns the status (unknown/pending/queued) of a batch of transactions
// identified by their hashes.
func (pool *LegacyPool) Status(hash common.Hash) txpool.TxStatus {
	_ = "STUB: not implemented"
	return *new(txpool.TxStatus)
}

// already validated

// Get returns a transaction if it is contained in the pool and nil otherwise.
func (pool *LegacyPool) Get(hash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// get returns a transaction if it is contained in the pool and nil otherwise.
func (pool *LegacyPool) get(hash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	return nil

	// Has returns an indicator whether txpool has a transaction cached with the
	// given hash.
}

func (pool *LegacyPool) Has(hash common.Hash) bool { _ = "STUB: not implemented"; return false }

func (pool *LegacyPool) HasLocal(hash common.Hash) bool { _ = "STUB: not implemented"; return false }

// removeTx removes a single transaction from the queue, moving all subsequent
// transactions back to the future queue.
//
// In unreserve is false, the account will not be relinquished to the main txpool
// even if there are no more references to it. This is used to handle a race when
// a tx being added, and it evicts a previously scheduled tx from the same account,
// which could lead to a premature release of the lock.
//
// Returns the number of transactions removed from the pending queue.
func (pool *LegacyPool) removeTx(hash common.Hash, outofbound bool, unreserve bool) int {
	_ = "STUB: not implemented"
	// Fetch the transaction we wish to delete
	return 0
}

// already validated during insertion

// If after deletion there are no more transactions belonging to this account,
// relinquish the address reservation. It's a bit convoluted do this, via a
// defer, but it's safer vs. the many return pathways.

// Remove it from the list of known transactions

// Remove the transaction from the pending lists and reset the account nonce

// If no more pending transactions are left, remove the list

// Postpone any invalidated transactions

// Internal shuffle shouldn't touch the lookup set.

// Update the account nonce if needed

// Reduce the pending counter

// Transaction is in the future queue

// Reduce the queued counter

// requestReset requests a pool reset to the new head block.
// The returned channel is closed when the reset has occurred.
func (pool *LegacyPool) requestReset(oldHead *types.Header, newHead *types.Header) chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// requestPromoteExecutables requests transaction promotion checks for the given addresses.
// The returned channel is closed when the promotion checks have occurred.
func (pool *LegacyPool) requestPromoteExecutables(set *accountSet) chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// queueTxEvent enqueues a transaction event to be sent in the next reorg run.
func (pool *LegacyPool) queueTxEvent(tx *types.Transaction) { _ = "STUB: not implemented"; return }

// scheduleReorgLoop schedules runs of reset and promoteExecutables. Code above should not
// call those methods directly, but request them being run using requestReset and
// requestPromoteExecutables instead.
func (pool *LegacyPool) scheduleReorgLoop() { _ = "STUB: not implemented"; return }

// non-nil while runReorg is active

// Launch next background reorg if needed

// Run the background reorg and announcements

// Prepare everything for the next round of reorg

// Reset request: update head if request is already pending.

// Promote request: update address set if request is already pending.

// Queue up the event, but don't schedule a reorg. It's up to the caller to
// request one later if they want the events sent.

// Wait for current run to finish.

// runReorg runs reset and promoteExecutables on behalf of scheduleReorgLoop.
func (pool *LegacyPool) runReorg(done chan struct{}, reset *txpoolResetRequest, dirtyAccounts *accountSet, events map[common.Address]*sortedMap) {
	_ = "STUB: not implemented"
	return
}

// Only dirty accounts need to be promoted, unless we're resetting.
// For resets, all addresses in the tx queue will be promoted and
// the flatten operation can be avoided.

// Reset from the old head to the new, rescheduling any reorged transactions

// Nonces were reset, discard any events that became stale

// Reset needs promote for all addresses

// Check for pending transactions for every account that sent new ones

// If a new block appeared, validate the pool of pending transactions. This will
// remove any transaction that has been included in the block or was invalidated
// because of another transaction (e.g. higher gas price).

// Update all accounts to the latest known pending nonce

// Ensure pool.queue and pool.pending sizes stay within the configured limits.

// Reset change counter

// Notify subsystems for newly added transactions

// reset retrieves the current state of the blockchain and ensures the content
// of the transaction pool is valid with regard to the chain state.
func (pool *LegacyPool) reset(oldHead, newHead *types.Header) {
	_ = "STUB: not implemented"
	// If we're reorging an old state, reinject all dropped transactions
	return
}

// If the reorg is too deep, avoid doing it (will happen during fast sync)

// Reorg seems shallow enough to pull in all transactions into memory

// This can happen if a setHead is performed, where we simply discard the old
// head from the chain.
// If that is the case, we don't have the lost transactions anymore, and
// there's nothing to add

// If we reorged to a same or higher number, then it's not a case of setHead

// If the reorg ended up on a lower number, it's indicative of setHead being the cause

// We still need to update the current state s.th. the lost transactions can be readded by the user

// if the new head is nil, it means that something happened between
// the firing of newhead-event and _now_: most likely a
// reorg caused by sync-reversion or explicit sethead back to an
// earlier block.

// Initialize the internal state to the current head

// Special case during testing

// when we reset txPool we should explicitly check if fee struct for min base fee has changed
// so that we can correctly drop txs with < minBaseFee from tx pool.

// Inject any transactions discarded due to reorgs

// promoteExecutables moves transactions that have become processable from the
// future queue to the set of pending transactions. During this process, all
// invalidated transactions (low nonce, low balance) are deleted.
func (pool *LegacyPool) promoteExecutables(accounts []common.Address) []*types.Transaction {
	_ = "STUB: not implemented"
	// Track the promoted transactions to broadcast them at once
	return nil
}

// Iterate over all accounts and promote any executable transactions

// Just in case someone calls with a non existing account

// Drop all transactions that are deemed too old (low nonce)

// Drop all transactions that are too costly (low balance or out of gas)

// Gather all executable transactions and promote them

// Drop all transactions over the allowed limit

// Mark all the items dropped as removed

// Delete the entire queue entry if it became empty.

// truncatePending removes transactions from the pending queue if the pool is above the
// pending limit. The algorithm tries to reduce transaction counts by an approximately
// equal number for all for accounts with many pending transactions.
func (pool *LegacyPool) truncatePending() { _ = "STUB: not implemented"; return }

// Assemble a spam order to penalize large transactors first

// Only evict transactions from high rollers

// Gradually drop transactions from offenders

// Retrieve the next offender if not local address

// Equalize balances until all the same or below threshold

// Calculate the equalization threshold for all current offenders

// Iteratively reduce all offenders until below limit or threshold reached

// Drop the transaction from the global pools too

// Update the account nonce to the dropped transaction

// If still above threshold, reduce to limit or min allowance

// Drop the transaction from the global pools too

// Update the account nonce to the dropped transaction

// truncateQueue drops the oldest transactions in the queue if the pool is above the global queue limit.
func (pool *LegacyPool) truncateQueue() { _ = "STUB: not implemented"; return }

// Sort all accounts with queued transactions by heartbeat

// don't drop locals

// Drop transactions until the total is below the limit or only locals remain

// Drop all transactions if they are less than the overflow

// Otherwise drop only last few transactions

// demoteUnexecutables removes invalid and processed transactions from the pools
// executable/pending queue and any subsequent transactions that become unexecutable
// are moved back into the future queue.
//
// Note: transactions are not marked as removed in the priced list because re-heaping
// is always explicitly triggered by SetBaseFee and it would be unnecessary and wasteful
// to trigger a re-heap is this function
func (pool *LegacyPool) demoteUnexecutables() {
	_ = "STUB: not implemented"
	// Iterate over all accounts and demote any non-executable transactions
	return
}

// Drop all transactions that are deemed too old (low nonce)

// Drop all transactions that are too costly (low balance or out of gas), and queue any invalids back for later

// Internal shuffle shouldn't touch the lookup set.

// If there's a gap in front, alert (should never happen) and postpone all transactions

// Internal shuffle shouldn't touch the lookup set.

// Delete the entire pending entry if it became empty.

func (pool *LegacyPool) startPeriodicFeeUpdate() { _ = "STUB: not implemented"; return }

// Call updateBaseFee here to ensure that there is not a [baseFeeUpdateInterval] delay
// when starting up in ApricotPhase3 before the base fee is updated.

func (pool *LegacyPool) periodicBaseFeeUpdate() { _ = "STUB: not implemented"; return }

// Sleep until its time to start the periodic base fee update or the tx pool is shutting down

// Return early if shutting down

// Update the base fee every [baseFeeUpdateInterval]
// and shutdown when [generalShutdownChan] is closed by Stop()

// updateBaseFee updates the base fee in the tx pool based on the current head block.
// should only be called when the chain is in Subnet EVM.
func (pool *LegacyPool) updateBaseFee() { _ = "STUB: not implemented"; return }

// assumes lock is already held
// should only be called when the chain is in Subnet EVM.
func (pool *LegacyPool) updateBaseFeeAt(head *types.Header) error {
	_ = "STUB: not implemented"
	return nil
}

// addressByHeartbeat is an account address tagged with its last activity timestamp.
type addressByHeartbeat struct {
	address   common.Address
	heartbeat time.Time
}

type addressesByHeartbeat []addressByHeartbeat

func (a addressesByHeartbeat) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a addressesByHeartbeat) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (a addressesByHeartbeat) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// accountSet is simply a set of addresses to check for existence, and a signer
// capable of deriving addresses from transactions.
type accountSet struct {
	accounts map[common.Address]struct{}
	signer   types.Signer
	cache    *[]common.Address
}

// newAccountSet creates a new address set with an associated signer for sender
// derivations.
func newAccountSet(signer types.Signer, addrs ...common.Address) *accountSet {
	_ = "STUB: not implemented"
	return nil
}

// contains checks if a given address is contained within the set.
func (as *accountSet) contains(addr common.Address) bool { _ = "STUB: not implemented"; return false }

// containsTx checks if the sender of a given tx is within the set. If the sender
// cannot be derived, this method returns false.
func (as *accountSet) containsTx(tx *types.Transaction) bool {
	_ = "STUB: not implemented"
	return false
}

// add inserts a new address into the set to track.
func (as *accountSet) add(addr common.Address) { _ = "STUB: not implemented"; return }

// addTx adds the sender of tx into the set.
func (as *accountSet) addTx(tx *types.Transaction) { _ = "STUB: not implemented"; return }

// flatten returns the list of addresses within this set, also caching it for later
// reuse. The returned slice should not be changed!
func (as *accountSet) flatten() []common.Address { _ = "STUB: not implemented"; return nil }

// merge adds all addresses from the 'other' set into 'as'.
func (as *accountSet) merge(other *accountSet) { _ = "STUB: not implemented"; return }

// lookup is used internally by LegacyPool to track transactions while allowing
// lookup without mutex contention.
//
// Note, although this type is properly protected against concurrent access, it
// is **not** a type that should ever be mutated or even exposed outside of the
// transaction pool, since its internal state is tightly coupled with the pools
// internal mechanisms. The sole purpose of the type is to permit out-of-bound
// peeking into the pool in LegacyPool.Get without having to acquire the widely scoped
// LegacyPool.mu mutex.
//
// This lookup set combines the notion of "local transactions", which is useful
// to build upper-level structure.
type lookup struct {
	slots   int
	lock    sync.RWMutex
	locals  map[common.Hash]*types.Transaction
	remotes map[common.Hash]*types.Transaction
}

// newLookup returns a new lookup structure.
func newLookup() *lookup { _ = "STUB: not implemented"; return nil }

// Range calls f on each key and value present in the map. The callback passed
// should return the indicator whether the iteration needs to be continued.
// Callers need to specify which set (or both) to be iterated.
func (t *lookup) Range(f func(hash common.Hash, tx *types.Transaction, local bool) bool, local bool, remote bool) {
	_ = "STUB: not implemented"
	return
}

// Get returns a transaction if it exists in the lookup, or nil if not found.
func (t *lookup) Get(hash common.Hash) *types.Transaction { _ = "STUB: not implemented"; return nil }

// GetLocal returns a transaction if it exists in the lookup, or nil if not found.
func (t *lookup) GetLocal(hash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// GetRemote returns a transaction if it exists in the lookup, or nil if not found.
func (t *lookup) GetRemote(hash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// Count returns the current number of transactions in the lookup.
func (t *lookup) Count() int { _ = "STUB: not implemented"; return 0 }

// LocalCount returns the current number of local transactions in the lookup.
func (t *lookup) LocalCount() int { _ = "STUB: not implemented"; return 0 }

// RemoteCount returns the current number of remote transactions in the lookup.
func (t *lookup) RemoteCount() int { _ = "STUB: not implemented"; return 0 }

// Slots returns the current number of slots used in the lookup.
func (t *lookup) Slots() int { _ = "STUB: not implemented"; return 0 }

// Add adds a transaction to the lookup.
func (t *lookup) Add(tx *types.Transaction, local bool) { _ = "STUB: not implemented"; return }

// Remove removes a transaction from the lookup.
func (t *lookup) Remove(hash common.Hash) { _ = "STUB: not implemented"; return }

// RemoteToLocals migrates the transactions belongs to the given locals to locals
// set. The assumption is held the locals set is thread-safe to be used.
func (t *lookup) RemoteToLocals(locals *accountSet) int { _ = "STUB: not implemented"; return 0 }

// RemotesBelowTip finds all remote transactions below the given tip threshold.
func (t *lookup) RemotesBelowTip(threshold *big.Int) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// Only iterate remotes

// numSlots calculates the number of slots needed for a single transaction.
func numSlots(tx *types.Transaction) int { _ = "STUB: not implemented"; return 0 }
