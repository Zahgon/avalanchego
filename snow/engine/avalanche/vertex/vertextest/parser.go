// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertextest

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
)

var (
	errParse = errors.New("unexpectedly called Parse")

	_ vertex.Parser = (*Parser)(nil)
)

type Parser struct {
	T            *testing.T
	CantParseVtx bool
	ParseVtxF    func(context.Context, []byte) (avalanche.Vertex, error)
}

func (p *Parser) Default(cant bool) { _ = "STUB: not implemented"; return }

func (p *Parser) ParseVtx(ctx context.Context, b []byte) (avalanche.Vertex, error) {
	_ = "STUB: not implemented"
	return *new(avalanche.Vertex), nil
}
