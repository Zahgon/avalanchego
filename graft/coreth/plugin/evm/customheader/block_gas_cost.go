// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
)

var (
	ErrInsufficientBlockGas                     = errors.New("insufficient gas to cover the block cost")
	errInvalidExtraStateChangeContribution      = errors.New("invalid extra state change contribution")
	errInvalidBaseFeeApricotPhase4              = errors.New("invalid base fee in apricot phase 4")
	errInvalidRequiredBlockGasCostApricotPhase4 = errors.New("invalid block gas cost in apricot phase 4")
)

// BlockGasCost calculates the required block gas cost based on the parent
// header and the timestamp of the new block.
// Prior to AP4, the returned block gas cost will be nil.
// In Granite, the returned block gas cost will be 0.
func BlockGasCost(
	config *extras.ChainConfig,
	parent *types.Header,
	timestamp uint64,
) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// Treat an invalid parent/current time combination as 0 elapsed time.
//
// TODO: Does it even make sense to handle this? The timestamp should be
// verified to ensure this never happens.

// BlockGasCostWithStep calculates the required block gas cost based on the
// parent cost and the time difference between the parent block and new block.
//
// This is a helper function that allows the caller to manually specify the step
// value to use.
func BlockGasCostWithStep(
	parentCost *big.Int,
	step uint64,
	timeElapsed uint64,
) uint64 {
	_ = "STUB: not implemented"
	// Handle AP3/AP4 boundary by returning the minimum value as the boundary.
	return 0
}

// [ap4.MaxBlockGasCost] is <= MaxUint64, so we know that parentCost is
// always going to be a valid uint64.

func VerifyBlockFee(
	baseFee *big.Int,
	requiredBlockGasCost *big.Int,
	txs []*types.Transaction,
	receipts []*types.Receipt,
	extraStateChangeContribution *big.Int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the required block gas cost is 0, we don't need to verify the block fee

// Add in the external contribution

// Calculate the total excess (denominated in AVAX) over the base fee that was paid towards the block fee

// Each transaction contributes the excess over the baseFee towards the totalBlockFee
// This should be equivalent to the sum of the "priority fees" within EIP-1559.

// Multiply the [txFeePremium] by the gasUsed in the transaction since this gives the total AVAX that was paid
// above the amount required if the transaction had simply paid the minimum base fee for the block.
//
// Ex. LegacyTx paying a gas price of 100 gwei for 1M gas in a block with a base fee of 10 gwei.
// Total Fee = 100 gwei * 1M gas
// Minimum Fee = 10 gwei * 1M gas (minimum fee that would have been accepted for this transaction)
// Fee Premium = 90 gwei
// Total Overpaid = 90 gwei * 1M gas

// Calculate how much gas the [totalBlockFee] would purchase at the price level
// set by the base fee of this block.

// Require that the amount of gas purchased by the effective tips within the
// block covers at least `requiredBlockGasCost`.
//
// NOTE: To determine the required block fee, multiply
// `requiredBlockGasCost` by `baseFee`.
