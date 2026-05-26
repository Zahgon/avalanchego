// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var _ Block = (*ApricotAtomicBlock)(nil)

// ApricotAtomicBlock being accepted results in the atomic transaction contained
// in the block to be accepted and committed to the chain.
type ApricotAtomicBlock struct {
	CommonBlock `serialize:"true"`
	Tx          *txs.Tx `serialize:"true" json:"tx"`
}

func (b *ApricotAtomicBlock) initialize(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (b *ApricotAtomicBlock) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (b *ApricotAtomicBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *ApricotAtomicBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

func NewApricotAtomicBlock(
	parentID ids.ID,
	height uint64,
	tx *txs.Tx,
) (*ApricotAtomicBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
