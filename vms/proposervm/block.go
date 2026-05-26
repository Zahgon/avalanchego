// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

const (
	// allowable block issuance in the future
	maxSkew                         = 10 * time.Second
	bootstrappingWarningGracePeriod = 5 * time.Minute
)

var (
	errUnsignedChild            = errors.New("expected child to be signed")
	errUnexpectedBlockType      = errors.New("unexpected proposer block type")
	errInnerParentMismatch      = errors.New("inner parentID didn't match expected parent")
	errTimeNotMonotonic         = errors.New("time must monotonically increase")
	errPChainHeightNotMonotonic = errors.New("non monotonically increasing P-chain height")
	errPChainHeightNotReached   = errors.New("block P-chain height larger than current P-chain height")
	errTimeTooAdvanced          = errors.New("time is too far advanced")
	errEpochMismatch            = errors.New("epoch mismatch")
	errProposerWindowNotStarted = errors.New("proposer window hasn't started")
	errUnexpectedProposer       = errors.New("unexpected proposer for current window")
	errProposerMismatch         = errors.New("proposer mismatch")
	errProposersNotActivated    = errors.New("proposers haven't been activated yet")
	errPChainHeightTooLow       = errors.New("block P-chain height is too low")
	errEpochNotZero             = errors.New("epoch must not be provided prior to granite")
)

type Block interface {
	snowman.Block

	getInnerBlk() snowman.Block

	// After a state sync, we may need to update last accepted block data
	// without propagating any changes to the innerVM.
	// acceptOuterBlk and acceptInnerBlk allow controlling acceptance of outer
	// and inner blocks.
	acceptOuterBlk() error
	acceptInnerBlk(context.Context) error

	verifyPreForkChild(ctx context.Context, child *preForkBlock) error
	verifyPostForkChild(ctx context.Context, child *postForkBlock) error
	verifyPostForkOption(ctx context.Context, child *postForkOption) error

	buildChild(context.Context) (Block, error)

	pChainHeight(context.Context) (uint64, error)
	pChainEpoch(context.Context) (block.Epoch, error)
	selectChildPChainHeight(context.Context) (uint64, error)
}

type PostForkBlock interface {
	Block

	getStatelessBlk() block.Block
	setInnerBlk(snowman.Block)
}

// field of postForkBlock and postForkOption
type postForkCommonComponents struct {
	vm       *VM
	innerBlk snowman.Block
}

// Return the inner block's height
func (p *postForkCommonComponents) Height() uint64 { _ = "STUB: not implemented"; return 0 }

// Verify returns nil if:
// 1) [p]'s inner block is not an oracle block
// 2) [child]'s P-Chain height >= [parentPChainHeight]
// 3) [p]'s inner block is the parent of [c]'s inner block
// 4) [child]'s timestamp isn't before [p]'s timestamp
// 5) [child]'s timestamp is within the skew bound
// 6) [childPChainHeight] <= the current P-Chain height
// 7) [child]'s timestamp is within its proposer's window
// 8) [child] has a valid signature from its proposer
// 9) [child]'s inner block is valid
// 10) [child] has the expected epoch
func (p *postForkCommonComponents) Verify(
	ctx context.Context,
	parentTimestamp time.Time,
	parentPChainHeight uint64,
	parentEpoch block.Epoch,
	child *postForkBlock,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Some L1s that missed the Granite upgrade accepted blocks without
// epochs enabled. By enforcing this check only after syncing, new nodes
// are able to join the network.

// If the node is currently syncing - we don't assume that the P-chain
// has been synced up to this point yet.

// Return the child (a *postForkBlock) of this block
func (p *postForkCommonComponents) buildChild(
	ctx context.Context,
	parentID ids.ID,
	parentTimestamp time.Time,
	parentPChainHeight uint64,
	parentEpoch block.Epoch,
) (Block, error) {
	_ = "STUB: not implemented"
	// Child's timestamp is the later of now and this block's timestamp
	return *new(Block), nil
}

// The child's P-Chain height is proposed as the optimal P-Chain height that
// is at least the parent's P-Chain height

// Build the child

func (p *postForkCommonComponents) getInnerBlk() snowman.Block {
	_ = "STUB: not implemented"
	return *new(snowman.Block)
}

func (p *postForkCommonComponents) setInnerBlk(innerBlk snowman.Block) {
	_ = "STUB: not implemented"
	return
}

func verifyIsOracleBlock(ctx context.Context, b snowman.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyIsNotOracleBlock(ctx context.Context, b snowman.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *postForkCommonComponents) verifyPreDurangoBlockDelay(
	ctx context.Context,
	parentTimestamp time.Time,
	parentPChainHeight uint64,
	blk *postForkBlock,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *postForkCommonComponents) verifyPostDurangoBlockDelay(
	ctx context.Context,
	parentTimestamp time.Time,
	parentPChainHeight uint64,
	blk *postForkBlock,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// populate the slot for the block.

// find the expected proposer

// block should be unsigned

// block should be signed

func (p *postForkCommonComponents) shouldBuildSignedBlockPostDurango(
	ctx context.Context,
	parentID ids.ID,
	parentTimestamp time.Time,
	parentPChainHeight uint64,
	newTimestamp time.Time,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// build an unsigned block

// build a signed block

// It's not our turn to propose a block yet. This is likely caused by having
// previously notified the consensus engine to attempt to build a block on
// top of a block that is no longer the preferred block.

func (p *postForkCommonComponents) shouldBuildSignedBlockPreDurango(
	ctx context.Context,
	parentID ids.ID,
	parentTimestamp time.Time,
	parentPChainHeight uint64,
	newTimestamp time.Time,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// time for any node to build an unsigned block

// it's time for this node to propose a block. It'll be signed or
// unsigned depending on the delay

// It's not our turn to propose a block yet. This is likely caused by having
// previously notified the consensus engine to attempt to build a block on
// top of a block that is no longer the preferred block.

func (p *postForkCommonComponents) logWarnOrError() func(msg string, fields ...zap.Field) {
	_ = "STUB: not implemented"
	return nil
}

// logUnexpectedPChainError logs an unexpected P-chain failure: Warn when err
// wraps database.ErrClosed, Error otherwise.
//
// ErrClosed is expected during VM shutdown because the P-chain can stop and
// close its DB before chains that consult it finish stopping. Logging at Warn
// avoids noise in that normal path. If ErrClosed happens during normal
// operation, the failure will still be surfaced elsewhere.
func logUnexpectedPChainError(log logging.Logger, err error, msg string, fields ...zap.Field) {
	_ = "STUB: not implemented"
	// Caller skip so the caller field points at the call site, not here.
	return
}
