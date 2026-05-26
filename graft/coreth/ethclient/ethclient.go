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
// Copyright 2016 The go-ethereum Authors
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

// Package ethclient provides a client for the Ethereum RPC API.
package ethclient

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// DialContext connects a client to the given URL with context.
func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client { _ = "STUB: not implemented"; return nil }

// Close closes the underlying RPC connection.
func (ec *Client) Close() {
	_ = "STUB: not implemented"

	// Client gets the underlying RPC client.
	return
}

func (ec *Client) Client() *rpc.Client {
	_ = "STUB: not implemented"

	// Blockchain Access
	return nil
}

// ChainID retrieves the current chain ID for transaction replay protection.
func (ec *Client) ChainID(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockByHash returns the given full block.
//
// Note that loading full blocks requires two requests. Use HeaderByHash
// if you don't need all transactions or uncle headers.
func (ec *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (ec *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockNumber returns the most recent block number
func (ec *Client) BlockNumber(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PeerCount returns the number of p2p peers as reported by the net_peerCount method.
// func (ec *Client) PeerCount(ctx context.Context) (uint64, error) {
// 	var result hexutil.Uint64
// 	err := ec.c.CallContext(ctx, &result, "net_peerCount")
// 	return uint64(result), err
// }

// BlockReceipts returns the receipts of a given block number or hash.
func (ec *Client) BlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rpcBlock struct {
	Hash           common.Hash         `json:"hash"`
	Transactions   []rpcTransaction    `json:"transactions"`
	UncleHashes    []common.Hash       `json:"uncles"`
	Withdrawals    []*types.Withdrawal `json:"withdrawals,omitempty"`
	Version        uint32              `json:"version"`
	BlockExtraData *hexutil.Bytes      `json:"blockExtraData"`
}

func (ec *Client) getBlock(ctx context.Context, method string, args ...interface{}) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode header and transactions.

// When the block is not found, the API returns JSON null.

// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.

// Load uncles because they are not included in the block response.

// Fill the sender cache of transactions in the block.

// HeaderByHash returns the block header with the given hash.
func (ec *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (ec *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error { _ = "STUB: not implemented"; return nil }

// TransactionByHash returns the transaction with the given hash.
func (ec *Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
func (ec *Client) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	_ = "STUB: not implemented"
	// Try to load the address from the cache.
	return *new(common.Address), nil
}

// It was not found in cache, ask the server.

// TransactionCount returns the total number of transactions in the given block.
func (ec *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TransactionInBlock returns a single transaction at index in the given block.
func (ec *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (ec *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (ec *Client) SyncProgress(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// If not syncing, the response will be 'false'. To detect this
// we unmarshal into a boolean and return nil on success.
// If the chain is syncing, the engine will not forward the
// request to the chain and a non-nil err will be returned.

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (ec *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(ethereum.Subscription), nil
}

// Defensively prefer returning nil interface explicitly on error-path, instead
// of letting default golang behavior wrap it with non-nil interface that stores
// nil concrete type value.

// State Access

// NetworkID returns the network ID for this client.
func (ec *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (ec *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BalanceAtHash returns the wei balance of the given account.
func (ec *Client) BalanceAtHash(ctx context.Context, account common.Address, blockHash common.Hash) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (ec *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StorageAtHash returns the value of key in the contract storage of the given account.
func (ec *Client) StorageAtHash(ctx context.Context, account common.Address, key common.Hash, blockHash common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (ec *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CodeAtHash returns the contract code of the given account.
func (ec *Client) CodeAtHash(ctx context.Context, account common.Address, blockHash common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (ec *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NonceAtHash returns the account nonce of the given account.
func (ec *Client) NonceAtHash(ctx context.Context, account common.Address, blockHash common.Hash) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Filters

// FilterLogs executes a filter query.
func (ec *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (ec *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(ethereum.Subscription), nil
}

// Defensively prefer returning nil interface explicitly on error-path, instead
// of letting default golang behavior wrap it with non-nil interface that stores
// nil concrete type value.

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pending State

// PendingBalanceAt returns the wei balance of the given account in the pending state.
// func (ec *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
// 	var result hexutil.Big
// 	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, "pending")
// 	return (*big.Int)(&result), err
// }

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
// func (ec *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
// 	var result hexutil.Bytes
// 	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, "pending")
// 	return result, err
// }

// PendingCodeAt returns the contract code of the given account in the pending state.
// func (ec *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
// 	var result hexutil.Bytes
// 	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
// 	return result, err
// }

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
// func (ec *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
// 	var result hexutil.Uint64
// 	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")
// 	return uint64(result), err
// }

// PendingTransactionCount returns the total number of transactions in the pending state.
// func (ec *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
// 	var num hexutil.Uint
// 	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "pending")
// 	return uint(num), err
// }

// Contract Calling

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (ec *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CallContractAtHash is almost the same as CallContract except that it selects
// the block by block hash instead of block height.
func (ec *Client) CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
// func (ec *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
// 	var hex hexutil.Bytes
// 	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), "pending")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return hex, nil
// }

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (ec *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
// allow a timely execution of a transaction.
func (ec *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type feeHistoryResultMarshaling struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

// FeeHistory retrieves the fee market history.
func (ec *Client) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (ec *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func toBlockNumArg(number *big.Int) string { _ = "STUB: not implemented"; return "" }

// It's negative.

// It's negative and large, which is invalid.

func toCallArg(msg ethereum.CallMsg) interface{} { _ = "STUB: not implemented"; return nil }
