// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlist

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

var (
	ErrAdminAndEnabledAddress         = errors.New("cannot set address as both admin and enabled")
	ErrAdminAndManagerAddress         = errors.New("cannot set address as both admin and manager")
	ErrCannotAddManagersBeforeDurango = errors.New("cannot add managers before Durango")
	ErrDuplicateEnabledAddress        = errors.New("duplicate address in enabled list")
	ErrDuplicateAdminAddress          = errors.New("duplicate address in admin list")
	ErrDuplicateManagerAddress        = errors.New("duplicate address in manager list")
	ErrEnabledAndManagerAddress       = errors.New("cannot set address as both enabled and manager")
)

// AllowListConfig specifies the initial set of addresses with Admin or Enabled roles.
type AllowListConfig struct {
	AdminAddresses   []common.Address `json:"adminAddresses,omitempty"`   // initial admin addresses
	ManagerAddresses []common.Address `json:"managerAddresses,omitempty"` // initial manager addresses
	EnabledAddresses []common.Address `json:"enabledAddresses,omitempty"` // initial enabled addresses
}

// Configure initializes the address space of [precompileAddr] by initializing the role of each of
// the addresses in [AllowListAdmins].
func (c *AllowListConfig) Configure(_ precompileconfig.ChainConfig, precompileAddr common.Address, state contract.StateDB, _ contract.ConfigurationBlockContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify() should have been called before Configure()
// so we know manager role is activated

// Equal returns true iff [other] has the same admins in the same order in its allow list.
func (c *AllowListConfig) Equal(other *AllowListConfig) bool {
	_ = "STUB: not implemented"
	return false
}

// areEqualAddressLists returns true iff [a] and [b] have the same addresses in the same order.
func areEqualAddressLists(current []common.Address, other []common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify returns an error if there is an overlapping address between admin and enabled roles
func (c *AllowListConfig) Verify(chainConfig precompileconfig.ChainConfig, upgrade precompileconfig.Upgrade) error {
	_ = "STUB: not implemented"
	return nil
}

// tracks which addresses we have seen and their role

// check for duplicates in enabled list

// check for overlap between enabled and admin lists or duplicates in admin list

// If the config attempts to activate a manager before the Durango, fail verification

// check for overlap between admin and manager lists or duplicates in manager list
