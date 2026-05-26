// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapper

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ Poll = (*Majority)(nil)

// Majority implements the bootstrapping poll to filter the initial set of
// potentially acceptable blocks into a set of accepted blocks to sync to.
//
// Once the last accepted blocks have been fetched from the initial set of
// peers, the set of blocks are sent to all peers. Each peer is expected to
// filter the provided blocks and report which of them they consider accepted.
// If a majority of the peers report that a block is accepted, then the node
// will consider that block to be accepted by the network. This assumes that a
// majority of the network is correct. If a majority of the network is
// malicious, the node may accept an incorrect block.
type Majority struct {
	requests

	log         logging.Logger
	nodeWeights map[ids.NodeID]uint64

	// received maps the blockID to the total sum of weight that has reported
	// that block as accepted.
	received map[ids.ID]uint64
	accepted []ids.ID
}

func NewMajority(
	log logging.Logger,
	nodeWeights map[ids.NodeID]uint64,
	maxOutstanding int,
) *Majority {
	_ = "STUB: not implemented"
	return nil
}

func (m *Majority) RecordOpinion(_ context.Context, nodeID ids.NodeID, blkIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

// The chain router should have already dropped unexpected messages.

func (m *Majority) Result(context.Context) ([]ids.ID, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
