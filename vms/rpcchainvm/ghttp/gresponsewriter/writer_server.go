// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gresponsewriter

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"

	responsewriterpb "github.com/ava-labs/avalanchego/proto/pb/http/responsewriter"
)

var (
	errUnsupportedFlushing  = errors.New("response writer doesn't support flushing")
	errUnsupportedHijacking = errors.New("response writer doesn't support hijacking")

	_ responsewriterpb.WriterServer = (*Server)(nil)
)

// Server is an http.ResponseWriter that is managed over RPC.
type Server struct {
	responsewriterpb.UnsafeWriterServer
	writer http.ResponseWriter
}

// NewServer returns an http.ResponseWriter instance managed remotely
func NewServer(writer http.ResponseWriter) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Write(
	_ context.Context,
	req *responsewriterpb.WriteRequest,
) (*responsewriterpb.WriteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) WriteHeader(
	_ context.Context,
	req *responsewriterpb.WriteHeaderRequest,
) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Flush(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Hijack(context.Context, *emptypb.Empty) (*responsewriterpb.HijackResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
