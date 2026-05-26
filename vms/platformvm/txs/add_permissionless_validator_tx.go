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
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
)

var (
	_ ValidatorTx     = (*AddPermissionlessValidatorTx)(nil)
	_ ScheduledStaker = (*AddPermissionlessDelegatorTx)(nil)

	errEmptyNodeID             = errors.New("validator nodeID cannot be empty")
	errNoStake                 = errors.New("no stake")
	errInvalidSigner           = errors.New("invalid signer")
	errMultipleStakedAssets    = errors.New("multiple staked assets")
	errValidatorWeightMismatch = errors.New("validator weight mismatch")
)

// AddPermissionlessValidatorTx is an unsigned addPermissionlessValidatorTx
type AddPermissionlessValidatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// Describes the validator
	Validator `serialize:"true" json:"validator"`
	// ID of the subnet this validator is validating
	Subnet ids.ID `serialize:"true" json:"subnetID"`
	// If the [Subnet] is the primary network, [Signer] is the BLS key for this
	// validator. If the [Subnet] is not the primary network, this value is the
	// empty signer
	// Note: We do not enforce that the BLS key is unique across all validators.
	//       This means that validators can share a key if they so choose.
	//       However, a NodeID does uniquely map to a BLS key
	Signer signer.Signer `serialize:"true" json:"signer"`
	// Where to send staked tokens when done validating
	StakeOuts []*avax.TransferableOutput `serialize:"true" json:"stake"`
	// Where to send validation rewards when done validating
	ValidatorRewardsOwner fx.Owner `serialize:"true" json:"validationRewardsOwner"`
	// Where to send delegation rewards when done validating
	DelegatorRewardsOwner fx.Owner `serialize:"true" json:"delegationRewardsOwner"`
	// Fee this validator charges delegators as a percentage, times 10,000
	// For example, if this validator has DelegationShares=300,000 then they
	// take 30% of rewards from delegators
	DelegationShares uint32 `serialize:"true" json:"shares"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [AddPermissionlessValidatorTx]. Also sets the [ctx] to the given [vm.ctx] so
// that the addresses can be json marshalled into human readable format
func (tx *AddPermissionlessValidatorTx) InitCtx(ctx *snow.Context) {
	_ = "STUB: not implemented"
	return
}

func (tx *AddPermissionlessValidatorTx) SubnetID() ids.ID {
	_ = "STUB: not implemented"
	return *new(ids.ID)
}

func (tx *AddPermissionlessValidatorTx) NodeID() ids.NodeID {
	_ = "STUB: not implemented"
	return *new(ids.NodeID)
}

func (tx *AddPermissionlessValidatorTx) PublicKey() (*bls.PublicKey, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (tx *AddPermissionlessValidatorTx) PendingPriority() Priority {
	_ = "STUB: not implemented"
	return *new(Priority)
}

func (tx *AddPermissionlessValidatorTx) CurrentPriority() Priority {
	_ = "STUB: not implemented"
	return *new(Priority)
}

func (tx *AddPermissionlessValidatorTx) Stake() []*avax.TransferableOutput {
	_ = "STUB: not implemented"
	return nil
}

func (tx *AddPermissionlessValidatorTx) ValidationRewardsOwner() fx.Owner {
	_ = "STUB: not implemented"
	return *new(fx.Owner)
}

func (tx *AddPermissionlessValidatorTx) DelegationRewardsOwner() fx.Owner {
	_ = "STUB: not implemented"
	return *new(fx.Owner)
}

func (tx *AddPermissionlessValidatorTx) Shares() uint32 { _ = "STUB: not implemented"; return 0 }

// SyntacticVerify returns nil iff [tx] is valid
func (tx *AddPermissionlessValidatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

// Ensure there is provided stake

// cache that this is valid

func (tx *AddPermissionlessValidatorTx) Visit(visitor Visitor) error {
	_ = "STUB: not implemented"
	return nil
}
