// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/params"
	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/gastime"
	"github.com/ava-labs/avalanchego/vms/saevm/proxytime"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// SetInterimExecutionTime is expected to be called during execution of b's
// transactions, with the highest-known gas time. This MAY be at any resolution
// but MUST be monotonic.
func (b *Block) SetInterimExecutionTime(t *proxytime.Time[gas.Gas]) {
	_ = "STUB: not implemented"
	return
}

//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

//nolint:revive // struct-tag: canoto allows unexported fields
type executionResults struct {
	byGas         gastime.Time `canoto:"value,1"`
	baseFee       uint256.Int  `canoto:"fixed repeated uint,2"`
	receiptRoot   common.Hash  `canoto:"fixed bytes,3"`
	stateRootPost common.Hash  `canoto:"fixed bytes,4"`

	ephemeralExecutionResults // not for canotofication

	canotoData canotoData_executionResults
}

type ephemeralExecutionResults struct {
	// Wall-clock time is for metrics only and MAY be incorrect.
	byWall time.Time
	// Receipts are deliberately not stored by the canoto representation as they
	// are already in the database. All methods that read the stored canoto
	// either accept a [types.Receipts] for comparison against the
	// `receiptRoot`, or don't care about receipts at all. They are, however,
	// carried here for propagation to the settling block.
	receipts types.Receipts
}

func (e *executionResults) setBaseFee(bf *big.Int) error {
	_ = "STUB: not implemented"
	// genesis blocks
	return nil
}

// MarkExecuted marks the block as having been executed at the specified time(s)
// and with the specified results. It also sets the chain's head block to b. The
// [gastime.Time] MUST have already been scaled to the target applicable after
// the block, as defined by the relevant [hook.Points].
//
// MarkExecuted guarantees that state is persisted to the database before
// in-memory indicators of execution are updated. [Block.Executed] returning
// true and [Block.WaitUntilExecuted] returning cleanly are both therefore
// indicative of a successful database write by MarkExecuted. The atomic pointer
// to the last-executed block is updated before [Block.WaitUntilExecuted]
// returns.
//
// This method MUST NOT be called more than once. The wall-clock [time.Time] is
// for metrics only.
func (b *Block) MarkExecuted(
	db ethdb.Database,
	xdb saetypes.ExecutionResults,
	byGas *gastime.Time,
	byWall time.Time,
	baseFee *big.Int,
	receipts types.Receipts,
	stateRootPost common.Hash,
	lastExecuted *atomic.Pointer[Block],
) error {
	_ = "STUB: not implemented"
	return nil
}

// The final execution time is scaled to the new gas target but interim
// times are not, which can result in rounding errors. Scaling always
// rounds up, to maintain a monotonic clock, but we confirm for safety.
// The logger used in tests will also convert this to a failure.

var errMarkBlockExecutedAgain = errors.New("block re-marked as executed")

// markExecuted is the intersection point of [Block.MarkExecuted],
// [Block.MarkSynchronous], and [Block.RestoreExecutionArtefacts], all of which
// have side effects drawn from the same ordered set of events. This method
// exists to guarantee that the correct selection and ordering of events occurs,
// regardless of the upstream trigger. See documentation re ordering invariants
// for more information.
//
// The batch is `Write()`n (yeah, it's a word now) after all disk artefacts are
// persisted.
func (b *Block) markExecuted(batch ethdb.Batch, xdb saetypes.ExecutionResults, e *executionResults, setAsHeadBlock bool, lastExecuted *atomic.Pointer[Block]) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) markExecutedOnDisk(batch ethdb.Batch, xdb saetypes.ExecutionResults, e *executionResults, setAsHeadBlock bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) markExecutedAfterDiskArtefacts(e *executionResults, lastExecuted *atomic.Pointer[Block]) error {
	_ = "STUB: not implemented"
	// Memory
	return nil
}

// This is fatal because we corrupted the database's head block if we
// got here by [Block.MarkExecuted] being called twice (an invalid use
// of the API).

// Internal indicator

// External indicator

// SetAsHeadBlock calls all necessary [rawdb] write methods for setting the
// block as head. Said writes are not atomic and an [ethdb.Batch] SHOULD be used
// if this is a desired property.
func (b *Block) SetAsHeadBlock(kv ethdb.KeyValueWriter) { _ = "STUB: not implemented"; return }

// WaitUntilExecuted blocks until [Block.MarkExecuted] is called or the
// [context.Context] is cancelled.
func (b *Block) WaitUntilExecuted(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Executed reports whether [Block.MarkExecuted] has been called without
// resulting in an error.
func (b *Block) Executed() bool { _ = "STUB: not implemented"; return false }

// executionArtefact blocks until [Block.MarkExecuted] has been called and then
// returns the requested value. A warning is logged if the caller is blocked for
// longer than [saeparams.MaxQueueWallTime].
func executionArtefact[T any](b *Block, desc string, get func(*executionResults) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func (e *executionResults) executedByGasTime() *gastime.Time { _ = "STUB: not implemented"; return nil }
func (e *executionResults) executedByWallTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
func (e *executionResults) cloneBaseFee() *uint256.Int { _ = "STUB: not implemented"; return nil }
func (e *executionResults) cloneReceiptsSlice() types.Receipts {
	_ = "STUB: not implemented"
	return *new(types.Receipts)
}
func (e *executionResults) postExecutionStateRoot() common.Hash {
	_ = "STUB: not implemented"
	return *

	// ExecutedByGasTime blocks until [Block.MarkExecuted] has been called and
	// returns a clone of the gas time passed to it.
	new(common.Hash)
}

func (b *Block) ExecutedByGasTime() *gastime.Time { _ = "STUB: not implemented"; return nil }

// ExecutedByWallTime blocks until [Block.MarkExecuted] has been called and
// returns the wall time passed to it.
func (b *Block) ExecutedByWallTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// ExecutedBaseFee blocks until [Block.MarkExecuted] has been called and returns
// a clone of the base fee passed to it.
func (b *Block) ExecutedBaseFee() *uint256.Int { _ = "STUB: not implemented"; return nil }

// Receipts blocks until [Block.MarkExecuted] has been called and returns the
// receipts passed to it.
func (b *Block) Receipts() types.Receipts { _ = "STUB: not implemented"; return *new(types.Receipts) }

// PostExecutionStateRoot blocks until [Block.MarkExecuted] has been called and
// returns the state root passed to it.
func (b *Block) PostExecutionStateRoot() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// RestoreExecutionArtefacts reloads post-execution artefacts persisted by
// [Block.MarkExecuted] such that the block is in an equivalent state to when
// said function was originally called.
func (b *Block) RestoreExecutionArtefacts(db ethdb.Database, xdb saetypes.ExecutionResults, chainConfig *params.ChainConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// SAE does not support blob transactions.

func loadExecutionResults(xdb saetypes.ExecutionResults, blockNum uint64) (*executionResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func persistedExecutionArtefact[T any](xdb saetypes.ExecutionResults, blockNum uint64, get func(*executionResults) T) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// PostExecutionStateRoot mirrors the behaviour of
// [Block.RestoreExecutionArtefacts], without requiring a full [Block], and only
// returning the state root after execution.
func PostExecutionStateRoot(xdb saetypes.ExecutionResults, blockNum uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// ExecutionBaseFee mirrors the behaviour of [Block.RestoreExecutionArtefacts],
// without requiring a full [Block], and only returning the base fee when the
// block was executed (as against the worst-case prediction).
func ExecutionBaseFee(xdb saetypes.ExecutionResults, blockNum uint64) (*uint256.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
