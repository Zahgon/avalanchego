// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"context"
	"errors"

	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers"
	"github.com/ava-labs/avalanchego/ids"
)

var (
	_ handlers.LeafRequestHandler = (*uninitializedHandler)(nil)

	errUninitialized = errors.New("uninitialized handler")
)

type uninitializedHandler struct{}

func (*uninitializedHandler) OnLeafsRequest(context.Context, ids.NodeID, uint32, message.LeafsRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// atomicLeafHandler is a wrapper around handlers.LeafRequestHandler that allows for initialization after creation
}

type leafHandler struct {
	handlers.LeafRequestHandler
}

// NewAtomicLeafHandler returns a new uninitialized leafHandler that can be later initialized
func NewLeafHandler() *leafHandler { _ = "STUB: not implemented"; return nil }

// Initialize initializes the atomicLeafHandler with the provided atomicTrieDB, trieKeyLength, and networkCodec
func (a *leafHandler) Initialize(atomicTrieDB *triedb.Database, trieKeyLength int, networkCodec codec.Manager) {
	_ = "STUB: not implemented"
	return
}
