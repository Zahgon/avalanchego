// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package wallet

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/chain/p/builder"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"

	vmsigner "github.com/ava-labs/avalanchego/vms/platformvm/signer"
	walletsigner "github.com/ava-labs/avalanchego/wallet/chain/p/signer"
)

var _ Wallet = (*withOptions)(nil)

func WithOptions(
	wallet Wallet,
	options ...common.Option,
) Wallet {
	_ = "STUB: not implemented"
	return *new(Wallet)
}

type withOptions struct {
	wallet  Wallet
	options []common.Option
}

func (w *withOptions) Builder() builder.Builder {
	_ = "STUB: not implemented"
	return *new(builder.Builder)
}

func (w *withOptions) Signer() walletsigner.Signer {
	_ = "STUB: not implemented"
	return *new(walletsigner.Signer)
}

func (w *withOptions) IssueBaseTx(
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueAddValidatorTx(
	vdr *txs.Validator,
	rewardsOwner *secp256k1fx.OutputOwners,
	shares uint32,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueAddSubnetValidatorTx(
	vdr *txs.SubnetValidator,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueRemoveSubnetValidatorTx(
	nodeID ids.NodeID,
	subnetID ids.ID,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueAddDelegatorTx(
	vdr *txs.Validator,
	rewardsOwner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueCreateChainTx(
	subnetID ids.ID,
	genesis []byte,
	vmID ids.ID,
	fxIDs []ids.ID,
	chainName string,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueCreateSubnetTx(
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueTransferSubnetOwnershipTx(
	subnetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueConvertSubnetToL1Tx(
	subnetID ids.ID,
	chainID ids.ID,
	address []byte,
	validators []*txs.ConvertSubnetToL1Validator,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueRegisterL1ValidatorTx(
	balance uint64,
	proofOfPossession [bls.SignatureLen]byte,
	message []byte,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueSetL1ValidatorWeightTx(
	message []byte,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueIncreaseL1ValidatorBalanceTx(
	validationID ids.ID,
	balance uint64,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueDisableL1ValidatorTx(
	validationID ids.ID,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueImportTx(
	sourceChainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueExportTx(
	chainID ids.ID,
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueTransformSubnetTx(
	subnetID ids.ID,
	assetID ids.ID,
	initialSupply uint64,
	maxSupply uint64,
	minConsumptionRate uint64,
	maxConsumptionRate uint64,
	minValidatorStake uint64,
	maxValidatorStake uint64,
	minStakeDuration time.Duration,
	maxStakeDuration time.Duration,
	minDelegationFee uint32,
	minDelegatorStake uint64,
	maxValidatorWeightFactor byte,
	uptimeRequirement uint32,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueAddPermissionlessValidatorTx(
	vdr *txs.SubnetValidator,
	signer vmsigner.Signer,
	assetID ids.ID,
	validationRewardsOwner *secp256k1fx.OutputOwners,
	delegationRewardsOwner *secp256k1fx.OutputOwners,
	shares uint32,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueAddPermissionlessDelegatorTx(
	vdr *txs.SubnetValidator,
	assetID ids.ID,
	rewardsOwner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueUnsignedTx(
	utx txs.UnsignedTx,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) IssueTx(
	tx *txs.Tx,
	options ...common.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}
