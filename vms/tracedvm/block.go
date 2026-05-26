// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracedvm

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ snowman.Block           = (*tracedBlock)(nil)
	_ snowman.OracleBlock     = (*tracedBlock)(nil)
	_ block.WithVerifyContext = (*tracedBlock)(nil)

	errExpectedBlockWithVerifyContext = errors.New("expected block.WithVerifyContext")
)

type tracedBlock struct {
	snowman.Block

	vm *blockVM
}

func (b *tracedBlock) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *tracedBlock) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *tracedBlock) Reject(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *tracedBlock) Options(ctx context.Context) ([2]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *tracedBlock) ShouldVerifyWithContext(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *tracedBlock) VerifyWithContext(ctx context.Context, blockCtx *block.Context) error {
	_ = "STUB: not implemented"
	return nil
}
