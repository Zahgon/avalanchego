// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"github.com/ava-labs/avalanchego/ids"
)

var _ Nnary = (*nnarySnowball)(nil)

func newNnarySnowball(alphaPreference int, terminationConditions []terminationCondition, choice ids.ID) nnarySnowball {
	_ = "STUB: not implemented"
	return *new(nnarySnowball)
}

// nnarySnowball is a naive implementation of a multi-color snowball instance
type nnarySnowball struct {
	// wrap the n-nary snowflake logic
	nnarySnowflake

	// preference is the choice with the largest number of polls which preferred
	// it. Ties are broken by switching choice lazily
	preference ids.ID

	// maxPreferenceStrength is the maximum value stored in [preferenceStrength]
	maxPreferenceStrength int

	// preferenceStrength tracks the total number of network polls which
	// preferred that choice
	preferenceStrength map[ids.ID]int
}

func (sb *nnarySnowball) Preference() ids.ID {
	_ = "STUB: not implemented"
	// It is possible, with low probability, that the snowflake preference is
	// not equal to the snowball preference when snowflake finalizes. However,
	// this case is handled for completion. Therefore, if snowflake is
	// finalized, then our finalized snowflake choice should be preferred.
	return *new(ids.ID)
}

func (sb *nnarySnowball) RecordPoll(count int, choice ids.ID) { _ = "STUB: not implemented"; return }

func (sb *nnarySnowball) String() string { _ = "STUB: not implemented"; return "" }
