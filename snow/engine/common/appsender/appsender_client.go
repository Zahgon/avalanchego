// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package appsender

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"

	appsenderpb "github.com/ava-labs/avalanchego/proto/pb/appsender"
)

var _ common.AppSender = (*Client)(nil)

type Client struct {
	client appsenderpb.AppSenderClient
}

// NewClient returns a client that is connected to a remote AppSender.
func NewClient(client appsenderpb.AppSenderClient) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SendAppRequest(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SendAppError(ctx context.Context, nodeID ids.NodeID, requestID uint32, errorCode int32, errorMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SendAppGossip(
	ctx context.Context,
	config common.SendConfig,
	msg []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}
