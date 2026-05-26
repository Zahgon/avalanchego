// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ block.Parser  = (*parseAcceptor)(nil)
	_ snowman.Block = (*blockAcceptor)(nil)
)

type parseAcceptor struct {
	parser      block.Parser
	ctx         *snow.ConsensusContext
	numAccepted prometheus.Counter
}

func (p *parseAcceptor) ParseBlock(ctx context.Context, bytes []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

type blockAcceptor struct {
	snowman.Block

	ctx         *snow.ConsensusContext
	numAccepted prometheus.Counter
}

func (b *blockAcceptor) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
