// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/evm/acp176"
)

// feeStateBeforeBlock takes the previous header and the timestamp of its child
// block and calculates the fee state before the child block is executed.
func feeStateBeforeBlock(
	config *extras.ChainConfig,
	parent *types.Header,
	timeMS uint64,
) (acp176.State, error) {
	_ = "STUB: not implemented"
	return *new(acp176.State), nil
}

// If the parent block was running with ACP-176, we start with the
// resulting fee state from the parent block. It is assumed that the
// parent has been verified, so the claimed fee state equals the actual
// fee state.

// feeStateAfterBlock takes the previous header and returns the fee state after
// the execution of the provided child.
func feeStateAfterBlock(
	config *extras.ChainConfig,
	parent *types.Header,
	header *types.Header,
	desiredTargetExcess *gas.Gas,
) (acp176.State, error) {
	_ = "STUB: not implemented"
	// Calculate the gas state after the parent block
	return *new(acp176.State), nil
}

// Consume the gas used by the block

// If the desired target excess is specified, move the target excess as much
// as possible toward that desired value.
