// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/block"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var (
	_ Diff     = (*diff)(nil)
	_ Versions = stateGetter{}

	ErrMissingParentState = errors.New("missing parent state")
)

type Diff interface {
	Chain

	Apply(Chain)
}

type diff struct {
	parentID      ids.ID
	stateVersions Versions

	// map of modified UTXOID -> *UTXO if the UTXO is nil, it has been removed
	modifiedUTXOs map[ids.ID]*avax.UTXO
	addedTxs      map[ids.ID]*txs.Tx     // map of txID -> tx
	addedBlockIDs map[uint64]ids.ID      // map of height -> blockID
	addedBlocks   map[ids.ID]block.Block // map of blockID -> block

	lastAccepted ids.ID
	timestamp    time.Time
}

func NewDiff(
	parentID ids.ID,
	stateVersions Versions,
) (Diff, error) {
	_ = "STUB: not implemented"
	return *new(Diff), nil
}

type stateGetter struct {
	state Chain
}

func (s stateGetter) GetState(ids.ID) (Chain, bool) {
	_ = "STUB: not implemented"
	return *new(Chain), false
}

func NewDiffOn(parentState Chain) (Diff, error) { _ = "STUB: not implemented"; return *new(Diff), nil }

func (d *diff) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *diff) AddUTXO(utxo *avax.UTXO) { _ = "STUB: not implemented"; return }

func (d *diff) DeleteUTXO(utxoID ids.ID) { _ = "STUB: not implemented"; return }

func (d *diff) GetTx(txID ids.ID) (*txs.Tx, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *diff) AddTx(tx *txs.Tx) { _ = "STUB: not implemented"; return }

func (d *diff) GetBlockIDAtHeight(height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (d *diff) GetBlock(blkID ids.ID) (block.Block, error) {
	_ = "STUB: not implemented"
	return *new(block.Block), nil
}

func (d *diff) AddBlock(blk block.Block) { _ = "STUB: not implemented"; return }

func (d *diff) GetLastAccepted() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (d *diff) SetLastAccepted(lastAccepted ids.ID) { _ = "STUB: not implemented"; return }

func (d *diff) GetTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (d *diff) SetTimestamp(t time.Time) { _ = "STUB: not implemented"; return }

func (d *diff) Apply(state Chain) { _ = "STUB: not implemented"; return }
