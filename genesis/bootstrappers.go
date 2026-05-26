// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/json"
	"fmt"
	"net/netip"

	_ "embed"

	"github.com/ava-labs/avalanchego/ids"
)

var (
	//go:embed bootstrappers.json
	bootstrappersPerNetworkJSON []byte

	bootstrappersPerNetwork map[string][]Bootstrapper
)

func init() {
	if err := json.Unmarshal(bootstrappersPerNetworkJSON, &bootstrappersPerNetwork); err != nil {
		panic(fmt.Sprintf("failed to decode bootstrappers.json %v", err))
	}
}

// Represents the relationship between the nodeID and the nodeIP.
// The bootstrapper is sometimes called "anchor" or "beacon" node.
type Bootstrapper struct {
	ID ids.NodeID     `json:"id"`
	IP netip.AddrPort `json:"ip"`
}

// GetBootstrappers returns all default bootstrappers for the provided network.
func GetBootstrappers(networkID uint32) []Bootstrapper { _ = "STUB: not implemented"; return nil }

// SampleBootstrappers returns the some beacons this node should connect to
func SampleBootstrappers(networkID uint32, count int) []Bootstrapper {
	_ = "STUB: not implemented"
	return nil
}
