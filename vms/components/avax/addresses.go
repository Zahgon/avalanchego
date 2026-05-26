// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ AddressManager = (*addressManager)(nil)

	ErrMismatchedChainIDs = errors.New("mismatched chainIDs")
)

type AddressManager interface {
	// ParseLocalAddress takes in an address for this chain and produces the ID
	ParseLocalAddress(addrStr string) (ids.ShortID, error)

	// ParseAddress takes in an address and produces the ID of the chain it's
	// for and the ID of the address
	ParseAddress(addrStr string) (ids.ID, ids.ShortID, error)

	// FormatLocalAddress takes in a raw address and produces the formatted
	// address for this chain
	FormatLocalAddress(addr ids.ShortID) (string, error)

	// FormatAddress takes in a chainID and a raw address and produces the
	// formatted address for that chain
	FormatAddress(chainID ids.ID, addr ids.ShortID) (string, error)
}

type addressManager struct {
	ctx *snow.Context
}

func NewAddressManager(ctx *snow.Context) AddressManager {
	_ = "STUB: not implemented"
	return *new(AddressManager)
}

func (a *addressManager) ParseLocalAddress(addrStr string) (ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ShortID), nil
}

func (a *addressManager) ParseAddress(addrStr string) (ids.ID, ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), *new(ids.ShortID), nil
}

func (a *addressManager) FormatLocalAddress(addr ids.ShortID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *addressManager) FormatAddress(chainID ids.ID, addr ids.ShortID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ParseLocalAddresses(a AddressManager, addrStrs []string) (set.Set[ids.ShortID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseServiceAddress get address ID from address string, being it either localized (using address manager,
// doing also components validations), or not localized.
// If both attempts fail, reports error from localized address parsing
func ParseServiceAddress(a AddressManager, addrStr string) (ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ShortID), nil
}

// ParseServiceAddresses get addresses IDs from addresses strings, being them either localized or not
func ParseServiceAddresses(a AddressManager, addrStrs []string) (set.Set[ids.ShortID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
