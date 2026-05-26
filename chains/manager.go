// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"crypto"
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/api/server"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/networking/handler"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/networking/timeout"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/buffer"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/fx"
	"github.com/ava-labs/avalanchego/vms/nftfx"
	"github.com/ava-labs/avalanchego/vms/propertyfx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	timetracker "github.com/ava-labs/avalanchego/snow/networking/tracker"
)

const (
	ChainLabel = "chain"

	defaultChannelSize = 1
	initialQueueSize   = 3

	avalancheNamespace    = constants.PlatformName + metric.NamespaceSeparator + "avalanche"
	handlerNamespace      = constants.PlatformName + metric.NamespaceSeparator + "handler"
	meterchainvmNamespace = constants.PlatformName + metric.NamespaceSeparator + "meterchainvm"
	meterdagvmNamespace   = constants.PlatformName + metric.NamespaceSeparator + "meterdagvm"
	proposervmNamespace   = constants.PlatformName + metric.NamespaceSeparator + "proposervm"
	p2pNamespace          = constants.PlatformName + metric.NamespaceSeparator + "p2p"
	snowmanNamespace      = constants.PlatformName + metric.NamespaceSeparator + "snowman"
	stakeNamespace        = constants.PlatformName + metric.NamespaceSeparator + "stake"
)

var (
	// Commonly shared VM DB prefix
	VMDBPrefix = []byte("vm")

	// Bootstrapping prefixes for LinearizableVMs
	VertexDBPrefix              = []byte("vertex")
	VertexBootstrappingDBPrefix = []byte("vertex_bs")
	TxBootstrappingDBPrefix     = []byte("tx_bs")
	BlockBootstrappingDBPrefix  = []byte("interval_block_bs")

	// Bootstrapping prefixes for ChainVMs
	ChainBootstrappingDBPrefix = []byte("interval_bs")

	errUnknownVMType           = errors.New("the vm should have type avalanche.DAGVM or snowman.ChainVM")
	errCreatePlatformVM        = errors.New("attempted to create a chain running the PlatformVM")
	errNotBootstrapped         = errors.New("subnets not bootstrapped")
	errPartialSyncAsAValidator = errors.New("partial sync should not be configured for a validator")

	fxs = map[ids.ID]fx.Factory{
		secp256k1fx.ID: &secp256k1fx.Factory{},
		nftfx.ID:       &nftfx.Factory{},
		propertyfx.ID:  &propertyfx.Factory{},
	}

	_ Manager = (*manager)(nil)
)

// Manager manages the chains running on this node.
// It can:
//   - Create a chain
//   - Add a registrant. When a chain is created, each registrant calls
//     RegisterChain with the new chain as the argument.
//   - Manage the aliases of chains
type Manager interface {
	ids.Aliaser

	// Queues a chain to be created in the future after chain creator is unblocked.
	// This is only called from the P-chain thread to create other chains
	// Queued chains are created only after P-chain is bootstrapped.
	// This assumes only chains in tracked subnets are queued.
	QueueChainCreation(ChainParameters)

	// Add a registrant [r]. Every time a chain is
	// created, [r].RegisterChain([new chain]) is called.
	AddRegistrant(Registrant)

	// Given an alias, return the ID of the chain associated with that alias
	Lookup(string) (ids.ID, error)

	// Given an alias, return the ID of the VM associated with that alias
	LookupVM(string) (ids.ID, error)

	// Returns true iff the chain with the given ID exists and is finished bootstrapping
	IsBootstrapped(ids.ID) bool

	// Starts the chain creator with the initial platform chain parameters, must
	// be called once.
	StartChainCreator(platformChain ChainParameters) error

	Shutdown()
}

// ChainParameters defines the chain being created
type ChainParameters struct {
	// The ID of the chain being created.
	ID ids.ID
	// ID of the subnet that validates this chain.
	SubnetID ids.ID
	// The genesis data of this chain's ledger.
	GenesisData []byte
	// The ID of the vm this chain is running.
	VMID ids.ID
	// The IDs of the feature extensions this chain is running.
	FxIDs []ids.ID
	// Invariant: Only used when [ID] is the P-chain ID.
	CustomBeacons validators.Manager
}

type chain struct {
	Name    string
	Context *snow.ConsensusContext
	VM      common.VM
	Handler handler.Handler
}

// ChainConfig is configuration settings for the current execution.
// [Config] is the user-provided config blob for the chain.
// [Upgrade] is a chain-specific blob for coordinating upgrades.
type ChainConfig struct {
	Config  []byte
	Upgrade []byte
}

type ManagerConfig struct {
	SybilProtectionEnabled bool
	StakingTLSSigner       crypto.Signer
	StakingTLSCert         *staking.Certificate
	StakingBLSKey          bls.Signer
	TracingEnabled         bool
	// Must not be used unless [TracingEnabled] is true as this may be nil.
	Tracer                    trace.Tracer
	Log                       logging.Logger
	LogFactory                logging.Factory
	VMManager                 *vms.Manager // Manage mappings from vm ID --> vm
	BlockAcceptorGroup        snow.AcceptorGroup
	TxAcceptorGroup           snow.AcceptorGroup
	VertexAcceptorGroup       snow.AcceptorGroup
	DB                        database.Database
	MsgCreator                message.OutboundMsgBuilder // message creator, shared with network
	Router                    router.Router              // Routes incoming messages to the appropriate chain
	Net                       network.Network            // Sends consensus messages to other validators
	Validators                validators.Manager         // Validators validating on this chain
	NodeID                    ids.NodeID                 // The ID of this node
	NetworkID                 uint32                     // ID of the network this node is connected to
	PartialSyncPrimaryNetwork bool
	Server                    server.Server // Handles HTTP API calls
	AtomicMemory              *atomic.Memory
	AVAXAssetID               ids.ID
	XChainID                  ids.ID           // ID of the X-Chain,
	CChainID                  ids.ID           // ID of the C-Chain,
	CriticalChains            set.Set[ids.ID]  // Chains that can't exit gracefully
	TimeoutManager            *timeout.Manager // Manages request timeouts when sending messages to other validators
	Health                    health.Registerer
	ProposerMinBlockDelay     time.Duration
	SubnetConfigs             map[ids.ID]subnets.Config // ID -> SubnetConfig
	ChainConfigs              map[string]ChainConfig    // alias -> ChainConfig
	// ShutdownNodeFunc allows the chain manager to issue a request to shutdown the node
	ShutdownNodeFunc func(exitCode int)
	MeterVMEnabled   bool // Should each VM be wrapped with a MeterVM

	Metrics        metrics.MultiGatherer
	MeterDBMetrics metrics.MultiGatherer

	FrontierPollFrequency   time.Duration
	ConsensusAppConcurrency int

	// Max Time to spend fetching a container and its
	// ancestors when responding to a GetAncestors
	BootstrapMaxTimeGetAncestors time.Duration
	// Max number of containers in an ancestors message sent by this node.
	BootstrapAncestorsMaxContainersSent int
	// This node will only consider the first [AncestorsMaxContainersReceived]
	// containers in an ancestors message it receives.
	BootstrapAncestorsMaxContainersReceived int

	Upgrades upgrade.Config

	// Tracks CPU/disk usage caused by each peer.
	ResourceTracker timetracker.ResourceTracker

	StateSyncBeacons []ids.NodeID

	ChainDataDir string

	Subnets *Subnets
}

type manager struct {
	// Note: The string representation of a chain's ID is also considered to be an alias of the chain
	// That is, [chainID].String() is an alias for the chain, too
	ids.Aliaser
	ManagerConfig

	pChainProgress *lock.ProgressSubscription[uint64]

	// Those notified when a chain is created
	registrants []Registrant

	// queue that holds chain create requests
	chainsQueue buffer.BlockingDeque[ChainParameters]
	// unblocks chain creator to start processing the queue
	unblockChainCreatorCh chan struct{}
	// shutdown the chain creator goroutine if the queue hasn't started to be
	// processed.
	chainCreatorShutdownCh chan struct{}
	chainCreatorExited     sync.WaitGroup

	chainsLock sync.Mutex
	// Key: Chain's ID
	// Value: The chain
	chains map[ids.ID]handler.Handler

	// snowman++ related interface to allow validators retrieval
	validatorState validators.State

	avalancheGatherer    metrics.MultiGatherer            // chainID
	handlerGatherer      metrics.MultiGatherer            // chainID
	meterChainVMGatherer metrics.MultiGatherer            // chainID
	meterDAGVMGatherer   metrics.MultiGatherer            // chainID
	proposervmGatherer   metrics.MultiGatherer            // chainID
	p2pGatherer          metrics.MultiGatherer            // chainID
	snowmanGatherer      metrics.MultiGatherer            // chainID
	stakeGatherer        metrics.MultiGatherer            // chainID
	vmGatherer           map[ids.ID]metrics.MultiGatherer // vmID -> chainID
}

// New returns a new Manager
func New(config *ManagerConfig) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

// QueueChainCreation queues a chain creation request
// Invariant: Tracked Subnet must be checked before calling this function
func (m *manager) QueueChainCreation(chainParams ChainParameters) {
	_ = "STUB: not implemented"
	return
}

// createChain creates and starts the chain
//
// Note: it is expected for the subnet to already have the chain registered as
// bootstrapping before this function is called
func (m *manager) createChain(chainParams ChainParameters) { _ = "STUB: not implemented"; return }

// Note: buildChain builds all chain's relevant objects (notably engine and handler)
// but does not start their operations. Starting of the handler (which could potentially
// issue some internal messages), is delayed until chain dispatching is started and
// the chain is registered in the manager. This ensures that no message generated by handler
// upon start is dropped.

// Shut down if we fail to create a required chain (i.e. X, P or C)

// Register the health check for this chain regardless of if it was
// created or not. This attempts to notify the node operator that their
// node may not be properly validating the subnet they expect to be
// validating.

// Associate the newly created chain with its default alias

// Notify those who registered to be notified when a new chain is created

// Allows messages to be routed to the new chain. If the handler hasn't been
// started and a message is forwarded, then the message will block until the
// handler is started.

// Register bootstrapped health checks after P chain has been added to
// chains.
//
// Note: Registering this after the chain has been tracked prevents a race
//       condition between the health check and adding the first chain to
//       the manager.

// Tell the chain to start processing messages.
// If the X, P, or C Chain panics, do not attempt to recover

// Create a chain
func (m *manager) buildChain(chainParams ChainParameters, sb subnets.Subnet) (*chain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create this chain's data directory

// Create the log and context of the chain

// Get a factory for the vm we want to use on our chain

// Create the chain

// TODO: Shutdown VM if an error occurs

// Register the chain with the timeout manager

func (m *manager) AddRegistrant(r Registrant) { _ = "STUB: not implemented"; return }

// Create a DAG-based blockchain that uses Avalanche
func (m *manager) createAvalancheChain(
	ctx *snow.ConsensusContext,
	genesisData []byte,
	vdrs validators.Manager,
	vm vertex.LinearizableVMWithEngine,
	fxs []*common.Fx,
	sb subnets.Subnet,
) (*chain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Passes messages from the avalanche engines to the network

// Passes messages from the snowman engines to the network

// Handles serialization/deserialization of vertices and also the
// persistence of vertices

// The only difference between using avalancheMessageSender and
// snowmanMessageSender here is where the metrics will be placed. Because we
// end up using this sender after the linearization, we pass in
// snowmanMessageSender here.

// Initialize the ProposerVM and the vm wrapped inside it

// A default subnet configuration will be present if explicit configuration is not provided

// X-chain uses this value

// Note: this does not use [dagVM] to ensure we use the [vm]'s height index.

// Note: vmWrappingProposerVM is the VM that the Snowman engines should be
// using.

// Note: linearizableVM is the VM that the Avalanche engines should be
// using.

// sanity check

// Asynchronously passes messages from the network to the consensus engine

// Create engine, bootstrapper and state-syncer in this order,
// to make sure start callbacks are duly initialized

// create bootstrap gear

// create engine gear

// create bootstrap gear

// Register health check for this chain

// Create a linear chain using the Snowman consensus engine
func (m *manager) createSnowmanChain(
	ctx *snow.ConsensusContext,
	genesisData []byte,
	vdrs validators.Manager,
	beacons validators.Manager,
	vm block.ChainVM,
	fxs []*common.Fx,
	sb subnets.Subnet,
) (*chain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Passes messages from the consensus engine to the network

// If [m.validatorState] is nil then we are creating the P-Chain. Since the
// P-Chain is the first chain to be created, we can use it to initialize
// required interfaces for the other chains

// Wrap the validator state with a cached state so that P-chain lookups
// are cached.

// Notice that this context is left unlocked. This is because the
// lock will already be held when accessing these values on the
// P-chain.

// Initialize validatorState for future chains.

// Wrap the validator state with a cached state so that the P-chain lock
// isn't grabbed when lookups are cached.

// Set this func only for platform
//
// The snowman bootstrapper ensures this function is only executed once, so
// we don't need to be concerned about closing this channel multiple times.

// Initialize the ProposerVM and the vm wrapped inside it

// A default subnet configuration will be present if explicit configuration is not provided

// Most chains default to 0

// sanity check

// Asynchronously passes messages from the network to the consensus engine

// Create engine, bootstrapper and state-syncer in this order,
// to make sure start callbacks are duly initialized

// create bootstrap gear

// create state sync gear

// must be > 50%

// Register health checks

func (m *manager) IsBootstrapped(id ids.ID) bool { _ = "STUB: not implemented"; return false }

func (m *manager) registerBootstrappedHealthChecks() error { _ = "STUB: not implemented"; return nil }

// We should only report unhealthy if the node is partially syncing the
// primary network and is a validator.

// Note: The health check is skipped during bootstrapping to allow a
// node to sync the network even if it was previously a validator.

// Starts chain creation loop to process queued chains
func (m *manager) StartChainCreator(platformParams ChainParameters) error {
	_ = "STUB: not implemented"
	// Add the P-Chain to the Primary Network
	return nil
}

// The P-chain is created synchronously to ensure that `VM.Initialize` has
// finished before returning from this function. This is required because
// the P-chain initializes state that the rest of the node initialization
// depends on.

func (m *manager) dispatchChainCreator() { _ = "STUB: not implemented"; return }

// This channel will be closed when Shutdown is called on the manager.

// Handle chain creations

// Get the next chain we should create.
// Dequeue waits until an element is pushed, so this is not
// busy-looping.

// queue is closed, return directly

// Shutdown stops all the chains
func (m *manager) Shutdown() { _ = "STUB: not implemented"; return }

// LookupVM returns the ID of the VM associated with an alias
func (m *manager) LookupVM(alias string) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// Notify registrants [those who want to know about the creation of chains]
// that the specified chain has been created
func (m *manager) notifyRegistrants(name string, ctx *snow.ConsensusContext, vm common.VM) {
	_ = "STUB: not implemented"
	return
}

// getChainConfig returns value of a entry by looking at ID key and alias key
// it first searches ID key, then falls back to it's corresponding primary alias
func (m *manager) getChainConfig(id ids.ID) (ChainConfig, error) {
	_ = "STUB: not implemented"
	return *new(ChainConfig), nil
}

func (m *manager) getOrMakeVMGatherer(vmID ids.ID) (metrics.MultiGatherer, error) {
	_ = "STUB: not implemented"
	return *new(metrics.MultiGatherer), nil
}

func getLastAcceptedHeight(vm block.ChainVM) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
