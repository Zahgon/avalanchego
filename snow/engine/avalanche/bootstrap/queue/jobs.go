// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package queue

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
)

const progressUpdateFrequency = 30 * time.Second

// Jobs tracks a series of jobs that form a DAG of dependencies.
type Jobs struct {
	// db ensures that database updates are atomically updated.
	db *versiondb.Database
	// state writes the job queue to [db].
	state *state
	// Measures the ETA until bootstrapping finishes in nanoseconds.
	etaMetric prometheus.Gauge
}

// New attempts to create a new job queue from the provided database.
func New(
	db database.Database,
	metricsNamespace string,
	metricsRegisterer prometheus.Registerer,
) (*Jobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetParser tells this job queue how to parse jobs from the database.
func (j *Jobs) SetParser(parser Parser) error { _ = "STUB: not implemented"; return nil }

func (j *Jobs) Has(jobID ids.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Returns how many pending jobs are waiting in the queue.
		nil
}

func (j *Jobs) PendingJobs() uint64 { _ = "STUB: not implemented"; return 0 }

// Push adds a new job to the queue. Returns true if [job] was added to the queue and false
// if [job] was already in the queue.
func (j *Jobs) Push(ctx context.Context, job Job) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Store this job into the database.

// This job needs to block on a set of dependencies.

// This job doesn't have any dependencies, so it should be placed onto the
// executable stack.

func (j *Jobs) ExecuteAll(
	ctx context.Context,
	chainCtx *snow.ConsensusContext,
	halter common.Haltable,
	restarted bool,
	acceptors ...snow.Acceptor,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Disable and clear state caches to prevent us from attempting to execute
// a vertex that was previously parsed, but not saved to the VM. Some VMs
// may only persist containers when they are accepted. This is a stop-gap
// measure to ensure the job will be re-parsed before executing until the VM
// provides a more explicit interface for freeing parsed blocks.
// TODO remove DisableCaching when VM provides better interface for freeing
// blocks.

// Note that acceptor.Accept must be called before executing [job] to
// honor Acceptor.Accept's invariant.

// Periodically print progress

// Now that executing has finished, zero out the ETA.

func (j *Jobs) Clear() error { _ = "STUB: not implemented"; return nil }

// Commit the versionDB to the underlying database.
func (j *Jobs) Commit() error { _ = "STUB: not implemented"; return nil }

type JobsWithMissing struct {
	*Jobs

	// keep the missing ID set in memory to avoid unnecessary database reads and
	// writes.
	missingIDs                            set.Set[ids.ID]
	removeFromMissingIDs, addToMissingIDs set.Set[ids.ID]
}

func NewWithMissing(
	db database.Database,
	metricsNamespace string,
	metricsRegisterer prometheus.Registerer,
) (*JobsWithMissing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetParser tells this job queue how to parse jobs from the database.
func (jm *JobsWithMissing) SetParser(ctx context.Context, parser Parser) error {
	_ = "STUB: not implemented"
	return nil
}

func (jm *JobsWithMissing) Clear() error { _ = "STUB: not implemented"; return nil }

func (jm *JobsWithMissing) Has(jobID ids.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Push adds a new job to the queue. Returns true if [job] was added to the queue and false
// if [job] was already in the queue.
func (jm *JobsWithMissing) Push(ctx context.Context, job Job) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Store this job into the database.

// This job needs to block on a set of dependencies.

// This job doesn't have any dependencies, so it should be placed onto the
// executable stack.

// AddMissingID adds [jobID] to missingIDs
func (jm *JobsWithMissing) AddMissingID(jobIDs ...ids.ID) { _ = "STUB: not implemented"; return }

// RemoveMissingID removes [jobID] from missingIDs
func (jm *JobsWithMissing) RemoveMissingID(jobIDs ...ids.ID) { _ = "STUB: not implemented"; return }

func (jm *JobsWithMissing) MissingIDs() []ids.ID { _ = "STUB: not implemented"; return nil }

func (jm *JobsWithMissing) NumMissingIDs() int { _ = "STUB: not implemented"; return 0 }

// Commit the versionDB to the underlying database.
func (jm *JobsWithMissing) Commit() error { _ = "STUB: not implemented"; return nil }

// cleanRunnableStack iterates over the jobs on the runnable stack and resets any job
// that has missing dependencies to block on those dependencies.
// Note: the jobs queue ensures that no job with missing dependencies will be placed
// on the runnable stack in the first place.
// However, for specific VM implementations blocks may be committed via a two stage commit
// (ex. platformvm Proposal and Commit/Abort blocks). This can cause an issue where if the first stage
// is executed immediately before the node dies, it will be removed from the runnable stack
// without writing the state transition to the VM's database. When the node restarts, the
// VM will not have marked the first block (the proposal block as accepted), but it could
// have already been removed from the jobs queue. cleanRunnableStack handles this case.
func (jm *JobsWithMissing) cleanRunnableStack(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If the job has missing dependencies, remove it from the runnable stack

// Add the missing dependencies to the set that needs to be fetched.
