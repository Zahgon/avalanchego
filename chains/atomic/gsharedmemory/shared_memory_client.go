// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsharedmemory

import (
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"

	sharedmemorypb "github.com/ava-labs/avalanchego/proto/pb/sharedmemory"
)

var _ atomic.SharedMemory = (*Client)(nil)

// Client is atomic.SharedMemory that talks over RPC.
type Client struct {
	client sharedmemorypb.SharedMemoryClient
}

// NewClient returns shared memory connected to remote shared memory
func NewClient(client sharedmemorypb.SharedMemoryClient) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Get(peerChainID ids.ID, keys [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Indexed(
	peerChainID ids.ID,
	traits [][]byte,
	startTrait,
	startKey []byte,
	limit int,
) (
	[][]byte,
	[]byte,
	[]byte,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (c *Client) Apply(requests map[ids.ID]*atomic.Requests, batches ...database.Batch) error {
	_ = "STUB: not implemented"
	return nil
}
