// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gruntime

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"

	pb "github.com/ava-labs/avalanchego/proto/pb/vm/runtime"
)

var _ pb.RuntimeServer = (*Server)(nil)

// Server is a VM runtime initializer controlled by RPC.
type Server struct {
	pb.UnsafeRuntimeServer
	runtime runtime.Initializer
}

func NewServer(runtime runtime.Initializer) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Initialize(ctx context.Context, req *pb.InitializeRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
