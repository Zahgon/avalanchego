// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/avm/block"
)

const SyncBound = 10 * time.Second

var (
	_ snowman.Block = (*Block)(nil)

	ErrUnexpectedMerkleRoot        = errors.New("unexpected merkle root")
	ErrTimestampBeyondSyncBound    = errors.New("proposed timestamp is too far in the future relative to local time")
	ErrEmptyBlock                  = errors.New("block contains no transactions")
	ErrChildBlockEarlierThanParent = errors.New("proposed timestamp before current chain time")
	ErrConflictingBlockTxs         = errors.New("block contains conflicting transactions")
	ErrIncorrectHeight             = errors.New("block has incorrect height")
	ErrBlockNotFound               = errors.New("block not found")
)

// Exported for testing in avm package.
type Block struct {
	block.Block
	manager *manager
}

func (b *Block) Verify(context.Context) error { _ = "STUB: not implemented"; return nil }

// This block has already been verified.

// Currently we don't populate the blocks merkle root.

// Only allow timestamp to reasonably far forward

// Syntactic verification is generally pretty fast, so we verify this first
// before performing any possible DB reads.

// Verify that the parent exists.

// Verify that currentBlkHeight = parentBlkHeight + 1.

// The proposed timestamp must not be before the parent's timestamp.

// Verify that the tx is valid according to the current state of the
// chain.

// Apply the txs state changes to the state.
//
// Note: This must be done inside the same loop as semantic verification
// to ensure that semantic verification correctly accounts for
// transactions that occurred earlier in the block.

// Verify that the transaction we just executed didn't consume inputs
// that were already imported in a previous transaction.

// Now that the tx would be marked as accepted, we should add it to the
// state for the next transaction in the block.

// Add/merge in the atomic requests represented by [tx]

// Verify that none of the transactions consumed any inputs that were
// already imported in a currently processing block.

// Now that the block has been executed, we can add the block data to the
// state diff.

func (b *Block) Accept(context.Context) error { _ = "STUB: not implemented"; return nil }

// Update the state to reflect the changes made in [onAcceptState].

// Note that this method writes [batch] to the database.

func (b *Block) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }
