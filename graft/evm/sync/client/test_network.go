// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
)

var _ Network = (*testNetwork)(nil)

type testNetwork struct {
	// captured request data
	numCalls uint

	// response testing for RequestAny and Request calls
	response       [][]byte
	callback       func() // callback is called prior to processing each test call
	requestErr     []error
	nodesRequested []ids.NodeID
}

func (*testNetwork) P2PNetwork() *p2p.Network { _ = "STUB: not implemented"; return nil }

func (*testNetwork) Sample(context.Context, int) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}

func (t *testNetwork) SendSyncedAppRequestAny(_ context.Context, _ []byte) ([]byte, ids.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.NodeID), nil
}

func (t *testNetwork) SendSyncedAppRequest(_ context.Context, nodeID ids.NodeID, _ []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *testNetwork) processTest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *testNetwork) testResponse(times uint8, callback func(), response []byte) {
	_ = "STUB: not implemented"
	return
}

func (t *testNetwork) testResponses(callback func(), responses ...[]byte) {
	_ = "STUB: not implemented"
	return
}

func (*testNetwork) RegisterResponse(ids.NodeID, float64) { _ = "STUB: not implemented"; return }

func (*testNetwork) RegisterFailure(ids.NodeID) { _ = "STUB: not implemented"; return }
