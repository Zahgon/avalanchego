// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/libevm/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/semaphore"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/version"
)

// Minimum amount of time to handle a request
const (
	minRequestHandlingDuration = 100 * time.Millisecond
	maxValidatorSetStaleness   = time.Minute
)

var (
	_ Network              = (*network)(nil)
	_ validators.Connector = (*network)(nil)
	_ common.AppHandler    = (*network)(nil)

	errAcquiringSemaphore = errors.New("error acquiring semaphore")
	errEmptyNodeID        = errors.New("cannot send request to empty nodeID")
	errExpiredRequest     = errors.New("expired request")
	errNoPeersFound       = errors.New("no peers found")

	timeUntilDeadline = metrics.GetOrRegisterTimer("net_req_time_until_deadline", nil)
	droppedRequests   = metrics.GetOrRegisterCounter("net_req_deadline_dropped", nil)
)

// SyncedNetworkClient defines ability to send request / response through the Network
type SyncedNetworkClient interface {
	// SendSyncedAppRequestAny synchronously sends request to an arbitrary peer.
	// Returns response bytes, the ID of the chosen peer, and ErrRequestFailed if
	// the request should be retried.
	SendSyncedAppRequestAny(ctx context.Context, request []byte) ([]byte, ids.NodeID, error)

	// SendSyncedAppRequest synchronously sends request to the selected nodeID
	// Returns response bytes, and ErrRequestFailed if the request should be retried.
	SendSyncedAppRequest(ctx context.Context, nodeID ids.NodeID, request []byte) ([]byte, error)

	// RegisterResponse records a successful response from nodeID with the
	// observed bandwidth (response bytes divided by request time).
	RegisterResponse(nodeID ids.NodeID, bandwidth float64)

	// RegisterFailure records a failed response from nodeID.
	RegisterFailure(nodeID ids.NodeID)
}

type Network interface {
	validators.Connector
	common.AppHandler
	p2p.NodeSampler

	SyncedNetworkClient

	// SendAppRequestAny sends request to an arbitrary peer.
	// Returns the ID of the chosen peer, and an error if no peer is available.
	SendAppRequestAny(ctx context.Context, message []byte, handler message.ResponseHandler) (ids.NodeID, error)

	// SendAppRequest sends message to given nodeID, notifying handler when there's a response or timeout
	SendAppRequest(ctx context.Context, nodeID ids.NodeID, message []byte, handler message.ResponseHandler) error

	// Shutdown stops all peer channel listeners and marks the node to have stopped
	// n.Start() can be called again but the peers will have to be reconnected
	// by calling OnPeerConnected for each peer
	Shutdown()

	// SetRequestHandler sets the provided request handler as the request handler
	SetRequestHandler(handler message.RequestHandler)

	// Size returns the size of the network in number of connected peers
	Size() uint32

	P2PNetwork() *p2p.Network
	P2PValidators() *p2p.Validators
}

// network is an implementation of Network that processes message requests for
// each peer in linear fashion
type network struct {
	lock                       sync.RWMutex                       // lock for mutating state of this Network struct
	requestIDGen               uint32                             // requestID counter used to track outbound requests
	outstandingRequestHandlers map[uint32]message.ResponseHandler // maps avalanchego requestID => message.ResponseHandler
	activeAppRequests          *semaphore.Weighted                // controls maximum number of active outbound requests
	sdkNetwork                 *p2p.Network                       // SDK network (avalanchego p2p) for sending messages to peers
	appSender                  common.AppSender                   // avalanchego AppSender for sending messages
	codec                      codec.Manager                      // Codec used for parsing messages
	appRequestHandler          message.RequestHandler             // maps request type => handler
	peers                      *p2p.PeerTracker                   // tracking of peers & bandwidth

	// Set to true when Shutdown is called, after which all operations on this
	// struct are no-ops.
	//
	// Invariant: Even though `closed` is an atomic, `lock` is required to be
	// held when sending requests to guarantee that the network isn't closed
	// during these calls. This is because closing the network cancels all
	// outstanding requests, which means we must guarantee never to register a
	// request that will never be fulfilled or cancelled.
	closed utils.Atomic[bool]

	p2pValidators *p2p.Validators
}

// NewNetwork constructs a [Network] uniting the legacy synchronous message
// handling with the new [p2p.Network].
func NewNetwork(
	ctx *snow.Context,
	appSender common.AppSender,
	codec codec.Manager,
	maxActiveAppRequests int64,
	registerer prometheus.Registerer,
) (Network, error) {
	_ = "STUB: not implemented"
	return *new(Network), nil
}

// Sample returns a random sample of connected peers.
// `limit` is ignored, and one peer will be returned.
// The peer returned may not be a validator - to sample validators,
// use [p2p.Validators.Sample] instead.
func (n *network) Sample(_ context.Context, limit int) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}

// SendAppRequestAny synchronously sends request to an arbitrary peer.
// Returns the ID of the chosen peer, and an error if no peer is available.
func (n *network) SendAppRequestAny(ctx context.Context, request []byte, handler message.ResponseHandler) (ids.NodeID, error) {
	_ = "STUB: not implemented"
	// If the context was cancelled, we can skip sending this request.
	return *new(ids.NodeID), nil
}

// Take a slot from total [activeAppRequests] and block until a slot becomes available.

// SendAppRequest sends request message bytes to specified nodeID, notifying the responseHandler on response or failure
func (n *network) SendAppRequest(ctx context.Context, nodeID ids.NodeID, request []byte, responseHandler message.ResponseHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// If the context was cancelled, we can skip sending this request.

// Take a slot from total [activeAppRequests] and block until a slot becomes available.

// sendAppRequest sends request message bytes to specified nodeID and adds responseHandler to outstandingRequestHandlers
// so that it can be invoked when the network receives either a response or failure message.
// Assumes nodeID is never the local node since the peer tracker is configured with the local node in its ignoredNodes set.
// Releases active requests semaphore if there was an error in sending the request
// Returns an error if appSender is unable to make the request.
// Assumes write lock is held
func (n *network) sendAppRequest(ctx context.Context, nodeID ids.NodeID, request []byte, responseHandler message.ResponseHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Other error paths return an error so the synced caller bails before
// WaitForResult. Here we return nil, so unblock it explicitly.

// If the context was cancelled, we can skip sending this request.

// Send app request to [nodeID].
// On failure, release the slot from [activeAppRequests] and delete request
// from [outstandingRequestHandlers]
//
// Cancellation is removed from this context to avoid erroring unexpectedly.
// SendAppRequest should be non-blocking and any error other than context
// cancellation is unexpected.
//
// This guarantees that the network should never receive an unexpected
// AppResponse.

// AppRequest is called by avalanchego -> VM when there is an incoming AppRequest from a peer
// error returned by this function is expected to be treated as fatal by the engine
// returns error if the requestHandler returns an error
// sends a response back to the sender if length of response returned by the handler is >0
// expects the deadline to not have been passed
func (n *network) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Drop the request if we already missed the deadline to respond.

// We make a new context here because we don't want to cancel the context
// passed into n.AppSender.SendAppResponse below

// Return a fatal error

// Propagate fatal error

// AppResponse is invoked when there is a response received from a peer regarding a request
// Error returned by this function is expected to be treated as fatal by the engine
// If [requestID] is not known, this function will emit a log and return a nil error.
// If the response handler returns an error it is propagated as a fatal error.
func (n *network) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// We must release the slot

// AppRequestFailed can be called by the avalanchego -> VM in following cases:
// - node is benched
// - failed to send message to [nodeID] due to a network issue
// - request times out before a response is provided
// error returned by this function is expected to be treated as fatal by the engine
// returns error only when the response handler returns an error
func (n *network) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error {
	_ = "STUB: not implemented"
	return nil
}

// We must release the slot

// timeUntil calculates the time until deadline and returns an error if the deadline has passed.
// This function updates metrics for app requests.
// This is called by [AppRequest].
func timeUntil(deadline time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	// calculate how much time is left until the deadline
	return *new(time.Time), nil
}

// bufferedDeadline is half the time till actual deadline so that the message has a reasonable chance
// of completing its processing and sending the response to the peer.

// check if we have enough time to handle this request

// markRequestFulfilled fetches the handler for [requestID] and marks the request with [requestID] as having been fulfilled.
// This is called by either [AppResponse] or [AppRequestFailed].
// Assumes that the write lock is not held.
func (n *network) markRequestFulfilled(requestID uint32) (message.ResponseHandler, bool) {
	_ = "STUB: not implemented"
	return *new(message.ResponseHandler), false
}

// mark message as processed

// AppGossip is called by avalanchego -> VM when there is an incoming AppGossip
// from a peer. An error returned by this function is treated as fatal by the
// engine.
func (n *network) AppGossip(ctx context.Context, nodeID ids.NodeID, gossipBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Connected adds the given nodeID to the peer list so that it can receive messages
func (n *network) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

// [p2p.PeerTracker] filters out self via the ignoredNodes set passed at construction.

// Disconnected removes given [nodeID] from the peer list
func (n *network) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnect is idempotent. Safe to call even if Connected was filtered out.

// Shutdown marks the network as closed and fails all outstanding requests.
func (n *network) Shutdown() { _ = "STUB: not implemented"; return }

// clean up any pending requests

// make sure all waiting threads are unblocked

// mark network as closed

func (n *network) SetRequestHandler(handler message.RequestHandler) {
	_ = "STUB: not implemented"
	return
}

func (n *network) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (n *network) RegisterResponse(nodeID ids.NodeID, bandwidth float64) {
	_ = "STUB: not implemented"
	return
}

func (n *network) RegisterFailure(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// SendSyncedAppRequestAny synchronously sends request to an arbitrary peer.
// Returns response bytes, the ID of the chosen peer, and ErrRequestFailed if
// the request should be retried.
func (n *network) SendSyncedAppRequestAny(ctx context.Context, request []byte) ([]byte, ids.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.NodeID), nil
}

// SendSyncedAppRequest synchronously sends request to the specified nodeID
// Returns response bytes and ErrRequestFailed if the request should be retried.
func (n *network) SendSyncedAppRequest(ctx context.Context, nodeID ids.NodeID, request []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// P2PNetwork returns the p2p network
func (n *network) P2PNetwork() *p2p.Network { _ = "STUB: not implemented"; return nil }

// P2PValidators returns the p2p validators
func (n *network) P2PValidators() *p2p.Validators { _ = "STUB: not implemented"; return nil }

// invariant: peer/network must use explicitly even request ids.
// for this reason, [n.requestID] is initialized as zero and incremented by 2.
// This is for backwards-compatibility while the SDK router exists with the
// legacy coreth handlers to avoid a (very) narrow edge case where request ids
// can overlap, resulting in a dropped timeout.
func (n *network) nextRequestID() uint32 { _ = "STUB: not implemented"; return 0 }

// IsNetworkRequest checks if the given requestID is a request for this network handler (even-numbered requestIDs)
// SDK requests are odd-numbered requestIDs
// (see invariant: https://github.com/ava-labs/avalanchego/blob/v1.13.0/network/p2p/router.go#L83)
func IsNetworkRequest(requestID uint32) bool { _ = "STUB: not implemented"; return false }
