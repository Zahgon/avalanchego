// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vms

import (
	"context"
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var ErrNotFound = errors.New("not found")

// A Factory creates new instances of a VM
type Factory interface {
	New(logging.Logger) (interface{}, error)
}

// Manager tracks a collection of VM factories, their aliases, and their
// versions.
// It has the following functionality:
//
//  1. Register a VM factory. To register a VM is to associate its ID with a
//     VMFactory which, when New() is called upon it, creates a new instance of
//     that VM.
//  2. Get a VM factory. Given the ID of a VM that has been registered, return
//     the factory that the ID is associated with.
//  3. Manage the aliases of VMs
//  4. Manage the versions of VMs
type Manager struct {
	// Note: The string representation of a VM's ID is also considered to be an
	// alias of the VM. That is, [vmID].String() is an alias for [vmID].
	ids.Aliaser

	log logging.Logger

	lock sync.RWMutex

	// Key: A VM's ID
	// Value: A factory that creates new instances of that VM
	factories map[ids.ID]Factory

	// Key: A VM's ID
	// Value: version the VM returned
	versions map[ids.ID]string
}

// NewManager returns an instance of a VM manager
func NewManager(log logging.Logger, aliaser ids.Aliaser) *Manager {
	_ = "STUB: not implemented"
	return nil
}

// Return a factory that can create new instances of the vm whose ID is vmID.
func (m *Manager) GetFactory(vmID ids.ID) (Factory, error) {
	_ = "STUB: not implemented"
	return *new(Factory), nil
}

// Map vmID to factory. factory creates new instances of the vm whose
// ID is vmID.
func (m *Manager) RegisterFactory(ctx context.Context, vmID ids.ID, factory Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// Drop the shutdown error to surface the original error

// ListFactories returns all the IDs that have had factories registered.
func (m *Manager) ListFactories() ([]ids.ID, error) { _ = "STUB: not implemented"; return nil, nil }

// Versions returns the primary alias of the VM mapped to the reported
// version of the VM for all the registered VMs that reported versions.
func (m *Manager) Versions() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
