// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"

	txfee "github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
	validatorfee "github.com/ava-labs/avalanchego/vms/platformvm/validators/fee"
)

func NextBlockTime(
	config validatorfee.Config,
	state Chain,
	clk *mockable.Clock,
) (time.Time, bool, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), false, nil
}

// [timestamp] = max(now, parentTime)

// If the NextStakerChangeTime is after timestamp, then we shouldn't return
// that the time was capped.

// timeWasCapped means that [timestamp] was reduced to [nextStakerChangeTime]

// [timestamp] = min(max(now, parentTime), nextStakerChangeTime)

// GetNextStakerChangeTime returns the next time a staker will be either added
// to or removed from the validator set. If the next staker change time is
// further in the future than [nextTime], then [nextTime] is returned.
func GetNextStakerChangeTime(
	config validatorfee.Config,
	state Chain,
	nextTime time.Time,
) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// If the iterator is empty, skip it

func getNextL1ValidatorEvictionTime(
	config validatorfee.Config,
	state Chain,
	nextTime time.Time,
) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// If there are no L1 validators, return

// Calculate the remaining funds that the next validator to evict has.

// GetActiveL1ValidatorsIterator iterates in order of increasing
// EndAccumulatedFee, so the first L1 validator is the next L1 validator
// to evict.

// Calculate how many seconds the remaining funds can last for.

// PickFeeCalculator creates either a simple or a dynamic fee calculator,
// depending on the active upgrade.
//
// PickFeeCalculator does not modify [state].
func PickFeeCalculator(config *config.Internal, state Chain) txfee.Calculator {
	_ = "STUB: not implemented"
	return *new(txfee.Calculator)
}
