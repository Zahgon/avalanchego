// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"os"
	"sync"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/prometheus/client_golang/prometheus"

	// Force-load precompiles to trigger registration
	_ "github.com/ava-labs/avalanchego/graft/coreth/precompile/registry"
	// Force-load tracer engine to trigger registration
	//
	// We must import this package (not referenced elsewhere) so that the native "callTracer"
	// is added to a map of client-accessible tracers. In geth, this is done
	// inside of cmd/geth.
	_ "github.com/ava-labs/libevm/eth/tracers/js"
	_ "github.com/ava-labs/libevm/eth/tracers/native"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/core/txpool"
	"github.com/ava-labs/avalanchego/graft/coreth/eth"
	"github.com/ava-labs/avalanchego/graft/coreth/eth/ethconfig"
	"github.com/ava-labs/avalanchego/graft/coreth/miner"
	"github.com/ava-labs/avalanchego/graft/coreth/network"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/config"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/graft/coreth/warp"
	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/avalanchego/graft/evm/sync/client"
	"github.com/ava-labs/avalanchego/graft/evm/sync/engine"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/profiler"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/chain"

	corethlog "github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/log"
	avalanchegossip "github.com/ava-labs/avalanchego/network/p2p/gossip"
	commonEng "github.com/ava-labs/avalanchego/snow/engine/common"
	avalancheUtils "github.com/ava-labs/avalanchego/utils"
)

var (
	_ block.ChainVM                      = (*VM)(nil)
	_ block.BuildBlockWithContextChainVM = (*VM)(nil)
	_ block.StateSyncableVM              = (*VM)(nil)
	_ client.EthBlockParser              = (*VM)(nil)
	_ engine.BlockAcceptor               = (*VM)(nil)
)

const (
	secpCacheSize          = 1024
	decidedCacheSize       = 10 * units.MiB
	missingCacheSize       = 50
	unverifiedCacheSize    = 5 * units.MiB
	bytesToIDCacheSize     = 5 * units.MiB
	warpSignatureCacheSize = 500

	// Prefixes for metrics gatherers
	ethMetricsPrefix        = "eth"
	sdkMetricsPrefix        = "sdk"
	chainStateMetricsPrefix = "chain_state"
)

// Define the API endpoints for the VM
const (
	adminEndpoint  = "/admin"
	ethRPCEndpoint = "/rpc"
	ethWSEndpoint  = "/ws"
)

var (
	// Set last accepted key to be longer than the keys used to store accepted block IDs.
	lastAcceptedKey = []byte("last_accepted_key")
	acceptedPrefix  = []byte("snowman_accepted")
	metadataPrefix  = []byte("metadata")
	warpPrefix      = []byte("warp")
	ethDBPrefix     = []byte("ethdb")
)

var (
	errInvalidBlock                               = errors.New("invalid block")
	errInvalidNonce                               = errors.New("invalid nonce")
	errUnclesUnsupported                          = errors.New("uncles unsupported")
	errNilBaseFeeApricotPhase3                    = errors.New("nil base fee is invalid after apricotPhase3")
	errNilBlockGasCostApricotPhase4               = errors.New("nil blockGasCost is invalid after apricotPhase4")
	errInvalidHeaderPredicateResults              = errors.New("invalid header predicate results")
	errInitializingLogger                         = errors.New("failed to initialize logger")
	errShuttingDownVM                             = errors.New("shutting down VM")
	errFirewoodSnapshotCacheDisabled              = errors.New("snapshot cache must be disabled for Firewood")
	errFirewoodOfflinePruningUnsupported          = errors.New("offline pruning is not supported for Firewood")
	errFirewoodMissingTrieRepopulationUnsupported = errors.New("missing trie repopulation is not supported for Firewood")
)

var originalStderr *os.File

// legacyApiNames maps pre geth v1.10.20 api names to their updated counterparts.
// used in attachEthService for backward configuration compatibility.
var legacyAPINames = map[string]string{
	"internal-public-eth":              "internal-eth",
	"internal-public-blockchain":       "internal-blockchain",
	"internal-public-transaction-pool": "internal-transaction",
	"internal-public-tx-pool":          "internal-tx-pool",
	"internal-public-debug":            "internal-debug",
	"internal-private-debug":           "internal-debug",
	"internal-public-account":          "internal-account",
	"internal-private-personal":        "internal-personal",

	"public-eth":        "eth",
	"public-eth-filter": "eth-filter",
	"private-admin":     "admin",
	"public-debug":      "debug",
	"private-debug":     "debug",
}

func init() {
	// Preserve [os.Stderr] prior to the call in plugin/main.go to plugin.Serve(...).
	// Preserving the log level allows us to update the root handler while writing to the original
	// [os.Stderr] that is being piped through to the logger via the rpcchainvm.
	originalStderr = os.Stderr
}

// VM implements the snowman.ChainVM interface
type VM struct {
	ctx *snow.Context
	// [cancel] may be nil until [snow.NormalOp] starts
	cancel context.CancelFunc
	// *chain.State helps to implement the VM interface by wrapping blocks
	// with an efficient caching layer.
	*chain.State

	config config.Config

	chainID     *big.Int
	genesisHash common.Hash
	chainConfig *params.ChainConfig
	ethConfig   ethconfig.Config

	// Extension Points
	extensionConfig *extension.Config

	// pointers to eth constructs
	eth        *eth.Ethereum
	txPool     *txpool.TxPool
	blockChain *core.BlockChain
	miner      *miner.Miner

	// [versiondb] is the VM's current versioned database
	versiondb *versiondb.Database

	// metadataDB is used to store one off keys.
	metadataDB database.Database

	// [chaindb] is the database supplied to the Ethereum backend
	chaindb ethdb.Database

	// [acceptedBlockDB] is the database to store the last accepted
	// block.
	acceptedBlockDB database.Database

	// [warpDB] is used to store warp message signatures
	// set to a prefixDB with the prefix [warpPrefix]
	warpDB database.Database

	// builderLock is used to synchronize access to the block builder,
	// as it is uninitialized at first and is only initialized when onNormalOperationsStarted is called.
	builderLock sync.Mutex
	builder     *blockBuilder

	clock *mockable.Clock

	shutdownChan chan struct{}
	shutdownWg   sync.WaitGroup

	// Continuous Profiler
	profiler profiler.ContinuousProfiler

	network.Network
	networkCodec codec.Manager

	// Metrics
	sdkMetrics *prometheus.Registry

	bootstrapped avalancheUtils.Atomic[bool]
	IsPlugin     bool

	stateSyncDone chan struct{}

	logger corethlog.Logger
	// State sync server and client
	engine.Server
	engine.Client

	// Avalanche Warp Messaging backend
	// Used to serve BLS signatures of warp messages over RPC
	warpBackend warp.Backend

	ethTxPushGossiper avalancheUtils.Atomic[*avalanchegossip.PushGossiper[*GossipEthTx]]

	chainAlias string
	// RPC handlers (should be stopped before closing chaindb)
	rpcHandlers []interface{ Stop() }
}

// Initialize implements the snowman.ChainVM interface
func (vm *VM) Initialize(
	_ context.Context,
	chainCtx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	_ []byte,
	configBytes []byte,
	_ []*commonEng.Fx,
	appSender commonEng.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create logger

// fallback to ChainID string instead of erroring

// Enable debug-level metrics that might impact runtime performance

// Initialize the database

// vm.ChainConfig() should be available for wrapping VMs before vm.initializeChain()

// must create genesis hash before [vm.ReadLastAccepted]

// Set minimum price for mining and default gas price oracle value to the min
// gas price to prevent so transactions and blocks all use the correct fees

// If we re-enable txpool journaling, we should also add the saved local
// transactions to the p2p gossip on startup.
// disable journal

// Firewood does not support iterators, so the snapshot cannot be constructed

// Create directory for offline pruning

// Initialize warp backend

// clear warpdb on initialization if config enabled

// Add p2p warp message warpHandler

func parseGenesis(ctx *snow.Context, bytes []byte) (*core.Genesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Populate the Avalanche config extras.

// If Durango is scheduled, schedule the Warp Precompile at the same time.

// Align all the Ethereum upgrades to the Avalanche upgrades

func (vm *VM) initializeMetrics() error { _ = "STUB: not implemented"; return nil }

func (vm *VM) initializeChain(lastAcceptedHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// If the gas target is specified, calculate the desired target excess and
// use it during block creation.

// Set the gas parameters for the tx pool to the minimum gas price for the
// latest upgrade.

// initializeStateSync initializes the vm for performing state sync and responding to peer requests.
// If state sync is disabled, this function will wipe any ongoing summary from
// disk to ensure that we do not continue syncing from an invalid snapshot.
func (vm *VM) initializeStateSync(lastAcceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Create standalone EVM TrieDB (read only) for serving leafs requests.
// We create a standalone TrieDB here, so that it has a standalone cache from the one
// used by the node when processing blocks.

// register default leaf request handler for state trie

// parse nodeIDs from state sync IDs in vm config

// Initialize the state sync client

// TODO clean up how this is passed around

// If StateSync is disabled, clear any ongoing summary so that we will not attempt to resume
// sync using a snapshot that has been modified by the node running normal operations.

func (vm *VM) initChainState(lastAcceptedBlock *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Register chain state metrics

func (vm *VM) SetState(_ context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

// onBootstrapStarted marks this VM as bootstrapping
func (vm *VM) onBootstrapStarted() error { _ = "STUB: not implemented"; return nil }

// After starting bootstrapping, do not attempt to resume a previous state sync.

// Ensure snapshots are initialized before bootstrapping (i.e., if state sync is skipped).
// Note calling this function has no effect if snapshots are already initialized.

// onNormalOperationsStarted marks this VM as bootstrapped
func (vm *VM) onNormalOperationsStarted() error { _ = "STUB: not implemented"; return nil }

// Initialize goroutines related to block building
// once we enter normal operation as there is no need to handle mempool gossip before this point.

// initBlockBuilding starts goroutines to manage block building
func (vm *VM) initBlockBuilding() error { _ = "STUB: not implemented"; return nil }

// NOTE: gossip network must be initialized first otherwise ETH tx gossip will not work.

func (vm *VM) WaitForEvent(ctx context.Context) (commonEng.Message, error) {
	_ = "STUB: not implemented"
	return *new(commonEng.Message), nil
}

// Block building is not initialized yet, so we haven't finished syncing or bootstrapping.

// Shutdown implements the snowman.ChainVM interface
func (vm *VM) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop RPC handlers before eth.Stop which will close the database

// buildBlock builds a block to be wrapped by ChainState
func (vm *VM) buildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) buildBlockWithContext(_ context.Context, proposerVMBlockCtx *block.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// Note: the status of block is set by ChainState

// Verify is called on a non-wrapped block here, such that this
// does not add [blk] to the processing blocks map in ChainState.
//
// TODO cache verification since Verify() will be called by the
// consensus engine as well.
//
// Note: this is only called when building a new block, so caching
// verification will only be a significant optimization for nodes
// that produce a large number of blocks.
// We call verify without writes here to avoid generating a reference
// to the blk state root in the triedb when we are going to call verify
// again from the consensus engine with writes enabled.
/*=writes*/

// parseBlock parses [b] into a block to be wrapped by ChainState.
func (vm *VM) parseBlock(_ context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// Note: the status of block is set by ChainState

// Performing syntactic verification in ParseBlock allows for
// short-circuiting bad blocks before they are processed by the VM.

func (vm *VM) ParseEthBlock(b []byte) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getBlock attempts to retrieve block [id] from the VM to be wrapped
// by ChainState.
func (vm *VM) getBlock(_ context.Context, id ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// If [ethBlock] is nil, return [database.ErrNotFound] here
// so that the miss is considered cacheable.

// Note: the status of block is set by ChainState

// GetAcceptedBlock attempts to retrieve block [blkID] from the VM. This method
// only returns accepted blocks.
func (vm *VM) GetAcceptedBlock(ctx context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// The provided block is not accepted.

// SetPreference sets what the current tail of the chain is
func (vm *VM) SetPreference(ctx context.Context, blkID ids.ID) error {
	_ = "STUB: not implemented"
	// Since each internal handler used by [vm.State] always returns a block
	// with non-nil ethBlock value, GetExtendedBlock should never return a
	// (*Block) with a nil ethBlock value.
	return nil
}

// GetBlockIDAtHeight returns the canonical block at [height].
// Note: the engine assumes that if a block is not found at [height], then
// [database.ErrNotFound] will be returned. This indicates that the VM has state
// synced and does not have all historical blocks available.
func (vm *VM) GetBlockIDAtHeight(_ context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (*VM) Version(context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CreateHandlers makes new http handlers that can handle API calls
func (vm *VM) CreateHandlers(context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*VM) NewHTTPHandler(context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (vm *VM) chainConfigExtra() *extras.ChainConfig { _ = "STUB: not implemented"; return nil }

func (vm *VM) rules(number *big.Int, time uint64) extras.Rules {
	_ = "STUB: not implemented"
	return *new(extras.Rules)
}

func (vm *VM) startContinuousProfiler() {
	_ = "STUB: not implemented"
	// If the profiler directory is empty, return immediately
	// without creating or starting a continuous profiler.
	return
}

// Wait for shutdownChan to be closed

// readLastAccepted reads the last accepted hash from [acceptedBlockDB] and returns the
// last accepted block hash and height by reading directly from [vm.chaindb] instead of relying
// on [chain].
// Note: assumes [vm.chaindb] and [vm.genesisHash] have been initialized.
func (vm *VM) ReadLastAccepted() (common.Hash, uint64, error) {
	_ = "STUB: not implemented"
	// Attempt to load last accepted block to determine if it is necessary to
	// initialize state with the genesis block.
	return *new(common.Hash), 0, nil
}

// If there is nothing in the database, return the genesis block hash and height

// attachEthService registers the backend RPC services provided by Ethereum
// to the provided handler under their assigned namespaces.
func attachEthService(handler *rpc.Server, apis []rpc.API, names []string) error {
	_ = "STUB: not implemented"
	return nil
}

// handle pre geth v1.10.20 api names as aliases for their updated values
// to allow configurations to be backwards compatible.

func (vm *VM) setPendingBlock(hash common.Hash) { _ = "STUB: not implemented"; return }

func (vm *VM) stateSyncEnabled(lastAcceptedHeight uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// if the config is set, use that

// enable state sync by default if the chain is empty.

func (vm *VM) PutLastAcceptedID(id ids.ID) error { _ = "STUB: not implemented"; return nil }
