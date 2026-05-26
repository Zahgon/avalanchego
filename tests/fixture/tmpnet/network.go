// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

// The Network type is defined in this file (orchestration) and
// network_config.go (reading/writing configuration).

const (
	// Constants defining the names of shell variables whose value can
	// configure network orchestration.
	RootNetworkDirEnvName = "TMPNET_ROOT_NETWORK_DIR"
	NetworkDirEnvName     = "TMPNET_NETWORK_DIR"

	// Message to log indicating where to look for metrics and logs for network
	MetricsAvailableMessage = "metrics and logs available via grafana (collectors must be running)"

	// This interval was chosen to avoid spamming node APIs during
	// startup, as smaller intervals (e.g. 50ms) seemed to noticeably
	// increase the time for a network's nodes to be seen as healthy.
	networkHealthCheckInterval = 200 * time.Millisecond

	// All temporary networks will use this arbitrary network ID by default.
	defaultNetworkID = 88888

	// eth address: 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC
	HardHatKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"

	// Default base grafana URI.
	DefaultBaseGrafanaURI = "https://avalabs.grafana.net/"

	// Default grafana URI used to construct metrics links. Can be overridden by setting GRAFANA_URI env var.
	defaultGrafanaURI = DefaultBaseGrafanaURI + "d/mabpvtq/avalanche-main-dashboard"
)

var (
	// Key expected to be funded for subnet-evm hardhat testing
	// TODO(marun) Remove when subnet-evm configures the genesis with this key.
	HardhatKey *secp256k1.PrivateKey

	errInsufficientNodes    = errors.New("at least one node is required")
	errMissingRuntimeConfig = errors.New("DefaultRuntimeConfig must not be empty")

	// Labels expected to be available in the environment when running in GitHub Actions
	githubLabels = []string{
		"gh_repo",
		"gh_workflow",
		"gh_run_id",
		"gh_run_number",
		"gh_run_attempt",
		"gh_job_id",
	}
)

func init() {
	hardhatKeyBytes, err := hex.DecodeString(HardHatKeyStr)
	if err != nil {
		panic(err)
	}
	HardhatKey, err = secp256k1.ToPrivateKey(hardhatKeyBytes)
	if err != nil {
		panic(err)
	}
}

// ConfigMap enables defining configuration in a format appropriate
// for round-tripping through JSON back to golang structs.
type ConfigMap map[string]any

// Collects the configuration for running a temporary avalanchego network
type Network struct {
	// Uniquely identifies the temporary network for metrics
	// collection. Distinct from avalanchego's concept of network ID
	// since the utility of special network ID values (e.g. to trigger
	// specific fork behavior in a given network) precludes requiring
	// unique network ID values across all temporary networks.
	UUID string

	// A string identifying the entity that started or maintains this
	// network. Useful for differentiating between networks when a
	// given CI job uses multiple networks.
	Owner string

	// Path where network configuration and data is stored
	Dir string

	// Id of the network. If zero, must be set in Genesis. Consider
	// using the GetNetworkID method if needing to retrieve the ID of
	// a running network.
	NetworkID uint32

	// Configuration common across nodes

	// Genesis for the network. If nil, NetworkID must be non-zero
	Genesis *genesis.UnparsedConfig

	// Configuration for primary subnets
	PrimarySubnetConfig ConfigMap

	// Configuration for primary network chains (P, X, C)
	PrimaryChainConfigs map[string]ConfigMap

	// Default configuration to use when creating new nodes
	DefaultFlags         FlagsMap
	DefaultRuntimeConfig NodeRuntimeConfig

	// Keys pre-funded in the genesis on both the X-Chain and the C-Chain
	PreFundedKeys []*secp256k1.PrivateKey

	// Nodes that constitute the network
	Nodes []*Node

	// Subnets that have been enabled on the network
	Subnets []*Subnet

	log logging.Logger
}

func NewDefaultNetwork(owner string) *Network { _ = "STUB: not implemented"; return nil }

// Ensure a real and absolute network dir so that node
// configuration that embeds the network path will continue to
// work regardless of symlink and working directory changes.
func toCanonicalDir(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func BootstrapNewNetwork(
	ctx context.Context,
	log logging.Logger,
	network *Network,
	rootNetworkDir string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Stops the nodes of the network configured in the provided directory.
func StopNetwork(ctx context.Context, log logging.Logger, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Restarts the nodes of the network configured in the provided directory.
func RestartNetwork(ctx context.Context, log logging.Logger, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Restart the provided nodes. Blocks on the nodes accepting API requests but not their health.
func restartNodes(ctx context.Context, nodes []*Node) error { _ = "STUB: not implemented"; return nil }

// Reads a network from the provided directory.
func ReadNetwork(ctx context.Context, log logging.Logger, dir string) (*Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initializes a new network with default configuration.
func (n *Network) EnsureDefaultConfig(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	// Populate runtime defaults before logging it
	return nil
}

// A UUID supports centralized metrics collection

// Ensure pre-funded keys if the genesis is not predefined

// Ensure primary chains are configured

// Creates the network on disk, generating its genesis and configuring its nodes in the process.
func (n *Network) Create(rootDir string) error {
	_ = "STUB: not implemented"
	// Ensure creation of the root dir
	return nil
}

// Use the default root dir

// A time-based name ensures consistent directory ordering

// Include the owner to differentiate networks created at similar times

// Ensure creation of the network dir

// Ensure the node is configured for use with the network and
// knows where to write its configuration.

// Ensure configuration on disk is current

func (n *Network) DefaultGenesis() (*genesis.UnparsedConfig, error) {
	_ = "STUB: not implemented"
	// Pre-fund known legacy keys to support ad-hoc testing. Usage of a legacy key will
	// require knowing the key beforehand rather than retrieving it from the set of pre-funded
	// keys exposed by a network. Since allocation will not be exclusive, a test using a
	// legacy key is unlikely to be a good candidate for parallel execution.
	return nil, nil
}

// Starts the specified nodes
func (n *Network) StartNodes(ctx context.Context, log logging.Logger, nodesToStart ...*Node) error {
	_ = "STUB: not implemented"
	return nil
}

// If starting all nodes except the bootstrap node (because the bootstrap node is already
// running), ensure that the health of the bootstrap node will be logged by including it in
// the set of nodes to wait for.

// Simplify output by only logging network start when starting all nodes or when starting
// the first node by itself to bootstrap subnet creation.

// Record the time before nodes are started to ensure visibility of subsequently collected metrics via the emitted link

// Provide a link to the main dashboard filtered by the uuid and showing results from now till whenever the link is viewed

// Write link to the network path

// Start the network for the first time
func (n *Network) Bootstrap(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil

	// Without the need to coordinate subnet configuration,
	// starting all nodes at once is the simplest option.
}

// The node that will be used to create subnets and bootstrap the network

// An existing sybil protection value that may need to be restored after subnet creation

// Reduce the cost of subnet creation for a network of multiple nodes by
// creating subnets with a single node with sybil protection
// disabled. This allows the creation of initial subnet state without
// requiring coordination between multiple nodes.

// If sybil protection is enabled, it should be re-enabled before the node is used to bootstrap the other nodes

// Ensure sybil protection is disabled for the bootstrap node.

// Affected node IDs can be ignored because the bootstrap node will
// always be restarted and other nodes have yet to start.

// Starts the provided node after configuring it for the network.
func (n *Network) StartNode(ctx context.Context, node *Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Attempt to stop an unhealthy node to provide some assurance to the caller
// that an error condition will not result in a lingering process.

// Stops all nodes in the network.
func (n *Network) Stop(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Ensure the node state is up-to-date
	return nil
}

// Initiate stop on all nodes

// Wait for stop to complete on all nodes

// Restarts all running nodes in the network.
func (n *Network) Restart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Waits for the provided nodes to become healthy.
func WaitForHealthyNodes(ctx context.Context, log logging.Logger, nodes []*Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensures the provided node has the configuration it needs to start. If the data dir is not
// set, it will be defaulted to [nodeParentDir]/[node ID].
func (n *Network) EnsureNodeConfig(node *Node) error {
	_ = "STUB: not implemented"
	// Ensure the node has access to network configuration
	return nil
}

// Ensure a data directory if not already set

// TrackedSubnetsForNode returns the subnet IDs for the given node
func (n *Network) TrackedSubnetsForNode(nodeID ids.NodeID) string {
	_ = "STUB: not implemented"
	return ""
}

// Subnet has not yet been created

// Only track subnets that this node validates

func (n *Network) GetSubnet(name string) *Subnet { _ = "STUB: not implemented"; return nil }

// Ensure that each subnet on the network is created. Returns the IDs of nodes whose configuration is affected by
// subnet creation so that they can be restarted if already running.
func (n *Network) CreateSubnets(ctx context.Context, log logging.Logger, apiNode *Node) (set.Set[ids.NodeID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The subnet already exists

// Allocate a pre-funded key and remove it from the network so it won't be used for
// other purposes

// Create the subnet on the network

// Persist the subnet configuration

// Ensure the pre-funded key changes are persisted to disk

// Track the set of nodes that will need to be restarted after subnet creation

// Update node configuration to track the new subnets

// Add validators for the subnet

// Collect the nodes intended to validate the subnet

// Wait for nodes to become subnet validators

// It should now be safe to create chains for the subnet

// If one or more of the subnets chains have explicit configuration, the
// subnet's validator nodes will need to be restarted for those nodes to read
// the newly written chain configuration and apply it to the chain(s).

func (n *Network) GetNode(nodeID ids.NodeID) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeURIs returns the accessible URIs of nodes in the network that are running and not ephemeral.
func (n *Network) GetNodeURIs() []NodeURI { _ = "STUB: not implemented"; return nil }

// GetAvailableNodeIDs returns the node IDs of nodes in the network that are running and not ephemeral.
func (n *Network) GetAvailableNodeIDs() []string { _ = "STUB: not implemented"; return nil }

// Retrieves bootstrap IPs and IDs for all non-ephemeral nodes except the skipped one
// (this supports collecting the bootstrap details for restarting a node).
//
// For consumption outside of avalanchego. Needs to be kept exported.
func (n *Network) GetBootstrapIPsAndIDs(skippedNode *Node) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ephemeral nodes are not guaranteed to stay running

// Node is not running

// GetNetworkID returns the effective ID of the network. If the network
// defines a genesis, the network ID in the genesis will be returned. If a
// genesis is not present (i.e. a network with a genesis included in the
// avalanchego binary - mainnet, testnet and local), the value of the
// NetworkID field will be returned
func (n *Network) GetNetworkID() uint32 { _ = "STUB: not implemented"; return 0 }

// GetGenesisFileContent returns the base64-encoded JSON-marshaled
// network genesis.
func (n *Network) GetGenesisFileContent() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetSubnetConfigContent returns the base64-encoded and
// JSON-marshaled map of subnetID to subnet configuration.
func (n *Network) GetSubnetConfigContent() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Collect configuration for non-primary subnets

// The subnet hasn't been created yet and it's not
// possible to supply configuration without an ID.

// GetChainConfigContent returns the base64-encoded and JSON-marshaled map of chain alias/ID
// to JSON-marshaled chain configuration for both primary and custom chains.
func (n *Network) GetChainConfigContent() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Collect custom chain configuration

// The chain hasn't been created yet and it's not possible to supply
// configuration without a chain ID.

// GetMonitoringLabels retrieves the map of labels and their values to be
// applied to metrics and logs collected from nodes and other collection
// targets for the network (including test workloads). Callers may need
// to set a unique value for the `instance` label to ensure a stable
// identity for the collection target.
func (n *Network) GetMonitoringLabels() map[string]string { _ = "STUB: not implemented"; return nil }

// This label must be set for compatibility with the expected
// filtering. Nodes should override this value.

// GetGitHubLabels returns a map of GitHub labels and their values if available.
func GetGitHubLabels() map[string]string { _ = "STUB: not implemented"; return nil }

// Waits until the provided nodes are healthy.
func waitForHealthy(ctx context.Context, log logging.Logger, nodes []*Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieves the root dir for tmpnet data.
func getTmpnetPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Retrieves the default root dir for storing networks and their
// configuration.
func getDefaultRootNetworkDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Retrieves the path to a reusable network path for the given owner.
func GetReusableNetworkPathForOwner(owner string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

const invalidRPCVersion = 0

// checkVMBinaries checks that VM binaries for the given subnets exist and optionally checks that VM
// binaries have the same rpcchainvm version as the indicated avalanchego binary.
func checkVMBinaries(log logging.Logger, subnets []*Subnet, config *ProcessRuntimeConfig) error {
	_ = "STUB: not implemented"
	return nil

	// Without subnets there are no VM binaries to check
}

// Check that the path exists

// Not possible to check the rpcchainvm version

// Check that the VM's rpcchainvm version matches avalanchego's version

type RPCChainVMVersion struct {
	RPCChainVM uint64 `json:"rpcchainvm"`
}

// getRPCVersion attempts to invoke the given command with the specified version arguments and
// retrieve an rpcchainvm version from its output.
func getRPCVersion(log logging.Logger, command string, versionArgs ...string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Ignore output before the opening brace to tolerate the case of a command being invoked
// with `go run` and the go toolchain emitting diagnostic logging before the version output.

// MetricsLinkForNetwork returns a link to the default metrics dashboard for the network
// with the given UUID. The start and end times are accepted as strings to support the
// use of Grafana's time range syntax (e.g. `now`, `now-1h`).
func MetricsLinkForNetwork(networkUUID string, startTime string, endTime string) string {
	_ = "STUB: not implemented"
	return ""
}

// NewGrafanaURI returns a Grafana dashboard URI.
func NewGrafanaURI(
	networkUUID string,
	startTime string,
	endTime string,
	grafanaURI string,
) string {
	_ = "STUB: not implemented"
	return ""
}
