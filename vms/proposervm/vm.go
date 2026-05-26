// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/utils/tree"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/proposervm/proposer"
	"github.com/ava-labs/avalanchego/vms/proposervm/state"
)

const (
	httpPathEndpoint = "/proposervm"
	HTTPHeaderRoute  = "proposervm"

	// DefaultMinBlockDelay should be kept as whole seconds because block
	// timestamps are only specific to the second.
	DefaultMinBlockDelay = time.Second
	// DefaultNumHistoricalBlocks as 0 results in never deleting any historical
	// blocks.
	DefaultNumHistoricalBlocks uint64 = 0

	innerBlkCacheSize = 64 * units.MiB
)

var (
	_ block.ChainVM         = (*VM)(nil)
	_ block.BatchedChainVM  = (*VM)(nil)
	_ block.StateSyncableVM = (*VM)(nil)

	dbPrefix = []byte("proposervm")
)

func cachedBlockSize(_ ids.ID, blk snowman.Block) int { _ = "STUB: not implemented"; return 0 }

type VM struct {
	block.ChainVM
	Config
	blockBuilderVM  block.BuildBlockWithContextChainVM
	setPreferenceVM block.SetPreferenceWithContextChainVM
	batchedVM       block.BatchedChainVM
	ssVM            block.StateSyncableVM

	state.State

	proposer.Windower
	tree.Tree
	mockable.Clock
	finishedBootstrappingAt time.Time

	ctx *snow.Context
	db  *versiondb.Database

	// Block ID --> Block
	// Each element is a block that passed verification but
	// hasn't yet been accepted/rejected
	verifiedBlocks map[ids.ID]PostForkBlock
	// Stateless block ID --> inner block.
	// Only contains post-fork blocks near the tip so that the cache doesn't get
	// filled with random blocks every time this node parses blocks while
	// processing a GetAncestors message from a bootstrapping node.
	innerBlkCache  cache.Cacher[ids.ID, snowman.Block]
	preferred      ids.ID
	consensusState snow.State

	// lastAcceptedTime is set to the last accepted PostForkBlock's timestamp
	// if the last accepted block has been a PostForkOption block since having
	// initialized the VM.
	lastAcceptedTime time.Time

	// lastAcceptedHeight is set to the last accepted PostForkBlock's height.
	lastAcceptedHeight uint64

	// proposerBuildSlotGauge reports the slot index when this node may attempt
	// to build a block.
	proposerBuildSlotGauge prometheus.Gauge

	// acceptedBlocksSlotHistogram reports the slots that accepted blocks were
	// proposed in.
	acceptedBlocksSlotHistogram prometheus.Histogram

	// lastAcceptedTimestampGaugeVec reports timestamps for the last-accepted
	// [postForkBlock] and its inner block.
	lastAcceptedTimestampGaugeVec *prometheus.GaugeVec
}

// New performs best when [minBlkDelay] is whole seconds. This is because block
// timestamps are only specific to the second.
func New(
	vm block.ChainVM,
	config Config,
) *VM {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Initialize(
	ctx context.Context,
	chainCtx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	upgradeBytes []byte,
	configBytes []byte,
	fxs []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// define the following ranges:
// (-inf, 0]
// (0, 1]
// (1, 2]
// (2, inf)
// the usage of ".5" before was to ensure we work around the limitation
// of comparing floating point of the same numerical value.

// Shutdown ops then propagate shutdown to innerVM
func (vm *VM) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (vm *VM) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) NewHTTPHandler(ctx context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (vm *VM) SetState(ctx context.Context, newState snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

// When finishing StateSyncing, if state sync has failed or was skipped,
// repairAcceptedChainByHeight rolls back the chain to the previously last
// accepted block. If state sync has completed successfully, this call is a
// no-op.

func (vm *VM) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) ParseBlock(ctx context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) ParseLocalBlock(ctx context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) GetBlock(ctx context.Context, id ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) SetPreference(ctx context.Context, preferred ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// If the inner VM implements SetPreferenceWithContext, use it to set the
// preference with the P-Chain height to be used to verify a child of the
// preferred block.

// The P-Chain height used to verify a child of the preferred block will
// potentially be different than the P-Chain height used to verify the
// preferred block if the preferred block seals the current epoch.

// The exact child timestamp doesn't matter here because we know Granite
// is already activated.

func (vm *VM) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

// If we are pre-fork or haven't finished bootstrapping yet, we should
// directly forward the inner VM's events.

// Wait until it is our turn to build a block.

// We should not call ChainVM.WaitForEvent here as it is possible
// that timeToBuild was capped less than the actual time for us to
// build a block. If it is actually our turn to build, timeToBuild
// will be <= 0 in the next iteration.

func (vm *VM) timeToBuild(ctx context.Context) (time.Time, bool, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), false, nil
}

// Block building is only supported if the consensus state is normal
// operations and the vm is not state syncing.
//
// TODO: Correctly handle dynamic state sync here. When the innerVM is
// dynamically state syncing, we should return here as well.

// Because the VM in marked as being in the [snow.NormalOp] state, we know
// that [VM.SetPreference] must have already been called.

// If the preferred block is pre-fork, we should wait for events on the
// innerVM.

// A nil error is returned here because it is possible that
// bootstrapping caused the last accepted block to move past the latest
// P-chain height. This will cause building blocks to return an error
// until the P-chain's height has advanced.

func (vm *VM) getPreDurangoSlotTime(
	ctx context.Context,
	blkHeight,
	pChainHeight uint64,
	parentTimestamp time.Time,
) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Note: The P-chain does not currently try to target any block time. It
// notifies the consensus engine as soon as a new block may be built. To
// avoid fast runs of blocks there is an additional minimum delay that
// validators can specify. This delay may be an issue for high performance,
// custom VMs. Until the P-chain is modified to target a specific block
// time, ProposerMinBlockDelay can be configured in the node config.

func (vm *VM) getPostDurangoSlotTime(
	ctx context.Context,
	blkHeight,
	pChainHeight,
	slot uint64,
	parentTimestamp time.Time,
) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Note: The P-chain does not currently try to target any block time. It
// notifies the consensus engine as soon as a new block may be built. To
// avoid fast runs of blocks there is an additional minimum delay that
// validators can specify. This delay may be an issue for high performance,
// custom VMs. Until the P-chain is modified to target a specific block
// time, ProposerMinBlockDelay can be configured in the node config.

func (vm *VM) LastAccepted(ctx context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *VM) repairAcceptedChainByHeight(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If the last accepted block isn't indexed yet, then the underlying
// chain is the only chain and there is nothing to repair.

// There is nothing to repair - as the heights match

// The inner vm must be behind the proposer vm, so we must roll the
// proposervm back.

// We are rolling back past the fork, so we should just forget about all
// of our proposervm indices.

// This fatal error can happen if NumHistoricalBlocks is set too
// aggressively and the inner vm rolled back before the oldest
// proposervm block.

func (vm *VM) setLastAcceptedMetadata(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If the last accepted block wasn't a PostFork block, then we don't
// initialize the metadata.

// Set the last accepted height

// If the last accepted block wasn't a PostForkOption, then we don't
// initialize the time.

func (vm *VM) parsePostForkBlock(ctx context.Context, b []byte, verifySignature bool) (PostForkBlock, error) {
	_ = "STUB: not implemented"
	return *new(PostForkBlock), nil
}

func (vm *VM) parsePreForkBlock(ctx context.Context, b []byte) (*preForkBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) getBlock(ctx context.Context, id ids.ID) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func (vm *VM) getPostForkBlock(ctx context.Context, blkID ids.ID) (PostForkBlock, error) {
	_ = "STUB: not implemented"
	return *new(PostForkBlock), nil
}

func (vm *VM) getPreForkBlock(ctx context.Context, blkID ids.ID) (*preForkBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) acceptPostForkBlock(blk PostForkBlock) error { _ = "STUB: not implemented"; return nil }

// Persist this block, its height index, and its status

func (vm *VM) verifyAndRecordInnerBlk(ctx context.Context, blockCtx *block.Context, postFork PostForkBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// We must update all of the mappings from postFork -> innerBlock to
// now point to originalInnerBlock.

// Invariant: If either [Verify] or [VerifyWithContext] returns nil, this
//            function must return nil. This maintains the inner block's
//            invariant that successful verification will eventually result
//            in accepted or rejected being called.

// This block needs to know the P-Chain height during verification.
// Note that [VerifyWithContext] with context may be called multiple
// times with multiple contexts.

// This isn't a [block.WithVerifyContext] so we only call [Verify] once.

// Since verification passed, we should ensure the inner block tree is
// populated.

// fujiOverridePChainHeightUntilHeight is the P-chain height at which the
// proposervm will no longer attempt to keep the P-chain height the same.
const fujiOverridePChainHeightUntilHeight = 200041

// fujiOverridePChainHeightUntilTimestamp is the timestamp at which the
// proposervm will no longer attempt to keep the P-chain height the same.
var fujiOverridePChainHeightUntilTimestamp = time.Date(2025, time.March, 7, 17, 0, 0, 0, time.UTC) // noon ET

func (vm *VM) selectChildPChainHeight(ctx context.Context, minPChainHeight uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// parseInnerBlock attempts to parse the provided bytes as an inner block. If
// the inner block happens to be cached, then the inner block will not be
// parsed.
func (vm *VM) parseInnerBlock(ctx context.Context, outerBlkID ids.ID, innerBlkBytes []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// Caches proposervm block ID --> inner block if the inner block's height
// is within [innerBlkCacheSize] of the last accepted block's height.
func (vm *VM) cacheInnerBlock(outerBlkID ids.ID, innerBlk snowman.Block) {
	_ = "STUB: not implemented"
	return
}
