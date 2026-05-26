// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import "github.com/ava-labs/avalanchego/ids"

// prefixGroup represents a bunch of IDs (stored in the members field),
// with a bit prefix.
// Each time the prefixGroup is split, it is divided into one or more prefixGroups
// according to the next bit index in the index field.
// Successively splitting prefixGroups yields a graph, with the first prefixGroup as the root.
type prefixGroup struct {
	// the bit index this prefixGroup would be split on by the next invocation of split().
	index int
	// the IDs of the prefixGroup
	members []ids.ID
	// children are prefixGroups that correspond to zero and one being the first bit of their members, respectively.
	children [2]*prefixGroup
}

// longestSharedPrefixes creates a prefixGroup that is the root of a graph
// of prefixGroup vertices.
// When iterating the graph, each prefixGroup vertex represents a shared bit prefix
// of IDs, and the members field contains all IDs from the given idList for which their bit prefix
// matches the prefix field.
func longestSharedPrefixes(idList []ids.ID) *prefixGroup {
	_ = "STUB: not implemented"
	// First thing - de-duplicate all ids that appear twice or more
	return nil
}

// Try to split each prefix group.
// Continue until all prefix groups cannot be split anymore.

// We cannot split this prefix group any longer, as the shared prefix ends in this bifurcation

// Else, there is no bifurcation, so swallow up your descendant

// Become your descendant

func determineDescendant(pg *prefixGroup) *prefixGroup { _ = "STUB: not implemented"; return nil }

// If both are nil, it's a programming error, so panic.

// bifurcationsWithCommonPrefix traverses the transitive descendants of this prefix group,
// and applies f() on the block IDs of each prefix group.
// Prefix groups with no descendants are skipped, as they do not represent any prefix.
// Prefix group without a prefix (root prefix group) are also skipped as they do not correspond
// to any instance of snowflake.
func (pg *prefixGroup) bifurcationsWithCommonPrefix(f func([]ids.ID)) {
	_ = "STUB: not implemented"
	return
}

// isBifurcation returns whether this prefixGroup has both zero and one bit descendants.
func (pg *prefixGroup) isBifurcation() bool { _ = "STUB: not implemented"; return false }

// canSplit returns whether this prefixGroup can be split.
func (pg *prefixGroup) canSplit() bool { _ = "STUB: not implemented"; return false }

// traverse invokes f() on this prefixGroup and all descendants in pre-order traversal.
func (pg *prefixGroup) traverse(f func(*prefixGroup)) { _ = "STUB: not implemented"; return }

// split splits the prefixGroup into two prefixGroups according
// to members and the next internal bit.
// All members in the current prefixGroup with bit zero in the next bit index are returned
// in the left result, and similarly for the bit one for the right result.
// Invariant: As long as the current prefixGroup can be split (canSplit() returns true),
// If canSplit() returned true on this prefixGroup, split() will never return (nil, nil),
// since it has at least two members, which means they either differ in the next bit index,
// in which case two prefixGroups would be returned, and otherwise they do not differ
// in the next bit, and then at least one prefixGroup would be returned.
func (pg *prefixGroup) split() { _ = "STUB: not implemented"; return }

// Split members according to their next bit

func deduplicate(in []ids.ID) []ids.ID { _ = "STUB: not implemented"; return nil }
