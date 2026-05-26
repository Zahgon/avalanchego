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
// Copyright 2021 The go-ethereum Authors
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

package ethapi

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto/kzg4844"
	ethparams "github.com/ava-labs/libevm/params"
)

var (
	maxBlobsPerTransaction = ethparams.MaxBlobGasPerBlock / ethparams.BlobTxBlobGasPerBlob
)

// TransactionArgs represents the arguments to construct a new transaction
// or a message call.
type TransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"`

	// For BlobTxType
	BlobFeeCap *hexutil.Big  `json:"maxFeePerBlobGas"`
	BlobHashes []common.Hash `json:"blobVersionedHashes,omitempty"`

	// For BlobTxType transactions with blob sidecar
	Blobs       []kzg4844.Blob       `json:"blobs"`
	Commitments []kzg4844.Commitment `json:"commitments"`
	Proofs      []kzg4844.Proof      `json:"proofs"`

	// This configures whether blobs are allowed to be passed.
	blobSidecarAllowed bool
}

// from retrieves the transaction sender address.
func (args *TransactionArgs) from() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

// data retrieves the transaction calldata. Input field is preferred.
func (args *TransactionArgs) data() []byte { _ = "STUB: not implemented"; return nil }

// setDefaults fills in default values for unspecified tx fields.
func (args *TransactionArgs) setDefaults(ctx context.Context, b Backend, skipGasEstimation bool) error {
	_ = "STUB: not implemented"
	return nil
}

// BlobTx fields

// create check

// Skip gas usage estimation if a precise gas limit is not critical, e.g., in non-transaction calls.

// Estimate the gas usage otherwise.
// These fields are immutable during the estimation, safe to
// pass the pointer directly.

// If chain id is provided, ensure it matches the local chain id. Otherwise, set the local
// chain id as the default.

type feeBackend interface {
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	CurrentHeader() *types.Header
	ChainConfig() *params.ChainConfig
}

// setFeeDefaults fills in default fee values for unspecified tx fields.
func (args *TransactionArgs) setFeeDefaults(ctx context.Context, b feeBackend) error {
	_ = "STUB: not implemented"
	return nil

	// Sanity check the EIP-4844 fee parameters.
}

// If both gasPrice and at least one of the EIP-1559 fee parameters are specified, error.

// If the tx has completely specified a fee mechanism, no default is needed.
// This allows users who are not yet synced past London to get defaults for
// other tx values. See https://github.com/ethereum/go-ethereum/pull/23274
// for more information.

// Sanity check the EIP-1559 fee parameters if present.

// No need to set anything, user already set MaxFeePerGas and MaxPriorityFeePerGas

// Sanity check the non-EIP-1559 fee parameters.

// Zero gas-price is not allowed after London fork

// No need to set anything, user already set GasPrice

// Now attempt to fill in default value depending on whether London is active or not.

// London is active, set maxPriorityFeePerGas and maxFeePerGas.

// London not active, set gas price.

// setCancunFeeDefaults fills in reasonable default fee values for unspecified fields.
func (args *TransactionArgs) setCancunFeeDefaults(ctx context.Context, head *types.Header, b feeBackend) error {
	_ = "STUB: not implemented"
	// Set maxFeePerBlobGas if it is missing.
	return nil
}

// ExcessBlobGas must be set for a Cancun block.

// Set the max fee to be 2 times larger than the previous block's blob base fee.
// The additional slack allows the tx to not become invalidated if the base
// fee is rising.

// setSubnetEVMFeeDefault fills in reasonable default fee values for unspecified fields.
func (args *TransactionArgs) setSubnetEVMFeeDefault(ctx context.Context, head *types.Header, b feeBackend) error {
	_ = "STUB: not implemented"
	// Set maxPriorityFeePerGas if it is missing.
	return nil
}

// Set maxFeePerGas if it is missing.

// Set the max fee to be 2 times larger than the previous block's base fee.
// The additional slack allows the tx to not become invalidated if the base
// fee is rising.

// Both EIP-1559 fee parameters are now set; sanity check them.

// setBlobTxSidecar adds the blob tx
func (args *TransactionArgs) setBlobTxSidecar(ctx context.Context, b Backend) error {
	_ = "STUB: not implemented"
	// No blobs, we're done.
	return nil
}

// Passing blobs is not allowed in all contexts, only in specific methods.

// Assume user provides either only blobs (w/o hashes), or
// blobs together with commitments and proofs.

// len(blobs) == len(commitments) == len(proofs) == len(hashes)

// Generate commitment and proof.

// ToMessage converts the transaction arguments to the Message type used by the
// core evm. This method is used in calls and traces that do not require a real
// live transaction.
func (args *TransactionArgs) ToMessage(globalGasCap uint64, baseFee *big.Int) (*core.Message, error) {
	_ = "STUB: not implemented"
	// Reject invalid combinations of pre- and post-1559 fee styles
	return nil, nil
}

// Set sender address or use zero address if none specified.

// Set default gas & gas price if none were set

// If there's no basefee, then it must be a non-1559 execution

// A basefee is provided, necessitating 1559-type execution

// User specified the legacy gas field, convert to 1559 gas typing

// User specified 1559 gas fields (or none), use those

// Backfill the legacy gasPrice for EVM execution, unless we're all zeroes

// toTransaction converts the arguments to a transaction.
// This assumes that setDefaults has been called.
func (args *TransactionArgs) toTransaction() *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// IsEIP4844 returns an indicator if the args contains EIP4844 fields.
func (args *TransactionArgs) IsEIP4844() bool { _ = "STUB: not implemented"; return false }
