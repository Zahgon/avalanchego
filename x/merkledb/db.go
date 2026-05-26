// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"context"
	"errors"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/maybe"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"

	merklesync "github.com/ava-labs/avalanchego/database/merkle/sync"
)

const (
	// TODO: name better
	rebuildViewSizeFractionOfCacheSize   = 50
	minRebuildViewSizePerCommit          = 1000
	rebuildIntermediateDeletionWriteSize = units.MiB
	valueNodePrefixLen                   = 1
	cacheEntryOverHead                   = 8
)

var (
	_ MerkleDB                                 = (*merkleDB)(nil)
	_ merklesync.DB[*RangeProof, *ChangeProof] = (MerkleDB)(nil)

	metadataPrefix         = []byte{0}
	valueNodePrefix        = []byte{1}
	intermediateNodePrefix = []byte{2}

	// cleanShutdownKey is used to flag that the database did (or did not)
	// previously shutdown correctly.
	//
	// If this key has value [hadCleanShutdown] it must be true that all
	// intermediate nodes of the trie are correctly populated on disk and that
	// the [rootDBKey] has the correct key for the root node.
	//
	// If this key has value [didNotHaveCleanShutdown] the intermediate nodes of
	// the trie may not be correct and the [rootDBKey] may not exist or point to
	// a node that node longer exists.
	//
	// Regardless of the value of [cleanShutdownKey], the value nodes must
	// always be persisted correctly.
	cleanShutdownKey        = []byte(string(metadataPrefix) + "cleanShutdown")
	rootDBKey               = []byte(string(metadataPrefix) + "root")
	hadCleanShutdown        = []byte{1}
	didNotHaveCleanShutdown = []byte{0}

	errSameRoot    = errors.New("start and end root are the same")
	errTooManyKeys = errors.New("response contains more than requested keys")
)

type ChangeProofer interface {
	// GetChangeProof returns a proof for a subset of the key/value changes in key range
	// [start, end] that occurred between [startRootID] and [endRootID].
	// Returns at most [maxLength] key/value pairs.
	// Returns [xsync.ErrInsufficientHistory] if this node has insufficient history
	// to generate the proof.
	// Returns ErrEmptyProof if [endRootID] is ids.Empty.
	// Note that [endRootID] == ids.Empty means the trie is empty
	// (i.e. we don't need a change proof.)
	// Returns [xsync.ErrNoEndRoot], if the history doesn't contain the [endRootID].
	GetChangeProof(
		ctx context.Context,
		startRootID ids.ID,
		endRootID ids.ID,
		start maybe.Maybe[[]byte],
		end maybe.Maybe[[]byte],
		maxLength int,
	) (*ChangeProof, error)

	// Returns nil iff all the following hold:
	//   - [start] <= [end].
	//   - [proof] is non-empty.
	//   - [proof.KeyChanges] is not longer than [maxLength].
	//   - All keys in [proof.KeyChanges] are in [start, end].
	//     If [start] is nothing, all keys are considered > [start].
	//     If [end] is nothing, all keys are considered < [end].
	//   - [proof.KeyChanges] are sorted in order of increasing key.
	//   - [proof.StartProof] and [proof.EndProof] are well-formed.
	//   - When the changes in [proof.KeyChanges] are applied,
	//     the root ID of the database is [expectedEndRootID].
	VerifyChangeProof(
		ctx context.Context,
		proof *ChangeProof,
		start maybe.Maybe[[]byte],
		end maybe.Maybe[[]byte],
		expectedEndRootID ids.ID,
		maxLength int,
	) error

	// CommitChangeProof commits the key/value pairs within the [proof] to the db.
	// [end] is the largest possible key in the range this [proof] covers.
	// [end] may be Nothing, meaning that there is no upper bound on the
	// range.
	// The key returned indicates the next key after the largest key in
	// the proof. If the database has all keys from the start of the proof
	// until [end], then Nothing is returned.
	CommitChangeProof(ctx context.Context, end maybe.Maybe[[]byte], proof *ChangeProof) (maybe.Maybe[[]byte], error)
}

type RangeProofer interface {
	// GetRangeProofAtRoot returns a proof for the key/value pairs in this trie within the range
	// [start, end] when the root of the trie was [rootID].
	// If [start] is Nothing, there's no lower bound on the range.
	// If [end] is Nothing, there's no upper bound on the range.
	// Returns ErrEmptyProof if [rootID] is ids.Empty.
	// Note that [rootID] == ids.Empty means the trie is empty
	// (i.e. we don't need a range proof.)
	// Returns [xsync.ErrNoEndRoot], if the history doesn't contain the [rootID].
	GetRangeProofAtRoot(
		ctx context.Context,
		rootID ids.ID,
		start maybe.Maybe[[]byte],
		end maybe.Maybe[[]byte],
		maxLength int,
	) (*RangeProof, error)

	// Returns nil iff all the following hold:
	//   - [start] <= [end].
	//   - [proof] is non-empty.
	//   - [proof.KeyChanges] is not longer than [maxLength].
	//   - All keys in [proof.KeyChanges] are in [start, end].
	//     If [start] is nothing, all keys are considered > [start].
	//     If [end] is nothing, all keys are considered < [end].
	//   - [proof.KeyChanges] are sorted in order of increasing key.
	//   - [proof.StartProof] and [proof.EndProof] are well-formed.
	//   - When the changes in [proof.KeyChanges] are applied,
	//     the root ID of the database is [expectedEndRootID].
	VerifyRangeProof(
		ctx context.Context,
		proof *RangeProof,
		start maybe.Maybe[[]byte],
		end maybe.Maybe[[]byte],
		expectedEndRootID ids.ID,
		maxLength int,
	) error

	// CommitRangeProof commits the key/value pairs within the [proof] to the db.
	// [start] is the smallest possible key in the range this [proof] covers.
	// [end] is the largest possible key in the range this [proof] covers.
	// [end] may be Nothing, meaning that there is no upper bound on the
	// range.
	// The key returned indicates the next key after the largest key in
	// the proof. If the database has all keys from the start of the proof
	// until [end], then Nothing is returned.
	CommitRangeProof(ctx context.Context, start, end maybe.Maybe[[]byte], proof *RangeProof) (maybe.Maybe[[]byte], error)
}

type Clearer interface {
	// Deletes all key/value pairs from the database
	// and clears the change history.
	Clear() error
}

type Prefetcher interface {
	// PrefetchPath attempts to load all trie nodes on the path of [key]
	// into the cache.
	PrefetchPath(key []byte) error

	// PrefetchPaths attempts to load all trie nodes on the paths of [keys]
	// into the cache.
	//
	// Using PrefetchPaths can be more efficient than PrefetchPath because
	// the underlying view used to compute each path can be reused.
	PrefetchPaths(keys [][]byte) error
}

type MerkleDB interface {
	database.Database
	Clearer
	View
	MerkleRootGetter
	ProofGetter
	ChangeProofer
	RangeProofer
	Prefetcher
}

func NewConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type Config struct {
	// BranchFactor determines the number of children each node can have.
	BranchFactor BranchFactor

	// Hasher defines the hash function to use when hashing the trie.
	//
	// If not specified, [DefaultHasher] will be used.
	Hasher Hasher

	// RootGenConcurrency is the number of goroutines to use when
	// generating a new state root.
	//
	// If 0 is specified, [runtime.NumCPU] will be used.
	RootGenConcurrency uint

	// The number of changes to the database that we store in memory in order to
	// serve change proofs.
	HistoryLength uint
	// The number of bytes used to cache nodes with values.
	ValueNodeCacheSize uint
	// The number of bytes used to cache nodes without values.
	IntermediateNodeCacheSize uint
	// The number of bytes used to store nodes without values in memory before forcing them onto disk.
	IntermediateWriteBufferSize uint
	// The number of bytes to write to disk when intermediate nodes are evicted
	// from the write buffer and written to disk.
	IntermediateWriteBatchSize uint
	// If [Reg] is nil, metrics are collected locally but not exported through
	// Prometheus.
	// This may be useful for testing.
	Reg        prometheus.Registerer
	Namespace  string
	TraceLevel TraceLevel
	Tracer     trace.Tracer
}

// merkleDB can only be edited by committing changes from a view.
type merkleDB struct {
	// Must be held when reading/writing fields.
	lock sync.RWMutex

	// Must be held when preparing work to be committed to the DB.
	// Used to prevent editing of the trie without restricting read access
	// until the full set of changes is ready to be written.
	// Should be held before taking [db.lock]
	commitLock sync.RWMutex

	// Contains all the key-value pairs stored by this database,
	// including metadata, intermediate nodes and value nodes.
	baseDB database.Database

	valueNodeDB        *valueNodeDB
	intermediateNodeDB *intermediateNodeDB

	// Stores change lists. Used to serve change proofs and construct
	// historical views of the trie.
	history *trieHistory

	// True iff the db has been closed.
	closed bool

	metrics metrics

	debugTracer trace.Tracer
	infoTracer  trace.Tracer

	// The root of this trie.
	// Nothing if the trie is empty.
	root maybe.Maybe[*node]

	rootID ids.ID

	// Valid children of this trie.
	childViews []*view

	// hashNodesKeyPool controls the number of goroutines that are created
	// inside [hashChangedNode] at any given time and provides slices for the
	// keys needed while hashing.
	hashNodesKeyPool *bytesPool

	tokenSize int

	hasher Hasher
}

// New returns a new merkle database.
func New(ctx context.Context, db database.Database, config Config) (MerkleDB, error) {
	_ = "STUB: not implemented"
	return *new(MerkleDB), nil
}

func newDatabase(
	ctx context.Context,
	db database.Database,
	config Config,
	metrics metrics,
) (*merkleDB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Share a bytes pool between the intermediateNodeDB and valueNodeDB to
// reduce memory allocations.

// If the marker wasn't found then the DB is being created for the first
// time and there is nothing to do.

// add current root to history (has no changes)

// mark that the db has not yet been cleanly closed

// Deletes every intermediate node and rebuilds them by re-adding every key/value.
// TODO: make this more efficient by only clearing out the stale portions of the trie.
func (db *merkleDB) rebuild(ctx context.Context, cacheSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete intermediate nodes.

// Add all key-value pairs back into the database.

// ensure valueIt is captured and release gets called on the latest copy of valueIt

// reset the iterator to prevent memory bloat

func (db *merkleDB) CommitChangeProof(ctx context.Context, end maybe.Maybe[[]byte], proof *ChangeProof) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If there's no changes to make, we don't have to commit anything.

// Commit the changes to the DB.

func (db *merkleDB) CommitRangeProof(ctx context.Context, start, end maybe.Maybe[[]byte], proof *RangeProof) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't need to lock [view] because nobody else has a reference to it.

func (db *merkleDB) Compact(start []byte, limit []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *merkleDB) Close() error { _ = "STUB: not implemented"; return nil }

// mark all children as no longer valid because the db has closed

// Flush intermediary nodes to disk.

// Write the root key

// Write the clean shutdown marker

func (db *merkleDB) PrefetchPaths(keys [][]byte) error { _ = "STUB: not implemented"; return nil }

func (db *merkleDB) PrefetchPath(key []byte) error { _ = "STUB: not implemented"; return nil }

func (db *merkleDB) prefetchPath(keyBytes []byte) error { _ = "STUB: not implemented"; return nil }

func (db *merkleDB) Get(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// this is a duplicate because the database interface doesn't support
	// contexts, which are used for tracing
	return nil, nil
}

func (db *merkleDB) GetValues(ctx context.Context, keys [][]byte) ([][]byte, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lock to ensure no commit happens during the reads.

// GetValue returns the value associated with [key].
// Returns database.ErrNotFound if it doesn't exist.
func (db *merkleDB) GetValue(ctx context.Context, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getValueCopy returns a copy of the value for the given [key].
// Returns database.ErrNotFound if it doesn't exist.
// Assumes [db.lock] is read locked.
func (db *merkleDB) getValueCopy(key Key) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getValue returns the value for the given [key].
// Returns database.ErrNotFound if it doesn't exist.
// Assumes [db.lock] isn't held.
func (db *merkleDB) getValue(key Key) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// getValueWithoutLock returns the value for the given [key].
// Returns database.ErrNotFound if it doesn't exist.
// Assumes [db.lock] is read locked.
func (db *merkleDB) getValueWithoutLock(key Key) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* hasValue */

func (db *merkleDB) GetMerkleRoot(ctx context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// Assumes [db.lock] or [db.commitLock] is read locked.
func (db *merkleDB) getMerkleRoot() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (db *merkleDB) GetProof(ctx context.Context, key []byte) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *merkleDB) GetRangeProof(
	ctx context.Context,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	maxLength int,
) (*RangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *merkleDB) GetRangeProofAtRoot(
	ctx context.Context,
	rootID ids.ID,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	maxLength int,
) (*RangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *merkleDB) GetChangeProof(
	ctx context.Context,
	startRootID ids.ID,
	endRootID ids.ID,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	maxLength int,
) (*ChangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [valueChanges] contains a subset of the keys that were added or had their
// values modified between [startRootID] to [endRootID].

// create a copy so edits of the []byte don't affect the db

// Since we hold [db.commitlock] we must still have sufficient
// history to recreate the trie at [endRootID].

// strip out any common nodes to reduce proof size

// Note that one of the following must be true:
//  - [result.StartProof] is non-empty.
//  - [result.EndProof] is non-empty.
//  - [result.KeyValues] is non-empty.
//  - [result.DeletedKeys] is non-empty.
// If all of these were false, it would mean that no
// [start] and [end] were given, and no diff between
// the trie at [startRootID] and [endRootID] was found.
// Since [startRootID] != [endRootID], this is impossible.

// NewView returns a new view on top of this Trie where the passed changes
// have been applied.
//
// Changes made to the view will only be reflected in the original trie if
// Commit is called.
//
// Assumes [db.commitLock] and [db.lock] aren't held.
func (db *merkleDB) NewView(
	_ context.Context,
	changes ViewChanges,
) (View, error) {
	_ = "STUB: not implemented"
	// ensure the db doesn't change while creating the new view
	return *new(View), nil
}

// ensure access to childViews is protected

func (db *merkleDB) Has(k []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db *merkleDB) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *merkleDB) NewBatch() database.Batch {
	_ = "STUB: not implemented"
	return *new(database.Batch)
}

func (db *merkleDB) NewIterator() database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *merkleDB) NewIteratorWithStart(start []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *merkleDB) NewIteratorWithPrefix(prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *merkleDB) NewIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *merkleDB) Put(k, v []byte) error { _ = "STUB: not implemented"; return nil }

// Same as [Put] but takes in a context used for tracing.
func (db *merkleDB) PutContext(ctx context.Context, k, v []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *merkleDB) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (db *merkleDB) DeleteContext(ctx context.Context, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Assumes values inside [ops] are safe to reference after the function
// returns. Assumes [db.lock] isn't held.
func (db *merkleDB) commitBatch(ops []database.BatchOp) error {
	_ = "STUB: not implemented"
	return nil
}

// commitView commits the changes in [trieToCommit] to [db].
// Assumes [trieToCommit]'s node IDs have been calculated.
// Assumes [db.commitLock] is held.
func (db *merkleDB) commitView(ctx context.Context, trieToCommit *view) error {
	_ = "STUB: not implemented"
	return nil
}

// invalidate all child views except for the view being committed

// move any child views of the committed trie onto the db

// Update root in database.

// moveChildViewsToDB removes any child views from the trieToCommit and moves
// them to the db.
//
// assumes [db.lock] is held
func (db *merkleDB) moveChildViewsToDB(trieToCommit *view) { _ = "STUB: not implemented"; return }

// applyChanges takes the [changes] and applies them to [db.intermediateNodeDB]
// and [valueNodeBatch].
//
// assumes [db.lock] is held
func (db *merkleDB) applyChanges(ctx context.Context, valueNodeBatch database.KeyValueWriterDeleter, changes *changeSummary) error {
	_ = "STUB: not implemented"
	return nil
}

// commitValueChanges is a thin wrapper around [valueNodeBatch.Write()] to
// provide tracing.
func (db *merkleDB) commitValueChanges(ctx context.Context, valueNodeBatch database.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// CommitToDB is a no-op for db since it is already in sync with itself.
// This exists to satisfy the View interface.
func (*merkleDB) CommitToDB(context.Context) error {
	_ = "STUB: not implemented"

	// This is defined on merkleDB instead of ChangeProof
	// because it accesses database internals.
	// Assumes [db.lock] isn't held.
	return nil
}

func (db *merkleDB) VerifyChangeProof(
	ctx context.Context,
	proof *ChangeProof,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	expectedEndRootID ids.ID,
	maxLength int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Update [endProofKey] with the largest key in [keyValues].

// Validate proof.

// Prevent commit writes to DB, but not prevent DB reads

// Ensure that the [startProof] has correct values.

// Ensure that the [endProof] has correct values.

// Prepare ops for the creation of the view.

// Don't need to lock [view] because nobody else has a reference to it.

// For all the nodes along the edges of the proofs, insert the children whose
// keys are less than [insertChildrenLessThan] or whose keys are greater
// than [insertChildrenGreaterThan] into the trie so that we get the
// expected root ID (if this proof is valid).

// Make sure we get the expected root.

func (db *merkleDB) VerifyRangeProof(
	ctx context.Context,
	proof *RangeProof,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
	expectedEndRootID ids.ID,
	maxLength int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Invalidates and removes any child views that aren't [exception].
// Assumes [db.lock] is held.
func (db *merkleDB) invalidateChildrenExcept(exception *view) { _ = "STUB: not implemented"; return }

// If the root is on disk, set [db.root] to it.
// Otherwise leave [db.root] as Nothing.
func (db *merkleDB) initializeRoot() error { _ = "STUB: not implemented"; return nil }

// Root isn't on disk.

// Root is on disk.

// First, see if root is an intermediate node.
/* hasValue */

// The root must be a value node.
/* hasValue */

// Returns a view of the trie as it was when it had root [rootID] for keys within range [start, end].
// If [start] is Nothing, there's no lower bound on the range.
// If [end] is Nothing, there's no upper bound on the range.
// Assumes [db.commitLock] is read locked.
func (db *merkleDB) getTrieAtRootForRange(
	rootID ids.ID,
	start maybe.Maybe[[]byte],
	end maybe.Maybe[[]byte],
) (Trie, error) {
	_ = "STUB: not implemented"
	// looking for the trie's current root id, so return the trie unmodified
	return *new(Trie), nil
}

// Returns all keys in range [start, end] that aren't in [keySet].
// If [start] is Nothing, then the range has no lower bound.
// If [end] is Nothing, then the range has no upper bound.
func (db *merkleDB) getKeysNotInSet(start, end maybe.Maybe[[]byte], keySet set.Set[string]) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a copy of the node with the given [key].
// hasValue determines which db the key is looked up in (intermediateNodeDB or valueNodeDB)
// This copy may be edited by the caller without affecting the database state.
// Returns database.ErrNotFound if the node doesn't exist.
// Assumes [db.lock] isn't held.
func (db *merkleDB) getEditableNode(key Key, hasValue bool) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns the node with the given [key].
// hasValue determines which db the key is looked up in (intermediateNodeDB or valueNodeDB)
// Editing the returned node affects the database state.
// Returns database.ErrNotFound if the node doesn't exist.
// Assumes [db.lock] is read locked.
func (db *merkleDB) getNode(key Key, hasValue bool) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Assumes [db.lock] or [db.commitLock] is read locked.
func (db *merkleDB) getRoot() maybe.Maybe[*node] { _ = "STUB: not implemented"; return nil }

func (db *merkleDB) Clear() error { _ = "STUB: not implemented"; return nil }

// Clear nodes from disk and caches

// Clear root

// Clear history

func (db *merkleDB) getTokenSize() int { _ = "STUB: not implemented"; return 0 }

// Returns [key] prefixed by [prefix].
// The returned *[]byte is taken from [bufferPool] and should be returned to it
// when the caller is done with it.
func addPrefixToKey(bufferPool *utils.BytesPool, prefix []byte, key []byte) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// cacheEntrySize returns a rough approximation of the memory consumed by storing the key and node.
func cacheEntrySize(key Key, n *node) int { _ = "STUB: not implemented"; return 0 }

// findNextKey returns the start of the key range that should be fetched next
// given that we just received a range/change proof that proved a range of
// key-value pairs ending at [lastReceivedKey].
//
// [rangeEnd] is the end of the range that we want to fetch.
//
// Returns Nothing if there are no more keys to fetch in [lastReceivedKey, rangeEnd].
//
// [endProof] is the end proof of the last proof received.
//
// Invariant: [lastReceivedKey] < [rangeEnd].
// If [rangeEnd] is Nothing it's considered > [lastReceivedKey].
func (db *merkleDB) findNextKey(
	largestHandledKey maybe.Maybe[[]byte],
	rangeEnd maybe.Maybe[[]byte],
	endProof []ProofNode,
) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The largest handled key isn't equal to the end of the range.
// Find the start of the next key range to fetch.
// Note that [largestHandledKey] can't be Nothing.
// Proof: Suppose it is. That means that we got a range/change proof that proved up to the
// greatest key-value pair in the database. That means we requested a proof with no upper
// bound. That is, [rangeEnd] is Nothing. Since we're here, [bothNothing] is false,
// which means [rangeEnd] isn't Nothing. Contradiction.

// We try to find the next key to fetch by looking at the end proof.
// If the end proof is empty, we have no information to use.
// Start fetching from the next key after [lastReceivedKey].

// We want the first key larger than the [lastReceivedKey].
// This is done by taking two proofs for the same key
// (one that was just received as part of a proof, and one from the local db)
// and traversing them from the longest key to the shortest key.
// For each node in these proofs, compare if the children of that node exist
// or have the same ID in the other proof.

// If the received proof is an exclusion proof, the last node may be for a
// key that is after the [lastReceivedKey].
// If the last received node's key is after the [lastReceivedKey], it can
// be removed to obtain a valid proof for a prefix of the [lastReceivedKey].

// update the proofKeyPath to be for the prefix

// get a proof for the same key as the received proof from the local db

// The local proof may also be an exclusion proof with an extra node.
// Remove this extra node if it exists to get a proof of the same key as the received proof

// Add sentinel node back into the localProofNodes, if it is missing.
// Required to ensure that a common node exists in both proofs

// Add sentinel node back into the endProof, if it is missing.
// Required to ensure that a common node exists in both proofs

// traverse the two proofs from the deepest nodes up to the sentinel node until a difference is found

// [deepestNode] is the proof node with the longest key (deepest in the trie) in the
// two proofs that hasn't been handled yet.
// [deepestNodeFromOtherProof] is the proof node from the other proof with
// the same key/depth if it exists, nil otherwise.

// select the deepest proof node from the two proofs

// there was a branch node in the received proof that isn't in the local proof
// see if the received proof node has children not present in the local proof

// we have dealt with this received node, so move on to the next received node

// there was a branch node in the local proof that isn't in the received proof
// see if the local proof node has children not present in the received proof

// we have dealt with this local node, so move on to the next local node

// the two nodes are at the same depth
// see if any of the children present in the local proof node are different
// from the children in the received proof node

// we have dealt with this local node and received node, so move on to the next nodes

// We only want to look at the children with keys greater than the proofKey.
// The proof key has the deepest node's key as a prefix,
// so only the next token of the proof key needs to be considered.

// If the deepest node has the same key as [proofKeyPath],
// then all of its children have keys greater than the proof key,
// so we can start at the 0 token.

// If the deepest node has a key shorter than the key being proven,
// we can look at the next token index of the proof key to determine which of that
// node's children have keys larger than [proofKeyPath].
// Any child with a token greater than the [proofKeyPath]'s token at that
// index will have a larger key.

// determine if there are any differences in the children for the deepest unhandled node of the two proofs

// If the nextKey is before or equal to the [lastReceivedKey]
// then we couldn't find a better answer than the [lastReceivedKey].
// Set the nextKey to [lastReceivedKey] + 0, which is the first key in
// the open range (lastReceivedKey, rangeEnd).

// If the [nextKey] is larger than the end of the range, return Nothing to signal that there is no next key in range

// the nextKey is within the open range (lastReceivedKey, rangeEnd), so return it

// findChildDifference returns the first child index that is different between node 1 and node 2 if one exists and
// a bool indicating if any difference was found
func findChildDifference(node1, node2 *ProofNode, startIndex int) (byte, bool) {
	_ = "STUB: not implemented"
	// Children indices >= [startIndex] present in at least one of the nodes.
	return 0, false
}

// if one node has a child and the other doesn't or the children ids don't match,
// return the current child index as the first difference

// there were no differences found
