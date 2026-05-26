// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2p

import (
	"context"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

var (
	_ validators.Connector = (*Network)(nil)
	_ common.AppHandler    = (*Network)(nil)
	_ NodeSampler          = (*PeerSampler)(nil)
	_ ConnectionHandler    = (*Peers)(nil)

	opLabel      = "op"
	handlerLabel = "handlerID"
	labelNames   = []string{opLabel, handlerLabel}
)

// ConnectionHandler handles peer connection events
type ConnectionHandler interface {
	// Connected is called when we connect to nodeID
	Connected(nodeID ids.NodeID)
	// Disconnected is called when we disconnect from nodeID
	Disconnected(nodeID ids.NodeID)
}

// NewNetwork returns an instance of Network
func NewNetwork(
	log logging.Logger,
	sender common.AppSender,
	registerer prometheus.Registerer,
	namespace string,
	connectionHandlers ...ConnectionHandler,
) (*Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Network exposes networking state and supports building p2p application
// protocols
type Network struct {
	sender             common.AppSender
	connectionHandlers []ConnectionHandler

	router *router
}

func (n *Network) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Network) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Network) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Network) AppGossip(ctx context.Context, nodeID ids.NodeID, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Network) Connected(_ context.Context, nodeID ids.NodeID, _ *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Network) Disconnected(_ context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// NewClient returns a Client that can be used to send messages for the
// corresponding protocol.
func (n *Network) NewClient(handlerID uint64, nodeSampler NodeSampler) *Client {
	_ = "STUB: not implemented"
	return nil
}

// AddHandler reserves an identifier for an application protocol
func (n *Network) AddHandler(handlerID uint64, handler Handler) error {
	_ = "STUB: not implemented"
	return nil
}

// Peers contains metadata about the current set of connected peers
type Peers struct {
	lock sync.RWMutex
	set  set.SampleableSet[ids.NodeID]
}

func (p *Peers) Connected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

func (p *Peers) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

func (p *Peers) Has(nodeID ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (p *Peers) Len() int { _ = "STUB: not implemented"; return 0 }

// Sample returns a pseudo-random sample of up to limit Peers
func (p *Peers) Sample(limit int) []ids.NodeID { _ = "STUB: not implemented"; return nil }

// PeerSampler implements NodeSampler
type PeerSampler struct {
	Peers *Peers
}

func (p PeerSampler) Sample(_ context.Context, limit int) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}

func ProtocolPrefix(handlerID uint64) []byte { _ = "STUB: not implemented"; return nil }
