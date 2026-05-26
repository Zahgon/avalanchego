// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

var _ Binary = (*binarySnowball)(nil)

func newBinarySnowball(alphaPreference int, terminationConditions []terminationCondition, choice int) binarySnowball {
	_ = "STUB: not implemented"
	return *new(binarySnowball)
}

// binarySnowball is the implementation of a binary snowball instance
type binarySnowball struct {
	// wrap the binary snowflake logic
	binarySnowflake

	// preference is the choice with the largest number of polls which preferred
	// the color. Ties are broken by switching choice lazily
	preference int

	// preferenceStrength tracks the total number of network polls which
	// preferred each choice
	preferenceStrength [2]int
}

func (sb *binarySnowball) Preference() int {
	_ = "STUB: not implemented"
	// It is possible, with low probability, that the snowflake preference is
	// not equal to the snowball preference when snowflake finalizes. However,
	// this case is handled for completion. Therefore, if snowflake is
	// finalized, then our finalized snowflake choice should be preferred.
	return 0
}

func (sb *binarySnowball) RecordPoll(count, choice int) { _ = "STUB: not implemented"; return }

func (sb *binarySnowball) String() string { _ = "STUB: not implemented"; return "" }
