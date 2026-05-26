// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/mempool"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	blockexecutor "github.com/ava-labs/avalanchego/vms/platformvm/block/executor"
	txexecutor "github.com/ava-labs/avalanchego/vms/platformvm/txs/executor"
)

const (
	// targetBlockSize is maximum number of transaction bytes to place into a
	// StandardBlock
	targetBlockSize = 128 * units.KiB

	// maxTimeToSleep is the maximum time to sleep between checking if a block
	// should be produced.
	maxTimeToSleep = time.Hour
)

var (
	_ Builder = (*builder)(nil)

	ErrEndOfTime                 = errors.New("program time is suspiciously far in the future")
	ErrNoPendingBlocks           = errors.New("no pending blocks")
	errMissingPreferredState     = errors.New("missing preferred block state")
	errCalculatingNextStakerTime = errors.New("failed calculating next staker time")
)

type Builder interface {
	smblock.BuildBlockWithContextChainVM
	// Add adds `tx` to the mempool and clears its dropped status.
	Add(tx *txs.Tx) error
	// Get returns the tx corresponding to `txID` and if it was present
	Get(txID ids.ID) (*txs.Tx, bool)
	// GetDropReason returns why `txID` was dropped
	GetDropReason(txID ids.ID) error
	// WaitForEvent blocks until the mempool has txs that are ready to build into
	// a block.
	WaitForEvent(ctx context.Context) (common.Message, error)

	// BuildBlock can be called to attempt to create a new block
	BuildBlock(context.Context) (snowman.Block, error)

	// PackAllBlockTxs returns an array of all txs that could be packed into a
	// valid block of infinite size. The returned txs are all verified against
	// the preferred state.
	//
	// Note: This function does not call the consensus engine.
	PackAllBlockTxs() ([]*txs.Tx, error)
}

// builder implements a simple builder to convert txs into valid blocks
type builder struct {
	*mempool.Mempool

	txExecutorBackend *txexecutor.Backend
	blkManager        blockexecutor.Manager
}

func New(
	mempool *mempool.Mempool,
	txExecutorBackend *txexecutor.Backend,
	blkManager blockexecutor.Manager,
) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builder) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	// We shouldn't call durationToSleep until the chain is marked as no longer
	// bootstrapping.
	return *new(common.Message), nil
}

// The next staker change is ready to be performed.

// Wait for a transaction in the mempool until there is a next staker
// change ready to be performed.

// Recheck the staker change time before returning

// Error could have been due to the parent context being cancelled
// or another unexpected error.

func (b *builder) durationToSleep() (time.Duration, error) {
	_ = "STUB: not implemented"
	// Grabbing the lock here enforces that this function is not called mid-way
	// through modifying of the state.
	return *new(time.Duration), nil
}

func (b *builder) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (b *builder) BuildBlockWithContext(
	ctx context.Context,
	blockContext *smblock.Context,
) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// Get the block to build on top of and retrieve the new block's context.

func (b *builder) PackAllBlockTxs() ([]*txs.Tx, error) { _ = "STUB: not implemented"; return nil, nil }

// [timestamp] is min(max(now, parent timestamp), next staker change time)
func buildBlock(
	ctx context.Context,
	builder *builder,
	parentID ids.ID,
	height uint64,
	timestamp time.Time,
	forceAdvanceTime bool,
	parentState state.Chain,
	pChainHeight uint64,
) (block.Block, error) {
	_ = "STUB: not implemented"
	return *new(block.Block), nil
}

// minCapacity is 0 as we want to honor the capacity in state.

// Try rewarding stakers whose staking period ends at the new chain time.
// This is done first to prioritize advancing the timestamp as quickly as
// possible.

// If there is no reason to build a block, don't.

// Issue a block with as many transactions as possible.

func packDurangoBlockTxs(
	ctx context.Context,
	parentID ids.ID,
	parentState state.Chain,
	mempool *mempool.Mempool,
	backend *txexecutor.Backend,
	manager blockexecutor.Manager,
	timestamp time.Time,
	pChainHeight uint64,
	remainingSize int,
) ([]*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func packEtnaBlockTxs(
	ctx context.Context,
	parentID ids.ID,
	parentState state.Chain,
	mempool *mempool.Mempool,
	backend *txexecutor.Backend,
	manager blockexecutor.Manager,
	timestamp time.Time,
	pChainHeight uint64,
	minCapacity gas.Gas,
) ([]*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func executeTx(
	ctx context.Context,
	parentID ids.ID,
	stateDiff *state.Diff,
	mempool *mempool.Mempool,
	backend *txexecutor.Backend,
	manager blockexecutor.Manager,
	pChainHeight uint64,
	inputs *set.Set[ids.ID],
	feeCalculator fee.Calculator,
	tx *txs.Tx,
) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Invariant: [tx] has already been syntactically verified.
		nil
}

// This log is a warn because the mempool should not have allowed this
// transaction to be included.

// getNextStakerToReward returns the next staker txID to remove from the staking
// set with a RewardValidatorTx rather than an AdvanceTimeTx. [chainTimestamp]
// is the timestamp of the chain at the time this validator would be getting
// removed and is used to calculate [shouldReward].
// Returns:
// - [txID] of the next staker to reward
// - [shouldReward] if the txID exists and is ready to be rewarded
// - [err] if something bad happened
func getNextStakerToReward(
	chainTimestamp time.Time,
	preferredState state.Chain,
) (ids.ID, bool, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), false, nil
}

// If the staker is a permissionless staker (not a permissioned subnet
// validator), it's the next staker we will want to remove with a
// RewardValidatorTx rather than an AdvanceTimeTx.

func NewRewardValidatorTx(ctx *snow.Context, txID ids.ID) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
