// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params/extras"
)

var (
	errInvalidGasUsed  = errors.New("invalid gas used")
	errInvalidGasLimit = errors.New("invalid gas limit")
)

type CalculateGasLimitFunc func(parentGasUsed, parentGasLimit, gasFloor, gasCeil uint64) uint64

// GasLimit takes the previous header and the timestamp of its child block and
// calculates the gas limit for the child block.
func GasLimit(
	config *extras.ChainConfig,
	feeConfig commontype.FeeConfig,
	parent *types.Header,
	timeMS uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// since all chains have activated Subnet-EVM,
// this code is not used in production. To avoid a dependency on the
// `core` package, this code is modified to just return the parent gas
// limit; which was valid to do prior to Subnet-EVM.

// VerifyGasUsed verifies that the gas used is less than or equal to the gas
// limit.
func VerifyGasUsed(
	config *extras.ChainConfig,
	feeConfig commontype.FeeConfig,
	parent *types.Header,
	header *types.Header,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyGasLimit verifies that the gas limit for the header is valid.
func VerifyGasLimit(
	config *extras.ChainConfig,
	feeConfig commontype.FeeConfig,
	parent *types.Header,
	header *types.Header,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify that the gas limit remains within allowed bounds

// GasCapacity takes the previous header and the timestamp of its child block
// and calculates the available gas that can be consumed in the child block.
func GasCapacity(
	config *extras.ChainConfig,
	feeConfig commontype.FeeConfig,
	parent *types.Header,
	timeMS uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
