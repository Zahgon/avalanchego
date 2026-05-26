// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
)

var (
	_ Client         = (*TestClient)(nil)
	_ EthBlockParser = (*testBlockParser)(nil)
)

type TestClient struct {
	codec          codec.Manager
	leafsHandler   handlers.LeafRequestHandler
	leavesReceived int32
	codesHandler   *handlers.CodeRequestHandler
	codeReceived   int32
	blocksHandler  *handlers.BlockRequestHandler
	blocksReceived int32
	// GetLeafsIntercept is called on every GetLeafs request if set to a non-nil callback.
	// The returned response will be returned by TestClient to the caller.
	GetLeafsIntercept func(req message.LeafsRequest, res message.LeafsResponse) (message.LeafsResponse, error)
	// GetCodesIntercept is called on every GetCode request if set to a non-nil callback.
	// The returned response will be returned by TestClient to the caller.
	GetCodeIntercept func(hashes []common.Hash, codeBytes [][]byte) ([][]byte, error)
	// GetBlocksIntercept is called on every GetBlocks request if set to a non-nil callback.
	// The returned response will be returned by TestClient to the caller.
	GetBlocksIntercept func(blockReq message.BlockRequest, blocks types.Blocks) (types.Blocks, error)
}

func NewTestClient(
	codec codec.Manager,
	leafHandler handlers.LeafRequestHandler,
	codesHandler *handlers.CodeRequestHandler,
	blocksHandler *handlers.BlockRequestHandler,
) *TestClient {
	_ = "STUB: not implemented"
	return nil
}

func (*TestClient) AddClient(uint64) *p2p.Client { _ = "STUB: not implemented"; return nil }

func (*TestClient) StateSyncNodes() []ids.NodeID { _ = "STUB: not implemented"; return nil }

func (ml *TestClient) GetLeafs(ctx context.Context, request message.LeafsRequest) (message.LeafsResponse, error) {
	_ = "STUB: not implemented"
	return *new(message.LeafsResponse), nil
}

// Increment the number of leaves received by the test client

func (ml *TestClient) LeavesReceived() int32 { _ = "STUB: not implemented"; return 0 }

func (ml *TestClient) GetCode(ctx context.Context, hashes []common.Hash) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ml *TestClient) CodeReceived() int32 { _ = "STUB: not implemented"; return 0 }

func (ml *TestClient) GetBlocks(ctx context.Context, blockHash common.Hash, height uint64, numParents uint16) ([]*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Actual client retries until the context is canceled.

// Hack to avoid duplicate code

func (ml *TestClient) BlocksReceived() int32 { _ = "STUB: not implemented"; return 0 }

type testBlockParser struct{}

func newTestBlockParser() *testBlockParser { _ = "STUB: not implemented"; return nil }

func (*testBlockParser) ParseEthBlock(b []byte) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
