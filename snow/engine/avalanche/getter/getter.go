// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package getter

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/utils/set"
)

// Get requests are always served, regardless node state (bootstrapping or normal operations).
var _ common.AllGetsServer = (*getter)(nil)

func New(
	storage vertex.Storage,
	sender common.Sender,
	log logging.Logger,
	maxTimeGetAncestors time.Duration,
	maxContainersGetAncestors int,
	reg prometheus.Registerer,
) (common.AllGetsServer, error) {
	_ = "STUB: not implemented"
	return *new(common.AllGetsServer), nil
}

type getter struct {
	storage                   vertex.Storage
	sender                    common.Sender
	log                       logging.Logger
	maxTimeGetAncestors       time.Duration
	maxContainersGetAncestors int

	getAncestorsVtxs metric.Averager
}

func (gh *getter) GetStateSummaryFrontier(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) GetAcceptedStateSummary(_ context.Context, nodeID ids.NodeID, requestID uint32, _ set.Set[uint64]) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) GetAcceptedFrontier(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) GetAccepted(_ context.Context, nodeID ids.NodeID, requestID uint32, _ set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) GetAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, vtxID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't have the requested vertex. Drop message.

//nolint:nilerr

// for BFS

// length, in bytes, of vertex and its ancestors
// vertex and its ancestors in BFS order
// IDs of vertices that have been in queue before

// pop

// Ensure response size isn't too large. Include wrappers.IntLen because the size of the message
// is included with each container, and the size is repr. by an int.

// reached maximum response size

// Don't have this vertex;ignore

// If already visited, ignore

func (gh *getter) Get(_ context.Context, nodeID ids.NodeID, requestID uint32, _ ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}
