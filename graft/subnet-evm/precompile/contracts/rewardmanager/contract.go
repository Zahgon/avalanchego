// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated
// This file is a generated precompile contract with stubbed abstract functions.

package rewardmanager

import (
	_ "embed"
	"errors"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"

	"github.com/ava-labs/libevm/common"
)

const (
	AllowFeeRecipientsGasCost      uint64 = contract.WriteGasCostPerSlot + allowlist.ReadAllowListGasCost // write 1 slot + read allow list
	AreFeeRecipientsAllowedGasCost uint64 = allowlist.ReadAllowListGasCost
	CurrentRewardAddressGasCost    uint64 = allowlist.ReadAllowListGasCost
	DisableRewardsGasCost          uint64 = contract.WriteGasCostPerSlot + allowlist.ReadAllowListGasCost // write 1 slot + read allow list
	SetRewardAddressGasCost        uint64 = contract.WriteGasCostPerSlot + allowlist.ReadAllowListGasCost // write 1 slot + read allow list
)

// Singleton StatefulPrecompiledContract and signatures.
var (
	ErrCannotAllowFeeRecipients      = errors.New("non-enabled cannot call allowFeeRecipients")
	ErrCannotAreFeeRecipientsAllowed = errors.New("non-enabled cannot call areFeeRecipientsAllowed")
	ErrCannotCurrentRewardAddress    = errors.New("non-enabled cannot call currentRewardAddress")
	ErrCannotDisableRewards          = errors.New("non-enabled cannot call disableRewards")
	ErrCannotSetRewardAddress        = errors.New("non-enabled cannot call setRewardAddress")

	ErrCannotEnableBothRewards = errors.New("cannot enable both fee recipients and reward address at the same time")
	ErrEmptyRewardAddress      = errors.New("reward address cannot be empty")
	ErrInvalidLen              = errors.New("invalid input length for setting reward address")

	// RewardManagerRawABI contains the raw ABI of RewardManager contract.
	//go:embed IRewardManager.abi
	RewardManagerRawABI string

	RewardManagerABI        = contract.ParseABI(RewardManagerRawABI)
	RewardManagerPrecompile = createRewardManagerPrecompile() // will be initialized by init function

	rewardAddressStorageKey        = common.Hash{'r', 'a', 's', 'k'}
	allowFeeRecipientsAddressValue = common.Hash{'a', 'f', 'r', 'a', 'v'}
)

// GetRewardManagerAllowListStatus returns the role of [address] for the RewardManager list.
func GetRewardManagerAllowListStatus(stateDB contract.StateDB, address common.Address) allowlist.Role {
	_ = "STUB: not implemented"
	return *new(allowlist.Role)
}

// SetRewardManagerAllowListStatus sets the permissions of [address] to [role] for the
// RewardManager list. Assumes [role] has already been verified as valid.
func SetRewardManagerAllowListStatus(stateDB contract.StateDB, address common.Address, role allowlist.Role) {
	_ = "STUB: not implemented"
	return
}

// PackAllowFeeRecipients packs the function selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackAllowFeeRecipients() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// EnableAllowFeeRecipients enables fee recipients.
func EnableAllowFeeRecipients(stateDB contract.StateDB) { _ = "STUB: not implemented"; return }

// DisableFeeRewards disables rewards and burns them by sending to Blackhole Address.
func DisableFeeRewards(stateDB contract.StateDB) { _ = "STUB: not implemented"; return }

func allowFeeRecipients(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// no input provided for this function

// Allow list is enabled and AllowFeeRecipients is a state-changer function.
// This part of the code restricts the function to be called only by enabled/admin addresses in the allow list.
// You can modify/delete this code if you don't want this function to be restricted by the allow list.

// Verify that the caller is in the allow list and therefore has the right to call this function.

// allow list code ends here.

// Return the packed output and the remaining gas

// PackAreFeeRecipientsAllowed packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackAreFeeRecipientsAllowed() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PackAreFeeRecipientsAllowedOutput attempts to pack given isAllowed of type bool
// to conform the ABI outputs.
func PackAreFeeRecipientsAllowedOutput(isAllowed bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func areFeeRecipientsAllowed(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// no input provided for this function

// Return the packed output and the remaining gas

// PackCurrentRewardAddress packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackCurrentRewardAddress() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PackCurrentRewardAddressOutput attempts to pack given rewardAddress of type common.Address
// to conform the ABI outputs.
func PackCurrentRewardAddressOutput(rewardAddress common.Address) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStoredRewardAddress returns the current value of the address stored under rewardAddressStorageKey.
// Returns an empty address and true if allow fee recipients is enabled, otherwise returns current reward address and false.
func GetStoredRewardAddress(stateDB contract.StateReader) (common.Address, bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), false
}

// StoreRewardAddress stores the given [val] under rewardAddressStorageKey.
func StoreRewardAddress(stateDB contract.StateDB, val common.Address) {
	_ = "STUB: not implemented"
	return
}

// PackSetRewardAddress packs [addr] of type common.Address into the appropriate arguments for setRewardAddress.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackSetRewardAddress(addr common.Address) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackSetRewardAddressInput attempts to unpack [input] into the common.Address type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
// if [useStrictMode] is true, it will return an error if the length of [input] is not divisible by 32
func UnpackSetRewardAddressInput(input []byte, useStrictMode bool) (common.Address, error) {
	_ = "STUB: not implemented"
	// Solidity does not always pack the input to the correct length, and allows
	// for extra padding bytes to be added to the end of the input. Therefore, we have removed
	// this check with Durango. We still need to keep this check for backwards compatibility.
	// However, as opposed to other precompiles, we only check that the length is divisible by 32,
	// since historical execution didn't enforce any particular length.
	return *new(common.Address), nil
}

func setRewardAddress(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// attempts to unpack [input] into the arguments to the SetRewardAddressInput.
// Assumes that [input] does not include selector
// do not use strict mode after Durango

// Allow list is enabled and SetRewardAddress is a state-changer function.
// This part of the code restricts the function to be called only by enabled/admin addresses in the allow list.
// You can modify/delete this code if you don't want this function to be restricted by the allow list.

// Verify that the caller is in the allow list and therefore has the right to call this function.

// allow list code ends here.

// if input is empty, return an error

// Add a log to be handled if this action is finalized.

// Return the packed output and the remaining gas

func currentRewardAddress(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// no input provided for this function

// Return the packed output and the remaining gas

// PackDisableRewards packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackDisableRewards() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func disableRewards(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// no input provided for this function

// Allow list is enabled and DisableRewards is a state-changer function.
// This part of the code restricts the function to be called only by enabled/admin addresses in the allow list.
// You can modify/delete this code if you don't want this function to be restricted by the allow list.

// Verify that the caller is in the allow list and therefore has the right to call this function.

// allow list code ends here.

// Return the packed output and the remaining gas

// createRewardManagerPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.
// Access to the getters/setters is controlled by an allow list for [precompileAddr].
func createRewardManagerPrecompile() contract.StatefulPrecompiledContract {
	_ = "STUB: not implemented"
	return *new(contract.StatefulPrecompiledContract)
}

// Construct the contract with no fallback function.
