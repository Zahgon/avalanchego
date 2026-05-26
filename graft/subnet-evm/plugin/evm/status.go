// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
)

var errUnknownStatus = errors.New("unknown status")

// Status ...
type Status uint32

// List of possible status values
// [Unknown] Zero value, means the status is not known
// [Dropped] means the transaction was in the mempool, but was dropped because it failed verification
// [Processing] means the transaction is in the mempool
// [Accepted] means the transaction was accepted
const (
	Unknown Status = iota
	Dropped
	Processing
	Accepted
)

// MarshalJSON ...
func (s Status) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON ...
func (s *Status) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Valid returns nil if the status is a valid status.
func (s Status) Valid() error { _ = "STUB: not implemented"; return nil }

func (s Status) String() string { _ = "STUB: not implemented"; return "" }
