// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/ava-labs/avalanchego/vms/avm/network"
)

var DefaultConfig = Config{
	Network:          network.DefaultConfig,
	ChecksumsEnabled: false,
}

type Config struct {
	Network          network.Config `json:"network"`
	ChecksumsEnabled bool           `json:"checksums-enabled"`
}

func ParseConfig(configBytes []byte) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}
