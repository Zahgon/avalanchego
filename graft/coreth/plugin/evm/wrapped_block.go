// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap0"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap1"
	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ snowman.Block           = (*wrappedBlock)(nil)
	_ block.WithVerifyContext = (*wrappedBlock)(nil)
	_ extension.ExtendedBlock = (*wrappedBlock)(nil)

	errMissingParentBlock                  = errors.New("missing parent block")
	errInvalidGasUsedRelativeToCapacity    = errors.New("invalid gas used relative to capacity")
	errTotalIntrinsicGasCostExceedsClaimed = errors.New("total intrinsic gas cost is greater than claimed gas used")
)

// Sentinel errors for header validation in this file
var (
	errInvalidExcessBlobGasBeforeCancun    = errors.New("invalid excessBlobGas before cancun")
	errInvalidBlobGasUsedBeforeCancun      = errors.New("invalid blobGasUsed before cancun")
	errInvalidParent                       = errors.New("parent header not found")
	errInvalidParentBeaconRootBeforeCancun = errors.New("invalid parentBeaconRoot before cancun")
	errInvalidExcessBlobGas                = errors.New("invalid excessBlobGas")
	errMissingParentBeaconRoot             = errors.New("header is missing parentBeaconRoot")
	errParentBeaconRootNonEmpty            = errors.New("invalid non-empty parentBeaconRoot")
	errBlobGasUsedNilInCancun              = errors.New("blob gas used must not be nil in Cancun")
	errBlobsNotEnabled                     = errors.New("blobs not enabled on avalanche networks")
)

var (
	ap0MinGasPrice = big.NewInt(ap0.MinGasPrice)
	ap1MinGasPrice = big.NewInt(ap1.MinGasPrice)
)

// wrappedBlock implements the snowman.wrappedBlock interface
type wrappedBlock struct {
	id        ids.ID
	ethBlock  *types.Block
	extension extension.BlockExtension
	vm        *VM
}

// wrapBlock returns a new Block wrapping the ethBlock type and implementing the snowman.Block interface
func wrapBlock(ethBlock *types.Block, vm *VM) (*wrappedBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ID implements the snowman.Block interface
func (b *wrappedBlock) ID() ids.ID {
	_ = "STUB: not implemented"

	// Accept implements the snowman.Block interface
	return *new(ids.ID)
}

func (b *wrappedBlock) Accept(context.Context) error {
	_ = "STUB: not implemented"

	// Although returning an error from Accept is considered fatal, it is good
	// practice to cleanup the batch we were modifying in the case of an error.
	return nil
}

// Call Accept for relevant precompile logs. Note we do this prior to
// calling Accept on the blockChain so any side effects (eg warp signatures)
// take place before the accepted log is emitted to subscribers.

// Get pending operations on the vm's versionDB so we can apply them atomically
// with the block extension's changes.

// Apply any changes atomically with other pending changes to
// the vm's versionDB.
// Accept flushes the changes in the batch to the database.

// If there is no extension, we still need to apply the changes to the versionDB

// handlePrecompileAccept calls Accept on any logs generated with an active precompile address that implements
// contract.Accepter
func (b *wrappedBlock) handlePrecompileAccept(rules extras.Rules) error {
	_ = "STUB: not implemented"
	// Short circuit early if there are no precompile accepters to execute
	return nil
}

// Read receipts from disk

// If there are no receipts, ReadReceipts may be nil, so we check the length and confirm the ReceiptHash
// is empty to ensure that missing receipts results in an error on accept.

// Reject implements the snowman.Block interface
// If [b] contains an atomic transaction, attempt to re-issue it
func (b *wrappedBlock) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }

// Parent implements the snowman.Block interface
func (b *wrappedBlock) Parent() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// Height implements the snowman.Block interface
func (b *wrappedBlock) Height() uint64 { _ = "STUB: not implemented"; return 0 }

// Timestamp implements the snowman.Block interface
func (b *wrappedBlock) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Verify implements the snowman.Block interface
func (b *wrappedBlock) Verify(context.Context) error { _ = "STUB: not implemented"; return nil }

// ShouldVerifyWithContext implements the block.WithVerifyContext interface
func (b *wrappedBlock) ShouldVerifyWithContext(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Short circuit early if there are no predicates to verify

// Check if any of the transactions in the block specify a precompile that enforces a predicate, which requires
// the ProposerVMBlockCtx.

// VerifyWithContext implements the block.WithVerifyContext interface
func (b *wrappedBlock) VerifyWithContext(_ context.Context, proposerVMBlockCtx *block.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the block is valid.
// Enforces that the predicates are valid within [predicateContext].
// Writes the block details to disk and the state to the trie manager iff writes=true.
func (b *wrappedBlock) verify(predicateContext *precompileconfig.PredicateContext, writes bool) error {
	_ = "STUB: not implemented"
	return nil
}

// The engine may call VerifyWithContext multiple times on the same block with different contexts.
// Since the engine will only call Accept/Reject once, we should only call InsertBlockManual once.
// Additionally, if a block is already in processing, then it has already passed verification and
// at this point we have checked the predicates are still valid in the different context so we
// can return nil.

// If this was not called with intention to writing to the database or
// got an error while inserting to blockchain, we may need to cleanup the extension.
// so that the extension can be garbage collected.

func (b *wrappedBlock) verifyIntrinsicGas() error {
	_ = "STUB: not implemented"
	// Verify claimed gas used fits within available capacity for this header.
	// This checks that the gas used is less than the block's capacity.
	return nil
}

// Verify that the claimed GasUsed is within the current capacity.

// Collect all intrinsic gas costs for all transactions in the block.

// Verify that the total intrinsic gas cost is less than or equal to the
// claimed GasUsed.

// semanticVerify verifies that a *Block is internally consistent.
func (b *wrappedBlock) semanticVerify(predicateContext *precompileconfig.PredicateContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure MinDelayExcess is consistent with rules and minimum block delay is enforced.

// Ensure Time and TimeMilliseconds are consistent with rules.

// If the VM is not marked as bootstrapped the other chains may also be
// bootstrapping and not have populated the required indices. Since
// bootstrapping only verifies blocks that have been canonically accepted by
// the network, these checks would be guaranteed to pass on a synced node.

// Verify that all the ICM messages are correctly marked as either valid
// or invalid.

// syntacticVerify verifies that a *Block is well-formed.
func (b *wrappedBlock) syntacticVerify() error { _ = "STUB: not implemented"; return nil }

// Skip verification of the genesis block since it should already be marked as accepted.

// Perform block and header sanity checks

// Verify the extra data is well-formed.

// Check that the tx hash in the header matches the body

// Check that the uncle hash in the header matches the body

// Coinbase must match the BlackholeAddr on C-Chain

// Block must not have any uncles

// Enforce minimum gas prices here prior to dynamic fees going into effect.

// If we are in ApricotPhase0, enforce each transaction has a minimum gas price of at least the LaunchMinGasPrice

// If we are prior to ApricotPhase3, enforce each transaction has a minimum gas price of at least the ApricotPhase1MinGasPrice

// Ensure BaseFee is non-nil as of ApricotPhase3.

// Make sure BlockGasCost is not nil
// NOTE: ethHeader.BlockGasCost correctness is checked in header verification

// Verify the existence / non-existence of excessBlobGas

// verifyPredicates verifies the predicates in the block are valid according to predicateContext.
func (b *wrappedBlock) verifyPredicates(predicateContext *precompileconfig.PredicateContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Bytes implements the snowman.Block interface
func (b *wrappedBlock) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *wrappedBlock) String() string { _ = "STUB: not implemented"; return "" }

func (b *wrappedBlock) GetEthBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (b *wrappedBlock) GetBlockExtension() extension.BlockExtension {
	_ = "STUB: not implemented"
	return *new(extension.BlockExtension)
}
