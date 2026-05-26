// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers/stats"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
)

// SyncDataProvider combines [BlockProvider] and [SnapshotProvider] for sync operations.
type SyncDataProvider interface {
	BlockProvider
	SnapshotProvider
}

// BlockProvider provides blocks by hash and height.
type BlockProvider interface {
	GetBlock(common.Hash, uint64) *types.Block
}

const (
	// parentLimit specifies how many parents to retrieve and send given a starting hash
	// This value overrides any specified limit in blockRequest.Parents if it is greater than this value
	parentLimit           = uint16(64)
	targetMessageByteSize = units.MiB - units.KiB // Target total block bytes slightly under original network codec max size of 1MB
)

// BlockRequestHandler is a peer.RequestHandler for message.BlockRequest
// serving requested blocks starting at specified hash
type BlockRequestHandler struct {
	stats         stats.BlockRequestHandlerStats
	blockProvider BlockProvider
	codec         codec.Manager
}

func NewBlockRequestHandler(blockProvider BlockProvider, codec codec.Manager, handlerStats stats.BlockRequestHandlerStats) *BlockRequestHandler {
	_ = "STUB: not implemented"
	return nil
}

// OnBlockRequest handles incoming message.BlockRequest, returning blocks as requested
// Never returns error
// Expects returned errors to be treated as FATAL
// Returns empty response or subset of requested blocks if ctx expires during fetch
// Assumes ctx is active
func (b *BlockRequestHandler) OnBlockRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, blockRequest message.BlockRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// override given Parents limit if it is greater than parentLimit

// ensure metrics are captured properly on all return paths

// we return whatever we have until ctx errors, limit is exceeded, or we reach the genesis block
// this will happen either when the ctx is cancelled or we hit the ctx deadline

// drop this request
