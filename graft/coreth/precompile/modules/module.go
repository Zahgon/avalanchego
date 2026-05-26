// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package modules

import (
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/coreth/precompile/contract"
)

type Module struct {
	// ConfigKey is the key used in json config files to specify this precompile config.
	ConfigKey string
	// Address returns the address where the stateful precompile is accessible.
	Address common.Address
	// Contract returns a thread-safe singleton that can be used as the StatefulPrecompiledContract when
	// this config is enabled.
	Contract contract.StatefulPrecompiledContract
	// Configurator is used to configure the stateful precompile when the config is enabled.
	contract.Configurator
}

type moduleArray []Module

func (m moduleArray) Len() int { _ = "STUB: not implemented"; return 0 }

func (m moduleArray) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (m moduleArray) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
