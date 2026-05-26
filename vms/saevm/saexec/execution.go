// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saexec

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/libevm/eventual"
	"github.com/ava-labs/libevm/params"
	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/gastime"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/saedb"
)

var errExecutorClosed = errors.New("saexec.Executor closed")

// Enqueue pushes a new block to the FIFO queue. If [Executor.Close] is called
// before [blocks.Block.Executed] returns true then there is no guarantee that
// the block will be executed.
func (e *Executor) Enqueue(ctx context.Context, block *blocks.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// If this happens then increase the channel's buffer size.

// `e.done` can also close due to [Executor.execute] errors.

const emergencyPlaybookLink = "https://github.com/ava-labs/strevm/issues/28"

func (e *Executor) processQueue() { _ = "STUB: not implemented"; return }

//nolint:gocritic // False positive, will not terminate the process

var errFatal = errors.New("fatal execution error")

func (e *Executor) execute(b *blocks.Block, log logging.Logger) error {
	_ = "STUB: not implemented"
	// If the VM were to encounter an error after enqueuing the block, we would
	// receive the same block twice for execution should consensus retry
	// acceptance.
	return nil
}

type (
	// ReceiptStore receives per-transaction receipts during block execution.
	// Only the [Executor] needs to provide a real implementation to [Execute]
	// and all other callers MUST use [NullReceiptStore].
	ReceiptStore interface {
		Load(common.Hash) (eventual.Value[*Receipt], bool)
	}

	// ExecutionResults holds the outputs of [Execute].
	ExecutionResults struct {
		BaseFee  *uint256.Int
		StateDB  *state.StateDB
		Signer   types.Signer
		BlockCtx vm.BlockContext
		Receipts types.Receipts
		FinishBy struct {
			Gas  *gastime.Time
			Wall time.Time
		}
	}
)

// Execute executes the transactions in the [blocks.Block], beginning from the
// post-execution state of the [blocks.Block.ParentBlock]. `maxNumTxs` limits
// the number of transactions to process, allowing partial execution for
// intra-block inspection.
//
// Although Execute does not call [blocks.Block.MarkExecuted] it does mutate
// consensus-critical internal values (e.g. interim execution time). A "live"
// accepted block (as against one recovered from the database) MUST NOT be
// passed directly to [Execute], only to [Executor.Enqueue].
func Execute(
	b *blocks.Block,
	sdbo saedb.StateDBOpener,
	maxNumTxs int,
	hooks hook.Points,
	config *params.ChainConfig,
	chainCtx core.ChainContext,
	receiptStore ReceiptStore,
	log logging.Logger,
) (*ExecutionResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*isMerge*/

// required by geth but irrelevant so max it out

// TODO(arr4n) investigate calling the same method on pending blocks in
// the queue. It's only worth it if [blocks.LastToSettleAt] regularly
// returns false, meaning that execution is blocking consensus.

// The [types.Header] that we pass to [core.ApplyTransaction] is
// modified to reduce gas price from the worst-case value agreed by
// consensus. This changes the hash, which is what is copied to receipts
// and logs.
//
// [core.ApplyTransaction] also doesn't set [types.Receipt.EffectiveGasPrice].
// Fixing both here avoids needing to call [types.Receipt.DeriveFields].

func (e *Executor) afterExecution(b *blocks.Block, r *ExecutionResults) error {
	_ = "STUB: not implemented"
	return nil
}

// Responsibility for untracking lies with the VM once it deems this block's
// post-execution state to no longer be consensus-critical.

// The strict ordering of the next 3 calls guarantees invariants that MUST
// NOT be broken:
//
// 1. [blocks.Block.MarkExecuted] guarantees disk then in-memory changes.
// 2. Internal indicator of last executed MUST follow in-memory change.
// 3. External indicator of last executed MUST follow internal indicator.
/* (2) */

// (3)

// NullReceiptStore is a no-op [ReceiptStore] for use when receipt broadcasting
// is not needed (e.g. state tracing).
type NullReceiptStore struct{}

var _ ReceiptStore = (*NullReceiptStore)(nil)

// Load always returns the zero value and false.
func (*NullReceiptStore) Load(common.Hash) (eventual.Value[*Receipt], bool) {
	_ = "STUB: not implemented"
	return nil, false
}
