// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"github.com/ava-labs/avalanchego/ids"
)

// TODO: Support configurable hash lengths
const HashLength = 32

var (
	SHA256Hasher Hasher = &sha256Hasher{}

	// If a Hasher isn't specified, this package defaults to using the
	// [SHA256Hasher].
	DefaultHasher = SHA256Hasher
)

type Hasher interface {
	// Returns the canonical hash of the non-nil [node].
	HashNode(node *node) ids.ID
	// Returns the canonical hash of [value].
	HashValue(value []byte) ids.ID
}

type sha256Hasher struct{}

// This method is performance critical. It is not expected to perform any memory
// allocations.
func (*sha256Hasher) HashNode(n *node) ids.ID {
	_ = "STUB: not implemented"

	// sha.Write always returns nil, so we ignore its return values.
	return *new(ids.ID)
}

// The hash length is larger than the maximum Uvarint length. This
// ensures binary.AppendUvarint doesn't perform any memory allocations.

// By directly calling sha.Write rather than passing sha around as an
// io.Writer, the compiler can perform sufficient escape analysis to avoid
// allocating buffers on the heap.

// Avoid allocating keys entirely if the node doesn't have any children.

// By allocating BranchFactorLargest rather than [numChildren], this
// slice is allocated on the stack rather than the heap.
// BranchFactorLargest is at least [numChildren] which avoids memory
// allocations.

// Ensure that the order of entries is correct.

// This method is performance critical. It is not expected to perform any memory
// allocations.
func (*sha256Hasher) HashValue(value []byte) ids.ID {
	_ = "STUB: not implemented"

	// sha.Write always returns nil, so we ignore its return values.
	return *new(ids.ID)
}
