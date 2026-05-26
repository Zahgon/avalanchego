// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package greader

import (
	"context"
	"io"

	readerpb "github.com/ava-labs/avalanchego/proto/pb/io/reader"
)

var _ readerpb.ReaderServer = (*Server)(nil)

// Server is an io.Reader that is managed over RPC.
type Server struct {
	readerpb.UnsafeReaderServer
	reader io.Reader
}

// NewServer returns an io.Reader instance managed remotely
func NewServer(reader io.Reader) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Read(_ context.Context, req *readerpb.ReadRequest) (*readerpb.ReadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sentinel errors must be special-cased through an error code
