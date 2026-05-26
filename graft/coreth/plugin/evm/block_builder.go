// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"sync"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/core/txpool"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"

	commonEng "github.com/ava-labs/avalanchego/snow/engine/common"
)

const (
	// Minimum amount of time to wait to retry building a block after a failed attempt.
	// It is assumed that the first block building attempt is already after the minimum delay time.
	RetryDelay = 100 * time.Millisecond
)

type blockBuilder struct {
	clock *mockable.Clock
	ctx   *snow.Context

	txPool       *txpool.TxPool
	extraMempool extension.BuilderMempool

	shutdownChan <-chan struct{}
	shutdownWg   *sync.WaitGroup

	pendingSignal *lock.Cond

	buildBlockLock sync.Mutex
	// lastBuildParentHash is the parent hash of the last block that was built.
	// This and lastBuildTime are used to ensure that we don't build blocks too frequently,
	// but at least after a minimum delay of minBlockBuildingRetryDelay.
	lastBuildParentHash common.Hash
	lastBuildTime       time.Time

	chainHeadHash   common.Hash
	mempoolHeadHash common.Hash
}

// NewBlockBuilder creates a new block builder. extraMempool is an optional mempool (can be nil) that
// can be used to add transactions to the block builder, in addition to the txPool.
func (vm *VM) NewBlockBuilder(extraMempool extension.BuilderMempool) *blockBuilder {
	_ = "STUB: not implemented"
	return nil
}

// handleGenerateBlock is called from the VM immediately after BuildBlock.
func (b *blockBuilder) handleGenerateBlock(currentParentHash common.Hash) {
	_ = "STUB: not implemented"
	return
}

// needToBuild returns true if there are outstanding transactions to be issued
// into a block.
func (b *blockBuilder) needToBuild() bool { _ = "STUB: not implemented"; return false }

// signalCanBuild notifies a block is expected to be built.
func (b *blockBuilder) signalCanBuild() { _ = "STUB: not implemented"; return }

// awaitSubmittedTxs waits for new transactions to be submitted
// and notifies the VM when the tx pool has transactions to be
// put into a new block.
func (b *blockBuilder) awaitSubmittedTxs() {
	_ = "STUB: not implemented"
	// txSubmitChan is invoked when new transactions are issued as well as on re-orgs which
	// may orphan transactions that were previously in a preferred block.
	return
}

// waitForEvent waits until a block needs to be built.
// It returns only after at least [minBlockBuildingRetryDelay] passed from the last time a block was built.
func (b *blockBuilder) waitForEvent(ctx context.Context, currentHeader *types.Header) (commonEng.Message, error) {
	_ = "STUB: not implemented"
	return *new(commonEng.Message), nil
}

// waitForNeedToBuild waits until needToBuild returns true.
// It returns the last time a block was built.
func (b *blockBuilder) waitForNeedToBuild(ctx context.Context) (time.Time, common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(common.Hash), nil
}

// calculateBlockBuildingDelay calculates the delay needed before building the next block.
func (b *blockBuilder) calculateBlockBuildingDelay(
	lastBuildTime time.Time,
	lastBuildParentHash common.Hash,
	currentHeader *types.Header,
) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// if last build time is zero, this is not a retry
// If this is a retry, we already have waited for the minimum next block time in a previous attempt,
// and only need to wait for the retry delay.

// If this is not a retry, we need to wait for the minimum next block time.

// minNextBlockTime calculates the minimum next block time based on the current header.
func minNextBlockTime(parent *types.Header) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// If the parent header has no min delay excess, there is nothing to wait for, because the rule does not apply
// to the block to be built.

// parent's delay excess is already verified by consensus
// so this should not overflow

func (b *blockBuilder) setChainHeadHash(hash common.Hash) { _ = "STUB: not implemented"; return }

func (b *blockBuilder) setMempoolHeadHash(hash common.Hash) { _ = "STUB: not implemented"; return }

func (b *blockBuilder) pendingPoolUpdate() bool { _ = "STUB: not implemented"; return false }
