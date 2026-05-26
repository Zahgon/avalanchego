// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package versionjson

import (
	"github.com/spf13/cobra"

	"github.com/ava-labs/avalanchego/ids"
)

type vmVersions struct {
	Name       string `json:"name"`
	VMID       ids.ID `json:"vmid"`
	Version    string `json:"version"`
	RPCChainVM uint64 `json:"rpcchainvm"`
}

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func versionFunc(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }
