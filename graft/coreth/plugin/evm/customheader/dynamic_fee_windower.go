// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap3"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap4"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap5"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/etna"
)

var (
	maxUint256Plus1 = new(big.Int).Lsh(common.Big1, 256)
	maxUint256      = new(big.Int).Sub(maxUint256Plus1, common.Big1)

	ap3MinBaseFee  = big.NewInt(ap3.MinBaseFee)
	ap4MinBaseFee  = big.NewInt(ap4.MinBaseFee)
	etnaMinBaseFee = big.NewInt(etna.MinBaseFee)

	ap3MaxBaseFee = big.NewInt(ap3.MaxBaseFee)
	ap4MaxBaseFee = big.NewInt(ap4.MaxBaseFee)

	ap3BaseFeeChangeDenominator = big.NewInt(ap3.BaseFeeChangeDenominator)
	ap5BaseFeeChangeDenominator = big.NewInt(ap5.BaseFeeChangeDenominator)

	errInvalidTimestamp = errors.New("invalid timestamp")
)

// baseFeeFromWindow should only be called if `timestamp` >= `config.ApricotPhase3Timestamp`
func baseFeeFromWindow(config *extras.ChainConfig, parent *types.Header, timestamp uint64) (*big.Int, error) {
	_ = "STUB: not implemented"
	// If the current block is the first EIP-1559 block, or it is the genesis block
	// return the initial slice and initial base fee.
	return nil, nil
}

// If AP5, use a less responsive BaseFeeChangeDenominator and a higher gas
// block limit

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

// If timeElapsed is greater than [ap3.WindowLen], apply the state
// transition to the base fee to account for the interval during which
// no blocks were produced.
//
// We use timeElapsed/[ap3.WindowLen], so that the transition is applied
// for every [ap3.WindowLen] seconds that has elapsed between the parent
// and this block.

// Because baseFeeDelta could actually be [common.Big1], we must not
// modify the existing value of `baseFeeDelta` but instead allocate
// a new one.

// Ensure that the base fee does not increase/decrease outside of the bounds

// feeWindow takes the previous header and the timestamp of its child block and
// calculates the expected fee window.
//
// feeWindow should only be called if timestamp >= config.ApricotPhase3Timestamp
func feeWindow(
	config *extras.ChainConfig,
	parent *types.Header,
	timestamp uint64,
) (ap3.Window, error) {
	_ = "STUB: not implemented"
	// If the current block is the first EIP-1559 block, or it is the genesis block
	// return the initial window.
	return *new(ap3.Window), nil
}

// Add in parent's consumed gas

// blockGasCost is not included in the fee window after AP5, so it is
// left as 0.

// At the start of a new network, the parent may not have a populated
// ExtDataGasUsed.

// The blockGasCost is paid by the effective tips in the block using
// the block's value of baseFee.
//
// Although the child block may be in AP5 here, the blockGasCost is
// still calculated using the AP4 step. This is different than the
// actual BlockGasCost calculation used for the child block. This
// behavior is kept to preserve the original behavior of this function.

// On the boundary of AP3 and AP4 or at the start of a new network, the
// parent may not have a populated ExtDataGasUsed.

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
