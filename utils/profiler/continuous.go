// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package profiler

import (
	"time"
)

// Config that is used to describe the options of the continuous profiler.
type Config struct {
	Dir         string        `json:"dir"`
	Enabled     bool          `json:"enabled"`
	Freq        time.Duration `json:"freq"`
	MaxNumFiles int           `json:"maxNumFiles"`
}

// ContinuousProfiler periodically captures CPU, memory, and lock profiles
type ContinuousProfiler interface {
	Dispatch() error
	Shutdown()
}

type continuousProfiler struct {
	profiler    *profiler
	freq        time.Duration
	maxNumFiles int

	// Dispatch returns when closer is closed
	closer chan struct{}
}

func NewContinuous(dir string, freq time.Duration, maxNumFiles int) ContinuousProfiler {
	_ = "STUB: not implemented"
	return *new(ContinuousProfiler)
}

func (p *continuousProfiler) Dispatch() error { _ = "STUB: not implemented"; return nil }

func (p *continuousProfiler) start() error { _ = "STUB: not implemented"; return nil }

func (p *continuousProfiler) stop() error { _ = "STUB: not implemented"; return nil }

func (p *continuousProfiler) rotate() error { _ = "STUB: not implemented"; return nil }

func (p *continuousProfiler) Shutdown() {
	_ = "STUB: not implemented"

	// Renames the file at [name] to [name].1, the file at [name].1 to [name].2, etc.
	// Assumes that there is a file at [name].
	return
}

func rotate(name string, maxNumFiles int) error { _ = "STUB: not implemented"; return nil }
