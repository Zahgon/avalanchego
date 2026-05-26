// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
)

var (
	_ lru.Evictable[ids.ID] = (*uniqueVertex)(nil)
	_ avalanche.Vertex      = (*uniqueVertex)(nil)

	errGetParents = errors.New("failed to get parents for vertex")
	errGetHeight  = errors.New("failed to get height for vertex")
	errGetTxs     = errors.New("failed to get txs for vertex")
)

// uniqueVertex acts as a cache for vertices in the database.
//
// If a vertex is loaded, it will have one canonical uniqueVertex. The vertex
// will eventually be evicted from memory, when the uniqueVertex is evicted from
// the cache. If the uniqueVertex has a function called again after this
// eviction, the vertex will be re-loaded from the database.
type uniqueVertex struct {
	serializer *Serializer

	id ids.ID
	v  *vertexState
}

// newUniqueVertex returns a uniqueVertex instance from [b] by checking the cache
// and then parsing the vertex bytes on a cache miss.
func newUniqueVertex(ctx context.Context, s *Serializer, b []byte) (*uniqueVertex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the vtx exists, then the vertex is already known

// If it wasn't in the cache parse the vertex and set it

// If the vertex has already been fetched,
// skip persisting the vertex.

// The vertex is newly parsed, so set the status
// and persist it.

func (vtx *uniqueVertex) refresh() { _ = "STUB: not implemented"; return }

// shallowRefresh checks the cache for the uniqueVertex and gets the
// most up-to-date status for [vtx]
// ensures that the status is up-to-date for this vertex
// inner vertex may be nil after calling shallowRefresh
func (vtx *uniqueVertex) shallowRefresh() { _ = "STUB: not implemented"; return }

// If someone is in the cache, they must be up-to-date

func (vtx *uniqueVertex) Evict() { _ = "STUB: not implemented"; return }

// make sure the parents can be garbage collected

func (vtx *uniqueVertex) setVertex(ctx context.Context, innerVtx vertex.StatelessVertex) error {
	_ = "STUB: not implemented"
	return nil
}

func (vtx *uniqueVertex) persist() error { _ = "STUB: not implemented"; return nil }

func (vtx *uniqueVertex) setStatus(status choices.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func (vtx *uniqueVertex) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (vtx *uniqueVertex) Key() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (vtx *uniqueVertex) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Should never traverse into parents of a decided vertex. Allows for the
// parents to be garbage collected

func (vtx *uniqueVertex) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }

// Should never traverse into parents of a decided vertex. Allows for the
// parents to be garbage collected

// TODO: run performance test to see if shallow refreshing
// (which will mean that refresh must be called in Bytes and Verify)
// improves performance
func (vtx *uniqueVertex) Status() choices.Status {
	_ = "STUB: not implemented"
	return *new(choices.Status)
}

func (vtx *uniqueVertex) Parents() ([]avalanche.Vertex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vtx *uniqueVertex) Height() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (vtx *uniqueVertex) Txs(ctx context.Context) ([]snowstorm.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vtx *uniqueVertex) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (vtx *uniqueVertex) String() string { _ = "STUB: not implemented"; return "" }

//nolint:perfsprint

//nolint:perfsprint

type vertexState struct {
	latest bool

	vtx    vertex.StatelessVertex
	status choices.Status

	parents []avalanche.Vertex
	txs     []snowstorm.Tx
}
