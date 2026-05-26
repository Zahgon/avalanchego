// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package archivedb

import "github.com/ava-labs/avalanchego/database"

var _ database.Batch = (*batch)(nil)

// batch is how a user performs modifications to the database.
//
// It consumes puts and deletes at a specified height. When committing, an
// atomic operation is created which registers the modifications at the
// specified height and updates the last tracked height to be equal to this
// batch's height.
type batch struct {
	db     *Database
	height uint64
	database.BatchOps
}

func (c *batch) Write() error { _ = "STUB: not implemented"; return nil }

func (c *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }
