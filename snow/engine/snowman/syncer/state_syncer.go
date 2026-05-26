// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package syncer

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

// maxOutstandingBroadcastRequests is the maximum number of requests to have
// outstanding when broadcasting.
const maxOutstandingBroadcastRequests = 50

var _ common.StateSyncer = (*stateSyncer)(nil)

// summary content as received from network, along with accumulated weight.
type weightedSummary struct {
	summary block.StateSummary
	weight  uint64
}

type stateSyncer struct {
	Config

	// list of NoOpsHandler for messages dropped by state syncer
	common.AcceptedFrontierHandler
	common.AcceptedHandler
	common.AncestorsHandler
	common.PutHandler
	common.QueryHandler
	common.ChitsHandler
	common.AppHandler
	common.SimplexHandler

	started bool

	// Tracks the last requestID that was used in a request
	requestID uint32

	stateSyncVM        block.StateSyncableVM
	onDoneStateSyncing func(ctx context.Context, lastReqID uint32) error

	// we track the (possibly nil) local summary to help engine
	// choosing among multiple validated summaries
	locallyAvailableSummary block.StateSummary

	// Holds the beacons that were sampled for the accepted frontier
	// Won't be consumed as seeders are reached out. Used to rescale
	// alpha for frontiers
	frontierSeeders validators.Manager
	// IDs of validators we should request state summary frontier from.
	// Will be consumed seeders are reached out for frontier.
	targetSeeders set.Set[ids.NodeID]
	// IDs of validators we requested a state summary frontier from
	// but haven't received a reply yet. ID is cleared if/when reply arrives.
	pendingSeeders set.Set[ids.NodeID]
	// IDs of validators that failed to respond with their state summary frontier
	failedSeeders set.Set[ids.NodeID]

	// IDs of validators we should request filtering the accepted state summaries from
	targetVoters set.Set[ids.NodeID]
	// IDs of validators we requested filtering the accepted state summaries from
	// but haven't received a reply yet. ID is cleared if/when reply arrives.
	pendingVoters set.Set[ids.NodeID]
	// IDs of validators that failed to respond with their filtered accepted state summaries
	failedVoters set.Set[ids.NodeID]

	// summaryID --> (summary, weight)
	weightedSummaries map[ids.ID]*weightedSummary

	// summaries received may be different even if referring to the same height
	// we keep a list of deduplicated height ready for voting
	summariesHeights       set.Set[uint64]
	uniqueSummariesHeights []uint64
}

func New(
	cfg Config,
	onDoneStateSyncing func(ctx context.Context, lastReqID uint32) error,
) common.StateSyncer {
	_ = "STUB: not implemented"
	return *new(common.StateSyncer)
}

func (ss *stateSyncer) Start(ctx context.Context, startReqID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (ss *stateSyncer) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (ss *stateSyncer) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// tryStartSyncing will start syncing the first time it is called while the
// startupTracker is reporting that the protocol should start.
func (ss *stateSyncer) tryStartSyncing(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ss *stateSyncer) StateSummaryFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, summaryBytes []byte) error {
	_ = "STUB: not implemented"
	// ignores any late responses
	return nil
}

// Mark that we received a response from [nodeID]

// retrieve summary ID and register frontier;
// make sure next beacons are reached out
// even in case invalid summaries are received

func (ss *stateSyncer) GetStateSummaryFrontierFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	// ignores any late responses
	return nil
}

// Mark that we didn't get a response from [nodeID]

func (ss *stateSyncer) receivedStateSummaryFrontier(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// still waiting on requests

// All nodes reached out for the summary frontier have responded or timed out.
// If enough of them have indeed responded we'll go ahead and ask
// each state syncer (not just a sample) to filter the list of state summaries
// that we were told are on the accepted frontier.
// If we got too many timeouts, we restart state syncing hoping that network
// problems will go away and we can collect a qualified frontier.
// We assume the frontier is qualified after an alpha proportion of frontier seeders have responded

func (ss *stateSyncer) AcceptedStateSummary(ctx context.Context, nodeID ids.NodeID, requestID uint32, summaryIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	// ignores any late responses
	return nil
}

// Mark that we received a response from [nodeID]

// wait on pending responses

// We've received the filtered accepted frontier from every state sync validator
// Drop all summaries without a sufficient weight behind them

// if we don't have enough weight for the state summary to be accepted then retry or fail the state sync

// retry the state sync if the weight is not enough to state sync

// if we had too many timeouts when asking for validator votes, we should restart
// state sync hoping for the network problems to go away; otherwise, we received
// enough (>= ss.Alpha) responses, but no state summary was supported by a majority
// of validators (i.e. votes are split between minorities supporting different state
// summaries), so there is no point in retrying state sync; we should move ahead to bootstrapping

// if we do not restart state sync, move on to bootstrapping.

// VM did not accept the summary, move on to bootstrapping.

// Summary was accepted and VM is state syncing.
// Engine will wait for notification of state sync done.

// Summary was accepted and VM is state syncing.
// Engine will continue into bootstrapping and the VM will sync in the
// background.

// selectSyncableStateSummary chooses a state summary from all
// the network validated summaries.
func (ss *stateSyncer) selectSyncableStateSummary() block.StateSummary {
	_ = "STUB: not implemented"
	return *new(block.StateSummary)
}

// by default pick highest summary, unless locallyAvailableSummary is still valid.
// In such case we pick locallyAvailableSummary to allow VM resuming state syncing.

func (ss *stateSyncer) GetAcceptedStateSummaryFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	// ignores any late responses
	return nil
}

// If we can't get a response from [nodeID], act as though they said that
// they think none of the containers we sent them in GetAccepted are
// accepted

// startup do start the whole state sync process by
// sampling frontier seeders, listing state syncers to request votes to
// and reaching out frontier seeders if any. Otherwise, it moves immediately
// to bootstrapping. Unlike Start, startup does not check
// whether sufficient stake amount is connected.
func (ss *stateSyncer) startup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// clear up messages trackers

// sample K beacons to retrieve frontier from

// Invariant: We never use the TxID or BLS keys populated here.

// list all beacons, to reach them for voting on frontier

// check if there is an ongoing state sync; if so add its state summary
// to the frontier to request votes on
// Note: database.ErrNotFound means there is no ongoing summary

// no action needed

// initiate messages exchange

// Ask up to [common.MaxOutstandingBroadcastRequests] state sync validators at a time
// to send their accepted state summary. It is called again until there are
// no more seeders to be reached in the pending set
func (ss *stateSyncer) sendGetStateSummaryFrontiers(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Ask up to [common.MaxOutstandingStateSyncRequests] syncers validators to send
// their filtered accepted frontier. It is called again until there are
// no more voters to be reached in the pending set.
func (ss *stateSyncer) sendGetAcceptedStateSummaries(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (ss *stateSyncer) Notify(ctx context.Context, msg common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (*stateSyncer) Gossip(context.Context) error { _ = "STUB: not implemented"; return nil }

func (ss *stateSyncer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*stateSyncer) Timeout(context.Context) error { _ = "STUB: not implemented"; return nil }

func (ss *stateSyncer) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ss *stateSyncer) IsEnabled(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// state sync is not implemented
		nil
}
