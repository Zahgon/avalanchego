// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

// Representation of a node stored in the database.
type dbNode struct {
	value    maybe.Maybe[[]byte]
	children map[byte]*child
}

type child struct {
	compressedKey Key
	id            ids.ID
	hasValue      bool
}

// node holds additional information on top of the dbNode that makes calculations easier to do
type node struct {
	dbNode
	key         Key
	valueDigest maybe.Maybe[[]byte]
}

// Returns a new node with the given [key] and no value.
func newNode(key Key) *node { _ = "STUB: not implemented"; return nil }

// Parse [nodeBytes] to a node and set its key to [key].
func parseNode(hasher Hasher, key Key, nodeBytes []byte) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns true iff this node has a value.
func (n *node) hasValue() bool { _ = "STUB: not implemented"; return false }

// Returns the byte representation of this node.
func (n *node) bytes() []byte { _ = "STUB: not implemented"; return nil }

// Set [n]'s value to [val].
func (n *node) setValue(hasher Hasher, val maybe.Maybe[[]byte]) { _ = "STUB: not implemented"; return }

func (n *node) setValueDigest(hasher Hasher) { _ = "STUB: not implemented"; return }

// Adds [child] as a child of [n].
// Assumes [child]'s key is valid as a child of [n].
// That is, [n.key] is a prefix of [child.key].
func (n *node) addChild(childNode *node, tokenSize int) { _ = "STUB: not implemented"; return }

func (n *node) addChildWithID(childNode *node, tokenSize int, childID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// Adds a child to [n] without a reference to the child node.
func (n *node) setChildEntry(index byte, childEntry *child) { _ = "STUB: not implemented"; return }

// Removes [child] from [n]'s children.
func (n *node) removeChild(child *node, tokenSize int) { _ = "STUB: not implemented"; return }

// clone Returns a copy of [n].
// Note: value isn't cloned because it is never edited, only overwritten
// if this ever changes, value will need to be copied as well
// it is safe to clone all fields because they are only written/read while one or both of the db locks are held
func (n *node) clone() *node { _ = "STUB: not implemented"; return nil }

// Returns the ProofNode representation of this node.
func (n *node) asProofNode() ProofNode { _ = "STUB: not implemented"; return *new(ProofNode) }
