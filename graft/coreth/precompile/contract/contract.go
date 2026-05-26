// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package contract

import (
	"github.com/ava-labs/libevm/common"
)

const (
	SelectorLen = 4
)

type RunStatefulPrecompileFunc func(accessibleState AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error)

// ActivationFunc defines a function that is used to determine if a function is active
// The return value is whether or not the function is active
type ActivationFunc func(AccessibleState) bool

// StatefulPrecompileFunction defines a function implemented by a stateful precompile
type StatefulPrecompileFunction struct {
	// selector is the 4 byte function selector for this function
	selector []byte
	// execute is performed when this function is selected
	execute RunStatefulPrecompileFunc
	// activation is checked before this function is executed
	activation ActivationFunc
}

func (f *StatefulPrecompileFunction) IsActivated(accessibleState AccessibleState) bool {
	_ = "STUB: not implemented"
	return false
}

// NewStatefulPrecompileFunction creates a stateful precompile function with the given arguments
func NewStatefulPrecompileFunction(selector []byte, execute RunStatefulPrecompileFunc) *StatefulPrecompileFunction {
	_ = "STUB: not implemented"
	return nil
}

func NewStatefulPrecompileFunctionWithActivator(selector []byte, execute RunStatefulPrecompileFunc, activation ActivationFunc) *StatefulPrecompileFunction {
	_ = "STUB: not implemented"
	return nil
}

// statefulPrecompileWithFunctionSelectors implements StatefulPrecompiledContract by using 4 byte function selectors to pass
// off responsibilities to internal execution functions.
// Note: because we only ever read from [functions] there no lock is required to make it thread-safe.
type statefulPrecompileWithFunctionSelectors struct {
	fallback  RunStatefulPrecompileFunc
	functions map[string]*StatefulPrecompileFunction
}

// NewStatefulPrecompileContract generates new StatefulPrecompile using [functions] as the available functions and [fallback]
// as an optional fallback if there is no input data. Note: the selector of [fallback] will be ignored, so it is required to be left empty.
func NewStatefulPrecompileContract(fallback RunStatefulPrecompileFunc, functions []*StatefulPrecompileFunction) (StatefulPrecompiledContract, error) {
	_ = "STUB: not implemented"
	// Construct the contract and populate [functions].
	return *new(StatefulPrecompiledContract), nil
}

// Run selects the function using the 4 byte function selector at the start of the input and executes the underlying function on the
// given arguments.
func (s *statefulPrecompileWithFunctionSelectors) Run(accessibleState AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	// If there is no input data present, call the fallback function if present.
	return nil, 0, nil
}

// Otherwise, an unexpected input size will result in an error.

// Use the function selector to grab the correct function

// Check if the function is activated
