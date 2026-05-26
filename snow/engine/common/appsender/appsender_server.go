// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package appsender

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ava-labs/avalanchego/snow/engine/common"

	appsenderpb "github.com/ava-labs/avalanchego/proto/pb/appsender"
)

var _ appsenderpb.AppSenderServer = (*Server)(nil)

type Server struct {
	appsenderpb.UnsafeAppSenderServer
	appSender common.AppSender
}

// NewServer returns a messenger connected to a remote channel
func NewServer(appSender common.AppSender) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) SendAppRequest(ctx context.Context, req *appsenderpb.SendAppRequestMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) SendAppResponse(ctx context.Context, req *appsenderpb.SendAppResponseMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) SendAppError(ctx context.Context, req *appsenderpb.SendAppErrorMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) SendAppGossip(ctx context.Context, req *appsenderpb.SendAppGossipMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
