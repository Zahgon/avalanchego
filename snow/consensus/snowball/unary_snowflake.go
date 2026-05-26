// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

var _ Unary = (*unarySnowflake)(nil)

func newUnarySnowflake(alphaPreference int, terminationConditions []terminationCondition) unarySnowflake {
	_ = "STUB: not implemented"
	return *new(unarySnowflake)
}

// unarySnowflake is the implementation of a unary snowflake instance
// Invariant:
// len(terminationConditions) == len(confidence)
// terminationConditions[i].alphaConfidence < terminationConditions[i+1].alphaConfidence
// terminationConditions[i].beta >= terminationConditions[i+1].beta
// confidence[i] >= confidence[i+1] (except after finalizing due to early termination)
type unarySnowflake struct {
	// alphaPreference is the threshold required to update the preference
	alphaPreference int

	// terminationConditions gives the ascending ordered list of alphaConfidence values
	// required to increment the corresponding confidence counter.
	// The corresponding beta values give the threshold required to finalize this instance.
	terminationConditions []terminationCondition

	// confidence is the number of consecutive successful polls for a given
	// alphaConfidence threshold.
	// This instance finalizes when confidence[i] >= terminationConditions[i].beta for any i
	confidence []int

	// finalized prevents the state from changing after the required number of
	// consecutive polls has been reached
	finalized bool
}

func (sf *unarySnowflake) RecordPoll(count int) { _ = "STUB: not implemented"; return }

// If I did not reach this alpha threshold, I did not
// reach any more alpha thresholds.
// Clear the remaining confidence counters.

// I reached this alpha threshold, increment the confidence counter
// and check if I can finalize.

func (sf *unarySnowflake) RecordUnsuccessfulPoll() { _ = "STUB: not implemented"; return }

func (sf *unarySnowflake) Finalized() bool { _ = "STUB: not implemented"; return false }

func (sf *unarySnowflake) Extend(choice int) Binary { _ = "STUB: not implemented"; return *new(Binary) }

func (sf *unarySnowflake) Clone() Unary { _ = "STUB: not implemented"; return *new(Unary) }

func (sf *unarySnowflake) String() string { _ = "STUB: not implemented"; return "" }
