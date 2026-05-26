// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"context"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/avalanchego/ids"
)

var _ Client = (*client)(nil)

type Client interface {
	GetMessage(ctx context.Context, messageID ids.ID) ([]byte, error)
	GetMessageSignature(ctx context.Context, messageID ids.ID) ([]byte, error)
	GetMessageAggregateSignature(ctx context.Context, messageID ids.ID, quorumNum uint64, subnetIDStr string) ([]byte, error)
	GetBlockSignature(ctx context.Context, blockID ids.ID) ([]byte, error)
	GetBlockAggregateSignature(ctx context.Context, blockID ids.ID, quorumNum uint64, subnetIDStr string) ([]byte, error)
}

// client implementation for interacting with EVM [chain]
type client struct {
	client *rpc.Client
}

// NewClient returns a Client for interacting with EVM [chain]
func NewClient(uri, chain string) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (c *client) GetMessage(ctx context.Context, messageID ids.ID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) GetMessageSignature(ctx context.Context, messageID ids.ID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) GetMessageAggregateSignature(ctx context.Context, messageID ids.ID, quorumNum uint64, subnetIDStr string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) GetBlockSignature(ctx context.Context, blockID ids.ID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) GetBlockAggregateSignature(ctx context.Context, blockID ids.ID, quorumNum uint64, subnetIDStr string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
