// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ethapi

import (
	"context"
	"math/big"

	"github.com/ava-labs/libevm/common/hexutil"
)

const (
	minGasTip      = 1 // 1 wei
	feeDenominator = 100
)

var (
	bigMinGasTip      = big.NewInt(minGasTip)
	bigFeeDenominator = big.NewInt(feeDenominator)
)

type PriceOptionConfig struct {
	SlowFeePercentage uint64
	FastFeePercentage uint64
	MaxTip            uint64
}

type Price struct {
	GasTip *hexutil.Big `json:"maxPriorityFeePerGas"`
	GasFee *hexutil.Big `json:"maxFeePerGas"`
}

type PriceOptions struct {
	Slow   *Price `json:"slow"`
	Normal *Price `json:"normal"`
	Fast   *Price `json:"fast"`
}

// TODO: This can be moved to AVAX/custom API

// SuggestPriceOptions returns suggestions for what to display to a user for
// current transaction fees.
func (s *EthereumAPI) SuggestPriceOptions(ctx context.Context) (*PriceOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the chain isn't running with dynamic fees, return nil.

// Double the baseFee estimate without modifying the original variable.

type feeSpeeds struct {
	slow   *big.Int
	normal *big.Int
	fast   *big.Int
}

// calculateFeeSpeeds returns the slow, normal, and fast price options for a
// given min, estimate, and max,
//
// slow   = max(slowFeePerc/100 * min(estimate, maxFee), minFee)
// normal = min(estimate, maxFee)
// fast   = fastFeePerc/100 * estimate
func calculateFeeSpeeds(
	minFee *big.Int,
	estimate *big.Int,
	maxFee *big.Int,
	slowFeePerc *big.Int,
	fastFeePerc *big.Int,
) feeSpeeds {
	_ = "STUB: not implemented"
	// Cap the fee to keep slow and normal options reasonable during fee spikes.
	return *new(feeSpeeds)
}
