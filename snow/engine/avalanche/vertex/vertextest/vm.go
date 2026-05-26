// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertextest

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block/blocktest"
)

var (
	errLinearize = errors.New("unexpectedly called Linearize")

	_ vertex.LinearizableVM = (*VM)(nil)
)

type VM struct {
	blocktest.VM

	CantLinearize, CantParse bool

	LinearizeF func(context.Context, ids.ID) error
	ParseTxF   func(context.Context, []byte) (snowstorm.Tx, error)
}

func (vm *VM) Default(cant bool) { _ = "STUB: not implemented"; return }

func (vm *VM) Linearize(ctx context.Context, stopVertexID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) ParseTx(ctx context.Context, b []byte) (snowstorm.Tx, error) {
	_ = "STUB: not implemented"
	return *new(snowstorm.Tx), nil
}
