// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
)

var (
	_ DelegatorTx     = (*AddPermissionlessDelegatorTx)(nil)
	_ ScheduledStaker = (*AddPermissionlessDelegatorTx)(nil)
)

// AddPermissionlessDelegatorTx is an unsigned addPermissionlessDelegatorTx
type AddPermissionlessDelegatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// Describes the validator
	Validator `serialize:"true" json:"validator"`
	// ID of the subnet this validator is validating
	Subnet ids.ID `serialize:"true" json:"subnetID"`
	// Where to send staked tokens when done validating
	StakeOuts []*avax.TransferableOutput `serialize:"true" json:"stake"`
	// Where to send staking rewards when done validating
	DelegationRewardsOwner fx.Owner `serialize:"true" json:"rewardsOwner"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [AddPermissionlessDelegatorTx]. Also sets the [ctx] to the given [vm.ctx] so
// that the addresses can be json marshalled into human readable format
func (tx *AddPermissionlessDelegatorTx) InitCtx(ctx *snow.Context) {
	_ = "STUB: not implemented"
	return
}

func (tx *AddPermissionlessDelegatorTx) SubnetID() ids.ID {
	_ = "STUB: not implemented"
	return *new(ids.ID)
}

func (tx *AddPermissionlessDelegatorTx) NodeID() ids.NodeID {
	_ = "STUB: not implemented"
	return *new(ids.NodeID)
}

func (*AddPermissionlessDelegatorTx) PublicKey() (*bls.PublicKey, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (tx *AddPermissionlessDelegatorTx) PendingPriority() Priority {
	_ = "STUB: not implemented"
	return *new(Priority)
}

func (tx *AddPermissionlessDelegatorTx) CurrentPriority() Priority {
	_ = "STUB: not implemented"
	return *new(Priority)
}

func (tx *AddPermissionlessDelegatorTx) Stake() []*avax.TransferableOutput {
	_ = "STUB: not implemented"
	return nil
}

func (tx *AddPermissionlessDelegatorTx) RewardsOwner() fx.Owner {
	_ = "STUB: not implemented"
	return *new(fx.Owner)
}

// SyntacticVerify returns nil iff [tx] is valid
func (tx *AddPermissionlessDelegatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

// Ensure there is provided stake

// cache that this is valid

func (tx *AddPermissionlessDelegatorTx) Visit(visitor Visitor) error {
	_ = "STUB: not implemented"
	return nil
}
