// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gwriter

import (
	"context"
	"io"

	writerpb "github.com/ava-labs/avalanchego/proto/pb/io/writer"
)

var _ writerpb.WriterServer = (*Server)(nil)

// Server is an http.Handler that is managed over RPC.
type Server struct {
	writerpb.UnsafeWriterServer
	writer io.Writer
}

// NewServer returns an http.Handler instance managed remotely
func NewServer(writer io.Writer) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Write(_ context.Context, req *writerpb.WriteRequest) (*writerpb.WriteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
