// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reward

import (
	"math/big"
	"time"
)

var _ Calculator = (*calculator)(nil)

type Calculator interface {
	Calculate(stakedDuration time.Duration, stakedAmount, currentSupply uint64) uint64
}

type calculator struct {
	maxSubMinConsumptionRate *big.Int
	minConsumptionRate       *big.Int
	mintingPeriod            *big.Int
	supplyCap                uint64
}

func NewCalculator(c Config) Calculator { _ = "STUB: not implemented"; return *new(Calculator) }

// Reward returns the amount of tokens to reward the staker with.
//
// RemainingSupply = SupplyCap - ExistingSupply
// PortionOfExistingSupply = StakedAmount / ExistingSupply
// PortionOfStakingDuration = StakingDuration / MaximumStakingDuration
// MintingRate = MinMintingRate + MaxSubMinMintingRate * PortionOfStakingDuration
// Reward = RemainingSupply * PortionOfExistingSupply * MintingRate * PortionOfStakingDuration
func (c *calculator) Calculate(stakedDuration time.Duration, stakedAmount, currentSupply uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// Split [totalAmount] into [totalAmount * shares percentage] and the remainder.
//
// Invariant: [shares] <= [PercentDenominator]
func Split(totalAmount uint64, shares uint32) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Delay rounding as long as possible for small numbers
