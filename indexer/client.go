// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

type Client struct {
	Requester rpc.EndpointRequester
}

// NewClient creates a client that can interact with an index via HTTP API
// calls.
// [uri] is the path to make API calls to.
// For example:
//   - http://1.2.3.4:9650/ext/index/C/block
//   - http://1.2.3.4:9650/ext/index/X/tx
func NewClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

// GetContainerRange returns the transactions at index [startIndex], [startIndex+1], ... , [startIndex+n-1]
// If [n] == 0, returns an empty response (i.e. null).
// If [startIndex] > the last accepted index, returns an error (unless the above apply.)
// If we run out of transactions, returns the ones fetched before running out.
func (c *Client) GetContainerRange(ctx context.Context, startIndex uint64, numToFetch int, options ...rpc.Option) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get a container by its index
func (c *Client) GetContainerByIndex(ctx context.Context, index uint64, options ...rpc.Option) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Get the most recently accepted container and its index
func (c *Client) GetLastAccepted(ctx context.Context, options ...rpc.Option) (Container, uint64, error) {
	_ = "STUB: not implemented"
	return *new(Container), 0, nil
}

// Returns 1 less than the number of containers accepted on this chain
func (c *Client) GetIndex(ctx context.Context, id ids.ID, options ...rpc.Option) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Returns true if the given container is accepted
func (c *Client) IsAccepted(ctx context.Context, id ids.ID, options ...rpc.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Get a container and its index by its ID
func (c *Client) GetContainerByID(ctx context.Context, id ids.ID, options ...rpc.Option) (Container, uint64, error) {
	_ = "STUB: not implemented"
	return *new(Container), 0, nil
}
