// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package factory

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils/logging"
)

// New creates a new database instance based on the provided configuration.
//
// It also wraps the database with a corruptable DB.
//
// dbName is the name of the database, either leveldb, memdb, or pebbledb.
// dbPath is the path to the database folder.
// readOnly indicates if the database should be read-only.
// dbConfig is the database configuration in JSON format.
func New(
	name string,
	path string,
	readOnly bool,
	config []byte,
	reg prometheus.Registerer,
	logger logging.Logger,
) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}
