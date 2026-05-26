// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var (
	ErrNoChangeAddress           = errors.New("no possible change address")
	ErrUnknownOutputType         = errors.New("unknown output type")
	ErrUnknownOwnerType          = errors.New("unknown owner type")
	ErrInsufficientAuthorization = errors.New("insufficient authorization")
	ErrInsufficientFunds         = errors.New("insufficient funds")

	_ Builder = (*builder)(nil)
)

// Builder provides a convenient interface for building unsigned P-chain
// transactions.
type Builder interface {
	// Context returns the configuration of the chain that this builder uses to
	// create transactions.
	Context() *Context

	// GetBalance calculates the amount of each asset that this builder has
	// control over.
	GetBalance(
		options ...common.Option,
	) (map[ids.ID]uint64, error)

	// GetImportableBalance calculates the amount of each asset that this
	// builder could import from the provided chain.
	//
	// - [chainID] specifies the chain the funds are from.
	GetImportableBalance(
		chainID ids.ID,
		options ...common.Option,
	) (map[ids.ID]uint64, error)

	// NewBaseTx creates a new simple value transfer.
	//
	// - [outputs] specifies all the recipients and amounts that should be sent
	//   from this transaction.
	NewBaseTx(
		outputs []*avax.TransferableOutput,
		options ...common.Option,
	) (*txs.BaseTx, error)

	// NewAddValidatorTx creates a new validator of the primary network.
	//
	// - [vdr] specifies all the details of the validation period such as the
	//   startTime, endTime, stake weight, and nodeID.
	// - [rewardsOwner] specifies the owner of all the rewards this validator
	//   may accrue during its validation period.
	// - [shares] specifies the fraction (out of 1,000,000) that this validator
	//   will take from delegation rewards. If 1,000,000 is provided, 100% of
	//   the delegation reward will be sent to the validator's [rewardsOwner].
	NewAddValidatorTx(
		vdr *txs.Validator,
		rewardsOwner *secp256k1fx.OutputOwners,
		shares uint32,
		options ...common.Option,
	) (*txs.AddValidatorTx, error)

	// NewAddSubnetValidatorTx creates a new validator of a subnet.
	//
	// - [vdr] specifies all the details of the validation period such as the
	//   startTime, endTime, sampling weight, nodeID, and subnetID.
	NewAddSubnetValidatorTx(
		vdr *txs.SubnetValidator,
		options ...common.Option,
	) (*txs.AddSubnetValidatorTx, error)

	// NewRemoveSubnetValidatorTx removes [nodeID] from the validator
	// set [subnetID].
	NewRemoveSubnetValidatorTx(
		nodeID ids.NodeID,
		subnetID ids.ID,
		options ...common.Option,
	) (*txs.RemoveSubnetValidatorTx, error)

	// NewAddDelegatorTx creates a new delegator to a validator on the primary
	// network.
	//
	// - [vdr] specifies all the details of the delegation period such as the
	//   startTime, endTime, stake weight, and validator's nodeID.
	// - [rewardsOwner] specifies the owner of all the rewards this delegator
	//   may accrue at the end of its delegation period.
	NewAddDelegatorTx(
		vdr *txs.Validator,
		rewardsOwner *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.AddDelegatorTx, error)

	// NewCreateChainTx creates a new chain in the named subnet.
	//
	// - [subnetID] specifies the subnet to launch the chain in.
	// - [genesis] specifies the initial state of the new chain.
	// - [vmID] specifies the vm that the new chain will run.
	// - [fxIDs] specifies all the feature extensions that the vm should be
	//   running with.
	// - [chainName] specifies a human readable name for the chain.
	NewCreateChainTx(
		subnetID ids.ID,
		genesis []byte,
		vmID ids.ID,
		fxIDs []ids.ID,
		chainName string,
		options ...common.Option,
	) (*txs.CreateChainTx, error)

	// NewCreateSubnetTx creates a new subnet with the specified owner.
	//
	// - [owner] specifies who has the ability to create new chains and add new
	//   validators to the subnet.
	NewCreateSubnetTx(
		owner *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.CreateSubnetTx, error)

	// NewTransferSubnetOwnershipTx changes the owner of the named subnet.
	//
	// - [subnetID] specifies the subnet to be modified
	// - [owner] specifies who has the ability to create new chains and add new
	//   validators to the subnet.
	NewTransferSubnetOwnershipTx(
		subnetID ids.ID,
		owner *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.TransferSubnetOwnershipTx, error)

	// NewConvertSubnetToL1Tx converts the subnet to a Permissionless L1.
	//
	// - [subnetID] specifies the subnet to be converted
	// - [chainID] specifies which chain the manager is deployed on
	// - [address] specifies the address of the manager
	// - [validators] specifies the initial L1 validators of the L1
	NewConvertSubnetToL1Tx(
		subnetID ids.ID,
		chainID ids.ID,
		address []byte,
		validators []*txs.ConvertSubnetToL1Validator,
		options ...common.Option,
	) (*txs.ConvertSubnetToL1Tx, error)

	// NewRegisterL1ValidatorTx adds a validator to an L1.
	//
	// - [balance] that the validator should allocate to continuous fees
	// - [proofOfPossession] is the BLS PoP for the key included in the Warp
	//   message
	// - [message] is the Warp message that authorizes this validator to be
	//   added
	NewRegisterL1ValidatorTx(
		balance uint64,
		proofOfPossession [bls.SignatureLen]byte,
		message []byte,
		options ...common.Option,
	) (*txs.RegisterL1ValidatorTx, error)

	// NewSetL1ValidatorWeightTx sets the weight of a validator on an L1.
	//
	// - [message] is the Warp message that authorizes this validator's weight
	//   to be changed
	NewSetL1ValidatorWeightTx(
		message []byte,
		options ...common.Option,
	) (*txs.SetL1ValidatorWeightTx, error)

	// NewIncreaseL1ValidatorBalanceTx increases the balance of a validator on
	// an L1 for the continuous fee.
	// the continuous fee.
	//
	// - [validationID] of the validator
	// - [balance] amount to increase the validator's balance by
	NewIncreaseL1ValidatorBalanceTx(
		validationID ids.ID,
		balance uint64,
		options ...common.Option,
	) (*txs.IncreaseL1ValidatorBalanceTx, error)

	// NewDisableL1ValidatorTx disables an L1 validator and returns the
	// remaining funds allocated to the continuous fee to the remaining balance
	// owner.
	//
	// - [validationID] of the validator to disable
	NewDisableL1ValidatorTx(
		validationID ids.ID,
		options ...common.Option,
	) (*txs.DisableL1ValidatorTx, error)

	// NewImportTx creates an import transaction that attempts to consume all
	// the available UTXOs and import the funds to [to].
	//
	// - [chainID] specifies the chain to be importing funds from.
	// - [to] specifies where to send the imported funds to.
	NewImportTx(
		chainID ids.ID,
		to *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.ImportTx, error)

	// NewExportTx creates an export transaction that attempts to send all the
	// provided [outputs] to the requested [chainID].
	//
	// - [chainID] specifies the chain to be exporting the funds to.
	// - [outputs] specifies the outputs to send to the [chainID].
	NewExportTx(
		chainID ids.ID,
		outputs []*avax.TransferableOutput,
		options ...common.Option,
	) (*txs.ExportTx, error)

	// NewTransformSubnetTx creates a transform subnet transaction that attempts
	// to convert the provided [subnetID] from a permissioned subnet to a
	// permissionless subnet. This transaction will convert
	// [maxSupply] - [initialSupply] of [assetID] to staking rewards.
	//
	// - [subnetID] specifies the subnet to transform.
	// - [assetID] specifies the asset to use to reward stakers on the subnet.
	// - [initialSupply] is the amount of [assetID] that will be in circulation
	//   after this transaction is accepted.
	// - [maxSupply] is the maximum total amount of [assetID] that should ever
	//   exist.
	// - [minConsumptionRate] is the rate that a staker will receive rewards
	//   if they stake with a duration of 0.
	// - [maxConsumptionRate] is the maximum rate that staking rewards should be
	//   consumed from the reward pool per year.
	// - [minValidatorStake] is the minimum amount of funds required to become a
	//   validator.
	// - [maxValidatorStake] is the maximum amount of funds a single validator
	//   can be allocated, including delegated funds.
	// - [minStakeDuration] is the minimum number of seconds a staker can stake
	//   for.
	// - [maxStakeDuration] is the maximum number of seconds a staker can stake
	//   for.
	// - [minValidatorStake] is the minimum amount of funds required to become a
	//   delegator.
	// - [maxValidatorWeightFactor] is the factor which calculates the maximum
	//   amount of delegation a validator can receive. A value of 1 effectively
	//   disables delegation.
	// - [uptimeRequirement] is the minimum percentage a validator must be
	//   online and responsive to receive a reward.
	NewTransformSubnetTx(
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
	) (*txs.TransformSubnetTx, error)

	// NewAddPermissionlessValidatorTx creates a new validator of the specified
	// subnet.
	//
	// - [vdr] specifies all the details of the validation period such as the
	//   subnetID, startTime, endTime, stake weight, and nodeID.
	// - [signer] if the subnetID is the primary network, this is the BLS key
	//   for this validator. Otherwise, this value should be the empty signer.
	// - [assetID] specifies the asset to stake.
	// - [validationRewardsOwner] specifies the owner of all the rewards this
	//   validator earns for its validation period.
	// - [delegationRewardsOwner] specifies the owner of all the rewards this
	//   validator earns for delegations during its validation period.
	// - [shares] specifies the fraction (out of 1,000,000) that this validator
	//   will take from delegation rewards. If 1,000,000 is provided, 100% of
	//   the delegation reward will be sent to the validator's [rewardsOwner].
	NewAddPermissionlessValidatorTx(
		vdr *txs.SubnetValidator,
		signer signer.Signer,
		assetID ids.ID,
		validationRewardsOwner *secp256k1fx.OutputOwners,
		delegationRewardsOwner *secp256k1fx.OutputOwners,
		shares uint32,
		options ...common.Option,
	) (*txs.AddPermissionlessValidatorTx, error)

	// NewAddPermissionlessDelegatorTx creates a new delegator of the specified
	// subnet on the specified nodeID.
	//
	// - [vdr] specifies all the details of the delegation period such as the
	//   subnetID, startTime, endTime, stake weight, and nodeID.
	// - [assetID] specifies the asset to stake.
	// - [rewardsOwner] specifies the owner of all the rewards this delegator
	//   earns during its delegation period.
	NewAddPermissionlessDelegatorTx(
		vdr *txs.SubnetValidator,
		assetID ids.ID,
		rewardsOwner *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.AddPermissionlessDelegatorTx, error)
}

type Backend interface {
	UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error)
	GetOwner(ctx context.Context, ownerID ids.ID) (fx.Owner, error)
}

type builder struct {
	addrs   set.Set[ids.ShortID]
	context *Context
	backend Backend
}

// New returns a new transaction builder.
//
//   - [addrs] is the set of addresses that the builder assumes can be used when
//     signing the transactions in the future.
//   - [context] provides the chain's configuration.
//   - [backend] provides the chain's state.
func New(
	addrs set.Set[ids.ShortID],
	context *Context,
	backend Backend,
) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builder) Context() *Context { _ = "STUB: not implemented"; return nil }

func (b *builder) GetBalance(
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) GetImportableBalance(
	chainID ids.ID,
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewBaseTx(
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.BaseTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort the outputs

func (b *builder) NewAddValidatorTx(
	vdr *txs.Validator,
	rewardsOwner *secp256k1fx.OutputOwners,
	shares uint32,
	options ...common.Option,
) (*txs.AddValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewAddSubnetValidatorTx(
	vdr *txs.SubnetValidator,
	options ...common.Option,
) (*txs.AddSubnetValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewRemoveSubnetValidatorTx(
	nodeID ids.NodeID,
	subnetID ids.ID,
	options ...common.Option,
) (*txs.RemoveSubnetValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewAddDelegatorTx(
	vdr *txs.Validator,
	rewardsOwner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.AddDelegatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewCreateChainTx(
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

func (b *builder) NewCreateSubnetTx(
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.CreateSubnetTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewTransferSubnetOwnershipTx(
	subnetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.TransferSubnetOwnershipTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewConvertSubnetToL1Tx(
	subnetID ids.ID,
	chainID ids.ID,
	address []byte,
	validators []*txs.ConvertSubnetToL1Validator,
	options ...common.Option,
) (*txs.ConvertSubnetToL1Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewRegisterL1ValidatorTx(
	balance uint64,
	proofOfPossession [bls.SignatureLen]byte,
	message []byte,
	options ...common.Option,
) (*txs.RegisterL1ValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewSetL1ValidatorWeightTx(
	message []byte,
	options ...common.Option,
) (*txs.SetL1ValidatorWeightTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewIncreaseL1ValidatorBalanceTx(
	validationID ids.ID,
	balance uint64,
	options ...common.Option,
) (*txs.IncreaseL1ValidatorBalanceTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewDisableL1ValidatorTx(
	validationID ids.ID,
	options ...common.Option,
) (*txs.DisableL1ValidatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewImportTx(
	sourceChainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.ImportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate over the unlocked UTXOs

// We couldn't spend this UTXO, so we skip to the next one

// sort imported inputs

// sort imported outputs

func (b *builder) NewExportTx(
	chainID ids.ID,
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.ExportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort exported outputs

func (b *builder) NewTransformSubnetTx(
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

func (b *builder) NewAddPermissionlessValidatorTx(
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

func (b *builder) NewAddPermissionlessDelegatorTx(
	vdr *txs.SubnetValidator,
	assetID ids.ID,
	rewardsOwner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.AddPermissionlessDelegatorTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) getBalance(
	chainID ids.ID,
	options *common.Options,
) (
	balance map[ids.ID]uint64,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate over the UTXOs

// This output is currently locked, so this output can't be
// burned.

// We couldn't spend this UTXO, so we skip to the next one

// spend takes in the requested burn amounts and the requested stake amounts.
//
//   - [toBurn] maps assetID to the amount of the asset to spend without
//     producing an output. This is typically used for fees. However, it can
//     also be used to consume some of an asset that will be produced in
//     separate outputs, such as ExportedOutputs. Only unlocked UTXOs are able
//     to be burned here.
//   - [toStake] maps assetID to the amount of the asset to spend and place into
//     the staked outputs. First locked UTXOs are attempted to be used for these
//     funds, and then unlocked UTXOs will be attempted to be used. There is no
//     preferential ordering on the unlock times.
//   - [excessAVAX] contains the amount of extra AVAX that spend can produce in
//     the change outputs in addition to the consumed and not burned AVAX.
//   - [complexity] contains the currently accrued transaction complexity that
//     will be used to calculate the required fees to be burned.
//   - [ownerOverride] optionally specifies the output owners to use for the
//     unlocked AVAX change output if no additional AVAX was needed to be
//     burned. If this value is nil, the default change owner is used.
func (b *builder) spend(
	toBurn map[ids.ID]uint64,
	toStake map[ids.ID]uint64,
	excessAVAX uint64,
	complexity gas.Dimensions,
	ownerOverride *secp256k1fx.OutputOwners,
	options *common.Options,
) (
	inputs []*avax.TransferableInput,
	changeOutputs []*avax.TransferableOutput,
	stakeOutputs []*avax.TransferableOutput,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Initialize the return values with empty slices to preserve backward
// compatibility of the json representation of transactions with no
// inputs or outputs.

// We couldn't spend this UTXO, so we skip to the next one

// This input had extra value, so some of it must be returned

// Add all the remaining stake amounts assuming unlocked UTXOs.

// AVAX is handled last to account for fees.

// We couldn't spend this UTXO, so we skip to the next one

// This input had extra value, so some of it must be returned

// If we don't need to burn or stake additional AVAX and we have
// consumed enough AVAX to pay the required fee, we should stop
// consuming UTXOs.

// We couldn't spend this UTXO, so we skip to the next one

// If we need to consume additional AVAX, we should be returning the
// change to the change address.

// Populated later if used

// It is worth adding the change output

// sort inputs
// sort the change outputs
// sort stake outputs

func (b *builder) authorize(ownerID ids.ID, options *common.Options) (*secp256k1fx.Input, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We can't authorize the subnet

func (b *builder) initCtx(tx txs.UnsignedTx) error { _ = "STUB: not implemented"; return nil }

type spendHelper struct {
	weights  gas.Dimensions
	gasPrice gas.Price

	toBurn     map[ids.ID]uint64
	toStake    map[ids.ID]uint64
	complexity gas.Dimensions

	inputs        []*avax.TransferableInput
	changeOutputs []*avax.TransferableOutput
	stakeOutputs  []*avax.TransferableOutput
}

func (s *spendHelper) addInput(input *avax.TransferableInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *spendHelper) addChangeOutput(output *avax.TransferableOutput) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *spendHelper) addStakedOutput(output *avax.TransferableOutput) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *spendHelper) addOutputComplexity(output *avax.TransferableOutput) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *spendHelper) shouldConsumeLockedAsset(assetID ids.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *spendHelper) shouldConsumeAsset(assetID ids.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *spendHelper) consumeLockedAsset(assetID ids.ID, amount uint64) uint64 {
	_ = "STUB: not implemented"
	// Stake any value that should be staked
	return 0
}

// Amount we still need to stake
// Amount available to stake

func (s *spendHelper) consumeAsset(assetID ids.ID, amount uint64) uint64 {
	_ = "STUB: not implemented"
	// Burn any value that should be burned
	return 0
}

// Amount we still need to burn
// Amount available to burn

// Stake any remaining value that should be staked

func (s *spendHelper) calculateFee() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *spendHelper) verifyAssetsConsumed() error { _ = "STUB: not implemented"; return nil }

type utxosByLocktime struct {
	unlocked []*avax.UTXO
	locked   []*avax.UTXO
}

// splitByLocktime separates the provided UTXOs into two slices:
// 1. UTXOs that are unlocked with the provided issuance time
// 2. UTXOs that are locked with the provided issuance time
func splitByLocktime(utxos []*avax.UTXO, minIssuanceTime uint64) utxosByLocktime {
	_ = "STUB: not implemented"
	return *new(utxosByLocktime)
}

type utxosByAssetID struct {
	requested []*avax.UTXO
	other     []*avax.UTXO
}

// splitByAssetID separates the provided UTXOs into two slices:
// 1. UTXOs with the provided assetID
// 2. UTXOs with a different assetID
func splitByAssetID(utxos []*avax.UTXO, assetID ids.ID) utxosByAssetID {
	_ = "STUB: not implemented"
	return *new(utxosByAssetID)
}

// unwrapOutput returns the *secp256k1fx.TransferOutput that was, potentially,
// wrapped by a *stakeable.LockOut.
//
// If the output was stakeable and locked, the locktime is returned. Otherwise,
// the locktime returned will be 0.
//
// If the output is not a, potentially wrapped, *secp256k1fx.TransferOutput, an
// error is returned.
func unwrapOutput(output verify.State) (*secp256k1fx.TransferOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
