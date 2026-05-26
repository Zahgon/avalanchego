// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/block"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

const (
	txCacheSize      = 8192
	blockIDCacheSize = 8192
	blockCacheSize   = 2048
)

var (
	utxoPrefix      = []byte("utxo")
	txPrefix        = []byte("tx")
	blockIDPrefix   = []byte("blockID")
	blockPrefix     = []byte("block")
	singletonPrefix = []byte("singleton")

	isInitializedKey = []byte{0x00}
	timestampKey     = []byte{0x01}
	lastAcceptedKey  = []byte{0x02}

	_ State = (*state)(nil)
)

type ReadOnlyChain interface {
	avax.UTXOGetter

	GetTx(txID ids.ID) (*txs.Tx, error)
	GetBlockIDAtHeight(height uint64) (ids.ID, error)
	GetBlock(blkID ids.ID) (block.Block, error)
	GetLastAccepted() ids.ID
	GetTimestamp() time.Time
}

type Chain interface {
	ReadOnlyChain
	avax.UTXOAdder
	avax.UTXODeleter

	AddTx(tx *txs.Tx)
	AddBlock(block block.Block)
	SetLastAccepted(blkID ids.ID)
	SetTimestamp(t time.Time)
}

// State persistently maintains a set of UTXOs, transaction, statuses, and
// singletons.
type State interface {
	Chain
	avax.UTXOReader

	IsInitialized() (bool, error)
	SetInitialized() error

	// InitializeChainState is called after the VM has been linearized. Calling
	// [GetLastAccepted] or [GetTimestamp] before calling this function will
	// return uninitialized data.
	//
	// Invariant: After the chain is linearized, this function is expected to be
	// called during startup.
	InitializeChainState(stopVertexID ids.ID, genesisTimestamp time.Time) error

	// Discard uncommitted changes to the database.
	Abort()

	// Commit changes to the base database.
	Commit() error

	// Returns a batch of unwritten changes that, when written, will commit all
	// pending changes to the base database.
	CommitBatch() (database.Batch, error)

	// Checksum returns the current state checksum.
	Checksum() ids.ID

	Close() error
}

/*
 * VMDB
 * |- utxos
 * | '-- utxoDB
 * |-. txs
 * | '-- txID -> tx bytes
 * |-. blockIDs
 * | '-- height -> blockID
 * |-. blocks
 * | '-- blockID -> block bytes
 * '-. singletons
 *   |-- initializedKey -> nil
 *   |-- timestampKey -> timestamp
 *   '-- lastAcceptedKey -> lastAccepted
 */
type state struct {
	parser block.Parser
	db     *versiondb.Database

	modifiedUTXOs map[ids.ID]*avax.UTXO // map of modified UTXOID -> *UTXO if the UTXO is nil, it has been removed
	utxoDB        database.Database
	utxoState     avax.UTXOState

	addedTxs map[ids.ID]*txs.Tx            // map of txID -> *txs.Tx
	txCache  cache.Cacher[ids.ID, *txs.Tx] // cache of txID -> *txs.Tx. If the entry is nil, it is not in the database
	txDB     database.Database

	addedBlockIDs map[uint64]ids.ID            // map of height -> blockID
	blockIDCache  cache.Cacher[uint64, ids.ID] // cache of height -> blockID. If the entry is ids.Empty, it is not in the database
	blockIDDB     database.Database

	addedBlocks map[ids.ID]block.Block            // map of blockID -> Block
	blockCache  cache.Cacher[ids.ID, block.Block] // cache of blockID -> Block. If the entry is nil, it is not in the database
	blockDB     database.Database

	// [lastAccepted] is the most recently accepted block.
	lastAccepted, persistedLastAccepted ids.ID
	timestamp, persistedTimestamp       time.Time
	singletonDB                         database.Database
}

func New(
	db *versiondb.Database,
	parser block.Parser,
	metrics prometheus.Registerer,
	trackChecksums bool,
) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (s *state) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *state) UTXOIDs(addr []byte, start ids.ID, limit int) ([]ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *state) AddUTXO(utxo *avax.UTXO) { _ = "STUB: not implemented"; return }

func (s *state) DeleteUTXO(utxoID ids.ID) { _ = "STUB: not implemented"; return }

func (s *state) GetTx(txID ids.ID) (*txs.Tx, error) { _ = "STUB: not implemented"; return nil, nil }

// The key was in the database

func (s *state) AddTx(tx *txs.Tx) { _ = "STUB: not implemented"; return }

func (s *state) GetBlockIDAtHeight(height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (s *state) GetBlock(blkID ids.ID) (block.Block, error) {
	_ = "STUB: not implemented"
	return *new(block.Block), nil
}

func (s *state) AddBlock(block block.Block) { _ = "STUB: not implemented"; return }

func (s *state) InitializeChainState(stopVertexID ids.ID, genesisTimestamp time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) initializeChainState(stopVertexID ids.ID, genesisTimestamp time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) IsInitialized() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *state) SetInitialized() error { _ = "STUB: not implemented"; return nil }

func (s *state) GetLastAccepted() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *state) SetLastAccepted(lastAccepted ids.ID) { _ = "STUB: not implemented"; return }

func (s *state) GetTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *state) SetTimestamp(t time.Time) { _ = "STUB: not implemented"; return }

func (s *state) Commit() error { _ = "STUB: not implemented"; return nil }

func (s *state) Abort() { _ = "STUB: not implemented"; return }

func (s *state) CommitBatch() (database.Batch, error) {
	_ = "STUB: not implemented"
	return *new(database.Batch), nil
}

func (s *state) Close() error { _ = "STUB: not implemented"; return nil }

func (s *state) write() error { _ = "STUB: not implemented"; return nil }

func (s *state) writeUTXOs() error { _ = "STUB: not implemented"; return nil }

func (s *state) writeTxs() error { _ = "STUB: not implemented"; return nil }

func (s *state) writeBlockIDs() error { _ = "STUB: not implemented"; return nil }

func (s *state) writeBlocks() error { _ = "STUB: not implemented"; return nil }

func (s *state) writeMetadata() error { _ = "STUB: not implemented"; return nil }

func (s *state) Checksum() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }
