// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers/stats"
	"github.com/ava-labs/avalanchego/ids"
)

// CodeRequestHandler is a peer.RequestHandler for message.CodeRequest
// serving requested contract code bytes
type CodeRequestHandler struct {
	codeReader ethdb.KeyValueReader
	codec      codec.Manager
	stats      stats.CodeRequestHandlerStats
}

func NewCodeRequestHandler(codeReader ethdb.KeyValueReader, codec codec.Manager, stats stats.CodeRequestHandlerStats) *CodeRequestHandler {
	_ = "STUB: not implemented"
	return nil
}

// OnCodeRequest handles request to retrieve contract code by its hash in message.CodeRequest
// Never returns error
// Returns nothing if code hash is not found
// Expects returned errors to be treated as FATAL
// Assumes ctx is active
func (n *CodeRequestHandler) OnCodeRequest(_ context.Context, nodeID ids.NodeID, requestID uint32, codeRequest message.CodeRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// always report code read time metric

func isUnique(hashes []common.Hash) bool { _ = "STUB: not implemented"; return false }
