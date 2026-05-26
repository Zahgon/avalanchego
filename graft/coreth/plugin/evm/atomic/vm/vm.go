// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"sync"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/coreth/consensus/dummy"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/txpool"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap5"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	avalanchedatabase "github.com/ava-labs/avalanchego/database"
	atomicstate "github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/state"
	avalanchegossip "github.com/ava-labs/avalanchego/network/p2p/gossip"
	avalanchecommon "github.com/ava-labs/avalanchego/snow/engine/common"
	avalancheutils "github.com/ava-labs/avalanchego/utils"
)

var (
	_ secp256k1fx.VM                     = (*VM)(nil)
	_ block.ChainVM                      = (*VM)(nil)
	_ block.BuildBlockWithContextChainVM = (*VM)(nil)
	_ block.StateSyncableVM              = (*VM)(nil)

	errAtomicGasExceedsLimit = errors.New("atomic gas used exceeds atomic gas limit")
)

const (
	secpCacheSize       = 1024
	defaultMempoolSize  = 4096
	targetAtomicTxsSize = 40 * units.KiB
	// maxAtomicTxMempoolGas is the maximum amount of gas that is allowed to be
	// used by an atomic transaction in the mempool. It is allowed to build
	// blocks with larger atomic transactions, but they will not be accepted
	// into the mempool.
	maxAtomicTxMempoolGas = ap5.AtomicGasLimit
	avaxEndpoint          = "/avax"
)

type VM struct {
	extension.InnerVM
	Ctx *snow.Context

	// TODO: unexport these fields
	SecpCache *secp256k1.RecoverCache
	Fx        secp256k1fx.Fx
	baseCodec codec.Registry

	// TODO: Remove Atomic prefix and unexport these fields
	AtomicMempool        *txpool.Mempool
	gossipHandler        p2p.Handler
	pullGossiper         *avalanchegossip.ValidatorGossiper
	AtomicTxPushGossiper *avalanchegossip.PushGossiper[*atomic.Tx]

	// AtomicTxRepository maintains two indexes on accepted atomic txs.
	// - txID to accepted atomic tx
	// - block height to list of atomic txs accepted on block at that height
	// TODO: unexport these fields
	AtomicTxRepository *atomicstate.AtomicRepository
	// AtomicBackend abstracts verification and processing of atomic transactions
	AtomicBackend *atomicstate.AtomicBackend

	// cancel may be nil until [snow.NormalOp] starts
	cancel     context.CancelFunc
	shutdownWg sync.WaitGroup

	clock        mockable.Clock
	bootstrapped avalancheutils.Atomic[bool]
}

func WrapVM(vm extension.InnerVM) *VM { _ = "STUB: not implemented"; return nil }

// Initialize implements the snowman.ChainVM interface
func (vm *VM) Initialize(
	ctx context.Context,
	chainCtx *snow.Context,
	db avalanchedatabase.Database,
	genesisBytes []byte,
	upgradeBytes []byte,
	configBytes []byte,
	fxs []*avalanchecommon.Fx,
	appSender avalanchecommon.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Set the chain config for mainnet/fuji chain IDs

// Free the memory of the extDataHash map

// Create the atomic extension structs
// some of them need to be initialized after the inner VM is initialized

// Create and pass the leaf handler to the atomic extension
// it will be initialized after the inner VM is initialized

// Initialize inner vm with the provided parameters

// initialize bonus blocks on mainnet

// initialize atomic repository

// Atomic backend is available now, we can initialize structs that depend on it

// so [vm.baseCodec] is a dummy codec use to fulfill the secp256k1fx VM
// interface. The fx will register all of its types, which can be safely
// ignored by the VM's codec.

func (vm *VM) SetState(ctx context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) onBootstrapStarted() error { _ = "STUB: not implemented"; return nil }

func (vm *VM) onNormalOperationsStarted() error { _ = "STUB: not implemented"; return nil }

func (vm *VM) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (vm *VM) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyTxAtTip verifies that [tx] is valid to be issued on top of the currently preferred block
func (vm *VM) verifyTxAtTip(tx *atomic.Tx) error { _ = "STUB: not implemented"; return nil }

// Note: we fetch the current block and then the state at that block instead of the current state directly
// since we need the header of the current block below.

// Return extremely detailed error since CalcBaseFee should never encounter an issue here

// We don’t need to revert the state here in case verifyTx errors, because
// preferredState is thrown away either way.

// verifyTx verifies that [tx] is valid to be issued into a block with parent block [parentHash]
// and validated at [state] using [rules] as the current rule set.
// Note: VerifyTx may modify [state]. If [state] needs to be properly maintained, the caller is responsible
// for reverting to the correct snapshot after calling this function. If this function is called with a
// throwaway state, then this is not necessary.
// TODO: unexport this function
func (vm *VM) verifyTx(tx *atomic.Tx, parentHash common.Hash, baseFee *big.Int, statedb *state.StateDB, rules extras.Rules) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyTxs verifies that [txs] are valid to be issued into a block with parent block [parentHash]
// using [rules] as the current rule set.
func (vm *VM) verifyTxs(txs []*atomic.Tx, parentHash common.Hash, baseFee *big.Int, height uint64, rules extras.Rules) error {
	_ = "STUB: not implemented"
	// Ensure that the parent was verified and inserted correctly.
	return nil
}

// If the ancestor is unknown, then the parent failed verification when
// it was called.
// If the ancestor is rejected, then this block shouldn't be inserted
// into the canonical chain because the parent will be missing.

// Ensure each tx in [txs] doesn't conflict with any other atomic tx in
// a processing ancestor block.

// CodecRegistry implements the secp256k1fx interface
func (vm *VM) CodecRegistry() codec.Registry {
	_ = "STUB: not implemented"
	return *

	// Clock implements the secp256k1fx interface
	new(codec.Registry)
}

func (vm *VM) Clock() *mockable.Clock {
	_ = "STUB: not implemented"

	// Logger implements the secp256k1fx interface
	return nil
}

func (vm *VM) Logger() logging.Logger { _ = "STUB: not implemented"; return *new(logging.Logger) }

func (vm *VM) createConsensusCallbacks() dummy.ConsensusCallbacks {
	_ = "STUB: not implemented"
	return *new(dummy.ConsensusCallbacks)
}

func (vm *VM) preBatchOnFinalizeAndAssemble(header *types.Header, state *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Take a snapshot of [state] before calling verifyTx so that if the transaction fails verification
// we can revert to [snapshot].
// Note: snapshot is taken inside the loop because you cannot revert to the same snapshot more than
// once.

// Discard the transaction from the mempool on failed verification.

// Discard the transaction from the mempool and error if the transaction
// cannot be marshalled. This should never happen.

// this could happen due to the async logic of geth tx pool

// assumes that we are in at least Apricot Phase 5.
func (vm *VM) postBatchOnFinalizeAndAssemble(
	header *types.Header,
	parent *types.Header,
	state *state.StateDB,
	txs []*types.Transaction,
) ([]byte, *big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Ensure that adding [tx] to the block will not exceed the block size soft limit.

// Note: we do not need to check if we are in at least ApricotPhase4 here because
// we assume that this function will only be called when the block is in at least
// ApricotPhase5.

// ensure `gasUsed + batchGasUsed` doesn't exceed `atomicGasLimit`

// Send [tx] back to the mempool's tx heap.

// Discard the transaction from the mempool since it will fail verification
// after this block has been accepted.
// Note: if the proposed block is not accepted, the transaction may still be
// valid, but we discard it early here based on the assumption that the proposed
// block will most likely be accepted.
// Discard the transaction from the mempool on failed verification.

// Discard the transaction from the mempool and reset the state to [snapshot]
// if it fails verification here.
// Note: prior to this point, we have not modified [state] so there is no need to
// revert to a snapshot if we discard the transaction prior to this point.

// Add the [txGasUsed] to the [batchGasUsed] when the [tx] has passed verification

// If there is a non-zero number of transactions, marshal them and return the byte slice
// for the block's extra data along with the contribution and gas used.

// If we fail to marshal the batch of atomic transactions for any reason,
// discard the entire set of current transactions.

// If there are no regular transactions and there were also no atomic transactions to be included,
// then the block is empty and should be considered invalid.

// this could happen due to the async logic of geth tx pool

// If there are no atomic transactions, but there is a non-zero number of regular transactions, then
// we return a nil slice with no contribution from the atomic transactions and a nil error.

func (vm *VM) onFinalizeAndAssemble(
	header *types.Header,
	parent *types.Header,
	state *state.StateDB,
	txs []*types.Transaction,
) ([]byte, *big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (vm *VM) onExtraStateChange(block *types.Block, parent *types.Header, statedb *state.StateDB) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// We cannot use chain config from InnerVM since it's not available when this function is called for the first time (bc.loadLastState).

// If [atomicBackend] is nil, the VM is still initializing and is reprocessing accepted blocks.

// Only verify and insert when processing at the chain tip; skip during historical replay.

// Verify [txs] do not conflict with themselves or ancestor blocks.

// Update the atomic backend with [txs] from this block.
//
// Note: The atomic trie canonically contains the duplicate operations
// from any bonus blocks.

// If there are no transactions, we can return early.

// If ApricotPhase4 is enabled, calculate the block fee contribution

// If ApricotPhase5 is enabled, enforce that the atomic gas used does not exceed the
// atomic gas limit.

func (vm *VM) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) BuildBlockWithContext(ctx context.Context, proposerVMBlockCtx *block.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// Handle errors and signal the mempool to take appropriate action

// Marks the current transactions from the mempool as being successfully issued
// into a block.

func (vm *VM) chainConfigExtra() *extras.ChainConfig { _ = "STUB: not implemented"; return nil }

func (vm *VM) rules(number *big.Int, time uint64) extras.Rules {
	_ = "STUB: not implemented"
	return *new(extras.Rules)
}

// CurrentRules returns the chain rules for the current block.
func (vm *VM) CurrentRules() extras.Rules { _ = "STUB: not implemented"; return *new(extras.Rules) }
