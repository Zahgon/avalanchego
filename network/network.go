// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"errors"
	"net"
	"net/netip"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/dialer"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/network/throttling"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/networking/sender"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/bloom"
	"github.com/ava-labs/avalanchego/utils/ips"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"

	safemath "github.com/ava-labs/avalanchego/utils/math"
)

const (
	PrimaryNetworkValidatorHealthKey = "primary network validator health"
	ConnectedPeersKey                = "connectedPeers"
	TimeSinceLastMsgReceivedKey      = "timeSinceLastMsgReceived"
	TimeSinceLastMsgSentKey          = "timeSinceLastMsgSent"
	SendFailRateKey                  = "sendFailRate"
)

var (
	_ Network = (*network)(nil)

	errNotValidator           = errors.New("node is not a validator")
	errExpectedProxy          = errors.New("expected proxy")
	errExpectedTCPProtocol    = errors.New("expected TCP protocol")
	errTrackingPrimaryNetwork = errors.New("cannot track primary network")
)

// Network defines the functionality of the networking library.
type Network interface {
	// All consensus messages can be sent through this interface. Thread safety
	// must be managed internally in the network.
	sender.ExternalSender

	// Has a health check
	health.Checker

	peer.Network

	// StartClose this network and all existing connections it has. Calling
	// StartClose multiple times is handled gracefully.
	StartClose()

	// Should only be called once, will run until either a fatal error occurs,
	// or the network is closed.
	Dispatch() error

	// Attempt to connect to this IP. The network will never stop attempting to
	// connect to this ID.
	ManuallyTrack(nodeID ids.NodeID, ip netip.AddrPort)

	// PeerInfo returns information about peers. If [nodeIDs] is empty, returns
	// info about all peers that have finished the handshake. Otherwise, returns
	// info about the peers in [nodeIDs] that have finished the handshake.
	PeerInfo(nodeIDs []ids.NodeID) []peer.Info

	// NodeUptime returns given node's primary network UptimeResults in the view of
	// this node's peer validators.
	NodeUptime() (UptimeResult, error)
}

type UptimeResult struct {
	// RewardingStakePercentage shows what percent of network stake thinks we're
	// above the uptime requirement.
	RewardingStakePercentage float64

	// WeightedAveragePercentage is the average perceived uptime of this node,
	// weighted by stake.
	// Note that this is different from RewardingStakePercentage, which shows
	// the percent of the network stake that thinks this node is above the
	// uptime requirement. WeightedAveragePercentage is weighted by uptime.
	// i.e If uptime requirement is 85 and a peer reports 40 percent it will be
	// counted (40*weight) in WeightedAveragePercentage but not in
	// RewardingStakePercentage since 40 < 85
	WeightedAveragePercentage float64
}

// To avoid potential deadlocks, we maintain that locks must be grabbed in the
// following order:
//
// 1. peersLock
// 2. manuallyTrackedIDsLock
//
// If a higher lock (e.g. manuallyTrackedIDsLock) is held when trying to grab a
// lower lock (e.g. peersLock) a deadlock could occur.
type network struct {
	config     *Config
	peerConfig *peer.Config
	metrics    *metrics

	outboundMsgThrottler throttling.OutboundMsgThrottler

	// Limits the number of connection attempts based on IP.
	inboundConnUpgradeThrottler throttling.InboundConnUpgradeThrottler
	// Listens for and accepts new inbound connections
	listener net.Listener
	// Makes new outbound connections
	dialer dialer.Dialer
	// Does TLS handshakes for inbound connections
	serverUpgrader peer.Upgrader
	// Does TLS handshakes for outbound connections
	clientUpgrader peer.Upgrader

	// ensures the close of the network only happens once.
	closeOnce sync.Once
	// Cancelled on close
	onCloseCtx context.Context
	// Call [onCloseCtxCancel] to cancel [onCloseCtx] during close()
	onCloseCtxCancel context.CancelFunc

	sendFailRateCalculator safemath.Averager

	// Tracks which peers know about which peers
	ipTracker *ipTracker
	peersLock sync.RWMutex
	// trackedIPs contains the set of IPs that we are currently attempting to
	// connect to. An entry is added to this set when we first start attempting
	// to connect to the peer. An entry is deleted from this set once we have
	// finished the handshake.
	trackedIPs      map[ids.NodeID]*trackedIP
	connectingPeers *peer.Set
	connectedPeers  *peer.Set
	closing         bool

	startupTime time.Time

	// router is notified about all peer [Connected] and [Disconnected] events
	// as well as all non-handshake peer messages.
	//
	// It is ensured that [Connected] and [Disconnected] are called in
	// consistent ways. Specifically, a peer starts in the disconnected
	// state and the network can change the peer's state from disconnected to
	// connected and back.
	//
	// It is ensured that [HandleInbound] is only called with a message from a
	// peer that is in the connected state.
	//
	// It is expected that the implementation of this interface can handle
	// concurrent calls to [Connected], [Disconnected], and [HandleInbound].
	router router.ExternalHandler
}

// NewNetwork returns a new Network implementation with the provided parameters.
func NewNetwork(
	config *Config,
	minCompatibleTime time.Time,
	msgCreator message.Creator,
	metricsRegisterer prometheus.Registerer,
	log logging.Logger,
	listener net.Listener,
	dialer dialer.Dialer,
	router router.ExternalHandler,
) (Network, error) {
	_ = "STUB: not implemented"
	return *

	// Wrap the listener to process the proxy header.
	new(Network), nil
}

// Do not perform any fuzzy matching, the header must be
// provided.

// Track all default bootstrappers to ensure their current IPs are gossiped
// like validator IPs.

// Track all recent validators to optimistically connect to them before the
// P-chain has finished syncing.

// This is set below.

func (n *network) Send(
	msg *message.OutboundMessage,
	config common.SendConfig,
	subnetID ids.ID,
	allower subnets.Allower,
) set.Set[ids.NodeID] {
	_ = "STUB: not implemented"
	return nil
}

// send to peers and update metrics
//
// Note: It is guaranteed that namedPeers and sampledPeers are disjoint.

// TODO: move send fail rate calculations into the peer metrics
// record metrics for success

// record metrics for failure

// HealthCheck returns information about several network layer health checks.
// 1) Information about health check results
// 2) An error if the health check reports unhealthy
func (n *network) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make sure we're connected to at least the minimum number of peers

// Make sure we've received an incoming message within the threshold

// Make sure we've sent an outgoing message within the threshold

// Make sure the message send failed rate isn't too high

// If we're a primary network validator, make sure we have ingress connections

// emit metrics about the lifetime of peer connections

// Network layer is healthy

func (n *network) IngressConnCount() int { _ = "STUB: not implemented"; return 0 }

// Connected is called after the peer finishes the handshake.
// Will not be called after [Disconnected] is called with this peer.
func (n *network) Connected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// AllowConnection returns true if this node should have a connection to the
// provided nodeID. If the node is attempting to connect to the minimum number
// of peers, then it should only connect if this node is a validator, or the
// peer is a validator/beacon.
func (n *network) AllowConnection(nodeID ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (n *network) Track(claimedIPPorts []*ips.ClaimedIPPort) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnected is called after the peer's handling has been shutdown.
// It is not guaranteed that [Connected] was previously called with [nodeID].
// It is guaranteed that [Connected] will not be called with [nodeID] after this
// call. Note that this is from the perspective of a single peer object, because
// a peer with the same ID can reconnect to this network instance.
func (n *network) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

func (n *network) KnownPeers() ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil,

		// There are 3 types of responses:
		//
		// - Respond with subnet IPs tracked by both ourselves and the peer
		//   - We do not consider ourself to be a primary network validator and do not track all subnets
		//
		// - Respond with all subnet IPs
		//   - The peer requests all peers
		//   - We believe ourself to be a primary network validator or we track all subnets
		//
		// - Respond with subnet IPs tracked by the peer
		//   - The peer does not request all peers
		//   - We believe ourself to be a primary network validator or we track all subnets
		//
		// The reason we allow the peer to request all peers is so that we can avoid
		// sending unnecessary data in the case that we consider them a primary network
		// validator but they do not consider themselves one.
		nil
}

func (n *network) Peers(
	peerID ids.NodeID,
	trackedSubnets set.Set[ids.ID],
	requestAllPeers bool,
	knownPeers *bloom.ReadFilter,
	salt []byte,
) []*ips.ClaimedIPPort {
	_ = "STUB: not implemented"
	return nil
}

// Only return IPs for subnets that we are tracking.

// Return IPs for all subnets.

// Dispatch starts accepting connections from other nodes attempting to connect
// to this node.
func (n *network) Dispatch() error {
	_ = "STUB: not implemented"
	// Periodically perform operations
	return nil
}

// Continuously accept new connections
// Returns error when n.Close() is called

// Sleep for a small amount of time to try to wait for the
// error to go away.

// Note: listener.Accept is rate limited outside of this package, so a
// peer can not just arbitrarily spin up goroutines here.

// Note: Calling [RemoteAddr] with the Proxy protocol enabled may
// block for up to ProxyReadHeaderTimeout. Therefore, we ensure to
// call this function inside the go-routine, rather than the main
// accept loop.

func (n *network) ManuallyTrack(nodeID ids.NodeID, ip netip.AddrPort) {
	_ = "STUB: not implemented"
	return
}

// If I'm currently connected to [nodeID] then they will have told me
// how to connect to them in the future, and I don't need to attempt to
// connect to them now.

func (n *network) track(ip *ips.ClaimedIPPort, trackAllSubnets bool) error {
	_ = "STUB: not implemented"
	// To avoid signature verification when the IP isn't needed, we
	// optimistically filter out IPs. This can result in us not tracking an IP
	// that we otherwise would have. This case can only happen if the node
	// became a validator between the time we verified the signature and when we
	// processed the IP; which should be very rare.
	//
	// Note: Avoiding signature verification when the IP isn't needed is a
	// **significant** performance optimization.
	return nil
}

// Perform all signature verification and hashing before grabbing the peer
// lock.

// If I'm currently connected to [nodeID] then I'll attempt to dial them
// when we disconnect.

// Stop tracking the old IP and start tracking the new one.

// getPeers returns a slice of connected peers from a set of [nodeIDs].
//
//   - [nodeIDs] the IDs of the peers that should be returned if they are
//     connected.
//   - [subnetID] the subnetID whose membership should be considered to
//     determine if the node is a validator.
//   - [allower] interface that determines if a node is allowed to connect to
//     the subnet based on its validator status.
func (n *network) getPeers(
	nodeIDs set.Set[ids.NodeID],
	subnetID ids.ID,
	allower subnets.Allower,
) []*peer.Peer {
	_ = "STUB: not implemented"
	return nil
}

// check if the peer is allowed to connect to the subnet

// samplePeers samples connected peers attempting to align with the number of
// requested validators, non-validators, and peers. This function will
// explicitly ignore nodeIDs already included in the send config.
func (n *network) samplePeers(
	config common.SendConfig,
	subnetID ids.ID,
	allower subnets.Allower,
) []*peer.Peer {
	_ = "STUB: not implemented"
	// As an optimization, if there are fewer validators than
	// [numValidatorsToSample], only attempt to sample [numValidatorsToSample]
	// validators to potentially avoid iterating over the entire peer set.
	return nil
}

// Only return peers that are tracking [subnetID]

// if the peer was already explicitly included, don't include in the
// sample

// check if the peer is allowed to connect to the subnet

func (n *network) disconnectedFromConnecting(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// The peer that is disconnecting from us didn't finish the handshake

func (n *network) disconnectedFromConnected(peer *peer.Peer, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

// The peer that is disconnecting from us finished the handshake

// dial will spin up a new goroutine and attempt to establish a connection with
// [nodeID] at [ip].
//
// If the connection established at [ip] doesn't match [nodeID]:
// - attempts to reach [nodeID] at [ip] will be halted.
// - the connection will be checked to see if the connection is desired or not.
//
// If [ip] has been flagged with [ip.stopTracking] then this goroutine will
// exit.
//
// If [nodeID] is marked as connecting or connected then this goroutine will
// exit.
//
// If [nodeID] is no longer marked as desired then this goroutine will exit and
// the entry in the [trackedIP]s set will be removed.
//
// If initiating a connection to [ip] fails, then dial will reattempt. However,
// there is a randomized exponential backoff to avoid spamming connection
// attempts.
func (n *network) dial(nodeID ids.NodeID, ip *trackedIP) { _ = "STUB: not implemented"; return }

// If we no longer desire a connect to nodeID, we should cleanup
// trackedIPs and this goroutine. This prevents a memory leak when
// the tracked nodeID leaves the validator set and is never able to
// be connected to.

// Typically [n.trackedIPs[nodeID]] will already equal [ip], but
// the reference to [ip] is refreshed to avoid any potential
// race conditions before removing the entry.

// While it may not be strictly needed to stop attempting to connect
// to an already connected peer here. It does prevent unnecessary
// outbound connections. Additionally, because the peer would
// immediately drop a duplicated connection, this prevents any
// "connection reset by peer" errors from interfering with the
// later duplicated connection check.

// Increase the delay that we will use for a future connection
// attempt.

// If the network is configured to disallow private IPs and the
// provided IP is private, we skip all attempts to initiate a
// connection.
//
// Invariant: We perform this check inside of the looping goroutine
// because this goroutine must clean up the trackedIPs entry if
// nodeID leaves the validator set. This is why we continue the loop
// rather than returning even though we will never initiate an
// outbound connection with this IP.

// upgrade the provided connection, which may be an inbound connection or an
// outbound connection, with the provided [upgrader].
//
// If the connection is successfully upgraded, [nil] will be returned.
//
// If the connection is desired by the node, then the resulting upgraded
// connection will be used to create a new peer. Otherwise the connection will
// be immediately closed.
func (n *network) upgrade(conn net.Conn, upgrader peer.Upgrader, isIngress bool) error {
	_ = "STUB: not implemented"
	return nil
}

// At this point we have successfully upgraded the connection and will
// return a nil error.

// peer.Start requires there is only ever one peer instance running with the
// same [peerConfig.InboundMsgThrottler]. This is guaranteed by the above
// de-duplications for [connectingPeers] and [connectedPeers].

func (n *network) PeerInfo(nodeIDs []ids.NodeID) []peer.Info { _ = "STUB: not implemented"; return nil }

func (n *network) StartClose() { _ = "STUB: not implemented"; return }

func (n *network) NodeUptime() (UptimeResult, error) {
	_ = "STUB: not implemented"
	return *new(UptimeResult), nil
}

// this is not a validator skip it.

// if this peer thinks we're above requirement add the weight

func (n *network) runTimers() { _ = "STUB: not implemented"; return }

// pullGossipPeerLists requests validators from peers in the network
func (n *network) pullGossipPeerLists() { _ = "STUB: not implemented"; return }

func (n *network) getLastReceived() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (n *network) getLastSent() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}
