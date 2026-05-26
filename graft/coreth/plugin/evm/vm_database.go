// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	avalanchedatabase "github.com/ava-labs/avalanchego/database"
)

// initializeDBs initializes the databases used by the VM.
// coreth always uses the avalanchego provided database.
func (vm *VM) initializeDBs(db avalanchedatabase.Database) {
	_ = "STUB: not implemented"
	// Use NewNested rather than New so that the structure of the database
	// remains the same regardless of the provided baseDB type.
	return
}

// Note warpDB is not part of versiondb because it is not necessary
// that warp signatures are committed to the database atomically with
// the last accepted block.

func (vm *VM) inspectDatabases() error { _ = "STUB: not implemented"; return nil }

func inspectDB(db avalanchedatabase.Database, label string) error {
	_ = "STUB: not implemented"
	return nil
}

// Totals

// Inspect key-value database first.

// Display the database statistic.
