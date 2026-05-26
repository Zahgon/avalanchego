// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocktest

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	errBuildBlock         = errors.New("unexpectedly called BuildBlock")
	errParseBlock         = errors.New("unexpectedly called ParseBlock")
	errGetBlock           = errors.New("unexpectedly called GetBlock")
	errLastAccepted       = errors.New("unexpectedly called LastAccepted")
	errGetBlockIDAtHeight = errors.New("unexpectedly called GetBlockIDAtHeight")

	_ block.ChainVM = (*VM)(nil)
)

// VM is a ChainVM that is useful for testing.
type VM struct {
	enginetest.VM

	CantBuildBlock,
	CantParseBlock,
	CantGetBlock,
	CantSetPreference,
	CantLastAccepted,
	CantGetBlockIDAtHeight bool

	BuildBlockF         func(context.Context) (snowman.Block, error)
	ParseBlockF         func(context.Context, []byte) (snowman.Block, error)
	GetBlockF           func(context.Context, ids.ID) (snowman.Block, error)
	SetPreferenceF      func(context.Context, ids.ID) error
	LastAcceptedF       func(context.Context) (ids.ID, error)
	GetBlockIDAtHeightF func(ctx context.Context, height uint64) (ids.ID, error)
}

func (vm *VM) Default(cant bool) { _ = "STUB: not implemented"; return }

func (vm *VM) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) ParseBlock(ctx context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) GetBlock(ctx context.Context, id ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) SetPreference(ctx context.Context, id ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) LastAccepted(ctx context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *VM) GetBlockIDAtHeight(ctx context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
