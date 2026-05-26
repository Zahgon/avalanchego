// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"reflect"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

// CodecVersion is the current default codec version
const CodecVersion = txs.CodecVersion

var _ Parser = (*parser)(nil)

type Parser interface {
	txs.Parser

	ParseBlock(bytes []byte) (Block, error)
	ParseGenesisBlock(bytes []byte) (Block, error)
}

type parser struct {
	txs.Parser
}

func NewParser(fxs []fxs.Fx) (Parser, error) { _ = "STUB: not implemented"; return *new(Parser), nil }

func NewCustomParser(
	typeToFxIndex map[reflect.Type]int,
	clock *mockable.Clock,
	log logging.Logger,
	fxs []fxs.Fx,
) (Parser, error) {
	_ = "STUB: not implemented"
	return *new(Parser), nil
}

func (p *parser) ParseBlock(bytes []byte) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func (p *parser) ParseGenesisBlock(bytes []byte) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func parse(cm codec.Manager, bytes []byte) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}
