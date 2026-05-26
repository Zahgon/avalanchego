// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/gas"
)

const Alias = "P"

type Context struct {
	NetworkID         uint32
	AVAXAssetID       ids.ID
	ComplexityWeights gas.Dimensions
	GasPrice          gas.Price
}

func NewSnowContext(networkID uint32, avaxAssetID ids.ID) (*snow.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
