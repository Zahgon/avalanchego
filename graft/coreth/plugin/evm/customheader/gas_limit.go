// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
)

var (
	errInvalidExtraDataGasUsed = errors.New("invalid extra data gas used")
	errInvalidGasUsed          = errors.New("invalid gas used")
	errInvalidGasLimit         = errors.New("invalid gas limit")
)

// GasLimit takes the previous header and the timestamp of its child block and
// calculates the gas limit for the child block.
func GasLimit(
	config *extras.ChainConfig,
	parent *types.Header,
	timeMS uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The gas limit is set to the maximum capacity, rather than the current
// capacity, to minimize the differences with upstream geth. During
// block building and gas usage calculations, the gas limit is checked
// against the current capacity.

// The gas limit prior Apricot Phase 1 started at the genesis value and
// migrated towards the [ap1.GasLimit] following the `core.CalcGasLimit`
// updates. However, since all chains have activated Apricot Phase 1,
// this code is not used in production. To avoid a dependency on the
// `core` package, this code is modified to just return the parent gas
// limit; which was valid to do prior to Apricot Phase 1.

// VerifyGasUsed verifies that the gas used is less than or equal to the gas
// limit.
func VerifyGasUsed(
	config *extras.ChainConfig,
	parent *types.Header,
	header *types.Header,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyGasLimit verifies that the gas limit for the header is valid.
func VerifyGasLimit(
	config *extras.ChainConfig,
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
	parent *types.Header,
	timeMS uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Prior to the F upgrade, the gas capacity is equal to the gas limit.
		nil
}

// RemainingAtomicGasCapacity returns the maximum amount ExtDataGasUsed could be
// on `header` while still being valid based on the initial capacity and
// consumed gas.
func RemainingAtomicGasCapacity(
	config *extras.ChainConfig,
	parent *types.Header,
	header *types.Header,
) (uint64, error) {
	_ = "STUB: not implemented"
	// Prior to the F upgrade, the atomic gas limit was a constant independent
	// of the evm gas used.
	return 0, nil
}
