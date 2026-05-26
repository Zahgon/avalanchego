// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stakeable

import (
	"errors"

	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var (
	errInvalidLocktime      = errors.New("invalid locktime")
	errNestedStakeableLocks = errors.New("shouldn't nest stakeable locks")
)

type LockOut struct {
	Locktime             uint64 `serialize:"true" json:"locktime"`
	avax.TransferableOut `serialize:"true" json:"output"`
}

func (s *LockOut) Addresses() [][]byte { _ = "STUB: not implemented"; return nil }

func (s *LockOut) Verify() error { _ = "STUB: not implemented"; return nil }

type LockIn struct {
	Locktime            uint64 `serialize:"true" json:"locktime"`
	avax.TransferableIn `serialize:"true" json:"input"`
}

func (s *LockIn) Verify() error { _ = "STUB: not implemented"; return nil }
