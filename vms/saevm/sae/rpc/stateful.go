// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/consensus"
	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/eth/tracers"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/vms/saevm/hook"
)

var noopRelease tracers.StateReleaseFunc = func() {}

// noEndOfBlockOps wraps [hook.Points] to suppress
// [hook.Points.EndOfBlockOps], used by the tracer to skip end-of-block
// operations during partial replay.
type noEndOfBlockOps struct {
	hook.Points
}

// EndOfBlockOps always returns nil.
func (noEndOfBlockOps) EndOfBlockOps(*types.Block) ([]hook.Op, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) RPCEVMTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *backend) RPCGasCap() uint64 { _ = "STUB: not implemented"; return 0 }

func (*backend) Engine() consensus.Engine { _ = "STUB: not implemented"; return *new(consensus.Engine) }

type coinbaseAsAuthor struct {
	consensus.Engine
}

func (*coinbaseAsAuthor) Author(h *types.Header) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (b *backend) GetEVM(ctx context.Context, msg *core.Message, sdb *state.StateDB, hdr *types.Header, cfg *vm.Config, bCtx *vm.BlockContext) *vm.EVM {
	_ = "STUB: not implemented"
	return nil
}

// StateAndHeaderByNumber performs the same faking as
// [backend.StateAndHeaderByNumberOrHash].
func (b *backend) StateAndHeaderByNumber(ctx context.Context, num rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// StateAndHeaderByNumberOrHash fakes the returned [types.Header] to contain
// post-execution results, mimicking a synchronous block. The [state.StateDB] is
// opened at the post-execution root, as carried by the faked header.
func (b *backend) StateAndHeaderByNumberOrHash(ctx context.Context, numOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// The API implementations expect this to be synchronous, sourcing the state
// root and the base fee from fields. At the time of writing, the returned
// header's hash is never used so it's safe to modify it.
//
// TODO(arr4n) the above assumption is brittle under geth/libevm updates;
// devise an approach to ensure that it is confirmed on each.

// TODO(arr4n) export [blocks.executionResults] to avoid multiple
// database reads and canoto unmarshallings here.

// StateAtBlock returns the state database after executing the given block. The
// reexec, base, readOnly, and preferDisk parameters are ignored because SAE
// does not implement geth's re-execution-from-archive strategy.
//
// Like geth, SAE only stores historical state roots, not full historical state.
// The underlying trie data must still be present in the state cache/DB for
// [state.New] to succeed. This means tracing is limited to recent blocks whose
// trie data has not been pruned (or requires an archival node for older blocks).
//
// Reference: https://geth.ethereum.org/docs/developers/evm-tracing#state-availability
//
//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func (b *backend) StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(tracers.StateReleaseFunc), nil
}

// StateAtTransaction returns the execution environment of a particular
// transaction within a block. It replays all preceding transactions to produce
// the state just before the target transaction, then returns the message and
// block context needed for tracing.
//
// Replay calls [saexec.Execute] - the same pipeline used by
// [saexec.Executor] - with [noEndOfBlockOps] to suppress end-of-block
// operations and [saexec.NullReceiptStore] to skip receipt broadcasting.
//
//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func (b *backend) StateAtTransaction(ctx context.Context, ethB *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(vm.BlockContext), nil, *new(tracers.StateReleaseFunc), nil
}

// The I(E) check above guarantees D(A) of the same block; see
// ../docs/invariants.md for details.

// Ancestry is irrelevant for the parent as we just want its
// post-execution artefacts.

// Replay transactions 0..txIndex-1 to produce the state just before the
// target transaction.

// postExecutionStateRoot returns the post-execution state root for the block
// identified by hash and number, checking in-memory blocks first, then falling
// back to disk.
func (b *backend) postExecutionStateRoot(hash common.Hash, num uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}
