// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gaspricemanager

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"

	_ "embed"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
)

//go:embed IGasPriceManager.abi
var GasPriceManagerRawABI string

var GasPriceManagerABI = contract.ParseABI(GasPriceManagerRawABI)

const (
	numGasPriceConfigField = 5 // validatorTargetGas, targetGas, staticPricing, minGasPrice, timeToDouble

	getGasPriceConfigGasCost              uint64 = contract.ReadGasCostPerSlot
	getGasPriceConfigLastChangedAtGasCost uint64 = contract.ReadGasCostPerSlot
	// GasPriceConfigUpdated has 2 topics (event sig + indexed sender) and non-indexed
	// data containing old + new GasPriceConfig. Each config has numGasPriceConfigField
	// static fields, each ABI-encoded as a 32-byte word, so the total data size
	// is numGasPriceConfigField * 32 bytes * 2 configs.
	gasPriceConfigUpdatedEventGasCost uint64 = contract.LogGas + contract.LogTopicGas*2 + contract.LogDataGas*numGasPriceConfigField*common.HashLength*2
	setGasPriceConfigGasCost          uint64 = allowlist.ReadAllowListGasCost +
		contract.WriteGasCostPerSlot*2 + // gas price config slot + lastChangedAt slot
		getGasPriceConfigGasCost + // reading old config for event
		gasPriceConfigUpdatedEventGasCost
)

var (
	errCannotSetGasPriceConfig = errors.New("non-enabled cannot call setGasPriceConfig")
	errInvalidABIConfig        = errors.New("failed to convert ABI config")
	errNilBlockNumber          = errors.New("block number cannot be nil")
)

var gasPriceManagerPrecompile = createGasPriceManagerPrecompile()

func createGasPriceManagerPrecompile() contract.StatefulPrecompiledContract {
	_ = "STUB: not implemented"
	return *new(contract.StatefulPrecompiledContract)
}

// nil = no fallback

// getGasPriceConfig

//nolint:revive // unused params are part of RunStatefulPrecompileFunc signature
func getGasPriceConfig(
	accessibleState contract.AccessibleState,
	caller common.Address,
	self common.Address, // EVM-semantic self; see libevm.AddressContext
	input []byte, // ignored
	suppliedGas uint64,
	readOnly bool, // ignored - method only reads
) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// PackGetGasPriceConfig packs the getGasPriceConfig calldata including the 4-byte selector.
func PackGetGasPriceConfig() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PackGetGasPriceConfigOutput ABI-encodes [config] as getGasPriceConfig return data.
func PackGetGasPriceConfigOutput(config commontype.GasPriceConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackGetGasPriceConfigOutput decodes ABI-encoded getGasPriceConfig return data.
func UnpackGetGasPriceConfigOutput(output []byte) (commontype.GasPriceConfig, error) {
	_ = "STUB: not implemented"
	return *new(commontype.GasPriceConfig), nil
}

// getGasPriceConfigLastChangedAt

//nolint:revive // unused params are part of RunStatefulPrecompileFunc signature
func getGasPriceConfigLastChangedAt(
	accessibleState contract.AccessibleState,
	caller common.Address,
	self common.Address, // EVM-semantic self; see libevm.AddressContext
	input []byte, // ignored
	suppliedGas uint64,
	readOnly bool, // ignored - method only reads
) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// PackGetGasPriceConfigLastChangedAt packs the calldata including the 4-byte selector.
func PackGetGasPriceConfigLastChangedAt() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PackGetGasPriceConfigLastChangedAtOutput ABI-encodes [blockNumber] as return data.
func PackGetGasPriceConfigLastChangedAtOutput(blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackGetGasPriceConfigLastChangedAtOutput decodes ABI-encoded return data.
func UnpackGetGasPriceConfigLastChangedAtOutput(output []byte) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setGasPriceConfig

func setGasPriceConfig(
	accessibleState contract.AccessibleState,
	caller common.Address,
	self common.Address, // EVM-semantic self; see libevm.AddressContext
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// PackSetGasPriceConfig packs [config] into ABI-encoded calldata including the 4-byte selector.
func PackSetGasPriceConfig(config commontype.GasPriceConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackSetGasPriceConfigInput assumes [input] does not include the 4-byte selector.
func UnpackSetGasPriceConfigInput(input []byte) (commontype.GasPriceConfig, error) {
	_ = "STUB: not implemented"
	// UnpackInputIntoInterface doesn't work for single-tuple arguments: copyAtomic
	// tries to assign the whole tuple to the first struct field (a bool), causing a
	// type mismatch. Use method.Inputs.Unpack + abi.ConvertType instead.
	return *new(commontype.GasPriceConfig), nil
}

// storageSlot returns a storage key with the "gasprm" namespace prefix
// left-aligned in the hash. This avoids collisions with AllowList role
// storage, which right-aligns 20-byte addresses via BytesToHash and
// therefore always has 12 leading zero bytes.
func storageSlot(key ...byte) common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

var (
	gasPriceConfigStorageKey       = storageSlot('g', 'p')
	gasPriceConfigLastChangedAtKey = storageSlot('l', 'c', 'a')
)

// GetGasPriceManagerAllowListStatus returns the role of `address` for the allowlist.
func GetGasPriceManagerAllowListStatus(stateDB contract.StateReader, contractAddr common.Address, address common.Address) allowlist.Role {
	_ = "STUB: not implemented"
	return *new(allowlist.Role)
}

// SetGasPriceManagerAllowListStatus assumes [role] has already been verified as valid.
// Roles are stored keyed by address hash, so precompile storage keys must not collide
// with address-derived keys.
func SetGasPriceManagerAllowListStatus(stateDB contract.StateDB, contractAddr common.Address, address common.Address, role allowlist.Role) {
	_ = "STUB: not implemented"
	return
}

// GetStoredGasPriceConfig returns the gas price config from contract storage.
// Configure always stores a value during activation, so the caller
// MUST NOT call this before the precompile has been configured.
func GetStoredGasPriceConfig(stateDB contract.StateReader, contractAddr common.Address) commontype.GasPriceConfig {
	_ = "STUB: not implemented"
	return *new(commontype.GasPriceConfig)
}

// GetGasPriceConfigLastChangedAt returns the block number of the last gas price config update.
func GetGasPriceConfigLastChangedAt(stateDB contract.StateReader, contractAddr common.Address) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// StoreGasPriceConfig validates and persists gasPriceConfig and blockNumber to contract storage.
func StoreGasPriceConfig(stateDB contract.StateDB, contractAddr common.Address, gasPriceConfig commontype.GasPriceConfig, blockNumber *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}
