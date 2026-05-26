// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package enginetest

import (
	"testing"

	"github.com/ava-labs/avalanchego/ids"
)

// BootstrapTracker is a test subnet
type BootstrapTracker struct {
	T *testing.T

	CantIsBootstrapped, CantBootstrapped, CantOnBootstrapCompleted bool

	IsBootstrappedF func() bool
	BootstrappedF   func(ids.ID)

	OnBootstrapCompletedF func() chan struct{}
}

// Default set the default callable value to [cant]
func (s *BootstrapTracker) Default(cant bool) { _ = "STUB: not implemented"; return }

// IsBootstrapped calls IsBootstrappedF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail. Defaults to returning false.
func (s *BootstrapTracker) IsBootstrapped() bool { _ = "STUB: not implemented"; return false }

// Bootstrapped calls BootstrappedF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *BootstrapTracker) Bootstrapped(chainID ids.ID) { _ = "STUB: not implemented"; return }

func (s *BootstrapTracker) AllBootstrapped() <-chan struct{} { _ = "STUB: not implemented"; return nil }
