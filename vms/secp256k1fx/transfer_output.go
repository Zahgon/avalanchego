// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	_ verify.State = (*TransferOutput)(nil)

	ErrNoValueOutput = errors.New("output has no value")
)

type TransferOutput struct {
	verify.IsState `json:"-"`

	Amt uint64 `serialize:"true" json:"amount"`

	OutputOwners `serialize:"true"`
}

// MarshalJSON marshals Amt and the embedded OutputOwners struct
// into a JSON readable format
// If OutputOwners cannot be serialized then this will return error
func (out *TransferOutput) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Amount returns the quantity of the asset this output consumes
func (out *TransferOutput) Amount() uint64 { _ = "STUB: not implemented"; return 0 }

func (out *TransferOutput) Verify() error { _ = "STUB: not implemented"; return nil }

func (out *TransferOutput) Owners() interface{} { _ = "STUB: not implemented"; return nil }
