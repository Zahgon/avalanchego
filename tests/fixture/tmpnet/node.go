// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"fmt"
	"net/netip"
	"time"

	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
)

// The Node type is defined in this file (node.go - orchestration) and
// node_config.go (reading/writing configuration).

const (
	defaultNodeTickerInterval = 50 * time.Millisecond
)

var (
	errMissingTLSKeyForNodeID = fmt.Errorf("failed to ensure node ID: missing value for %q", config.StakingTLSKeyContentKey)
	errMissingCertForNodeID   = fmt.Errorf("failed to ensure node ID: missing value for %q", config.StakingCertContentKey)
	errInvalidKeypair         = fmt.Errorf("%q and %q must be provided together or not at all", config.StakingTLSKeyContentKey, config.StakingCertContentKey)
)

// NodeRuntime defines the methods required to support running a node.
type NodeRuntime interface {
	readState(ctx context.Context) error
	GetAccessibleURI() string
	GetAccessibleStakingAddress(ctx context.Context) (netip.AddrPort, func(), error)
	Start(ctx context.Context) error
	InitiateStop(ctx context.Context) error
	WaitForStopped(ctx context.Context) error
	Restart(ctx context.Context) error
	IsHealthy(ctx context.Context) (bool, error)
}

// Configuration required to configure a node runtime. Only one of the fields should be set.
type NodeRuntimeConfig struct {
	Process *ProcessRuntimeConfig `json:"process,omitempty"`
	Kube    *KubeRuntimeConfig    `json:"kube,omitempty"`
}

// GetNetworkStartTimeout returns the timeout to use when starting a network.
func (c *NodeRuntimeConfig) GetNetworkStartTimeout(nodeCount int) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Processes are expected to start quickly, nodeCount is ignored

// Ensure sufficient time for scheduling and image pull

// Ensure sufficient time for the creation of autoscaled nodes

// Node supports configuring and running a node participating in a temporary network.
type Node struct {
	// Set by EnsureNodeID which is also called when the node is read.
	NodeID ids.NodeID

	// The set of flags used to start whose values are intended to deviate from the
	// default set of flags configured for the network.
	Flags FlagsMap

	// An ephemeral node is not expected to be a persistent member of the network and
	// should therefore not be used as for bootstrapping purposes.
	IsEphemeral bool

	// Optional, the configuration used to initialize the node runtime.
	// If not set, the network default will be used.
	RuntimeConfig *NodeRuntimeConfig

	// Runtime state, intended to be set by NodeRuntime
	URI            string
	StakingAddress netip.AddrPort

	// Defaults to [network dir]/[node id] if not set
	DataDir string

	// Initialized on demand
	runtime NodeRuntime

	network *Network
}

// Initializes a new node with only the data dir set
func NewNode() *Node { _ = "STUB: not implemented"; return nil }

// Initializes an ephemeral node using the provided config flags
func NewEphemeralNode(flags FlagsMap) *Node { _ = "STUB: not implemented"; return nil }

// Initializes the specified number of nodes.
func NewNodesOrPanic(count int) []*Node { _ = "STUB: not implemented"; return nil }

// Retrieves the runtime for the node.
func (n *Node) getRuntime() NodeRuntime { _ = "STUB: not implemented"; return *new(NodeRuntime) }

// Runtime configuration is validated during flag handling and network
// bootstrap so misconfiguration should be unusual.

// Retrieves the runtime configuration for the node, defaulting to the
// runtime configuration from the network if none is set for the node.
func (n *Node) getRuntimeConfig() NodeRuntimeConfig {
	_ = "STUB: not implemented"
	return *new(NodeRuntimeConfig)
}

// Runtime methods

func (n *Node) IsHealthy(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (n *Node) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *Node) InitiateStop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *Node) WaitForStopped(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *Node) Restart(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Ensure the config used to restart the node is persisted for future use
	return nil
}

func (n *Node) readState(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *Node) GetAccessibleURI() string { _ = "STUB: not implemented"; return "" }

func (n *Node) GetAccessibleStakingAddress(ctx context.Context) (netip.AddrPort, func(), error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil, nil
}

// Writes the current state of the metrics endpoint to disk
func (n *Node) SaveMetricsSnapshot(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil

	// No URI to request metrics from
}

//nolint:bodyclose // body is closed via rpc.CleanlyCloseBody

// Initiates node shutdown and waits for the node to stop.
func (n *Node) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Ensures staking and signing keys are generated if not already present and
// that the node ID (derived from the staking keypair) is set.
func (n *Node) EnsureKeys() error { _ = "STUB: not implemented"; return nil }

// Ensures a BLS signing key is generated if not already present.
func (n *Node) EnsureBLSSigningKey() error {
	_ = "STUB: not implemented"
	// Attempt to retrieve an existing key
	return nil
}

// Nothing to do

// Generate a new signing key

// Ensures a staking keypair is generated if not already present.
func (n *Node) EnsureStakingKeypair() error { _ = "STUB: not implemented"; return nil }

// Generate new keypair

// Only one of key and cert was provided

// Derives the nodes proof-of-possession. Requires the node to have a
// BLS signing key.
func (n *Node) GetProofOfPossession() (*signer.ProofOfPossession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Derives the node ID. Requires that a tls keypair is present.
func (n *Node) EnsureNodeID() error { _ = "STUB: not implemented"; return nil }

// GetUniqueID returns a globally unique identifier for the node.
func (n *Node) GetUniqueID() string { _ = "STUB: not implemented"; return "" }

// 8 characters should be enough to identify a node in the context of its network

// composeFlags determines the set of flags that should be used to
// start the node.
func (n *Node) composeFlags() (FlagsMap, error) {
	_ = "STUB: not implemented"
	return *new(FlagsMap), nil
}

// Apply the network defaults first so that they are not overridden

// Convert the network id to a string to ensure consistency in JSON round-tripping.

// Set the bootstrap configuration only for non-public networks
// Public networks should use avalanchego's built-in bootstrappers

// TODO(marun) Maybe avoid computing content flags for each node start?

// WaitForHealthy blocks until node health is true or an error (including context timeout) is observed.
func (n *Node) WaitForHealthy(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// getMonitoringLabels retrieves the map of labels and their values to be
// applied to metrics and logs collected from the node.
func (n *Node) getMonitoringLabels() map[string]string { _ = "STUB: not implemented"; return nil }

// Explicitly setting an instance label avoids the default
// behavior of using the node's URI since the URI isn't
// guaranteed stable (e.g. port may change after restart).

func (n *Node) IsRunning() bool { _ = "STUB: not implemented"; return false }
