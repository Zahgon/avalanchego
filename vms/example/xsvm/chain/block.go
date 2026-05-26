// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/utils/set"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	xsblock "github.com/ava-labs/avalanchego/vms/example/xsvm/block"
)

const maxClockSkew = 10 * time.Second

var (
	_ Block = (*block)(nil)

	errMissingParent         = errors.New("missing parent block")
	errMissingChild          = errors.New("missing child block")
	errParentNotVerified     = errors.New("parent block has not been verified")
	errFutureTimestamp       = errors.New("future timestamp")
	errTimestampBeforeParent = errors.New("timestamp before parent")
	errWrongHeight           = errors.New("wrong height")
)

type Block interface {
	snowman.Block
	smblock.WithVerifyContext

	// State intends to return the new chain state following this block's
	// acceptance. The new chain state is built (but not persisted) following a
	// block's verification to allow block's descendants verification before
	// being accepted.
	State() (database.Database, error)
}

type block struct {
	*xsblock.Stateless

	chain *chain

	id    ids.ID
	bytes []byte

	state               *versiondb.Database
	verifiedChildrenIDs set.Set[ids.ID]
}

func (b *block) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *block) Parent() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *block) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *block) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *block) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *block) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *block) Accept(context.Context) error { _ = "STUB: not implemented"; return nil }

// Following this block's acceptance, make sure that it's direct children
// point to the base state, which now also contains this block's changes.

func (b *block) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO: push transactions back into the mempool

func (b *block) ShouldVerifyWithContext(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *block) VerifyWithContext(ctx context.Context, blockContext *smblock.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// parent block must be verified or accepted

// This block's state is a versionDB built on top of it's parent state. This
// block's changes are pushed atomically to the parent state when accepted.

// Make sure to only state the state the first time we verify this block.

func (b *block) State() (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

// If this block isn't processing, then the child should never have had
// verify called on it.
