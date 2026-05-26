// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gresponsewriter

import (
	"bufio"
	"net"
	"net/http"

	responsewriterpb "github.com/ava-labs/avalanchego/proto/pb/http/responsewriter"
)

var (
	_ http.ResponseWriter = (*Client)(nil)
	_ http.Flusher        = (*Client)(nil)
	_ http.Hijacker       = (*Client)(nil)
)

// Client is an http.ResponseWriter that talks over RPC.
type Client struct {
	client responsewriterpb.WriterClient
	header http.Header
}

// NewClient returns a response writer connected to a remote response writer
func NewClient(header http.Header, client responsewriterpb.WriterClient) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (c *Client) Write(payload []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Client) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

// TODO: Is there a way to handle the error here?

func (c *Client) Flush() {
	_ = "STUB: not implemented"
	// TODO: is there a way to handle the error here?
	return
}

type addr struct {
	network string
	str     string
}

func (a *addr) Network() string { _ = "STUB: not implemented"; return "" }

func (a *addr) String() string { _ = "STUB: not implemented"; return "" }

func (c *Client) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}
