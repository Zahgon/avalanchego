// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
)

var _ Consensus = (*Flat)(nil)

func NewFlat(factory Factory, params Parameters, choice ids.ID) Consensus {
	_ = "STUB: not implemented"
	return *new(Consensus)
}

// Flat is a naive implementation of a multi-choice snow instance
type Flat struct {
	// wraps the n-nary snow logic
	Nnary

	// params contains all the configurations of a snow instance
	params Parameters
}

func (f *Flat) RecordPoll(votes bag.Bag[ids.ID]) bool { _ = "STUB: not implemented"; return false }
