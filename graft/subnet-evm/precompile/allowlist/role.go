// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlist

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
)

// 1. NoRole - this is equivalent to common.Hash{} and deletes the key from the DB when set
// 2. EnabledRole - allowed to call the precompile
// 3. Admin - allowed to both modify the allowlist and call the precompile
// 4. Manager - allowed to add and remove only enabled addresses and also call the precompile. (only after Durango)
var (
	NoRole      = Role(common.BigToHash(common.Big0))
	EnabledRole = Role(common.BigToHash(common.Big1))
	AdminRole   = Role(common.BigToHash(common.Big2))
	ManagerRole = Role(common.BigToHash(common.Big3))
	// Roles should be incremented and not changed.

	ErrInvalidRole = errors.New("invalid role")
)

// Enum constants for valid Role
type Role common.Hash

// IsNoRole returns true if [r] indicates no specific role.
func (r Role) IsNoRole() bool { _ = "STUB: not implemented"; return false }

// IsAdmin returns true if [r] indicates the permission to modify the allow list.
func (r Role) IsAdmin() bool { _ = "STUB: not implemented"; return false }

// IsEnabled returns true if [r] indicates that it has permission to access the resource.
func (r Role) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (r Role) CanModify(from, target Role) bool { _ = "STUB: not implemented"; return false }

func (r Role) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (r Role) Big() *big.Int { _ = "STUB: not implemented"; return nil }

func (r Role) Hash() common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

func (r Role) GetSetterFunctionName() (string, error) { _ = "STUB: not implemented"; return "", nil }

// String returns a string representation of [r].
func (r Role) String() string { _ = "STUB: not implemented"; return "" }

func FromBig(b *big.Int) (Role, error) { _ = "STUB: not implemented"; return *new(Role), nil }
