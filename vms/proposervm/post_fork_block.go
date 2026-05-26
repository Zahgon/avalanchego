// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

var _ PostForkBlock = (*postForkBlock)(nil)

type postForkBlock struct {
	block.SignedBlock
	postForkCommonComponents

	// slot of the proposer that produced this block.
	// It is populated in verifyPostDurangoBlockDelay.
	// It is used to report metrics during Accept.
	slot *uint64
}

// Accept:
// 1) Sets this blocks status to Accepted.
// 2) Persists this block in storage
// 3) Calls Reject() on siblings of this block and their descendants.
func (b *postForkBlock) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

const (
	innerBlockTypeMetricLabel = "inner"
	outerBlockTypeMetricLabel = "proposervm"
)

func (b *postForkBlock) updateLastAcceptedTimestampMetric(blockTypeLabel string, t time.Time) {
	_ = "STUB: not implemented"
	return
}

func (b *postForkBlock) acceptOuterBlk() error {
	_ = "STUB: not implemented"
	// Update in-memory references
	return nil
}

func (b *postForkBlock) acceptInnerBlk(ctx context.Context) error {
	_ = "STUB: not implemented"
	// mark the inner block as accepted and all conflicting inner blocks as
	// rejected
	return nil
}

func (b *postForkBlock) Reject(context.Context) error {
	_ = "STUB: not implemented"
	// We do not reject the inner block here because it may be accepted later
	return nil
}

// Return this block's parent, or a *missing.Block if
// we don't have the parent.
func (b *postForkBlock) Parent() ids.ID {
	_ = "STUB: not implemented"
	return *

	// If Verify() returns nil, Accept() or Reject() will eventually be called on
	// [b] and [b.innerBlk]
	new(ids.ID)
}

func (b *postForkBlock) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Return the two options for the block that follows [b]
func (b *postForkBlock) Options(ctx context.Context) ([2]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [b]'s innerBlk isn't an oracle block

// The inner block's child options

// Wrap the inner block's child option

// A post-fork block can never have a pre-fork child
func (*postForkBlock) verifyPreForkChild(context.Context, *preForkBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *postForkBlock) verifyPostForkChild(ctx context.Context, child *postForkBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *postForkBlock) verifyPostForkOption(ctx context.Context, child *postForkOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure [b]'s inner block is the parent of [child]'s inner block

// Return the child (a *postForkBlock) of this block
func (b *postForkBlock) buildChild(ctx context.Context) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func (b *postForkBlock) pChainHeight(context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *postForkBlock) pChainEpoch(context.Context) (block.Epoch, error) {
	_ = "STUB: not implemented"
	return *new(block.Epoch), nil
}

func (b *postForkBlock) selectChildPChainHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *postForkBlock) getStatelessBlk() block.Block {
	_ = "STUB: not implemented"
	return *new(block.Block)
}
