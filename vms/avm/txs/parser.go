// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"reflect"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
)

// CodecVersion is the current default codec version
const CodecVersion = 0

var _ Parser = (*parser)(nil)

type Parser interface {
	Codec() codec.Manager
	GenesisCodec() codec.Manager

	CodecRegistry() codec.Registry
	GenesisCodecRegistry() codec.Registry

	ParseTx(bytes []byte) (*Tx, error)
	ParseGenesisTx(bytes []byte) (*Tx, error)
}

type parser struct {
	cm  codec.Manager
	gcm codec.Manager
	c   linearcodec.Codec
	gc  linearcodec.Codec
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

func (p *parser) Codec() codec.Manager { _ = "STUB: not implemented"; return *new(codec.Manager) }

func (p *parser) GenesisCodec() codec.Manager {
	_ = "STUB: not implemented"
	return *new(codec.Manager)
}

func (p *parser) CodecRegistry() codec.Registry {
	_ = "STUB: not implemented"
	return *new(codec.Registry)
}

func (p *parser) GenesisCodecRegistry() codec.Registry {
	_ = "STUB: not implemented"
	return *new(codec.Registry)
}

func (p *parser) ParseTx(bytes []byte) (*Tx, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) ParseGenesisTx(bytes []byte) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parse(cm codec.Manager, signedBytes []byte) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
