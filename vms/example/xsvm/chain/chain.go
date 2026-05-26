// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"

	xsblock "github.com/ava-labs/avalanchego/vms/example/xsvm/block"
)

var _ Chain = (*chain)(nil)

type Chain interface {
	LastAccepted() ids.ID
	SetChainState(state snow.State)
	GetBlock(blkID ids.ID) (Block, error)

	// Creates a fully verifiable and executable block, which can be processed
	// by the consensus engine, from a stateless block.
	NewBlock(blk *xsblock.Stateless) (Block, error)
}

type chain struct {
	chainContext  *snow.Context
	acceptedState database.Database

	// chain state as driven by the consensus engine
	chainState snow.State

	lastAcceptedID ids.ID
	verifiedBlocks map[ids.ID]*block
}

func New(ctx *snow.Context, db database.Database) (Chain, error) {
	_ = "STUB: not implemented"
	// Load the last accepted block data. For a newly created VM, this will be
	// the genesis. It is assumed the genesis was processed and stored
	// previously during VM initialization.
	return *new(Chain), nil
}

func (c *chain) LastAccepted() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (c *chain) SetChainState(state snow.State) { _ = "STUB: not implemented"; return }

func (c *chain) GetBlock(blkID ids.ID) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func (c *chain) NewBlock(blk *xsblock.Stateless) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func (c *chain) getBlock(blkID ids.ID) (*block, error) { _ = "STUB: not implemented"; return nil, nil }
