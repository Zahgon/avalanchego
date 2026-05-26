// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var errNilMintOperation = errors.New("nil mint operation")

type MintOperation struct {
	MintInput      Input          `serialize:"true" json:"mintInput"`
	MintOutput     MintOutput     `serialize:"true" json:"mintOutput"`
	TransferOutput TransferOutput `serialize:"true" json:"transferOutput"`
}

func (op *MintOperation) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (op *MintOperation) Cost() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (op *MintOperation) Outs() []verify.State { _ = "STUB: not implemented"; return nil }

func (op *MintOperation) Verify() error { _ = "STUB: not implemented"; return nil }
