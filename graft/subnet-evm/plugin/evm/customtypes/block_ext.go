// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customtypes

import (
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/vms/evm/acp226"

	ethtypes "github.com/ava-labs/libevm/core/types"
)

func BlockGasCost(b *ethtypes.Block) *big.Int { _ = "STUB: not implemented"; return nil }

func BlockTimeMilliseconds(b *ethtypes.Block) *uint64 { _ = "STUB: not implemented"; return nil }

func BlockMinDelayExcess(b *ethtypes.Block) *acp226.DelayExcess {
	_ = "STUB: not implemented"
	return nil
}

func BlockTime(eth *ethtypes.Header) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
