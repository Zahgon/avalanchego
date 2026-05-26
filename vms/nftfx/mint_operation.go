// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"errors"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/vms/types"
)

var errNilMintOperation = errors.New("nil mint operation")

type MintOperation struct {
	MintInput secp256k1fx.Input           `serialize:"true" json:"mintInput"`
	GroupID   uint32                      `serialize:"true" json:"groupID"`
	Payload   types.JSONByteSlice         `serialize:"true" json:"payload"`
	Outputs   []*secp256k1fx.OutputOwners `serialize:"true" json:"outputs"`
}

func (op *MintOperation) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (op *MintOperation) Cost() (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Outs Returns []TransferOutput as []verify.State
		nil
}

func (op *MintOperation) Outs() []verify.State { _ = "STUB: not implemented"; return nil }

func (op *MintOperation) Verify() error { _ = "STUB: not implemented"; return nil }
