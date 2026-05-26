// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package profiler

import (
	"errors"
	"os"
)

const (
	// Name of file that CPU profile is written to when StartCPUProfiler called
	cpuProfileFile = "cpu.profile"
	// Name of file that memory profile is written to when MemoryProfile called
	memProfileFile = "mem.profile"
	// Name of file that lock profile is written to
	lockProfileFile = "lock.profile"
)

var (
	_ Profiler = (*profiler)(nil)

	errCPUProfilerRunning    = errors.New("cpu profiler already running")
	errCPUProfilerNotRunning = errors.New("cpu profiler doesn't exist")
)

// Profiler provides helper methods for measuring the current performance of
// this process
type Profiler interface {
	// StartCPUProfiler starts measuring the cpu utilization of this process
	StartCPUProfiler() error

	// StopCPUProfiler stops measuring the cpu utilization of this process
	StopCPUProfiler() error

	// MemoryProfile dumps the current memory utilization of this process
	MemoryProfile() error

	// LockProfile dumps the current lock statistics of this process
	LockProfile() error
}

type profiler struct {
	dir,
	cpuProfileName,
	memProfileName,
	lockProfileName string

	cpuProfileFile *os.File
}

func New(dir string) Profiler { _ = "STUB: not implemented"; return *new(Profiler) }

func newProfiler(dir string) *profiler { _ = "STUB: not implemented"; return nil }

func (p *profiler) StartCPUProfiler() error { _ = "STUB: not implemented"; return nil }

// Return the original error

func (p *profiler) StopCPUProfiler() error { _ = "STUB: not implemented"; return nil }

func (p *profiler) MemoryProfile() error { _ = "STUB: not implemented"; return nil }

// get up-to-date statistics

// Return the original error

func (p *profiler) LockProfile() error { _ = "STUB: not implemented"; return nil }

// Return the original error
