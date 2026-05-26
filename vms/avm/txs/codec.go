// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"reflect"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	_ codec.Registry = (*codecRegistry)(nil)
	_ secp256k1fx.VM = (*fxVM)(nil)
)

type codecRegistry struct {
	codecs      []codec.Registry
	index       int
	typeToIndex map[reflect.Type]int
}

func (cr *codecRegistry) RegisterType(val interface{}) error { _ = "STUB: not implemented"; return nil }

type fxVM struct {
	typeToFxIndex map[reflect.Type]int

	clock         *mockable.Clock
	log           logging.Logger
	codecRegistry codec.Registry
}

func (vm *fxVM) Clock() *mockable.Clock { _ = "STUB: not implemented"; return nil }

func (vm *fxVM) CodecRegistry() codec.Registry {
	_ = "STUB: not implemented"
	return *new(codec.Registry)
}

func (vm *fxVM) Logger() logging.Logger { _ = "STUB: not implemented"; return *new(logging.Logger) }
