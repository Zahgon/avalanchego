// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"errors"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var errNilMintOperation = errors.New("nil mint operation")

type MintOperation struct {
	MintInput   secp256k1fx.Input `serialize:"true" json:"mintInput"`
	MintOutput  MintOutput        `serialize:"true" json:"mintOutput"`
	OwnedOutput OwnedOutput       `serialize:"true" json:"ownedOutput"`
}

func (op *MintOperation) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (op *MintOperation) Cost() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (op *MintOperation) Outs() []verify.State { _ = "STUB: not implemented"; return nil }

func (op *MintOperation) Verify() error { _ = "STUB: not implemented"; return nil }
