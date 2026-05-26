// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nativeminter

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"

	_ "embed"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
)

const (
	mintInputLen = common.HashLength + common.HashLength

	MintGasCost = 30_000
)

type MintNativeCoinInput struct {
	Addr   common.Address
	Amount *big.Int
}

var (
	// Singleton StatefulPrecompiledContract for minting native assets by permissioned callers.
	ContractNativeMinterPrecompile contract.StatefulPrecompiledContract = createNativeMinterPrecompile()

	ErrCannotMint  = errors.New("non-enabled cannot mint")
	ErrInvalidLen  = errors.New("invalid input length for minting")
	ErrUnpackInput = errors.New("failed to unpack input")

	// NativeMinterRawABI contains the raw ABI of NativeMinter contract.
	//go:embed INativeMinter.abi
	NativeMinterRawABI string

	NativeMinterABI = contract.ParseABI(NativeMinterRawABI)
)

// GetContractNativeMinterStatus returns the role of [address] for the minter list.
func GetContractNativeMinterStatus(stateDB contract.StateDB, address common.Address) allowlist.Role {
	_ = "STUB: not implemented"
	return *new(allowlist.Role)
}

// SetContractNativeMinterStatus sets the permissions of [address] to [role] for the
// minter list. assumes [role] has already been verified as valid.
func SetContractNativeMinterStatus(stateDB contract.StateDB, address common.Address, role allowlist.Role) {
	_ = "STUB: not implemented"
	return
}

// PackMintNativeCoin packs [address] and [amount] into the appropriate arguments for mintNativeCoin.
func PackMintNativeCoin(address common.Address, amount *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackMintNativeCoinInput attempts to unpack [input] as address and amount.
// assumes that [input] does not include selector (omits first 4 func signature bytes)
// if [useStrictMode] is true, it will return an error if the length of [input] is not [mintInputLen]
func UnpackMintNativeCoinInput(input []byte, useStrictMode bool) (common.Address, *big.Int, error) {
	_ = "STUB: not implemented"
	// Initially we had this check to ensure that the input was the correct length.
	// However solidity does not always pack the input to the correct length, and allows
	// for extra padding bytes to be added to the end of the input. Therefore, we have removed
	// this check with Durango. We still need to keep this check for backwards compatibility.
	return *new(common.Address), nil, nil
}

// mintNativeCoin checks if the caller is permissioned for minting operation.
// The execution function parses the [input] into native coin amount and receiver address.
func mintNativeCoin(accessibleState contract.AccessibleState, caller common.Address, _ common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Verify that the caller is in the allow list and therefore has the right to call this function.

// if there is no address in the state, create one.

// Return an empty output and the remaining gas

// createNativeMinterPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.
// Access to the getters/setters is controlled by an allow list for ContractAddress.
func createNativeMinterPrecompile() contract.StatefulPrecompiledContract {
	_ = "STUB: not implemented"
	return *new(contract.StatefulPrecompiledContract)
}

// Construct the contract with no fallback function.
