// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/types"
)

var (
	ErrInvalidSubnetID = errors.New("invalid subnet ID")
	ErrInvalidWeight   = errors.New("invalid weight")
	ErrInvalidNodeID   = errors.New("invalid node ID")
	ErrInvalidOwner    = errors.New("invalid owner")
)

type PChainOwner struct {
	// The threshold number of `Addresses` that must provide a signature in
	// order for the `PChainOwner` to be considered valid.
	Threshold uint32 `serialize:"true" json:"threshold"`
	// The addresses that are allowed to sign to authenticate a `PChainOwner`.
	Addresses []ids.ShortID `serialize:"true" json:"addresses"`
}

// RegisterL1Validator adds a validator to the subnet.
type RegisterL1Validator struct {
	payload

	SubnetID              ids.ID                 `serialize:"true" json:"subnetID"`
	NodeID                types.JSONByteSlice    `serialize:"true" json:"nodeID"`
	BLSPublicKey          [bls.PublicKeyLen]byte `serialize:"true" json:"blsPublicKey"`
	Expiry                uint64                 `serialize:"true" json:"expiry"`
	RemainingBalanceOwner PChainOwner            `serialize:"true" json:"remainingBalanceOwner"`
	DisableOwner          PChainOwner            `serialize:"true" json:"disableOwner"`
	Weight                uint64                 `serialize:"true" json:"weight"`
}

func (r *RegisterL1Validator) Verify() error { _ = "STUB: not implemented"; return nil }

func (r *RegisterL1Validator) ValidationID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// NewRegisterL1Validator creates a new initialized RegisterL1Validator.
func NewRegisterL1Validator(
	subnetID ids.ID,
	nodeID ids.NodeID,
	blsPublicKey [bls.PublicKeyLen]byte,
	expiry uint64,
	remainingBalanceOwner PChainOwner,
	disableOwner PChainOwner,
	weight uint64,
) (*RegisterL1Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseRegisterL1Validator parses bytes into an initialized
// RegisterL1Validator.
func ParseRegisterL1Validator(b []byte) (*RegisterL1Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
