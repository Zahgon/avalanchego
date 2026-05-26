// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/rpc"
)

func (b *backend) CurrentBlock() *types.Header { _ = "STUB: not implemented"; return nil }

func (b *backend) HeaderByNumber(ctx context.Context, n rpc.BlockNumber) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) BlockByNumber(ctx context.Context, n rpc.BlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) GetBody(ctx context.Context, hash common.Hash, number rpc.BlockNumber) (*types.Body, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
