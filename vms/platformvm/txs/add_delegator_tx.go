// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
)

var (
	_ DelegatorTx     = (*AddDelegatorTx)(nil)
	_ ScheduledStaker = (*AddDelegatorTx)(nil)

	errDelegatorWeightMismatch = errors.New("delegator weight is not equal to total stake weight")
	errStakeMustBeAVAX         = errors.New("stake must be AVAX")
)

// AddDelegatorTx is an unsigned addDelegatorTx
type AddDelegatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// Describes the delegatee
	Validator `serialize:"true" json:"validator"`
	// Where to send staked tokens when done validating
	StakeOuts []*avax.TransferableOutput `serialize:"true" json:"stake"`
	// Where to send staking rewards when done validating
	DelegationRewardsOwner fx.Owner `serialize:"true" json:"rewardsOwner"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [UnsignedAddDelegatorTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *AddDelegatorTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (*AddDelegatorTx) SubnetID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (tx *AddDelegatorTx) NodeID() ids.NodeID { _ = "STUB: not implemented"; return *new(ids.NodeID) }

func (*AddDelegatorTx) PublicKey() (*bls.PublicKey, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (*AddDelegatorTx) PendingPriority() Priority { _ = "STUB: not implemented"; return *new(Priority) }

func (*AddDelegatorTx) CurrentPriority() Priority { _ = "STUB: not implemented"; return *new(Priority) }

func (tx *AddDelegatorTx) Stake() []*avax.TransferableOutput { _ = "STUB: not implemented"; return nil }

func (tx *AddDelegatorTx) RewardsOwner() fx.Owner { _ = "STUB: not implemented"; return *new(fx.Owner) }

// SyntacticVerify returns nil iff [tx] is valid
func (tx *AddDelegatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

// cache that this is valid

func (tx *AddDelegatorTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
