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
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	defaultBodyLimit = 5 * 1024 * 1024
	contentType      = "application/json"
)

// https://www.jsonrpc.org/historical/json-rpc-over-http.html#id13
var acceptedContentTypes = []string{contentType, "application/json-rpc", "application/jsonrequest"}

type httpConn struct {
	client    *http.Client
	url       string
	closeOnce sync.Once
	closeCh   chan interface{}
	mu        sync.Mutex // protects headers
	headers   http.Header
	auth      HTTPAuth
}

// httpConn implements ServerCodec, but it is treated specially by Client
// and some methods don't work. The panic() stubs here exist to ensure
// this special treatment is correct.

func (hc *httpConn) writeJSON(ctx context.Context, val interface{}, isError bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *httpConn) writeJSONSkipDeadline(context.Context, interface{}, bool, bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *httpConn) peerInfo() PeerInfo { _ = "STUB: not implemented"; return *new(PeerInfo) }

func (hc *httpConn) remoteAddr() string { _ = "STUB: not implemented"; return "" }

func (hc *httpConn) readBatch() ([]*jsonrpcMessage, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (hc *httpConn) close() { _ = "STUB: not implemented"; return }

func (hc *httpConn) closed() <-chan interface{} {
	_ = "STUB: not implemented"

	// HTTPTimeouts represents the configuration params for the HTTP RPC server.
	return nil
}

type HTTPTimeouts struct {
	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	ReadHeaderTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, ReadHeaderTimeout is used.
	IdleTimeout time.Duration
}

// DefaultHTTPTimeouts represents the default timeout values used if further
// configuration is not provided.
var DefaultHTTPTimeouts = HTTPTimeouts{
	ReadTimeout:       30 * time.Second,
	ReadHeaderTimeout: 30 * time.Second,
	WriteTimeout:      30 * time.Second,
	IdleTimeout:       120 * time.Second,
}

// DialHTTP creates a new RPC client that connects to an RPC server over HTTP.
func DialHTTP(endpoint string) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// DialHTTPWithClient creates a new RPC client that connects to an RPC server over HTTP
// using the provided HTTP Client.
//
// Deprecated: use DialOptions and the WithHTTPClient option.
func DialHTTPWithClient(endpoint string, client *http.Client) (*Client, error) {
	_ = "STUB: not implemented"
	// Sanity check URL so we don't end up with a client that will fail every request.
	return nil, nil
}

func newClientTransportHTTP(endpoint string, cfg *clientConfig) reconnectFunc {
	_ = "STUB: not implemented"
	return *new(reconnectFunc)
}

// cleanlyCloseBody avoids sending unnecessary RST_STREAM and PING frames by
// ensuring the whole body is read before being closed.
// See https://blog.cloudflare.com/go-and-enhance-your-calm/#reading-bodies-in-go-can-be-unintuitive
func cleanlyCloseBody(body io.ReadCloser) error { _ = "STUB: not implemented"; return nil }

func (c *Client) sendHTTP(ctx context.Context, op *requestOp, msg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) sendBatchHTTP(ctx context.Context, op *requestOp, msgs []*jsonrpcMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *httpConn) doRequest(ctx context.Context, msg interface{}) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// set headers

// do request

// httpServerConn turns a HTTP connection into a Conn.
type httpServerConn struct {
	io.Reader
	io.Writer
	r *http.Request
}

func (s *Server) newHTTPServerConn(r *http.Request, w http.ResponseWriter) ServerCodec {
	_ = "STUB: not implemented"
	return *new(ServerCodec)
}

// It's an error response and requires special treatment.
//
// In case of a timeout error, the response must be written before the HTTP
// server's write timeout occurs. So we need to flush the response. The
// Content-Length header also needs to be set to ensure the client knows
// when it has the full response.

// If this request is wrapped in a handler that might remove Content-Length (such
// as the automatic gzip we do in package node), we need to ensure the HTTP server
// doesn't perform chunked encoding. In case WriteTimeout is reached, the chunked
// encoding might not be finished correctly, and some clients do not like it when
// the final chunk is missing.

// Close does nothing and always returns nil.
func (t *httpServerConn) Close() error {
	_ = "STUB: not implemented"

	// RemoteAddr returns the peer address of the underlying connection.
	return nil
}

func (t *httpServerConn) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

// SetWriteDeadline does nothing and always returns nil.
func (t *httpServerConn) SetWriteDeadline(time.Time) error {
	_ = "STUB: not implemented"

	// ServeHTTP serves JSON-RPC requests over HTTP.
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Permit dumb empty requests for remote health-checks (AWS)
	return
}

// Create request-scoped context.

// All checks passed, create a codec that reads directly from the request body
// until EOF, writes the response to w, and orders the server to process a
// single request.

// validateRequest returns a non-zero response code and error message if the
// request is invalid.
func (s *Server) validateRequest(r *http.Request) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Allow OPTIONS (regardless of content-type)

// Check content-type

// Invalid content-type

// ContextRequestTimeout returns the request timeout derived from the given context.
func ContextRequestTimeout(ctx context.Context) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// If the context is an HTTP request context, use the server's WriteTimeout.

// When a write timeout is configured, we need to send the response message before
// the HTTP server cuts connection. So our internal timeout must be earlier than
// the server's true timeout.
//
// Note: Timeouts are sanitized to be a minimum of 1 second.
// Also see issue: https://github.com/golang/go/issues/47229
