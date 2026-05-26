// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracedvm

import (
	"context"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/trace"
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
	// ChainVM tags
	initializeTag              string
	buildBlockTag              string
	parseBlockTag              string
	getBlockTag                string
	setPreferenceTag           string
	lastAcceptedTag            string
	verifyTag                  string
	acceptTag                  string
	rejectTag                  string
	optionsTag                 string
	shouldVerifyWithContextTag string
	verifyWithContextTag       string
	// BuildBlockWithContextChainVM tags
	buildBlockWithContextTag string
	// SetPreferenceWithContextChainVM tags
	setPreferenceWithContextTag string
	// BatchedChainVM tags
	getAncestorsTag      string
	batchedParseBlockTag string
	// HeightIndexedChainVM tags
	getBlockIDAtHeightTag string
	// StateSyncableVM tags
	stateSyncEnabledTag           string
	getOngoingSyncStateSummaryTag string
	getLastStateSummaryTag        string
	parseStateSummaryTag          string
	getStateSummaryTag            string
	tracer                        trace.Tracer
}

func NewBlockVM(vm block.ChainVM, name string, tracer trace.Tracer) block.ChainVM {
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

func (vm *blockVM) ParseBlock(ctx context.Context, block []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *blockVM) GetBlock(ctx context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *blockVM) SetPreference(ctx context.Context, blkID ids.ID) error {
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
