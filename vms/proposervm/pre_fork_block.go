// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

var (
	_ Block = (*preForkBlock)(nil)

	errChildOfPreForkBlockHasProposer = errors.New("child of pre-fork block has proposer")
)

type preForkBlock struct {
	snowman.Block
	vm *VM
}

func (b *preForkBlock) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*preForkBlock) acceptOuterBlk() error { _ = "STUB: not implemented"; return nil }

func (b *preForkBlock) acceptInnerBlk(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *preForkBlock) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *preForkBlock) Options(ctx context.Context) ([2]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A pre-fork block's child options are always pre-fork blocks

func (b *preForkBlock) getInnerBlk() snowman.Block {
	_ = "STUB: not implemented"
	return *new(snowman.Block)
}

func (b *preForkBlock) verifyPreForkChild(ctx context.Context, child *preForkBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// This method only returns nil once (during the transition)
func (b *preForkBlock) verifyPostForkChild(ctx context.Context, child *postForkBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure [b] is the parent of [child]'s inner block

// A *preForkBlock can only have a *postForkBlock child
// if the *preForkBlock is the last *preForkBlock before activation takes effect
// (its timestamp is at or after the activation time)

// Child's timestamp must be at or after its parent's timestamp

// Child timestamp can't be too far in the future

// The first block after the fork should not have an epoch even if granite is activated.

// Verify the lack of signature on the node

// Verify the inner block and track it as verified

func (*preForkBlock) verifyPostForkOption(context.Context, *postForkOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *preForkBlock) selectChildPChainHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *preForkBlock) buildChild(ctx context.Context) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

// The chain hasn't forked yet

// The chain is currently forking

// The child's P-Chain height is proposed as the optimal P-Chain height that
// is at least the minimum height

func (*preForkBlock) pChainHeight(context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (*preForkBlock) pChainEpoch(context.Context) (block.Epoch, error) {
	_ = "STUB: not implemented"
	return *new(block.Epoch), nil
}
