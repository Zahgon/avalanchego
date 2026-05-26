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

package eth

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/consensus"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/eth/gasprice"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/eth/tracers"
	"github.com/ava-labs/libevm/accounts"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/bloombits"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/params"
)

var ErrUnfinalizedData = errors.New("cannot query unfinalized data")

// EthAPIBackend implements ethapi.Backend and tracers.Backend for full nodes
type EthAPIBackend struct {
	extRPCEnabled            bool
	allowUnprotectedTxs      bool
	allowUnprotectedTxHashes map[common.Hash]struct{} // Invariant: read-only after creation.
	allowUnfinalizedQueries  bool
	eth                      *Ethereum
	gpo                      *gasprice.Oracle

	// historicalProofQueryWindow is the number of blocks before the last accepted block to be accepted for
	// state queries when running archive mode.
	historicalProofQueryWindow uint64
}

// ChainConfig returns the active chain configuration.
func (b *EthAPIBackend) ChainConfig() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

// IsArchive returns true if the node is running in archive mode, false otherwise.
func (b *EthAPIBackend) IsArchive() bool { _ = "STUB: not implemented"; return false }

// HistoricalProofQueryWindow returns the number of blocks before the last accepted block to be accepted for state queries.
// It returns 0 to indicate to accept any block number for state queries.
func (b *EthAPIBackend) HistoricalProofQueryWindow() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *EthAPIBackend) IsAllowUnfinalizedQueries() bool { _ = "STUB: not implemented"; return false }

func (b *EthAPIBackend) SetAllowUnfinalizedQueries(allow bool) { _ = "STUB: not implemented"; return }

func (b *EthAPIBackend) CurrentBlock() *types.Header { _ = "STUB: not implemented"; return nil }

func (b *EthAPIBackend) LastAcceptedBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (b *EthAPIBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Treat requests for the pending, latest, or accepted block
// identically.

func (b *EthAPIBackend) GetFeeConfigAt(parent *types.Header) (commontype.FeeConfig, *big.Int, error) {
	_ = "STUB: not implemented"
	return *new(commontype.FeeConfig), nil, nil
}

func (b *EthAPIBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Treat requests for the pending, latest, or accepted block
// identically.

func (b *EthAPIBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBody returns body of a block. It does not resolve special block numbers.
func (b *EthAPIBackend) GetBody(ctx context.Context, hash common.Hash, number rpc.BlockNumber) (*types.Body, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) BadBlocks() ([]*types.Block, []*core.BadBlockReason) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	_ = "STUB: not implemented"
	// Request the block by its number and retrieve its state
	return nil, nil, nil
}

// Firewood can reconstruct historical state by walking back to a
// persisted revision and re-executing blocks forward, even if the
// state is not available in-memory or on disk.
//
// TODO: propagate the release function through the [ethapi.Backend]
// interface for deterministic cleanup. Currently, it is discarded
// because the interface does not return one. The underlying reconstructed
// revision will be freed by the GC when stateDb becomes unreachable,
// since stateDb holds the only references to it.

func (b *EthAPIBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Firewood can reconstruct historical state by walking back to a
// persisted revision and re-executing blocks forward, even if the
// state is not available in-memory or on disk.
//
// TODO: propagate the release function through the [ethapi.Backend]
// interface for deterministic cleanup. Currently, it is discarded
// because the interface does not return one. The underlying reconstructed
// revision will be freed by the GC when stateDb becomes unreachable,
// since stateDb holds the only references to it.

func (b *EthAPIBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	_ = "STUB: not implemented"
	return *new(types.Receipts), nil
}

func (b *EthAPIBackend) GetLogs(ctx context.Context, hash common.Hash, number uint64) ([][]*types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) GetEVM(ctx context.Context, msg *core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config, blockCtx *vm.BlockContext) *vm.EVM {
	_ = "STUB: not implemented"
	return nil
}

func (b *EthAPIBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeChainAcceptedEvent(ch chan<- core.ChainEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeAcceptedLogsEvent(ch chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SubscribeAcceptedTransactionEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// We only enqueue transactions for push gossip if they were submitted over the RPC and
// added to the mempool.

func (b *EthAPIBackend) GetPoolTransactions() (types.Transactions, error) {
	_ = "STUB: not implemented"
	return *new(types.Transactions), nil
}

func (b *EthAPIBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func (b *EthAPIBackend) GetTransaction(ctx context.Context, txHash common.Hash) (bool, *types.Transaction, common.Hash, uint64, uint64, error) {
	_ = "STUB: not implemented"
	return false, nil, *new(common.Hash), 0, 0, nil
}

// Respond as if the transaction does not exist if it is not yet in an
// accepted block. We explicitly choose not to error here to avoid breaking
// expectations with clients (expect an empty response when a transaction
// does not exist).

func (b *EthAPIBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *EthAPIBackend) Stats() (runnable int, blocked int) { _ = "STUB: not implemented"; return 0, 0 }

func (b *EthAPIBackend) TxPoolContent() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) TxPoolContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (b *EthAPIBackend) EstimateBaseFee(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *EthAPIBackend) FeeHistory(ctx context.Context, blockCount uint64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (firstBlock *big.Int, reward [][]*big.Int, baseFee []*big.Int, gasUsedRatio []float64, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func (b *EthAPIBackend) ChainDb() ethdb.Database {
	_ = "STUB: not implemented"
	return *new(ethdb.Database)
}

func (b *EthAPIBackend) EventMux() *event.TypeMux { _ = "STUB: not implemented"; return nil }

func (b *EthAPIBackend) AccountManager() *accounts.Manager { _ = "STUB: not implemented"; return nil }

func (b *EthAPIBackend) ExtRPCEnabled() bool { _ = "STUB: not implemented"; return false }

func (b *EthAPIBackend) UnprotectedAllowed(tx *types.Transaction) bool {
	_ = "STUB: not implemented"
	return false
}

// Check for special cased transaction hashes:
// Note: this map is read-only after creation, so it is safe to read from it on multiple threads.

// Check for "predictable pattern" (Nick's Signature: https://weka.medium.com/how-to-send-ether-to-11-440-people-187e332566b7)

func (b *EthAPIBackend) RPCGasCap() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *EthAPIBackend) RPCEVMTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *EthAPIBackend) RPCTxFeeCap() float64 { _ = "STUB: not implemented"; return 0 }

func (b *EthAPIBackend) BloomStatus() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

func (b *EthAPIBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	_ = "STUB: not implemented"
	return
}

func (b *EthAPIBackend) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	return *new(consensus.Engine)
}

func (b *EthAPIBackend) CurrentHeader() *types.Header { _ = "STUB: not implemented"; return nil }

func (b *EthAPIBackend) GetMaxBlocksPerRequest() int64 { _ = "STUB: not implemented"; return 0 }

func (b *EthAPIBackend) StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(tracers.StateReleaseFunc), nil
}

func (b *EthAPIBackend) StateAtNextBlock(ctx context.Context, parent, nextBlock *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(tracers.StateReleaseFunc), nil
}

func (b *EthAPIBackend) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(vm.BlockContext), nil, *new(tracers.StateReleaseFunc), nil
}

func (b *EthAPIBackend) isLatestAndAllowed(number rpc.BlockNumber) bool {
	_ = "STUB: not implemented"
	return false
}
