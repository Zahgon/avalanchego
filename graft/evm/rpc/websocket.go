// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"context"
	"net/http"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gorilla/websocket"
)

const (
	wsReadBuffer       = 1024
	wsWriteBuffer      = 1024
	wsPingInterval     = 30 * time.Second
	wsPingWriteTimeout = 5 * time.Second
	wsPongTimeout      = 30 * time.Second
	wsDefaultReadLimit = 32 * 1024 * 1024
)

var wsBufferPool = new(sync.Pool)

// WebsocketHandler returns a handler that serves JSON-RPC to WebSocket connections.
//
// allowedOrigins should be a comma-separated list of allowed origin URLs.
// To allow connections with any origin, pass "*".
func (s *Server) WebsocketHandler(allowedOrigins []string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Server) WebsocketHandlerWithDuration(allowedOrigins []string, apiMaxDuration, refillRate, maxStored time.Duration) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// wsHandshakeValidator returns a handler that verifies the origin during the
// websocket upgrade process. When a '*' is specified as an allowed origins all
// connections are accepted.
func wsHandshakeValidator(allowedOrigins []string) func(*http.Request) bool {
	_ = "STUB: not implemented"
	return nil
}

// allow localhost if no allowedOrigins are specified.

// Skip origin verification if no Origin header is present. The origin check
// is supposed to protect against browser based attacks. Browsers always set
// Origin. Non-browser software can put anything in origin and checking it doesn't
// provide additional security.

// Verify origin against allow list.

type wsHandshakeError struct {
	err    error
	status string
}

func (e wsHandshakeError) Error() string { _ = "STUB: not implemented"; return "" }

func originIsAllowed(allowedOrigins mapset.Set[string], browserOrigin string) bool {
	_ = "STUB: not implemented"
	return false
}

func ruleAllowsOrigin(allowedOrigin string, browserOrigin string) bool {
	_ = "STUB: not implemented"
	return false
}

func parseOriginURL(origin string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// DialWebsocketWithDialer creates a new RPC client using WebSocket.
//
// The context is used for the initial connection establishment. It does not
// affect subsequent interactions with the client.
//
// Deprecated: use DialOptions and the WithWebsocketDialer option.
func DialWebsocketWithDialer(ctx context.Context, endpoint, origin string, dialer websocket.Dialer) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DialWebsocket creates a new RPC client that communicates with a JSON-RPC server
// that is listening on the given endpoint.
//
// The context is used for the initial connection establishment. It does not
// affect subsequent interactions with the client.
func DialWebsocket(ctx context.Context, endpoint, origin string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newClientTransportWS(endpoint string, cfg *clientConfig) (reconnectFunc, error) {
	_ = "STUB: not implemented"
	return *new(reconnectFunc), nil
}

func wsClientHeaders(endpoint, origin string) (string, http.Header, error) {
	_ = "STUB: not implemented"
	return "", *new(http.Header), nil
}

type websocketCodec struct {
	*jsonCodec
	conn *websocket.Conn
	info PeerInfo

	wg           sync.WaitGroup
	pingReset    chan struct{}
	pongReceived chan struct{}
}

func newWebsocketCodec(conn *websocket.Conn, host string, req http.Header, readLimit int64) ServerCodec {
	_ = "STUB: not implemented"
	return *new(ServerCodec)
}

// Fill in connection details.

// Start pinger.

func (wc *websocketCodec) close() { _ = "STUB: not implemented"; return }

func (wc *websocketCodec) peerInfo() PeerInfo { _ = "STUB: not implemented"; return *new(PeerInfo) }

func (wc *websocketCodec) writeJSON(ctx context.Context, v interface{}, isError bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (wc *websocketCodec) writeJSONSkipDeadline(ctx context.Context, v interface{}, isError bool, skip bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Notify pingLoop to delay the next idle ping.

// pingLoop sends periodic ping frames when the connection is idle.
func (wc *websocketCodec) pingLoop() { _ = "STUB: not implemented"; return }
