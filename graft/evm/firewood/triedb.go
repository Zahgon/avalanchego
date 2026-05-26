// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package firewood

import (
	"errors"
	"sync"

	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/stateconf"
	"github.com/ava-labs/libevm/metrics"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/trie/triestate"
	"github.com/ava-labs/libevm/triedb"
	"github.com/ava-labs/libevm/triedb/database"
)

const firewoodDir = "firewood"

var (
	_ triedb.DBConstructor = TrieDBConfig{}.BackendConstructor
	_ triedb.DBOverride    = (*TrieDB)(nil)

	hashCount              = metrics.GetOrRegisterCounter("firewood/triedb/hash/count", nil)
	hashTimer              = metrics.GetOrRegisterCounter("firewood/triedb/hash/time", nil)
	commitCount            = metrics.GetOrRegisterCounter("firewood/triedb/commit/count", nil)
	commitTimer            = metrics.GetOrRegisterCounter("firewood/triedb/commit/time", nil)
	proposeOnDiskCount     = metrics.GetOrRegisterCounter("firewood/triedb/propose/disk/count", nil)
	proposeOnProposeCount  = metrics.GetOrRegisterCounter("firewood/triedb/propose/proposal/count", nil)
	explicitlyDroppedCount = metrics.GetOrRegisterCounter("firewood/triedb/drop/count", nil)

	ErrNoRevisionFound = errors.New("no revision found")

	errNoProposalFound         = errors.New("no proposal found")
	errUnexpectedProposalFound = errors.New("unexpected proposal found")
)

// TrieDB is a triedb.DBOverride implementation backed by Firewood.
// It acts as a HashDB for backwards compatibility with most of the blockchain code.
type TrieDB struct {
	proposals

	// The underlying Firewood database, used for storing proposals and revisions.
	// This is exported as read-only, with knowledge that the consumer will not close it
	// and the latest state can be modified at any time.
	Firewood *ffi.Database
}

type proposals struct {
	sync.Mutex

	byStateRoot map[common.Hash][]*proposal
	// The proposal tree tracks the structure of the current proposals, and which proposals are children of which.
	// This is used to ensure that we can dereference proposals correctly and commit the correct ones
	// in the case of duplicate state roots.
	// The root of the tree is stored here, and represents the top-most layer on disk.
	tree *proposal

	// possible temporarily holds proposals created during a trie update.
	// This is cleared after the update is complete and the proposals have been sent to the database.
	// It's unexpected for multiple updates to this to occur simultaneously, but a lock is used to ensure safety.
	possible map[possibleKey]*proposal
}

type possibleKey struct {
	parentBlockHash, root common.Hash //nolint:unused // It is used as a map key
}

// A proposal carries a Firewood FFI proposal (i.e. Rust-owned memory).
// The Firewood library adds a finalizer to the proposal handle to ensure that
// the memory is freed when the Go object is garbage collected. However, because
// we form a tree of proposals, the `proposal.Proposal` field may be the only
// reference to a given proposal. To ensure that all proposals in the tree
// can be freed in a finalizer, this cannot be included in the tree structure.
type proposal struct {
	*proposalMeta
	handle *ffi.Proposal
}

type proposalMeta struct {
	parent      *proposalMeta
	children    []*proposalMeta
	blockHashes map[common.Hash]struct{} // All corresponding block hashes
	root        common.Hash
	height      uint64
}

type TrieDBConfig struct {
	DatabaseDir       string
	CacheSizeBytes    uint
	RevisionsInMemory uint // must be >= 2
	CacheStrategy     ffi.CacheStrategy
	Archive           bool
	// DeferredCommitInterval must be < RevisionsInMemory as otherwise, it's
	// possible to reap the latest persisted revision.
	DeferredCommitInterval uint64
}

// DefaultConfig returns a sensible TrieDBConfig with the given directory.
// The default config is:
//   - CacheSizeBytes: 1MB
//   - RevisionsInMemory: 128
//   - CacheStrategy: [ffi.CacheAllReads]
//   - DeferredCommitInterval: 64
func DefaultConfig(dir string) TrieDBConfig { _ = "STUB: not implemented"; return *new(TrieDBConfig) }

// 1MB

// BackendConstructor implements the [triedb.DBConstructor] interface.
// It creates a new Firewood database with the given configuration.
// Any error during creation will cause the program to exit.
func (c TrieDBConfig) BackendConstructor(ethdb.Database) triedb.DBOverride {
	_ = "STUB: not implemented"
	return *new(triedb.DBOverride)
}

// New creates a new Firewood database with the given configuration.
// The database will not be opened on error.
//
// If config.DeferredCommitInterval >= config.RevisionsInMemory, then
// config.DeferredCommitInterval is set to config.RevisionsInMemory - 1 to
// uphold the invariant that DeferredCommitInterval < RevisionsInMemory.
func New(config TrieDBConfig) (*TrieDB, error) { _ = "STUB: not implemented"; return nil, nil }

// The Firewood constructor will check that commitCount is nonzero.

// validateDir ensures that the given directory exists and is a directory.
func validateDir(dir string) error { _ = "STUB: not implemented"; return nil }

// SetHashAndHeight sets the committed block hashes and height in memory.
// This must be called at startup to initialize the in-memory state if the
// database is non-empty (e.g. restart, state sync)
func (t *TrieDB) SetHashAndHeight(blockHash common.Hash, height uint64) {
	_ = "STUB: not implemented"
	return
}

// ClearAll resets the database to an empty state, without a genesis block committed.
func (t *TrieDB) ClearAll() error { _ = "STUB: not implemented"; return nil }

// Scheme returns the scheme of the database.
// However, to avoid a slow deletion in `libevm` `StateDB`, it returns [rawdb.HashScheme].
func (*TrieDB) Scheme() string { _ = "STUB: not implemented"; return "" }

// Initialized checks whether a non-empty genesis block has been written.
func (t *TrieDB) Initialized(common.Hash) bool { _ = "STUB: not implemented"; return false }

// Size is a no-op because Firewood does not track storage size.
// All memory management is handled internally by Firewood.
func (*TrieDB) Size() (common.StorageSize, common.StorageSize) {
	_ = "STUB: not implemented"

	// Reference is no-op because proposals are only referenced when created.
	// Additionally, internal nodes do not need tracked by consumers.
	return *new(common.StorageSize), *new(common.StorageSize)
}

func (*TrieDB) Reference(common.Hash, common.Hash) {
	_ = "STUB: not implemented"

	// Dereference is no-op because proposals will be removed automatically.
	// Additionally, internal nodes do not need tracked by consumers.
	return
}

func (*TrieDB) Dereference(common.Hash) {
	_ = "STUB: not implemented"

	// Cap is a no-op because it isn't supported by Firewood.
	return
}

func (*TrieDB) Cap(common.StorageSize) error {
	_ = "STUB: not implemented"

	// Close closes the database, freeing all associated resources.
	// This may hang for a short period while waiting for finalizers to complete.
	// If it does not close as expected, this indicates that there are still references
	// to proposals or revisions in memory, and an error will be returned.
	// The database should not be used after calling Close, but it is safe to call multiple times.
	return nil
}

func (t *TrieDB) Close() error { _ = "STUB: not implemented"; return nil }

// already closed

// All remaining proposals can explicitly be dropped.

// encourage finalizers to run before we wait, otherwise the database won't close properly.

// We must provide a context to close since it may hang while waiting for the finalizers to complete.

// Update updates the database to the given root at the given height.
// The parent block hash and block hash must be provided in the options.
// A proposal must have already been created from [accountTrie.Commit] with the same root,
// parent root, and height.
// If no such proposal exists, an error will be returned.
//
// Unlike for HashDB and PathDB, `Commit` must be called even if if the root is unchanged.
func (t *TrieDB) Update(root, parent common.Hash, height uint64, _ *trienode.MergedNodeSet, _ *triestate.Set, opts ...stateconf.TrieDBUpdateOption) error {
	_ = "STUB: not implemented"
	// We require block hashes to be provided for all blocks in production.
	// However, many tests cannot reasonably provide a block blockHash for genesis, so we allow it to be omitted.
	return nil
}

// The rest of the operations except key-value arranging must occur with a lock

// It's possible that we are committing a proposal on top of an empty genesis block in testing.
// In this case, we can still find the proposal by looking for the empty block hash

// If we have already created an identical proposal, we can skip adding it again.

// All unused proposals can be cleared, since we are already tracking an identical one.

// Track the proposal context in the tree and map.

// Now, all unused proposals have no other references, since we didn't store them
// in the proposal map or tree, so they will be garbage collected.
// Any proposals with a different root were mistakenly created, so they can be freed as well.

// Check if this proposal already exists.
// During reorgs, we may have already tracked this block hash.
// Additionally, we may have coincidentally created an identical proposal with a different block hash.
func (ps *proposals) exists(root, block, parentBlock common.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

// If the block hash is already tracked, we can skip proposing this again.

// We have an identical proposal, but should ensure the hash is tracked with this proposal.

// Commit persists a proposal as a revision to the database.
//
// Any time this is called, we expect either:
//  1. The root is the same as the current root of the database (empty block during bootstrapping)
//  2. We have created a valid propsal with that root, and it is of height +1 above the proposal tree root.
//     Additionally, this will be unique.
//
// Afterward, we know that no other proposal at this height can be committed, so we can dereference all
// children in the the other branches of the proposal tree.
//
// Unlike for HashDB and PathDB, `Commit` must be called even if if the root is unchanged.
func (t *TrieDB) Commit(root common.Hash, report bool) error { _ = "STUB: not implemented"; return nil }

// The proposal has been committed.

// On success, we should remove all children of the committed proposal.
// They will never be committed.

func (ps *proposals) findProposalToCommitWhenLocked(root common.Hash) (*proposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This should never happen, as we ensure that we don't create duplicate proposals in `propose`.

// createProposal creates a new proposal from the given layer
func (t *TrieDB) createProposal(parent *proposal, ops []ffi.BatchOp) (*proposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Edge case: we know the genesis block has an empty parent hash.

// cleanupCommittedProposal dereferences the proposal and removes it from the proposal map.
// It also recursively dereferences all children of the proposal.
func (ps *proposals) cleanupCommittedProposal(p *proposal) { _ = "STUB: not implemented"; return }

// Since this propose has been committed, it doesn't need dropped.

// Internally removes all references of the proposal from the database.
// Frees the associated Rust memory for the proposal and all its children.
// Should only be accessed with the proposal lock held.
func (ps *proposals) removeProposalAndChildren(p *proposalMeta) { _ = "STUB: not implemented"; return }

// removeProposalFromMap removes the proposal from the state root map.
// The proposal lock must be held when calling this function.
// The Rust memory is explicitly freed if drop is true.
func (ps *proposals) removeProposalFromMap(meta *proposalMeta, drop bool) {
	_ = "STUB: not implemented"
	return
}

// pointer comparison - guaranteed to be unique

// createProposals calculates the hash if the set of keys and values are
// proposed from the given parent root.
// All proposals created will be tracked for future use.
func (t *TrieDB) createProposals(parentRoot common.Hash, ops []ffi.BatchOp) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Must prevent a simultaneous `Commit`, as it alters the proposal tree/disk state.

// number of proposals created.
// The resulting root hash, should match between proposals

// Propose from the database root.

// Find any proposal with the given parent root.
// Since we are only using the proposal to find the root hash,
// we can use the first proposal found.

// This should never occur, as to process a block, there must be a revision to read from.

// Reader retrieves a node reader belonging to the given state root.
// An error will be returned if the requested state is not available.
func (t *TrieDB) Reader(root common.Hash) (database.Reader, error) {
	_ = "STUB: not implemented"
	return *new(database.Reader), nil
}

// reader is a state reader of Database which implements the Reader interface.
type reader struct {
	revision *ffi.Revision
}

// Node retrieves the trie node with the given node hash. No error will be
// returned if the node is not found.
func (r *reader) Node(_ common.Hash, path []byte, _ common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
