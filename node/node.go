// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"crypto"
	"errors"
	"io"
	"net/netip"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/api/server"
	"github.com/ava-labs/avalanchego/chains"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/config/node"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/indexer"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/nat"
	"github.com/ava-labs/avalanchego/network"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/networking/benchlist"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/networking/timeout"
	"github.com/ava-labs/avalanchego/snow/networking/tracker"
	"github.com/ava-labs/avalanchego/snow/uptime"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/dynamicip"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/utils/profiler"
	"github.com/ava-labs/avalanchego/utils/resource"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/registry"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"
)

const (
	stakingPortName = constants.AppName + "-staking"
	httpPortName    = constants.AppName + "-http"

	ipResolutionTimeout = 30 * time.Second

	apiNamespace             = constants.PlatformName + metric.NamespaceSeparator + "api"
	benchlistNamespace       = constants.PlatformName + metric.NamespaceSeparator + "benchlist"
	dbNamespace              = constants.PlatformName + metric.NamespaceSeparator + "db"
	healthNamespace          = constants.PlatformName + metric.NamespaceSeparator + "health"
	meterDBNamespace         = constants.PlatformName + metric.NamespaceSeparator + "meterdb"
	networkNamespace         = constants.PlatformName + metric.NamespaceSeparator + "network"
	processNamespace         = constants.PlatformName + metric.NamespaceSeparator + "process"
	requestsNamespace        = constants.PlatformName + metric.NamespaceSeparator + "requests"
	resourceTrackerNamespace = constants.PlatformName + metric.NamespaceSeparator + "resource_tracker"
	responsesNamespace       = constants.PlatformName + metric.NamespaceSeparator + "responses"
	rpcchainvmNamespace      = constants.PlatformName + metric.NamespaceSeparator + "rpcchainvm"
	systemResourcesNamespace = constants.PlatformName + metric.NamespaceSeparator + "system_resources"
	upgradeNamespace         = constants.PlatformName + metric.NamespaceSeparator + "upgrade"
)

var (
	genesisHashKey     = []byte("genesisID")
	ungracefulShutdown = []byte("ungracefulShutdown")

	indexerDBPrefix = []byte{0x00}

	errInvalidTLSKey        = errors.New("invalid TLS key")
	errShuttingDown         = errors.New("server shutting down")
	errNoValidators         = errors.New("no validators in the current validator set")
	errUpgradeNeeded        = errors.New("unknown network upgrade detected")
	errUpgradeWithinTheDay  = errors.New("unknown network upgrade detected - update as soon as possible")
	errUpgradeWithinTheHour = errors.New("imminent network upgrade detected - update immediately")
)

// New returns an instance of Node
func New(
	config *node.Config,
	logFactory logging.Factory,
	logger logging.Logger,
) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Configure the bootstrappers

// Set up tracer

// Start the API Server

// Start the Metrics API

// Set up the node's database

// Initialize shared memory

// message.Creator is shared between networking, chainManager and the engine.
// It must be initiated before networking (initNetworking), chain manager (initChainManager)
// and the engine (initChains) but after the metrics (initMetricsAPI)
// message.Creator currently record metrics under network namespace

// Set up networking layer.

// Start the Health API
// Has to be initialized before chain manager
// [n.Net] must already be set

// Set up the chain manager

// Initialize the VM registry.

// Start the Admin API

// Start the Info API

// Start the Platform chain

// Node is an instance of an Avalanche node.
type Node struct {
	Log          logging.Logger
	VMFactoryLog logging.Logger
	LogFactory   logging.Factory

	// This node's unique ID used when communicating with other nodes
	// (in consensus, for example)
	ID ids.NodeID

	StakingTLSSigner crypto.Signer
	StakingTLSCert   *staking.Certificate
	StakingSigner    bls.Signer

	// Storage for this node
	DB database.Database

	router     nat.Router
	portMapper *nat.Mapper
	ipUpdater  dynamicip.Updater

	chainRouter router.Router

	// Profiles the process. Nil if continuous profiling is disabled.
	profiler profiler.ContinuousProfiler

	// Indexes blocks, transactions and blocks
	indexer indexer.Indexer

	// Manages shared memory
	sharedMemory *atomic.Memory

	// Monitors node health and runs health checks
	health health.Health

	// Build and parse messages, for both network layer and chain manager
	msgCreator message.Creator

	// Manages network timeouts
	timeoutManager *timeout.Manager

	// Manages creation of blockchains and routing messages to them
	chainManager chains.Manager

	// Manages validator benching
	benchlistManager benchlist.Manager

	uptimeCalculator uptime.LockedCalculator

	// dispatcher for events as they happen in consensus
	BlockAcceptorGroup  snow.AcceptorGroup
	TxAcceptorGroup     snow.AcceptorGroup
	VertexAcceptorGroup snow.AcceptorGroup

	// Net runs the networking stack
	Net network.Network

	// The staking address will optionally be written to a process context
	// file to enable other nodes to be configured to use this node as a
	// beacon.
	stakingAddress netip.AddrPort

	// tlsKeyLogWriterCloser is a debug file handle that writes all the TLS
	// session keys. This value should only be non-nil during debugging.
	tlsKeyLogWriterCloser io.WriteCloser

	// this node's initial connections to the network
	bootstrappers validators.Manager

	// current validators of the network
	vdrs validators.Manager

	apiURI string

	// Handles HTTP API calls
	APIServer server.Server

	// This node's configuration
	Config *node.Config

	tracer trace.Tracer

	// ensures that we only close the node once.
	shutdownOnce sync.Once

	// True if node is shutting down or is done shutting down
	shuttingDown utils.Atomic[bool]

	// Sets the exit code
	shuttingDownExitCode utils.Atomic[int]

	// Metrics Registerer
	MetricsGatherer        metrics.MultiGatherer
	MeterDBMetricsGatherer metrics.MultiGatherer

	VMAliaser ids.Aliaser
	VMManager *vms.Manager

	// VM endpoint registry
	VMRegistry registry.VMRegistry

	// Manages shutdown of a VM process
	runtimeManager runtime.Manager

	resourceManager resource.Manager

	// Tracks the CPU/disk usage caused by processing
	// messages of each peer.
	resourceTracker tracker.ResourceTracker

	// Specifies how much CPU usage each peer can cause before
	// we rate-limit them.
	cpuTargeter tracker.Targeter

	// Specifies how much disk usage each peer can cause before
	// we rate-limit them.
	diskTargeter tracker.Targeter

	// Closed when a sufficient amount of bootstrap nodes are connected to
	onSufficientlyConnected chan struct{}
}

/*
 ******************************************************************************
 *************************** P2P Networking Section ***************************
 ******************************************************************************
 */

// Initialize the networking layer.
// Assumes [n.vdrs], [n.CPUTracker], and [n.CPUTargeter] have been initialized.
func (n *Node) initNetworking(reg prometheus.Registerer) error {
	_ = "STUB: not implemented"
	// Providing either loopback address - `::1` for ipv6 and `127.0.0.1` for ipv4 - as the listen
	// host will avoid the need for a firewall exception on recent MacOS:
	//
	//   - MacOS requires a manually-approved firewall exception [1] for each version of a given
	//   binary that wants to bind to all interfaces (i.e. with an address of `:[port]`). Each
	//   compiled version of avalanchego requires a separate exception to be allowed to bind to all
	//   interfaces.
	//
	//   - A firewall exception is not required to bind to a loopback interface, but the only way for
	//   Listen() to bind to loopback for both ipv4 and ipv6 is to bind to all interfaces [2] which
	//   requires an exception.
	//
	//   - Thus, the only way to start a node on MacOS without approving a firewall exception for the
	//   avalanchego binary is to bind to loopback by specifying the host to be `::1` or `127.0.0.1`.
	//
	// 1: https://apple.stackexchange.com/questions/393715/do-you-want-the-application-main-to-accept-incoming-network-connections-pop
	// 2: https://github.com/golang/go/issues/56998
	return nil
}

// Wrap listener so it will only accept a certain number of incoming connections per second

// Record the bound address to enable inclusion in process context file.

// Use the specified public IP.

// Use dynamic IP resolution.

// Use that to resolve our public IP.

// Regularly update our public IP and port mappings.

// We allow nodes to gossip unknown ACPs in case the current ACPs constant
// becomes out of date.

// Create chain router

// Configure benchlist

// Sybil protection is disabled so we don't have a txID that added us as
// a validator. Because each validator needs a txID associated with it,
// we hack one together by just padding our nodeID with zeroes.

// add node configs to network config

// Write process context to the configured path. Supports the use of
// dynamically chosen network ports with local network orchestration.
func (n *Node) writeProcessContext() error { _ = "STUB: not implemented"; return nil }

// Write the process context to disk

// Set by network initialization

// Dispatch starts the node's servers.
// Returns when the node exits.
func (n *Node) Dispatch() error { _ = "STUB: not implemented"; return nil }

// Start the HTTP API server

// When [n].Shutdown() is called, [n.APIServer].Close() is called.
// This causes [n.APIServer].Dispatch() to return an error.
// If that happened, don't log/return an error here.

// If the API server isn't running, shut down the node.
// If node is already shutting down, this does not trigger shutdown again,
// and blocks until Shutdown returns.

// Log a warning if we aren't able to connect to a sufficient portion of
// nodes.

// Add state sync nodes to the peer network

// Add bootstrap nodes to the peer network

// Start P2P connections

// If the P2P server isn't running, shut down the node.
// If node is already shutting down, this does not trigger shutdown again,
// and blocks until Shutdown returns.

// Remove the process context file to communicate to an orchestrator
// that the node is no longer running.

/*
 ******************************************************************************
 *********************** End P2P Networking Section ***************************
 ******************************************************************************
 */

func (n *Node) initDatabase() error { _ = "STUB: not implemented"; return nil }

// Prior to v1.10.15, the only on-disk database was leveldb, and its
// files went to [dbPath]/[networkID]/v1.4.5.

// dbFolderName is appended to the database path given in the config

// Set the node IDs of the peers this node should first connect to
func (n *Node) initBootstrappers() error { _ = "STUB: not implemented"; return nil }

// Note: The beacon connection manager will treat all beaconIDs as
//       equal.
// Invariant: We never use the TxID or BLS keys populated here.

// Create the EventDispatcher used for hooking events
// into the general process flow.
func (n *Node) initEventDispatchers() { _ = "STUB: not implemented"; return }

// Initialize [n.indexer].
// Should only be called after [n.DB], [n.DecisionAcceptorGroup],
// [n.ConsensusAcceptorGroup], [n.Log], [n.APIServer], [n.chainManager] are
// initialized
func (n *Node) initIndexer() error { _ = "STUB: not implemented"; return nil }

// TODO put exit code here

// Chain manager will notify indexer when a chain is created

// Initializes the Platform chain.
// Its genesis data specifies the other chains that should be created.
func (n *Node) initChains(genesisBytes []byte) error { _ = "STUB: not implemented"; return nil }

// Specifies other chains to create

// Start the chain creator with the Platform Chain

func (n *Node) initMetrics() error { _ = "STUB: not implemented"; return nil }

func (n *Node) initNAT() { _ = "STUB: not implemented"; return }

// initAPIServer initializes the server that handles HTTP calls
func (n *Node) initAPIServer() error { _ = "STUB: not implemented"; return nil }

// An empty host is treated as a wildcard to match all addresses, so it is
// considered public.

// Don't open the HTTP port if the HTTP server is private

// Add the default VM aliases
func (n *Node) addDefaultVMAliases() error { _ = "STUB: not implemented"; return nil }

// Create the chainManager and register the following VMs:
// AVM, Simple Payments DAG, Simple Payments Chain, and Platform VM
// Assumes n.DBManager, n.vdrs all initialized (non-nil)
func (n *Node) initChainManager(avaxAssetID ids.ID) error { _ = "STUB: not implemented"; return nil }

// If any of these chains die, the node shuts down

// Routes incoming messages from peers to the appropriate chain

// Notify the API server when new chains are created

// initVMs initializes the VMs Avalanche supports + any additional vms installed as plugins.
func (n *Node) initVMs() error { _ = "STUB: not implemented"; return nil }

// If sybil protection is disabled, we provide the P-chain its own local
// validator manager that will not be used by the rest of the node. This
// allows the node's validator sets to be determined by network connections.

// Register the VMs that Avalanche supports

// initialize vm runtime manager

// initialize the vm registry

// register any vms that need to be installed as plugins from disk

// initSharedMemory initializes the shared memory for cross chain interaction
func (n *Node) initSharedMemory() { _ = "STUB: not implemented"; return }

// initMetricsAPI initializes the Metrics API
// Assumes n.APIServer is already set
func (n *Node) initMetricsAPI() error { _ = "STUB: not implemented"; return nil }

// Current state of process metrics.

// Go process metrics using debug.GCStats.

// initAdminAPI initializes the Admin API service
// Assumes n.log, n.chainManager, and n.ValidatorAPI already initialized
func (n *Node) initAdminAPI() error { _ = "STUB: not implemented"; return nil }

// initProfiler initializes the continuous profiling
func (n *Node) initProfiler() { _ = "STUB: not implemented"; return }

func (n *Node) initInfoAPI() error { _ = "STUB: not implemented"; return nil }

// initHealthAPI initializes the Health API service
// Assumes n.Log, n.Net, n.APIServer, n.HTTPLog already initialized
func (n *Node) initHealthAPI() error { _ = "STUB: not implemented"; return nil }

// TODO: add database health to liveness check

// confirm that the node has enough disk space to continue operating
// if there is too little disk space remaining, first report unhealthy and then shutdown the node

// TODO: This healthcheck calls both n.vdrs.GetMap and n.Net.PeerInfo which
// are expensive calls. This could be rewritten as an event based monitor to
// avoid expensive iteration.

// upgrade time -> stake weight

// log at the rate of the health check

// Give chains aliases as specified by the genesis information
func (n *Node) initChainAliases(genesisBytes []byte) error { _ = "STUB: not implemented"; return nil }

// APIs aliases as specified by the genesis information
func (n *Node) initAPIAliases(genesisBytes []byte) error { _ = "STUB: not implemented"; return nil }

// Initialize [n.resourceManager].
func (n *Node) initResourceManager() error { _ = "STUB: not implemented"; return nil }

// Initialize [n.cpuTargeter].
// Assumes [n.resourceTracker] is already initialized.
func (n *Node) initCPUTargeter(
	config *tracker.TargeterConfig,
) {
	_ = "STUB: not implemented"
	return
}

// Initialize [n.diskTargeter].
// Assumes [n.resourceTracker] is already initialized.
func (n *Node) initDiskTargeter(
	config *tracker.TargeterConfig,
) {
	_ = "STUB: not implemented"
	return
}

// newStakingSigner returns a BLS signer based on the provided validated configuration.
func newStakingSigner(cfg node.StakingSignerConfig) (bls.Signer, error) {
	_ = "STUB: not implemented"
	return *new(bls.Signer), nil
}

// Shutdown this node
// May be called multiple times
// All calls to shutdownOnce.Do block until the first call returns
func (n *Node) Shutdown(exitCode int) { _ = "STUB: not implemented"; return }

// only set the exit code once

func (n *Node) shutdown() { _ = "STUB: not implemented"; return }

// Passes if the node is not shutting down

// Ensure all runtimes are shutdown

func (n *Node) ExitCode() int { _ = "STUB: not implemented"; return 0 }
