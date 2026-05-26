// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/buffer"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

// stores previous trie states
type trieHistory struct {
	// Root ID --> The most recent insert number resulting in [rootID].
	lastChangesInsertNumber map[ids.ID]uint64

	// Maximum number of previous roots/changes to store in [history].
	maxHistoryLen int

	// Contains the history.
	// Sorted by increasing order of insertion.
	// Contains at most [maxHistoryLen] values.
	history buffer.Deque[*changeSummaryAndInsertNumber]
}

// Tracks the beginning and ending state of a value.
type change[T any] struct {
	before T
	after  T
}

// Wrapper around a changeSummary that allows comparison
// of when the change was made.
type changeSummaryAndInsertNumber struct {
	*changeSummary
	// Another changeSummaryAndInsertNumber with a greater
	// [insertNumber] means that change was after this one.
	insertNumber uint64
}

// Tracks all the node and value changes that resulted in the rootID.
type changeSummary struct {
	// The ID of the trie after these changes.
	rootID ids.ID
	// The root before/after this change.
	// Set in [applyValueChanges].
	rootChange change[maybe.Maybe[*node]]
	nodes      map[Key]*change[*node]

	keyChanges map[Key]*change[maybe.Maybe[[]byte]]
	sortedKeys []Key
}

func newChangeSummary(estimatedSize int) *changeSummary { _ = "STUB: not implemented"; return nil }

func newTrieHistory(maxHistoryLookback int) *trieHistory { _ = "STUB: not implemented"; return nil }

func (th *trieHistory) getNextInsertNumber() uint64 { _ = "STUB: not implemented"; return 0 }

func (th *trieHistory) getRootChanges(root ids.ID) (*changeSummaryAndInsertNumber, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// The difference between the last index in [th.history] and the index of [rootChanges].

// The index in [th.history] of the latest change resulting in [root].

type valueChange struct {
	key    Key
	change *change[maybe.Maybe[[]byte]]
}

// Returns up to [maxLength] sorted changes with keys in
// [start, end] that occurred between [startRoot] and [endRoot].
// If [start] is Nothing, there's no lower bound on the range.
// If [end] is Nothing, there's no upper bound on the range.
// Returns [sync.ErrInsufficientHistory] if the history is insufficient
// to generate the proof.
// Returns [sync.ErrNoEndRoot], if the history doesn't contain the [endRootID].
func (th *trieHistory) getValueChanges(
	startRoot ids.ID,
	endRoot ids.ID,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	maxLength int,
) ([]valueChange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [endRootChanges] is the last change in the history resulting in [endRoot].

// Confirm there's a change resulting in [startRoot] before
// a change resulting in [endRoot] in the history.
// [startRootChanges] is the last appearance of [startRoot].

// The insert number of the last element in [th.history].

// The index within [th.history] of its last element.

// The difference between the last index in [th.history] and the index of [endRootChanges].

// The index in [th.history] of the latest change resulting in [endRoot].

// [startRootChanges] happened after [endRootChanges].
// However, that is just the *latest* change resulting in [startRoot].
// Attempt to find a change resulting in [startRoot] before [endRootChanges].
//
// Translate the insert number to the index in [th.history] so we can iterate
// backward from [endRootChanges].

// [startRootChanges] is now the last change resulting in
// [startRoot] before [endRootChanges].

// historyChangesIndex is used for tracking keyChanges index from each historical root.

// historyChangesIndexHeap is used to traverse the changes sorted by ASC [key] and ASC [insertNumber].

// For each element in the history in the range between [startRoot]'s
// last appearance (exclusive) and [endRoot]'s last appearance (inclusive),
// add the changes to keys in [start, end] to [combinedChanges].
// Only the key-value pairs with the greatest [maxLength] keys will be kept.
// The difference between the index of [startRootChanges] and [endRootChanges] in [th.history].

// The index of the last change resulting in [startRoot]
// which occurs before [endRootChanges].

// Push in the heap first key in [startKey, endKey] for each historical root.

// Binary search for [startKey] index, or the index where [startKey] would appear.

// [startKey] is after last key of [sortedKeyChanges].

// [keyChange] is after [endKey].

// [startKeyIndex] is the index of the first key in [startKey, endKey] from [sortedKeyChanges].

// Used for combining the changes of all the historical changes, for the current smallest key.

// Skip processing the current [historyRootChanges] if we are after [endKey].

// If there are remaining changes in the current [historyRootChanges], push to minheap.

// Same key, update [after] value.

// New key

// Add the last [currentKeyChange] to [combinedKeyChanges] if there is an actual change.

// If we have [maxLength] changes, we can return the current [combinedKeyChanges].

// Add the last [currentKeyChange] to [combinedKeyChanges] if there is an actual change.

// Returns the changes to go from the current trie state back to the requested [rootID]
// for the keys in [start, end].
// If [start] is Nothing, all keys are considered > [start].
// If [end] is Nothing, all keys are considered < [end].
func (th *trieHistory) getChangesToGetToRoot(rootID ids.ID, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte]) (*changeSummary, error) {
	_ = "STUB: not implemented"
	// [lastRootChange] is the last change in the history resulting in [rootID].
	return nil, nil
}

// Go backward from the most recent change in the history up to but
// not including the last change resulting in [rootID].
// Record each change in [combinedChanges].

// Binary search for [startKey] index.

// Update existing [after] with current [before]

// Remove changed key, if there is a no-op.

// record the provided set of changes in the history
func (th *trieHistory) record(changes *changeSummary) {
	_ = "STUB: not implemented"
	// we aren't recording history so noop
	return
}

// This change causes us to go over our lookback limit.
// Remove the oldest set of changes.

// The removed change was the most recent resulting in this root ID.
// (Note: this if is for situations when the same root could appear twice in history)

// Add [changes] to the sorted change list.

// Mark that this is the most recent change resulting in [changes.rootID].
