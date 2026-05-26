// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package packages

import (
	"github.com/ava-labs/avalanchego/utils/set"
)

// GetDependencies takes a fully qualified package name and returns a map of all
// its recursive package imports (including itself) in the same format.
func GetDependencies(packageName string) (set.Set[string], error) {
	_ = "STUB: not implemented"
	// Configure the load mode to include dependencies
	return nil, nil
}

// collectDeps is recursive

// Avoid re-processing the same dependency

// Start collecting dependencies
