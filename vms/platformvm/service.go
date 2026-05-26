// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/status"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/validators/fee"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/vms/types"

	avajson "github.com/ava-labs/avalanchego/utils/json"
	platformapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
)

const (
	// Max number of addresses that can be passed in as argument to GetUTXOs
	maxGetUTXOsAddrs = 1024

	// Max number of addresses that can be passed in as argument to GetStake
	maxGetStakeAddrs = 256

	// Max number of items allowed in a page
	maxPageSize = 1024

	// Note: Staker attributes cache should be large enough so that no evictions
	// happen when the API loops through all stakers.
	stakerAttributesCacheSize = 100_000
)

var (
	errMissingDecisionBlock       = errors.New("should have a decision block within the past two blocks")
	errPrimaryNetworkIsNotASubnet = errors.New("the primary network isn't a subnet")
	errNoAddresses                = errors.New("no addresses provided")
	errMissingBlockchainID        = errors.New("argument 'blockchainID' not given")
)

// Service defines the API calls that can be made to the platform chain
type Service struct {
	vm                    *VM
	addrManager           avax.AddressManager
	stakerAttributesCache *lru.Cache[ids.ID, *stakerAttributes]
}

// All attributes are optional and may not be filled for each stakerTx.
type stakerAttributes struct {
	shares                 uint32
	rewardsOwner           fx.Owner
	validationRewardsOwner fx.Owner
	delegationRewardsOwner fx.Owner
	proofOfPossession      *signer.ProofOfPossession
}

// GetHeight returns the height of the last accepted block
func (s *Service) GetHeight(r *http.Request, _ *struct{}, response *api.GetHeightResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetProposedHeight returns the current ProposerVM height
func (s *Service) GetProposedHeight(r *http.Request, _ *struct{}, reply *api.GetHeightResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type GetBalanceRequest struct {
	Addresses []string `json:"addresses"`
}

// Note: We explicitly duplicate AVAX out of the maps to ensure backwards
// compatibility.
type GetBalanceResponse struct {
	// Balance, in nAVAX, of the address
	Balance             avajson.Uint64            `json:"balance"`
	Unlocked            avajson.Uint64            `json:"unlocked"`
	LockedStakeable     avajson.Uint64            `json:"lockedStakeable"`
	LockedNotStakeable  avajson.Uint64            `json:"lockedNotStakeable"`
	Balances            map[ids.ID]avajson.Uint64 `json:"balances"`
	Unlockeds           map[ids.ID]avajson.Uint64 `json:"unlockeds"`
	LockedStakeables    map[ids.ID]avajson.Uint64 `json:"lockedStakeables"`
	LockedNotStakeables map[ids.ID]avajson.Uint64 `json:"lockedNotStakeables"`
	UTXOIDs             []*avax.UTXOID            `json:"utxoIDs"`
}

// GetBalance gets the balance of an address
func (s *Service) GetBalance(_ *http.Request, args *GetBalanceRequest, response *GetBalanceResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func newJSONBalanceMap(balanceMap map[ids.ID]uint64) map[ids.ID]avajson.Uint64 {
	_ = "STUB: not implemented"
	return nil
}

// Index is an address and an associated UTXO.
// Marks a starting or stopping point when fetching UTXOs. Used for pagination.
type Index struct {
	Address string `json:"address"` // The address as a string
	UTXO    string `json:"utxo"`    // The UTXO ID as a string
}

// GetUTXOs returns the UTXOs controlled by the given addresses
func (s *Service) GetUTXOs(_ *http.Request, args *api.GetUTXOsArgs, response *api.GetUTXOsReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSubnetArgs are the arguments to GetSubnet
type GetSubnetArgs struct {
	// ID of the subnet to retrieve information about
	SubnetID ids.ID `json:"subnetID"`
}

// GetSubnetResponse is the response from calling GetSubnet
type GetSubnetResponse struct {
	// whether it is permissioned or not
	IsPermissioned bool `json:"isPermissioned"`
	// subnet auth information for a permissioned subnet
	ControlKeys []string       `json:"controlKeys"`
	Threshold   avajson.Uint32 `json:"threshold"`
	Locktime    avajson.Uint64 `json:"locktime"`
	// subnet transformation tx ID for an elastic subnet
	SubnetTransformationTxID ids.ID `json:"subnetTransformationTxID"`
	// subnet conversion information for an L1
	ConversionID   ids.ID              `json:"conversionID"`
	ManagerChainID ids.ID              `json:"managerChainID"`
	ManagerAddress types.JSONByteSlice `json:"managerAddress"`
}

func (s *Service) GetSubnet(_ *http.Request, args *GetSubnetArgs, response *GetSubnetResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// APISubnet is a representation of a subnet used in API calls
type APISubnet struct {
	// ID of the subnet
	ID ids.ID `json:"id"`

	// Each element of [ControlKeys] the address of a public key.
	// A transaction to add a validator to this subnet requires
	// signatures from [Threshold] of these keys to be valid.
	ControlKeys []string       `json:"controlKeys"`
	Threshold   avajson.Uint32 `json:"threshold"`
}

// GetSubnetsArgs are the arguments to GetSubnets
type GetSubnetsArgs struct {
	// IDs of the subnets to retrieve information about
	// If omitted, gets all subnets
	IDs []ids.ID `json:"ids"`
}

// GetSubnetsResponse is the response from calling GetSubnets
type GetSubnetsResponse struct {
	// Each element is a subnet that exists
	// Null if there are no subnets other than the primary network
	Subnets []APISubnet `json:"subnets"`
}

// GetSubnets returns the subnets whose ID are in [args.IDs]
// The response will include the primary network
func (s *Service) GetSubnets(_ *http.Request, args *GetSubnetsArgs, response *GetSubnetsResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// all subnets

// Include primary network

// GetStakingAssetIDArgs are the arguments to GetStakingAssetID
type GetStakingAssetIDArgs struct {
	SubnetID ids.ID `json:"subnetID"`
}

// GetStakingAssetIDResponse is the response from calling GetStakingAssetID
type GetStakingAssetIDResponse struct {
	AssetID ids.ID `json:"assetID"`
}

// GetStakingAssetID returns the assetID of the token used to stake on the
// provided subnet
func (s *Service) GetStakingAssetID(_ *http.Request, args *GetStakingAssetIDArgs, response *GetStakingAssetIDResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCurrentValidatorsArgs are the arguments for calling GetCurrentValidators
type GetCurrentValidatorsArgs struct {
	// Subnet we're listing the validators of
	// If omitted, defaults to primary network
	SubnetID ids.ID `json:"subnetID"`
	// NodeIDs of validators to request. If [NodeIDs]
	// is empty, it fetches all current validators. If
	// some nodeIDs are not currently validators, they
	// will be omitted from the response.
	NodeIDs []ids.NodeID `json:"nodeIDs"`
}

// GetCurrentValidatorsReply are the results from calling GetCurrentValidators.
// Each validator contains a list of delegators to itself.
type GetCurrentValidatorsReply struct {
	Validators []any `json:"validators"`
}

func (s *Service) loadStakerTxAttributes(txID ids.ID) (*stakerAttributes, error) {
	_ = "STUB: not implemented"
	// Lookup tx from the cache first.
	return nil, nil
}

// Tx not available in cache; pull it from disk and populate the cache.

// GetCurrentValidators returns the current validators. If a single nodeID
// is provided, full delegators information is also returned. Otherwise only
// delegators' number and total weight is returned.
func (s *Service) GetCurrentValidators(request *http.Request, args *GetCurrentValidatorsArgs, reply *GetCurrentValidatorsReply) error {
	_ = "STUB: not implemented"
	return nil
}

// Create set of nodeIDs

// Check if subnet is L1

// Subnet is not L1, get validators for the subnet

// Subnet is L1, get validators for L1

func (s *Service) getL1Validators(
	ctx context.Context,
	subnetID ids.ID,
	nodeIDs set.Set[ids.NodeID],
) ([]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) getPrimaryOrSubnetValidators(subnetID ids.ID, nodeIDs set.Set[ids.NodeID]) ([]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validator's node ID as string --> Delegators to them

// Include all nodes

// TODO: avoid iterating over delegators here.

// nothing to do, continue

// TODO: avoid iterating over delegators when numNodeIDs > 1.

// Transform this to a percentage (0-100) to make it consistent
// with observedUptime in info.peers API

// If we are handling multiple nodeIDs, we don't return the
// delegator information.

// handle delegators' information

// If we are expected to populate the delegators field, we should
// always return a non-nil value.

// queried a specific validator, load all of its delegators

type GetL1ValidatorArgs struct {
	ValidationID ids.ID `json:"validationID"`
}

type GetL1ValidatorReply struct {
	platformapi.APIL1Validator
	SubnetID ids.ID `json:"subnetID"`
	// Height is the height of the last accepted block
	Height avajson.Uint64 `json:"height"`
}

// GetL1Validator returns the L1 validator if it exists
func (s *Service) GetL1Validator(r *http.Request, args *GetL1ValidatorArgs, reply *GetL1ValidatorReply) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) convertL1ValidatorToAPI(vdr state.L1Validator) (platformapi.APIL1Validator, error) {
	_ = "STUB: not implemented"
	return *new(platformapi.APIL1Validator), nil
}

// GetCurrentSupplyArgs are the arguments for calling GetCurrentSupply
type GetCurrentSupplyArgs struct {
	SubnetID ids.ID `json:"subnetID"`
}

// GetCurrentSupplyReply are the results from calling GetCurrentSupply
type GetCurrentSupplyReply struct {
	Supply avajson.Uint64 `json:"supply"`
	Height avajson.Uint64 `json:"height"`
}

// GetCurrentSupply returns an upper bound on the supply of AVAX in the system
func (s *Service) GetCurrentSupply(r *http.Request, args *GetCurrentSupplyArgs, reply *GetCurrentSupplyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// SampleValidatorsArgs are the arguments for calling SampleValidators
type SampleValidatorsArgs struct {
	// Number of validators in the sample
	Size avajson.Uint16 `json:"size"`

	// ID of subnet to sample validators from
	// If omitted, defaults to the primary network
	SubnetID ids.ID `json:"subnetID"`
}

// SampleValidatorsReply are the results from calling Sample
type SampleValidatorsReply struct {
	Validators []ids.NodeID `json:"validators"`
}

// SampleValidators returns a sampling of the list of current validators
func (s *Service) SampleValidators(_ *http.Request, args *SampleValidatorsArgs, reply *SampleValidatorsReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockchainStatusArgs is the arguments for calling GetBlockchainStatus
// [BlockchainID] is the ID of or an alias of the blockchain to get the status of.
type GetBlockchainStatusArgs struct {
	BlockchainID string `json:"blockchainID"`
}

// GetBlockchainStatusReply is the reply from calling GetBlockchainStatus
// [Status] is the blockchain's status.
type GetBlockchainStatusReply struct {
	Status status.BlockchainStatus `json:"status"`
}

// GetBlockchainStatus gets the status of a blockchain with the ID [args.BlockchainID].
func (s *Service) GetBlockchainStatus(r *http.Request, args *GetBlockchainStatusArgs, reply *GetBlockchainStatusReply) error {
	_ = "STUB: not implemented"
	return nil
}

// if its aliased then vm created this chain.

func (s *Service) nodeValidates(blockchainID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (s *Service) chainExists(ctx context.Context, blockID ids.ID, chainID ids.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ValidatedByArgs is the arguments for calling ValidatedBy
type ValidatedByArgs struct {
	// ValidatedBy returns the ID of the Subnet validating the blockchain with this ID
	BlockchainID ids.ID `json:"blockchainID"`
}

// ValidatedByResponse is the reply from calling ValidatedBy
type ValidatedByResponse struct {
	// ID of the Subnet validating the specified blockchain
	SubnetID ids.ID `json:"subnetID"`
}

// ValidatedBy returns the ID of the Subnet that validates [args.BlockchainID]
func (s *Service) ValidatedBy(r *http.Request, args *ValidatedByArgs, response *ValidatedByResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidatesArgs are the arguments to Validates
type ValidatesArgs struct {
	SubnetID ids.ID `json:"subnetID"`
}

// ValidatesResponse is the response from calling Validates
type ValidatesResponse struct {
	BlockchainIDs []ids.ID `json:"blockchainIDs"`
}

// Validates returns the IDs of the blockchains validated by [args.SubnetID]
func (s *Service) Validates(_ *http.Request, args *ValidatesArgs, response *ValidatesResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the chains that exist

// APIBlockchain is the representation of a blockchain used in API calls
type APIBlockchain struct {
	// Blockchain's ID
	ID ids.ID `json:"id"`

	// Blockchain's (non-unique) human-readable name
	Name string `json:"name"`

	// Subnet that validates the blockchain
	SubnetID ids.ID `json:"subnetID"`

	// Virtual Machine the blockchain runs
	VMID ids.ID `json:"vmID"`
}

// GetBlockchainsResponse is the response from a call to GetBlockchains
type GetBlockchainsResponse struct {
	// blockchains that exist
	Blockchains []APIBlockchain `json:"blockchains"`
}

// GetBlockchains returns all of the blockchains that exist
func (s *Service) GetBlockchains(_ *http.Request, _ *struct{}, response *GetBlockchainsResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) IssueTx(_ *http.Request, args *api.FormattedTx, response *api.JSONTxID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) GetTx(_ *http.Request, args *api.GetTxArgs, response *api.GetTxReply) error {
	_ = "STUB: not implemented"
	return nil
}

type GetTxStatusArgs struct {
	TxID ids.ID `json:"txID"`
}

type GetTxStatusResponse struct {
	Status status.Status `json:"status"`
	// Reason this tx was dropped.
	// Only non-empty if Status is dropped
	Reason string `json:"reason,omitempty"`
}

// GetTxStatus gets a tx's status
func (s *Service) GetTxStatus(_ *http.Request, args *GetTxStatusArgs, response *GetTxStatusResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// Found the status. Report it.

// The status of this transaction is not in the database - check if the tx
// is in the preferred block's db. If so, return that it's processing.

// Found the status in the preferred block's db. Report tx is processing.

// Found the tx in the mempool. Report tx is processing.

// Note: we check if tx is dropped only after having looked for it
// in the database and the mempool, because dropped txs may be re-issued.

// The tx isn't being tracked by the node.

// The tx was recently dropped because it was invalid.

type GetStakeArgs struct {
	api.JSONAddresses
	ValidatorsOnly bool                `json:"validatorsOnly"`
	Encoding       formatting.Encoding `json:"encoding"`
}

// GetStakeReply is the response from calling GetStake.
type GetStakeReply struct {
	Staked  avajson.Uint64            `json:"staked"`
	Stakeds map[ids.ID]avajson.Uint64 `json:"stakeds"`
	// String representation of staked outputs
	// Each is of type avax.TransferableOutput
	Outputs []string `json:"stakedOutputs"`
	// Encoding of [Outputs]
	Encoding formatting.Encoding `json:"encoding"`
}

// GetStake returns the amount of nAVAX that [args.Addresses] have cumulatively
// staked on the Primary Network.
//
// This method assumes that each stake output has only owner
// This method assumes only AVAX can be staked
// This method only concerns itself with the Primary Network, not subnets
// TODO: Improve the performance of this method by maintaining this data
// in a data structure rather than re-calculating it by iterating over stakers
func (s *Service) GetStake(_ *http.Request, args *GetStakeArgs, response *GetStakeReply) error {
	_ = "STUB: not implemented"
	return nil
}

// Iterates over current stakers

// Iterates over pending stakers

// GetMinStakeArgs are the arguments for calling GetMinStake.
type GetMinStakeArgs struct {
	SubnetID ids.ID `json:"subnetID"`
}

// GetMinStakeReply is the response from calling GetMinStake.
type GetMinStakeReply struct {
	//  The minimum amount of tokens one must bond to be a validator
	MinValidatorStake avajson.Uint64 `json:"minValidatorStake"`
	// Minimum stake, in nAVAX, that can be delegated on the primary network
	MinDelegatorStake avajson.Uint64 `json:"minDelegatorStake"`
}

// GetMinStake returns the minimum staking amount in nAVAX.
func (s *Service) GetMinStake(_ *http.Request, args *GetMinStakeArgs, reply *GetMinStakeReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTotalStakeArgs are the arguments for calling GetTotalStake
type GetTotalStakeArgs struct {
	// Subnet we're getting the total stake
	// If omitted returns Primary network weight
	SubnetID ids.ID `json:"subnetID"`
}

// GetTotalStakeReply is the response from calling GetTotalStake.
type GetTotalStakeReply struct {
	// Deprecated: Use Weight instead.
	Stake avajson.Uint64 `json:"stake"`

	Weight avajson.Uint64 `json:"weight"`
}

// GetTotalStake returns the total amount staked on the Primary Network
func (s *Service) GetTotalStake(_ *http.Request, args *GetTotalStakeArgs, reply *GetTotalStakeReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRewardUTXOsReply defines the GetRewardUTXOs replies returned from the API
type GetRewardUTXOsReply struct {
	// Number of UTXOs returned
	NumFetched avajson.Uint64 `json:"numFetched"`
	// The UTXOs
	UTXOs []string `json:"utxos"`
	// Encoding specifies the encoding format the UTXOs are returned in
	Encoding formatting.Encoding `json:"encoding"`
}

// GetRewardUTXOs returns the UTXOs that were rewarded after the provided
// transaction's staking period ended.
func (s *Service) GetRewardUTXOs(_ *http.Request, args *api.GetTxArgs, reply *GetRewardUTXOsReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTimestampReply is the response from GetTimestamp
type GetTimestampReply struct {
	// Current timestamp
	Timestamp time.Time `json:"timestamp"`
}

// GetTimestamp returns the current timestamp on chain.
func (s *Service) GetTimestamp(_ *http.Request, _ *struct{}, reply *GetTimestampReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAllValidatorsAtArgs are the arguments for GetAllValidatorsAt
type GetAllValidatorsAtArgs struct {
	Height platformapi.Height `json:"height"`
}

// GetAllValidatorsAtReply is the response from GetAllValidatorsAt
type GetAllValidatorsAtReply struct {
	ValidatorSets map[ids.ID]validators.WarpSet `json:"validatorSets"`
}

// GetAllValidatorsAt returns the canonical validator sets of
// all chains with at least one active validator at the specified
// height or at proposerVM height if set to [platformapi.ProposedHeight].
func (s *Service) GetAllValidatorsAt(r *http.Request, args *GetAllValidatorsAtArgs, reply *GetAllValidatorsAtReply) error {
	_ = "STUB: not implemented"
	return nil
}

// If args.Height is the sentinel value for proposed height, gets the proposed height and return it,
// else returns the input height.
func (s *Service) getQueryHeight(ctx context.Context, heightArg platformapi.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetValidatorsAtArgs is the response from GetValidatorsAt
type GetValidatorsAtArgs struct {
	Height   platformapi.Height `json:"height"`
	SubnetID ids.ID             `json:"subnetID"`
}

type jsonGetValidatorOutput struct {
	PublicKey *string        `json:"publicKey"`
	Weight    avajson.Uint64 `json:"weight"`
}

func (v *GetValidatorsAtReply) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *GetValidatorsAtReply) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// GetValidatorsAtReply is the response from GetValidatorsAt
type GetValidatorsAtReply struct {
	Validators map[ids.NodeID]*validators.GetValidatorOutput
}

// GetValidatorsAt returns the weights of the validator set of a provided subnet
// at the specified height.
func (s *Service) GetValidatorsAt(r *http.Request, args *GetValidatorsAtArgs, reply *GetValidatorsAtReply) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) GetBlock(_ *http.Request, args *api.GetBlockArgs, response *api.GetBlockResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockByHeight returns the block at the given height.
func (s *Service) GetBlockByHeight(_ *http.Request, args *api.GetBlockByHeightArgs, response *api.GetBlockResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetFeeConfig returns the dynamic fee config of the chain.
func (s *Service) GetFeeConfig(_ *http.Request, _ *struct{}, reply *gas.Config) error {
	_ = "STUB: not implemented"
	return nil
}

type GetFeeStateReply struct {
	gas.State
	Price gas.Price `json:"price"`
	Time  time.Time `json:"timestamp"`
}

// GetFeeState returns the current fee state of the chain.
func (s *Service) GetFeeState(_ *http.Request, _ *struct{}, reply *GetFeeStateReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetValidatorFeeConfig returns the validator fee config of the chain.
func (s *Service) GetValidatorFeeConfig(_ *http.Request, _ *struct{}, reply *fee.Config) error {
	_ = "STUB: not implemented"
	return nil
}

type GetValidatorFeeStateReply struct {
	Excess gas.Gas   `json:"excess"`
	Price  gas.Price `json:"price"`
	Time   time.Time `json:"timestamp"`
}

// GetValidatorFeeState returns the current validator fee state of the chain.
func (s *Service) GetValidatorFeeState(_ *http.Request, _ *struct{}, reply *GetValidatorFeeStateReply) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) getAPIOwner(owner *secp256k1fx.OutputOwners) (*platformapi.Owner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Takes in a staker and a set of addresses
// Returns:
// 1) The total amount staked by addresses in [addrs]
// 2) The staked outputs
func getStakeHelper(tx *txs.Tx, addrs set.Set[ids.ShortID], totalAmountStaked map[ids.ID]uint64) []avax.TransferableOutput {
	_ = "STUB: not implemented"
	return nil
}

// Go through all of the staked outputs

// This output can only be used for staking until [stakeOnlyUntil]

// Check whether this output is owned by one of the given addresses

// This output isn't owned by one of the given addresses. Ignore.

func toPlatformStaker(staker *state.Staker) platformapi.Staker {
	_ = "STUB: not implemented"
	return *new(platformapi.Staker)
}
