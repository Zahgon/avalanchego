// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var _ UnsignedTx = (*AdvanceTimeTx)(nil)

// AdvanceTimeTx is a transaction to increase the chain's timestamp.
// When the chain's timestamp is updated (a AdvanceTimeTx is accepted and
// followed by a commit block) the staker set is also updated accordingly.
// It must be that:
// - proposed timestamp > [current chain time]
// - proposed timestamp <= [time for next staker set change]
type AdvanceTimeTx struct {
	// Unix time this block proposes increasing the timestamp to
	Time uint64 `serialize:"true" json:"time"`

	unsignedBytes []byte // Unsigned byte representation of this data
}

func (tx *AdvanceTimeTx) SetBytes(unsignedBytes []byte) { _ = "STUB: not implemented"; return }

func (tx *AdvanceTimeTx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (*AdvanceTimeTx) InitCtx(*snow.Context) {
	_ = "STUB: not implemented"

	// Timestamp returns the time this block is proposing the chain should be set to
	return
}

func (tx *AdvanceTimeTx) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (*AdvanceTimeTx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

func (*AdvanceTimeTx) Outputs() []*avax.TransferableOutput { _ = "STUB: not implemented"; return nil }

func (*AdvanceTimeTx) SyntacticVerify(*snow.Context) error { _ = "STUB: not implemented"; return nil }

func (tx *AdvanceTimeTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
