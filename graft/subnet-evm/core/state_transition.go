// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
)

// ExecutionResult includes all output after executing given evm
// message no matter the execution itself is successful or not.
type ExecutionResult struct {
	UsedGas     uint64 // Total used gas, not including the refunded gas
	RefundedGas uint64 // Total gas refunded after execution
	Err         error  // Any error encountered during the execution(listed in core/vm/errors.go)
	ReturnData  []byte // Returned data from evm(function result or data supplied with revert opcode)
}

// Unwrap returns the internal evm error which allows us for further
// analysis outside.
func (result *ExecutionResult) Unwrap() error {
	_ = "STUB: not implemented"

	// Failed returns the indicator whether the execution is successful or not
	return nil
}

func (result *ExecutionResult) Failed() bool { _ = "STUB: not implemented"; return false }

// Return is a helper function to help caller distinguish between revert reason
// and function return. Return returns the data after execution if no error occurs.
func (result *ExecutionResult) Return() []byte { _ = "STUB: not implemented"; return nil }

// Revert returns the concrete revert reason if the execution is aborted by `REVERT`
// opcode. Note the reason can be nil if no data supplied with revert opcode.
func (result *ExecutionResult) Revert() []byte { _ = "STUB: not implemented"; return nil }

// A Message contains the data derived from a single transaction that is relevant to state
// processing.
type Message struct {
	To            *common.Address
	From          common.Address
	Nonce         uint64
	Value         *big.Int
	GasLimit      uint64
	GasPrice      *big.Int
	GasFeeCap     *big.Int
	GasTipCap     *big.Int
	Data          []byte
	AccessList    types.AccessList
	BlobGasFeeCap *big.Int
	BlobHashes    []common.Hash

	// When SkipAccountChecks is true, the message nonce is not checked against the
	// account nonce in state. It also disables checking that the sender is an EOA.
	// This field will be set to true for operations like RPC eth_call.
	SkipAccountChecks bool
}

// TransactionToMessage converts a transaction into a Message.
func TransactionToMessage(tx *types.Transaction, s types.Signer, baseFee *big.Int) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If baseFee provided, set gasPrice to effectiveGasPrice.

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any EVM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(evm *vm.EVM, msg *Message, gp *GasPool) (*ExecutionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StateTransition represents a state transition.
//
// == The State Transitioning Model
//
// A state transition is a change made when a transaction is applied to the current world
// state. The state transitioning model does all the necessary work to work out a valid new
// state root.
//
//  1. Nonce handling
//  2. Pre pay gas
//  3. Create a new state object if the recipient is nil
//  4. Value transfer
//
// == If contract creation ==
//
//	4a. Attempt to run transaction data
//	4b. If valid, use result as code for the new state object
//
// == end ==
//
//  5. Run Script section
//  6. Derive new state root
type StateTransition struct {
	gp           *GasPool
	msg          *Message
	gasRemaining uint64
	initialGas   uint64
	state        vm.StateDB
	evm          *vm.EVM
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(evm *vm.EVM, msg *Message, gp *GasPool) *StateTransition {
	_ = "STUB: not implemented"
	return nil
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

/* contract creation */

func (st *StateTransition) buyGas() error { _ = "STUB: not implemented"; return nil }

// Check that the user has enough funds to cover blobGasUsed * tx.BlobGasFeeCap

// Pay for blobGasUsed * actual blob fee

func (st *StateTransition) preCheck() error {
	_ = "STUB: not implemented"
	// Only check transactions that are not fake
	return nil
}

// Make sure this transaction's nonce is correct.

// Make sure the sender is an EOA

// Check that the sender is on the tx allow list if enabled

// Make sure that transaction gasFeeCap is greater than the baseFee (post london)

// Skip the checks if gas fields are zero and baseFee was explicitly disabled (eth_call)

// This will panic if baseFee is nil, but basefee presence is verified
// as part of header validation.

// Check the blob version validity

// The to field of a blob tx type is mandatory, and a `BlobTx` transaction internally
// has it as a non-nillable value, so any msg derived from blob transaction has it non-nil.
// However, messages created through RPC (eth_call) don't have this restriction.

// Check that the user is paying at least the current blob fee

// Skip the checks if gas fields are zero and blobBaseFee was explicitly disabled (eth_call)

// This will panic if blobBaseFee is nil, but blobBaseFee presence
// is verified as part of header validation.

// TransitionDb will transition the state by applying the current message and
// returning the evm execution result with following fields.
//
//   - used gas: total gas used (including gas being refunded)
//   - returndata: the returned data from evm
//   - concrete execution error: various EVM errors which abort the execution, e.g.
//     ErrOutOfGas, ErrExecutionReverted
//
// However if any consensus issue encountered, return the error directly with
// nil evm execution result.
func (st *StateTransition) TransitionDb() (*ExecutionResult, error) {
	_ = "STUB: not implemented"
	// First check this message satisfies all consensus rules before
	// applying the message. The rules include these clauses
	//
	// 1. the nonce of the message caller is correct
	// 2. caller has enough balance to cover transaction fee(gaslimit * gasprice)
	// 3. the amount of gas required is available in the block
	// 4. the message caller is on the tx allow list (if enabled)
	// 5. the purchased gas is enough to cover intrinsic usage
	// 6. there is no overflow when calculating intrinsic gas
	// 7. caller has enough balance to cover asset transfer for **topmost** call
	return nil, nil
}

// Check clauses 1-4, buy gas if everything is correct

// Check clauses 4-5, subtract intrinsic gas if everything is correct

// Check clause 6

// Check whether the init code size has been exceeded.

// Execute the preparatory steps for state transition which includes:
// - prepare accessList(post-berlin)
// - reset transient storage(eip 1153)

// store in case execution invalidated

// vm errors do not effect consensus and are therefore not assigned to err

// Increment the nonce for the next transaction

func (st *StateTransition) refundGas(subnetEVM bool) uint64 {
	_ = "STUB: not implemented"

	// Inspired by: https://gist.github.com/holiman/460f952716a74eeb9ab358bb1836d821#gistcomment-3642048
	return 0
}

// Apply refund counter, capped to half of the used gas.

// Return ETH for remaining gas, exchanged at the original rate.

// Also return remaining gas to the block gas counter so it is
// available for the next transaction.

// gasUsed returns the amount of gas used up by the state transition.
func (st *StateTransition) gasUsed() uint64 { _ = "STUB: not implemented"; return 0 }

// blobGasUsed returns the amount of blob gas used by the message.
func (st *StateTransition) blobGasUsed() uint64 { _ = "STUB: not implemented"; return 0 }
