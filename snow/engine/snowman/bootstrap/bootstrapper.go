// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman/bootstrapper"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/bootstrap/interval"
	"github.com/ava-labs/avalanchego/utils/bimap"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/timer"
	"github.com/ava-labs/avalanchego/version"
)

const (
	// Delay bootstrapping to avoid potential CPU burns
	bootstrappingDelay = 10 * time.Second

	// statusUpdateFrequency is how many containers should be processed between
	// logs
	statusUpdateFrequency = 5000

	// maxOutstandingBroadcastRequests is the maximum number of requests to have
	// outstanding when broadcasting.
	maxOutstandingBroadcastRequests = 50

	// minimumLogInterval is the minimum time between log entries to avoid noise
	minimumLogInterval = 5 * time.Second

	epsilon = 1e-6 // small amount to add to time to avoid division by 0
)

var (
	_ common.BootstrapableEngine = (*Bootstrapper)(nil)

	errUnexpectedTimeout = errors.New("unexpected timeout fired")
)

// bootstrapper repeatedly performs the bootstrapping protocol.
//
//  1. Wait until a sufficient amount of stake is connected.
//  2. Sample a small number of nodes to get the last accepted block ID
//  3. Verify against the full network that the last accepted block ID received
//     in step 2 is an accepted block.
//  4. Sync the full ancestry of the last accepted block.
//  5. Execute all the fetched blocks that haven't already been executed.
//  6. Restart the bootstrapping protocol until the number of blocks being
//     accepted during a bootstrapping round stops decreasing.
//
// Note: Because of step 6, the bootstrapping protocol will generally be
// performed multiple times.
//
// Invariant: The VM is not guaranteed to be initialized until Start has been
// called, so it must be guaranteed the VM is not used until after Start.
type Bootstrapper struct {
	Config

	*metrics
	TimeoutRegistrar common.TimeoutRegistrar
	// list of NoOpsHandler for messages dropped by bootstrapper
	common.StateSummaryFrontierHandler
	common.AcceptedStateSummaryHandler
	common.PutHandler
	common.QueryHandler
	common.ChitsHandler
	common.AppHandler
	common.SimplexHandler

	requestID uint32 // Tracks the last requestID that was used in a request

	started   bool
	restarted bool

	minority bootstrapper.Poll
	majority bootstrapper.Poll

	// Greatest height of the blocks passed in startSyncing
	tipHeight uint64
	// Height of the last accepted block when bootstrapping starts
	startingHeight uint64
	// Number of blocks that were fetched on startSyncing
	initiallyFetched uint64
	// Time that startSyncing was last called
	startTime time.Time
	// Time of the last progress update for accurate ETA calculation
	lastProgressUpdateTime time.Time

	// ETA tracker for more accurate time estimates
	etaTracker *timer.EtaTracker

	// tracks which validators were asked for which containers in which requests
	outstandingRequests     *bimap.BiMap[common.Request, ids.ID]
	outstandingRequestTimes map[common.Request]time.Time

	// number of state transitions executed
	executedStateTransitions uint64
	awaitingTimeout          bool

	tree            *interval.Tree
	missingBlockIDs set.Set[ids.ID]

	// bootstrappedOnce ensures that the [Bootstrapped] callback is only invoked
	// once, even if bootstrapping is retried.
	bootstrappedOnce sync.Once

	// Called when bootstrapping is done on a specific chain
	onFinished func(ctx context.Context, lastReqID uint32) error

	nonVerifyingParser block.Parser
}

func New(config Config, onFinished func(ctx context.Context, lastReqID uint32) error) (*Bootstrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bootstrapper) Context() *snow.ConsensusContext { _ = "STUB: not implemented"; return nil }

func (b *Bootstrapper) Clear(context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *Bootstrapper) Start(ctx context.Context, startReqID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Set the starting height

func (b *Bootstrapper) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// tryStartBootstrapping will start bootstrapping the first time it is called
// while the startupTracker is reporting that the protocol should start.
func (b *Bootstrapper) tryStartBootstrapping(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) startBootstrapping(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) sendBootstrappingMessagesOrFinish(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// We haven't finalized the accepted frontier, so we should wait for the
// outstanding requests.

// We haven't finalized the accepted set, so we should wait for the
// outstanding requests.

// Invariant: These functions are mutually recursive. However, when
// [startBootstrapping] calls [sendMessagesOrFinish], it is guaranteed
// to exit when sending GetAcceptedFrontier requests.

func (b *Bootstrapper) AcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) GetAcceptedFrontierFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) Accepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) GetAcceptedFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) startSyncing(ctx context.Context, acceptedBlockIDs []ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: if `GetBlock` returns an error other than
// `database.ErrNotFound`, then the error should be propagated.

// Add the first sample to the EtaTracker to establish an accurate baseline
// It's okay to call this a few times if startSyncing is called more than once.

// Process received blocks

// Get block [blkID] and its ancestors from a validator
func (b *Bootstrapper) fetch(ctx context.Context, blkID ids.ID) error {
	_ = "STUB: not implemented"
	// Make sure we haven't already requested this block
	return nil
}

// If we aren't connected to any peers, we send a request to ourself
// which is guaranteed to fail. We send this message to use the message
// timeout as a retry mechanism. Once we are connected to another node
// again we will select them to sample from.

// request block and ancestors

// Ancestors handles the receipt of multiple containers. Should be received in
// response to a GetAncestors message to [nodeID] with request ID [requestID]
func (b *Bootstrapper) Ancestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, blks [][]byte) error {
	_ = "STUB: not implemented"
	// Make sure this is in response to a request we made
	return nil
}

// this message isn't in response to a request we made

// Send another request for this

// the provided blocks couldn't be parsed

// TODO: Calculate bandwidth based on the blocks that were persisted to
// disk.

func (b *Bootstrapper) GetAncestorsFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// This node timed out their request.

// Send another request for this

// process a series of consecutive blocks starting at [blk].
//
//   - blk is a block that is assumed to have been marked as acceptable by the
//     bootstrapping engine.
//   - ancestors is a set of blocks that can be used to optimistically lookup
//     parent blocks. This enables the engine to process multiple blocks without
//     relying on the VM to have stored blocks during `ParseBlock`.
func (b *Bootstrapper) process(
	ctx context.Context,
	blk snowman.Block,
	ancestors map[ids.ID]snowman.Block,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Update metrics and log statuses

// Check if it's time to log progress (both progress-based and time-based frequency)

// Update the last progress update time and previous progress for next iteration

// Only log if we have a valid ETA estimate

// Lower log level for restarted bootstrapping.

// Attempt to fetch the newly discovered block

// tryStartExecuting executes all pending blocks if there are no more blocks
// being fetched. After executing all pending blocks it will either restart
// bootstrapping, or transition into normal operations.
func (b *Bootstrapper) tryStartExecuting(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If a fatal error has occurred, include the last accepted block
// information.

// Note that executedBlocks < c*previouslyExecuted ( 0 <= c < 1 ) is enforced
// so that the bootstrapping process will terminate even as new blocks are
// being issued.

// If there is an additional callback, notify them that this chain has been
// synced.

// Notify the subnet that this chain is synced

// If the subnet hasn't finished bootstrapping, this chain should remain
// syncing.

// Restart bootstrapping after [bootstrappingDelay] to keep up to date
// on the latest tip.

func (b *Bootstrapper) getLastAccepted(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (b *Bootstrapper) Timeout() error { _ = "STUB: not implemented"; return nil }

func (b *Bootstrapper) restartBootstrapping(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) Notify(_ context.Context, msg common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bootstrapper) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Bootstrapper) Gossip(context.Context) error { _ = "STUB: not implemented"; return nil }
