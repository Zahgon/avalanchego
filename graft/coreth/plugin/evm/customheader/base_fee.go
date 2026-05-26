// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"math/big"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
)

// BaseFee takes the previous header and the timestamp of its child block and
// calculates the expected base fee for the child block.
//
// Prior to AP3, the returned base fee will be nil.
func BaseFee(
	config *extras.ChainConfig,
	parent *types.Header,
	timeMS uint64,
) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prior to AP3 the expected base fee is nil.

// EstimateNextBaseFee attempts to estimate the base fee of a block built at
// `timestamp` on top of `parent`.
//
// If timestamp is before parent.Time or the AP3 activation time, then timestamp
// is set to the maximum of parent.Time and the AP3 activation time.
//
// Warning: This function should only be used in estimation and should not be
// used when calculating the canonical base fee for a block.
func EstimateNextBaseFee(
	config *extras.ChainConfig,
	parent *types.Header,
	timeMS uint64,
) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
