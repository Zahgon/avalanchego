// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customrawdb

import (
	"errors"

	"github.com/ava-labs/libevm/ethdb"
)

// FirewoodScheme is the scheme for the Firewood storage scheme.
const FirewoodScheme = "firewood"

// errStateSchemeConflict indicates the provided state scheme conflicts with
// what is on disk.
var errStateSchemeConflict = errors.New("state scheme conflict")

// ParseStateScheme parses the state scheme from the provided string.
func ParseStateScheme(provided string, db ethdb.Database) (string, error) {
	_ = "STUB: not implemented"
	// Check for custom scheme
	return "", nil
}

// Valid scheme on db mismatched

// If no conflicting scheme is found, is valid.

// Check for valid eth scheme
