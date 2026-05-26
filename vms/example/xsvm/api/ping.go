// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"context"

	"connectrpc.com/connect"

	"github.com/ava-labs/avalanchego/connectproto/pb/xsvm"
	"github.com/ava-labs/avalanchego/connectproto/pb/xsvm/xsvmconnect"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var _ xsvmconnect.PingHandler = (*PingService)(nil)

type PingService struct {
	Log logging.Logger
}

func (p *PingService) Ping(_ context.Context, request *connect.Request[xsvm.PingRequest]) (*connect.Response[xsvm.PingReply], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PingService) StreamPing(_ context.Context, server *connect.BidiStream[xsvm.StreamPingRequest, xsvm.StreamPingReply]) error {
	_ = "STUB: not implemented"
	return nil
}

// Client closed the send stream
