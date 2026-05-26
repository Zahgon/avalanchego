// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"errors"
	"math"

	"github.com/ava-labs/avalanchego/utils/json"
)

type Height json.Uint64

const (
	ProposedHeightJSON = `"proposed"`
	ProposedHeight     = math.MaxUint64
)

var errInvalidHeight = errors.New("invalid height")

func (h Height) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *Height) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	// First check for known string values
	return nil
}

// Otherwise, unmarshal as a uint64

// MaxUint64 is reserved for proposed height, so return an error if supplied
// numerically.

func (h Height) IsProposed() bool { _ = "STUB: not implemented"; return false }
