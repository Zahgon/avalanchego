// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ava-labs/avalanchego/graft/evm/sync/engine"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/eth"
)

var _ engine.ChainContext = (*chainContextAdapter)(nil)

// chainContextAdapter adapts *eth.Ethereum to implement engine.ChainContext.
// This is needed because eth.Ethereum.BlockChain() returns *core.BlockChain
// while the interface expects engine.BlockChain.
type chainContextAdapter struct {
	eth *eth.Ethereum
}

// newChainContextAdapter creates a new adapter for the given Ethereum instance.
func newChainContextAdapter(eth *eth.Ethereum) engine.ChainContext {
	_ = "STUB: not implemented"
	return *new(engine.ChainContext)
}

// BloomIndexer returns the bloom indexer, which implements engine.BloomIndexer.
func (a *chainContextAdapter) BloomIndexer() engine.BloomIndexer {
	_ = "STUB: not implemented"
	return *new(engine.BloomIndexer)
}

// BlockChain returns the blockchain, which implements engine.BlockChain.
func (a *chainContextAdapter) BlockChain() engine.BlockChain {
	_ = "STUB: not implemented"
	return *new(engine.BlockChain)
}
