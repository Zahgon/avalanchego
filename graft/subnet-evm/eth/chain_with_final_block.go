// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//nolint:unused
package eth

import (
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
)

const blocksToKeep = 604_800 // Approx. 2 weeks worth of blocks assuming 2s block time

type chainWithFinalBlock struct {
	*core.BlockChain
}

// CurrentFinalBlock returns the current block below which blobs should not
// be maintained anymore for reorg purposes.
func (c *chainWithFinalBlock) CurrentFinalBlock() *types.Header {
	_ = "STUB: not implemented"
	return nil
}
