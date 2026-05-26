// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow/networking/benchlist"
	"github.com/ava-labs/avalanchego/snow/networking/handler"
	"github.com/ava-labs/avalanchego/snow/networking/timeout"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/version"
)

var (
	errUnknownChain  = errors.New("received message for unknown chain")
	errUnallowedNode = errors.New("received message from non-allowed node")
	errClosing       = errors.New("router is closing")

	_ Router              = (*ChainRouter)(nil)
	_ benchlist.Benchable = (*ChainRouter)(nil)
)

type requestEntry struct {
	// When this request was registered
	time time.Time
	// The type of request that was made
	op message.Op
	// The engine type of the request that was made
	engineType p2p.EngineType

	// handled is true iff the message was already handled by the router.
	// When a peer is benched, the router immediately sends a failed response to the
	// consensus engine and marks handled as true, so that it can send and track the request/response,
	// but skip sending the response to consensus.
	handled bool
}

type peer struct {
	version *version.Application
	// The subnets that this peer is currently tracking
	trackedSubnets set.Set[ids.ID]
}

// ChainRouter routes incoming messages from the validator network
// to the consensus engines that the messages are intended for.
// Note that consensus engines are uniquely identified by the ID of the chain
// that they are working on.
// Invariant: P-chain must be registered before processing any messages
type ChainRouter struct {
	clock         mockable.Clock
	log           logging.Logger
	lock          sync.Mutex
	closing       bool
	chainHandlers map[ids.ID]handler.Handler

	// It is only safe to call [RegisterResponse] with the router lock held. Any
	// other calls to the timeout manager with the router lock held could cause
	// a deadlock because the timeout manager will call Benched and Unbenched.
	timeoutManager *timeout.Manager

	closeTimeout time.Duration
	myNodeID     ids.NodeID
	peers        map[ids.NodeID]*peer
	// node ID --> chains that node is benched on
	// invariant: if a node is benched on any chain, it is treated as disconnected on all chains
	benched                map[ids.NodeID]set.Set[ids.ID]
	criticalChains         set.Set[ids.ID]
	sybilProtectionEnabled bool
	onFatal                func(exitCode int)
	metrics                *routerMetrics
	// Parameters for doing health checks
	healthConfig HealthConfig
	// aggregator of requests based on their time
	timedRequests *linked.Hashmap[ids.RequestID, *requestEntry]
}

// Initialize the router.
//
// When this router receives an incoming message, it cancels the timeout in
// [timeouts] associated with the request that caused the incoming message, if
// applicable.
func (cr *ChainRouter) Initialize(
	nodeID ids.NodeID,
	log logging.Logger,
	timeoutManager *timeout.Manager,
	closeTimeout time.Duration,
	criticalChains set.Set[ids.ID],
	sybilProtectionEnabled bool,
	trackedSubnets set.Set[ids.ID],
	onFatal func(exitCode int),
	healthConfig HealthConfig,
	reg prometheus.Registerer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Mark myself as connected

// Register metrics

// RegisterRequest marks that we should expect to receive a reply for a request
// from the given node's [chainID] and
// the reply should have the given requestID.
//
// The type of message we expect is [op].
//
// Every registered request must be cleared either by receiving a valid reply
// and passing it to the appropriate chain or by a timeout.
// This method registers a timeout that calls such methods if we don't get a
// reply in time.
func (cr *ChainRouter) RegisterRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
	op message.Op,
	timeoutMsg *message.InboundMessage,
	engineType p2p.EngineType,
) {
	_ = "STUB: not implemented"
	return
}

// When we receive a response message type (Chits, Put, Accepted, etc.)
// we validate that we actually sent the corresponding request.
// Give this request a unique ID so we can do that validation.

// Add to the set of unfulfilled requests

// Determine whether we should include the latency of this request in our
// measurements.
// - Don't measure messages from ourself since these don't go over the
//   network.
// - Don't measure Puts because an adversary can cause us to issue a Get
//   request to them and not respond, causing a timeout, skewing latency
//   measurements.

// Register a timeout to fire if we don't get a reply in time.

// Note: timeout manager prepends this function with a notification to the benchlist manager
// of the failure (timeout).

func (cr *ChainRouter) HandleInbound(ctx context.Context, msg *message.InboundMessage) {
	_ = "STUB: not implemented"
	return
}

func (cr *ChainRouter) HandleInternal(ctx context.Context, msg *message.InboundMessage) {
	_ = "STUB: not implemented"
	// handleMessage is called in a separate goroutine because internal messages
	// may be sent while holding the chain's context lock. To enforce the
	// expected lock ordering, we must not grab the chain router lock while
	// holding the chain's context lock.
	return
}

// handleMessage routes a message to the specified chain. Messages may be
// unrequested, responses, or timeouts. The internal flag indicates whether the
// message is being sent from an internal component, such as due to a timeout,
// or if the message originated from a remote peer.
func (cr *ChainRouter) handleMessage(ctx context.Context, msg *message.InboundMessage, internal, timeout bool) {
	_ = "STUB: not implemented"
	return
}

// Get the chain, if it exists

// Note: engineType is not guaranteed to be one of the explicitly named
// enum values. If it was not specified it defaults to UNSPECIFIED.

// This was a duplicated message

// External failures and timeout-fired failures both clear the
// outstanding request.
//
// External failures (`AppError`) also remove the request from the
// timeout-manager so the timeout callback does not fire later.
// Timeout-fired failures do not need to remove the request from the
// timeout-manager because we are currently executing the callback,
// which already clears the request.
//
// Early internal failures caused by benching, disconnect, or a request
// to self are delivered once, but the request remains outstanding
// until the real response arrives or the timeout fires.

// This was a duplicated message

// Prevent duplicate handling of the request

// Pass the failure to the chain

// We didn't request this message.

// Calculate how long it took [nodeID] to reply

// Tell the timeout manager we got a response

// This message was already marked as handled internally. Skip pushing
// to the chain.

// Pass the response to the chain

// Shutdown shuts down this router
func (cr *ChainRouter) Shutdown(ctx context.Context) { _ = "STUB: not implemented"; return }

// AddChain registers the specified chain so that incoming
// messages can be routed to it
func (cr *ChainRouter) AddChain(ctx context.Context, chain handler.Handler) {
	_ = "STUB: not implemented"
	return
}

// Notify connected validators

// If this validator is benched on any chain, treat them as disconnected
// on all chains

// If this peer isn't running this chain, then we shouldn't mark them as
// connected

// Connected routes an incoming notification that a validator was just connected
func (cr *ChainRouter) Connected(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// If this validator is benched on any chain, treat them as disconnected on all chains

// TODO: fire up an event when validator state changes i.e when they leave
// set, disconnect. we cannot put an L1 validator check here since
// Disconnected would not be handled properly.
//
// When sybil protection is disabled, we only want this clause to happen
// once. Therefore, we only update the chains during the connection of the
// primary network, which is guaranteed to happen for every peer.

// If sybil protection is disabled, send a Connected message to
// every chain when connecting to the primary network.

// Disconnected routes an incoming notification that a validator was connected
func (cr *ChainRouter) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// TODO: fire up an event when validator state changes i.e when they leave
// set, disconnect. we cannot put an L1 validator check here since
// if a validator connects then it leaves validator-set, it would not be
// disconnected properly.

// Benched routes an incoming notification that a validator was benched
func (cr *ChainRouter) Benched(chainID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

// If the set already existed, then the node was previously benched.

// This will disconnect the node from all subnets when issued to P-chain.
// Even if there is no chain in the subnet.

// Unbenched routes an incoming notification that a validator was just unbenched
func (cr *ChainRouter) Unbenched(chainID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

// This node is still benched

// HealthCheck returns results of router health checks. Returns:
// 1) Information about health check results
// 2) An error if the health check reports unhealthy
func (cr *ChainRouter) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check for long running requests

// The router is not healthy

// RemoveChain removes the specified chain so that incoming
// messages can't be routed to it
func (cr *ChainRouter) removeChain(ctx context.Context, chainID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (cr *ChainRouter) clearRequest(
	op message.Op,
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) (ids.RequestID, *requestEntry) {
	_ = "STUB: not implemented"
	// Create the request ID of the request we sent that this message is (allegedly) in response to.
	return *new(ids.RequestID), nil
}

// Mark that an outstanding request has been fulfilled
