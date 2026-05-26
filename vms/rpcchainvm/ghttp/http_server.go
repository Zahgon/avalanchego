// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ghttp

import (
	"bytes"
	"context"
	"net/http"

	httppb "github.com/ava-labs/avalanchego/proto/pb/http"
)

var (
	_ httppb.HTTPServer   = (*Server)(nil)
	_ http.ResponseWriter = (*ResponseWriter)(nil)
)

// Server is an http.Handler that is managed over RPC.
type Server struct {
	httppb.UnsafeHTTPServer
	handler http.Handler
}

// NewServer returns an http.Handler instance managed remotely
func NewServer(handler http.Handler) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Handle(ctx context.Context, req *httppb.HTTPRequest) (*httppb.HTTPResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create the request with the current context

// always true per https://pkg.go.dev/crypto/tls#ConnectionState

// HandleSimple handles http requests over http2 using a simple request response model.
// Websockets are not supported.
func (s *Server) HandleSimple(ctx context.Context, r *httppb.HandleSimpleHTTPRequest) (*httppb.HandleSimpleHTTPResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResponseWriter struct {
	body       *bytes.Buffer
	header     http.Header
	statusCode int
}

// newResponseWriter returns very basic implementation of the http.ResponseWriter
func newResponseWriter() *ResponseWriter { _ = "STUB: not implemented"; return nil }

func (w *ResponseWriter) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (w *ResponseWriter) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *ResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (w *ResponseWriter) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (w *ResponseWriter) Body() *bytes.Buffer { _ = "STUB: not implemented"; return nil }
