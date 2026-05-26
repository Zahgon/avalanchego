// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2p

import (
	"context"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ ValidatorSet      = (*Validators)(nil)
	_ ValidatorSubset   = (*Validators)(nil)
	_ NodeSampler       = (*Validators)(nil)
	_ ConnectionHandler = (*Validators)(nil)
)

type ValidatorSet interface {
	Len(ctx context.Context) int
	Has(ctx context.Context, nodeID ids.NodeID) bool // TODO return error
}

type ValidatorSubset interface {
	Top(ctx context.Context, percentage float64) []ids.NodeID // TODO return error
}

func NewValidators(
	log logging.Logger,
	subnetID ids.ID,
	validators validators.State,
	maxValidatorSetStaleness time.Duration,
) *Validators {
	_ = "STUB: not implemented"
	return nil
}

// Validators contains a set of nodes that are staking.
type Validators struct {
	log                      logging.Logger
	subnetID                 ids.ID
	validators               validators.State
	maxValidatorSetStaleness time.Duration

	lock                sync.RWMutex
	peers               set.Set[ids.NodeID]
	connectedValidators set.Set[ids.NodeID]
	validatorList       []validator
	validatorSet        set.Set[ids.NodeID]
	totalWeight         uint64
	lastUpdated         time.Time
}

type validator struct {
	nodeID ids.NodeID
	weight uint64
}

func (v validator) Compare(other validator) int { _ = "STUB: not implemented"; return 0 }

// Sort in decreasing order of stake

// getCurrentValidators must not be called with Validators.lock held to avoid a
// potential deadlock.
//
// getCurrentValidators calls [validators.State] which grabs the context lock.
// [Validators.Connected] and [Validators.Disconnected] are called with the
// context lock.
func (v *Validators) getCurrentValidators(ctx context.Context) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore inactive ACP-77 validators.

// refresh must not be called with Validators.lock held.
func (v *Validators) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

// Even though validatorList may be nil, truncating will not panic.

// Sample returns a random sample of connected validators
func (v *Validators) Sample(ctx context.Context, limit int) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}

// Top returns the top [percentage] of validators, regardless of if they are
// connected or not.
func (v *Validators) Top(ctx context.Context, percentage float64) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}

// bound percentage inside [0, 1]

// Has returns if nodeID is a connected validator
func (v *Validators) Has(ctx context.Context, nodeID ids.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

// Len returns the number of connected validators.
func (v *Validators) Len(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

func (v *Validators) Connected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

func (v *Validators) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }
