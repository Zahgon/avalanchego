// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package xsvm

import (
	"context"
	"net/http"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/builder"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/chain"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/genesis"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ smblock.ChainVM                      = (*VM)(nil)
	_ smblock.BuildBlockWithContextChainVM = (*VM)(nil)
)

type VM struct {
	*p2p.Network

	chainContext *snow.Context
	db           database.Database
	genesis      *genesis.Genesis

	chain   chain.Chain
	builder builder.Builder
}

func (vm *VM) Initialize(
	_ context.Context,
	chainContext *snow.Context,
	db database.Database,
	genesisBytes []byte,
	_ []byte,
	_ []byte,
	_ []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Allow signing of all warp messages. This is not typically safe, but is
// allowed for this example.

func (vm *VM) SetState(_ context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*VM) Version(context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (vm *VM) CreateHandlers(context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) NewHTTPHandler(context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (*VM) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) GetBlock(_ context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) ParseBlock(_ context.Context, blkBytes []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

func (vm *VM) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) SetPreference(_ context.Context, preferred ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) LastAccepted(context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *VM) BuildBlockWithContext(ctx context.Context, blockContext *smblock.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) GetBlockIDAtHeight(_ context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
