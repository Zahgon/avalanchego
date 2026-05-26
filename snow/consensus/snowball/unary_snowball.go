// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

var _ Unary = (*unarySnowball)(nil)

func newUnarySnowball(alphaPreference int, terminationConditions []terminationCondition) unarySnowball {
	_ = "STUB: not implemented"
	return *new(unarySnowball)
}

// unarySnowball is the implementation of a unary snowball instance
type unarySnowball struct {
	// wrap the unary snowflake logic
	unarySnowflake

	// preferenceStrength tracks the total number of polls with a preference
	preferenceStrength int
}

func (sb *unarySnowball) RecordPoll(count int) { _ = "STUB: not implemented"; return }

func (sb *unarySnowball) Extend(choice int) Binary { _ = "STUB: not implemented"; return *new(Binary) }

func (sb *unarySnowball) Clone() Unary { _ = "STUB: not implemented"; return *new(Unary) }

func (sb *unarySnowball) String() string { _ = "STUB: not implemented"; return "" }
