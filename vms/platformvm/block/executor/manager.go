// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/metrics"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/executor"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/mempool"
	"github.com/ava-labs/avalanchego/vms/platformvm/validators"

	snowmanblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ Manager = (*manager)(nil)

	ErrChainNotSynced              = errors.New("chain not synced")
	ErrImportTxWhilePartialSyncing = errors.New("issuing an import tx is not allowed while partial syncing")
)

type Manager interface {
	state.Versions

	// Returns the ID of the most recently accepted block.
	LastAccepted() ids.ID

	SetPreference(blkID ids.ID, blockCtx *snowmanblock.Context)
	Preferred() ids.ID

	GetBlock(blkID ids.ID) (snowman.Block, error)
	GetStatelessBlock(blkID ids.ID) (block.Block, error)
	NewBlock(block.Block) snowman.Block

	// VerifyTx verifies that the transaction can be issued based on the currently
	// preferred state. This should *not* be used to verify transactions in a block.
	VerifyTx(tx *txs.Tx) error

	// VerifyUniqueInputs verifies that the inputs are not duplicated in the
	// provided blk or any of its ancestors pinned in memory.
	VerifyUniqueInputs(blkID ids.ID, inputs set.Set[ids.ID]) error
}

func NewManager(
	mempool *mempool.Mempool,
	metrics metrics.Metrics,
	s *state.State,
	txExecutorBackend *executor.Backend,
	validatorManager *validators.Manager,
) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

type manager struct {
	*backend
	acceptor block.Visitor
	rejector block.Visitor

	preferred         ids.ID
	preferredCtx      *snowmanblock.Context
	txExecutorBackend *executor.Backend
}

func (m *manager) GetBlock(blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (m *manager) GetStatelessBlock(blkID ids.ID) (block.Block, error) {
	_ = "STUB: not implemented"
	return *new(block.Block), nil
}

func (m *manager) NewBlock(blk block.Block) snowman.Block {
	_ = "STUB: not implemented"
	return *new(snowman.Block)
}

func (m *manager) SetPreference(blkID ids.ID, blockCtx *snowmanblock.Context) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) Preferred() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (m *manager) VerifyTx(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

// If partial sync is enabled, this node isn't guaranteed to have the full
// UTXO set from shared memory. To avoid issuing invalid transactions,
// issuance of an ImportTx during this state is completely disallowed.

// TODO: After the mempool is updated, convert this check to use the
// maximum mempool capacity.

func (m *manager) VerifyUniqueInputs(blkID ids.ID, inputs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}
