// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	_ "embed"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

const addWarpMessageBaseGasCost uint64 = 20_000 // Cost of producing and serving a BLS Signature

var (
	preGraniteGasConfig = GasConfig{
		GetBlockchainID: 2,

		GetVerifiedWarpMessageBase: 2,
		PerWarpSigner:              500,
		PerWarpMessageChunk:        3_200,
		VerifyPredicateBase:        200_000,

		// Sum of base log gas cost, cost of producing 3 topics, and
		// producing + serving a BLS Signature (sign + trie write).
		//
		// Note: using trie write for the gas cost results in a conservative
		// overestimate since the message is stored in a flat database that can
		// be cleaned up after a period of time instead of the EVM trie.
		SendWarpMessageBase: contract.LogGas + 3*contract.LogTopicGas + addWarpMessageBaseGasCost + contract.WriteGasCostPerSlot,
		PerWarpMessageByte:  contract.LogDataGas,
	}
	// graniteGasConfig updated the gas costs of warp operations to target a
	// processing speed of around 100 mgas/s after the introduction of epoching
	// which allows for pre-calculating the warp validator sets for all active
	// networks.
	graniteGasConfig = GasConfig{
		GetBlockchainID: 200,

		GetVerifiedWarpMessageBase: 750,
		PerWarpSigner:              250,
		PerWarpMessageChunk:        512, // matches call data byte cost
		VerifyPredicateBase:        125_000,

		// Unchanged during Granite.
		SendWarpMessageBase: preGraniteGasConfig.SendWarpMessageBase,
		PerWarpMessageByte:  preGraniteGasConfig.PerWarpMessageByte,
	}
)

type GasConfig struct {
	// Cost to call getBlockchainID
	GetBlockchainID uint64

	// Base cost of entering getVerifiedWarpMessage
	GetVerifiedWarpMessageBase uint64
	// Gas cost per warp signer in the validator set
	PerWarpSigner uint64
	// Gas cost per chunk of the warp message (each chunk is 32 bytes)
	PerWarpMessageChunk uint64
	// Gas cost to verify a BLS signature
	VerifyPredicateBase uint64

	// Base cost of entering sendWarpMessage
	SendWarpMessageBase uint64
	// PerWarpMessageByte cost accounts for producing a message of a given size
	PerWarpMessageByte uint64
}

var (
	errInvalidSendInput  = errors.New("invalid sendWarpMessage input")
	errInvalidIndexInput = errors.New("invalid index to specify warp message")
)

// Singleton StatefulPrecompiledContract and signatures.
var (
	// WarpRawABI contains the raw ABI of Warp contract.
	//go:embed IWarpMessenger.abi
	WarpRawABI string

	WarpABI = contract.ParseABI(WarpRawABI)

	WarpPrecompile = createWarpPrecompile()
)

// WarpBlockHash is an auto generated low-level Go binding around an user-defined struct.
type WarpBlockHash struct {
	SourceChainID common.Hash
	BlockHash     common.Hash
}

type GetVerifiedWarpBlockHashOutput struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}

// WarpMessage is an auto generated low-level Go binding around an user-defined struct.
type WarpMessage struct {
	SourceChainID       common.Hash
	OriginSenderAddress common.Address
	Payload             []byte
}

type GetVerifiedWarpMessageOutput struct {
	Message WarpMessage
	Valid   bool
}

type SendWarpMessageEventData struct {
	Message []byte
}

// PackGetBlockchainID packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetBlockchainID() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PackGetBlockchainIDOutput attempts to pack given blockchainID of type common.Hash
// to conform the ABI outputs.
func PackGetBlockchainIDOutput(blockchainID common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getBlockchainID returns the snow Chain Context ChainID of this blockchain.
//
//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func getBlockchainID(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Return the packed output and the remaining gas

// UnpackGetVerifiedWarpBlockHashInput attempts to unpack [input] into the uint32 type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackGetVerifiedWarpBlockHashInput(input []byte) (uint32, error) {
	_ = "STUB: not implemented"
	// We don't use strict mode here because it was disabled with Durango.
	// Since Warp will be deployed after Durango, we don't need to validate padding length.
	return 0, nil
}

// PackGetVerifiedWarpBlockHash packs [index] of type uint32 into the appropriate arguments for getVerifiedWarpBlockHash.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetVerifiedWarpBlockHash(index uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PackGetVerifiedWarpBlockHashOutput attempts to pack given [outputStruct] of type GetVerifiedWarpBlockHashOutput
// to conform the ABI outputs.
func PackGetVerifiedWarpBlockHashOutput(outputStruct GetVerifiedWarpBlockHashOutput) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackGetVerifiedWarpBlockHashOutput attempts to unpack [output] as GetVerifiedWarpBlockHashOutput
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackGetVerifiedWarpBlockHashOutput(output []byte) (GetVerifiedWarpBlockHashOutput, error) {
	_ = "STUB: not implemented"
	return *new(GetVerifiedWarpBlockHashOutput), nil
}

//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func getVerifiedWarpBlockHash(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// UnpackGetVerifiedWarpMessageInput attempts to unpack [input] into the uint32 type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackGetVerifiedWarpMessageInput(input []byte) (uint32, error) {
	_ = "STUB: not implemented"
	// We don't use strict mode here because it was disabled with Durango.
	// Since Warp will be deployed after Durango, we don't need to validate padding length.
	return 0, nil
}

// PackGetVerifiedWarpMessage packs [index] of type uint32 into the appropriate arguments for getVerifiedWarpMessage.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetVerifiedWarpMessage(index uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PackGetVerifiedWarpMessageOutput attempts to pack given [outputStruct] of type GetVerifiedWarpMessageOutput
// to conform the ABI outputs.
func PackGetVerifiedWarpMessageOutput(outputStruct GetVerifiedWarpMessageOutput) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackGetVerifiedWarpMessageOutput attempts to unpack [output] as GetVerifiedWarpMessageOutput
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackGetVerifiedWarpMessageOutput(output []byte) (GetVerifiedWarpMessageOutput, error) {
	_ = "STUB: not implemented"
	return *new(GetVerifiedWarpMessageOutput), nil
}

// getVerifiedWarpMessage retrieves the pre-verified warp message from the predicate storage slots and returns
// the expected ABI encoding of the message to the caller.
//
//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func getVerifiedWarpMessage(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// UnpackSendWarpMessageInput attempts to unpack [input] as []byte
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackSendWarpMessageInput(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// We don't use strict mode here because it was disabled with Durango.
	// Since Warp will be deployed after Durango, we don't need to validate padding length.
	return nil, nil
}

// PackSendWarpMessage packs [inputStruct] of type []byte into the appropriate arguments for sendWarpMessage.
func PackSendWarpMessage(payloadData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PackSendWarpMessageOutput attempts to pack given messageID of type common.Hash
// to conform the ABI outputs.
func PackSendWarpMessageOutput(messageID common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackSendWarpMessageOutput attempts to unpack given [output] into the common.Hash type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackSendWarpMessageOutput(output []byte) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// sendWarpMessage constructs an Avalanche Warp Message containing an AddressedPayload and emits a log to signal validators that they should
// be willing to sign this message.
func sendWarpMessage(accessibleState contract.AccessibleState, caller common.Address, _ common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// This gas cost includes buffer room because it is based off of the total size of the input instead of the produced payload.
// This ensures that we charge gas before we unpack the variable sized input.

// unpack the arguments

// Add a log to be handled if this action is finalized.

// Return the packed message ID and the remaining gas

// PackSendWarpMessageEvent packs the given arguments into SendWarpMessage events including topics and data.
func PackSendWarpMessageEvent(sourceAddress common.Address, unsignedMessageID common.Hash, unsignedMessageBytes []byte) ([]common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// UnpackSendWarpEventDataToMessage attempts to unpack event [data] as warp.UnsignedMessage.
func UnpackSendWarpEventDataToMessage(data []byte) (*warp.UnsignedMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createWarpPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.
func createWarpPrecompile() contract.StatefulPrecompiledContract {
	_ = "STUB: not implemented"
	return *new(contract.StatefulPrecompiledContract)
}

// Construct the contract with no fallback function.

func CurrentGasConfig(rules precompileconfig.Rules) GasConfig {
	_ = "STUB: not implemented"
	return *new(GasConfig)
}
