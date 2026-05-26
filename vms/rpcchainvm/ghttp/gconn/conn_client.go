// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gconn

import (
	"io"
	"net"
	"time"

	connpb "github.com/ava-labs/avalanchego/proto/pb/net/conn"
)

var _ net.Conn = (*Client)(nil)

// Client is an implementation of a connection that talks over RPC.
type Client struct {
	client  connpb.ConnClient
	local   net.Addr
	remote  net.Addr
	toClose []io.Closer
}

// NewClient returns a connection connected to a remote connection
func NewClient(client connpb.ConnClient, local, remote net.Addr, toClose ...io.Closer) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Client) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Client) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Client) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Client) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Client) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Client) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
