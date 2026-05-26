// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package commontype

import (
	"errors"
	"math/big"
)

var (
	ErrGasLimitTooLow         = errors.New("gasLimit cannot be less than or equal to 0")
	ErrGasLimitNil            = errors.New("gasLimit cannot be nil")
	ErrMinBlockGasCostTooHigh = errors.New("minBlockGasCost cannot be greater than maxBlockGasCost")

	errMinBaseFeeNil                             = errors.New("minBaseFee cannot be nil")
	errTargetGasNil                              = errors.New("targetGas cannot be nil")
	errBaseFeeChangeDenominatorNil               = errors.New("baseFeeChangeDenominator cannot be nil")
	errMinBlockGasCostNil                        = errors.New("minBlockGasCost cannot be nil")
	errMaxBlockGasCostNil                        = errors.New("maxBlockGasCost cannot be nil")
	errBlockGasCostStepNil                       = errors.New("blockGasCostStep cannot be nil")
	errTargetBlockRateTooLow                     = errors.New("targetBlockRate cannot be less than or equal to 0")
	errMinBaseFeeNegative                        = errors.New("minBaseFee cannot be less than 0")
	errTargetGasTooLow                           = errors.New("targetGas cannot be less than or equal to 0")
	errBaseFeeChangeDenominatorTooLow            = errors.New("baseFeeChangeDenominator cannot be less than or equal to 0")
	errMinBlockGasCostNegative                   = errors.New("minBlockGasCost cannot be less than 0")
	errBlockGasCostStepNegative                  = errors.New("blockGasCostStep cannot be less than 0")
	errMaxBlockGasCostNotUint64                  = errors.New("maxBlockGasCost is not a valid uint64")
	errGasLimitExceedsHashLength                 = errors.New("gasLimit exceeds hash length")
	errTargetBlockRateExceedsHashLength          = errors.New("targetBlockRate exceeds hash length")
	errMinBaseFeeExceedsHashLength               = errors.New("minBaseFee exceeds hash length")
	errTargetGasExceedsHashLength                = errors.New("targetGas exceeds hash length")
	errBaseFeeChangeDenominatorExceedsHashLength = errors.New("baseFeeChangeDenominator exceeds hash length")
	errMinBlockGasCostExceedsHashLength          = errors.New("minBlockGasCost exceeds hash length")
	errMaxBlockGasCostExceedsHashLength          = errors.New("maxBlockGasCost exceeds hash length")
	errBlockGasCostStepExceedsHashLength         = errors.New("blockGasCostStep exceeds hash length")
)

// FeeConfig specifies the parameters for the dynamic fee algorithm, which determines the gas limit, base fee, and block gas cost of blocks
// on the network.
//
// The dynamic fee algorithm simply increases fees when the network is operating at a utilization level above the target and decreases fees
// when the network is operating at a utilization level below the target.
// This struct is used by Genesis and Fee Manager precompile.
// Any modification of this struct has direct affect on the precompiled contract
// and changes should be carefully handled in the precompiled contract code.
type FeeConfig struct {
	// GasLimit sets the max amount of gas consumed per block.
	GasLimit *big.Int `json:"gasLimit,omitempty"`

	// TargetBlockRate sets the target rate of block production in seconds.
	// A target of 2 will target producing a block every 2 seconds.
	TargetBlockRate uint64 `json:"targetBlockRate,omitempty"`

	// The minimum base fee sets a lower bound on the EIP-1559 base fee of a block.
	// Since the block's base fee sets the minimum gas price for any transaction included in that block, this effectively sets a minimum
	// gas price for any transaction.
	MinBaseFee *big.Int `json:"minBaseFee,omitempty"`

	// When the dynamic fee algorithm observes that network activity is above/below the [TargetGas], it increases/decreases the base fee proportionally to
	// how far above/below the target actual network activity is.

	// TargetGas specifies the targeted amount of gas (including block gas cost) to consume within a rolling 10s window.
	TargetGas *big.Int `json:"targetGas,omitempty"`
	// The BaseFeeChangeDenominator divides the difference between actual and target utilization to determine how much to increase/decrease the base fee.
	// This means that a larger denominator indicates a slower changing, stickier base fee, while a lower denominator will allow the base fee to adjust
	// more quickly.
	BaseFeeChangeDenominator *big.Int `json:"baseFeeChangeDenominator,omitempty"`

	// MinBlockGasCost sets the minimum amount of gas to charge for the production of a block.
	MinBlockGasCost *big.Int `json:"minBlockGasCost,omitempty"`
	// MaxBlockGasCost sets the maximum amount of gas to charge for the production of a block.
	MaxBlockGasCost *big.Int `json:"maxBlockGasCost,omitempty"`
	// BlockGasCostStep determines how much to increase/decrease the block gas cost depending on the amount of time elapsed since the previous block.
	// If the block is produced at the target rate, the block gas cost will stay the same as the block gas cost for the parent block.
	// If it is produced faster/slower, the block gas cost will be increased/decreased by the step value for each second faster/slower than the target
	// block rate accordingly.
	// Note: if the BlockGasCostStep is set to a very large number, it effectively requires block production to go no faster than the TargetBlockRate.
	//
	// Ex: if a block is produced two seconds faster than the target block rate, the block gas cost will increase by 2 * BlockGasCostStep.
	BlockGasCostStep *big.Int `json:"blockGasCostStep,omitempty"`
}

// represents an empty fee config without any field
var EmptyFeeConfig = FeeConfig{}

// Verify checks fields of this config to ensure a valid fee configuration is provided.
func (f *FeeConfig) Verify() error { _ = "STUB: not implemented"; return nil }

// Equal checks if given [other] is same with this FeeConfig.
func (f *FeeConfig) Equal(other *FeeConfig) bool { _ = "STUB: not implemented"; return false }

// checkByteLens checks byte lengths against common.HashLen (32 bytes) and returns error
func (f *FeeConfig) checkByteLens() error { _ = "STUB: not implemented"; return nil }

func isBiggerThanHashLen(bigint *big.Int) bool { _ = "STUB: not implemented"; return false }
