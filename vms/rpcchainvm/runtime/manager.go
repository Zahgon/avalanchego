// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runtime

import (
	"context"
	"sync"
)

// Manages tracking and shutdown of VM runtimes.
type manager struct {
	lock     sync.Mutex
	runtimes []Stopper
}

// NewManager returns manager of VM runtimes.
//
// TODO: If a runtime exits before the call to `manager.Stop`, it would be nice
// to remove it from the current set.
func NewManager() Manager { _ = "STUB: not implemented"; return *new(Manager) }

func (m *manager) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *manager) TrackRuntime(runtime Stopper) { _ = "STUB: not implemented"; return }
