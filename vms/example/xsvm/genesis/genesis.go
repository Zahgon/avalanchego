// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/block"
)

type Genesis struct {
	Timestamp   int64        `serialize:"true" json:"timestamp"`
	Allocations []Allocation `serialize:"true" json:"allocations"`
}

type Allocation struct {
	Address ids.ShortID `serialize:"true" json:"address"`
	Balance uint64      `serialize:"true" json:"balance"`
}

func Parse(bytes []byte) (*Genesis, error) { _ = "STUB: not implemented"; return nil, nil }

func Block(genesis *Genesis) (*block.Stateless, error) { _ = "STUB: not implemented"; return nil, nil }
