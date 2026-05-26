// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2p

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	ErrRequestPending = errors.New("request pending")
	ErrNoPeers        = errors.New("no peers")
)

// AppResponseCallback is called upon receiving an AppResponse for an AppRequest
// issued by Client.
// Callers should check [err] to see whether the AppRequest failed or not.
type AppResponseCallback func(
	ctx context.Context,
	nodeID ids.NodeID,
	responseBytes []byte,
	err error,
)

type Client struct {
	handlerIDStr  string
	handlerPrefix []byte
	router        *router
	sender        common.AppSender
	// nodeSampler is used to select nodes to route Client.AppRequestAny to
	nodeSampler NodeSampler
}

// AppRequestAny issues an AppRequest to an arbitrary node decided by Client.
// If a specific node needs to be requested, use AppRequest instead.
// See AppRequest for more docs.
func (c *Client) AppRequestAny(
	ctx context.Context,
	appRequestBytes []byte,
	onResponse AppResponseCallback,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AppRequest issues an arbitrary request to a node.
// [onResponse] is invoked upon an error or a response.
func (c *Client) AppRequest(
	ctx context.Context,
	nodeIDs set.Set[ids.NodeID],
	appRequestBytes []byte,
	onResponse AppResponseCallback,
) error {
	_ = "STUB: not implemented"
	// Cancellation is removed from this context to avoid erroring unexpectedly.
	// SendAppRequest should be non-blocking and any error other than context
	// cancellation is unexpected.
	//
	// This guarantees that the router should never receive an unexpected
	// AppResponse.
	return nil
}

// AppGossip sends a gossip message to a random set of peers.
func (c *Client) AppGossip(
	ctx context.Context,
	config common.SendConfig,
	appGossipBytes []byte,
) error {
	_ = "STUB: not implemented"
	// Cancellation is removed from this context to avoid erroring unexpectedly.
	// SendAppGossip should be non-blocking and any error other than context
	// cancellation is unexpected.
	return nil
}

// PrefixMessage prefixes the original message with the protocol identifier.
//
// Only gossip and request messages need to be prefixed.
// Response messages don't need to be prefixed because request ids are tracked
// which map to the expected response handler.
func PrefixMessage(prefix, msg []byte) []byte { _ = "STUB: not implemented"; return nil }
