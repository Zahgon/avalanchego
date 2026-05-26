// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nativeasset

import (
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/vm"

	"github.com/ava-labs/avalanchego/graft/coreth/precompile/contract"
)

// PrecompiledContractsApricot contains the default set of pre-compiled Ethereum
// contracts used in the Istanbul release and the stateful precompiled contracts
// added for the Avalanche Apricot release.
// Apricot is incompatible with the YoloV3 Release since it does not include the
// BLS12-381 Curve Operations added to the set of precompiled contracts

var (
	GenesisContractAddr    = common.HexToAddress("0x0100000000000000000000000000000000000000")
	NativeAssetBalanceAddr = common.HexToAddress("0x0100000000000000000000000000000000000001")
	NativeAssetCallAddr    = common.HexToAddress("0x0100000000000000000000000000000000000002")
)

// NativeAssetBalance is a precompiled contract used to retrieve the native asset balance
type NativeAssetBalance struct {
	GasCost uint64
}

// PackNativeAssetBalanceInput packs the arguments into the required input data for a transaction to be passed into
// the native asset balance contract.
func PackNativeAssetBalanceInput(address common.Address, assetID common.Hash) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnpackNativeAssetBalanceInput attempts to unpack [input] into the arguments to the native asset balance precompile
func UnpackNativeAssetBalanceInput(input []byte) (common.Address, common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(common.Hash), nil
}

// Run implements StatefulPrecompiledContract
func (b *NativeAssetBalance) Run(accessibleState contract.AccessibleState, _ common.Address, _ common.Address, input []byte, suppliedGas uint64, _ bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	// input: encodePacked(address 20 bytes, assetID 32 bytes)
	return nil, 0, nil
}

// NativeAssetCall atomically transfers a native asset to a recipient address as well as calling that
// address
type NativeAssetCall struct {
	GasCost           uint64
	CallNewAccountGas uint64
}

// PackNativeAssetCallInput packs the arguments into the required input data for a transaction to be passed into
// the native asset contract.
// Assumes that [assetAmount] is non-nil.
func PackNativeAssetCallInput(address common.Address, assetID common.Hash, assetAmount *big.Int, callData []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnpackNativeAssetCallInput attempts to unpack [input] into the arguments to the native asset call precompile
func UnpackNativeAssetCallInput(input []byte) (common.Address, common.Hash, *big.Int, []byte, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(common.Hash), nil, nil, nil
}

// Run implements [contract.StatefulPrecompiledContract]
func (c *NativeAssetCall) Run(accessibleState contract.AccessibleState, caller common.Address, _ common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// This precompile will be wrapped in a libevm `legacy.PrecompiledStatefulContract`, which
// allows for the deprecated pattern of returning remaining gas by calling
// env.UseGas() on the difference between gas in and gas out. Since we call
// UseGas() ourselves, we therefore return `suppliedGas` unchanged to stop
// the legacy wrapper from double-counting spends.

// run implements the contract logic, using `env.Gas()` and `env.UseGas()` in
// place of `suppliedGas` and returning `remainingGas`, respectively. This
// avoids mixing gas-accounting patterns when using `env.Call()`.
func (c *NativeAssetCall) run(env vm.PrecompileEnvironment, stateDB contract.StateDB, caller common.Address, input []byte, readOnly bool) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: it is not possible for a negative `assetAmount` to be passed in here due to the fact that decoding a
// byte slice into a [*big.Int] will always return a positive value, as documented on [big.Int.SetBytes].

// Send `assetAmount` of `assetID` to `to` address

// When an error was returned by the EVM or when setting the creation code
// above we revert to the snapshot and consume any gas remaining. Additionally
// when we're in homestead this also counts for code storage gas errors.

type DeprecatedContract struct{}

func (*DeprecatedContract) Run(_ contract.AccessibleState, _ common.Address, _ common.Address, _ []byte, suppliedGas uint64, _ bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
