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

package miner

import (
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/core/txpool"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/holiman/uint256"
)

// txWithMinerFee wraps a transaction with its gas price or effective miner gasTipCap
type txWithMinerFee struct {
	tx   *txpool.LazyTransaction
	from common.Address
	fees *uint256.Int
}

// newTxWithMinerFee creates a wrapped transaction, calculating the effective
// miner gasTipCap if a base fee is provided.
// Returns error in case of a negative effective miner gasTipCap.
func newTxWithMinerFee(tx *txpool.LazyTransaction, from common.Address, baseFee *uint256.Int) (*txWithMinerFee, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// txByPriceAndTime implements both the sort and the heap interface, making it useful
// for all at once sorting as well as individually adding and removing elements.
type txByPriceAndTime []*txWithMinerFee

func (s txByPriceAndTime) Len() int { _ = "STUB: not implemented"; return 0 }
func (s txByPriceAndTime) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// If the prices are equal, use the time the transaction was first seen for
	// deterministic sorting
	return false
}

func (s txByPriceAndTime) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s *txByPriceAndTime) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (s *txByPriceAndTime) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// transactionsByPriceAndNonce represents a set of transactions that can return
// transactions in a profit-maximizing sorted order, while supporting removing
// entire batches of transactions for non-executable accounts.
type transactionsByPriceAndNonce struct {
	txs     map[common.Address][]*txpool.LazyTransaction // Per account nonce-sorted list of transactions
	heads   txByPriceAndTime                             // Next transaction for each unique account (price heap)
	signer  types.Signer                                 // Signer for the set of transactions
	baseFee *uint256.Int                                 // Current base fee
}

// newTransactionsByPriceAndNonce creates a transaction set that can retrieve
// price sorted transactions in a nonce-honouring way.
//
// Note, the input map is reowned so the caller should not interact any more with
// if after providing it to the constructor.
func newTransactionsByPriceAndNonce(signer types.Signer, txs map[common.Address][]*txpool.LazyTransaction, baseFee *big.Int) *transactionsByPriceAndNonce {
	_ = "STUB: not implemented"
	// Convert the basefee from header format to uint256 format
	return nil
}

// Initialize a price and received time based heap with the head transactions

// Assemble and return the transaction set

// Peek returns the next transaction by price.
func (t *transactionsByPriceAndNonce) Peek() (*txpool.LazyTransaction, *uint256.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shift replaces the current best head with the next one from the same account.
func (t *transactionsByPriceAndNonce) Shift() { _ = "STUB: not implemented"; return }

// Pop removes the best transaction, *not* replacing it with the next one from
// the same account. This should be used when a transaction cannot be executed
// and hence all subsequent ones should be discarded from the same account.
func (t *transactionsByPriceAndNonce) Pop() {
	_ = "STUB: not implemented"

	// Empty returns if the price heap is empty. It can be used to check it simpler
	// than calling peek and checking for nil return.
	return
}

func (t *transactionsByPriceAndNonce) Empty() bool { _ = "STUB: not implemented"; return false }

// Clear removes the entire content of the heap.
func (t *transactionsByPriceAndNonce) Clear() { _ = "STUB: not implemented"; return }
