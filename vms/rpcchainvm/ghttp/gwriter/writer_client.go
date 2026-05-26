// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gwriter

import (
	"io"

	writerpb "github.com/ava-labs/avalanchego/proto/pb/io/writer"
)

var _ io.Writer = (*Client)(nil)

// Client is an io.Writer that talks over RPC.
type Client struct{ client writerpb.WriterClient }

// NewClient returns a writer connected to a remote writer
func NewClient(client writerpb.WriterClient) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
