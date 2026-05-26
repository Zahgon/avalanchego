// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2p

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
)

// Standardized identifiers for application protocol handlers
const (
	TxGossipHandlerID = iota
	AtomicTxGossipHandlerID
	// SignatureRequestHandlerID is specified in ACP-118: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/118-warp-signature-request
	SignatureRequestHandlerID
	FirewoodProofHandlerID
)

var (
	_ Handler = (*NoOpHandler)(nil)
	_ Handler = (*TestHandler)(nil)
	_ Handler = (*ValidatorHandler)(nil)

	errPeriodMustBePositive             = errors.New("period must be positive")
	errRequestsPerPeerMustBeNonNegative = errors.New("requests-per-peer must be non-negative")
)

// Handler is the server-side logic for virtual machine application protocols.
type Handler interface {
	// AppGossip is called when handling an AppGossip message.
	AppGossip(
		ctx context.Context,
		nodeID ids.NodeID,
		gossipBytes []byte,
	)
	// AppRequest is called when handling an AppRequest message.
	// Sends a response with the response corresponding to [requestBytes] or
	// an application-defined error.
	AppRequest(
		ctx context.Context,
		nodeID ids.NodeID,
		deadline time.Time,
		requestBytes []byte,
	) ([]byte, *common.AppError)
}

// NoOpHandler drops all messages
type NoOpHandler struct{}

func (NoOpHandler) AppGossip(context.Context, ids.NodeID, []byte) {
	_ = "STUB: not implemented"
	return
}

func (NoOpHandler) AppRequest(context.Context, ids.NodeID, time.Time, []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DynamicThrottlerHandler struct {
	handler         *ThrottlerHandler
	validatorSet    ValidatorSet
	requestsPerPeer float64

	throttler                  *SlidingWindowThrottler
	throttleLimitMetric        prometheus.Gauge
	lock                       sync.Mutex
	prevNumConnectedValidators int
}

func (d *DynamicThrottlerHandler) AppGossip(
	ctx context.Context,
	nodeID ids.NodeID,
	gossipBytes []byte,
) {
	_ = "STUB: not implemented"
	return
}

func (d *DynamicThrottlerHandler) AppRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	deadline time.Time,
	requestBytes []byte,
) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DynamicThrottlerHandler) checkUpdateThrottlingLimit(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// guaranteed to not overflow an int

// Throttle anything beyond 4 standard deviations which should throttle
// anything beyond the 99.994 percentile of expected requests.

func (d *DynamicThrottlerHandler) setLimit(limit float64) { _ = "STUB: not implemented"; return }

// NewDynamicThrottlerHandler wraps a handler with defaults.
// Period is the throttling evaluation period during which this node is
// expecting each peer to make requestsPerPeer requests to the network. The
// throttling limit is dynamically updated to be inversely proportional to the
// number of connected network validators.
func NewDynamicThrottlerHandler(
	log logging.Logger,
	handler Handler,
	validatorSet ValidatorSet,
	period time.Duration,
	requestsPerPeer float64,
	metrics prometheus.Registerer,
	namespace string,
) (*DynamicThrottlerHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Throttling limit will be initialized when a request is handled

func NewValidatorHandler(
	handler Handler,
	validatorSet ValidatorSet,
	log logging.Logger,
) *ValidatorHandler {
	_ = "STUB: not implemented"
	return nil
}

// ValidatorHandler drops messages from non-validators
type ValidatorHandler struct {
	handler      Handler
	validatorSet ValidatorSet
	log          logging.Logger
}

func (v ValidatorHandler) AppGossip(ctx context.Context, nodeID ids.NodeID, gossipBytes []byte) {
	_ = "STUB: not implemented"
	return
}

func (v ValidatorHandler) AppRequest(ctx context.Context, nodeID ids.NodeID, deadline time.Time, requestBytes []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// responder automatically sends the response for a given request
type responder struct {
	Handler
	handlerID uint64
	log       logging.Logger
	sender    common.AppSender
}

// AppRequest calls the underlying handler and sends back the response to nodeID
func (r *responder) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type TestHandler struct {
	AppGossipF  func(ctx context.Context, nodeID ids.NodeID, gossipBytes []byte)
	AppRequestF func(ctx context.Context, nodeID ids.NodeID, deadline time.Time, requestBytes []byte) ([]byte, *common.AppError)
}

func (t TestHandler) AppGossip(ctx context.Context, nodeID ids.NodeID, gossipBytes []byte) {
	_ = "STUB: not implemented"
	return
}

func (t TestHandler) AppRequest(ctx context.Context, nodeID ids.NodeID, deadline time.Time, requestBytes []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}
