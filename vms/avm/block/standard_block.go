// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"time"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

var _ Block = (*StandardBlock)(nil)

type StandardBlock struct {
	// parent's ID
	PrntID ids.ID `serialize:"true" json:"parentID"`
	// This block's height. The genesis block is at height 0.
	Hght uint64 `serialize:"true" json:"height"`
	Time uint64 `serialize:"true" json:"time"`
	Root ids.ID `serialize:"true" json:"merkleRoot"`
	// List of transactions contained in this block.
	Transactions []*txs.Tx `serialize:"true" json:"txs"`

	BlockID ids.ID `json:"id"`
	bytes   []byte
}

func (b *StandardBlock) initialize(bytes []byte, cm codec.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *StandardBlock) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (b *StandardBlock) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *StandardBlock) Parent() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *StandardBlock) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *StandardBlock) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *StandardBlock) MerkleRoot() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *StandardBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *StandardBlock) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func NewStandardBlock(
	parentID ids.ID,
	height uint64,
	timestamp time.Time,
	txs []*txs.Tx,
	cm codec.Manager,
) (*StandardBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We serialize this block as a pointer so that it can be deserialized into
// a Block
