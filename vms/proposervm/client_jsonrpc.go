// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"

	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

// JSONRPCClient for interacting with the jsonrpc API.
type JSONRPCClient struct {
	Requester rpc.EndpointRequester
}

// NewJSONRPCClient returns a Client for interacting with the jsonrpc API.
//
// The provided chain should be the chainID or an alias. Such as "P" for the
// P-Chain.
func NewJSONRPCClient(uri string, chain string) *JSONRPCClient {
	_ = "STUB: not implemented"
	return nil
}

// GetProposedHeight returns the P-chain height this node would propose in the
// next block.
func (j *JSONRPCClient) GetProposedHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetCurrentEpoch returns the current epoch information.
func (j *JSONRPCClient) GetCurrentEpoch(ctx context.Context, options ...rpc.Option) (block.Epoch, error) {
	_ = "STUB: not implemented"
	return *new(block.Epoch), nil
}
