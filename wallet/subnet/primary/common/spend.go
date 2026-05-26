// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

// MatchOwners attempts to match a list of addresses up to the provided
// threshold.
func MatchOwners(
	owners *secp256k1fx.OutputOwners,
	addrs set.Set[ids.ShortID],
	minIssuanceTime uint64,
) ([]uint32, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
