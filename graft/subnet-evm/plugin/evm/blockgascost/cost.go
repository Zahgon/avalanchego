// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// blockgascost implements the block gas cost logic
package blockgascost

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
)

// BlockGasCost calculates the required block gas cost.
//
// cost = parentCost + step * (TargetBlockRate - timeElapsed)
//
// The returned cost is clamped to [MinBlockGasCost, MaxBlockGasCost].
func BlockGasCost(
	feeConfig commontype.FeeConfig,
	parentCost uint64,
	step uint64,
	timeElapsed uint64,
) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// This is technically dead code because [MinBlockGasCost] is 0, but it
// makes the code more clear.
