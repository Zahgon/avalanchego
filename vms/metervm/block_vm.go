// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ block.ChainVM                         = (*blockVM)(nil)
	_ block.BuildBlockWithContextChainVM    = (*blockVM)(nil)
	_ block.SetPreferenceWithContextChainVM = (*blockVM)(nil)
	_ block.BatchedChainVM                  = (*blockVM)(nil)
	_ block.StateSyncableVM                 = (*blockVM)(nil)
)

type blockVM struct {
	block.ChainVM
	buildBlockVM    block.BuildBlockWithContextChainVM
	setPreferenceVM block.SetPreferenceWithContextChainVM
	batchedVM       block.BatchedChainVM
	ssVM            block.StateSyncableVM

	blockMetrics
	registry prometheus.Registerer
}

func NewBlockVM(
	vm block.ChainVM,
	reg prometheus.Registerer,
) block.ChainVM {
	_ = "STUB: not implemented"
	return *new(block.ChainVM)
}

func (vm *blockVM) Initialize(
	ctx context.Context,
	chainCtx *snow.Context,
	db database.Database,
	genesisBytes,
	upgradeBytes,
	configBytes []byte,
	fxs []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *blockVM) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *blockVM) ParseBlock(ctx context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *blockVM) GetBlock(ctx context.Context, id ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *blockVM) SetPreference(ctx context.Context, id ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *blockVM) LastAccepted(ctx context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *blockVM) GetBlockIDAtHeight(ctx context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
