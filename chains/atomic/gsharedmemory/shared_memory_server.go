// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsharedmemory

import (
	"context"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database"

	sharedmemorypb "github.com/ava-labs/avalanchego/proto/pb/sharedmemory"
)

var _ sharedmemorypb.SharedMemoryServer = (*Server)(nil)

// Server is shared memory that is managed over RPC.
type Server struct {
	sharedmemorypb.UnsafeSharedMemoryServer
	sm atomic.SharedMemory
	db database.Database
}

// NewServer returns shared memory connected to remote shared memory
func NewServer(sm atomic.SharedMemory, db database.Database) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Get(
	_ context.Context,
	req *sharedmemorypb.GetRequest,
) (*sharedmemorypb.GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Indexed(
	_ context.Context,
	req *sharedmemorypb.IndexedRequest,
) (*sharedmemorypb.IndexedResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Apply(
	_ context.Context,
	req *sharedmemorypb.ApplyRequest,
) (*sharedmemorypb.ApplyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
