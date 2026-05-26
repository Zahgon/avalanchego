// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gruntime

import (
	"context"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"

	pb "github.com/ava-labs/avalanchego/proto/pb/vm/runtime"
)

var _ runtime.Initializer = (*Client)(nil)

// Client is a VM runtime initializer.
type Client struct {
	client pb.RuntimeClient
}

func NewClient(client pb.RuntimeClient) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Initialize(ctx context.Context, protocolVersion uint, vmAddr string) error {
	_ = "STUB: not implemented"
	return nil
}
