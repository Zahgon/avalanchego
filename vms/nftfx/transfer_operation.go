// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"errors"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var errNilTransferOperation = errors.New("nil transfer operation")

type TransferOperation struct {
	Input  secp256k1fx.Input `serialize:"true" json:"input"`
	Output TransferOutput    `serialize:"true" json:"output"`
}

func (op *TransferOperation) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (op *TransferOperation) Cost() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (op *TransferOperation) Outs() []verify.State { _ = "STUB: not implemented"; return nil }

func (op *TransferOperation) Verify() error { _ = "STUB: not implemented"; return nil }
