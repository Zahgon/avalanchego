// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2ptest

import (
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
)

func NewSelfClient(t *testing.T, ctx context.Context, nodeID ids.NodeID, handler p2p.Handler) *p2p.Client {
	_ = "STUB: not implemented"
	return nil
}

// NewClient generates a client-server pair and returns the client used to
// communicate with a server with the specified handler
func NewClient(
	t *testing.T,
	ctx context.Context,
	clientNodeID ids.NodeID,
	clientHandler p2p.Handler,
	serverNodeID ids.NodeID,
	serverHandler p2p.Handler,
) *p2p.Client {
	_ = "STUB: not implemented"
	return nil
}

// NewClientWithPeers generates a client to communicate to a set of peers
func NewClientWithPeers(
	t *testing.T,
	ctx context.Context,
	clientNodeID ids.NodeID,
	clientHandler p2p.Handler,
	peers map[ids.NodeID]p2p.Handler,
) *p2p.Client {
	_ = "STUB: not implemented"
	return nil
}

// Send the request asynchronously to avoid deadlock when the server
// sends the response back to the client

// Send the request asynchronously to avoid deadlock when the server
// sends the response back to the client

// Send the request asynchronously to avoid deadlock when the server
// sends the response back to the client
