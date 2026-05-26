// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package firewood

import (
	"errors"

	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/triedb/database"
)

var (
	_ state.Trie = (*reconstructedAccountTrie)(nil)

	errNilReconstructed = errors.New("nil Reconstructed")
)

// reconstructedAccountTrie implements [state.Trie] backed by an [ffi.Reconstructed] view.
// Like [accountTrie], it accumulates BatchOps from writes. Unlike [accountTrie],
// Hash() chains Reconstruct() calls instead of creating proposals.
//
// If computeRootOnHash is true, Hash and Commit apply pending writes and then
// compute the updated reconstructed root. Otherwise, they apply pending writes
// but return the previously cached root, avoiding the expensive root computation
// when the caller only needs Hash as a state-flush point and validates the final
// root separately.
//
// Not concurrent-safe (matching Reconstructed's guarantees).
type reconstructedAccountTrie struct {
	baseTrie
	recon             *ffi.Reconstructed
	computeRootOnHash bool
}

// newReconstructedAccountTrie creates a new reconstructed account trie.
// The caller retains ownership of the [ffi.Reconstructed] handle and must ensure
// it outlives the trie.
// computeRootOnHash controls whether Hash and Commit update the cached root.
func newReconstructedAccountTrie(recon *ffi.Reconstructed, computeRootOnHash bool) (*reconstructedAccountTrie, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hash applies the accumulated ops to the reconstructed trie and returns a root.
// If computeRootOnHash is true, the cached root is refreshed from the
// reconstructed trie after applying the ops. If computeRootOnHash is false, the
// ops are applied but the cached root is returned unchanged, avoiding root
// computation until the caller validates the final root separately.
// If there are no changes since the last call, the cached root is returned.
// On error, the zero hash is returned.
func (r *reconstructedAccountTrie) Hash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (r *reconstructedAccountTrie) hash() (common.Hash, error) {
	_ = "STUB: not implemented"

	// Reconstruct() mutates the receiver in place with the new state.
	return *new(common.Hash), nil
}

// Unlike accountTrie, updateOps must be cleared because Reconstruct()
// is incremental (mutates in place), whereas createProposals() replays
// all ops from the parent root each time.

// Commit applies the accumulated ops to the reconstructed trie and returns a
// root with an empty [trienode.NodeSet]. If computeRootOnHash is true, the
// cached root is refreshed from the reconstructed trie after applying the ops.
// If computeRootOnHash is false, the ops are applied but the cached root is
// returned unchanged. No persistence occurs; reconstructed views exist only in
// memory and are not committed to the Firewood database.
func (r *reconstructedAccountTrie) Commit(bool) (common.Hash, *trienode.NodeSet, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

var _ database.Reader = (*reconstructedReader)(nil)

// reconstructedReader adapts an [ffi.Reconstructed] to the [database.Reader] interface.
// The underlying [ffi.Reconstructed] may be mutated by Reconstruct() calls, which
// changes what Get() returns.
type reconstructedReader struct {
	reconstructed *ffi.Reconstructed
}

// Node retrieves the value at the given path from the reconstructed view.
func (r *reconstructedReader) Node(_ common.Hash, path []byte, _ common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
