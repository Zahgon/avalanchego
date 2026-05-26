// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlisttest

import (
	"testing"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

// RunAllowListEventTests runs the standard AllowList event emission tests.
// This can be used by any precompile that uses the AllowList pattern.
func RunAllowListEventTests(
	t *testing.T,
	precompileCfg precompileconfig.Config,
	contractAddress common.Address,
	adminAuth *bind.TransactOpts,
	fundedAddrs ...common.Address,
) {
	_ = "STUB: not implemented"
	return
}
