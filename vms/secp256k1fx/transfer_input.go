// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/avalanchego/snow"
)

var ErrNoValueInput = errors.New("input has no value")

type TransferInput struct {
	Amt   uint64 `serialize:"true" json:"amount"`
	Input `serialize:"true"`
}

func (*TransferInput) InitCtx(*snow.Context) {
	_ = "STUB: not implemented"

	// Amount returns the quantity of the asset this input produces
	return
}

func (in *TransferInput) Amount() uint64 {
	_ = "STUB: not implemented"

	// Verify this input is syntactically valid
	return 0
}

func (in *TransferInput) Verify() error { _ = "STUB: not implemented"; return nil }
