// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	ErrNilOutput            = errors.New("nil output")
	ErrOutputUnspendable    = errors.New("output is unspendable")
	ErrOutputUnoptimized    = errors.New("output representation should be optimized")
	ErrAddrsNotSortedUnique = errors.New("addresses not sorted and unique")
)

type OutputOwners struct {
	verify.IsNotState `json:"-"`

	Locktime  uint64        `serialize:"true" json:"locktime"`
	Threshold uint32        `serialize:"true" json:"threshold"`
	Addrs     []ids.ShortID `serialize:"true" json:"addresses"`

	// ctx is used in MarshalJSON to convert Addrs into human readable
	// format with ChainID and NetworkID. Unexported because we don't use
	// it outside this object.
	ctx *snow.Context
}

// InitCtx allows addresses to be formatted into their human readable format
// during json marshalling.
func (out *OutputOwners) InitCtx(ctx *snow.Context) {
	_ = "STUB: not implemented"

	// MarshalJSON marshals OutputOwners as JSON with human readable addresses.
	// OutputOwners.InitCtx must be called before marshalling this or one of
	// the parent objects to json. Uses the OutputOwners.ctx method to format
	// the addresses. Returns errMarshal error if OutputOwners.ctx is not set.
	return
}

func (out *OutputOwners) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Fields returns JSON keys in a map that can be used with marshal JSON
// to serialize OutputOwners struct
func (out *OutputOwners) Fields() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for each [addr] in [Addrs] we attempt to format it given
// the [out.ctx] object

// we expect these addresses to be valid, return error
// if they are not

// Addresses returns the addresses that manage this output
func (out *OutputOwners) Addresses() [][]byte { _ = "STUB: not implemented"; return nil }

// AddressesSet returns addresses as a set
func (out *OutputOwners) AddressesSet() set.Set[ids.ShortID] { _ = "STUB: not implemented"; return nil }

// Equals returns true if the provided owners create the same condition
func (out *OutputOwners) Equals(other *OutputOwners) bool { _ = "STUB: not implemented"; return false }

func (out *OutputOwners) Verify() error { _ = "STUB: not implemented"; return nil }

func (out *OutputOwners) Sort() { _ = "STUB: not implemented"; return }

// formatAddress formats a given [addr] into human readable format using
// [ChainID] and [NetworkID] if a non-nil [ctx] is provided. If [ctx] is not
// provided, the address will be returned in cb58 format.
func formatAddress(ctx *snow.Context, addr ids.ShortID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
