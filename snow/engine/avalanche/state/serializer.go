// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package state manages the meta-data required by consensus for an avalanche
// dag.
package state

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

const (
	dbCacheSize = 10000
	idCacheSize = 1000
)

var (
	errUnknownVertex = errors.New("unknown vertex")
	errWrongChainID  = errors.New("wrong ChainID in vertex")
)

var _ vertex.Manager = (*Serializer)(nil)

// Serializer manages the state of multiple vertices
type Serializer struct {
	SerializerConfig
	versionDB *versiondb.Database
	state     *prefixedState
	edge      set.Set[ids.ID]
}

type SerializerConfig struct {
	ChainID ids.ID
	VM      vertex.DAGVM
	DB      database.Database
	Log     logging.Logger
}

func NewSerializer(config SerializerConfig) vertex.Manager {
	_ = "STUB: not implemented"
	return *new(vertex.Manager)
}

func (s *Serializer) ParseVtx(ctx context.Context, b []byte) (avalanche.Vertex, error) {
	_ = "STUB: not implemented"
	return *new(avalanche.Vertex), nil
}

func (s *Serializer) BuildStopVtx(
	ctx context.Context,
	parentIDs []ids.ID,
) (avalanche.Vertex, error) {
	_ = "STUB: not implemented"
	return *new(avalanche.Vertex), nil
}

// setVertex handles the case where this vertex already exists even
// though we just made it

func (s *Serializer) GetVtx(_ context.Context, vtxID ids.ID) (avalanche.Vertex, error) {
	_ = "STUB: not implemented"
	return *new(avalanche.Vertex), nil
}

func (s *Serializer) Edge(context.Context) []ids.ID { _ = "STUB: not implemented"; return nil }

func (s *Serializer) parseVertex(b []byte) (vertex.StatelessVertex, error) {
	_ = "STUB: not implemented"
	return *new(vertex.StatelessVertex), nil
}

func (s *Serializer) getUniqueVertex(vtxID ids.ID) (*uniqueVertex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Serializer) StopVertexAccepted(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
