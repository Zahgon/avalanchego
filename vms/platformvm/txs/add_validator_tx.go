// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
)

var (
	_ ValidatorTx     = (*AddValidatorTx)(nil)
	_ ScheduledStaker = (*AddValidatorTx)(nil)

	errTooManyShares = fmt.Errorf("a staker can only require at most %d shares from delegators", reward.PercentDenominator)
)

// AddValidatorTx is an unsigned addValidatorTx
type AddValidatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// Describes the delegatee
	Validator `serialize:"true" json:"validator"`
	// Where to send staked tokens when done validating
	StakeOuts []*avax.TransferableOutput `serialize:"true" json:"stake"`
	// Where to send staking rewards when done validating
	RewardsOwner fx.Owner `serialize:"true" json:"rewardsOwner"`
	// Fee this validator charges delegators as a percentage, times 10,000
	// For example, if this validator has DelegationShares=300,000 then they
	// take 30% of rewards from delegators
	DelegationShares uint32 `serialize:"true" json:"shares"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [AddValidatorTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *AddValidatorTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (*AddValidatorTx) SubnetID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (tx *AddValidatorTx) NodeID() ids.NodeID { _ = "STUB: not implemented"; return *new(ids.NodeID) }

func (*AddValidatorTx) PublicKey() (*bls.PublicKey, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (*AddValidatorTx) PendingPriority() Priority { _ = "STUB: not implemented"; return *new(Priority) }

func (*AddValidatorTx) CurrentPriority() Priority { _ = "STUB: not implemented"; return *new(Priority) }

func (tx *AddValidatorTx) Stake() []*avax.TransferableOutput { _ = "STUB: not implemented"; return nil }

func (tx *AddValidatorTx) ValidationRewardsOwner() fx.Owner {
	_ = "STUB: not implemented"
	return *new(fx.Owner)
}

func (tx *AddValidatorTx) DelegationRewardsOwner() fx.Owner {
	_ = "STUB: not implemented"
	return *new(fx.Owner)
}

func (tx *AddValidatorTx) Shares() uint32 { _ = "STUB: not implemented"; return 0 }

// SyntacticVerify returns nil iff [tx] is valid
func (tx *AddValidatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

// Ensure delegators shares are in the allowed amount

// cache that this is valid

func (tx *AddValidatorTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
