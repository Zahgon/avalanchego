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

var (
	ErrExistingAppProtocol = errors.New("existing app protocol")
	ErrUnrequestedResponse = errors.New("unrequested response")

	_ common.AppHandler = (*router)(nil)
)

type pendingAppRequest struct {
	handlerID string
	callback  AppResponseCallback
}

type metrics struct {
	msgTime  *prometheus.GaugeVec
	msgCount *prometheus.CounterVec
}

func (m *metrics) observe(labels prometheus.Labels, start time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// router routes incoming application messages to the corresponding registered
// app handler. App messages must be made using the registered handler's
// corresponding Client.
type router struct {
	log     logging.Logger
	sender  common.AppSender
	metrics metrics

	lock               sync.RWMutex
	handlers           map[uint64]*responder
	pendingAppRequests map[uint32]pendingAppRequest
	requestID          uint32
}

// newRouter returns a new instance of Router
func newRouter(
	log logging.Logger,
	sender common.AppSender,
	metrics metrics,
) *router {
	_ = "STUB: not implemented"
	return nil
}

// invariant: sdk uses odd-numbered requestIDs

func (r *router) addHandler(handlerID uint64, handler Handler) error {
	_ = "STUB: not implemented"
	return nil
}

// AppRequest routes an AppRequest to a Handler based on the handler prefix. The
// message is dropped if no matching handler can be found.
//
// Any error condition propagated outside Handler application logic is
// considered fatal
func (r *router) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Send an error back to the requesting peer. Invalid requests that we
// cannot parse a handler id for are handled the same way as requests
// for which we do not have a registered handler.

// call the corresponding handler and send back a response to nodeID

// AppRequestFailed routes an AppRequestFailed message to the callback
// corresponding to requestID.
//
// Any error condition propagated outside Handler application logic is
// considered fatal
func (r *router) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error {
	_ = "STUB: not implemented"
	return nil
}

// we should never receive a timeout without a corresponding requestID

// AppResponse routes an AppResponse message to the callback corresponding to
// requestID.
//
// Any error condition propagated outside Handler application logic is
// considered fatal
func (r *router) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// we should never receive a timeout without a corresponding requestID

// AppGossip routes an AppGossip message to a Handler based on the handler
// prefix. The message is dropped if no matching handler can be found.
//
// Any error condition propagated outside Handler application logic is
// considered fatal
func (r *router) AppGossip(ctx context.Context, nodeID ids.NodeID, gossip []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse parses a gossip or request message and maps it to a corresponding
// handler if present.
//
// Returns:
// - The unprefixed protocol message.
// - The protocol responder.
// - The protocol metric name.
// - A boolean indicating that parsing succeeded.
//
// Invariant: Assumes [r.lock] isn't held.
func (r *router) parse(prefixedMsg []byte) ([]byte, *responder, string, bool) {
	_ = "STUB: not implemented"
	return nil, nil, "", false
}

// Invariant: Assumes [r.lock] isn't held.
func (r *router) clearAppRequest(requestID uint32) (pendingAppRequest, bool) {
	_ = "STUB: not implemented"
	return *new(pendingAppRequest), false
}

// Parse a gossip or request message.
//
// Returns:
// - The protocol ID.
// - The unprefixed protocol message.
// - A boolean indicating that parsing succeeded.
func ParseMessage(msg []byte) (uint64, []byte, bool) {
	_ = "STUB: not implemented"
	return 0, nil, false
}
