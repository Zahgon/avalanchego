// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

var _ PostForkBlock = (*postForkOption)(nil)

// The parent of a *postForkOption must be a *postForkBlock.
type postForkOption struct {
	block.Block
	postForkCommonComponents

	timestamp time.Time
}

func (b *postForkOption) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *postForkOption) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *postForkOption) acceptOuterBlk() error { _ = "STUB: not implemented"; return nil }

func (b *postForkOption) acceptInnerBlk(ctx context.Context) error {
	_ = "STUB: not implemented"
	// mark the inner block as accepted and all conflicting inner blocks as
	// rejected
	return nil
}

func (b *postForkOption) Reject(context.Context) error {
	_ = "STUB: not implemented"
	// we do not reject the inner block here because that block may be contained
	// in the proposer block that causing this block to be rejected.
	return nil
}

func (b *postForkOption) Parent() ids.ID {
	_ = "STUB: not implemented"
	return *

	// If Verify returns nil, Accept or Reject is eventually called on [b] and
	// [b.innerBlk].
	new(ids.ID)
}

func (b *postForkOption) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*postForkOption) verifyPreForkChild(context.Context, *preForkBlock) error {
	_ = "STUB: not implemented"
	// A *preForkBlock's parent must be a *preForkBlock
	return nil
}

func (b *postForkOption) verifyPostForkChild(ctx context.Context, child *postForkBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (*postForkOption) verifyPostForkOption(context.Context, *postForkOption) error {
	_ = "STUB: not implemented"
	// A *postForkOption's parent can't be a *postForkOption
	return nil
}

func (b *postForkOption) buildChild(ctx context.Context) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

// This block's P-Chain height is its parent's P-Chain height
func (b *postForkOption) pChainHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *postForkOption) pChainEpoch(ctx context.Context) (block.Epoch, error) {
	_ = "STUB: not implemented"
	return *new(block.Epoch), nil
}

func (b *postForkOption) selectChildPChainHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *postForkOption) getStatelessBlk() block.Block {
	_ = "STUB: not implemented"
	return *new(block.Block)
}
