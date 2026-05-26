// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gconn

import (
	"context"
	"net"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm/grpcutils"

	connpb "github.com/ava-labs/avalanchego/proto/pb/net/conn"
)

var _ connpb.ConnServer = (*Server)(nil)

// Server is an http.Conn that is managed over RPC.
type Server struct {
	connpb.UnsafeConnServer
	conn   net.Conn
	closer *grpcutils.ServerCloser
}

// NewServer returns an http.Conn managed remotely
func NewServer(conn net.Conn, closer *grpcutils.ServerCloser) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Read(_ context.Context, req *connpb.ReadRequest) (*connpb.ReadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sentinel errors must be special-cased through an error code

func (s *Server) Write(_ context.Context, req *connpb.WriteRequest) (*connpb.WriteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Close(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) SetDeadline(_ context.Context, req *connpb.SetDeadlineRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) SetReadDeadline(_ context.Context, req *connpb.SetDeadlineRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) SetWriteDeadline(_ context.Context, req *connpb.SetDeadlineRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
