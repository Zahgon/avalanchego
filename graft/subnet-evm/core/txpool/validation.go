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
	"math/big"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	ethparams "github.com/ava-labs/libevm/params"
)

var (
	// blobTxMinBlobGasPrice is the big.Int version of the configured protocol
	// parameter to avoid constucting a new big integer for every transaction.
	blobTxMinBlobGasPrice = big.NewInt(ethparams.BlobTxMinBlobGasprice)
)

// ValidationOptions define certain differences between transaction validation
// across the different pools without having to duplicate those checks.
type ValidationOptions struct {
	Config *params.ChainConfig // Chain configuration to selectively validate based on current fork rules

	Accept  uint8    // Bitmap of transaction types that should be accepted for the calling pool
	MaxSize uint64   // Maximum size of a transaction that the caller can meaningfully handle
	MinTip  *big.Int // Minimum gas tip needed to allow a transaction into the caller pool
}

// ValidateTransaction is a helper method to check whether a transaction is valid
// according to the consensus rules, but does not check state-dependent validation
// (balance, nonce, etc).
//
// This check is public to allow different transaction pools to check the basic
// rules without duplicating code and running the risk of missed updates.
func ValidateTransaction(tx *types.Transaction, head *types.Header, signer types.Signer, opts *ValidationOptions) error {
	_ = "STUB: not implemented"
	// Ensure transactions not implemented by the calling pool are rejected
	return nil
}

// Before performing any expensive validations, sanity check that the tx is
// smaller than the maximum limit the pool can meaningfully handle

// Ensure only transactions that have been enabled are accepted

// Check whether the init code size has been exceeded

// Transactions can't be negative. This may never happen using RLP decoded
// transactions but may occur for transactions created using the RPC.

// Ensure the transaction doesn't exceed the current block limit gas

// Sanity check for extremely large numbers (supported by RLP or RPC)

// Ensure gasFeeCap is greater than or equal to gasTipCap

// Make sure the transaction is signed properly

// Ensure the transaction has more gas than the bare minimum needed to cover
// the transaction metadata

// Ensure the gasprice is high enough to cover the requirement of the calling pool

// Ensure the blob fee cap satisfies the minimum blob gas price

// Ensure the number of items in the blob transaction and various side
// data match up before doing any expensive validations

// Ensure commitments, proofs and hashes are valid

func validateBlobSidecar(hashes []common.Hash, sidecar *types.BlobTxSidecar) error {
	_ = "STUB: not implemented"
	return nil
}

// Blob quantities match up, validate that the provers match with the
// transaction hash before getting to the cryptography

// Blob commitments match with the hashes in the transaction, verify the
// blobs themselves via KZG

// ValidationOptionsWithState define certain differences between stateful transaction
// validation across the different pools without having to duplicate those checks.
type ValidationOptionsWithState struct {
	State *state.StateDB // State database to check nonces and balances against

	// FirstNonceGap is an optional callback to retrieve the first nonce gap in
	// the list of pooled transactions of a specific account. If this method is
	// set, nonce gaps will be checked and forbidden. If this method is not set,
	// nonce gaps will be ignored and permitted.
	FirstNonceGap func(addr common.Address) uint64

	// UsedAndLeftSlots is a mandatory callback to retrieve the number of tx slots
	// used and the number still permitted for an account. New transactions will
	// be rejected once the number of remaining slots reaches zero.
	UsedAndLeftSlots func(addr common.Address) (int, int)

	// ExistingExpenditure is a mandatory callback to retrieve the cumulative
	// cost of the already pooled transactions to check for overdrafts.
	ExistingExpenditure func(addr common.Address) *big.Int

	// ExistingCost is a mandatory callback to retrieve an already pooled
	// transaction's cost with the given nonce to check for overdrafts.
	ExistingCost func(addr common.Address, nonce uint64) *big.Int

	Rules      params.Rules
	MinimumFee *big.Int
}

// ValidateTransactionWithState is a helper method to check whether a transaction
// is valid according to the pool's internal state checks (balance, nonce, gaps).
//
// This check is public to allow different transaction pools to check the stateful
// rules without duplicating code and running the risk of missed updates.
func ValidateTransactionWithState(tx *types.Transaction, signer types.Signer, opts *ValidationOptionsWithState) error {
	_ = "STUB: not implemented"
	// Ensure the transaction adheres to nonce ordering
	return nil
}

// already validated (and cached), but cleaner to check

// Drop the transaction if the gas fee cap is below the pool's minimum fee

// Ensure the transaction doesn't produce a nonce gap in pools that do not
// support arbitrary orderings

// Ensure the transactor has enough funds to cover the transaction costs

// Ensure the transactor has enough funds to cover for replacements or nonce
// expansions without overdrafts

// Transaction takes a new nonce value out of the pool. Ensure it doesn't
// overflow the number of permitted transactions from a single account
// (i.e. max cancellable via out-of-bound transaction).

// If the tx allow list is enabled, return an error if the from address is not allow listed.
