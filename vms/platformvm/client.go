// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/status"
	"github.com/ava-labs/avalanchego/vms/platformvm/validators/fee"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	platformapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
)

type Client struct {
	Requester rpc.EndpointRequester
}

func NewClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

// GetHeight returns the current block height.
func (c *Client) GetHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetProposedHeight returns the current height of this node's proposer VM.
func (c *Client) GetProposedHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetBalance returns the balance of addrs.
//
// Deprecated: GetUTXOs should be used instead.
func (c *Client) GetBalance(ctx context.Context, addrs []ids.ShortID, options ...rpc.Option) (*GetBalanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUTXOs returns the byte representation of the UTXOs controlled by addrs.
func (c *Client) GetUTXOs(
	ctx context.Context,
	addrs []ids.ShortID,
	limit uint32,
	startAddress ids.ShortID,
	startUTXOID ids.ID,
	options ...rpc.Option,
) ([][]byte, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}

// GetAtomicUTXOs returns the byte representation of the atomic UTXOs controlled
// by addrs from sourceChain.
func (c *Client) GetAtomicUTXOs(
	ctx context.Context,
	addrs []ids.ShortID,
	sourceChain string,
	limit uint32,
	startAddress ids.ShortID,
	startUTXOID ids.ID,
	options ...rpc.Option,
) ([][]byte, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}

// GetSubnetClientResponse is the response from calling GetSubnet on the client
type GetSubnetClientResponse struct {
	// whether it is permissioned or not
	IsPermissioned bool
	// subnet auth information for a permissioned subnet
	ControlKeys []ids.ShortID
	Threshold   uint32
	Locktime    uint64
	// subnet transformation tx ID for a permissionless subnet
	SubnetTransformationTxID ids.ID
	// subnet conversion information for an L1
	ConversionID   ids.ID
	ManagerChainID ids.ID
	ManagerAddress []byte
}

// GetSubnet returns information about the specified subnet.
func (c *Client) GetSubnet(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (GetSubnetClientResponse, error) {
	_ = "STUB: not implemented"
	return *new(GetSubnetClientResponse), nil
}

// ClientSubnet is a representation of a subnet used in client methods
type ClientSubnet struct {
	// ID of the subnet
	ID ids.ID
	// Each element of [ControlKeys] the address of a public key.
	// A transaction to add a validator to this subnet requires
	// signatures from [Threshold] of these keys to be valid.
	ControlKeys []ids.ShortID
	Threshold   uint32
}

// GetSubnets returns information about the specified subnets
//
// Deprecated: Subnets should be fetched from a dedicated indexer.
func (c *Client) GetSubnets(ctx context.Context, ids []ids.ID, options ...rpc.Option) ([]ClientSubnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStakingAssetID returns the assetID of the asset used for staking on the
// subnet corresponding to subnetID.
func (c *Client) GetStakingAssetID(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// GetCurrentValidators returns the list of current validators for subnetID.
func (c *Client) GetCurrentValidators(
	ctx context.Context,
	subnetID ids.ID,
	nodeIDs []ids.NodeID,
	options ...rpc.Option,
) ([]ClientPermissionlessValidator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// L1Validator is the response from calling GetL1Validator on the API client.
type L1Validator struct {
	SubnetID              ids.ID
	NodeID                ids.NodeID
	PublicKey             *bls.PublicKey
	RemainingBalanceOwner *secp256k1fx.OutputOwners
	DeactivationOwner     *secp256k1fx.OutputOwners
	StartTime             uint64
	Weight                uint64
	MinNonce              uint64
	// Balance is the remaining amount of AVAX this L1 validator has for paying
	// the continuous fee.
	Balance uint64
}

// GetL1Validator returns the requested L1 validator with validationID and the
// height at which it was calculated.
func (c *Client) GetL1Validator(
	ctx context.Context,
	validationID ids.ID,
	options ...rpc.Option,
) (L1Validator, uint64, error) {
	_ = "STUB: not implemented"
	return *new(L1Validator), 0, nil
}

// GetCurrentSupply returns an upper bound on the supply of AVAX in the system
// along with the chain height.
func (c *Client) GetCurrentSupply(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SampleValidators returns the nodeIDs of a sample of sampleSize validators
// from the current validator set for subnetID.
func (c *Client) SampleValidators(ctx context.Context, subnetID ids.ID, sampleSize uint16, options ...rpc.Option) ([]ids.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockchainStatus returns the current status of blockchainID.
func (c *Client) GetBlockchainStatus(ctx context.Context, blockchainID string, options ...rpc.Option) (status.BlockchainStatus, error) {
	_ = "STUB: not implemented"
	return *new(status.BlockchainStatus), nil
}

// ValidatedBy returns the subnetID that validates blockchainID.
func (c *Client) ValidatedBy(ctx context.Context, blockchainID ids.ID, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// Validates returns the list of blockchains that are validated by subnetID.
func (c *Client) Validates(ctx context.Context, subnetID ids.ID, options ...rpc.Option) ([]ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockchains returns the list of all blockchains on the platform.
//
// Deprecated: Blockchains should be fetched from a dedicated indexer.
func (c *Client) GetBlockchains(ctx context.Context, options ...rpc.Option) ([]APIBlockchain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IssueTx issues the transaction and returns its txID.
func (c *Client) IssueTx(ctx context.Context, txBytes []byte, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// GetTx returns the byte representation of txID.
func (c *Client) GetTx(ctx context.Context, txID ids.ID, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTxStatus returns the status of txID.
func (c *Client) GetTxStatus(ctx context.Context, txID ids.ID, options ...rpc.Option) (*GetTxStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStake returns the amount of nAVAX that addrs have cumulatively staked on
// the Primary Network.
//
// Deprecated: Stake should be calculated using GetTx and GetCurrentValidators.
func (c *Client) GetStake(
	ctx context.Context,
	addrs []ids.ShortID,
	validatorsOnly bool,
	options ...rpc.Option,
) (map[ids.ID]uint64, [][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetMinStake returns the minimum staking amount in nAVAX for validators and
// delegators respectively.
func (c *Client) GetMinStake(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// GetTotalStake returns the total amount (in nAVAX) staked on the network.
func (c *Client) GetTotalStake(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetRewardUTXOs returns the reward UTXOs for a transaction.
//
// Deprecated: GetRewardUTXOs should be fetched from a dedicated indexer.
func (c *Client) GetRewardUTXOs(ctx context.Context, args *api.GetTxArgs, options ...rpc.Option) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTimestamp returns the current chain timestamp.
func (c *Client) GetTimestamp(ctx context.Context, options ...rpc.Option) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// GetAllValidatorsAt returns the canonical validator sets of
// all chains with at least one active validator at the specified
// height or at proposerVM height if set to [platformapi.ProposedHeight].
func (c *Client) GetAllValidatorsAt(
	ctx context.Context,
	height platformapi.Height,
	options ...rpc.Option,
) (map[ids.ID]validators.WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetValidatorsAt returns the weights of the validator set of a provided subnet
// at the specified height or at proposerVM height if set to
// [platformapi.ProposedHeight].
func (c *Client) GetValidatorsAt(
	ctx context.Context,
	subnetID ids.ID,
	height platformapi.Height,
	options ...rpc.Option,
) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlock returns blockID.
func (c *Client) GetBlock(ctx context.Context, blockID ids.ID, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockByHeight returns the block at the given height.
func (c *Client) GetBlockByHeight(ctx context.Context, height uint64, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFeeConfig returns the dynamic fee config.
func (c *Client) GetFeeConfig(ctx context.Context, options ...rpc.Option) (*gas.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFeeState returns the current fee state.
func (c *Client) GetFeeState(ctx context.Context, options ...rpc.Option) (
	gas.State,
	gas.Price,
	time.Time,
	error,
) {
	_ = "STUB: not implemented"
	return *new(gas.State), *new(gas.Price), *new(time.Time), nil
}

// GetValidatorFeeConfig returns the validator fee config.
func (c *Client) GetValidatorFeeConfig(ctx context.Context, options ...rpc.Option) (*fee.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetValidatorFeeState returns the current validator fee state.
func (c *Client) GetValidatorFeeState(ctx context.Context, options ...rpc.Option) (
	gas.Gas,
	gas.Price,
	time.Time,
	error,
) {
	_ = "STUB: not implemented"
	return *new(gas.Gas), *new(gas.Price), *new(time.Time), nil
}

func (c *Client) AwaitTxAccepted(ctx context.Context, txID ids.ID, freq time.Duration, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSubnetOwners returns a map of subnet ID to current subnet's owner
func (c *Client) GetSubnetOwners(ctx context.Context, subnetIDs ...ids.ID) (map[ids.ID]fx.Owner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDeactivationOwners returns a map of validation ID to deactivation owners
func (c *Client) GetDeactivationOwners(ctx context.Context, validationIDs ...ids.ID) (map[ids.ID]fx.Owner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetOwners returns the union of GetSubnetOwners and GetDeactivationOwners.
func (c *Client) GetOwners(ctx context.Context, subnetIDs []ids.ID, validationIDs []ids.ID) (map[ids.ID]fx.Owner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
