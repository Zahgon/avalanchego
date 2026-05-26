// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database/versiondb"
)

var (
	chainStatePrefix  = []byte("chain")
	blockStatePrefix  = []byte("block")
	heightIndexPrefix = []byte("height")
)

type State interface {
	ChainState
	BlockState
	HeightIndex
}

type state struct {
	ChainState
	BlockState
	HeightIndex
}

func New(db *versiondb.Database) State { _ = "STUB: not implemented"; return *new(State) }

func NewMetered(db *versiondb.Database, namespace string, metrics prometheus.Registerer) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}
