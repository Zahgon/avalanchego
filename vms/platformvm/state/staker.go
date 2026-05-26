// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"time"

	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var _ btree.LessFunc[*Staker] = (*Staker).Less

// Staker contains all information required to represent a validator or
// delegator in the current and pending validator sets.
// Invariant: Staker's size is bounded to prevent OOM DoS attacks.
type Staker struct {
	TxID            ids.ID
	NodeID          ids.NodeID
	PublicKey       *bls.PublicKey
	SubnetID        ids.ID
	Weight          uint64
	StartTime       time.Time
	EndTime         time.Time
	PotentialReward uint64

	// NextTime is the next time this staker will be moved from a validator set.
	// If the staker is in the pending validator set, NextTime will equal
	// StartTime. If the staker is in the current validator set, NextTime will
	// equal EndTime.
	NextTime time.Time

	// Priority specifies how to break ties between stakers with the same
	// NextTime. This ensures that stakers created by the same transaction type
	// are grouped together. The ordering of these groups is documented in
	// [priorities.go] and depends on if the stakers are in the pending or
	// current validator set.
	Priority txs.Priority
}

// Equals returns true if this staker is equal to the provided staker.
// If s.Less(other) and other.Less(s) are both false, then it doesn't mean that s.Equals(other) is true.
func (s *Staker) Equals(other *Staker) bool { _ = "STUB: not implemented"; return false }

// A *Staker is considered to be less than another *Staker when:
//
//  1. If its NextTime is before the other's.
//  2. If the NextTimes are the same, the *Staker with the lesser priority is the
//     lesser one.
//  3. If the priorities are also the same, the one with the lesser txID is
//     lesser.
func (s *Staker) Less(than *Staker) bool { _ = "STUB: not implemented"; return false }

func NewCurrentStaker(
	txID ids.ID,
	staker txs.Staker,
	startTime time.Time,
	potentialReward uint64,
) (*Staker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPendingStaker(txID ids.ID, staker txs.ScheduledStaker) (*Staker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
