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

package core

import (
	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/triedb"
)

// CurrentHeader retrieves the current head header of the canonical chain. The
// header is retrieved from the HeaderChain's internal cache.
func (bc *BlockChain) CurrentHeader() *types.Header { _ = "STUB: not implemented"; return nil }

// CurrentBlock retrieves the current head block of the canonical chain. The
// block is retrieved from the blockchain's internal cache.
func (bc *BlockChain) CurrentBlock() *types.Header { _ = "STUB: not implemented"; return nil }

// HasHeader checks if a block header is present in the database or not, caching
// it if present.
func (bc *BlockChain) HasHeader(hash common.Hash, number uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// GetHeader retrieves a block header from the database by hash and number,
// caching it if found.
func (bc *BlockChain) GetHeader(hash common.Hash, number uint64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

// GetHeaderByHash retrieves a block header from the database by hash, caching it if
// found.
func (bc *BlockChain) GetHeaderByHash(hash common.Hash) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

// GetHeaderByNumber retrieves a block header from the database by number,
// caching it (associated with its hash) if found.
func (bc *BlockChain) GetHeaderByNumber(number uint64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

// GetBody retrieves a block body (transactions and uncles) from the database by
// hash, caching it if found.
func (bc *BlockChain) GetBody(hash common.Hash) *types.Body {
	_ = "STUB: not implemented"
	// Short circuit if the body's already in the cache, retrieve otherwise
	return nil
}

// Cache the found body for next time and return

// HasBlock checks if a block is fully present in the database or not.
func (bc *BlockChain) HasBlock(hash common.Hash, number uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// HasFastBlock checks if a fast block is fully present in the database or not.
func (bc *BlockChain) HasFastBlock(hash common.Hash, number uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// GetBlock retrieves a block from the database by hash and number,
// caching it if found.
func (bc *BlockChain) GetBlock(hash common.Hash, number uint64) *types.Block {
	_ = "STUB: not implemented"
	// Short circuit if the block's already in the cache, retrieve otherwise
	return nil
}

// Cache the found block for next time and return

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *BlockChain) GetBlockByHash(hash common.Hash) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockByNumber retrieves a block from the database by number, caching it
// (associated with its hash) if found.
func (bc *BlockChain) GetBlockByNumber(number uint64) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// GetBlocksFromHash returns the block corresponding to hash and up to n-1 ancestors.
// [deprecated by eth/62]
func (bc *BlockChain) GetBlocksFromHash(hash common.Hash, n int) (blocks []*types.Block) {
	_ = "STUB: not implemented"
	return nil
}

// GetReceiptsByHash retrieves the receipts for all transactions in a given block.
func (bc *BlockChain) GetReceiptsByHash(hash common.Hash) types.Receipts {
	_ = "STUB: not implemented"
	return *new(types.Receipts)
}

// GetCanonicalHash returns the canonical hash for a given block number
func (bc *BlockChain) GetCanonicalHash(number uint64) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// GetTransactionLookup retrieves the lookup along with the transaction
// itself associate with the given transaction hash.
//
// An error will be returned if the transaction is not found, and background
// indexing for transactions is still in progress. The transaction might be
// reachable shortly once it's indexed.
//
// A null will be returned in the transaction is not found and background
// transaction indexing is already finished. The transaction is not existent
// from the node's perspective.
func (bc *BlockChain) GetTransactionLookup(hash common.Hash) (*rawdb.LegacyTxLookupEntry, *types.Transaction, error) {
	_ = "STUB: not implemented"
	// Short circuit if the txlookup already in the cache, retrieve otherwise
	return nil, nil, nil
}

// The transaction is already indexed, the transaction is either
// not existent or not in the range of index, returning null.

// HasState checks if state trie is fully present in the database or not.
func (bc *BlockChain) HasState(hash common.Hash) bool { _ = "STUB: not implemented"; return false }

// HasBlockAndState checks if a block and associated state trie is fully present
// in the database or not, caching it if present.
func (bc *BlockChain) HasBlockAndState(hash common.Hash, number uint64) bool {
	_ = "STUB: not implemented"
	// Check first that the block itself is known
	return false
}

// State returns a new mutable state based on the current HEAD block.
func (bc *BlockChain) State() (*state.StateDB, error) { _ = "STUB: not implemented"; return nil, nil }

// StateAt returns a new mutable state based on a particular point in time.
func (bc *BlockChain) StateAt(root common.Hash) (*state.StateDB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Config retrieves the chain's fork configuration.
func (bc *BlockChain) Config() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

// Engine retrieves the blockchain's consensus engine.
func (bc *BlockChain) Engine() consensus.Engine {
	_ = "STUB: not implemented"

	// Snapshots returns the blockchain snapshot tree.
	return *new(consensus.Engine)
}

func (bc *BlockChain) Snapshots() *snapshot.Tree {
	_ = "STUB: not implemented"

	// Validator returns the current validator.
	return nil
}

func (bc *BlockChain) Validator() Validator {
	_ = "STUB: not implemented"
	return *

	// Processor returns the current processor.
	new(Validator)
}

func (bc *BlockChain) Processor() Processor {
	_ = "STUB: not implemented"
	return *

	// StateCache returns the caching database underpinning the blockchain instance.
	new(Processor)
}

func (bc *BlockChain) StateCache() state.Database {
	_ = "STUB: not implemented"
	return *

	// GasLimit returns the gas limit of the current HEAD block.
	new(state.Database)
}

func (bc *BlockChain) GasLimit() uint64 { _ = "STUB: not implemented"; return 0 }

// Genesis retrieves the chain's genesis block.
func (bc *BlockChain) Genesis() *types.Block { _ = "STUB: not implemented"; return nil }

// GetVMConfig returns the block chain VM config.
func (bc *BlockChain) GetVMConfig() *vm.Config { _ = "STUB: not implemented"; return nil }

// TrieDB retrieves the low level trie database used for data storage.
func (bc *BlockChain) TrieDB() *triedb.Database {
	_ = "STUB: not implemented"

	// HeaderChain returns the underlying header chain.
	return nil
}

func (bc *BlockChain) HeaderChain() *HeaderChain {
	_ = "STUB: not implemented"

	// SubscribeRemovedLogsEvent registers a subscription of RemovedLogsEvent.
	return nil
}

func (bc *BlockChain) SubscribeRemovedLogsEvent(ch chan<- RemovedLogsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeChainEvent registers a subscription of ChainEvent.
func (bc *BlockChain) SubscribeChainEvent(ch chan<- ChainEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeChainHeadEvent registers a subscription of ChainHeadEvent.
func (bc *BlockChain) SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeChainSideEvent registers a subscription of ChainSideEvent.
func (bc *BlockChain) SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeLogsEvent registers a subscription of []*types.Log.
func (bc *BlockChain) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeBlockProcessingEvent registers a subscription of bool where true means
// block processing has started while false means it has stopped.
func (bc *BlockChain) SubscribeBlockProcessingEvent(ch chan<- bool) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeChainAcceptedEvent registers a subscription of ChainEvent.
func (bc *BlockChain) SubscribeChainAcceptedEvent(ch chan<- ChainEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeAcceptedLogsEvent registers a subscription of accepted []*types.Log.
func (bc *BlockChain) SubscribeAcceptedLogsEvent(ch chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeAcceptedTransactionEvent registers a subscription of accepted transactions
func (bc *BlockChain) SubscribeAcceptedTransactionEvent(ch chan<- NewTxsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// GetLogs fetches all logs from a given block.
func (bc *BlockChain) GetLogs(hash common.Hash, number uint64) [][]*types.Log {
	_ = "STUB: not implemented"
	return nil
}

// this cache is thread-safe
