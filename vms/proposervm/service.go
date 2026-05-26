// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"net/http"

	"connectrpc.com/connect"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/connectproto/pb/proposervm/proposervmconnect"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"

	pb "github.com/ava-labs/avalanchego/connectproto/pb/proposervm"
	avajson "github.com/ava-labs/avalanchego/utils/json"
)

var _ proposervmconnect.ProposerVMHandler = (*connectrpcService)(nil)

type connectrpcService struct {
	vm *VM
}

func (c *connectrpcService) GetProposedHeight(ctx context.Context, r *connect.Request[pb.GetProposedHeightRequest]) (*connect.Response[pb.GetProposedHeightReply], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *connectrpcService) GetCurrentEpoch(ctx context.Context, r *connect.Request[pb.GetCurrentEpochRequest]) (*connect.Response[pb.GetCurrentEpochReply], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type jsonrpcService struct {
	vm *VM
}

func (j *jsonrpcService) GetProposedHeight(r *http.Request, _ *struct{}, reply *api.GetHeightResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type GetEpochResponse struct {
	Number       avajson.Uint64 `json:"number"`
	StartTime    avajson.Uint64 `json:"startTime"`
	PChainHeight avajson.Uint64 `json:"pChainHeight"`
}

func (j *jsonrpcService) GetCurrentEpoch(r *http.Request, _ *struct{}, reply *GetEpochResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) getCurrentEpoch(ctx context.Context) (block.Epoch, error) {
	_ = "STUB: not implemented"
	return *new(block.Epoch), nil
}
