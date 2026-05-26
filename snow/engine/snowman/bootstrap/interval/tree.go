// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package interval

import (
	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/database"
)

// TODO: Benchmark what degree to use.
const treeDegree = 2

// Tree implements a set of numbers by tracking intervals. It supports adding
// and removing new values. It also allows checking if a value is included in
// the set.
//
// Tree is more space efficient than a map implementation if the values that it
// contains are continuous. The tree takes O(n) space where n is the number of
// continuous ranges that have been inserted into the tree.
//
// Add, Remove, and Contains all run in O(log n) where n is the number of
// continuous ranges that have been inserted into the tree.
type Tree struct {
	knownHeights *btree.BTreeG[*Interval]
	// If knownHeights contains the full range [0, MaxUint64], then
	// numKnownHeights overflows to 0.
	numKnownHeights uint64
}

// NewTree creates a new interval tree from the provided database.
//
// It is assumed that persisted intervals are non-overlapping. Providing a
// database with overlapping intervals will result in undefined behavior of the
// structure.
func NewTree(db database.Iteratee) (*Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) Add(db database.KeyValueWriterDeleter, height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// height is already in the tree

// the upper and lower ranges should be merged

// the upper range should be extended by one on the lower side

// the lower range should be extended by one on the upper side

func (t *Tree) Remove(db database.KeyValueWriterDeleter, height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// height isn't in the tree

func (t *Tree) Contains(height uint64) bool { _ = "STUB: not implemented"; return false }

func (t *Tree) Flatten() []*Interval { _ = "STUB: not implemented"; return nil }

// Len returns the number of heights in the tree; not the number of intervals.
//
// Because Len returns a uint64 and is describing the number of values in the
// range of uint64s, it will return 0 if the tree contains the full interval
// [0, MaxUint64].
func (t *Tree) Len() uint64 { _ = "STUB: not implemented"; return 0 }
