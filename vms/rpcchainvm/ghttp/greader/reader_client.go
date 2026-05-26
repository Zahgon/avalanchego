// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package greader

import (
	"io"

	readerpb "github.com/ava-labs/avalanchego/proto/pb/io/reader"
)

var _ io.Reader = (*Client)(nil)

// Client is a reader that talks over RPC.
type Client struct{ client readerpb.ReaderClient }

// NewClient returns a reader connected to a remote reader
func NewClient(client readerpb.ReaderClient) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Sentinel errors must be special-cased through an error code
