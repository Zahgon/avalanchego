// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/vms/types"
)

const (
	// MaxPayloadSize is the maximum size that can be placed into a payload
	MaxPayloadSize = units.KiB
)

var (
	errNilTransferOutput              = errors.New("nil transfer output")
	errPayloadTooLarge                = errors.New("payload too large")
	_                    verify.State = (*TransferOutput)(nil)
)

type TransferOutput struct {
	verify.IsState `json:"-"`

	GroupID                  uint32              `serialize:"true" json:"groupID"`
	Payload                  types.JSONByteSlice `serialize:"true" json:"payload"`
	secp256k1fx.OutputOwners `serialize:"true"`
}

// MarshalJSON marshals Amt and the embedded OutputOwners struct
// into a JSON readable format
// If OutputOwners cannot be serialized then this will return error
func (out *TransferOutput) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (out *TransferOutput) Verify() error { _ = "STUB: not implemented"; return nil }
