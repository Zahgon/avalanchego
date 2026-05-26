// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
)

// Builder builds a vertex given a set of parentIDs and transactions.
type Builder interface {
	// Build a new stop vertex from the parents
	BuildStopVtx(ctx context.Context, parentIDs []ids.ID) (avalanche.Vertex, error)
}

// Build a new stateless vertex from the contents of a vertex
func Build(
	chainID ids.ID,
	height uint64,
	parentIDs []ids.ID,
	txs [][]byte,
) (StatelessVertex, error) {
	_ = "STUB: not implemented"
	return *new(StatelessVertex), nil
}

// Build a new stateless vertex from the contents of a vertex
func BuildStopVertex(chainID ids.ID, height uint64, parentIDs []ids.ID) (StatelessVertex, error) {
	_ = "STUB: not implemented"
	return *new(StatelessVertex), nil
}

func buildVtx(
	chainID ids.ID,
	height uint64,
	parentIDs []ids.ID,
	txs [][]byte,
	verifyFunc func(innerStatelessVertex) error,
	stopVertex bool,
) (StatelessVertex, error) {
	_ = "STUB: not implemented"
	return *new(StatelessVertex), nil
}

// use new codec version for the "StopVertex"
