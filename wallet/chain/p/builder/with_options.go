// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var _ Builder = (*withOptions)(nil)

type withOptions struct {
	builder Builder
	options []common.Option
}

// WithOptions returns a new builder that will use the given options by default.
//
//   - [builder] is the builder that will be called to perform the underlying
//     operations.
//   - [options] will be provided to the builder in addition to the options
//     provided in the method calls.
func WithOptions(builder Builder, options ...common.Option) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (w *withOptions) Context() *Context { _ = "STUB: not implemented"; return nil }

func (w *withOptions) GetBalance(
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) GetImportableBalance(
	chainID ids.ID,
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewBaseTx(
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.BaseTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewAddValidatorTx(
	vdr *txs.Validator,
	rewardsOwner *secp256k1fx.OutputOwners,
	shares uint32,
	options ...common.Option,
) (*txs.AddValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewAddSubnetValidatorTx(
	vdr *txs.SubnetValidator,
	options ...common.Option,
) (*txs.AddSubnetValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewRemoveSubnetValidatorTx(
	nodeID ids.NodeID,
	subnetID ids.ID,
	options ...common.Option,
) (*txs.RemoveSubnetValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewAddDelegatorTx(
	vdr *txs.Validator,
	rewardsOwner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.AddDelegatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewCreateChainTx(
	subnetID ids.ID,
	genesis []byte,
	vmID ids.ID,
	fxIDs []ids.ID,
	chainName string,
	options ...common.Option,
) (*txs.CreateChainTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewCreateSubnetTx(
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.CreateSubnetTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewTransferSubnetOwnershipTx(
	subnetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.TransferSubnetOwnershipTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewConvertSubnetToL1Tx(
	subnetID ids.ID,
	chainID ids.ID,
	address []byte,
	validators []*txs.ConvertSubnetToL1Validator,
	options ...common.Option,
) (*txs.ConvertSubnetToL1Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewRegisterL1ValidatorTx(
	balance uint64,
	proofOfPossession [bls.SignatureLen]byte,
	message []byte,
	options ...common.Option,
) (*txs.RegisterL1ValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewSetL1ValidatorWeightTx(
	message []byte,
	options ...common.Option,
) (*txs.SetL1ValidatorWeightTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewIncreaseL1ValidatorBalanceTx(
	validationID ids.ID,
	balance uint64,
	options ...common.Option,
) (*txs.IncreaseL1ValidatorBalanceTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewDisableL1ValidatorTx(
	validationID ids.ID,
	options ...common.Option,
) (*txs.DisableL1ValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewImportTx(
	sourceChainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.ImportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewExportTx(
	chainID ids.ID,
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.ExportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewTransformSubnetTx(
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
) (*txs.TransformSubnetTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewAddPermissionlessValidatorTx(
	vdr *txs.SubnetValidator,
	signer signer.Signer,
	assetID ids.ID,
	validationRewardsOwner *secp256k1fx.OutputOwners,
	delegationRewardsOwner *secp256k1fx.OutputOwners,
	shares uint32,
	options ...common.Option,
) (*txs.AddPermissionlessValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withOptions) NewAddPermissionlessDelegatorTx(
	vdr *txs.SubnetValidator,
	assetID ids.ID,
	rewardsOwner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.AddPermissionlessDelegatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
