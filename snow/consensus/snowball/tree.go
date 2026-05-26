// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
)

var (
	_ Consensus = (*Tree)(nil)
	_ node      = (*unaryNode)(nil)
	_ node      = (*binaryNode)(nil)
)

func NewTree(factory Factory, params Parameters, choice ids.ID) Consensus {
	_ = "STUB: not implemented"
	return *new(Consensus)
}

// The initial state has no conflicts

// Tree implements the Consensus interface by using a modified patricia tree.
type Tree struct {
	// node is the root that represents the first snow instance in the tree,
	// and contains references to all the other snow instances in the tree.
	node

	// params contains all the configurations of a snow instance
	params Parameters

	// shouldReset is used as an optimization to prevent needless tree
	// traversals. If a snow instance does not get an alpha majority, that
	// instance needs to reset by calling RecordUnsuccessfulPoll. Because the
	// tree splits votes based on the branch, when an instance doesn't get an
	// alpha majority none of the children of this instance can get an alpha
	// majority. To avoid calling RecordUnsuccessfulPoll on the full sub-tree of
	// a node that didn't get an alpha majority, shouldReset is used to indicate
	// that any later traversal into this sub-tree should call
	// RecordUnsuccessfulPoll before performing any other action.
	shouldReset bool

	// factory is used to produce new snow instances as needed
	factory Factory
}

func (t *Tree) Add(choice ids.ID) { _ = "STUB: not implemented"; return }

// Make sure that we haven't already decided against this new id

func (t *Tree) RecordPoll(votes bag.Bag[ids.ID]) bool {
	_ = "STUB: not implemented"
	// Get the assumed decided prefix of the root node.
	return false
}

// If any of the bits differ from the preference in this prefix, the vote is
// for a rejected operation. So, we filter out these invalid votes.

// Now that the votes have been restricted to valid votes, pass them into
// the first snow instance

// Because we just passed the reset into the snow instance, we should no
// longer reset.

func (t *Tree) RecordUnsuccessfulPoll() { _ = "STUB: not implemented"; return }

func (t *Tree) String() string { _ = "STUB: not implemented"; return "" }

type node interface {
	// Preference returns the preferred choice of this sub-tree
	Preference() ids.ID
	// Return the number of assumed decided bits of this node
	DecidedPrefix() int
	// Adds a new choice to vote on
	// Returns the new node
	Add(newChoice ids.ID) node
	// Apply the votes, reset the model if needed
	// Returns the new node and whether the vote was successful
	RecordPoll(votes bag.Bag[ids.ID], shouldReset bool) (newChild node, successful bool)
	// Returns true if consensus has been reached on this node
	Finalized() bool

	Printable() (string, []node)
}

// unary is a node with either no children, or a single child. It handles the
// voting on a range of identical, unary, snow instances.
type unaryNode struct {
	// tree references the tree that contains this node
	tree *Tree

	// preference is the choice that is preferred at every branch in this
	// sub-tree
	preference ids.ID

	// decidedPrefix is the last bit in the prefix that is assumed to be decided
	decidedPrefix int // Will be in the range [0, 255)

	// commonPrefix is the last bit in the prefix that this node transitively
	// references
	commonPrefix int // Will be in the range (decidedPrefix, 256)

	// snow wraps the unary decision logic
	snow Unary

	// shouldReset is used as an optimization to prevent needless tree
	// traversals. It is the continuation of shouldReset in the Tree struct.
	shouldReset bool

	// child is the, possibly nil, node that votes on the next bits in the
	// decision
	child node
}

func (u *unaryNode) Preference() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (u *unaryNode) DecidedPrefix() int { _ = "STUB: not implemented"; return 0 }

// This is by far the most complicated function in this algorithm.
// The intuition is that this instance represents a series of consecutive unary
// snowball instances, and this function's purpose is convert one of these unary
// snowball instances into a binary snowball instance.
// There are 5 possible cases.
//
//  1. None of these instances should be split, we should attempt to split a
//     child
//
//     For example, attempting to insert the value "00001" in this node:
//
//     +-------------------+ <-- This node will not be split
//     |                   |
//     |       0 0 0       |
//     |                   |
//     +-------------------+ <-- Pass the add to the child
//     ^
//     |
//
//     Results in:
//
//     +-------------------+
//     |                   |
//     |       0 0 0       |
//     |                   |
//     +-------------------+ <-- With the modified child
//     ^
//     |
//
//  2. This instance represents a series of only one unary instance and it must
//     be split.
//
//     This will return a binary choice, with one child the same as my child,
//     and another (possibly nil child) representing a new chain to the end of
//     the hash
//
//     For example, attempting to insert the value "1" in this tree:
//
//     +-------------------+
//     |                   |
//     |         0         |
//     |                   |
//     +-------------------+
//
//     Results in:
//
//     +-------------------+
//     |         |         |
//     |    0    |    1    |
//     |         |         |
//     +-------------------+
//
//  3. This instance must be split on the first bit
//
//     This will return a binary choice, with one child equal to this instance
//     with decidedPrefix increased by one, and another representing a new
//     chain to the end of the hash
//
//     For example, attempting to insert the value "10" in this tree:
//
//     +-------------------+
//     |                   |
//     |        0 0        |
//     |                   |
//     +-------------------+
//
//     Results in:
//
//     +-------------------+
//     |         |         |
//     |    0    |    1    |
//     |         |         |
//     +-------------------+
//     ^         ^
//     /           \
//     +-------------------+ +-------------------+
//     |                   | |                   |
//     |         0         | |         0         |
//     |                   | |                   |
//     +-------------------+ +-------------------+
//
//  4. This instance must be split on the last bit
//
//     This will modify this unary choice. The commonPrefix is decreased by
//     one. The child is set to a binary instance that has a child equal to
//     the current child and another child equal to a new unary instance to
//     the end of the hash
//
//     For example, attempting to insert the value "01" in this tree:
//
//     +-------------------+
//     |                   |
//     |        0 0        |
//     |                   |
//     +-------------------+
//
//     Results in:
//
//     +-------------------+
//     |                   |
//     |         0         |
//     |                   |
//     +-------------------+
//     ^
//     |
//     +-------------------+
//     |         |         |
//     |    0    |    1    |
//     |         |         |
//     +-------------------+
//
//  5. This instance must be split on an interior bit
//
//     This will modify this unary choice. The commonPrefix is set to the
//     interior bit. The child is set to a binary instance that has a child
//     equal to this unary choice with the decidedPrefix equal to the interior
//     bit and another child equal to a new unary instance to the end of the
//     hash
//
//     For example, attempting to insert the value "010" in this tree:
//
//     +-------------------+
//     |                   |
//     |       0 0 0       |
//     |                   |
//     +-------------------+
//
//     Results in:
//
//     +-------------------+
//     |                   |
//     |         0         |
//     |                   |
//     +-------------------+
//     ^
//     |
//     +-------------------+
//     |         |         |
//     |    0    |    1    |
//     |         |         |
//     +-------------------+
//     ^         ^
//     /           \
//     +-------------------+ +-------------------+
//     |                   | |                   |
//     |         0         | |         0         |
//     |                   | |                   |
//     +-------------------+ +-------------------+
//
//nolint:gci,gofmt,gofumpt // this comment is formatted as intended
func (u *unaryNode) Add(newChoice ids.ID) node { _ = "STUB: not implemented"; return *new(node) }

// Only happens if the tree is finalized, or it's a leaf node

// If the first difference doesn't exist, then this node shouldn't be
// split

// Because this node will finalize before any children could
// finalize, it must be that the newChoice will match my child's
// prefix

// if u.child is nil, then we are attempting to add the same choice into
// the tree, which should be a noop

// The difference was found, so this node must be split
// The currently preferred bit

// The new child assumes this branch has decided in its favor
// The new child has no conflicts under this branch

// This node was only voting over one bit. (Case 2. from above)

// This node was split on the first bit. (Case 3. from above)

// This node was split on the last bit. (Case 4. from above)

// This node was split on an interior bit. (Case 5. from above)

func (u *unaryNode) RecordPoll(votes bag.Bag[ids.ID], reset bool) (node, bool) {
	_ = "STUB: not implemented"
	// We are guaranteed that the votes are of IDs that have previously been
	// added. This ensures that the provided votes all have the same bits in the
	// range [u.decidedPrefix, u.commonPrefix) as in u.preference.
	return *new(node), false
}

// If my parent didn't get enough votes previously, then neither did I

// Make sure my child is also reset correctly

// We are guaranteed that u.commonPrefix will equal
// u.child.DecidedPrefix(). Otherwise, there must have been a
// decision under this node, which isn't possible because
// beta1 <= beta2. That means that filtering the votes between
// u.commonPrefix and u.child.DecidedPrefix() would always result in
// the same set being returned.

// If I'm now decided, return my child

// The child's preference may have changed

// Now that I have passed my votes to my child, I don't need to reset
// them

func (u *unaryNode) Finalized() bool { _ = "STUB: not implemented"; return false }

func (u *unaryNode) Printable() (string, []node) { _ = "STUB: not implemented"; return "", nil }

// binaryNode is a node with either no children, or two children. It handles the
// voting of a single, binary, snow instance.
type binaryNode struct {
	// tree references the tree that contains this node
	tree *Tree

	// preferences are the choices that are preferred at every branch in their
	// sub-tree
	preferences [2]ids.ID

	// bit is the index in the id of the choice this node is deciding on
	bit int // Will be in the range [0, 256)

	// snow wraps the binary decision logic
	snow Binary

	// shouldReset is used as an optimization to prevent needless tree
	// traversals. It is the continuation of shouldReset in the Tree struct.
	shouldReset [2]bool

	// children are the, possibly nil, nodes that vote on the next bits in the
	// decision
	children [2]node
}

func (b *binaryNode) Preference() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *binaryNode) DecidedPrefix() int { _ = "STUB: not implemented"; return 0 }

func (b *binaryNode) Add(id ids.ID) node { _ = "STUB: not implemented"; return *new(node) }

// If child is nil, then we are running an instance on the last bit. Finding
// two hashes that are equal up to the last bit would be really cool though.
// Regardless, the case is handled

// If child is nil, then the id has already been added to the tree, so
// nothing should be done
// If the decided prefix isn't matched, then a previous decision has made
// the id that is being added to have already been rejected

func (b *binaryNode) RecordPoll(votes bag.Bag[ids.ID], reset bool) (node, bool) {
	_ = "STUB: not implemented"
	// The list of votes we are passed is split into votes for bit 0 and votes
	// for bit 1
	return *new(node), false
}

// We only care about which bit is set if a successful poll can happen

// 1-bit isn't set here because it is set below anyway

// They didn't get the threshold of votes

// The winning child didn't get enough votes either

// If we are decided here, that means we must have decided due
// to this poll. Therefore, we must have decided on bit.

// We passed the reset down

func (b *binaryNode) Finalized() bool { _ = "STUB: not implemented"; return false }

func (b *binaryNode) Printable() (string, []node) { _ = "STUB: not implemented"; return "", nil }
