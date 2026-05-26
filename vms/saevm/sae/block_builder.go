// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/libevm/core/txpool"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/saexec"
	"github.com/ava-labs/avalanchego/vms/saevm/txgossip"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// blockBuilder hides [blockBuilderG]'s generic type behind non-generic methods.
type blockBuilder interface {
	// new constructs a [blocks.Block] with the provided arguments. It is
	// allowed for parent and lastSettled to be nil as an indication that the
	// block hasn't yet been verified.
	new(eth *types.Block, parent, lastSettled *blocks.Block) (*blocks.Block, error)
	// build a new block on top of the provided parent. The block context MAY be
	// nil.
	build(ctx context.Context, bCtx *block.Context, parent *blocks.Block) (*blocks.Block, error)
	// rebuild attempts to build a block identical to the provided block. If the
	// provided block contains any invalid components, those components will be
	// set to their valid counterparts in the returned block. The block context
	// MAY be nil.
	rebuild(ctx context.Context, bCtx *block.Context, parent, block *blocks.Block) (*blocks.Block, error)
}

type blockBuilderG[T hook.Transaction] struct {
	hooks   hook.PointsG[T]
	now     func() time.Time
	log     logging.Logger
	exec    *saexec.Executor
	mempool *txgossip.Set
	source  saetypes.BlockSource
}

func (b *blockBuilderG[_]) new(eth *types.Block, parent, lastSettled *blocks.Block) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *blockBuilderG[_]) build(
	ctx context.Context,
	bCtx *block.Context,
	parent *blocks.Block,
) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *blockBuilderG[_]) rebuild(
	ctx context.Context,
	bCtx *block.Context,
	parent *blocks.Block,
	block *blocks.Block,
) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	// Moving sender caching into [VM.ParseBlock] is not robust as there is
	// insufficient spam protection, so it must be done here.
	return nil, nil
}

// asynchronous

var (
	errBlockTimeUnderMinimum = errors.New("block time under minimum allowed time")
	errBlockTimeBeforeParent = errors.New("block time before parent time")
	errBlockTimeAfterMaximum = errors.New("block time after maximum allowed time")
	errExecutionLagging      = errors.New("execution lagging for settlement")
)

// buildWithTxs implements the block-building logic shared by [blockBuilder.build]
// and [blockBuilder.rebuild]. The block context MAY be nil.
func (b *blockBuilderG[T]) buildWithTxs(
	ctx context.Context,
	bCtx *block.Context,
	parent *blocks.Block,
	pendingTxs func(txpool.PendingFilter) []*txgossip.LazyTransaction,
	builder hook.BlockBuilder[T],
) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A full queue is a normal mode of operation (backpressure working as
// intended) so should not be a warning.

// If we don't have enough gas remaining in the block for the minimum
// gas amount, we are done including transactions.

// The [saexec.Executor] checks the worst-case balance before tx
// execution so we MUST record it at the equivalent point, before
// ApplyTx().

// TODO(StephenButtolph): Return additional information from
// [hook.PointsG.PotentialEndOfBlockOps] to terminate the loop early
// when there is insufficient block space remaining.

func lastToSettle(
	hooks hook.Points,
	hdr *types.Header,
	parent *blocks.Block,
	now time.Time,
	log logging.Logger,
) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// It is allowed for [hook.BlockBuilder] to further constrain the allowed
// block times. However, every block MUST at least satisfy these basic
// sanity checks.

// Underflow of Add(-tau) is prevented by the above check.
