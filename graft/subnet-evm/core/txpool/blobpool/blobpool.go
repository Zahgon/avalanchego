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
// Copyright 2022 The go-ethereum Authors
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

// Package blobpool implements the EIP-4844 blob transaction pool.
package blobpool

import (
	"math/big"
	"sync"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core/txpool"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	ethparams "github.com/ava-labs/libevm/params"
	"github.com/holiman/billy"
	"github.com/holiman/uint256"
)

const (
	// blobSize is the protocol constrained byte size of a single blob in a
	// transaction. There can be multiple of these embedded into a single tx.
	blobSize = ethparams.BlobTxFieldElementsPerBlob * ethparams.BlobTxBytesPerFieldElement

	// maxBlobsPerTransaction is the maximum number of blobs a single transaction
	// is allowed to contain. Whilst the spec states it's unlimited, the block
	// data slots are protocol bound, which implicitly also limit this.
	maxBlobsPerTransaction = ethparams.MaxBlobGasPerBlock / ethparams.BlobTxBlobGasPerBlob

	// txAvgSize is an approximate byte size of a transaction metadata to avoid
	// tiny overflows causing all txs to move a shelf higher, wasting disk space.
	txAvgSize = 4 * 1024

	// txMaxSize is the maximum size a single transaction can have, outside
	// the included blobs. Since blob transactions are pulled instead of pushed,
	// and only a small metadata is kept in ram, the rest is on disk, there is
	// no critical limit that should be enforced. Still, capping it to some sane
	// limit can never hurt.
	txMaxSize = 1024 * 1024

	// maxTxsPerAccount is the maximum number of blob transactions admitted from
	// a single account. The limit is enforced to minimize the DoS potential of
	// a private tx cancelling publicly propagated blobs.
	//
	// Note, transactions resurrected by a reorg are also subject to this limit,
	// so pushing it down too aggressively might make resurrections non-functional.
	maxTxsPerAccount = 16

	// pendingTransactionStore is the subfolder containing the currently queued
	// blob transactions.
	pendingTransactionStore = "queue"

	// limboedTransactionStore is the subfolder containing the currently included
	// but not yet finalized transaction blobs.
	limboedTransactionStore = "limbo"
)

// blobTxMeta is the minimal subset of types.BlobTx necessary to validate and
// schedule the blob transactions into the following blocks. Only ever add the
// bare minimum needed fields to keep the size down (and thus number of entries
// larger with the same memory consumption).
type blobTxMeta struct {
	hash common.Hash // Transaction hash to maintain the lookup table
	id   uint64      // Storage ID in the pool's persistent store
	size uint32      // Byte size in the pool's persistent store

	nonce      uint64       // Needed to prioritize inclusion order within an account
	costCap    *uint256.Int // Needed to validate cumulative balance sufficiency
	execTipCap *uint256.Int // Needed to prioritize inclusion order across accounts and validate replacement price bump
	execFeeCap *uint256.Int // Needed to validate replacement price bump
	blobFeeCap *uint256.Int // Needed to validate replacement price bump
	execGas    uint64       // Needed to check inclusion validity before reading the blob
	blobGas    uint64       // Needed to check inclusion validity before reading the blob

	basefeeJumps float64 // Absolute number of 1559 fee adjustments needed to reach the tx's fee cap
	blobfeeJumps float64 // Absolute number of 4844 fee adjustments needed to reach the tx's blob fee cap

	evictionExecTip      *uint256.Int // Worst gas tip across all previous nonces
	evictionExecFeeJumps float64      // Worst base fee (converted to fee jumps) across all previous nonces
	evictionBlobFeeJumps float64      // Worse blob fee (converted to fee jumps) across all previous nonces
}

// newBlobTxMeta retrieves the indexed metadata fields from a blob transaction
// and assembles a helper struct to track in memory.
func newBlobTxMeta(id uint64, size uint32, tx *types.Transaction) *blobTxMeta {
	_ = "STUB: not implemented"
	return nil
}

// BlobPool is the transaction pool dedicated to EIP-4844 blob transactions.
//
// Blob transactions are special snowflakes that are designed for a very specific
// purpose (rollups) and are expected to adhere to that specific use case. These
// behavioural expectations allow us to design a transaction pool that is more robust
// (i.e. resending issues) and more resilient to DoS attacks (e.g. replace-flush
// attacks) than the generic tx pool. These improvements will also mean, however,
// that we enforce a significantly more aggressive strategy on entering and exiting
// the pool:
//
//   - Blob transactions are large. With the initial design aiming for 128KB blobs,
//     we must ensure that these only traverse the network the absolute minimum
//     number of times. Broadcasting to sqrt(peers) is out of the question, rather
//     these should only ever be announced and the remote side should request it if
//     it wants to.
//
//   - Block blob-space is limited. With blocks being capped to a few blob txs, we
//     can make use of the very low expected churn rate within the pool. Notably,
//     we should be able to use a persistent disk backend for the pool, solving
//     the tx resend issue that plagues the generic tx pool, as long as there's no
//     artificial churn (i.e. pool wars).
//
//   - Purpose of blobs are layer-2s. Layer-2s are meant to use blob transactions to
//     commit to their own current state, which is independent of Ethereum mainnet
//     (state, txs). This means that there's no reason for blob tx cancellation or
//     replacement, apart from a potential basefee / miner tip adjustment.
//
//   - Replacements are expensive. Given their size, propagating a replacement
//     blob transaction to an existing one should be aggressively discouraged.
//     Whilst generic transactions can start at 1 Wei gas cost and require a 10%
//     fee bump to replace, we suggest requiring a higher min cost (e.g. 1 gwei)
//     and a more aggressive bump (100%).
//
//   - Cancellation is prohibitive. Evicting an already propagated blob tx is a huge
//     DoS vector. As such, a) replacement (higher-fee) blob txs mustn't invalidate
//     already propagated (future) blob txs (cumulative fee); b) nonce-gapped blob
//     txs are disallowed; c) the presence of blob transactions exclude non-blob
//     transactions.
//
//   - Malicious cancellations are possible. Although the pool might prevent txs
//     that cancel blobs, blocks might contain such transaction (malicious miner
//     or flashbotter). The pool should cap the total number of blob transactions
//     per account as to prevent propagating too much data before cancelling it
//     via a normal transaction. It should nonetheless be high enough to support
//     resurrecting reorged transactions. Perhaps 4-16.
//
//   - Local txs are meaningless. Mining pools historically used local transactions
//     for payouts or for backdoor deals. With 1559 in place, the basefee usually
//     dominates the final price, so 0 or non-0 tip doesn't change much. Blob txs
//     retain the 1559 2D gas pricing (and introduce on top a dynamic blob gas fee),
//     so locality is moot. With a disk backed blob pool avoiding the resend issue,
//     there's also no need to save own transactions for later.
//
//   - No-blob blob-txs are bad. Theoretically there's no strong reason to disallow
//     blob txs containing 0 blobs. In practice, admitting such txs into the pool
//     breaks the low-churn invariant as blob constraints don't apply anymore. Even
//     though we could accept blocks containing such txs, a reorg would require moving
//     them back into the blob pool, which can break invariants.
//
//   - Dropping blobs needs delay. When normal transactions are included, they
//     are immediately evicted from the pool since they are contained in the
//     including block. Blobs however are not included in the execution chain,
//     so a mini reorg cannot re-pool "lost" blob transactions. To support reorgs,
//     blobs are retained on disk until they are finalised.
//
//   - Blobs can arrive via flashbots. Blocks might contain blob transactions we
//     have never seen on the network. Since we cannot recover them from blocks
//     either, the engine_newPayload needs to give them to us, and we cache them
//     until finality to support reorgs without tx losses.
//
// Whilst some constraints above might sound overly aggressive, the general idea is
// that the blob pool should work robustly for its intended use case and whilst
// anyone is free to use blob transactions for arbitrary non-rollup use cases,
// they should not be allowed to run amok the network.
//
// Implementation wise there are a few interesting design choices:
//
//   - Adding a transaction to the pool blocks until persisted to disk. This is
//     viable because TPS is low (2-4 blobs per block initially, maybe 8-16 at
//     peak), so natural churn is a couple MB per block. Replacements doing O(n)
//     updates are forbidden and transaction propagation is pull based (i.e. no
//     pileup of pending data).
//
//   - When transactions are chosen for inclusion, the primary criteria is the
//     signer tip (and having a basefee/data fee high enough of course). However,
//     same-tip transactions will be split by their basefee/datafee, preferring
//     those that are closer to the current network limits. The idea being that
//     very relaxed ones can be included even if the fees go up, when the closer
//     ones could already be invalid.
//
// When the pool eventually reaches saturation, some old transactions - that may
// never execute - will need to be evicted in favor of newer ones. The eviction
// strategy is quite complex:
//
//   - Exceeding capacity evicts the highest-nonce of the account with the lowest
//     paying blob transaction anywhere in the pooled nonce-sequence, as that tx
//     would be executed the furthest in the future and is thus blocking anything
//     after it. The smallest is deliberately not evicted to avoid a nonce-gap.
//
//   - Analogously, if the pool is full, the consideration price of a new tx for
//     evicting an old one is the smallest price in the entire nonce-sequence of
//     the account. This avoids malicious users DoSing the pool with seemingly
//     high paying transactions hidden behind a low-paying blocked one.
//
//   - Since blob transactions have 3 price parameters: execution tip, execution
//     fee cap and data fee cap, there's no singular parameter to create a total
//     price ordering on. What's more, since the base fee and blob fee can move
//     independently of one another, there's no pre-defined way to combine them
//     into a stable order either. This leads to a multi-dimensional problem to
//     solve after every block.
//
//   - The first observation is that comparing 1559 base fees or 4844 blob fees
//     needs to happen in the context of their dynamism. Since these fees jump
//     up or down in ~1.125 multipliers (at max) across blocks, comparing fees
//     in two transactions should be based on log1.125(fee) to eliminate noise.
//
//   - The second observation is that the basefee and blobfee move independently,
//     so there's no way to split mixed txs on their own (A has higher base fee,
//     B has higher blob fee). Rather than look at the absolute fees, the useful
//     metric is the max time it can take to exceed the transaction's fee caps.
//     Specifically, we're interested in the number of jumps needed to go from
//     the current fee to the transaction's cap:
//
//     jumps = log1.125(txfee) - log1.125(basefee)
//
//   - The third observation is that the base fee tends to hover around rather
//     than swing wildly. The number of jumps needed from the current fee starts
//     to get less relevant the higher it is. To remove the noise here too, the
//     pool will use log(jumps) as the delta for comparing transactions.
//
//     delta = sign(jumps) * log(abs(jumps))
//
//   - To establish a total order, we need to reduce the dimensionality of the
//     two base fees (log jumps) to a single value. The interesting aspect from
//     the pool's perspective is how fast will a tx get executable (fees going
//     down, crossing the smaller negative jump counter) or non-executable (fees
//     going up, crossing the smaller positive jump counter). As such, the pool
//     cares only about the min of the two delta values for eviction priority.
//
//     priority = min(deltaBasefee, deltaBlobfee)
//
//   - The above very aggressive dimensionality and noise reduction should result
//     in transaction being grouped into a small number of buckets, the further
//     the fees the larger the buckets. This is good because it allows us to use
//     the miner tip meaningfully as a splitter.
//
//   - For the scenario where the pool does not contain non-executable blob txs
//     anymore, it does not make sense to grant a later eviction priority to txs
//     with high fee caps since it could enable pool wars. As such, any positive
//     priority will be grouped together.
//
//     priority = min(deltaBasefee, deltaBlobfee, 0)
//
// Optimisation tradeoffs:
//
//   - Eviction relies on 3 fee minimums per account (exec tip, exec cap and blob
//     cap). Maintaining these values across all transactions from the account is
//     problematic as each transaction replacement or inclusion would require a
//     rescan of all other transactions to recalculate the minimum. Instead, the
//     pool maintains a rolling minimum across the nonce range. Updating all the
//     minimums will need to be done only starting at the swapped in/out nonce
//     and leading up to the first no-change.
type BlobPool struct {
	config  Config                 // Pool configuration
	reserve txpool.AddressReserver // Address reserver to ensure exclusivity across subpools

	store  billy.Database // Persistent data store for the tx metadata and blobs
	stored uint64         // Useful data size of all transactions on disk
	limbo  *limbo         // Persistent data store for the non-finalized blobs

	signer types.Signer // Transaction signer to use for sender recovery
	chain  BlockChain   // Chain object to access the state through

	head   *types.Header  // Current head of the chain
	state  *state.StateDB // Current state at the head of the chain
	gasTip *uint256.Int   // Currently accepted minimum gas tip

	lookup map[common.Hash]uint64           // Lookup table mapping hashes to tx billy entries
	index  map[common.Address][]*blobTxMeta // Blob transactions grouped by accounts, sorted by nonce
	spent  map[common.Address]*uint256.Int  // Expenditure tracking for individual accounts
	evict  *evictHeap                       // Heap of cheapest accounts for eviction when full

	discoverFeed event.Feed // Event feed to send out new tx events on pool discovery (reorg excluded)
	insertFeed   event.Feed // Event feed to send out new tx events on pool inclusion (reorg included)

	lock sync.RWMutex // Mutex protecting the pool during reorg handling
}

// New creates a new blob transaction pool to gather, sort and filter inbound
// blob transactions from the network.
func New(config Config, chain BlockChain) *BlobPool {
	_ = "STUB: not implemented"
	// Sanitize the input to ensure no vulnerable gas prices are set
	return nil
}

// Create the transaction pool with its initial settings

// Filter returns whether the given transaction can be consumed by the blob pool.
func (p *BlobPool) Filter(tx *types.Transaction) bool { _ = "STUB: not implemented"; return false }

// Init sets the gas price needed to keep a transaction in the pool and the chain
// head to allow balance / nonce checks. The transaction journal will be loaded
// from disk and filtered based on the provided starting settings.
func (p *BlobPool) Init(gasTip uint64, head *types.Header, reserve txpool.AddressReserver) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize the state with head block, or fallback to empty one in
// case the head state is not available (might occur when node is not
// fully synced).

// Index all transactions on disk and delete anything unprocessable

// Sort the indexed transactions by nonce and delete anything gapped, create
// the eviction heap of anyone still standing

// basefee = uint256.MustFromBig(eip1559.CalcBaseFee(p.chain.Config(), p.head))

// Pool initialized, attach the blob limbo to it to track blobs included
// recently but not yet finalized

// Set the configured gas tip, triggering a filtering of anything just loaded

// Since the user might have modified their pool's capacity, evict anything
// above the current allowance

// Update the metrics and return the constructed pool

// Close closes down the underlying persistent store.
func (p *BlobPool) Close() error { _ = "STUB: not implemented"; return nil }

// Close might be invoked due to error in constructor, before p,limbo is set

// parseTransaction is a callback method on pool creation that gets called for
// each transaction on disk to create the in-memory metadata index.
func (p *BlobPool) parseTransaction(id uint64, size uint32, blob []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// This path is impossible unless the disk data representation changes
// across restarts. For that ever improbable case, recover gracefully
// by ignoring this data entry.

// This path is only possible after a crash, where deleted items are not
// removed via the normal shutdown-startup procedure and thus may get
// partially resurrected.

// This path is impossible unless the signature validity changes across
// restarts. For that ever improbable case, recover gracefully by ignoring
// this data entry.

// recheck verifies the pool's content for a specific account and drops anything
// that does not fit anymore (dangling or filled nonce, overdraft).
func (p *BlobPool) recheck(addr common.Address, inclusions map[common.Hash]uint64) {
	_ = "STUB: not implemented"
	// Sort the transactions belonging to the account so reinjects can be simpler
	return
}

// during reorgs, we might find new accounts

// If there is a gap between the chain state and the blob pool, drop
// all the transactions as they are non-executable. Similarly, if the
// entire tx range was included, drop all.

// Included transactions blobs need to be moved to the limbo

// only during reorgs will the heap be initialized

// If there is overlap between the chain state and the blob pool, drop
// anything below the current state

// Included transactions blobs need to be moved to the limbo

// Iterate over the transactions to initialize their eviction thresholds
// and to detect any nonce gaps

// If there's no nonce gap, initialize the eviction thresholds as the
// minimum between the cumulative thresholds and the current tx fees

// Sanity check that there's no double nonce. This case would generally
// be a coding error, so better know about it.
//
// Also, Billy behind the blobpool does not journal deletes. A process
// crash would result in previously deleted entities being resurrected.
// That could potentially cause a duplicate nonce to appear.

// Otherwise if there's a nonce gap evict all later transactions

// Ensure that there's no over-draft, this is expected to happen when some
// transactions get included without publishing on the network

// Evict the highest nonce transactions until the pending set falls under
// the account's available balance

// only during reorgs will the heap be initialized

// Sanity check that no account can have more queued transactions than the
// DoS protection threshold.

// Evict the highest nonce transactions until the pending set falls under
// the account's transaction cap

// Included cheap transactions might have left the remaining ones better from
// an eviction point, fix any potential issues in the heap.

// offload removes a tracked blob transaction from the pool and moves it into the
// limbo for tracking until finality.
//
// The method may log errors for various unexpected scenarios but will not return
// any of it since there's no clear error case. Some errors may be due to coding
// issues, others caused by signers mining MEV stuff or swapping transactions. In
// all cases, the pool needs to continue operating.
func (p *BlobPool) offload(addr common.Address, nonce uint64, id uint64, inclusions map[common.Hash]uint64) {
	_ = "STUB: not implemented"
	return
}

// Reset implements txpool.SubPool, allowing the blob pool's internal state to be
// kept in sync with the main transaction pool's internal state.
func (p *BlobPool) Reset(oldHead, newHead *types.Header) { _ = "STUB: not implemented"; return }

// Run the reorg between the old and new head and figure out which accounts
// need to be rechecked and which transactions need to be readded

// Blindly push all the lost transactions back into the pool

// Recheck the account's pooled transactions to drop included and
// invalidated ones

// Flush out any blobs from limbo that are older than the latest finality

// Reset the price heap for the new set of basefee/blobfee pairs

// basefee = uint256.MustFromBig(eip1559.CalcBaseFee(p.chain.Config(), newHead))

// reorg assembles all the transactors and missing transactions between an old
// and new head to figure out which account's tx set needs to be rechecked and
// which transactions need to be requeued.
//
// The transactionblock inclusion infos are also returned to allow tracking any
// just-included blocks by block number in the limbo.
func (p *BlobPool) reorg(oldHead, newHead *types.Header) (map[common.Address][]*types.Transaction, map[common.Hash]uint64) {
	_ = "STUB: not implemented"
	// If the pool was not yet initialized, don't do anything
	return nil, nil
}

// If the reorg is too deep, avoid doing it (will happen during snap sync)

// Reorg seems shallow enough to pull in all transactions into memory

// if the new head is nil, it means that something happened between
// the firing of newhead-event and _now_: most likely a
// reorg caused by sync-reversion or explicit sethead back to an
// earlier block.

// This can happen if a setHead is performed, where we simply discard
// the old head from the chain. If that is the case, we don't have the
// lost transactions anymore, and there's nothing to add.

// If we reorged to a same or higher number, then it's not a case
// of setHead

// If the reorg ended up on a lower number, it's indicative of setHead
// being the cause

// Both old and new blocks exist, traverse through the progression chain
// and accumulate the transactors and transactions

// Generate the set of transactions per address to pull back into the pool,
// also updating the rest along the way

// Generate the set that was lost to reinject into the pool

// Update the set that was already reincluded to track the blocks in limbo

// reinject blindly pushes a transaction previously included in the chain - and
// just reorged out - into the pool. The transaction is assumed valid (having
// been in the chain), thus the only validation needed is nonce sorting and over-
// draft checks after injection.
//
// Note, the method will not initialize the eviction cache values as those will
// be done once for all transactions belonging to an account after all individual
// transactions are injected back into the pool.
func (p *BlobPool) reinject(addr common.Address, txhash common.Hash) error {
	_ = "STUB: not implemented"
	// Retrieve the associated blob from the limbo. Without the blobs, we cannot
	// add the transaction back into the pool as it is not mineable.
	return nil
}

// TODO: seems like an easy optimization here would be getting the serialized tx
// from limbo instead of re-serializing it here.

// Serialize the transaction back into the primary datastore.

// Update the indices and metrics

// SetGasTip implements txpool.SubPool, allowing the blob pool's gas requirements
// to be kept in sync with the main transaction pool's gas requirements.
func (p *BlobPool) SetGasTip(tip *big.Int) { _ = "STUB: not implemented"; return }

// Store the new minimum gas tip

// If the min miner fee increased, remove transactions below the new threshold

// Drop the offending transaction

// Drop everything afterwards, no gaps allowed

// Clear out the dropped transactions from the index

// Clear out the transactions from the data store

// validateTx checks whether a transaction is valid according to the consensus
// rules and adheres to some heuristic limits of the local node (price and size).
func (p *BlobPool) validateTx(tx *types.Transaction) error {
	_ = "STUB: not implemented"
	// Ensure the transaction adheres to basic pool filters (type, size, tip) and
	// consensus rules
	return nil
}

// Ensure the transaction adheres to the stateful pool filters (nonce, balance)

// Nonce gaps are not permitted in the blob pool, the first gap will
// be the next nonce shifted by however many transactions we already
// have pooled.

// If the transaction replaces an existing one, ensure that price bumps are
// adhered to.

// already validated above

// Account can support the replacement, but the price bump must also be met

// Has returns an indicator whether subpool has a transaction cached with the
// given hash.
func (p *BlobPool) Has(hash common.Hash) bool { _ = "STUB: not implemented"; return false }

func (p *BlobPool) HasLocal(hash common.Hash) bool {
	_ = "STUB: not implemented"
	// TODO: add support to check local transactions
	return false
}

// Get returns a transaction if it is contained in the pool, or nil otherwise.
func (p *BlobPool) Get(hash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	// Track the amount of time waiting to retrieve a fully resolved blob tx from
	// the pool and the amount of time actually spent on pulling the data from disk.
	return nil
}

// Pull the blob from disk and return an assembled response

// Add inserts a set of blob transactions into the pool if they pass validation (both
// consensus validity and pool restrictions).
func (p *BlobPool) Add(txs []*types.Transaction, local bool, sync bool) []error {
	_ = "STUB: not implemented"
	return nil
}

// Add inserts a new blob transaction into the pool if it passes validation (both
// consensus validity and pool restrictions).
func (p *BlobPool) add(tx *types.Transaction) (err error) {
	_ = "STUB: not implemented"
	// The blob pool blocks on adding a transaction. This is because blob txs are
	// only even pulled from the network, so this method will act as the overload
	// protection for fetches.
	return nil
}

// Ensure the transaction is valid from all perspectives

// If the address is not yet known, request exclusivity to track the account
// only by this subpool until all transactions are evicted
// already validated above

// If the transaction is rejected by some post-validation check, remove
// the lock on the reservation set.
//
// Note, `err` here is the named error return, which will be initialized
// by a return statement before running deferred methods. Take care with
// removing or subscoping err as it will break this clause.

// Transaction permitted into the pool from a nonce and cost perspective,
// insert it into the database and update the indices

// Transaction replaces a previously queued one

// Shitty situation, but try to recover gracefully instead of going boom

// Update the transaction index

// Transaction extends previously scheduled ones

// Recompute the rolling eviction fields. In case of a replacement, this will
// recompute all subsequent fields. In case of an append, this will only do
// the fresh calculation.

// The first transaction will always use itself

// Subsequent transactions will use a rolling calculation

// Update the eviction heap with the new information:
//   - If the transaction is from a new account, add it to the heap
//   - If the account had a singleton tx replaced, update the heap (new price caps)
//   - If the account has a transaction replaced or appended, update the heap if significantly changed

// 1 tx and not a new acc, must be replacement

// replacement or new append

// need math.Abs, can go up and down

// If the pool went over the allowed data limit, evict transactions until
// we're again below the threshold

// drop removes the worst transaction from the pool. It is primarily used when a
// freshly added transaction overflows the pool and needs to evict something. The
// method is also called on startup if the user resizes their storage, might be an
// expensive run but it should be fine-ish.
func (p *BlobPool) drop() {
	_ = "STUB: not implemented"
	// Peek at the account with the worse transaction set to evict from (Go's heap
	// stores the minimum at index zero of the heap slice) and retrieve it's last
	// transaction.
	return
}

// cannot call drop on empty pool

// Remove the transaction from the pool's index

// Remove the transaction from the pool's eviction heap:
//   - If the entire account was dropped, pop off the address
//   - Otherwise, if the new tail has better eviction caps, fix the heap

// new tail, surely exists

// no need for math.Abs, monotonic decreasing

// Remove the transaction from the data store

// Pending retrieves all currently processable transactions, grouped by origin
// account and sorted by nonce.
//
// The transactions can also be pre-filtered by the dynamic fee components to
// reduce allocations and load on downstream subsystems.
func (p *BlobPool) Pending(filter txpool.PendingFilter) map[common.Address][]*txpool.LazyTransaction {
	_ = "STUB: not implemented"
	// If only plain transactions are requested, this pool is unsuitable as it
	// contains none, don't even bother.
	return nil
}

// Track the amount of time waiting to retrieve the list of pending blob txs
// from the pool and the amount of time actually spent on assembling the data.
// The latter will be pretty much moot, but we've kept it to have symmetric
// across all user operations.

// If transaction filtering was requested, discard badly priced ones

// basefee too low, cannot be included, discard rest of txs from the account

// allowed or remaining tip too low, cannot be included, discard rest of txs from the account

// blobfee too low, cannot be included, discard rest of txs from the account

// Transaction was accepted according to the filter, append to the pending list

// TODO(karalabe): Maybe save these and use that?

// IteratePending iterates over [pool.pending] until [f] returns false.
// The caller must not modify [tx]. Returns false if iteration was interrupted.
func (pool *BlobPool) IteratePending(f func(tx *types.Transaction) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *BlobPool) SetMinFee(minFee *big.Int) {
	_ = "STUB: not implemented"

	// updateStorageMetrics retrieves a bunch of stats from the data store and pushes
	// them out as metrics.
	return
}

func (p *BlobPool) updateStorageMetrics() { _ = "STUB: not implemented"; return }

// updateLimboMetrics retrieves a bunch of stats from the limbo store and pushes
// them out as metrics.
func (p *BlobPool) updateLimboMetrics() { _ = "STUB: not implemented"; return }

// SubscribeTransactions registers a subscription for new transaction events,
// supporting feeding only newly seen or also resurrected transactions.
func (p *BlobPool) SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// Nonce returns the next nonce of an account, with all transactions executable
// by the pool already applied on top.
func (p *BlobPool) Nonce(addr common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

// Stats retrieves the current pool stats, namely the number of pending and the
// number of queued (non-executable) transactions.
func (p *BlobPool) Stats() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// No non-executable txs in the blob pool

// Content retrieves the data content of the transaction pool, returning all the
// pending as well as queued transactions, grouped by account and sorted by nonce.
//
// For the blob pool, this method will return nothing for now.
// TODO(karalabe): Abstract out the returned metadata.
func (p *BlobPool) Content() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContentFrom retrieves the data content of the transaction pool, returning the
// pending as well as queued transactions of this address, grouped by nonce.
//
// For the blob pool, this method will return nothing for now.
// TODO(karalabe): Abstract out the returned metadata.
func (p *BlobPool) ContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Locals retrieves the accounts currently considered local by the pool.
//
// There is no notion of local accounts in the blob pool.
func (p *BlobPool) Locals() []common.Address { _ = "STUB: not implemented"; return nil }

// Status returns the known status (unknown/pending/queued) of a transaction
// identified by their hashes.
func (p *BlobPool) Status(hash common.Hash) txpool.TxStatus {
	_ = "STUB: not implemented"
	return *new(txpool.TxStatus)
}
