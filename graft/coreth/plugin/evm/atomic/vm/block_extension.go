// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
)

var (
	_ extension.BlockExtension  = (*blockExtension)(nil)
	_ extension.BlockExtender   = (*blockExtender)(nil)
	_ atomic.AtomicBlockContext = (*blockExtension)(nil)
)

var (
	errNilEthBlock            = errors.New("nil ethBlock")
	ErrMissingUTXOs           = errors.New("missing UTXOs")
	ErrEmptyBlock             = errors.New("empty block")
	errAtomicExtractionFailed = errors.New("atomic tx extraction failed")
)

type blockExtender struct {
	extDataHashes map[common.Hash]common.Hash
	vm            *VM
}

type blockExtension struct {
	atomicTxs     []*atomic.Tx
	blockExtender *blockExtender
	block         extension.ExtendedBlock
}

// newBlockExtender returns a new block extender.
func newBlockExtender(
	extDataHashes map[common.Hash]common.Hash,
	vm *VM,
) *blockExtender {
	_ = "STUB: not implemented"
	return nil
}

// Note: we need VM here to access the atomic backend that
// could be initialized later in the VM.

// NewBlockExtension returns a new block extension.
func (be *blockExtender) NewBlockExtension(b extension.ExtendedBlock) (extension.BlockExtension, error) {
	_ = "STUB: not implemented"
	return *new(extension.BlockExtension), nil
}

// Extract atomic transactions from the block

// SyntacticVerify checks the syntactic validity of the block. This is called by the wrapper
// block manager's SyntacticVerify method.
func (be *blockExtension) SyntacticVerify(rules extras.Rules) error {
	_ = "STUB: not implemented"
	return nil
}

// should not happen

// If there is no extra data, check that there is no extra data in the hash map either to ensure we do not
// have a block that is unexpectedly missing extra data.

// If there is extra data, check to make sure that the extra data hash matches the expected extra data hash for this
// block

// Verify the ExtDataHash field

// Block must not be empty

// If we are in ApricotPhase4, ensure that ExtDataGasUsed is populated correctly.

// After the F upgrade, the extDataGasUsed field is validated by
// [header.VerifyGasUsed].

// We perform this check manually here to avoid the overhead of having to
// reparse the atomicTx in `CalcExtDataGasUsed`.
// Charge the atomic tx fixed fee as of ApricotPhase5

// SemanticVerify checks the semantic validity of the block. This is called by the wrapper
// block manager's SemanticVerify method.
func (be *blockExtension) SemanticVerify() error { _ = "STUB: not implemented"; return nil }

// Verify that the UTXOs named in import txs are present in shared
// memory.
//
// This does not fully verify that this block can spend these UTXOs.
// However, it guarantees that any block that fails the later checks was
// built by an incorrect block proposer. This ensures that we only mark
// blocks as BAD BLOCKs if they were incorrectly generated.

// Accept is called when the block is accepted. This is called by the wrapper
// block manager's Accept method. The acceptedBatch contains the changes that
// were made to the database as a result of accepting the block, and it's flushed
// to the database in this method.
func (be *blockExtension) Accept(acceptedBatch database.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove the accepted transaction from the mempool

// Update VM state for atomic txs in this block. This includes updating the
// atomic tx repo, atomic trie, and shared memory.

// should never occur since [b] must be verified before calling Accept

// Apply any shared memory changes atomically with other pending batched changes

// Reject is called when the block is rejected. This is called by the wrapper
// block manager's Reject method.
func (be *blockExtension) Reject() error { _ = "STUB: not implemented"; return nil }

// Re-issue the transaction in the mempool, continue even if it fails

// should never occur since [b] must be verified before calling Reject

// CleanupVerified is called when the block is cleaned up after a failed insertion.
func (be *blockExtension) CleanupVerified() { _ = "STUB: not implemented"; return }

// If the state isn't found, no need to reject (it was never verified).

// atomicState.Reject() never returns an error in practice.
//nolint:errcheck // Reject never returns an error

// AtomicTxs returns the atomic transactions in this block.
func (be *blockExtension) AtomicTxs() []*atomic.Tx { _ = "STUB: not implemented"; return nil }

// verifyUTXOsPresent verifies all atomic UTXOs consumed by the block are
// present in shared memory.
func (be *blockExtension) verifyUTXOsPresent(atomicTxs []*atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// verify UTXOs named in import txs are present in shared memory.
