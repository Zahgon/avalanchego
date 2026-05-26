// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import "github.com/ava-labs/avalanchego/database"

var _ database.Batch = (*batch)(nil)

type batch struct {
	database.BatchOps

	db *merkleDB
}

// Assumes [b.db.lock] isn't held.
func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

func (b *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }
