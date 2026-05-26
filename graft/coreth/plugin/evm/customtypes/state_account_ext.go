// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customtypes

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"

	ethtypes "github.com/ava-labs/libevm/core/types"
)

type isMultiCoin bool

func IsMultiCoin(s *state.StateDB, addr common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func SetMultiCoin(s *state.StateDB, addr common.Address, to bool) {
	_ = "STUB: not implemented"
	return
}

func IsAccountMultiCoin(s ethtypes.StateOrSlimAccount) bool {
	_ = "STUB: not implemented"
	return false
}
