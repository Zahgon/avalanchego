// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package simulated

import (
	"time"

	"github.com/ava-labs/avalanchego/graft/coreth/eth"
	"github.com/ava-labs/avalanchego/graft/coreth/eth/ethconfig"
	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
	"github.com/ava-labs/avalanchego/graft/coreth/interfaces"
	"github.com/ava-labs/avalanchego/graft/coreth/node"
	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

var _ eth.PushGossiper = (*fakePushGossiper)(nil)

type fakePushGossiper struct{}

func (*fakePushGossiper) Add(*types.Transaction) {
	_ = "STUB: not implemented"

	// Client exposes the methods provided by the Ethereum RPC client.
	return
}

type Client interface {
	ethereum.BlockNumberReader
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.GasPricer1559
	ethereum.FeeHistoryReader
	ethereum.LogFilterer
	interfaces.AcceptedStateReader
	interfaces.AcceptedContractCaller
	ethereum.TransactionReader
	ethereum.TransactionSender
	ethereum.ChainIDReader
}

// simClient wraps ethclient. This exists to prevent extracting ethclient.Client
// from the Client interface returned by Backend.
type simClient struct {
	*ethclient.Client
}

// Backend is a simulated blockchain. You can use it to test your contracts or
// other code that interacts with the Ethereum chain.
type Backend struct {
	eth    *eth.Ethereum
	client simClient
	clock  *mockable.Clock
	server *rpc.Server
}

// NewBackend creates a new simulated blockchain that can be used as a backend for
// contract bindings in unit tests.
//
// A simulated backend always uses chainID 1337.
func NewBackend(alloc types.GenesisAlloc, options ...func(nodeConf *node.Config, ethConf *ethconfig.Config)) *Backend {
	_ = "STUB: not implemented"
	return nil
}

// Create the default configurations for the outer node shell and the Ethereum
// service to mutate with the options afterwards

// Assemble the Ethereum stack to run the chain with

// this should never happen

// this should never happen

// newWithNode sets up a simulated backend on an existing node. The provided node
// must not be started and will be started by this method.
func newWithNode(stack *node.Node, conf *eth.Config, blockPeriod uint64) (*Backend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close shuts down the simBackend.
// The simulated backend can't be used afterwards.
func (n *Backend) Close() error { _ = "STUB: not implemented"; return nil }

// Commit seals a block and moves the chain forward to a new empty block.
func (n *Backend) Commit(accept bool) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (n *Backend) buildBlock(accept bool, gap uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (n *Backend) acceptAncestors(block *types.Block) error { _ = "STUB: not implemented"; return nil }

// Accept all ancestors of the block

// Rollback removes all pending transactions, reverting to the last committed state.
func (n *Backend) Rollback() {
	_ = "STUB: not implemented"
	// Flush all transactions from the transaction pools
	return
}

// Fork creates a side-chain that can be used to simulate reorgs.
//
// This function should be called with the ancestor block where the new side
// chain should be started. Transactions (old and new) can then be applied on
// top and Commit-ed.
//
// Note, the side-chain will only become canonical (and trigger the events) when
// it becomes longer. Until then CallContract will still operate on the current
// canonical chain.
//
// There is a % chance that the side chain becomes canonical at the same length
// to simulate live network behavior.
func (n *Backend) Fork(parentHash common.Hash) error { _ = "STUB: not implemented"; return nil }

// Wait for tx pool to reorg, then flush the tx pool

// AdjustTime changes the block timestamp and creates a new block.
// It can only be called on empty blocks.
func (n *Backend) AdjustTime(adjustment time.Duration) error { _ = "STUB: not implemented"; return nil }

// Client returns a client that accesses the simulated chain.
func (n *Backend) Client() Client { _ = "STUB: not implemented"; return *new(Client) }
