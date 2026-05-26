// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/bloom"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/ips"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

const (
	saltSize                       = 32
	minCountEstimate               = 128
	targetFalsePositiveProbability = .001
	maxFalsePositiveProbability    = .01
	// By setting maxIPEntriesPerNode > 1, we allow nodes to update their IP at
	// least once per bloom filter reset.
	maxIPEntriesPerNode = 2

	untrackedTimestamp = -2
	olderTimestamp     = -1
	sameTimestamp      = 0
	newerTimestamp     = 1
	newTimestamp       = 2
)

var _ validators.ManagerCallbackListener = (*ipTracker)(nil)

func newIPTracker(
	trackedSubnets set.Set[ids.ID],
	log logging.Logger,
	registerer prometheus.Registerer,
	connectToAllValidators bool,
) (*ipTracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A node is tracked if any of the following conditions are met:
// - The node was manually tracked
// - The node is a validator on any subnet
type trackedNode struct {
	// manuallyTracked tracks if this node's connection was manually requested.
	manuallyTracked bool
	// validatedSubnets contains all the subnets that this node is a validator
	// of, including potentially the primary network.
	validatedSubnets set.Set[ids.ID]
	// subnets contains the subset of [subnets] that the local node also tracks,
	// including potentially the primary network.
	trackedSubnets set.Set[ids.ID]
	// ip is the most recently known IP of this node.
	ip *ips.ClaimedIPPort
}

func (n *trackedNode) wantsConnection() bool { _ = "STUB: not implemented"; return false }

func (n *trackedNode) canDelete() bool { _ = "STUB: not implemented"; return false }

type connectedNode struct {
	// trackedSubnets contains all the subnets that this node is syncing,
	// including the primary network.
	trackedSubnets set.Set[ids.ID]
	// ip this node claimed when connecting. The IP is not necessarily the same
	// IP as in the tracked map.
	ip *ips.ClaimedIPPort
}

type gossipableSubnet struct {
	numGossipableIPs prometheus.Gauge

	// manuallyGossipable contains the nodeIDs of all nodes whose IP was
	// manually configured to be gossiped for this subnet.
	manuallyGossipable set.Set[ids.NodeID]

	// gossipableIDs contains the nodeIDs of all nodes whose IP could be
	// gossiped. This is a superset of manuallyGossipable.
	gossipableIDs set.Set[ids.NodeID]

	// An IP is marked as gossipable if all of the following conditions are met:
	// - The node is a validator or was manually requested to be gossiped
	// - The node is connected
	// - The node reported that they are syncing this subnet
	// - The IP the node connected with is its latest IP
	gossipableIndices map[ids.NodeID]int
	gossipableIPs     []*ips.ClaimedIPPort
}

func (s *gossipableSubnet) setGossipableIP(ip *ips.ClaimedIPPort) {
	_ = "STUB: not implemented"
	return
}

func (s *gossipableSubnet) removeGossipableIP(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// If we aren't removing the last IP, we need to swap the last IP with the
// IP we are removing so that the slice is contiguous.

// [maxNumIPs] applies to the total number of IPs returned, including the IPs
// initially provided in [ips].
// [ips] and [nodeIDs] are extended and returned with the additional IPs added.
func (s *gossipableSubnet) getGossipableIPs(
	exceptNodeID ids.NodeID,
	exceptIPs *bloom.ReadFilter,
	salt []byte,
	maxNumIPs int,
	ips []*ips.ClaimedIPPort,
	nodeIDs set.Set[ids.NodeID],
) ([]*ips.ClaimedIPPort, set.Set[ids.NodeID]) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *gossipableSubnet) canDelete() bool { _ = "STUB: not implemented"; return false }

type ipTracker struct {
	// trackedSubnets does not include the primary network.
	trackedSubnets    set.Set[ids.ID]
	log               logging.Logger
	numTrackedPeers   prometheus.Gauge
	numGossipableIPs  prometheus.Gauge // IPs are not deduplicated across subnets
	numTrackedSubnets prometheus.Gauge
	bloomMetrics      *bloom.Metrics

	lock    sync.RWMutex
	tracked map[ids.NodeID]*trackedNode

	// The bloom filter contains the most recent tracked IPs to avoid
	// unnecessary IP gossip.
	bloom *bloom.Filter
	// To prevent validators from causing the bloom filter to have too many
	// false positives, we limit each validator to maxIPEntriesPerValidator in
	// the bloom filter.
	bloomAdditions map[ids.NodeID]int // Number of IPs added to the bloom
	bloomSalt      []byte
	maxBloomCount  int

	// Connected tracks the information of currently connected peers, including
	// tracked and untracked nodes.
	connected map[ids.NodeID]*connectedNode
	// subnet tracks all the subnets that have at least one gossipable ID.
	subnet map[ids.ID]*gossipableSubnet

	connectToAllValidators bool
}

// ManuallyTrack marks the provided nodeID as being desirable to connect to.
//
// In order for a node to learn about these nodeIDs, other nodes in the network
// must have marked them as gossipable.
//
// Even if nodes disagree on the set of manually tracked nodeIDs, they will not
// introduce persistent network gossip.
func (i *ipTracker) ManuallyTrack(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// ManuallyGossip marks the provided nodeID as being desirable to connect to and
// marks the IPs that this node provides as being valid to gossip.
//
// In order to avoid persistent network gossip, it's important for nodes in the
// network to agree upon manually gossiped nodeIDs.
func (i *ipTracker) ManuallyGossip(subnetID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

// WantsConnection returns true if any of the following conditions are met:
//  1. The node has been manually tracked.
//  2. The node has been manually gossiped on a tracked subnet.
//  3. The node is currently a validator on a tracked subnet.
//  4. The node is currently a validator on any subnet and connectToAllValidators is true.
func (i *ipTracker) WantsConnection(nodeID ids.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

// ShouldVerifyIP is used as an optimization to avoid unnecessary IP
// verification. It returns true if all of the following conditions are met:
//  1. The provided IP is from a node whose connection is desired.
//  2. This IP is newer than the most recent IP we know of for the node.
func (i *ipTracker) ShouldVerifyIP(
	ip *ips.ClaimedIPPort,
	trackAllSubnets bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// This would be the first IP
// This would be a newer IP

// AddIP attempts to update the node's IP to the provided IP. This function
// assumes the provided IP has been verified. Returns true if all of the
// following conditions are met:
//  1. The provided IP is from a node whose connection is desired on a tracked
//     subnet.
//  2. This IP is newer than the most recent IP we know of for the node.
//  3. The node is a validator and connectToAllValidators is true.
//
// If this IP is replacing a gossipable IP, this IP will also be marked as
// gossipable.
func (i *ipTracker) AddIP(ip *ips.ClaimedIPPort) bool { _ = "STUB: not implemented"; return false }

// GetIP returns the most recent IP of the provided nodeID. Returns true if all
// of the following conditions are met:
//  1. There is currently an IP for the provided nodeID.
//  2. The provided IP is from a node whose connection is desired on a tracked
//     subnet.
func (i *ipTracker) GetIP(nodeID ids.NodeID) (*ips.ClaimedIPPort, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Connected is called when a connection is established. The peer should have
// provided [ip] during the handshake.
func (i *ipTracker) Connected(ip *ips.ClaimedIPPort, trackedSubnets set.Set[ids.ID]) {
	_ = "STUB: not implemented"
	return
}

func (i *ipTracker) addIP(ip *ips.ClaimedIPPort) (int, *trackedNode) {
	_ = "STUB: not implemented"
	return 0, nil
}

// This is the first IP we've heard from the validator, so it is the
// most recent.

// This IP is older than the previously known IP.

// This IP is equal to the previously known IP.

// This IP is newer than the previously known IP.

func (i *ipTracker) setGossipableIP(ip *ips.ClaimedIPPort, trackedSubnets set.Set[ids.ID]) {
	_ = "STUB: not implemented"
	return
}

// Disconnected is called when a connection to the peer is closed.
func (i *ipTracker) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

func (i *ipTracker) OnValidatorAdded(subnetID ids.ID, nodeID ids.NodeID, _ *bls.PublicKey, _ ids.ID, _ uint64) {
	_ = "STUB: not implemented"
	return
}

// If [subnetID] is nil, the nodeID is being manually tracked.
func (i *ipTracker) addTrackableID(nodeID ids.NodeID, subnetID *ids.ID) {
	_ = "STUB: not implemented"
	return
}

// Because we previously weren't tracking this nodeID, the IP from the
// connection is guaranteed to be the most up-to-date IP that we know.

func (i *ipTracker) addGossipableID(nodeID ids.NodeID, subnetID ids.ID, manuallyGossiped bool) {
	_ = "STUB: not implemented"
	return
}

func (*ipTracker) OnValidatorWeightChanged(ids.ID, ids.NodeID, uint64, uint64) {
	_ = "STUB: not implemented"
	return
}

func (i *ipTracker) OnValidatorRemoved(subnetID ids.ID, nodeID ids.NodeID, _ uint64) {
	_ = "STUB: not implemented"
	return
}

func (i *ipTracker) updateMostRecentTrackedIP(node *trackedNode, ip *ips.ClaimedIPPort) {
	_ = "STUB: not implemented"
	return
}

// If the validator set is growing rapidly, we should increase the size of
// the bloom filter.

// ResetBloom prunes the current bloom filter. This must be called periodically
// to ensure that validators that change their IPs are updated correctly and
// that validators that left the validator set are removed.
func (i *ipTracker) ResetBloom() error { _ = "STUB: not implemented"; return nil }

// Bloom returns the binary representation of the bloom filter along with the
// random salt.
func (i *ipTracker) Bloom() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// resetBloom creates a new bloom filter with a reasonable size for the current
// validator set size. This function additionally populates the new bloom filter
// with the current most recently known IPs of validators.
func (i *ipTracker) resetBloom() error { _ = "STUB: not implemented"; return nil }

func getGossipableIPs[T any](
	i *ipTracker,
	iter map[ids.ID]T, // The values in this map aren't actually used.
	allowed func(ids.ID) bool,
	exceptNodeID ids.NodeID,
	exceptIPs *bloom.ReadFilter,
	salt []byte,
	maxNumIPs int,
) []*ips.ClaimedIPPort {
	_ = "STUB: not implemented"
	return nil
}
