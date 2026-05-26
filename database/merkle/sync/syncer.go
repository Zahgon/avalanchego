// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/maybe"

	pb "github.com/ava-labs/avalanchego/proto/pb/sync"
)

const (
	DefaultRequestKeyLimit      = MaxKeyValuesLimit
	DefaultRequestByteSizeLimit = maxByteSizeLimit
	initialRetryWait            = 10 * time.Millisecond
	maxRetryWait                = time.Second
	retryWaitFactor             = 1.5 // Larger --> timeout grows more quickly
	logInterval                 = time.Minute
)

var (
	ErrAlreadyStarted                 = errors.New("cannot start a Syncer that has already been started")
	ErrAlreadyClosed                  = errors.New("Syncer is closed")
	ErrNoRangeProofMarshalerProvided  = errors.New("range proof marshaler is a required field of the sync config")
	ErrNoChangeProofMarshalerProvided = errors.New("change proof marshaler is a required field of the sync config")
	ErrNoProofClientProvided          = errors.New("proof client is a required field of the sync config")
	ErrNoDatabaseProvided             = errors.New("sync database is a required field of the sync config")
	ErrNoLogProvided                  = errors.New("log is a required field of the sync config")
	ErrZeroWorkLimit                  = errors.New("simultaneous work limit must be greater than 0")
	ErrFinishedWithUnexpectedRoot     = errors.New("finished syncing with an unexpected root")
	errInvalidRangeProof              = errors.New("failed to verify range proof")
	errInvalidChangeProof             = errors.New("failed to verify change proof")
	errTooManyBytes                   = errors.New("response contains more than requested bytes")
	errUnexpectedResponseType         = errors.New("unexpected response type")
)

type priority byte

// Note that [highPriority] > [medPriority] > [lowPriority].
const (
	lowPriority priority = iota + 1
	medPriority
	highPriority
	retryPriority
)

// Signifies that we should sync the range [start, end].
// nil [start] means there is no lower bound.
// nil [end] means there is no upper bound.
// [localRootID] is the ID of the root of this range in our database.
// If we have no local root for this range, [localRootID] is ids.Empty.
type workItem struct {
	start       maybe.Maybe[[]byte]
	end         maybe.Maybe[[]byte]
	priority    priority
	localRootID ids.ID
	attempt     int
	queueTime   time.Time
}

func (w *workItem) requestFailed() { _ = "STUB: not implemented"; return }

// Overflow check

func newWorkItem(localRootID ids.ID, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], priority priority, queueTime time.Time) *workItem {
	_ = "STUB: not implemented"
	return nil
}

type Syncer[R any, C any] struct {
	// The database to sync.
	db DB[R, C]

	// Must be held when accessing [config.TargetRoot].
	syncTargetLock sync.RWMutex
	config         Config[R, C]

	workLock sync.Mutex
	// The number of work items currently being processed.
	// Namely, the number of goroutines executing [doWork].
	// [workLock] must be held when accessing [processingWorkItems].
	processingWorkItems int
	// [workLock] must be held while accessing [unprocessedWork].
	unprocessedWork *workHeap
	// Signalled when:
	// - An item is added to [unprocessedWork].
	// - An item is added to [processedWork].
	// - Close() is called.
	// [workLock] is its inner lock.
	unprocessedWorkCond *lock.Cond
	// [workLock] must be held while accessing [processedWork].
	processedWork *workHeap

	// When this is closed:
	// - [closed] is true.
	// - [cancelCtx] was called.
	// - [workToBeDone] and [completedWork] are closed.
	doneChan chan struct{}

	errLock sync.Mutex
	// If non-nil, there was a fatal error.
	// [errLock] must be held when accessing [fatalError].
	fatalError error

	// Cancels all currently processing work items.
	cancelCtx context.CancelFunc

	// Set to true when StartSyncing is called.
	syncing   bool
	closeOnce sync.Once

	stateSyncNodeIdx uint32
	metrics          SyncMetrics
}

// TODO remove non-config values out of this struct
type Config[R any, C any] struct {
	RangeProofMarshaler   Marshaler[R]
	ChangeProofMarshaler  Marshaler[C]
	ProofClient           *p2p.Client
	SimultaneousWorkLimit int
	Log                   logging.Logger
	TargetRoot            ids.ID
	EmptyRoot             ids.ID
	StateSyncNodes        []ids.NodeID
}

func NewSyncer[R any, C any](
	db DB[R, C],
	config Config[R, C],
	registerer prometheus.Registerer,
) (*Syncer[R, C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sync initiates the trie syncing process and blocks until one of the following occurs:
//   - [Syncer.Sync] is complete.
//   - [Syncer.Sync] fatally errored.
//   - `ctx` is canceled.
//
// If `ctx` is canceled, returns [context.Context.Err].
func (s *Syncer[_, _]) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Blocks until syncing completes, errors, or the context is canceled.

// There was a fatal error.

// This should never happen.

// setup initiates the work queue and enables cancellation through a new context.
func (s *Syncer[_, _]) setup(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Add work item to fetch the entire key range.
// Note that this will be the first work item to be processed.

// workLoop awaits signal on [s.unprocessedWorkCond], which indicates that there
// is work to do or syncing completes.  If there is work, workLoop will dispatch a goroutine to do
// the work.
// Assumes [s.workLock] is not held.
func (s *Syncer[_, _]) workLoop(ctx context.Context) {
	_ = "STUB: not implemented"

	// Invariant: [s.workLock] is held when this goroutine begins.
	return
}

// Keep doing work until we're closed, done or [ctx] is canceled.

// Invariant: [s.workLock] is held here.

// [s.workLock] released by defer.

// We're already processing the maximum number of work items.
// Wait until one of them finishes or the ctx is canceled.

// There's no work to do, and there are no work items being processed
// which could cause work to be added, so we're done.
// [s.workLock] released by defer.

// No work to do, but in-flight work may yet produce more.
// Wait returns when work is added or ctx is canceled.

func (s *Syncer[_, _]) logProgress(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Syncer[_, _]) getProgress(root ids.ID) float64 { _ = "STUB: not implemented"; return 0 }

// close is called when there is a fatal error or sync is complete.
// [workLock] must be held
func (s *Syncer[_, _]) close() { _ = "STUB: not implemented"; return }

// ensure any goroutines waiting for work from the heaps gets released

// signal all code waiting on the sync to complete

func (s *Syncer[_, _]) finishWorkItem() { _ = "STUB: not implemented"; return }

// Processes [item] by fetching a change or range proof.
func (s *Syncer[_, _]) doWork(ctx context.Context, work *workItem) {
	_ = "STUB: not implemented"
	// Backoff for failed requests accounting for time this job has already
	// spent waiting in the unprocessed queue
	return
}

// Check if we can start this work item before the context deadline

// the keys in this range have not been downloaded, so get all key/values

// the keys in this range have already been downloaded, but the root changed, so get all changes

// Fetch and apply the change proof given by [work].
// Assumes [s.workLock] is not held.
func (s *Syncer[_, _]) requestChangeProof(ctx context.Context, work *workItem) {
	_ = "STUB: not implemented"
	return
}

// Start root is the same as the end root, so we're done.

// The trie is empty after this change.
// Delete all the key-value pairs in the range.

// TODO log responses

// Fetch and apply the range proof given by [work].
// Assumes [s.workLock] is not held.
func (s *Syncer[_, _]) requestRangeProof(ctx context.Context, work *workItem) {
	_ = "STUB: not implemented"
	return
}

// TODO log responses

func (s *Syncer[_, _]) sendRequest(
	ctx context.Context,
	client *p2p.Client,
	requestBytes []byte,
	onResponse p2p.AppResponseCallback,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the next nodeID to query using the [nodeIdx] offset.
// If we're out of nodes, loop back to 0.
// We do this try to query a different node each time if possible.

func (s *Syncer[_, _]) retryWork(work *workItem) { _ = "STUB: not implemented"; return }

// Returns an error if we should drop the response
func (s *Syncer[_, _]) shouldHandleResponse(
	bytesLimit uint32,
	responseBytes []byte,
	err error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO can we remove this?

// If we're closed, don't apply the proof.

func (s *Syncer[R, _]) handleRangeProofResponse(
	ctx context.Context,
	targetRootID ids.ID,
	work *workItem,
	request *pb.RangeProofRequest,
	responseBytes []byte,
	err error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// A change proof returned is unexpected.

// Replace all the key-value pairs in the DB from start to end with values from the response.

func (s *Syncer[R, C]) handleChangeProofResponse(
	ctx context.Context,
	targetRootID ids.ID,
	work *workItem,
	request *pb.ChangeProofRequest,
	responseBytes []byte,
	err error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// The server had enough history to send us a change proof

// if the proof wasn't empty, apply changes to the sync DB

// The server did not have enough history to send us a change proof
// so they sent a range proof instead.

// Add all the key-value pairs we got to the database.

func (s *Syncer[_, _]) error() error { _ = "STUB: not implemented"; return nil }

func (s *Syncer[_, _]) UpdateSyncTarget(syncTargetRoot ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// the target hasn't changed, so there is nothing to do

// move all completed ranges into the work heap with high priority

// Note that [s.processedWork].Close() hasn't
// been called because we have [s.workLock]
// and we checked that [s.closed] is false.

// Only signal once because we only have 1 goroutine
// waiting on [s.unprocessedWorkCond].

func (s *Syncer[_, _]) getTargetRoot() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// Record that there was a fatal error and begin shutting down.
func (s *Syncer[_, _]) setError(err error) { _ = "STUB: not implemented"; return }

// Call in goroutine because we might be holding [s.workLock]

// Mark that we've fetched all the key-value pairs in the range
// [workItem.start, largestHandledKey] for the trie with root [rootID].
//
// If [workItem.start] is Nothing, then we've fetched all the key-value
// pairs up to and including [largestHandledKey].
//
// If [largestHandledKey] is Nothing, then we've fetched all the key-value
// pairs at and after [workItem.start].
//
// [proofOfLargestKey] is the end proof for the range/change proof
// that gave us the range up to and including [largestHandledKey].
//
// Assumes [s.workLock] is not held.
func (s *Syncer[_, _]) completeWorkItem(
	work *workItem,
	largestHandledKey maybe.Maybe[[]byte],
	rootID ids.ID,
) {
	_ = "STUB: not implemented"
	// largestHandledKey being Nothing indicates that the entire range has been completed
	return
}

// the full range wasn't completed, so enqueue a new work item for the range [nextStartKey, workItem.end]

// Process [work] while holding [syncTargetLock] to ensure that object
// is added to the right queue, even if a target update is triggered

// the root has changed, so reinsert with high priority

// completed the range [work.start, lastKey], log and record in the completed work heap

// Queue the given key range to be fetched and applied.
// If there are sufficiently few unprocessed/processing work items,
// splits the range into two items and queues them both.
// Assumes [s.workLock] is not held.
func (s *Syncer[_, _]) enqueueWork(work *workItem) { _ = "STUB: not implemented"; return }

// There are too many work items already, don't split the range

// Split the remaining range into to 2.
// Find the middle point.

// The range is too small to split.
// If we didn't have this check we would add work items
// [start, start] and [start, end]. Since start <= end, this would
// violate the invariant of [s.unprocessedWork] and [s.processedWork]
// that there are no overlapping ranges.

// first item gets higher priority than the second to encourage finished ranges to grow
// rather than start a new range that is not contiguous with existing completed ranges

// find the midpoint between two keys
// start is expected to be less than end
// Nothing/nil [start] is treated as all 0's
// Nothing/nil [end] is treated as all 255's
func midPoint(startMaybe, endMaybe maybe.Maybe[[]byte]) maybe.Maybe[[]byte] {
	_ = "STUB: not implemented"
	return nil
}

// This check deals with cases where the end has a 255(or is nothing which is treated as all 255s) and the start key ends 255.
// For example, midPoint([255], nothing) should be [255, 127], not [255].
// The result needs the extra byte added on to the end to deal with the fact that the naive midpoint between 255 and 255 would be 255

// if total is odd, when we divide, we will lose the .5,
// record that in the leftover for the next digits

// find the midpoint between the start and the end

// larger than byte can hold, so carry over to previous byte

func calculateBackoff(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
