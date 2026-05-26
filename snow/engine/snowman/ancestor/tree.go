// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ancestor

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ Tree = (*tree)(nil)

// Tree manages a (potentially partial) view of a tree.
//
// For example, assume this is the full tree:
//
//		A
//	  /	  \
//	B		D
//	|		|
//	C		E
//
// A partial view of this tree may be:
//
//		A
//	  /
//	B		D
//	|		|
//	C		E
//
// Or:
//
//	B		D
//	|		|
//	C		E
//
// This structure is designed to update and traverse these partial views.
type Tree interface {
	// Add a mapping from blkID to parentID.
	//
	// Invariant: blkID must not be equal to parentID
	// Invariant: a given blkID must only ever have one parentID
	Add(blkID ids.ID, parentID ids.ID)

	// Has returns if blkID's parentID is known by the tree.
	Has(blkID ids.ID) bool

	// GetAncestor returns the oldest known ancestor of blkID. If there is no
	// known parentID of blkID, blkID will be returned.
	GetAncestor(blkID ids.ID) ids.ID

	// Remove the mapping from blkID to its parentID from the tree.
	Remove(blkID ids.ID)

	// RemoveDescendants removes blkID from the tree along with all of its known
	// descendants.
	RemoveDescendants(blkID ids.ID)

	// Len returns the total number of blkID to parentID mappings that are
	// currently tracked by the tree.
	Len() int
}

type tree struct {
	childToParent    map[ids.ID]ids.ID
	parentToChildren map[ids.ID]set.Set[ids.ID]
}

func NewTree() Tree { _ = "STUB: not implemented"; return *new(Tree) }

func (t *tree) Add(blkID ids.ID, parentID ids.ID) { _ = "STUB: not implemented"; return }

func (t *tree) Has(blkID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (t *tree) GetAncestor(blkID ids.ID) ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// this is the furthest parent available, break loop and return blkID

// continue to loop with parentID

func (t *tree) Remove(blkID ids.ID) { _ = "STUB: not implemented"; return }

// remove blkID from children

// this parent has no more children, remove it from map

func (t *tree) RemoveDescendants(blkID ids.ID) { _ = "STUB: not implemented"; return }

// get children of child

func (t *tree) Len() int { _ = "STUB: not implemented"; return 0 }
