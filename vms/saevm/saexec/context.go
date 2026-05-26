// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saexec

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/consensus"
	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/utils/logging"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

var _ core.ChainContext = (*chainContext)(nil)

type chainContext struct {
	headers saetypes.HeaderSource
	recent  *lru.Cache[uint64, *types.Header]
	log     logging.Logger
}

func (c *chainContext) GetHeader(h common.Hash, n uint64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

// eth_call on historical state will miss the cache but we still need to
// support BLOCKHASH.

// We explicitly DO NOT populate the cache with these historical values
// because they'll evict the recent headers, which are populated by
// [Executor.execute] for use by BLOCKHASH in newly executed blocks.

func (c *chainContext) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	// This is serious enough that it needs to be investigated immediately, but
	// not enough to be fatal. It will also cause tests to fail if ever called,
	// so we can catch it early.
	return *new(consensus.Engine)
}
