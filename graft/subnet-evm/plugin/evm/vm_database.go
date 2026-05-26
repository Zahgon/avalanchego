// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm/config"
	"github.com/ava-labs/avalanchego/utils/logging"

	avalanchedatabase "github.com/ava-labs/avalanchego/database"
)

const (
	dbMetricsPrefix = "db"
	meterDBGatherer = "meterdb"
)

type DatabaseConfig struct {
	// If true, all writes are to memory and are discarded at shutdown.
	ReadOnly bool `json:"readOnly"`

	// Path to database
	Path string `json:"path"`

	// Name of the database type to use
	Name string `json:"name"`

	// Config bytes (JSON) for the database
	// See relevant (pebbledb, leveldb) config options
	Config []byte `json:"-"`
}

// initializeDBs initializes the databases used by the VM.
// If [useStandaloneDB] is true, the chain will use a standalone database for its state.
// Otherwise, the chain will use the provided [avaDB] for its state.
func (vm *VM) initializeDBs(avaDB avalanchedatabase.Database) error {
	_ = "STUB: not implemented"

	// skip standalone database initialization if we are running in unit tests
	return nil
}

// first initialize the accepted block database to check if we need to use a standalone database

// If we are using a standalone database, we need to create a new database
// for the chain state.

// Use NewNested rather than New so that the structure of the database
// remains the same regardless of the provided baseDB type.

// Note warpDB and validatorsDB are not part of versiondb because it is not necessary
// that they are committed to the database atomically with
// the last accepted block.
// [warpDB] is used to store warp message signatures
// set to a prefixDB with the prefix [warpPrefix]

// [validatorsDB] is used to store the current validator set and uptimes
// set to a prefixDB with the prefix [validatorsDBPrefix]

func (vm *VM) inspectDatabases() error { _ = "STUB: not implemented"; return nil }

// useStandaloneDatabase returns true if the chain can and should use a standalone database
// other than given by [db] in Initialize()
func (vm *VM) useStandaloneDatabase(acceptedDB avalanchedatabase.Database) (bool, error) {
	_ = "STUB: not implemented"
	// no config provided, use default
	return false, nil
}

// check if the chain can use a standalone database

// If there is nothing in the database, we can use the standalone database

// getDatabaseConfig returns the database configuration for the chain
// to be used by separate, standalone database.
func getDatabaseConfig(config config.Config, chainDataDir string) (DatabaseConfig, error) {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig), nil
}

func inspectDB(db avalanchedatabase.Database, label string) error {
	_ = "STUB: not implemented"
	return nil
}

// Totals

// Inspect key-value database first.

// Display the database statistic.

func newStandaloneDatabase(dbConfig DatabaseConfig, gatherer metrics.MultiGatherer, logger logging.Logger) (avalanchedatabase.Database, error) {
	_ = "STUB: not implemented"
	return *new(avalanchedatabase.Database), nil
}

// If the database is pebble, we need to set the config
// to use no sync. Sync mode in pebble has an issue with OSs like MacOS.

// Default to "no sync" for pebble db

// Marshal the config back to bytes to ensure that new defaults are applied
