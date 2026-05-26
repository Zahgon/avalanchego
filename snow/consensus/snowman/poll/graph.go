// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
)

// voteVertex encapsulates an ID and points to its parent and descendants.
type voteVertex struct {
	id          ids.ID
	parent      *voteVertex
	descendants []*voteVertex
}

// traverse invokes f() on this voteVertex and all its descendants in pre-order traversal.
func (v *voteVertex) traverse(f func(*voteVertex)) { _ = "STUB: not implemented"; return }

// voteGraph is a collection of vote vertices.
// it must be initialized via buildVoteGraph() and not explicitly.
type voteGraph struct {
	vertexCount int
	leaves      []*voteVertex
	roots       []*voteVertex
}

// buildVoteGraph receives as input a function that returns the ID of a block, or Empty if unknown,
// as well as a bag of IDs (The bag is for enforcing uniqueness among the IDs in contrast to a list).
// It returns a voteGraph where each vertex corresponds to an ID and is linked to vertices
// according to what getParent() returns for each ID.
func buildVoteGraph(getParent func(ids.ID) (ids.ID, bool), votes bag.Bag[ids.ID]) voteGraph {
	_ = "STUB: not implemented"
	// Build a graph out of the vertices that correspond to the IDs of the votes.
	return *new(voteGraph)
}

// This case is typically dead code, as it can only happen if
// getParent returns new mappings that were not returned earlier in
// the loop. But it is handled for completeness.

// addIDAndAncestorsToGraph adds an ID and its ancestors to the graph
func addIDAndAncestorsToGraph(
	getParent func(ids.ID) (ids.ID, bool),
	id ids.ID,
	id2Vertex map[ids.ID]*voteVertex,
) {
	_ = "STUB: not implemented"

	// If the ID has already been added, no need to add it or its ancestors,
	// as we already did so in previous iterations.
	return
}

// Add the ID to the graph.

// Attempt to add the parent of the ID to the graph.

// If this parent isn't found, it must be already finalized, so don't
// add it to the graph.

// If the parent is not finalized we can vote on it.

// traverse traverses over all vertices in the voteGraph in pre-order traversal.
func (vg *voteGraph) traverse(f func(*voteVertex)) { _ = "STUB: not implemented"; return }

// topologicalSortTraversal invokes f() on all vote vertices in the voteGraph according to
// the topological order of the vertices.
func (vg *voteGraph) topologicalSortTraversal(f func(*voteVertex)) {
	_ = "STUB: not implemented"
	// We hold a counter for each vertex
	return
}

// Iterate from leaves to roots and apply f() over each vertex

func findLeaves(idToVertex map[ids.ID]*voteVertex) []*voteVertex {
	_ = "STUB: not implemented"
	return nil
}

// computeTransitiveVoteCountGraph receives a vote graph and corresponding votes for each vertex ID.
// Returns a new bag where element represents the number of votes for transitive descendents in the graph.
func computeTransitiveVoteCountGraph(graph *voteGraph, votes bag.Bag[ids.ID]) bag.Bag[ids.ID] {
	_ = "STUB: not implemented"
	return nil
}

// Traverse from the leaves to the roots and recursively add the number of votes of descendents to each parent.
