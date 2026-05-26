// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package modules

import (
	"errors"
	"fmt"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/evm/constants"
	"github.com/ava-labs/avalanchego/graft/evm/utils"
)

var (
	// registeredModules is a list of Module to preserve order
	// for deterministic iteration
	registeredModules = make([]Module, 0)

	reservedRanges = []utils.AddressRange{
		{
			Start: common.HexToAddress("0x0100000000000000000000000000000000000000"),
			End:   common.HexToAddress("0x01000000000000000000000000000000000000ff"),
		},
		{
			Start: common.HexToAddress("0x0200000000000000000000000000000000000000"),
			End:   common.HexToAddress("0x02000000000000000000000000000000000000ff"),
		},
		{
			Start: common.HexToAddress("0x0300000000000000000000000000000000000000"),
			End:   common.HexToAddress("0x03000000000000000000000000000000000000ff"),
		},
	}

	errBlackholeAddress          = fmt.Errorf("cannot register module that overlaps with blackhole address %s", constants.BlackholeAddr)
	errAddressNotInReservedRange = errors.New("address is not in a reserved range for custom precompiles")
)

// ReservedAddress returns true if [addr] is in a reserved range for custom precompiles
func ReservedAddress(addr common.Address) bool { _ = "STUB: not implemented"; return false }

// RegisterModule registers a stateful precompile module
func RegisterModule(stm Module) error { _ = "STUB: not implemented"; return nil }

// sort by address to ensure deterministic iteration

func GetPrecompileModuleByAddress(address common.Address) (Module, bool) {
	_ = "STUB: not implemented"
	return *new(Module), false
}

func GetPrecompileModule(key string) (Module, bool) {
	_ = "STUB: not implemented"
	return *new(Module), false
}

func RegisteredModules() []Module { _ = "STUB: not implemented"; return nil }

func insertSortedByAddress(data []Module, stm Module) []Module {
	_ = "STUB: not implemented"
	return nil
}
