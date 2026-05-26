// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

const (
	boolLen   = 1
	trueByte  = 1
	falseByte = 0
)

var (
	trueBytes  = []byte{trueByte}
	falseBytes = []byte{falseByte}

	errChildIndexTooLarge = errors.New("invalid child index. Must be less than branching factor")
	errLeadingZeroes      = errors.New("varint has leading zeroes")
	errInvalidBool        = errors.New("decoded bool is neither true nor false")
	errNonZeroKeyPadding  = errors.New("key partial byte should be padded with 0s")
	errExtraSpace         = errors.New("trailing buffer space")
	errIntOverflow        = errors.New("value overflows int")
	errTooManyChildren    = errors.New("too many children")
)

func childSize(index byte, childEntry *child) int {
	_ = "STUB: not implemented"
	// * index
	// * child ID
	// * child key
	// * bool indicating whether the child has a value
	return 0
}

// based on the implementation of encodeUint which uses binary.PutUvarint
func uintSize(value uint64) int { _ = "STUB: not implemented"; return 0 }

func keySize(p Key) int { _ = "STUB: not implemented"; return 0 }

// Assumes [n] is non-nil.
func encodedDBNodeSize(n *dbNode) int {
	_ = "STUB: not implemented"
	// * number of children
	// * bool indicating whether [n] has a value
	// * the value (optional)
	// * children
	return 0
}

// for each non-nil entry, we add the additional size of the child entry

// Assumes [n] is non-nil.
func encodeDBNode(n *dbNode) []byte { _ = "STUB: not implemented"; return nil }

// Avoid allocating keys entirely if the node doesn't have any children.

// By allocating BranchFactorLargest rather than [numChildren], this slice
// is allocated on the stack rather than the heap. BranchFactorLargest is
// at least [numChildren] which avoids memory allocations.

// Ensure that the order of entries is correct.

func encodeKey(key Key) []byte { _ = "STUB: not implemented"; return nil }

type codecWriter struct {
	b []byte
}

func (w *codecWriter) Bool(v bool) { _ = "STUB: not implemented"; return }

func (w *codecWriter) Uvarint(v uint64) { _ = "STUB: not implemented"; return }

func (w *codecWriter) ID(v ids.ID) { _ = "STUB: not implemented"; return }

func (w *codecWriter) Bytes(v []byte) { _ = "STUB: not implemented"; return }

func (w *codecWriter) MaybeBytes(v maybe.Maybe[[]byte]) { _ = "STUB: not implemented"; return }

func (w *codecWriter) Key(v Key) { _ = "STUB: not implemented"; return }

// Assumes [n] is non-nil.
func decodeDBNode(b []byte, n *dbNode) error { _ = "STUB: not implemented"; return nil }

func decodeKey(b []byte) (Key, error) { _ = "STUB: not implemented"; return *new(Key), nil }

type codecReader struct {
	b []byte
	// copy is used to flag to the reader if it is required to copy references
	// to [b].
	copy bool
}

func (r *codecReader) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *codecReader) Uvarint() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// To ensure decoding is canonical, we check for leading zeroes in the
// varint.
// The last byte of the varint includes the most significant bits.
// If the last byte is 0, then the number should have been encoded more
// efficiently by removing this leading zero.

func (r *codecReader) ID() (ids.ID, error) { _ = "STUB: not implemented"; return *new(ids.ID), nil }

func (r *codecReader) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *codecReader) MaybeBytes() (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *codecReader) Key() (Key, error) { _ = "STUB: not implemented"; return *new(Key), nil }

// Confirm that the padding bits in the partial byte are 0.
// We want to only look at the bits to the right of the last token,
// which is at index length-1.
// Generate a mask where the (result.length % 8) left bits are 0.
