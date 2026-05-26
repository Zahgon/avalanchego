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
// Copyright 2021 The go-ethereum Authors
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

// Package subnetevmclient provides an RPC client for subnet-evm-specific APIs.
package subnetevmclient

import (
	"context"
	"math/big"
	"runtime"
	"runtime/debug"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// Client is a wrapper around rpc.Client that implements geth-specific functionality.
//
// If you want to use the standardized Ethereum RPC functionality, use ethclient.Client instead.
type Client struct {
	c *rpc.Client
}

// New creates a client that uses the given RPC client.
func New(c *rpc.Client) *Client {
	_ = "STUB: not implemented"

	// CreateAccessList tries to create an access list for a specific transaction based on the
	// current pending state of the blockchain.
	return nil
}

func (ec *Client) CreateAccessList(ctx context.Context, msg ethereum.CallMsg) (*types.AccessList, uint64, string, error) {
	_ = "STUB: not implemented"
	return nil, 0, "", nil
}

// AccountResult is the result of a GetProof operation.
type AccountResult struct {
	Address      common.Address  `json:"address"`
	AccountProof []string        `json:"accountProof"`
	Balance      *big.Int        `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        uint64          `json:"nonce"`
	StorageHash  common.Hash     `json:"storageHash"`
	StorageProof []StorageResult `json:"storageProof"`
}

// StorageResult provides a proof for a key-value pair.
type StorageResult struct {
	Key   string   `json:"key"`
	Value *big.Int `json:"value"`
	Proof []string `json:"proof"`
}

// GetProof returns the account and storage values of the specified account including the Merkle-proof.
// The block number can be nil, in which case the value is taken from the latest known block.
func (ec *Client) GetProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*AccountResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Avoid keys being 'null'.

// Turn hexutils back to normal datatypes

// OverrideAccount specifies the state of an account to be overridden.
type OverrideAccount struct {
	Nonce     uint64                      `json:"nonce"`
	Code      []byte                      `json:"code"`
	Balance   *big.Int                    `json:"balance"`
	State     map[common.Hash]common.Hash `json:"state"`
	StateDiff map[common.Hash]common.Hash `json:"stateDiff"`
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
//
// overrides specifies a map of contract states that should be overwritten before executing
// the message call.
// Please use ethclient.CallContract instead if you don't need the override functionality.
func (ec *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int, overrides *map[common.Address]OverrideAccount) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GCStats retrieves the current garbage collection stats from a geth node.
func (ec *Client) GCStats(ctx context.Context) (*debug.GCStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MemStats retrieves the current memory stats from a geth node.
func (ec *Client) MemStats(ctx context.Context) (*runtime.MemStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubscribePendingTransactions subscribes to new pending transactions.
func (ec *Client) SubscribePendingTransactions(ctx context.Context, ch chan<- common.Hash) (*rpc.ClientSubscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toCallArg(msg ethereum.CallMsg) interface{} { _ = "STUB: not implemented"; return nil }

func toOverrideMap(overrides *map[common.Address]OverrideAccount) interface{} {
	_ = "STUB: not implemented"
	return nil
}
