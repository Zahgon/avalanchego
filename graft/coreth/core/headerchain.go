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
// Copyright 2015 The go-ethereum Authors
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
	mrand "math/rand"
	"sync/atomic"

	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/lru"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
)

const (
	headerCacheLimit = 512
	tdCacheLimit     = 1024
	numberCacheLimit = 2048
)

// HeaderChain implements the basic block header chain logic that is shared by
// core.BlockChain and light.LightChain. It is not usable in itself, only as
// a part of either structure.
//
// HeaderChain is responsible for maintaining the header chain including the
// header query and updating.
//
// The components maintained by headerchain includes:
// (1) header (2) block hash -> number mapping (3) canonical number -> hash mapping
// and (4) head header flag.
//
// It is not thread safe either, the encapsulating chain structures should do
// the necessary mutex locking/unlocking.
type HeaderChain struct {
	config *params.ChainConfig

	chainDb       ethdb.Database
	genesisHeader *types.Header

	currentHeader     atomic.Value // Current head of the header chain (may be above the block chain!)
	currentHeaderHash common.Hash  // Hash of the current head of the header chain (prevent recomputing all the time)

	headerCache         *lru.Cache[common.Hash, *types.Header]
	numberCache         *lru.Cache[common.Hash, uint64]  // most recent block numbers
	acceptedNumberCache FIFOCache[uint64, *types.Header] // most recent accepted heights to headers (only modified in accept)

	rand   *mrand.Rand
	engine consensus.Engine
}

// NewHeaderChain creates a new HeaderChain structure. ProcInterrupt points
// to the parent's interrupt semaphore.
func NewHeaderChain(chainDb ethdb.Database, config *params.ChainConfig, cacheConfig *CacheConfig, engine consensus.Engine) (*HeaderChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Seed a fast but crypto originating random generator

// GetBlockNumber retrieves the block number belonging to the given hash
// from the cache or database
func (hc *HeaderChain) GetBlockNumber(hash common.Hash) *uint64 {
	_ = "STUB: not implemented"
	return nil
}

// GetHeader retrieves a block header from the database by hash and number,
// caching it if found.
func (hc *HeaderChain) GetHeader(hash common.Hash, number uint64) *types.Header {
	_ = "STUB: not implemented"
	// Short circuit if the header's already in the cache, retrieve otherwise
	return nil
}

// Cache the found header for next time and return

// GetHeaderByHash retrieves a block header from the database by hash, caching it if
// found.
func (hc *HeaderChain) GetHeaderByHash(hash common.Hash) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

// HasHeader checks if a block header is present in the database or not.
// In theory, if header is present in the database, all relative components
// like td and hash->number should be present too.
func (hc *HeaderChain) HasHeader(hash common.Hash, number uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// GetHeaderByNumber retrieves a block header from the database by number,
// caching it (associated with its hash) if found.
func (hc *HeaderChain) GetHeaderByNumber(number uint64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

func (hc *HeaderChain) GetCanonicalHash(number uint64) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// CurrentHeader retrieves the current head header of the canonical chain. The
// header is retrieved from the HeaderChain's internal cache.
func (hc *HeaderChain) CurrentHeader() *types.Header { _ = "STUB: not implemented"; return nil }

// SetCurrentHeader sets the in-memory head header marker of the canonical chan
// as the given header.
func (hc *HeaderChain) SetCurrentHeader(head *types.Header) { _ = "STUB: not implemented"; return }

// SetGenesis sets a new genesis block header for the chain
func (hc *HeaderChain) SetGenesis(head *types.Header) { _ = "STUB: not implemented"; return }

// Config retrieves the header chain's chain configuration.
func (hc *HeaderChain) Config() *params.ChainConfig {
	_ = "STUB: not implemented"

	// Engine retrieves the header chain's consensus engine.
	return nil
}

func (hc *HeaderChain) Engine() consensus.Engine {
	_ = "STUB: not implemented"

	// GetBlock implements consensus.ChainReader, and returns nil for every input as
	// a header chain does not have blocks available for retrieval.
	return *new(consensus.Engine)
}

func (hc *HeaderChain) GetBlock(hash common.Hash, number uint64) *types.Block {
	_ = "STUB: not implemented"
	return nil
}
