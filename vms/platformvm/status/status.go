// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package status

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/vms/components/verify"
)

// List of possible status values:
// - [Unknown] The transaction is not known
// - [Committed] The transaction was proposed and committed
// - [Aborted] The transaction was proposed and aborted
// - [Processing] The transaction was proposed and is currently in the preferred chain
// - [Dropped] The transaction was dropped due to failing verification
const (
	Unknown    Status = 0
	Committed  Status = 4
	Aborted    Status = 5
	Processing Status = 6
	Dropped    Status = 8
)

var (
	errUnknownStatus = errors.New("unknown status")

	_ json.Marshaler    = Status(0)
	_ verify.Verifiable = Status(0)
	_ fmt.Stringer      = Status(0)
)

type Status uint32

func (s Status) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Status) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Verify that this is a valid status.
func (s Status) Verify() error { _ = "STUB: not implemented"; return nil }

func (s Status) String() string { _ = "STUB: not implemented"; return "" }
