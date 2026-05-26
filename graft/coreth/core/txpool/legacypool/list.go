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
// Copyright 2016 The go-ethereum Authors
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

package legacypool

import (
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ava-labs/libevm/core/types"
	"github.com/holiman/uint256"
)

// nonceHeap is a heap.Interface implementation over 64bit unsigned integers for
// retrieving sorted transactions from the possibly gapped future queue.
type nonceHeap []uint64

func (h nonceHeap) Len() int           { _ = "STUB: not implemented"; return 0 }
func (h nonceHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (h nonceHeap) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *nonceHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (h *nonceHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// sortedMap is a nonce->transaction hash map with a heap based index to allow
// iterating over the contents in a nonce-incrementing way.
type sortedMap struct {
	items   map[uint64]*types.Transaction // Hash map storing the transaction data
	index   *nonceHeap                    // Heap of nonces of all the stored transactions (non-strict mode)
	cache   types.Transactions            // Cache of the transactions already sorted
	cacheMu sync.Mutex                    // Mutex covering the cache
}

// newSortedMap creates a new nonce-sorted transaction map.
func newSortedMap() *sortedMap { _ = "STUB: not implemented"; return nil }

// Get retrieves the current transactions associated with the given nonce.
func (m *sortedMap) Get(nonce uint64) *types.Transaction { _ = "STUB: not implemented"; return nil }

// Put inserts a new transaction into the map, also updating the map's nonce
// index. If a transaction already exists with the same nonce, it's overwritten.
func (m *sortedMap) Put(tx *types.Transaction) { _ = "STUB: not implemented"; return }

// Forward removes all transactions from the map with a nonce lower than the
// provided threshold. Every removed transaction is returned for any post-removal
// maintenance.
func (m *sortedMap) Forward(threshold uint64) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// Pop off heap items until the threshold is reached

// If we had a cached order, shift the front

// Filter iterates over the list of transactions and removes all of them for which
// the specified function evaluates to true.
// Filter, as opposed to 'filter', re-initialises the heap after the operation is done.
// If you want to do several consecutive filterings, it's therefore better to first
// do a .filter(func1) followed by .Filter(func2) or reheap()
func (m *sortedMap) Filter(filter func(*types.Transaction) bool) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// If transactions were removed, the heap and cache are ruined

func (m *sortedMap) reheap() { _ = "STUB: not implemented"; return }

// filter is identical to Filter, but **does not** regenerate the heap. This method
// should only be used if followed immediately by a call to Filter or reheap()
func (m *sortedMap) filter(filter func(*types.Transaction) bool) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// Collect all the transactions to filter out

// Cap places a hard limit on the number of items, returning all transactions
// exceeding that limit.
func (m *sortedMap) Cap(threshold int) types.Transactions {
	_ = "STUB: not implemented"
	// Short circuit if the number of items is under the limit
	return *new(types.Transactions)
}

// Otherwise gather and drop the highest nonce'd transactions

// The sorted m.index slice is still a valid heap, so there is no need to
// reheap after deleting tail items.

// If we had a cache, shift the back

// Remove deletes a transaction from the maintained map, returning whether the
// transaction was found.
func (m *sortedMap) Remove(nonce uint64) bool {
	_ = "STUB: not implemented"
	// Short circuit if no transaction is present
	return false
}

// Otherwise delete the transaction and fix the heap index

// Ready retrieves a sequentially increasing list of transactions starting at the
// provided nonce that is ready for processing. The returned transactions will be
// removed from the list.
//
// Note, all transactions with nonces lower than start will also be returned to
// prevent getting into an invalid state. This is not something that should ever
// happen but better to be self correcting than failing!
func (m *sortedMap) Ready(start uint64) types.Transactions {
	_ = "STUB: not implemented"
	// Short circuit if no transactions are available
	return *new(types.Transactions)
}

// Otherwise start accumulating incremental transactions

// Len returns the length of the transaction map.
func (m *sortedMap) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *sortedMap) flatten() types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// If the sorting was not cached yet, create and cache it

// Flatten creates a nonce-sorted slice of transactions based on the loosely
// sorted internal representation. The result of the sorting is cached in case
// it's requested again before any modifications are made to the contents.
func (m *sortedMap) Flatten() types.Transactions {
	_ = "STUB: not implemented"
	return *

	// Copy the cache to prevent accidental modification
	new(types.Transactions)
}

// LastElement returns the last element of a flattened list, thus, the
// transaction with the highest nonce
func (m *sortedMap) LastElement() *types.Transaction { _ = "STUB: not implemented"; return nil }

// list is a "list" of transactions belonging to an account, sorted by account
// nonce. The same type can be used both for storing contiguous transactions for
// the executable/pending queue; and for storing gapped transactions for the non-
// executable/future queue, with minor behavioral changes.
type list struct {
	strict bool       // Whether nonces are strictly continuous or not
	txs    *sortedMap // Heap indexed sorted hash map of the transactions

	costcap   *uint256.Int // Price of the highest costing transaction (reset only if exceeds balance)
	gascap    uint64       // Gas limit of the highest spending transaction (reset only if exceeds block limit)
	totalcost *uint256.Int // Total cost of all transactions in the list
}

// newList creates a new transaction list for maintaining nonce-indexable fast,
// gapped, sortable transaction lists.
func newList(strict bool) *list { _ = "STUB: not implemented"; return nil }

// Contains returns whether the  list contains a transaction
// with the provided nonce.
func (l *list) Contains(nonce uint64) bool { _ = "STUB: not implemented"; return false }

// Add tries to insert a new transaction into the list, returning whether the
// transaction was accepted, and if yes, any previous transaction it replaced.
//
// If the new transaction is accepted into the list, the lists' cost and gas
// thresholds are also potentially updated.
func (l *list) Add(tx *types.Transaction, priceBump uint64) (bool, *types.Transaction) {
	_ = "STUB: not implemented"
	// If there's an older better transaction, abort
	return false, nil
}

// thresholdFeeCap = oldFC  * (100 + priceBump) / 100

// thresholdTip    = oldTip * (100 + priceBump) / 100

// We have to ensure that both the new fee cap and tip are higher than the
// old ones as well as checking the percentage threshold to ensure that
// this is accurate for low (Wei-level) gas price replacements.

// Old is being replaced, subtract old cost

// Add new tx cost to totalcost

// Otherwise overwrite the old transaction with the current one

// Forward removes all transactions from the list with a nonce lower than the
// provided threshold. Every removed transaction is returned for any post-removal
// maintenance.
func (l *list) Forward(threshold uint64) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// Filter removes all transactions from the list with a cost or gas limit higher
// than the provided thresholds. Every removed transaction is returned for any
// post-removal maintenance. Strict-mode invalidated transactions are also
// returned.
//
// This method uses the cached costcap and gascap to quickly decide if there's even
// a point in calculating all the costs or if the balance covers all. If the threshold
// is lower than the costgas cap, the caps will be reset to a new high after removing
// the newly invalidated transactions.
func (l *list) Filter(costLimit *uint256.Int, gasLimit uint64) (types.Transactions, types.Transactions) {
	_ = "STUB: not implemented"
	// If all transactions are below the threshold, short circuit
	return *new(types.Transactions), *new(types.Transactions)
}

// Lower the caps to the thresholds

// Filter out all the transactions above the account's funds

// If the list was strict, filter anything above the lowest nonce

// Reset total cost

// Cap places a hard limit on the number of items, returning all transactions
// exceeding that limit.
func (l *list) Cap(threshold int) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// Remove deletes a transaction from the maintained list, returning whether the
// transaction was found, and also returning any transaction invalidated due to
// the deletion (strict mode only).
func (l *list) Remove(tx *types.Transaction) (bool, types.Transactions) {
	_ = "STUB: not implemented"
	// Remove the transaction from the set
	return false, *new(types.Transactions)
}

// In strict mode, filter out non-executable transactions

// Ready retrieves a sequentially increasing list of transactions starting at the
// provided nonce that is ready for processing. The returned transactions will be
// removed from the list.
//
// Note, all transactions with nonces lower than start will also be returned to
// prevent getting into an invalid state. This is not something that should ever
// happen but better to be self correcting than failing!
func (l *list) Ready(start uint64) types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}

// Len returns the length of the transaction list.
func (l *list) Len() int {
	_ = "STUB: not implemented"

	// Empty returns whether the list of transactions is empty or not.
	return 0
}

func (l *list) Empty() bool { _ = "STUB: not implemented"; return false }

// Flatten creates a nonce-sorted slice of transactions based on the loosely
// sorted internal representation. The result of the sorting is cached in case
// it's requested again before any modifications are made to the contents.
func (l *list) Flatten() types.Transactions {
	_ = "STUB: not implemented"
	return *

	// LastElement returns the last element of a flattened list, thus, the
	// transaction with the highest nonce
	new(types.Transactions)
}

func (l *list) LastElement() *types.Transaction { _ = "STUB: not implemented"; return nil }

// subTotalCost subtracts the cost of the given transactions from the
// total cost of all transactions.
func (l *list) subTotalCost(txs []*types.Transaction) { _ = "STUB: not implemented"; return }

// priceHeap is a heap.Interface implementation over transactions for retrieving
// price-sorted transactions to discard when the pool fills up. If baseFee is set
// then the heap is sorted based on the effective tip based on the given base fee.
// If baseFee is nil then the sorting is based on gasFeeCap.
type priceHeap struct {
	baseFee *big.Int // heap should always be re-sorted after baseFee is changed
	list    []*types.Transaction
}

func (h *priceHeap) Len() int      { _ = "STUB: not implemented"; return 0 }
func (h *priceHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (h *priceHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (h *priceHeap) cmp(a, b *types.Transaction) int { _ = "STUB: not implemented"; return 0 }

// Compare effective tips if baseFee is specified

// Compare fee caps if baseFee is not specified or effective tips are equal

// Compare tips if effective tips and fee caps are equal

func (h *priceHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (h *priceHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// pricedList is a price-sorted heap to allow operating on transactions pool
// contents in a price-incrementing way. It's built upon the all transactions
// in txpool but only interested in the remote part. It means only remote transactions
// will be considered for tracking, sorting, eviction, etc.
//
// Two heaps are used for sorting: the urgent heap (based on effective tip in the next
// block) and the floating heap (based on gasFeeCap). Always the bigger heap is chosen for
// eviction. Transactions evicted from the urgent heap are first demoted into the floating heap.
// In some cases (during a congestion, when blocks are full) the urgent heap can provide
// better candidates for inclusion while in other cases (at the top of the baseFee peak)
// the floating heap is better. When baseFee is decreasing they behave similarly.
type pricedList struct {
	// Number of stale price points to (re-heap trigger).
	stales atomic.Int64

	all              *lookup    // Pointer to the map of all transactions
	urgent, floating priceHeap  // Heaps of prices of all the stored **remote** transactions
	reheapMu         sync.Mutex // Mutex asserts that only one routine is reheaping the list
}

const (
	// urgentRatio : floatingRatio is the capacity ratio of the two queues
	urgentRatio   = 4
	floatingRatio = 1
)

// newPricedList creates a new price-sorted transaction heap.
func newPricedList(all *lookup) *pricedList { _ = "STUB: not implemented"; return nil }

// Put inserts a new transaction into the heap.
func (l *pricedList) Put(tx *types.Transaction, local bool) { _ = "STUB: not implemented"; return }

// Insert every new transaction to the urgent heap first; Discard will balance the heaps

// Removed notifies the prices transaction list that an old transaction dropped
// from the pool. The list will just keep a counter of stale objects and update
// the heap if a large enough ratio of transactions go stale.
func (l *pricedList) Removed(count int) {
	_ = "STUB: not implemented"
	// Bump the stale counter, but exit if still too low (< 25%)
	return
}

// Seems we've reached a critical number of stale transactions, reheap

// Underpriced checks whether a transaction is cheaper than (or as cheap as) the
// lowest priced (remote) transaction currently being tracked.
func (l *pricedList) Underpriced(tx *types.Transaction) bool {
	_ = "STUB: not implemented"
	// Note: with two queues, being underpriced is defined as being worse than the worst item
	// in all non-empty queues if there is any. If both queues are empty then nothing is underpriced.
	return false
}

// underpricedFor checks whether a transaction is cheaper than (or as cheap as) the
// lowest priced (remote) transaction in the given heap.
func (l *pricedList) underpricedFor(h *priceHeap, tx *types.Transaction) bool {
	_ = "STUB: not implemented"
	// Discard stale price points if found at the heap start
	return false
}

// Removed or migrated

// Check if the transaction is underpriced or not

// There is no remote transaction at all.

// If the remote transaction is even cheaper than the
// cheapest one tracked locally, reject it.

// Discard finds a number of most underpriced transactions, removes them from the
// priced list and returns them for further removal from the entire pool.
// If noPending is set to true, we will only consider the floating list
//
// Note local transaction won't be considered for eviction.
func (l *pricedList) Discard(slots int, force bool) (types.Transactions, bool) {
	_ = "STUB: not implemented"
	return *new(types.Transactions), false
}

// Remote underpriced transactions to drop

// Discard stale transactions if found during cleanup

// Removed or migrated

// Non stale transaction found, move to floating heap

// Stop if both heaps are empty

// Discard stale transactions if found during cleanup

// Removed or migrated

// Non stale transaction found, discard it

// If we still can't make enough room for the new transaction

// Reheap forcibly rebuilds the heap based on the current remote transaction set.
func (l *pricedList) Reheap() { _ = "STUB: not implemented"; return }

// Only iterate remotes

// balance out the two heaps by moving the worse half of transactions into the
// floating heap
// Note: Discard would also do this before the first eviction but Reheap can do
// is more efficiently. Also, Underpriced would work suboptimally the first time
// if the floating queue was empty.

// SetBaseFee updates the base fee and triggers a re-heap. Note that Removed is not
// necessary to call right before SetBaseFee when processing a new block.
func (l *pricedList) SetBaseFee(baseFee *big.Int) { _ = "STUB: not implemented"; return }
