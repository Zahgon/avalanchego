// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"context"
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

const (
	initKeyValuesSize        = 256
	defaultPreallocationSize = 100
)

var (
	_ View = (*view)(nil)

	ErrCommitted                  = errors.New("view has been committed")
	ErrInvalid                    = errors.New("the trie this view was based on has changed, rendering this view invalid")
	ErrPartialByteLengthWithValue = errors.New(
		"the underlying db only supports whole number of byte keys, so cannot record changes with partial byte lengths",
	)
	ErrVisitPathToKey         = errors.New("failed to visit expected node during insertion")
	ErrStartAfterEnd          = errors.New("start key > end key")
	ErrNoChanges              = errors.New("no changes provided")
	ErrParentNotDatabase      = errors.New("parent trie is not database")
	ErrNodesAlreadyCalculated = errors.New("cannot modify the trie after the node changes have been calculated")
)

type view struct {
	// If true, this view has been committed.
	// [commitLock] must be held while accessing this field.
	committed  bool
	commitLock sync.RWMutex

	// valueChangesApplied is used to enforce that no changes are made to the
	// trie after the nodes have been calculated
	valueChangesApplied utils.Atomic[bool]

	// applyValueChangesOnce prevents node calculation from occurring multiple
	// times
	applyValueChangesOnce sync.Once

	// Controls the view's validity related fields.
	// Must be held while reading/writing [childViews], [invalidated], and [parentTrie].
	// Only use to lock current view or descendants of the current view
	// DO NOT grab the [validityTrackingLock] of any ancestor trie while this is held.
	validityTrackingLock sync.RWMutex

	// If true, this view has been invalidated and can't be used.
	//
	// Invariant: This view is marked as invalid before any of its ancestors change.
	// Since we ensure that all subviews are marked invalid before making an invalidating change
	// then if we are still valid at the end of the function, then no corrupting changes could have
	// occurred during execution.
	// Namely, if we have a method with:
	//
	// *Code Accessing Ancestor State*
	//
	// if v.isInvalid() {
	//     return ErrInvalid
	//  }
	// return [result]
	//
	// If the invalidated check passes, then we're guaranteed that no ancestor changes occurred
	// during the code that accessed ancestor state and the result of that work is still valid
	//
	// [validityTrackingLock] must be held when reading/writing this field.
	invalidated bool

	// the uncommitted parent trie of this view
	// [validityTrackingLock] must be held when reading/writing this field.
	parentTrie View

	// The valid children of this view.
	// [validityTrackingLock] must be held when reading/writing this field.
	childViews []*view

	// Changes made to this view.
	// May include nodes that haven't been updated
	// but will when their ID is recalculated.
	changes *changeSummary

	db *merkleDB

	// The root of the trie represented by this view.
	root maybe.Maybe[*node]

	tokenSize int
}

// NewView returns a new view on top of this view where the passed changes
// have been applied.
// Adds the new view to [v.childViews].
// Assumes [v.commitLock] isn't held.
func (v *view) NewView(
	ctx context.Context,
	changes ViewChanges,
) (View, error) {
	_ = "STUB: not implemented"
	return *new(View), nil
}

// Creates a new view with the given [parentTrie].
func newView(
	db *merkleDB,
	parentTrie View,
	changes ViewChanges,
) (*view, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func recordValueChange(v *view, keyChanges map[Key]*change[maybe.Maybe[[]byte]], key Key, value maybe.Maybe[[]byte]) error {
	_ = "STUB: not implemented"
	// update the existing change if it exists
	return nil
}

// grab the before value

// Creates a view of the db at a historical root using the provided [changes].
// Returns ErrNoChanges if [changes] is empty.
func newViewWithChanges(
	db *merkleDB,
	changes *changeSummary,
) (*view, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// since this is a set of historical changes, all nodes have already been calculated
// since no new changes have occurred, no new calculations need to be done

func (v *view) getTokenSize() int { _ = "STUB: not implemented"; return 0 }

func (v *view) getRoot() maybe.Maybe[*node] {
	_ = "STUB: not implemented"

	// applyValueChanges generates the node changes from the value changes. It then
	// hashes the changed nodes to calculate the new trie.
	//
	// Cancelling [ctx] doesn't cancel the operation. It's used only for tracing.
	return nil
}

func (v *view) applyValueChanges(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Create the span inside the once wrapper to make traces more useful.
// Otherwise, spans would be created during calls where the IDs are not
// re-calculated.

// Note we're setting [err] defined outside this function.

// ensure no ancestor changes occurred during execution

func (v *view) calculateNodeChanges(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Add all the changed key/values to the nodes of the trie

func (v *view) hashChangedNodes(ctx context.Context) { _ = "STUB: not implemented"; return }

// If there are no children, we can avoid allocating [keyBuffer].

// Allocate [keyBuffer] and populate it with the root node's key.

// Calculates the ID of all descendants of [n] which need to be recalculated,
// and then calculates the ID of [n] itself.
//
// Returns a potentially expanded [keyBuffer]. By returning this value this
// function is able to have a maximum total number of allocations shared across
// multiple invocations.
//
// Invariant: [keyBuffer] must be populated with [n]'s key and have sufficient
// length to contain any of [n]'s child keys.
func (v *view) hashChangedNode(n *node, keyBuffer []byte) (ids.ID, []byte) {
	_ = "STUB: not implemented"

	// childBuffer is allocated on the stack.
	return *new(ids.ID), nil
}

// We track the last byte of [n.key] so that we can reset the value for
// each key. This is needed because the child buffer may get ORed at
// this byte.

// We use [wg] to wait until all descendants of [n] have been updated.

// This loop is optimized to avoid allocations when calculating the
// [childKey] by reusing [keyBuffer] and leaving the first [bytesForKey-1]
// bytes unmodified.

// It is safe to use byteSliceToString because [childBuffer] is not
// modified while [childIndexAsKey] is in use.

// Because [keyBuffer] may have been modified in a prior iteration of
// this loop, it is not guaranteed that its length is at least
// [bytesNeeded(totalBitLength)]. However, that's fine. The below
// slicing would only panic if the buffer didn't have sufficient
// capacity.

// We don't need to copy this node's key. It's assumed to already be
// correct; except for the last byte. We must make sure the last byte of
// the key is set correctly because extendIntoBuffer may OR bits from
// the extension and overwrite the last byte. However, extendIntoBuffer
// does not modify the first [bytesForKey-1] bytes of [keyBuffer].

// It is safe to use byteSliceToString because [keyBuffer] is not
// modified while [childKey] is in use.

// This child wasn't changed.

// If there are no children of the childNode, we can avoid constructing
// the buffer for the child keys.

// Try updating the child and its descendants in a goroutine.

// We're at the goroutine limit; do the work in this goroutine.
//
// We can skip copying the key here because [keyBuffer] is already
// constructed to be childNode's key.

// Wait until all descendants of [n] have been updated.

// The IDs [n]'s descendants are up to date so we can calculate [n]'s ID.

// setKeyBuffer expands [keyBuffer] to have sufficient size for any of [n]'s
// child keys and populates [n]'s key into [keyBuffer]. If [keyBuffer] already
// has sufficient size, this function will not perform any memory allocations.
func (v *view) setKeyBuffer(n *node, keyBuffer []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// setLengthForChildren expands [keyBuffer] to have sufficient size for any of
// [n]'s child keys.
func (v *view) setLengthForChildren(n *node, keyBuffer []byte) []byte {
	_ = "STUB: not implemented"
	// Calculate the size of the largest child key of this node.
	return nil
}

func setBytesLength(b []byte, size int) []byte { _ = "STUB: not implemented"; return nil }

// GetProof returns a proof that [bytesPath] is in or not in trie [t].
func (v *view) GetProof(ctx context.Context, key []byte) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRangeProof returns a range proof for (at least part of) the key range [start, end].
// The returned proof's [KeyValues] has at most [maxLength] values.
// [maxLength] must be > 0.
func (v *view) GetRangeProof(
	ctx context.Context,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	maxLength int,
) (*RangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CommitToDB commits changes from this view to the underlying DB.
func (v *view) CommitToDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Commits the changes from [trieToCommit] to the db.
// Assumes that its parent view has already been committed to the db.
// Assumes [v.db.commitLock] is held.
func (v *view) commitToDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Call this here instead of in [v.db.commitView] because doing so there
// would be a deadlock.

// Assumes [v.validityTrackingLock] isn't held.
func (v *view) isInvalid() bool { _ = "STUB: not implemented"; return false }

// Invalidates this view and all descendants.
// Assumes [v.validityTrackingLock] isn't held.
func (v *view) invalidate() { _ = "STUB: not implemented"; return }

// after invalidating the children, they no longer need to be tracked

func (v *view) updateParent(newParent View) { _ = "STUB: not implemented"; return }

// GetMerkleRoot returns the ID of the root of this view.
func (v *view) GetMerkleRoot(ctx context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (v *view) GetValues(ctx context.Context, keys [][]byte) ([][]byte, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetValue returns the value for the given [key].
// Returns database.ErrNotFound if it doesn't exist.
func (v *view) GetValue(ctx context.Context, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getValueCopy returns a copy of the value for the given [key].
// Returns database.ErrNotFound if it doesn't exist.
func (v *view) getValueCopy(key Key) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *view) getValue(key Key) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// if we don't have local copy of the value, then grab a copy from the parent trie

// ensure no ancestor changes occurred during execution

// Must not be called after [applyValueChanges] has returned.
func (v *view) remove(key Key) error { _ = "STUB: not implemented"; return nil }

// confirm a node exists with a value

// [key] isn't in the trie.

// [key] doesn't have a value.

// if the node exists and contains a value
// mark all ancestor for change
// grab parent and grandparent nodes for path compression

// if the removed node has no children, the node can be removed from the trie

// We deleted the root. The trie is empty now.

// Note [parent] != nil since [nodeToDelete.key] != [v.root.key].
// i.e. There's the root and at least one more node.

// merge the parent node and its child into a single node if possible

// merge this node and its parent into a single node if possible

// Merges [n] with its [parent] if [n] has only one child and no value.
// If [parent] is nil, [n] is the root node and [v.root] is updated to [n].
// Assumes at least one of the following is true:
// * [n] has a value.
// * [n] has children.
// Must not be called after [applyValueChanges] has returned.
func (v *view) compressNodePath(parent, n *node) error { _ = "STUB: not implemented"; return nil }

// We know from above that [n] has no value.
/* hasValue */

// There is only one child, but we don't know the index.
// "Cycle" over the key/values to find the only child.
// Note this iteration once because len(node.children) == 1.

// Get a copy of the node matching the passed key from the view.
// Used by views to get nodes from their ancestors.
func (v *view) getEditableNode(key Key, hadValue bool) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// grab the node in question

// ensure no ancestor changes occurred during execution

// return a clone of the node, so it can be edited without affecting this view

// insert a key/value pair into the correct node of the trie.
// Must not be called after [applyValueChanges] has returned.
func (v *view) insert(
	key Key,
	value maybe.Maybe[[]byte],
) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the trie is empty, so create a new root node.

// Find the node that most closely matches [key].

// Need to recalculate ID for all nodes on path to [key].

// [v.root.key] isn't a prefix of [key].

/*offset*/

// Call addChildWithID instead of addChild so the old root is added
// to the new root with the correct ID.
// TODO:
// [oldRootID] shouldn't need to be calculated here.
// Either oldRootID should already be calculated or will be calculated at the end with the other nodes
// Initialize the v.changes.rootID during newView and then use that here instead of oldRootID

// a node with that exact key already exists so update its value

// closestNode was already marked as changed in the ancestry loop above

// A node with the exact key doesn't exist so determine the portion of the
// key that hasn't been matched yet
// Note that [key] has prefix [closestNode.key], so [key] must be longer
// and the following index won't OOB.

// there are no existing nodes along the key [key], so create a new node to insert [value]

// if we have reached this point, then the [key] we are trying to insert and
// the existing path node have some common prefix.
// a new branching node will be created that will represent this common prefix and
// have the existing path node and the value being inserted as children.

// generate the new branch node
// find how many tokens are common between the existing child's compressed key and
// the current key(offset by the closest node's key),
// then move all the common tokens into the branch node

// Since the compressed key is shorter than the common prefix,
// we should have visited [existingChildEntry] in [visitPathToKey].

// the branch node has exactly the key to be inserted as its key, so set the value on the branch node

// the key to be inserted is a child of the branch node
// create a new node and add the value to it

// add the existing child onto the branch node

func getLengthOfCommonPrefix(first, second Key, secondOffset int, tokenSize int) int {
	_ = "STUB: not implemented"
	return 0
}

// Records that a node has been created.
// Must not be called after [applyValueChanges] has returned.
func (v *view) recordNewNode(after *node) error { _ = "STUB: not implemented"; return nil }

/* newNode */

// Records that an existing node has been changed.
// Must not be called after [applyValueChanges] has returned.
func (v *view) recordNodeChange(after *node) error { _ = "STUB: not implemented"; return nil }

/* newNode */

// Records that the node associated with the given key has been deleted.
// Must not be called after [applyValueChanges] has returned.
func (v *view) recordNodeDeleted(after *node, hadValue bool) error {
	_ = "STUB: not implemented"
	return nil
}

/* newNode */

// Records that the node associated with the given key has been changed.
// If it is an existing node, record what its value was before it was changed.
// Must not be called after [applyValueChanges] has returned.
func (v *view) recordKeyChange(key Key, after *node, hadValue bool, newNode bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieves a node with the given [key].
// If the node is fetched from [v.parentTrie] and [id] isn't empty,
// sets the node's ID to [id].
// If the node is loaded from the baseDB, [hasValue] determines which database the node is stored in.
// Returns database.ErrNotFound if the node doesn't exist.
func (v *view) getNode(key Key, hasValue bool) (*node, error) {
	_ = "STUB: not implemented"
	// check for the key within the changed nodes
	return nil, nil
}

// get the node from the parent trie and store a local copy

// Get the parent trie of the view
func (v *view) getParentTrie() View { _ = "STUB: not implemented"; return *new(View) }
