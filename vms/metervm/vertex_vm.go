// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/snow/engine/common"
)

var (
	_ vertex.LinearizableVMWithEngine = (*vertexVM)(nil)
	_ snowstorm.Tx                    = (*meterTx)(nil)
)

func NewVertexVM(
	vm vertex.LinearizableVMWithEngine,
	reg prometheus.Registerer,
) vertex.LinearizableVMWithEngine {
	_ = "STUB: not implemented"
	return *new(vertex.LinearizableVMWithEngine)
}

type vertexVM struct {
	vertex.LinearizableVMWithEngine
	vertexMetrics
	registry prometheus.Registerer
}

func (vm *vertexVM) Initialize(
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

func (vm *vertexVM) ParseTx(ctx context.Context, b []byte) (snowstorm.Tx, error) {
	_ = "STUB: not implemented"
	return *new(snowstorm.Tx), nil
}

type meterTx struct {
	snowstorm.Tx

	vm *vertexVM
}

func (mtx *meterTx) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mtx *meterTx) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mtx *meterTx) Reject(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
