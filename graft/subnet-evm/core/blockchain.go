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

// Package core implements the Ethereum consensus protocol.
package core

import (
	"errors"
	"io"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/consensus"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/lru"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/metrics"
	"github.com/ava-labs/libevm/triedb"

	// Force libevm metrics of the same name to be registered first.
	_ "github.com/ava-labs/libevm/core"
)

// ====== If resolving merge conflicts ======
//
// All calls to metrics.NewRegistered*() for metrics also defined in libevm/core have been
// replaced either with:
//   - metrics.GetOrRegister*() to get a metric already registered in libevm/core, or register it
//     here otherwise
//   - [getOrOverrideAsRegisteredCounter] to get a metric already registered in libevm/core
//     only if it is a [metrics.Counter]. If it is not, the metric is unregistered and registered
//     as a [metrics.Counter] here.
//
// These replacements ensure the same metrics are shared between the two packages.
var (
	accountReadTimer         = getOrOverrideAsRegisteredCounter("chain/account/reads", nil)
	accountHashTimer         = getOrOverrideAsRegisteredCounter("chain/account/hashes", nil)
	accountUpdateTimer       = getOrOverrideAsRegisteredCounter("chain/account/updates", nil)
	accountCommitTimer       = getOrOverrideAsRegisteredCounter("chain/account/commits", nil)
	storageReadTimer         = getOrOverrideAsRegisteredCounter("chain/storage/reads", nil)
	storageHashTimer         = getOrOverrideAsRegisteredCounter("chain/storage/hashes", nil)
	storageUpdateTimer       = getOrOverrideAsRegisteredCounter("chain/storage/updates", nil)
	storageCommitTimer       = getOrOverrideAsRegisteredCounter("chain/storage/commits", nil)
	snapshotAccountReadTimer = getOrOverrideAsRegisteredCounter("chain/snapshot/account/reads", nil)
	snapshotStorageReadTimer = getOrOverrideAsRegisteredCounter("chain/snapshot/storage/reads", nil)
	snapshotCommitTimer      = getOrOverrideAsRegisteredCounter("chain/snapshot/commits", nil)

	triedbCommitTimer = getOrOverrideAsRegisteredCounter("chain/triedb/commits", nil)

	blockInsertTimer            = metrics.GetOrRegisterCounter("chain/block/inserts", nil)
	blockSignatureRecoveryTimer = metrics.GetOrRegisterCounter("chain/block/signature/recovery", nil)
	blockInsertCount            = metrics.GetOrRegisterCounter("chain/block/inserts/count", nil)
	blockContentValidationTimer = metrics.GetOrRegisterCounter("chain/block/validations/content", nil)
	blockStateInitTimer         = metrics.GetOrRegisterCounter("chain/block/inits/state", nil)
	blockExecutionTimer         = metrics.GetOrRegisterCounter("chain/block/executions", nil)
	blockTrieOpsTimer           = metrics.GetOrRegisterCounter("chain/block/trie", nil)
	blockValidationTimer        = metrics.GetOrRegisterCounter("chain/block/validations/state", nil)
	blockWriteTimer             = metrics.GetOrRegisterCounter("chain/block/writes", nil)

	acceptorQueueGauge            = metrics.GetOrRegisterGauge("chain/acceptor/queue/size", nil)
	acceptorWorkTimer             = metrics.GetOrRegisterCounter("chain/acceptor/work", nil)
	acceptorWorkCount             = metrics.GetOrRegisterCounter("chain/acceptor/work/count", nil)
	processedBlockGasUsedCounter  = metrics.GetOrRegisterCounter("chain/block/gas/used/processed", nil)
	lastAcceptedBlockBaseFeeGauge = metrics.NewRegisteredGauge("chain/block/fee/basefee", nil)
	blockTotalFeesGauge           = metrics.NewRegisteredGauge("chain/block/fee/total", nil)
	acceptedBlockGasUsedCounter   = metrics.GetOrRegisterCounter("chain/block/gas/used/accepted", nil)
	badBlockCounter               = metrics.GetOrRegisterCounter("chain/block/bad/count", nil)

	txUnindexTimer      = metrics.GetOrRegisterCounter("chain/txs/unindex", nil)
	acceptedTxsCounter  = metrics.GetOrRegisterCounter("chain/txs/accepted", nil)
	processedTxsCounter = metrics.GetOrRegisterCounter("chain/txs/processed", nil)

	acceptedLogsCounter  = metrics.GetOrRegisterCounter("chain/logs/accepted", nil)
	processedLogsCounter = metrics.GetOrRegisterCounter("chain/logs/processed", nil)

	latestMinDelayGauge       = metrics.NewRegisteredGauge("chain/latest/mindelay", nil)
	latestMinDelayExcessGauge = metrics.NewRegisteredGauge("chain/latest/mindelay/excess", nil)

	ErrRefuseToCorruptArchiver = errors.New("node has operated with pruning disabled, shutting down to prevent missing tries")

	errFutureBlockUnsupported  = errors.New("future block insertion not supported")
	errCacheConfigNotSpecified = errors.New("must specify cache config")
	errInvalidOldChain         = errors.New("invalid old chain")
	errInvalidNewChain         = errors.New("invalid new chain")
)

const (
	bodyCacheLimit           = 256
	blockCacheLimit          = 256
	receiptsCacheLimit       = 32
	txLookupCacheLimit       = 1024
	feeConfigCacheLimit      = 256
	coinbaseConfigCacheLimit = 256
	badBlockLimit            = 10

	// BlockChainVersion ensures that an incompatible database forces a resync from scratch.
	//
	// Changelog:
	//
	// - Version 4
	//   The following incompatible database changes were added:
	//   * the `BlockNumber`, `TxHash`, `TxIndex`, `BlockHash` and `Index` fields of log are deleted
	//   * the `Bloom` field of receipt is deleted
	//   * the `BlockIndex` and `TxIndex` fields of txlookup are deleted
	// - Version 5
	//  The following incompatible database changes were added:
	//    * the `TxHash`, `GasCost`, and `ContractAddress` fields are no longer stored for a receipt
	//    * the `TxHash`, `GasCost`, and `ContractAddress` fields are computed by looking up the
	//      receipts' corresponding block
	// - Version 6
	//  The following incompatible database changes were added:
	//    * Transaction lookup information stores the corresponding block number instead of block hash
	// - Version 7
	//  The following incompatible database changes were added:
	//    * Use freezer as the ancient database to maintain all ancient data
	// - Version 8
	//  The following incompatible database changes were added:
	//    * New scheme for contract code in order to separate the codes and trie nodes
	BlockChainVersion uint64 = 8

	// statsReportLimit is the time limit during import and export after which we
	// always print out progress. This avoids the user wondering what's going on.
	statsReportLimit = 8 * time.Second

	// trieCleanCacheStatsNamespace is the namespace to surface stats from the trie
	// clean cache's underlying fastcache.
	trieCleanCacheStatsNamespace = "hashdb/memcache/clean/fastcache"
)

// cacheableFeeConfig encapsulates fee configuration itself and the block number that it has changed at,
// in order to cache them together.
type cacheableFeeConfig struct {
	feeConfig     commontype.FeeConfig
	lastChangedAt *big.Int
}

// cacheableCoinbaseConfig encapsulates coinbase address itself and allowFeeRecipient flag,
// in order to cache them together.
type cacheableCoinbaseConfig struct {
	coinbaseAddress    common.Address
	allowFeeRecipients bool
}

// CacheConfig contains the configuration values for the trie database
// and state snapshot these are resident in a blockchain.
type CacheConfig struct {
	TrieCleanLimit                  int     // Memory allowance (MB) to use for caching trie nodes in memory
	TrieDirtyLimit                  int     // Memory limit (MB) at which to block on insert and force a flush of dirty trie nodes to disk
	TrieDirtyCommitTarget           int     // Memory limit (MB) to target for the dirties cache before invoking commit
	TriePrefetcherParallelism       int     // Max concurrent disk reads trie prefetcher should perform at once
	CommitInterval                  uint64  // Commit the trie every [CommitInterval] blocks.
	Pruning                         bool    // Whether to disable trie write caching and GC altogether (archive node)
	AcceptorQueueLimit              int     // Blocks to queue before blocking during acceptance
	PopulateMissingTries            *uint64 // If non-nil, sets the starting height for re-generating historical tries.
	PopulateMissingTriesParallelism int     // Number of readers to use when trying to populate missing tries.
	AllowMissingTries               bool    // Whether to allow an archive node to run with pruning enabled
	SnapshotDelayInit               bool    // Whether to initialize snapshots on startup or wait for external call (= StateSyncEnabled)
	SnapshotLimit                   int     // Memory allowance (MB) to use for caching snapshot entries in memory
	SnapshotVerify                  bool    // Verify generated snapshots
	Preimages                       bool    // Whether to store preimage of trie key to the disk
	AcceptedCacheSize               int     // Depth of accepted headers cache and accepted logs cache at the accepted tip
	TransactionHistory              uint64  // Number of recent blocks for which to maintain transaction lookup indices
	SkipTxIndexing                  bool    // Whether to skip transaction indexing
	StateHistory                    uint64  // Number of blocks from head whose state histories are reserved.
	StateScheme                     string  // Scheme used to store ethereum states and merkle tree nodes on top

	ChainDataDir    string // Directory to store chain data in (used by Firewood)
	SnapshotNoBuild bool   // Whether the background generation is allowed
	SnapshotWait    bool   // Wait for snapshot construction on startup. TODO(karalabe): This is a dirty hack for testing, nuke it
}

// triedbConfig derives the configures for trie database.
func (c *CacheConfig) triedbConfig() *triedb.Config { _ = "STUB: not implemented"; return nil }

// ChainDataDir may not be set during some tests, where this path won't be called.

// must be at least 2

// DefaultCacheConfig are the default caching values if none are specified by the
// user (also used during testing).
var DefaultCacheConfig = &CacheConfig{
	TrieCleanLimit:            256,
	TrieDirtyLimit:            256,
	TrieDirtyCommitTarget:     20, // 20% overhead in memory counting (this targets 16 MB)
	TriePrefetcherParallelism: 16,
	Pruning:                   true,
	CommitInterval:            4096,
	AcceptorQueueLimit:        64, // Provides 2 minutes of buffer (2s block target) for a commit delay
	SnapshotLimit:             256,
	AcceptedCacheSize:         32,
	StateHistory:              32, // Default state history size
	StateScheme:               rawdb.HashScheme,
}

// DefaultCacheConfigWithScheme returns a deep copied default cache config with
// a provided trie node scheme.
func DefaultCacheConfigWithScheme(scheme string) *CacheConfig {
	_ = "STUB: not implemented"
	return nil
}

// TODO: remove this once if Firewood supports snapshots

// no snapshot allowed for firewood

// txLookup is wrapper over transaction lookup along with the corresponding
// transaction object.
type txLookup struct {
	lookup      *rawdb.LegacyTxLookupEntry
	transaction *types.Transaction
}

// BlockChain represents the canonical chain given a database with a genesis
// block. The Blockchain manages chain imports, reverts, chain reorganisations.
//
// Importing blocks in to the block chain happens according to the set of rules
// defined by the two stage Validator. Processing of blocks is done using the
// Processor which processes the included transaction. The validation of the state
// is done in the second part of the Validator. Failing results in aborting of
// the import.
//
// The BlockChain also helps in returning blocks from **any** chain included
// in the database as well as blocks that represents the canonical chain. It's
// important to note that GetBlock can return any block and does not need to be
// included in the canonical one where as GetBlockByNumber always represents the
// canonical chain.
type BlockChain struct {
	chainConfig *params.ChainConfig // Chain & network configuration
	cacheConfig *CacheConfig        // Cache configuration for pruning

	db           ethdb.Database   // Low level persistent database to store final content in
	snaps        *snapshot.Tree   // Snapshot tree for fast trie leaf access
	triedb       *triedb.Database // The database handler for maintaining trie nodes.
	stateCache   state.Database   // State database to reuse between imports (contains state cache)
	txIndexer    *txIndexer       // Transaction indexer, might be nil if not enabled
	stateManager TrieWriter

	hc                *HeaderChain
	rmLogsFeed        event.Feed
	chainFeed         event.Feed
	chainSideFeed     event.Feed
	chainHeadFeed     event.Feed
	chainAcceptedFeed event.Feed
	logsFeed          event.Feed
	logsAcceptedFeed  event.Feed
	blockProcFeed     event.Feed
	txAcceptedFeed    event.Feed
	scope             event.SubscriptionScope
	genesis           *Genesis
	genesisBlock      *types.Block

	// This mutex synchronizes chain write operations.
	// Readers don't need to take it, they can just read the database.
	chainmu sync.RWMutex

	currentBlock atomic.Pointer[types.Header] // Current head of the block chain

	bodyCache           *lru.Cache[common.Hash, *types.Body]              // Cache for the most recent block bodies
	receiptsCache       *lru.Cache[common.Hash, []*types.Receipt]         // Cache for the most recent receipts per block
	blockCache          *lru.Cache[common.Hash, *types.Block]             // Cache for the most recent entire blocks
	txLookupCache       *lru.Cache[common.Hash, txLookup]                 // Cache for the most recent transaction lookup data.
	badBlocks           *lru.Cache[common.Hash, *badBlock]                // Cache for bad blocks
	feeConfigCache      *lru.Cache[common.Hash, *cacheableFeeConfig]      // Cache for the most recent feeConfig lookup data.
	coinbaseConfigCache *lru.Cache[common.Hash, *cacheableCoinbaseConfig] // Cache for the most recent coinbaseConfig lookup data.

	stopping atomic.Bool // false if chain is running, true when stopped

	engine    consensus.Engine
	validator Validator // Block and state validator interface
	processor Processor // Block transaction processor interface
	vmConfig  vm.Config

	lastAccepted *types.Block // Prevents reorgs past this height

	senderCacher *TxSenderCacher

	// [acceptorQueue] is a processing queue for the Acceptor. This is
	// different than [chainAcceptedFeed], which is sent an event after an accepted
	// block is processed (after each loop of the accepted worker). If there is a
	// clean shutdown, all items inserted into the [acceptorQueue] will be processed.
	acceptorQueue chan *types.Block

	// [acceptorClosingLock], and [acceptorClosed] are used
	// to synchronize the closing of the [acceptorQueue] channel.
	//
	// Because we can't check if a channel is closed without reading from it
	// (which we don't want to do as we may remove a processing block), we need
	// to use a second variable to ensure we don't close a closed channel.
	acceptorClosingLock sync.RWMutex
	acceptorClosed      bool

	// [acceptorWg] is used to wait for the acceptorQueue to clear. This is used
	// during shutdown and in tests.
	acceptorWg sync.WaitGroup

	// [wg] is used to wait for the async blockchain processes to finish on shutdown.
	wg sync.WaitGroup

	// quit channel is used to listen for when the blockchain is shut down to close
	// async processes.
	// WaitGroups are used to ensure that async processes have finished during shutdown.
	quit chan struct{}

	// [acceptorTip] is the last block processed by the acceptor. This is
	// returned as the LastAcceptedBlock() to ensure clients get only fully
	// processed blocks. This may be equal to [lastAccepted].
	acceptorTip     *types.Block
	acceptorTipLock sync.Mutex

	// [flattenLock] prevents the [acceptor] from flattening snapshots while
	// a block is being verified.
	flattenLock sync.Mutex

	// [acceptedLogsCache] stores recently accepted logs to improve the performance of eth_getLogs.
	acceptedLogsCache FIFOCache[common.Hash, [][]*types.Log]

	// [txIndexTailLock] is used to synchronize the updating of the tx index tail.
	txIndexTailLock sync.Mutex
}

// NewBlockChain returns a fully initialised block chain using information
// available in the database. It initialises the default Ethereum Validator and
// Processor.
func NewBlockChain(
	db ethdb.Database, cacheConfig *CacheConfig, genesis *Genesis, engine consensus.Engine,
	vmConfig vm.Config, lastAcceptedHash common.Hash, skipChainConfigCheckCompatible bool,
) (*BlockChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open trie database with provided config

// Setup the genesis block, commit the provided genesis specification
// to database if the genesis block is not present yet, or load the
// stored one from database.
// Note: In go-ethereum, the code rewinds the chain on an incompatible config upgrade.
// We don't do this and expect the node operator to always update their node's configuration
// before network upgrades take effect.

// Create the state manager

// Re-generate current block state if it is missing

// After loading the last state (and reprocessing if necessary), we are
// guaranteed that [acceptorTip] is equal to [lastAccepted].
//
// It is critical to update this vaue before performing any state repairs so
// that all accepted blocks can be considered.

// Make sure the state associated with the block is available

// Populate missing tries if required

// If snapshot initialization is delayed for fast sync, skip initializing it here.
// This assumes that no blocks will be processed until ResetState is called to initialize
// the state of fast sync.

// Load any existing snapshot, regenerating it if loading failed (if not
// already initialized in recovery)

// Warm up [hc.acceptedNumberCache] and [acceptedLogsCache]

// if txlookup limit is 0 (uindexing disabled), we don't need to repair the tx index tail.

// Start processing accepted blocks effects in the background

// Start tx indexer if it's enabled.

// writeBlockAcceptedIndices writes any indices that must be persisted for accepted block.
// This includes the following:
// - transaction lookup indices
// - updating the acceptor tip index
func (bc *BlockChain) writeBlockAcceptedIndices(b *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *BlockChain) batchBlockAcceptedIndices(batch ethdb.Batch, b *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// flattenSnapshot attempts to flatten a block of [hash] to disk.
func (bc *BlockChain) flattenSnapshot(postAbortWork func() error, hash common.Hash) error {
	_ = "STUB: not implemented"
	// If snapshots are not initialized, perform [postAbortWork] immediately.
	return nil
}

// Abort snapshot generation before pruning anything from trie database
// (could occur in AcceptTrie)

// Perform work after snapshot generation is aborted (typically trie updates)

// Ensure we avoid flattening the snapshot while we are processing a block, or
// block execution will fallback to reading from the trie (which is much
// slower).

// Flatten the entire snap Trie to disk
//
// Note: This resumes snapshot generation.

// warmAcceptedCaches fetches previously accepted headers and logs from disk to
// pre-populate [hc.acceptedNumberCache] and [acceptedLogsCache].
func (bc *BlockChain) warmAcceptedCaches() { _ = "STUB: not implemented"; return }

// This could occur if we haven't accepted any blocks yet

// last accepted lookback is inclusive, so we reduce size by 1

// This could happen if a node state-synced

// TODO: handle blocks written to disk during state sync

// startAcceptor starts processing items on the [acceptorQueue]. If a [nil]
// object is placed on the [acceptorQueue], the [startAcceptor] will exit.
func (bc *BlockChain) startAcceptor() { _ = "STUB: not implemented"; return }

// Update acceptor tip and transaction lookup index
// Write this prior to state changes to allow easier reconstruction in `reprocessState`.

// Ensure [hc.acceptedNumberCache] and [acceptedLogsCache] have latest content

// Update the acceptor tip before sending events to ensure that any client acting based off of
// the events observes the updated acceptorTip on subsequent requests

// Update accepted feeds

// Note: in contrast to most accepted metrics, we increment the accepted log metrics in the acceptor queue because
// the logs are already processed in the acceptor queue.

// addAcceptorQueue adds a new *types.Block to the [acceptorQueue]. This will
// block if there are [AcceptorQueueLimit] items in [acceptorQueue].
func (bc *BlockChain) addAcceptorQueue(b *types.Block) {
	_ = "STUB: not implemented"
	// We only acquire a read lock here because it is ok to add items to the
	// [acceptorQueue] concurrently.
	return
}

// DrainAcceptorQueue blocks until all items in [acceptorQueue] have been
// processed.
func (bc *BlockChain) DrainAcceptorQueue() { _ = "STUB: not implemented"; return }

// stopAcceptor sends a signal to the Acceptor to stop processing accepted
// blocks. The Acceptor will exit once all items in [acceptorQueue] have been
// processed.
func (bc *BlockChain) stopAcceptor() { _ = "STUB: not implemented"; return }

// If [acceptorClosed] is already false, we should just return here instead
// of attempting to close [acceptorQueue] more than once (will cause
// a panic).
//
// This typically happens when a test calls [stopAcceptor] directly (prior to
// shutdown) and then [stopAcceptor] is called again in shutdown.

// Although nothing should be added to [acceptorQueue] after
// [acceptorClosed] is updated, we close the channel so the Acceptor
// goroutine exits.

func (bc *BlockChain) InitializeSnapshots() { _ = "STUB: not implemented"; return }

// SenderCacher returns the *TxSenderCacher used within the core package.
func (bc *BlockChain) SenderCacher() *TxSenderCacher { _ = "STUB: not implemented"; return nil }

// loadLastState loads the last known chain state from the database. This method
// assumes that the chain manager mutex is held.
func (bc *BlockChain) loadLastState(lastAcceptedHash common.Hash) error {
	_ = "STUB: not implemented"
	// Initialize genesis state
	return nil
}

// Restore the last known head block

// Make sure the entire head block is available

// Everything seems to be fine, set as the head block

// Restore the last known head header

// Otherwise, set the last accepted block and perform a re-org.

// This ensures that the head block is updated to the last accepted block on startup

// reprocessState is necessary to ensure that the last accepted state is
// available. The state may not be available if it was not committed due
// to an unclean shutdown.

func (bc *BlockChain) loadGenesisState() error {
	_ = "STUB: not implemented"
	// Prepare the genesis block and reinitialise the chain
	return nil
}

// Last update all in-memory chain markers

// Export writes the active chain to the given writer.
func (bc *BlockChain) Export(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// ExportN writes a subset of the active chain to the given writer.
func (bc *BlockChain) ExportN(w io.Writer, first uint64, last uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportCallback invokes [callback] for every block from [first] to [last] in order.
func (bc *BlockChain) ExportCallback(callback func(block *types.Block) error, first uint64, last uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// writeHeadBlock injects a new head block into the current block chain. This method
// assumes that the block is indeed a true head. It will also reset the head
// header to this very same block if they are older or if they are on a different side chain.
//
// Note, this function assumes that the `mu` mutex is held!
func (bc *BlockChain) writeHeadBlock(block *types.Block) {
	_ = "STUB: not implemented"
	// If the block is on a side chain or an unknown one, force other heads onto it too
	// Add the block to the canonical chain number scheme and mark as the head
	return
}

// Flush the whole batch into the disk, exit the node if failed

// Update all in-memory chain markers in the last step

// ValidateCanonicalChain confirms a canonical chain is well-formed.
func (bc *BlockChain) ValidateCanonicalChain() error {
	_ = "STUB: not implemented"
	// Ensure all accepted blocks are fully processed
	return nil
}

// Lookup the full block to get the transactions

// Transactions are only indexed beneath the last accepted block, so we only check
// that the transactions have been indexed, if we are checking below the last accepted
// block.

// Ensure that all of the transactions have been stored correctly in the canonical
// chain

// stopWithoutSaving stops the blockchain service. If any imports are currently in progress
// it will abort them using the procInterrupt. This method stops all running
// goroutines, but does not do all the post-stop work of persisting data.
// OBS! It is generally recommended to use the Stop method!
// This method has been exposed to allow tests to stop the blockchain while simulating
// a crash.
func (bc *BlockChain) stopWithoutSaving() { _ = "STUB: not implemented"; return }

// Signal shutdown tx indexer.

// Wait for accepted feed to process all remaining items

// Stop senderCacher's goroutines

// Unsubscribe all subscriptions registered from blockchain.

// Waiting for background processes to complete

// Stop stops the blockchain service. If any imports are currently in progress
// it will abort them using the procInterrupt.
func (bc *BlockChain) Stop() { _ = "STUB: not implemented"; return }

// Stop snapshot generation and release resources

// Ensure that the in-memory trie nodes are journaled to disk properly.

// Close the trie database, release all the held resources as the last step.

// SetPreference attempts to update the head block to be the provided block and
// emits a ChainHeadEvent if successful. This function will handle all reorg
// side effects, if necessary.
//
// Note: This function should ONLY be called on blocks that have already been
// inserted into the chain.
//
// Assumes [bc.chainmu] is not held by the caller.
func (bc *BlockChain) SetPreference(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// setPreference attempts to update the head block to be the provided block and
// emits a ChainHeadEvent if successful. This function will handle all reorg
// side effects, if necessary.
//
// Assumes [bc.chainmu] is held by the caller.
func (bc *BlockChain) setPreference(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Return early if the current block is already the block
// we are trying to write.

// writeKnownBlock updates the head block and will handle any reorg side
// effects automatically.

// Send a ChainHeadEvent if we end up altering
// the head block. Many internal aysnc processes rely on
// receiving these events (i.e. the TxPool).

// LastConsensusAcceptedBlock returns the last block to be marked as accepted. It may or
// may not yet be processed.
func (bc *BlockChain) LastConsensusAcceptedBlock() *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// LastAcceptedBlock returns the last block to be marked as accepted and is
// processed.
//
// Note: During initialization, [acceptorTip] is equal to [lastAccepted].
func (bc *BlockChain) LastAcceptedBlock() *types.Block { _ = "STUB: not implemented"; return nil }

// Accept sets a minimum height at which no reorg can pass. Additionally,
// this function may trigger a reorg if the block being accepted is not in the
// canonical chain.
//
// Assumes [bc.chainmu] is not held by the caller.
func (bc *BlockChain) Accept(block *types.Block) error { _ = "STUB: not implemented"; return nil }

// The parent of [block] must be the last accepted block.

// If the canonical hash at the block height does not match the block we are
// accepting, we need to trigger a reorg.

// Enqueue block in the acceptor

// TotalFees computes total consumed fees in wei. Block transactions and receipts have to have the same order.
func TotalFees(block *types.Block, receipts []*types.Receipt) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// legacy block, no baseFee

// TotalFees computes total consumed fees in ether. Block transactions and receipts have to have the same order.
func TotalFeesFloat(block *types.Block, receipts []*types.Receipt) (*big.Float, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bc *BlockChain) Reject(block *types.Block) error { _ = "STUB: not implemented"; return nil }

// Reject Trie

// Remove the block since its data is no longer needed

// Remove the block from the block cache (ignore return value of whether it was in the cache)

// writeKnownBlock updates the head block flag with a known block
// and introduces chain reorg if necessary.
func (bc *BlockChain) writeKnownBlock(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// writeCanonicalBlockWithLogs writes the new head [block] and emits events
// for the new head block.
func (bc *BlockChain) writeCanonicalBlockWithLogs(block *types.Block, logs []*types.Log) {
	_ = "STUB: not implemented"
	return
}

// newTip returns a boolean indicating if the block should be appended to
// the canonical chain.
func (bc *BlockChain) newTip(block *types.Block) bool { _ = "STUB: not implemented"; return false }

// writeBlockAndSetHead persists the block and associated state to the database
// and optimistically updates the canonical chain if [block] extends the current
// canonical chain.
// writeBlockAndSetHead expects to be the last verification step during InsertBlock
// since it creates a reference that will only be cleaned up by Accept/Reject.
func (bc *BlockChain) writeBlockAndSetHead(block *types.Block, parentRoot common.Hash, receipts []*types.Receipt, logs []*types.Log, state *state.StateDB) error {
	_ = "STUB: not implemented"
	return nil
}

// If [block] represents a new tip of the canonical chain, we optimistically add it before
// setPreference is called. Otherwise, we consider it a side chain block.

// writeBlockWithState writes the block and all associated state to the database,
// but it expects the chain mutex to be held.
func (bc *BlockChain) writeBlockWithState(block *types.Block, parentRoot common.Hash, receipts []*types.Receipt, state *state.StateDB) error {
	_ = "STUB: not implemented"
	// Irrelevant of the canonical status, write the block itself to the database.
	//
	// Note all the components of block(hash->number map, header, body, receipts)
	// should be written atomically. BlockBatch is used for containing all components.
	return nil
}

// Commit all cached state changes into underlying memory database.

// If node is running in path mode, skip explicit gc operation
// which is unnecessary in this mode.

// Note: if InsertTrie must be the last step in verification that can return an error.
// This allows [stateManager] to assume that if it inserts a trie without returning an
// error then the block has passed verification and either AcceptTrie/RejectTrie will
// eventually be called on [root] unless a fatal error occurs. It does not assume that
// the node will not shutdown before either AcceptTrie/RejectTrie is called.

// InsertChain attempts to insert the given batch of blocks in to the canonical
// chain or, otherwise, create a fork. If an error is returned it will return
// the index number of the failing block as well an error describing what went
// wrong.
//
// After insertion is done, all accumulated events will be fired.
func (bc *BlockChain) InsertChain(chain types.Blocks) (int, error) {
	_ = "STUB: not implemented"
	// Sanity check that we have something meaningful to import
	return 0, nil
}

// Do a sanity check that the provided chain is actually ordered and linked.

// Pre-checks passed, start the full block imports

func (bc *BlockChain) InsertBlock(block *types.Block) error { _ = "STUB: not implemented"; return nil }

func (bc *BlockChain) InsertBlockManual(block *types.Block, writes bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *BlockChain) insertBlock(block *types.Block, writes bool) error {
	_ = "STUB: not implemented"
	return nil
}

// even if the block is already known, we still need to generate the
// snapshot layer and add a reference to the triedb, so we re-execute
// the block. Note that insertBlock should only be called on a block
// once if it returns nil

// If an ancestor has been pruned, then this block cannot be acceptable.

// Future blocks are not supported, but should not be reported, so we return an error
// early here

// Some other error occurred, abort

// No validation errors for the block

// Retrieve the parent block to determine which root to build state on

// Instantiate the statedb to use for processing transactions
//
// NOTE: Flattening a snapshot during block execution requires fetching state
// entries directly from the trie (much slower).

// Enable prefetching to pull in trie node paths while processing transactions

// Process block using the parent state as reference point

// Validate the state using the default validator

// Update the metrics touched during block processing and validation
// Account reads are complete(in processing)
// Storage reads are complete(in processing)
// Account reads are complete(in processing)
// Storage reads are complete(in processing)
// Account updates are complete(in validation)
// Storage updates are complete(in validation)
// Account hashes are complete(in validation)
// Storage hashes are complete(in validation)
// The time spent on tries hashing
// The time spent on tries update
// The time spent on account read
// The time spent on storage read
// The time spent on EVM processing
// The time spent on block validation
// The time spent on trie operations

// If [writes] are disabled, skip [writeBlockWithState] so that we do not write the block
// or the state trie to disk.
// Note: in pruning mode, this prevents us from generating a reference to the state root.

// Write the block to the chain and get the status.
// writeBlockWithState (called within writeBlockAndSethead) creates a reference that
// will be cleaned up in Accept/Reject so we need to ensure an error cannot occur
// later in verification, since that would cause the referenced root to never be dereferenced.

// Update the metrics touched during block commit
// Account commits are complete, we can mark them
// Storage commits are complete, we can mark them
// Snapshot commits are complete, we can mark them
// Trie database commits are complete, we can mark them

// collectUnflattenedLogs collects the logs that were generated or removed during
// the processing of a block.
func (bc *BlockChain) collectUnflattenedLogs(b *types.Block, removed bool) [][]*types.Log {
	_ = "STUB: not implemented"
	return nil
}

// Note: gross but this needs to be initialized here because returning nil will be treated specially as an incorrect
// error case downstream.

// collectLogs collects the logs that were generated or removed during
// the processing of a block. These logs are later announced as deleted or reborn.
func (bc *BlockChain) collectLogs(b *types.Block, removed bool) []*types.Log {
	_ = "STUB: not implemented"
	return nil
}

// reorg takes two blocks, an old chain and a new chain and will reconstruct the
// blocks and inserts them to be part of the new canonical chain and accumulates
// potential missing transactions and post an event about them.
func (bc *BlockChain) reorg(oldHead *types.Header, newHead *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Reduce the longer chain to the same number as the shorter one

// Old chain is longer, gather all transactions and logs as deleted ones

// New chain is longer, stash all blocks away for subsequent insertion

// Both sides of the reorg are at the same number, reduce both until the common
// ancestor is found

// If the common ancestor was found, bail out

// Remove an old block as well as stash away a new block

// Step back with both chains

// If the commonBlock is less than the last accepted height, we return an error
// because performing a reorg would mean removing an accepted block from the
// canonical chain.

// Ensure the user sees large reorgs

// Reset the tx lookup cache in case to clear stale txlookups.
// This is done before writing any new chain data to avoid the
// weird scenario that canonical chain is changed while the
// stale lookups are still cached.

// Insert the new chain(except the head block(reverse order)),
// taking care of the proper incremental order.

// Insert the block in the canonical way, re-writing history

// Delete any canonical number assignments above the new head

// Use the height of [newHead] to determine which canonical hashes to remove
// in case the new chain is shorter than the old chain, in which case
// there may be hashes set on the canonical chain that were invalidated
// but not yet overwritten by the re-org.

// Send out events for logs from the old canon chain, and 'reborn'
// logs from the new canon chain. The number of logs can be very
// high, so the events are sent in batches of size around 512.

// Deleted logs + blocks:

// Also send event for blocks removed from the canon chain.

// Collect deleted logs for notification

// New logs:

type badBlock struct {
	block  *types.Block
	reason *BadBlockReason
}

type BadBlockReason struct {
	ChainConfig *params.ChainConfig `json:"chainConfig"`
	Receipts    types.Receipts      `json:"receipts"`
	Number      uint64              `json:"number"`
	Hash        common.Hash         `json:"hash"`
	Error       string              `json:"error"`
}

func (b *BadBlockReason) String() string { _ = "STUB: not implemented"; return "" }

// BadBlocks returns a list of the last 'bad blocks' that the client has seen on the network and the BadBlockReason
// that caused each to be reported as a bad block.
// BadBlocks ensures that the length of the blocks and the BadBlockReason slice have the same length.
func (bc *BlockChain) BadBlocks() ([]*types.Block, []*BadBlockReason) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addBadBlock adds a bad block to the bad-block LRU cache
func (bc *BlockChain) addBadBlock(block *types.Block, reason *BadBlockReason) {
	_ = "STUB: not implemented"
	return
}

// reportBlock logs a bad block error.
func (bc *BlockChain) reportBlock(block *types.Block, receipts types.Receipts, err error) {
	_ = "STUB: not implemented"
	return
}

// reprocessBlock reprocesses a previously accepted block. This is often used
// to regenerate previously pruned state tries.
func (bc *BlockChain) reprocessBlock(parent *types.Block, current *types.Block) (common.Hash, error) {
	_ = "STUB: not implemented"
	// Retrieve the parent block and its state to execute block
	return *new(common.Hash), nil
}

// We don't simply use [NewWithSnapshot] here because it doesn't return an
// error if [bc.snaps != nil] and [bc.snaps.Snapshot(parentRoot) == nil].

// Enable prefetching to pull in trie node paths while processing transactions

// Process previously stored block

// Validate the state using the default validator

// Commit all cached state changes into underlying memory database.

func (bc *BlockChain) commitWithSnap(
	current *types.Block, parentRoot common.Hash, statedb *state.StateDB,
) (common.Hash, error) {
	_ = "STUB: not implemented"
	// We pass through block hashes to the Update calls to ensure that we can uniquely
	// identify states despite identical state roots.
	return *new(common.Hash), nil
}

// Upstream does not perform a snapshot update if the root is the same as the
// parent root, however here the snapshots are based on the block hash, so
// this update is necessary. Note blockHashes are passed here as well.

// Because Firewood relies on tracking block hashes in a tree, we need to notify the
// database that this block is empty.

// initSnapshot instantiates a Snapshot instance and adds it to [bc]
func (bc *BlockChain) initSnapshot(b *types.Header) { _ = "STUB: not implemented"; return }

// If we are starting from genesis, generate the original snapshot disk layer
// up front, so we can use it while executing blocks in bootstrapping. This
// also avoids a costly async generation process when reaching tip.
//
// Additionally, we should always repair a snapshot if starting at genesis
// if [SnapshotLimit] > 0.

// reprocessState reprocesses the state up to [block], iterating through its ancestors until
// it reaches a block with a state committed to the database. reprocessState does not use
// snapshots since the disk layer for snapshots will most likely be above the last committed
// state that reprocessing will start from.
func (bc *BlockChain) reprocessState(current *types.Block, reexec uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// If no acceptor tip exists, treat it as empty hash (not initialized).

// The acceptor tip is up to date either if it matches the current hash, or it has not been
// initialized (i.e., this node has not accepted any blocks asynchronously).

// If the state is already available and the acceptor tip is up to date, skip re-processing.

// If the acceptorTip is a non-empty hash, jump re-processing back to the acceptor tip to ensure that
// we re-process at a minimum from the last processed accepted block.
// Note: we do not have a guarantee that the last trie on disk will be at a height <= acceptorTip.
// Since we need to re-process from at least the acceptorTip to ensure indices are updated correctly
// we must start searching for the block to start re-processing at the acceptorTip.
// This may occur if we are running in archive mode where every block's trie is committed on insertion
// or during an unclean shutdown.

// Find a historic available state root

// TODO: handle canceled context

// Try committing genesis state, and reprocess from there

// State was available at historical point, regenerate

// Note: we add 1 since in each iteration, we attempt to re-execute the next block.

// Firewood requires every root to be committed, and archival nodes
// expect every state to always be available.

// TODO: handle canceled context

// Print progress logs if long enough time elapsed

// Retrieve the next block to regenerate and process it

// Initialize snapshot if required (prevents full snapshot re-generation in
// the case of unclean shutdown)

// TODO: switch to checking the snapshot block hash markers here to ensure that when we re-process the block, we have the opportunity to apply
// a snapshot diff layer that we may have been in the middle of committing during shutdown. This will prevent snapshot re-generation in the case
// that the node stops mid-way through snapshot flattening (performed across multiple DB batches).
// If snapshot initialization is delayed due to state sync, skip initializing snaps here

// Set [writeIndices] to true, so that the indices will be updated from the last accepted tip onwards.

// Reprocess next block using previously fetched data

// Write any unsaved indices to disk

// Flatten snapshot if initialized, holding a reference to the state root until the next block
// is processed.

func (bc *BlockChain) recommitGenesis() error { _ = "STUB: not implemented"; return nil }

// clear all state, allows rebuilding genesis on top

func (bc *BlockChain) protectTrieIndex() error { _ = "STUB: not implemented"; return nil }

// populateMissingTries iterates from [bc.cacheConfig.PopulateMissingTries] (defaults to 0)
// to [LastAcceptedBlock] and persists all tries to disk that are not already on disk. This is
// used to fill trie index gaps in an "archive" node without resyncing from scratch.
//
// NOTE: Assumes the genesis root and last accepted root are written to disk
func (bc *BlockChain) populateMissingTries() error { _ = "STUB: not implemented"; return nil }

// Do not allow the config to specify a starting point above the last accepted block.

// If we are starting from the genesis, increment the start height by 1 so we don't attempt to re-process
// the genesis block.

// Print progress logs if long enough time elapsed

// TODO: handle canceled context

// Commit root to disk so that it can be accessed directly

// Write marker to DB to indicate populate missing tries finished successfully.
// Note: writing the marker here means that we do allow consecutive runs of re-populating
// missing tries if it does not finish during the prior run.

// CleanBlockRootsAboveLastAccepted gathers the blocks that may have previously been in processing above the
// last accepted block and wipes their block roots from disk to mark their tries as inaccessible.
// This is used prior to pruning to ensure that all of the tries that may still be in processing are marked
// as inaccessible and mirrors the handling of middle roots in the geth offline pruning implementation.
// This is not strictly necessary, but maintains a soft assumption.
func (bc *BlockChain) CleanBlockRootsAboveLastAccepted() error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up any block roots above the last accepted block before we start pruning.
// Note: this takes the place of middleRoots in the geth implementation since we do not
// track processing block roots via snapshot journals in the same way.

// If there is a block above the last accepted block with an identical state root, we
// explicitly remove it from the set to ensure we do not corrupt the last accepted trie.

// Delete the processing root from disk to mark the trie as inaccessible (no need to handle this in a batch).

// gatherBlockRootsAboveLastAccepted iterates forward from the last accepted block and returns a list of all block roots
// for any blocks that were inserted above the last accepted block.
// Given that we never insert a block into the chain unless all of its ancestors have been inserted, this should gather
// all of the block roots for blocks inserted above the last accepted block that may have been in processing at some point
// in the past and are therefore potentially still acceptable.
// Note: there is an edge case where the node dies while the consensus engine is rejecting a branch of blocks since the
// consensus engine will reject the lowest ancestor first. In this case, these blocks will not be considered acceptable in
// the future.
// Ex.
//
//	   A
//	 /   \
//	B     C
//	|
//	D
//	|
//	E
//	|
//	F
//
// The consensus engine accepts block C and proceeds to reject the other branch in order (B, D, E, F).
// If the consensus engine dies after rejecting block D, block D will be deleted, such that the forward iteration
// may not find any blocks at this height and will not reach the previously processing blocks E and F.
func (bc *BlockChain) gatherBlockRootsAboveLastAccepted() map[common.Hash]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// If there are no block hashes at [height], then there should be no further acceptable blocks
// past this point.

// Fetch the blocks and append their roots.

// TODO: split extras to blockchain_extra.go

// ResetToStateSyncedBlock reinitializes the state of the blockchain
// to the trie represented by [block.Root()] after updating
// in-memory and on disk current block pointers to [block].
// Only should be called after state sync has completed.
func (bc *BlockChain) ResetToStateSyncedBlock(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Update head block and snapshot pointers on disk

// if txlookup limit is 0 (uindexing disabled), we don't need to repair the tx index tail.

// Update all in-memory chain markers

// Create the state manager

// Make sure the state associated with the block is available

// CacheConfig returns a reference to [bc.cacheConfig]
//
// This is used by [miner] to set prefetch parallelism
// during block building.
func (bc *BlockChain) CacheConfig() *CacheConfig { _ = "STUB: not implemented"; return nil }

func (bc *BlockChain) repairTxIndexTail(newTail uint64) error {
	_ = "STUB: not implemented"
	return nil
}
