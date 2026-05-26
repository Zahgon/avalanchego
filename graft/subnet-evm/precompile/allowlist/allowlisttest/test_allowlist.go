// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlisttest

import (
	"testing"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/core/extstate"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/modules"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompiletest"

	ethtypes "github.com/ava-labs/libevm/core/types"
)

var (
	TestAdminAddr   = common.HexToAddress("0x0000000000000000000000000000000000000011")
	TestEnabledAddr = common.HexToAddress("0x0000000000000000000000000000000000000022")
	TestNoRoleAddr  = common.HexToAddress("0x0000000000000000000000000000000000000033")
	TestManagerAddr = common.HexToAddress("0x0000000000000000000000000000000000000044")
)

func AllowListTests(_ testing.TB, module modules.Module) []precompiletest.PrecompileTest {
	_ = "STUB: not implemented"
	return nil
}

// Check logs are stored in state

// Check logs are stored in state

// Check logs are stored in state

// Check logs are stored in state

// Check logs are stored in state

// Check logs are stored in state

// SetDefaultRoles returns a BeforeHook that sets roles TestAdminAddr and TestEnabledAddr
// to have the AdminRole and EnabledRole respectively.
func SetDefaultRoles(contractAddress common.Address) func(t testing.TB, state *extstate.StateDB) {
	_ = "STUB: not implemented"
	return nil
}

func RunPrecompileWithAllowListTests(t *testing.T, module modules.Module, tests []precompiletest.PrecompileTest) {
	_ = "STUB: not implemented"

	// Add the contract specific tests to the map of tests to run.
	return
}

func assertSetRoleEvent(t testing.TB, logs []*ethtypes.Log, role allowlist.Role, addr common.Address, caller common.Address, oldRole allowlist.Role) {
	_ = "STUB: not implemented"
	return
}
