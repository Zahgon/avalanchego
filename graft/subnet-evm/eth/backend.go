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
// Copyright 2014 The go-ethereum Authors
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

// Package eth implements the Ethereum protocol.
package eth

import (
	"sync"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/consensus"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core/txpool"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/eth/ethconfig"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/internal/ethapi"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/internal/shutdowncheck"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/miner"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/node"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/libevm/accounts"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/bloombits"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/event"
)

// Config contains the configuration options of the ETH protocol.
// Deprecated: use ethconfig.Config instead.
type Config = ethconfig.Config

var DefaultSettings Settings = Settings{MaxBlocksPerRequest: 2000}

type Settings struct {
	MaxBlocksPerRequest int64 // Maximum number of blocks to serve per getLogs request
}

// PushGossiper sends pushes pending transactions to peers until they are
// removed from the mempool.
type PushGossiper interface {
	Add(*types.Transaction)
}

// Ethereum implements the Ethereum full node service.
type Ethereum struct {
	config *Config

	// Handlers
	txPool *txpool.TxPool

	blockchain *core.BlockChain
	gossiper   PushGossiper

	// DB interfaces
	chainDb ethdb.Database // Block chain database

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests     chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer      *core.ChainIndexer             // Bloom indexer operating during block imports
	closeBloomHandler chan struct{}

	APIBackend *EthAPIBackend

	miner     *miner.Miner
	etherbase common.Address

	networkID     uint64
	netRPCService *ethapi.NetAPI

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)

	shutdownTracker *shutdowncheck.ShutdownTracker // Tracks if and when the node has shutdown ungracefully

	stackRPCs []rpc.API

	settings Settings // Settings for Ethereum API
}

// roundUpCacheSize returns [input] rounded up to the next multiple of [allocSize]
func roundUpCacheSize(input int, allocSize int) int { _ = "STUB: not implemented"; return 0 }

// New creates a new Ethereum object (including the
// initialisation of the common Ethereum object)
func New(
	stack *node.Node,
	config *Config,
	gossiper PushGossiper,
	chainDb ethdb.Database,
	settings Settings,
	lastAcceptedHash common.Hash,
	engine consensus.Engine,
	clock *mockable.Clock,
) (*Ethereum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// round TrieCleanCache and SnapshotCache up to nearest 64MB, since fastcache will mmap
// memory in 64MBs chunks.

// Try to recover offline state pruning only in hash-based.

// Note: RecoverPruning must be called to handle the case that we are midway through offline pruning.
// If the data directory is changed in between runs preventing RecoverPruning from performing its job correctly,
// it may cause DB corruption.
// Since RecoverPruning will only continue a pruning run that already began, we do not need to ensure that
// reprocessState has already been called and completed successfully. To ensure this, we must maintain
// that Prune is only run after reprocessState has finished successfully.

// If the context is not set, avoid a panic. Only necessary during firewood use.

// Free airdrop data to save memory usage

// config.BlobPool.Datadir = ""
// blobPool := blobpool.New(config.BlobPool, &chainWithFinalBlock{eth.blockchain})

//, blobPool})

// Start the RPC service

// Successful startup; push a marker and check previous unclean shutdowns.

// APIs return the collection of RPC services the ethereum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Ethereum) APIs() []rpc.API { _ = "STUB: not implemented"; return nil }

// Append tracing APIs

// Add the APIs from the node

// Create [filterSystem] with the log cache size set in the config.

// Append all the local APIs and return

func (s *Ethereum) Etherbase() (eb common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// SetEtherbase sets the mining reward address.
func (s *Ethereum) SetEtherbase(etherbase common.Address) { _ = "STUB: not implemented"; return }

func (s *Ethereum) Miner() *miner.Miner { _ = "STUB: not implemented"; return nil }

func (s *Ethereum) AccountManager() *accounts.Manager { _ = "STUB: not implemented"; return nil }
func (s *Ethereum) BlockChain() *core.BlockChain      { _ = "STUB: not implemented"; return nil }
func (s *Ethereum) TxPool() *txpool.TxPool            { _ = "STUB: not implemented"; return nil }
func (s *Ethereum) EventMux() *event.TypeMux          { _ = "STUB: not implemented"; return nil }
func (s *Ethereum) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	return *new(consensus.Engine)
}
func (s *Ethereum) ChainDb() ethdb.Database { _ = "STUB: not implemented"; return *new(ethdb.Database) }

func (s *Ethereum) NetVersion() uint64               { _ = "STUB: not implemented"; return 0 }
func (s *Ethereum) ArchiveMode() bool                { _ = "STUB: not implemented"; return false }
func (s *Ethereum) BloomIndexer() *core.ChainIndexer { _ = "STUB: not implemented"; return nil }

// Start implements node.Lifecycle, starting all internal goroutines needed by the
// Ethereum protocol implementation.
func (s *Ethereum) Start() {
	_ = "STUB: not implemented"
	// Start the bloom bits servicing goroutines
	return
}

// Regularly update shutdown marker

// Stop implements node.Lifecycle, terminating all internal goroutines used by the
// Ethereum protocol.
// FIXME remove error from type if this will never return an error
func (s *Ethereum) Stop() error { _ = "STUB: not implemented"; return nil }

// Clean shutdown marker as the last thing before closing db

func (s *Ethereum) LastAcceptedBlock() *types.Block { _ = "STUB: not implemented"; return nil }

// precheckPopulateMissingTries returns an error if config flags should prevent
// [populateMissingTries]
//
// NOTE: [populateMissingTries] is called from [New] to ensure all
// state is repaired before any async processes (specifically snapshot re-generation)
// are started which could interfere with historical re-generation.
func (s *Ethereum) precheckPopulateMissingTries() error { _ = "STUB: not implemented"; return nil }

// Delete the populate missing tries marker to indicate that the node started with
// populate missing tries disabled.

// Note: Time Marker is written inside of [populateMissingTries] once it
// succeeds inside of [NewBlockChain]

func (s *Ethereum) handleOfflinePruning(cacheConfig *core.CacheConfig, gspec *core.Genesis, vmConfig vm.Config, lastAcceptedHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete the offline pruning marker to indicate that the node started with offline pruning disabled.

// Perform offline pruning after NewBlockChain has been called to ensure that we have rolled back the chain
// to the last accepted block before pruning begins.
// If offline pruning marker is on disk, then we force the node to be started with offline pruning disabled
// before allowing another run of offline pruning.

// Clean up middle roots

// Allow the blockchain to be garbage collected immediately, since we will shut down the chain after offline pruning completes.

// Note: Time Marker is written inside of [Prune] before compaction begins
// (considered an optional optimization)
