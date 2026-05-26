// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman/poll"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/common/tracker"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/ancestor"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/job"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/bimap"
	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	nonVerifiedCacheSize = 64 * units.MiB
	errInsufficientStake = "insufficient connected stake"
)

var _ common.Engine = (*Engine)(nil)

func cachedBlockSize(_ ids.ID, blk snowman.Block) int { _ = "STUB: not implemented"; return 0 }

// Engine implements the Engine interface by attempting to fetch all
// Engine dependencies.
type Engine struct {
	Config
	*metrics

	// list of NoOpsHandler for messages dropped by engine
	common.StateSummaryFrontierHandler
	common.AcceptedStateSummaryHandler
	common.AcceptedFrontierHandler
	common.AcceptedHandler
	common.AncestorsHandler
	common.AppHandler
	common.SimplexHandler
	validators.Connector

	requestID uint32

	// track outstanding preference requests
	polls poll.Set

	// blocks that have we have sent get requests for but haven't yet received
	blkReqs            *bimap.BiMap[common.Request, ids.ID]
	blkReqSourceMetric map[common.Request]prometheus.Counter

	// blocks that are queued to be issued to consensus once missing dependencies are fetched
	// Block ID --> Block
	pending map[ids.ID]snowman.Block

	// Block ID --> Parent ID
	unverifiedIDToAncestor ancestor.Tree

	// Block ID --> Block.
	//
	// A block is put into this cache if its ancestry was fetched, but the block
	// was not able to be issued. A block may fail to be issued if verification
	// on the block or one of its ancestors returns an error.
	unverifiedBlockCache cache.Cacher[ids.ID, snowman.Block]

	// acceptedFrontiers of the other validators of this chain
	acceptedFrontiers tracker.Accepted

	// operations that are blocked on a block being issued. This could be
	// issuing another block, responding to a query, or applying votes to consensus
	blocked *job.Scheduler[ids.ID]

	// number of times build block needs to be called once the number of
	// processing blocks has gone below the optimal number.
	pendingBuildBlocks int
}

func New(config Config) (*Engine, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Engine) Gossip(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// repoll is called here to unblock the engine if it previously errored
// when attempting to issue a query. This can happen if a subnet was
// temporarily misconfigured and there were no validators.

// Uniform sampling is used here to reduce bandwidth requirements of
// nodes with a large amount of stake weight.

func (e *Engine) Put(ctx context.Context, nodeID ids.NodeID, requestID uint32, blkBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// because GetFailed doesn't utilize the assumption that we actually
// sent a Get message, we can safely call GetFailed here to potentially
// abandon the request.

// We assume that [blk] is useless because it doesn't match what we
// expected.

// This can happen if this block was provided to this engine while a Get
// request was outstanding. For example, the block may have been locally
// built or the node may have received a PushQuery with this block.
//
// Note: It is still possible this block will be issued here, because
// the block may have previously failed verification.

// issue the block into consensus. If the block has already been issued,
// this will be a noop. If this block has missing dependencies, vdr will
// receive requests to fill the ancestry. dependencies that have already
// been fetched, but with missing dependencies themselves won't be requested
// from the vdr.

func (e *Engine) GetFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	// We don't assume that this function is called after a failed Get message.
	// Check to see if we have an outstanding request and also get what the
	// request was for if it exists.
	return nil
}

// Because the get request was dropped, we no longer expect blkID to be
// issued.

func (e *Engine) PullQuery(ctx context.Context, nodeID ids.NodeID, requestID uint32, blkID ids.ID, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to issue [blkID] to consensus.
// If we're missing an ancestor, request it from [vdr]

func (e *Engine) PushQuery(ctx context.Context, nodeID ids.NodeID, requestID uint32, blkBytes []byte, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// If parsing fails, we just drop the request, as we didn't ask for it

// issue the block into consensus. If the block has already been issued,
// this will be a noop. If this block has missing dependencies, nodeID will
// receive requests to fill the ancestry. dependencies that have already
// been fetched, but with missing dependencies themselves won't be requested
// from the vdr.

func (e *Engine) Chits(ctx context.Context, nodeID ids.NodeID, requestID uint32, preferredID ids.ID, preferredIDAtHeight ids.ID, acceptedID ids.ID, acceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: The order of [responseOptions] must be [preferredID] then
// (optionally) [preferredIDAtHeight]. During vote application, the
// first vote that can be applied will be used. So, the votes should be
// populated in order of decreasing height.

// Will record chits once [preferredID] and [preferredIDAtHeight] have been
// issued into consensus

// Wait until [preferredID] and [preferredIDAtHeight] have been issued to
// consensus before applying this chit.

func (e *Engine) QueryFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) Notify(ctx context.Context, msg common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// the pending txs message means we should attempt to build a block.

func (e *Engine) Context() *snow.ConsensusContext { _ = "STUB: not implemented"; return nil }

func (e *Engine) Start(ctx context.Context, startReqID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// initialize consensus to the last accepted blockID

// to maintain the invariant that oracle blocks are issued in the correct
// preferences, we need to handle the case that we are bootstrapping into an oracle block

// if there aren't blocks we need to deliver on startup, we need to set
// the preference to the last accepted block

// note that deliver will set the VM's preference

func (e *Engine) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) executeDeferredWork(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) getBlock(ctx context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (e *Engine) sendChits(ctx context.Context, nodeID ids.NodeID, requestID uint32, requestedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

// If we aren't fully verifying blocks, only vote for blocks that are widely
// preferred by the validator set.

// Because we only return accepted state here, it's fairly likely
// that the requested height is higher than the last accepted block.
// That means that this code path is actually quite common.

// If this chain is pruning historical blocks, it's expected for a
// node to be unable to fetch some block IDs. In this case, we fall
// back to returning the last accepted ID.
//
// Because it is possible for a byzantine node to spam requests at
// old heights on a pruning network, we log this as debug. However,
// this case is unexpected to be hit by correct peers.

// If the requested height is higher than our preferred tip, we
// don't prefer anything at the requested height yet.

// Build blocks if they have been requested and the number of processing blocks
// is less than optimal.
func (e *Engine) buildBlocks(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// The newly created block should be built on top of the preferred block.
// Otherwise, the new block doesn't have the best chance of being confirmed.

// TODO: Technically this may incorrectly log a warning if the block
// that was just built caused votes to be applied such that the block
// was rejected or was accepted along with one of its children. This
// should be cleaned up to never produce an invalid warning.

// Issue another poll to the network, asking what it prefers given the block we prefer.
// Helps move consensus along.
func (e *Engine) repoll(ctx context.Context) {
	_ = "STUB: not implemented"
	// if we are issuing a repoll, we should gossip our current preferences to
	// propagate the most likely branch as quickly as possible
	return
}

// issueFromByID attempts to issue the branch ending with a block [blkID] into
// consensus.
// If we do not have [blkID], request it.
func (e *Engine) issueFromByID(
	ctx context.Context,
	nodeID ids.NodeID,
	blkID ids.ID,
	issuedMetric prometheus.Counter,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the block is not locally available, request it from the peer.

//nolint:nilerr

// issueFrom attempts to issue the branch ending with block [blkID] to
// consensus.
// If a dependency is missing, it will be requested it from [nodeID].
func (e *Engine) issueFrom(
	ctx context.Context,
	nodeID ids.NodeID,
	blk snowman.Block,
	issuedMetric prometheus.Counter,
) error {
	_ = "STUB: not implemented"
	// issue [blk] and its ancestors to consensus.
	return nil
}

// If we don't have this ancestor, request it from [nodeID]

// If the block is not locally available, request it from the peer.

//nolint:nilerr

// Remove any outstanding requests for this block

// If this block isn't pending, make sure nothing is blocked on it.

// issueWithAncestors attempts to issue the branch ending with [blk] to
// consensus.
// If a dependency is missing and the dependency hasn't been requested, the
// issuance will be abandoned.
func (e *Engine) issueWithAncestors(
	ctx context.Context,
	blk snowman.Block,
	issuedMetric prometheus.Counter,
) error {
	_ = "STUB: not implemented"

	// issue [blk] and its ancestors into consensus
	return nil
}

// There's an outstanding request for this block. We can wait for that
// request to succeed or fail.

// If the block wasn't already issued, we have no reason to expect that it
// will be able to be issued.

// Issue [blk] to consensus once its ancestors have been issued.
// If [push] is true, a push query will be used. Otherwise, a pull query will be
// used.
func (e *Engine) issue(
	ctx context.Context,
	nodeID ids.NodeID,
	blk snowman.Block,
	push bool,
	issuedMetric prometheus.Counter,
) error {
	_ = "STUB: not implemented"

	// mark that the block is queued to be added to consensus once its ancestors have been
	return nil
}

// Remove any outstanding requests for this block

// Will add [blk] to consensus once its ancestors have been

// We know that shouldIssueBlock(blk) is true. This means that parent is
// either the last accepted block or is not decided.

// Request that [vdr] send us block [blkID]
func (e *Engine) sendRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	blkID ids.ID,
	issuedMetric prometheus.Counter,
) {
	_ = "STUB: not implemented"
	// There is already an outstanding request for this block
	return
}

// Send a query for this block. If push is set to true, blkBytes will be used to
// send a PushQuery. Otherwise, blkBytes will be ignored and a PullQuery will be
// sent.
func (e *Engine) sendQuery(
	ctx context.Context,
	blkID ids.ID,
	blkBytes []byte,
	push bool,
) {
	_ = "STUB: not implemented"
	return
}

func (e *Engine) abortDueToInsufficientConnectedStake(blkID ids.ID) bool {
	_ = "STUB: not implemented"
	return false
}

// issue [blk] to consensus
// If [push] is true, a push query will be used. Otherwise, a pull query will be
// used.
func (e *Engine) deliver(
	ctx context.Context,
	nodeID ids.NodeID,
	blk snowman.Block,
	push bool,
	issuedMetric prometheus.Counter,
) error {
	_ = "STUB: not implemented"
	// we are no longer waiting on adding the block to consensus, so it is no
	// longer pending
	return nil
}

// If the parent isn't processing or the last accepted block, then this
// block is effectively rejected.
// Additionally, if [blkID] is already in the processing set, it
// shouldn't be added to consensus again.

// By ensuring that the parent is either processing or accepted, it is
// guaranteed that the parent was successfully verified. This means that
// calling Verify on this block is allowed.

// Add all the oracle blocks if they exist. We call verify on all the blocks
// and add them to consensus before marking anything as fulfilled to avoid
// any potential reentrant bugs.

// If the block is now preferred, query the network for its preferences
// with this new block.

// It's possible that the blocks we just added to consensus were decided
// immediately by votes that were pending their issuance. If this is the
// case, we should not be requesting any chits.

// If we should issue multiple queries at the same time, we need to repoll

func (e *Engine) markAsUnverified(blk snowman.Block) {
	_ = "STUB: not implemented"
	// If this block is processing, we don't need to add it to non-verifieds.
	return
}

// We might still need this block so we can bubble votes to the parent.
//
// If the non-verified set contains the parentID, then we know that the
// parent is not decided and therefore blk is not decided.
// Similarly, if the parent is processing, then the parent is not decided
// and therefore blk is not decided.

// addUnverifiedBlockToConsensus returns whether the block was added and an
// error if one occurred while adding it to consensus.
func (e *Engine) addUnverifiedBlockToConsensus(
	ctx context.Context,
	nodeID ids.NodeID,
	blk snowman.Block,
	issuedMetric prometheus.Counter,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// make sure this block is valid

// if verify fails, then all descendants are also invalid

// getProcessingAncestor finds [initialVote]'s most recent ancestor that is
// processing in consensus. If no ancestor could be found, false is returned.
//
// Note: If [initialVote] is processing, then [initialVote] will be returned.
func (e *Engine) getProcessingAncestor(initialVote ids.ID) (ids.ID, bool) {
	_ = "STUB: not implemented"
	// If [bubbledVote] != [initialVote], it is guaranteed that [bubbledVote] is
	// in processing. Otherwise, we attempt to iterate through any blocks we
	// have at our disposal as a best-effort mechanism to find a valid ancestor.
	return *new(ids.ID), false
}

// If we haven't cached the block, drop [vote].

// shouldIssueBlock returns true if the provided block should be enqueued for
// issuance. If the block is already decided, already enqueued, or has already
// been issued, this function will return false.
func (e *Engine) shouldIssueBlock(blk snowman.Block) bool { _ = "STUB: not implemented"; return false }

// If the block is already pending, don't issue it again.
// If the block was previously issued, don't issue it again.

// canDependOn reports true if it is guaranteed for the provided block ID to
// eventually either be fulfilled or abandoned.
func (e *Engine) canDependOn(blkID ids.ID) bool { _ = "STUB: not implemented"; return false }

// canIssueChildOn reports true if it is valid for a child of parentID to be
// verified and added to consensus.
func (e *Engine) canIssueChildOn(parentID ids.ID) bool { _ = "STUB: not implemented"; return false }

// isDecided reports true if the provided block's height implies that the block
// is either Accepted or Rejected.
func (e *Engine) isDecided(blk snowman.Block) bool { _ = "STUB: not implemented"; return false }

// block is either accepted or rejected

// This is guaranteed not to underflow because the above check ensures
// [height] > 0.

// the parent was rejected

// PChainProgressUpdater is used to update the P-Chain height progress of this chain.
// It's only used by the engine instance that runs the P-chain.
type PChainProgressUpdater interface {
	SetProgress(height uint64)
}
