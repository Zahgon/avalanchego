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
// Copyright 2015 The go-ethereum Authors
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
//
// NOTE: this piece of code is modified by Ted Yin.
// The modification is also licensed under the same LGPL.

package miner

import (
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

const (
	// Leaves 256 KBs for other sections of the block (limit is 2MB).
	// This should suffice for atomic txs, proposervm header, and serialization overhead.
	targetTxsSize = 1792 * units.KiB
)

var ErrInsufficientGasCapacityToBuild = errors.New("insufficient gas capacity to build block")

// environment is the worker's current environment and holds all of the current state information.
type environment struct {
	signer  types.Signer
	state   *state.StateDB // apply state changes here
	tcount  int            // tx count in cycle
	gasPool *core.GasPool  // available gas used to pack transactions

	parent   *types.Header
	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt
	sidecars []*types.BlobTxSidecar
	blobs    int
	size     uint64

	rules            params.Rules
	predicateContext *precompileconfig.PredicateContext
	// predicateResults contains the results of checking the predicates for each transaction in the miner.
	// The results are accumulated as transactions are executed by the miner and set on the BlockContext.
	// If a transaction is dropped, its results must explicitly be removed from predicateResults in the same
	// way that the gas pool and state is reset.
	predicateResults predicate.BlockResults

	start time.Time // Time that block building began
}

// worker is the main object which takes care of submitting new work to consensus engine
// and gathering the sealing result.
type worker struct {
	config      *Config
	chainConfig *params.ChainConfig
	engine      consensus.Engine
	eth         Backend
	chain       *core.BlockChain

	// Feeds
	// TODO remove since this will never be written to
	pendingLogsFeed event.Feed

	// Subscriptions
	mux        *event.TypeMux // TODO replace
	mu         sync.RWMutex   // The lock used to protect the coinbase and extra fields
	coinbase   common.Address
	clock      *mockable.Clock // Allows us mock the clock for testing
	beaconRoot *common.Hash    // TODO: set to empty hash, retained for upstream compatibility and future use
}

func newWorker(config *Config, chainConfig *params.ChainConfig, engine consensus.Engine, eth Backend, mux *event.TypeMux, clock *mockable.Clock) *worker {
	_ = "STUB: not implemented"
	return nil
}

// setEtherbase sets the etherbase used to initialize the block coinbase field.
func (w *worker) setEtherbase(addr common.Address) { _ = "STUB: not implemented"; return }

// commitNewWork generates several new sealing tasks based on the parent block.
func (w *worker) commitNewWork(predicateContext *precompileconfig.PredicateContext) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply EIP-4844, EIP-4788.

// For the first post-fork block, both parent.data_gas_used and parent.excess_data_gas are evaluated as 0

// Ensure we always stop prefetcher after block building is complete.

// Configure any upgrades that should go into effect during this block.

// Retrieve the pending transactions pre-filtered by the 1559/4844 dynamic fees

// Split the pending transactions into locals and remotes.

// Fill the block with all available pending transactions.

func (w *worker) createCurrentEnvironment(predicateContext *precompileconfig.PredicateContext, parent *types.Header, header *types.Header, tstart time.Time) (*environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If Fortuna has activated, enforce we only build a block when there is a minimum
// built up gas capacity.

// ACP-176 maintains a gas capacity bucket that fills up over time separate from the gas limit.
// Since smaller transactions can consume the available gas capacity every second, this can lead
// to transactions above this size stalling.
// To mitigate this, the block builder returns an error and waits until the gas capacity bucket
// has refilled to the desired minimum.
// We set that minimum to MaxCapacity - GasCapacityPerSecond to be available before building a block.
// This avoids waiting past the point where the gas capacity bucket is already full. If we waited
// until the MaxCapacity, then waiting beyond that point would "overflow" the bucket. Since we do
// not allow the bucket to overflow, this would waste potential gas capacity.

// Since the GasLimit is configurable, only wait up to the GasLimit as of the Cortina upgrade.
// There is no need to continue scaling up beyond the GasLimit guaranteed prior to ACP-176 activation.

func (w *worker) commitTransaction(env *environment, tx *types.Transaction, coinbase common.Address) ([]*types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *worker) commitBlobTransaction(env *environment, tx *types.Transaction, coinbase common.Address) ([]*types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Checking against blob gas limit: It's kind of ugly to perform this check here, but there
// isn't really a better place right now. The blob gas limit is checked at block validation time
// and not during execution. This means core.ApplyTransaction will not return an error if the
// tx has too many blobs. So we have to explicitly check it here.

// applyTransaction runs the transaction. If execution fails, state and gas pool are reverted.
func (w *worker) applyTransaction(env *environment, tx *types.Transaction, coinbase common.Address) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete results by setting to nil

func (w *worker) commitTransactions(env *environment, plainTxs, blobTxs *transactionsByPriceAndNonce, coinbase common.Address) {
	_ = "STUB: not implemented"

	// If we don't have enough gas for any further transactions then we're done.
	return
}

// If we don't have enough blob space for any further blob transactions,
// skip that list altogether

// Fall though to pick up any plain txs

// If we don't have enough blob space for any further blob transactions,
// skip that list altogether

// Fall though to pick up any plain txs

// Retrieve the next transaction and abort if all done.

// If we don't have enough space for the next transaction, skip the account.

// Transaction seems to fit, pull it up from the pool

// Abort transaction if it won't fit in the block and continue to search for a smaller
// transction that will fit.

// Error may be ignored here. The error has already been checked
// during transaction acceptance is the transaction pool.

// Check whether the tx is replay protected. If we're not in the EIP155 hf
// phase, start ignoring the sender until we do.

// Start executing the transaction

// New head notification data race between the transaction pool and miner, shift

// Transaction is regarded as invalid, drop all consecutive transactions from
// the same sender because of `nonce-too-high` clause.

// commit runs any post-transaction state modifications, assembles the final block
// and commits new work if consensus engine is running.
func (w *worker) commit(env *environment) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deep copy receipts here to avoid interaction between different tasks.

func (w *worker) handleResult(env *environment, block *types.Block, createdAt time.Time, unfinishedReceipts []*types.Receipt) (*types.Block, error) {
	_ = "STUB: not implemented"
	// Short circuit when receiving duplicate result caused by resubmitting.
	return nil, nil
}

// Different block could share same sealhash, deep copy here to prevent write-write conflict.

// add block location fields

// Update the block hash in all logs since it is now available and not when the
// receipt/log of individual transactions were created.

// Note: the miner no longer emits a NewMinedBlock event. Instead the caller
// is responsible for running any additional verification and then inserting
// the block with InsertChain, which will also emit a new head event.

// copyReceipts makes a deep copy of the given receipts.
func copyReceipts(receipts []*types.Receipt) []*types.Receipt {
	_ = "STUB: not implemented"
	return nil
}

// totalFees computes total consumed miner fees in Wei. Block transactions and receipts have to have the same order.
func totalFees(block *types.Block, receipts []*types.Receipt) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// Note in coreth the coinbase payment is (baseFee + effectiveGasTip) * gasUsed

// Prior to activation of EIP-1559, the coinbase payment was gasPrice * gasUsed
