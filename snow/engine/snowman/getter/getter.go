// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package getter

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/utils/set"
)

// Get requests are always served, regardless node state (bootstrapping or normal operations).
var _ common.AllGetsServer = (*getter)(nil)

func New(
	vm block.ChainVM,
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
	vm   block.ChainVM
	ssVM block.StateSyncableVM // can be nil

	sender common.Sender
	log    logging.Logger
	// Max time to spend fetching a container and its ancestors when responding
	// to a GetAncestors
	maxTimeGetAncestors time.Duration
	// Max number of containers in an ancestors message sent by this node.
	maxContainersGetAncestors int

	getAncestorsBlks metric.Averager
}

func (gh *getter) GetStateSummaryFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	// Note: we do not check if gh.ssVM.StateSyncEnabled since we want all
	// nodes, including those disabling state sync to serve state summaries if
	// these are available
	return nil
}

func (gh *getter) GetAcceptedStateSummary(ctx context.Context, nodeID ids.NodeID, requestID uint32, heights set.Set[uint64]) error {
	_ = "STUB: not implemented"
	// If there are no requested heights, then we can return the result
	// immediately, regardless of if the underlying VM implements state sync.
	return nil
}

// Note: we do not check if gh.ssVM.StateSyncEnabled since we want all
// nodes, including those disabling state sync to serve state summaries if
// these are available

func (gh *getter) GetAcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) GetAccepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) GetAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (gh *getter) Get(ctx context.Context, nodeID ids.NodeID, requestID uint32, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// If we failed to get the block, that means either an unexpected error
// has occurred, [vdr] is not following the protocol, or the
// block has been pruned.

// Respond to the validator with the fetched block and the same requestID.
