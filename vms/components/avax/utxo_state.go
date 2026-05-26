// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/linkeddb"
	"github.com/ava-labs/avalanchego/ids"
)

const (
	utxoCacheSize  = 8192
	indexCacheSize = 64
)

var (
	utxoPrefix  = []byte("utxo")
	indexPrefix = []byte("index")
)

// UTXOState is a thin wrapper around a database to provide, caching,
// serialization, and de-serialization for UTXOs.
type UTXOState interface {
	UTXOReader
	UTXOWriter

	// Checksum returns the current UTXOChecksum.
	Checksum() ids.ID
}

// UTXOReader is a thin wrapper around a database to provide fetching of UTXOs.
type UTXOReader interface {
	UTXOGetter

	// UTXOIDs returns the slice of IDs associated with [addr], starting after
	// [previous].
	// If [previous] is not in the list, starts at beginning.
	// Returns at most [limit] IDs.
	UTXOIDs(addr []byte, previous ids.ID, limit int) ([]ids.ID, error)
}

// UTXOGetter is a thin wrapper around a database to provide fetching of a UTXO.
type UTXOGetter interface {
	// GetUTXO attempts to load a utxo.
	GetUTXO(utxoID ids.ID) (*UTXO, error)
}

type UTXOAdder interface {
	AddUTXO(utxo *UTXO)
}

type UTXODeleter interface {
	DeleteUTXO(utxoID ids.ID)
}

// UTXOWriter is a thin wrapper around a database to provide storage and
// deletion of UTXOs.
type UTXOWriter interface {
	// PutUTXO saves the provided utxo to storage.
	PutUTXO(utxo *UTXO) error

	// DeleteUTXO deletes the provided utxo.
	DeleteUTXO(utxoID ids.ID) error
}

type utxoState struct {
	codec codec.Manager

	// UTXO ID -> *UTXO. If the *UTXO is nil the UTXO doesn't exist
	utxoCache cache.Cacher[ids.ID, *UTXO]
	utxoDB    database.Database

	indexDB    database.Database
	indexCache cache.Cacher[string, linkeddb.LinkedDB]

	trackChecksum bool
	checksum      ids.ID
}

func NewUTXOState(
	db database.Database,
	codec codec.Manager,
	trackChecksum bool,
) (UTXOState, error) {
	_ = "STUB: not implemented"
	return *new(UTXOState), nil
}

func NewMeteredUTXOState(
	db database.Database,
	codec codec.Manager,
	metrics prometheus.Registerer,
	trackChecksum bool,
) (UTXOState, error) {
	_ = "STUB: not implemented"
	return *new(UTXOState), nil
}

func (s *utxoState) GetUTXO(utxoID ids.ID) (*UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The key was in the database

func (s *utxoState) PutUTXO(utxo *UTXO) error { _ = "STUB: not implemented"; return nil }

func (s *utxoState) DeleteUTXO(utxoID ids.ID) error { _ = "STUB: not implemented"; return nil }

func (s *utxoState) UTXOIDs(addr []byte, start ids.ID, limit int) ([]ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *utxoState) Checksum() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *utxoState) getIndexDB(addr []byte) linkeddb.LinkedDB {
	_ = "STUB: not implemented"
	return *new(linkeddb.LinkedDB)
}

func (s *utxoState) initChecksum() error { _ = "STUB: not implemented"; return nil }

func (s *utxoState) updateChecksum(modifiedID ids.ID) { _ = "STUB: not implemented"; return }
