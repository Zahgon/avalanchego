// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/bimap"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

const (
	// We cache processed vertices where height = c * stripeDistance for c = {1,2,3...}
	// This forms a "stripe" of cached DAG vertices at height stripeDistance, 2*stripeDistance, etc.
	// This helps to limit the number of repeated DAG traversals performed
	stripeDistance = 2000
	stripeWidth    = 5
	cacheSize      = 100000

	// statusUpdateFrequency is how many containers should be processed between
	// logs
	statusUpdateFrequency = 5000

	// maxOutstandingGetAncestorsRequests is the maximum number of GetAncestors
	// sent but not yet responded to/failed
	maxOutstandingGetAncestorsRequests = 10

	epsilon = 1e-6 // small amount to add to time to avoid division by 0
)

var _ common.BootstrapableEngine = (*Bootstrapper)(nil)

func New(
	config Config,
	onFinished func(ctx context.Context, lastReqID uint32) error,
	reg prometheus.Registerer,
) (*Bootstrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: To align with the Snowman invariant, it should be guaranteed the VM is
// not used until after the Bootstrapper has been Started.
type Bootstrapper struct {
	Config

	// list of NoOpsHandler for messages dropped by Bootstrapper
	common.StateSummaryFrontierHandler
	common.AcceptedStateSummaryHandler
	common.AcceptedFrontierHandler
	common.AcceptedHandler
	common.PutHandler
	common.QueryHandler
	common.ChitsHandler
	common.AppHandler
	common.SimplexHandler

	metrics

	// tracks which validators were asked for which containers in which requests
	outstandingRequests     *bimap.BiMap[common.Request, ids.ID]
	outstandingRequestTimes map[common.Request]time.Time

	// IDs of vertices that we will send a GetAncestors request for once we are
	// not at the max number of outstanding requests
	needToFetch set.Set[ids.ID]

	// Contains IDs of vertices that have recently been processed
	processedCache *lru.Cache[ids.ID, struct{}]

	// Tracks the last requestID that was used in a request
	requestID uint32

	// Called when bootstrapping is done on a specific chain
	onFinished func(ctx context.Context, lastReqID uint32) error
}

func (b *Bootstrapper) Context() *snow.ConsensusContext { _ = "STUB: not implemented"; return nil }

func (b *Bootstrapper) Clear(context.Context) error { _ = "STUB: not implemented"; return nil }

// Ancestors handles the receipt of multiple containers. Should be received in
// response to a GetAncestors message to [nodeID] with request ID [requestID].
// Expects vtxs[0] to be the vertex requested in the corresponding GetAncestors.
func (b *Bootstrapper) Ancestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, vtxs [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// this message isn't in response to a request we made

// All vertices added to [verticesToProcess] have received transitive votes
// from the accepted frontier.

// Persists the vtx

// No need to fetch this vertex since we have it now

// TODO: Calculate bandwidth based on the vertices that were persisted to
// disk.

func (b *Bootstrapper) GetAncestorsFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// This node timed out their request.

// Send another request for the vertex

func (b *Bootstrapper) Connected(
	ctx context.Context,
	nodeID ids.NodeID,
	nodeVersion *version.Application,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Bootstrapper) Timeout(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Bootstrapper) Gossip(context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *Bootstrapper) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Bootstrapper) Notify(context.Context, common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) Start(ctx context.Context, startReqID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// If the network was already linearized, don't attempt to linearize it
// again.

// If a stop vertex is well known, accept that.

// If a stop vertex isn't well known, treat the current state as the final
// DAG state.
//
// Note: This is used to linearize networks that were created after the
// linearization occurred.

func (b *Bootstrapper) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add the vertices in [vtxIDs] to the set of vertices that we need to fetch,
// and then fetch vertices (and their ancestors) until either there are no more
// to fetch or we are at the maximum number of outstanding requests.
func (b *Bootstrapper) fetch(ctx context.Context, vtxIDs ...ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Length checked in predicate above

// Make sure we haven't already requested this vertex

// Make sure we don't already have this vertex

// If we aren't connected to any peers, we send a request to ourself
// which is guaranteed to fail. We send this message to use the
// message timeout as a retry mechanism. Once we are connected to
// another node again we will select them to sample from.

// request vertex and ancestors

// Process the vertices in [vtxs].
func (b *Bootstrapper) process(ctx context.Context, vtxs ...avalanche.Vertex) error {
	_ = "STUB: not implemented"
	// Vertices that we need to process prioritized by vertices that are unknown
	// or the furthest down the DAG. Unknown vertices are prioritized to ensure
	// that once we have made it below a certain height in DAG traversal we do
	// not need to reset and repeat DAG traversals.
	return nil
}

// only process a vertex if we haven't already

// We don't have this vertex locally. Mark that we need to fetch it.

// Add to queue of vertices to execute when bootstrapping finishes.

// If the vertex is already on the queue, then we have already
// pushed [vtx]'s transactions and traversed into its parents.

// Add to queue of txs to execute when bootstrapping finishes.

// Periodically print progress

// Process the parents of this vertex (traverse up the DAG)

// But only if we haven't processed the parent

// See comment for stripeDistance

// Set new height and reset [vtxHeightSet]

// startSyncing starts bootstrapping. Process the vertices in [accepterContainerIDs].
func (b *Bootstrapper) startSyncing(ctx context.Context, acceptedContainerIDs []ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Append the list of accepted container IDs to pendingContainerIDs to ensure
// we iterate over every container that must be traversed.

// Process this vertex.

// We don't have this vertex. Mark that we have to fetch it.

// checkFinish repeatedly executes pending transactions and requests new frontier blocks until there aren't any new ones
// after which it finishes the bootstrap process
func (b *Bootstrapper) checkFinish(ctx context.Context) error {
	_ = "STUB: not implemented"
	// If we still need to fetch vertices, we can't finish
	return nil
}

// Invariant: edge will only be the stop vertex

// A vertex is less than another vertex if it is unknown. Ties are broken by
// prioritizing vertices that have a greater height.
func vertexLess(i, j avalanche.Vertex) bool { _ = "STUB: not implemented"; return false }

// Treat errors on retrieving the height as if the vertex is not fetched
