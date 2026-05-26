// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
)

const (
	vtxID uint64 = iota
	vtxStatusID
	edgeID
)

var uniqueEdgeID = ids.Empty.Prefix(edgeID)

type prefixedState struct {
	state *state

	vtx, status cache.Cacher[ids.ID, ids.ID]
	uniqueVtx   *lru.Deduplicator[ids.ID, *uniqueVertex]
}

func newPrefixedState(state *state, idCacheSizes int) *prefixedState {
	_ = "STUB: not implemented"
	return nil
}

func (s *prefixedState) UniqueVertex(vtx *uniqueVertex) *uniqueVertex {
	_ = "STUB: not implemented"
	return nil
}

func (s *prefixedState) Vertex(id ids.ID) vertex.StatelessVertex {
	_ = "STUB: not implemented"
	return *new(vertex.StatelessVertex)
}

func (s *prefixedState) SetVertex(vtx vertex.StatelessVertex) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *prefixedState) Status(id ids.ID) choices.Status {
	_ = "STUB: not implemented"
	return *new(choices.Status)
}

func (s *prefixedState) SetStatus(id ids.ID, status choices.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *prefixedState) Edge() []ids.ID { _ = "STUB: not implemented"; return nil }

func (s *prefixedState) SetEdge(frontier []ids.ID) error { _ = "STUB: not implemented"; return nil }
