// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertextest

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
)

var (
	errGet                = errors.New("unexpectedly called Get")
	errEdge               = errors.New("unexpectedly called Edge")
	errStopVertexAccepted = errors.New("unexpectedly called StopVertexAccepted")

	_ vertex.Storage = (*Storage)(nil)
)

type Storage struct {
	T                                            *testing.T
	CantGetVtx, CantEdge, CantStopVertexAccepted bool
	GetVtxF                                      func(context.Context, ids.ID) (avalanche.Vertex, error)
	EdgeF                                        func(context.Context) []ids.ID
	StopVertexAcceptedF                          func(context.Context) (bool, error)
}

func (s *Storage) Default(cant bool) { _ = "STUB: not implemented"; return }

func (s *Storage) GetVtx(ctx context.Context, vtxID ids.ID) (avalanche.Vertex, error) {
	_ = "STUB: not implemented"
	return *new(avalanche.Vertex), nil
}

func (s *Storage) Edge(ctx context.Context) []ids.ID { _ = "STUB: not implemented"; return nil }

func (s *Storage) StopVertexAccepted(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
