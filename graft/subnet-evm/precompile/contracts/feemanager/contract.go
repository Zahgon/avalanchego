// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package feemanager

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"

	_ "embed"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
)

const (
	minFeeConfigFieldKey = iota + 1
	// add new fields below this
	// must preserve order of these fields
	gasLimitKey = iota
	targetBlockRateKey
	minBaseFeeKey
	targetGasKey
	baseFeeChangeDenominatorKey
	minBlockGasCostKey
	maxBlockGasCostKey
	blockGasCostStepKey
	// add new fields above this
	numFeeConfigField = iota - 1

	// [numFeeConfigField] fields in FeeConfig struct
	feeConfigInputLen = common.HashLength * numFeeConfigField

	SetFeeConfigGasCost     uint64 = contract.WriteGasCostPerSlot * (numFeeConfigField + 1) // plus one for setting last changed at
	GetFeeConfigGasCost     uint64 = contract.ReadGasCostPerSlot * numFeeConfigField
	GetLastChangedAtGasCost uint64 = contract.ReadGasCostPerSlot
)

var (

	// Singleton StatefulPrecompiledContract for setting fee configs by permissioned callers.
	FeeManagerPrecompile contract.StatefulPrecompiledContract = createFeeManagerPrecompile()

	feeConfigLastChangedAtKey = common.Hash{'l', 'c', 'a'}

	ErrCannotChangeFee = errors.New("non-enabled cannot change fee config")
	ErrInvalidLen      = errors.New("invalid input length for fee config Input")
	ErrUnpackInput     = errors.New("failed to unpack input")
	ErrUnpackOutput    = errors.New("failed to unpack output")

	// IFeeManagerRawABI contains the raw ABI of FeeManager contract.
	//go:embed IFeeManager.abi
	FeeManagerRawABI string

	FeeManagerABI = contract.ParseABI(FeeManagerRawABI)
)

// FeeConfigABIStruct is the ABI struct for FeeConfig type.
type FeeConfigABIStruct struct {
	GasLimit                 *big.Int
	TargetBlockRate          *big.Int
	MinBaseFee               *big.Int
	TargetGas                *big.Int
	BaseFeeChangeDenominator *big.Int
	MinBlockGasCost          *big.Int
	MaxBlockGasCost          *big.Int
	BlockGasCostStep         *big.Int
}

// GetFeeManagerStatus returns the role of [address] for the fee config manager list.
func GetFeeManagerStatus(stateDB contract.StateReader, address common.Address) allowlist.Role {
	_ = "STUB: not implemented"
	return *new(allowlist.Role)
}

// SetFeeManagerStatus sets the permissions of [address] to [role] for the
// fee config manager list. assumes [role] has already been verified as valid.
func SetFeeManagerStatus(stateDB contract.StateDB, address common.Address, role allowlist.Role) {
	_ = "STUB: not implemented"
	return
}

// GetStoredFeeConfig returns fee config from contract storage in given state
func GetStoredFeeConfig(stateDB contract.StateReader) commontype.FeeConfig {
	_ = "STUB: not implemented"
	return *new(commontype.FeeConfig)
}

// This should never encounter an unknown fee config key

func GetFeeConfigLastChangedAt(stateDB contract.StateReader) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// StoreFeeConfig stores given [feeConfig] and block number in the [blockContext] to the [stateDB].
// A validation on [feeConfig] is done before storing.
func StoreFeeConfig(stateDB contract.StateDB, feeConfig commontype.FeeConfig, blockContext contract.ConfigurationBlockContext) error {
	_ = "STUB: not implemented"
	return nil
}

// This should never encounter an unknown fee config key

// PackSetFeeConfig packs [inputStruct] of type SetFeeConfigInput into the appropriate arguments for setFeeConfig.
func PackSetFeeConfig(input commontype.FeeConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackSetFeeConfigInput attempts to unpack [input] as SetFeeConfigInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
// if [useStrictMode] is true, it will return an error if the length of [input] is not [feeConfigInputLen]
func UnpackSetFeeConfigInput(input []byte, useStrictMode bool) (commontype.FeeConfig, error) {
	_ = "STUB: not implemented"
	// Initially we had this check to ensure that the input was the correct length.
	// However solidity does not always pack the input to the correct length, and allows
	// for extra padding bytes to be added to the end of the input. Therefore, we have removed
	// this check with the Durango. We still need to keep this check for backwards compatibility.
	return *new(commontype.FeeConfig), nil
}

// setFeeConfig checks if the caller has permissions to set the fee config.
// The execution function parses [input] into FeeConfig structure and sets contract storage accordingly.
func setFeeConfig(accessibleState contract.AccessibleState, caller common.Address, _ common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// do not use strict mode after Durango

// Verify that the caller is in the allow list and therefore has the right to call this function.

// Return an empty output and the remaining gas

// PackGetFeeConfig packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetFeeConfig() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PackGetFeeConfigOutput attempts to pack given [outputStruct] of type GetFeeConfigOutput
// to conform the ABI outputs.
func PackGetFeeConfigOutput(output commontype.FeeConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackGetFeeConfigOutput attempts to unpack [output] as GetFeeConfigOutput
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackGetFeeConfigOutput(output []byte, skipLenCheck bool) (commontype.FeeConfig, error) {
	_ = "STUB: not implemented"
	return *new(commontype.FeeConfig), nil
}

// getFeeConfig returns the stored fee config as an output.
// The execution function reads the contract state for the stored fee config and returns the output.
//
//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func getFeeConfig(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Return the fee config as output and the remaining gas

// PackGetFeeConfigLastChangedAt packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetFeeConfigLastChangedAt() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PackGetFeeConfigLastChangedAtOutput attempts to pack given blockNumber of type *big.Int
// to conform the ABI outputs.
func PackGetFeeConfigLastChangedAtOutput(blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackGetFeeConfigLastChangedAtOutput attempts to unpack given [output] into the *big.Int type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackGetFeeConfigLastChangedAtOutput(output []byte) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getFeeConfigLastChangedAt returns the block number that fee config was last changed in.
// The execution function reads the contract state for the stored block number and returns the output.
//
//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func getFeeConfigLastChangedAt(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// createFeeManagerPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.
// Access to the getters/setters is controlled by an allow list for ContractAddress.
func createFeeManagerPrecompile() contract.StatefulPrecompiledContract {
	_ = "STUB: not implemented"
	return *new(contract.StatefulPrecompiledContract)
}

// Construct the contract with no fallback function.
