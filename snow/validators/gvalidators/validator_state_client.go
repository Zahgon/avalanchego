// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gvalidators

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"

	pb "github.com/ava-labs/avalanchego/proto/pb/validatorstate"
)

var (
	_                             validators.State = (*Client)(nil)
	errFailedPublicKeyDeserialize                  = errors.New("couldn't deserialize public key")
)

type Client struct {
	client pb.ValidatorStateClient
}

func NewClient(client pb.ValidatorStateClient) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) GetMinimumHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Client) GetCurrentHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Client) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (c *Client) GetWarpValidatorSets(
	ctx context.Context,
	height uint64,
) (map[ids.ID]validators.WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublicKeyFromValidUncompressedBytes is used rather than
// PublicKeyFromCompressedBytes because it is significantly faster
// due to the avoidance of decompression and key re-verification. We
// can safely assume that the BLS Public Keys are verified before
// being added to the P-Chain and served by the gRPC server.

func (c *Client) GetCurrentValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) (map[ids.ID]*validators.GetCurrentValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// PublicKeyFromValidUncompressedBytes is used rather than
// PublicKeyFromCompressedBytes because it is significantly faster
// due to the avoidance of decompression and key re-verification. We
// can safely assume that the BLS Public Keys are verified before
// being added to the P-Chain and served by the gRPC server.

func warpValidatorsFromProto(proto []*pb.WarpValidator) ([]*validators.Warp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
