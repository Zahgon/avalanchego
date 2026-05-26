// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package choices

import (
	"errors"
)

var errUnknownStatus = errors.New("unknown status")

type Status uint32

// List of possible status values
// [Unknown] Zero value, means the operation is not known
// [Processing] means the operation is known, but hasn't been decided yet
// [Rejected] means the operation will never be accepted
// [Accepted] means the operation was accepted
const (
	Unknown Status = iota
	Processing
	Rejected
	Accepted
)

func (s Status) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Status) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Fetched returns true if the status has been set.
func (s Status) Fetched() bool { _ = "STUB: not implemented"; return false }

// Decided returns true if the status is Rejected or Accepted.
func (s Status) Decided() bool { _ = "STUB: not implemented"; return false }

// Valid returns nil if the status is a valid status.
func (s Status) Valid() error { _ = "STUB: not implemented"; return nil }

func (s Status) String() string { _ = "STUB: not implemented"; return "" }

// Bytes returns the byte repr. of this status
func (s Status) Bytes() []byte { _ = "STUB: not implemented"; return nil }
