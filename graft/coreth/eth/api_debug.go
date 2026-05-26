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

package eth

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/graft/coreth/internal/ethapi"
	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
)

var errFirewoodNotSupported = errors.New("firewood triedb scheme does not yet support this operation")

// DebugAPI is the collection of Ethereum full node APIs for debugging the
// protocol.
type DebugAPI struct {
	eth *Ethereum
}

// NewDebugAPI creates a new DebugAPI instance.
func NewDebugAPI(eth *Ethereum) *DebugAPI { _ = "STUB: not implemented"; return nil }

// DumpBlock retrieves the entire state of the database at a given block.
func (api *DebugAPI) DumpBlock(blockNr rpc.BlockNumber) (state.Dump, error) {
	_ = "STUB: not implemented"
	return *new(state.Dump), nil
}

// Sanity limit over RPC

// Preimage is a debug API function that returns the preimage for a sha3 hash, if known.
func (api *DebugAPI) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetBadBlocks returns a list of the last 'bad blocks' that the client has seen on the network
// and returns them as a JSON list of block hashes.
func (api *DebugAPI) GetBadBlocks(ctx context.Context) ([]*ethapi.BadBlockArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountRangeMaxResults is the maximum number of results to be returned per call
const AccountRangeMaxResults = 256

// AccountRange enumerates all accounts in the given block and start point in paging request
func (api *DebugAPI) AccountRange(blockNrOrHash rpc.BlockNumberOrHash, start hexutil.Bytes, maxResults int, nocode, nostorage, incompletes bool) (state.Dump, error) {
	_ = "STUB: not implemented"
	return *new(state.Dump), nil
}

// StorageRangeResult is the result of a debug_storageRangeAt API call.
type StorageRangeResult struct {
	Storage storageMap   `json:"storage"`
	NextKey *common.Hash `json:"nextKey"` // nil if Storage includes the last key in the trie.
}

type storageMap map[common.Hash]storageEntry

type storageEntry struct {
	Key   *common.Hash `json:"key"`
	Value common.Hash  `json:"value"`
}

// StorageRangeAt returns the storage at the given block height and transaction index.
func (api *DebugAPI) StorageRangeAt(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (StorageRangeResult, error) {
	_ = "STUB: not implemented"
	return *new(StorageRangeResult), nil
}

func storageRangeAt(statedb *state.StateDB, root common.Hash, address common.Address, start []byte, maxResult int) (StorageRangeResult, error) {
	_ = "STUB: not implemented"
	return *new(StorageRangeResult), nil
}

// empty storage

// Add the 'next key' so clients can continue downloading.

// GetModifiedAccountsByNumber returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
func (api *DebugAPI) GetModifiedAccountsByNumber(startNum uint64, endNum *uint64) ([]common.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetModifiedAccountsByHash returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
func (api *DebugAPI) GetModifiedAccountsByHash(startHash common.Hash, endHash *common.Hash) ([]common.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) getModifiedAccounts(startBlock, endBlock *types.Block) ([]common.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccessibleState returns the first number where the node has accessible
// state on disk. Note this being the post-state of that block and the pre-state
// of the next block.
// The (from, to) parameters are the sequence of blocks to search, which can go
// either forwards or backwards
func (api *DebugAPI) GetAccessibleState(from, to rpc.BlockNumber) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We don't have state for pending (-2), so treat it as latest

func (api *DebugAPI) isFirewood() bool { _ = "STUB: not implemented"; return false }
