// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params/extras"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm/upgrade/subnetevm"
)

var (
	maxUint256Plus1 = new(big.Int).Lsh(common.Big1, 256)
	maxUint256      = new(big.Int).Sub(maxUint256Plus1, common.Big1)

	errInvalidTimestamp = errors.New("invalid timestamp")
)

// baseFeeFromWindow should only be called if `timestamp` >= `config.SubnetEVMTimestamp`
func baseFeeFromWindow(config *extras.ChainConfig, feeConfig commontype.FeeConfig, parent *types.Header, timestamp uint64) (*big.Int, error) {
	_ = "STUB: not implemented"
	// If the current block is the first EIP-1559 block, or it is the genesis block
	// return the initial slice and initial base fee.
	return nil, nil
}

// Calculate the amount of gas consumed within the rollup window.

// If the parent block used exactly its target gas, the baseFee stays
// the same.
//
// For legacy reasons, this is true even if the baseFee would have
// otherwise been clamped to a different range.

// If the parent block used more gas than its target, the baseFee should increase.

// Otherwise if the parent block used less gas than its target, the baseFee should decrease.

// This should never happen as the fee window calculations should
// have already failed, but it is kept for clarity.

// If timeElapsed is greater than [subnetevm.WindowLen], apply the state
// transition to the base fee to account for the interval during which
// no blocks were produced.
//
// We use timeElapsed/[subnetevm.WindowLen], so that the transition is applied
// for every [subnetevm.WindowLen] seconds that has elapsed between the parent
// and this block.

// Because baseFeeDelta could actually be [common.Big1], we must not
// modify the existing value of `baseFeeDelta` but instead allocate
// a new one.

// Ensure that the base fee does not increase/decrease outside of the bounds

// feeWindow takes the previous header and the timestamp of its child block and
// calculates the expected fee window.
//
// feeWindow should only be called if timestamp >= config.SubnetEVMTimestamp
func feeWindow(
	config *extras.ChainConfig,
	parent *types.Header,
	timestamp uint64,
) (subnetevm.Window, error) {
	_ = "STUB: not implemented"
	// If the current block is the first EIP-1559 block, or it is the genesis block
	// return the initial window.
	return *new(subnetevm.Window), nil
}

// Compute the new state of the gas rolling window.

// roll the window over by the timeElapsed to generate the new rollup
// window.

// selectBigWithinBounds returns [value] if it is within the bounds:
// lowerBound <= value <= upperBound or the bound at either end if [value]
// is outside of the defined boundaries.
func selectBigWithinBounds(lowerBound, value, upperBound *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}
