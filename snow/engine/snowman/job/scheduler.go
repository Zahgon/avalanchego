// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package job provides a Scheduler to manage and execute Jobs with
// dependencies.
package job

import "context"

// Job is a unit of work that can be executed based on the result of resolving
// requested dependencies.
type Job[T any] interface {
	Execute(ctx context.Context, fulfilled []T, abandoned []T) error
}

type job[T comparable] struct {
	// Once all dependencies are resolved, the job will be executed.
	numUnresolved int
	fulfilled     []T
	abandoned     []T
	job           Job[T]
}

// Scheduler implements a dependency graph for jobs. Jobs can be registered with
// dependencies, and once all dependencies are resolved, the job will be
// executed.
type Scheduler[T comparable] struct {
	// dependents maps a dependency to the jobs that depend on it.
	dependents map[T][]*job[T]
}

func NewScheduler[T comparable]() *Scheduler[T] { _ = "STUB: not implemented"; return nil }

// Schedule a job to be executed once all of its dependencies are resolved. If a
// job is scheduled with no dependencies, it's executed immediately.
//
// In order to prevent a memory leak, all dependencies must eventually either be
// fulfilled or abandoned.
//
// While registering a job with duplicate dependencies is discouraged, it is
// allowed.
func (s *Scheduler[T]) Schedule(ctx context.Context, userJob Job[T], dependencies ...T) error {
	_ = "STUB: not implemented"
	return nil
}

// NumDependencies returns the number of dependencies that jobs are currently
// blocking on.
func (s *Scheduler[_]) NumDependencies() int { _ = "STUB: not implemented"; return 0 }

// Fulfill a dependency. If all dependencies for a job are resolved, the job
// will be executed.
//
// It is safe to call the scheduler during the execution of a job.
func (s *Scheduler[T]) Fulfill(ctx context.Context, dependency T) error {
	_ = "STUB: not implemented"
	return nil
}

// Abandon a dependency. If all dependencies for a job are resolved, the job
// will be executed.
//
// It is safe to call the scheduler during the execution of a job.
func (s *Scheduler[T]) Abandon(ctx context.Context, dependency T) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scheduler[T]) resolveDependency(
	ctx context.Context,
	dependency T,
	fulfilled bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
