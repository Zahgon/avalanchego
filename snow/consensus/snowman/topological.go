// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/utils/bag"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	errDuplicateAdd            = errors.New("duplicate block add")
	errUnknownParentBlock      = errors.New("unknown parent block")
	errTooManyProcessingBlocks = errors.New("too many processing blocks")
	errBlockProcessingTooLong  = errors.New("block processing too long")
	errAcceptanceTimeTooHigh   = errors.New("acceptance time too high")

	maxAcceptanceTime = 15 * time.Second

	_ Factory   = (*TopologicalFactory)(nil)
	_ Consensus = (*Topological)(nil)
)

// TopologicalFactory implements Factory by returning a topological struct
type TopologicalFactory struct {
	factory snowball.Factory
}

func (tf TopologicalFactory) New() Consensus { _ = "STUB: not implemented"; return *new(Consensus) }

// Topological implements the Snowman interface by using a tree tracking the
// strongly preferred branch. This tree structure amortizes network polls to
// vote on more than just the next block.
type Topological struct {
	Factory snowball.Factory

	metrics *metrics

	// pollNumber is the number of times RecordPolls has been called
	pollNumber uint64

	// ctx is the context this snowman instance is executing in
	ctx *snow.ConsensusContext

	// params are the parameters that should be used to initialize snowball
	// instances
	params snowball.Parameters

	lastAcceptedID     ids.ID
	lastAcceptedHeight uint64

	// blocks stores the last accepted block and all the pending blocks
	blocks map[ids.ID]*snowmanBlock // blockID -> snowmanBlock

	// preferredIDs stores the set of IDs that are currently preferred.
	preferredIDs set.Set[ids.ID]

	// preferredHeights maps a height to the currently preferred block ID at
	// that height.
	preferredHeights map[uint64]ids.ID // height -> blockID

	// preference is the preferred block with highest height
	preference ids.ID

	// Used in [calculateInDegree] and.
	// Should only be accessed in that method.
	// We use this one instance of set.Set instead of creating a
	// new set.Set during each call to [calculateInDegree].
	leaves set.Set[ids.ID]

	// Kahn nodes used in [calculateInDegree] and [markAncestorInDegrees].
	// Should only be accessed in those methods.
	// We use this one map instead of creating a new map
	// during each call to [calculateInDegree].
	kahnNodes map[ids.ID]kahnNode
}

// Used to track the kahn topological sort status
type kahnNode struct {
	// inDegree is the number of children that haven't been processed yet. If
	// inDegree is 0, then this node is a leaf
	inDegree int
	// votes for all the children of this node, so far
	votes bag.Bag[ids.ID]
}

// Used to track which children should receive votes
type votes struct {
	// parentID is the parent of all the votes provided in the votes bag
	parentID ids.ID
	// votes for all the children of the parent
	votes bag.Bag[ids.ID]
}

func (ts *Topological) Initialize(
	ctx *snow.ConsensusContext,
	params snowball.Parameters,
	lastAcceptedID ids.ID,
	lastAcceptedHeight uint64,
	lastAcceptedTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ts *Topological) NumProcessing() int { _ = "STUB: not implemented"; return 0 }

func (ts *Topological) Add(blk Block) error { _ = "STUB: not implemented"; return nil }

// Make sure a block is not inserted twice.

// add the block as a child of its parent, and add the block to the tree

// If we are extending the preference, this is the new preference

func (ts *Topological) Processing(blkID ids.ID) bool {
	_ = "STUB: not implemented"
	// The last accepted block is in the blocks map, so we first must ensure the
	// requested block isn't the last accepted block.
	return false
}

// If the block is in the map of current blocks and not the last accepted
// block, then it is currently processing.

func (ts *Topological) IsPreferred(blkID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (ts *Topological) LastAccepted() (ids.ID, uint64) {
	_ = "STUB: not implemented"
	return *new(ids.ID), 0
}

func (ts *Topological) Preference() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (ts *Topological) PreferenceAtHeight(height uint64) (ids.ID, bool) {
	_ = "STUB: not implemented"
	return *new(ids.ID), false
}

// The votes bag contains at most K votes for blocks in the tree. If there is a
// vote for a block that isn't in the tree, the vote is dropped.
//
// Votes are propagated transitively towards the genesis. All blocks in the tree
// that result in at least Alpha votes will record the poll on their children.
// Every other block will have an unsuccessful poll registered.
//
// After collecting which blocks should be voted on, the polls are registered
// and blocks are accepted/rejected as needed. The preference is then updated to
// equal the leaf on the preferred branch.
//
// To optimize the theoretical complexity of the vote propagation, a topological
// sort is done over the blocks that are reachable from the provided votes.
// During the sort, votes are pushed towards the genesis. To prevent iterating
// over all blocks that had unsuccessful polls, we set a flag on the block to
// know that any future traversal through that block should register an
// unsuccessful poll on that block and every descendant block.
//
// The complexity of this function is:
// - Runtime = 4 * |live set| + |votes|
// - Space = 2 * |live set| + |votes|
func (ts *Topological) RecordPoll(ctx context.Context, voteBag bag.Bag[ids.ID]) error {
	_ = "STUB: not implemented"
	// Register a new poll call
	return nil
}

// Since we received at least alpha votes, it's possible that
// we reached an alpha majority on a processing block.
// We must perform the traversals to calculate all block
// that reached an alpha majority.

// Populates [ts.kahnNodes] and [ts.leaves]
// Runtime = |live set| + |votes| ; Space = |live set| + |votes|

// Runtime = |live set| ; Space = |live set|

// Runtime = |live set| ; Space = Constant

// If the set of preferred IDs already contains the preference, then the
// preference is guaranteed to already be set correctly. This is because the
// value returned from vote reports the next preferred block after the last
// preferred block that was voted for. If this block was previously
// preferred, then we know that following the preferences down the chain
// will return the current preference.

// Runtime = 2 * |live set| ; Space = Constant

// Runtime = |live set| ; Space = Constant
// Traverse from the preferred ID to the last accepted ancestor.
//
// It is guaranteed that the first decided block we encounter is the last
// accepted block because the startBlock is the preferred block. The
// preferred block is guaranteed to either be the last accepted block or
// extend the accepted chain.

// Traverse from the preferred ID to the preferred child until there are no
// children.

// Invariant: Because the prior block had an initialized snowball
// instance, it must have a processing child. This guarantees that
// block.blk is non-nil here.

// HealthCheck returns information about the consensus health.
func (ts *Topological) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check average acceptance time

// .String() is needed here to ensure a human readable format

// takes in a list of votes and sets up the topological ordering. Returns the
// reachable section of the graph annotated with the number of inbound edges and
// the non-transitively applied votes. Also returns the list of leaf blocks.
func (ts *Topological) calculateInDegree(votes bag.Bag[ids.ID]) {
	_ = "STUB: not implemented"
	// Clear the Kahn node set
	return
}

// Clear the leaf set

// If the vote is for a block that isn't in the current pending set,
// then the vote is dropped

// If the vote is for the last accepted block, the vote is dropped

// The parent contains the snowball instance of its children

// Add the votes for this block to the parent's set of responses

// If the parent block already had registered votes, then there is no
// need to iterate into the parents

// If I've never seen this parent block before, it is currently a leaf.

// iterate through all the block's ancestors and set up the inDegrees of
// the blocks

// Increase the inDegree by one

// If we have already seen this block, then we shouldn't increase
// the inDegree of the ancestors through this block again.

// Nodes are only leaves if they have no inbound edges.

// convert the tree into a branch of snowball instances with at least alpha
// votes
func (ts *Topological) pushVotes() []votes { _ = "STUB: not implemented"; return nil }

// Pop one element of [leaves]

// Should never return false because we just
// checked that [ts.leaves] is non-empty.

// get the block and sort information about the block

// If there are at least Alpha votes, then this block needs to record
// the poll on the snowball instance

// If the block is accepted, then we don't need to push votes to the
// parent block

// Remove an inbound edge from the parent kahn node and push the votes.

// If the inDegree is zero, then the parent node is now a leaf

// apply votes to the branch that received an Alpha threshold and returns the
// next preferred block after the last preferred block that received an Alpha
// threshold.
func (ts *Topological) vote(ctx context.Context, voteStack []votes) (ids.ID, error) {
	_ = "STUB: not implemented"
	// If the voteStack is empty, then the full tree should falter. This won't
	// change the preferred branch.
	return *new(ids.ID), nil
}

// keep track of the new preferred block

// pop a vote off the stack

// get the block that we are going to vote on

// if the block we are going to vote on was already rejected, then
// we should stop applying the votes

// keep track of transitive falters to propagate to this block's
// children

// if the block was previously marked as needing to falter, the block
// should falter before applying the vote

// apply the votes for this snowball instance

// Only accept when you are finalized and a child of the last accepted
// block.

// by accepting the child of parentBlock, the last accepted block is
// no longer voteParentID, but its child. So, voteParentID can be
// removed from the tree.

// If we are on the preferred branch, then the parent's preference is
// the next block on the preferred branch.

// Get the ID of the child that is having a RecordPoll called. All other
// children will need to have their confidence reset. If there isn't a
// child having RecordPoll called, then the nextID will default to the
// nil ID.

// If we are on the preferred branch and the nextID is the preference of
// the snowball instance, then we are following the preferred branch.

// If there wasn't an alpha threshold on the branch (either on this vote
// or a past transitive vote), I should falter now.

// If we don't need to transitively falter and the child is going to
// have RecordPoll called on it, then there is no reason to reset
// the block's confidence

// If we finalized a child of the current block, then all other
// children will have been rejected and removed from the tree.
// Therefore, we need to make sure the child is still in the tree.

// If the child is ever voted for positively, the confidence
// must be reset first.

// Accepts the preferred child of the provided snowman block. By accepting the
// preferred child, all other children will be rejected. When these children are
// rejected, all their descendants will be rejected.
//
// We accept a block once its parent's snowball instance has finalized
// with it as the preference.
func (ts *Topological) acceptPreferredChild(ctx context.Context, n *snowmanBlock) error {
	_ = "STUB: not implemented"
	// We are finalizing the block's child, so we need to get the preference
	return nil
}

// Get the child and accept it

// Notify anyone listening that this block was accepted.

// Note that BlockAcceptor.Accept must be called before child.Accept to
// honor Acceptor.Accept's invariant.

// Update the last accepted values to the newly accepted block.

// Remove the decided block from the set of processing IDs, as its status
// now implies its preferredness.

// Because ts.blocks contains the last accepted block, we don't delete the
// block from the blocks map here.

// don't reject the block we just accepted

// Track which blocks have been directly rejected

// reject all the descendants of the blocks we just rejected

// Takes in a list of rejected ids and rejects all descendants of these IDs
func (ts *Topological) rejectTransitively(ctx context.Context, rejected []ids.ID) error {
	_ = "STUB: not implemented"
	// the rejected array is treated as a stack, with the next element at index
	// 0 and the last element at the end of the slice.
	return nil
}

// pop the rejected ID off the stack

// get the rejected node, and remove it from the tree

// add the newly rejected block to the end of the stack

func (ts *Topological) GetParent(id ids.ID) (ids.ID, bool) {
	_ = "STUB: not implemented"
	return *new(ids.ID), false
}
