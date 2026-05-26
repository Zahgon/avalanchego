// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gwarp

import (
	"context"

	"github.com/ava-labs/avalanchego/vms/platformvm/warp"

	pb "github.com/ava-labs/avalanchego/proto/pb/warp"
)

var _ pb.SignerServer = (*Server)(nil)

type Server struct {
	pb.UnsafeSignerServer
	signer warp.Signer
}

func NewServer(signer warp.Signer) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Sign(_ context.Context, unsignedMsg *pb.SignRequest) (*pb.SignResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
