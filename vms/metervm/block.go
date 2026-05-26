// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ snowman.Block           = (*meterBlock)(nil)
	_ snowman.OracleBlock     = (*meterBlock)(nil)
	_ block.WithVerifyContext = (*meterBlock)(nil)

	errExpectedBlockWithVerifyContext = errors.New("expected block.WithVerifyContext")
)

type meterBlock struct {
	snowman.Block

	vm *blockVM
}

func (mb *meterBlock) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mb *meterBlock) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mb *meterBlock) Reject(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mb *meterBlock) Options(ctx context.Context) ([2]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mb *meterBlock) ShouldVerifyWithContext(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (mb *meterBlock) VerifyWithContext(ctx context.Context, blockCtx *block.Context) error {
	_ = "STUB: not implemented"
	return nil
}
