// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/types"
)

const MaxSubnetAddressLength = 4096

var (
	_ UnsignedTx                                  = (*ConvertSubnetToL1Tx)(nil)
	_ utils.Sortable[*ConvertSubnetToL1Validator] = (*ConvertSubnetToL1Validator)(nil)

	ErrConvertPermissionlessSubnet         = errors.New("cannot convert a permissionless subnet")
	ErrAddressTooLong                      = errors.New("address is too long")
	ErrConvertMustIncludeValidators        = errors.New("conversion must include at least one validator")
	ErrConvertValidatorsNotSortedAndUnique = errors.New("conversion validators must be sorted and unique")
	ErrZeroWeight                          = errors.New("validator weight must be non-zero")
)

type ConvertSubnetToL1Tx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// ID of the Subnet to transform
	Subnet ids.ID `serialize:"true" json:"subnetID"`
	// Chain where the Subnet manager lives
	ChainID ids.ID `serialize:"true" json:"chainID"`
	// Address of the Subnet manager
	Address types.JSONByteSlice `serialize:"true" json:"address"`
	// Initial pay-as-you-go validators for the Subnet
	Validators []*ConvertSubnetToL1Validator `serialize:"true" json:"validators"`
	// Authorizes this conversion
	SubnetAuth verify.Verifiable `serialize:"true" json:"subnetAuthorization"`
}

func (tx *ConvertSubnetToL1Tx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

func (tx *ConvertSubnetToL1Tx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }

type ConvertSubnetToL1Validator struct {
	// NodeID of this validator
	NodeID types.JSONByteSlice `serialize:"true" json:"nodeID"`
	// Weight of this validator used when sampling
	Weight uint64 `serialize:"true" json:"weight"`
	// Initial balance for this validator
	Balance uint64 `serialize:"true" json:"balance"`
	// [Signer] is the BLS key for this validator.
	// Note: We do not enforce that the BLS key is unique across all validators.
	//       This means that validators can share a key if they so choose.
	//       However, a NodeID + Subnet does uniquely map to a BLS key
	Signer signer.ProofOfPossession `serialize:"true" json:"signer"`
	// Leftover $AVAX from the [Balance] will be issued to this owner once it is
	// removed from the validator set.
	RemainingBalanceOwner message.PChainOwner `serialize:"true" json:"remainingBalanceOwner"`
	// This owner has the authority to manually deactivate this validator.
	DeactivationOwner message.PChainOwner `serialize:"true" json:"deactivationOwner"`
}

func (v *ConvertSubnetToL1Validator) Compare(o *ConvertSubnetToL1Validator) int {
	_ = "STUB: not implemented"
	return 0
}

func (v *ConvertSubnetToL1Validator) Verify() error { _ = "STUB: not implemented"; return nil }
