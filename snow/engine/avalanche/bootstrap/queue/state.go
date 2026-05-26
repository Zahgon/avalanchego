// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package queue

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/linkeddb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

const (
	dependentsCacheSize = 1024
	jobsCacheSize       = 2048
)

var (
	runnableJobIDsPrefix = []byte("runnable")
	jobsPrefix           = []byte("jobs")
	dependenciesPrefix   = []byte("dependencies")
	missingJobIDsPrefix  = []byte("missing job IDs")
	metadataPrefix       = []byte("metadata")
	numJobsKey           = []byte("numJobs")
)

type state struct {
	parser         Parser
	runnableJobIDs linkeddb.LinkedDB
	cachingEnabled bool
	jobsCache      cache.Cacher[ids.ID, Job]
	jobsDB         database.Database
	// Should be prefixed with the jobID that we are attempting to find the
	// dependencies of. This prefixdb.Database should then be wrapped in a
	// linkeddb.LinkedDB to read the dependencies.
	dependenciesDB database.Database
	// This is a cache that tracks LinkedDB iterators that have recently been
	// made.
	dependentsCache cache.Cacher[ids.ID, linkeddb.LinkedDB]
	missingJobIDs   linkeddb.LinkedDB
	// This tracks the summary values of this state. Currently, this only
	// contains the last known checkpoint of how many jobs are currently in the
	// queue to execute.
	metadataDB database.Database
	// This caches the number of jobs that are currently in the queue to
	// execute.
	numJobs uint64
}

func newState(
	db database.Database,
	metricsNamespace string,
	metricsRegisterer prometheus.Registerer,
) (*state, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNumJobs(d database.Database, jobs database.Iteratee) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If we don't have a checkpoint, we need to initialize it.

func (s *state) Clear() error { _ = "STUB: not implemented"; return nil }

// clear runnableJobIDs

// clear jobs

// clear dependencies

// clear missing jobs IDs

// clear number of pending jobs

// AddRunnableJob adds [jobID] to the runnable queue
func (s *state) AddRunnableJob(jobID ids.ID) error { _ = "STUB: not implemented"; return nil }

// HasRunnableJob returns true if there is a job that can be run on the queue
func (s *state) HasRunnableJob() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// RemoveRunnableJob fetches and deletes the next job from the runnable queue
func (s *state) RemoveRunnableJob(ctx context.Context) (Job, error) {
	_ = "STUB: not implemented"
	return *new(Job), nil
}

// Guard rail to make sure we don't underflow.

// PutJob adds the job to the queue
func (s *state) PutJob(job Job) error { _ = "STUB: not implemented"; return nil }

// HasJob returns true if the job [id] is in the queue
func (s *state) HasJob(id ids.ID) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GetJob returns the job [id]
func (s *state) GetJob(ctx context.Context, id ids.ID) (Job, error) {
	_ = "STUB: not implemented"
	return *new(Job), nil
}

// AddDependency adds [dependent] as blocking on [dependency] being completed
func (s *state) AddDependency(dependency, dependent ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveDependencies removes the set of IDs that are blocking on the completion
// of [dependency] from the database and returns them.
func (s *state) RemoveDependencies(dependency ids.ID) ([]ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *state) DisableCaching() { _ = "STUB: not implemented"; return }

func (s *state) AddMissingJobIDs(missingIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) RemoveMissingJobIDs(missingIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) MissingJobIDs() ([]ids.ID, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *state) getDependentsDB(dependency ids.ID) linkeddb.LinkedDB {
	_ = "STUB: not implemented"
	return *new(linkeddb.LinkedDB)
}
