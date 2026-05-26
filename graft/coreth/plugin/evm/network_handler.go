// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"

	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers/stats"
	"github.com/ava-labs/avalanchego/ids"
)

var _ message.RequestHandler = (*networkHandler)(nil)

type LeafHandlers map[message.NodeType]handlers.LeafRequestHandler

type networkHandler struct {
	leafRequestHandlers LeafHandlers
	blockRequestHandler *handlers.BlockRequestHandler
	codeRequestHandler  *handlers.CodeRequestHandler
}

type LeafRequestTypeConfig struct {
	NodeType     message.NodeType
	NodeKeyLen   int
	TrieDB       *triedb.Database
	UseSnapshots bool
	MetricName   string
}

// newNetworkHandler constructs the handler for serving network requests.
func newNetworkHandler(
	provider handlers.SyncDataProvider,
	diskDB ethdb.KeyValueReader,
	networkCodec codec.Manager,
	leafRequestHandlers LeafHandlers,
	syncStats stats.HandlerStats,
) *networkHandler {
	_ = "STUB: not implemented"
	return nil
}

func (n networkHandler) HandleLeafsRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, leafsRequest message.LeafsRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n networkHandler) HandleBlockRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, blockRequest message.BlockRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n networkHandler) HandleCodeRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, codeRequest message.CodeRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
