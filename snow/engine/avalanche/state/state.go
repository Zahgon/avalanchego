// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/utils/logging"
)

type state struct {
	serializer *Serializer
	log        logging.Logger

	dbCache cache.Cacher[ids.ID, any]
	db      database.Database
}

// Vertex retrieves the vertex with the given id from cache/disk.
// Returns nil if it's not found.
// TODO this should return an error
func (s *state) Vertex(id ids.ID) vertex.StatelessVertex {
	_ = "STUB: not implemented"
	return *new(vertex.StatelessVertex)
}

// SetVertex persists the vertex to the database and returns an error if it
// fails to write to the db
func (s *state) SetVertex(id ids.ID, vtx vertex.StatelessVertex) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) Status(id ids.ID) choices.Status {
	_ = "STUB: not implemented"
	return *new(choices.Status)
}

// The key was in the database

// SetStatus sets the status of the vertex and returns an error if it fails to write to the db
func (s *state) SetStatus(id ids.ID, status choices.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) Edge(id ids.ID) []ids.ID { _ = "STUB: not implemented"; return nil }

// Cache the miss

// SetEdge sets the frontier and returns an error if it fails to write to the db
func (s *state) SetEdge(id ids.ID, frontier []ids.ID) error { _ = "STUB: not implemented"; return nil }
