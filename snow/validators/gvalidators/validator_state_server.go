// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gvalidators

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ava-labs/avalanchego/snow/validators"

	pb "github.com/ava-labs/avalanchego/proto/pb/validatorstate"
)

var _ pb.ValidatorStateServer = (*Server)(nil)

type Server struct {
	pb.UnsafeValidatorStateServer
	state validators.State
}

func NewServer(state validators.State) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) GetMinimumHeight(ctx context.Context, _ *emptypb.Empty) (*pb.GetMinimumHeightResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) GetCurrentHeight(ctx context.Context, _ *emptypb.Empty) (*pb.GetCurrentHeightResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) GetSubnetID(ctx context.Context, req *pb.GetSubnetIDRequest) (*pb.GetSubnetIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) GetWarpValidatorSets(ctx context.Context, req *pb.GetWarpValidatorSetsRequest) (*pb.GetWarpValidatorSetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) GetValidatorSet(ctx context.Context, req *pb.GetValidatorSetRequest) (*pb.GetValidatorSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Passing in the uncompressed bytes is a performance optimization
// to avoid the cost of calling PublicKeyFromCompressedBytes on the
// client side.

func (s *Server) GetCurrentValidatorSet(ctx context.Context, req *pb.GetCurrentValidatorSetRequest) (*pb.GetCurrentValidatorSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Passing in the uncompressed bytes is a performance optimization
// to avoid the cost of calling PublicKeyFromCompressedBytes on the
// client side.

func warpValidatorsToProto(vdrs []*validators.Warp) []*pb.WarpValidator {
	_ = "STUB: not implemented"
	return nil
}
