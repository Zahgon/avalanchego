// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowmantest

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
)

func MakeLastAcceptedBlockF(blks ...[]*Block) func(context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return nil
}

func MakeGetBlockIDAtHeightF(blks ...[]*Block) func(context.Context, uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return nil
}
