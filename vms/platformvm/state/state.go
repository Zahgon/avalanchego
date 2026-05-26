// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/btree"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/linkeddb"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/utils/iterator"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/maybe"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/metrics"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
	"github.com/ava-labs/avalanchego/vms/platformvm/status"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

const (
	defaultTreeDegree             = 2
	indexIterationLimit           = 4096
	indexIterationSleepMultiplier = 5
	indexIterationSleepCap        = 10 * time.Second
	indexLogFrequency             = 30 * time.Second
)

var (
	errValidatorSetAlreadyPopulated   = errors.New("validator set already populated")
	errIsNotSubnet                    = errors.New("is not a subnet")
	errMissingPrimaryNetworkValidator = errors.New("missing primary network validator")
	errDeleteOrder                    = errors.New("wrong deletion order")

	BlockIDPrefix                           = []byte("blockID")
	BlockPrefix                             = []byte("block")
	ValidatorsPrefix                        = []byte("validators")
	CurrentPrefix                           = []byte("current")
	PendingPrefix                           = []byte("pending")
	ValidatorPrefix                         = []byte("validator")
	DelegatorPrefix                         = []byte("delegator")
	SubnetValidatorPrefix                   = []byte("subnetValidator")
	SubnetDelegatorPrefix                   = []byte("subnetDelegator")
	ValidatorWeightDiffsBySubnetIDPrefix    = []byte("flatValidatorDiffs")
	ValidatorWeightDiffsByHeightPrefix      = []byte("flatValidatorDiffsByHeight")
	ValidatorPublicKeyDiffsBySubnetIDPrefix = []byte("flatPublicKeyDiffs")
	ValidatorPublicKeyDiffsByHeightPrefix   = []byte("flatPublicKeyDiffsByHeight")
	TxPrefix                                = []byte("tx")
	RewardUTXOsPrefix                       = []byte("rewardUTXOs")
	UTXOPrefix                              = []byte("utxo")
	SubnetPrefix                            = []byte("subnet")
	SubnetOwnerPrefix                       = []byte("subnetOwner")
	SubnetToL1ConversionPrefix              = []byte("subnetToL1Conversion")
	TransformedSubnetPrefix                 = []byte("transformedSubnet")
	SupplyPrefix                            = []byte("supply")
	ChainPrefix                             = []byte("chain")
	ExpiryReplayProtectionPrefix            = []byte("expiryReplayProtection")
	L1Prefix                                = []byte("l1")
	WeightsPrefix                           = []byte("weights")
	SubnetIDNodeIDPrefix                    = []byte("subnetIDNodeID")
	ActivePrefix                            = []byte("active")
	InactivePrefix                          = []byte("inactive")
	SingletonPrefix                         = []byte("singleton")

	TimestampKey         = []byte("timestamp")
	FeeStateKey          = []byte("fee state")
	L1ValidatorExcessKey = []byte("l1Validator excess")
	AccruedFeesKey       = []byte("accrued fees")
	CurrentSupplyKey     = []byte("current supply")
	LastAcceptedKey      = []byte("last accepted")
	HeightsIndexedKey    = []byte("heights indexed")
	InitializedKey       = []byte("initialized")
	BlocksReindexedKey   = []byte("blocks reindexed.3")

	emptyL1ValidatorCache = &cache.Empty[ids.ID, maybe.Maybe[L1Validator]]{}
)

// Chain collects all methods to manage the state of the chain for block
// execution.
type Chain interface {
	Expiry
	L1Validators
	Stakers
	avax.UTXOAdder
	avax.UTXOGetter
	avax.UTXODeleter

	GetTimestamp() time.Time
	SetTimestamp(tm time.Time)

	GetFeeState() gas.State
	SetFeeState(f gas.State)

	GetL1ValidatorExcess() gas.Gas
	SetL1ValidatorExcess(e gas.Gas)

	GetAccruedFees() uint64
	SetAccruedFees(f uint64)

	GetCurrentSupply(subnetID ids.ID) (uint64, error)
	SetCurrentSupply(subnetID ids.ID, cs uint64)

	AddRewardUTXO(txID ids.ID, utxo *avax.UTXO)

	AddSubnet(subnetID ids.ID)

	GetSubnetOwner(subnetID ids.ID) (fx.Owner, error)
	SetSubnetOwner(subnetID ids.ID, owner fx.Owner)

	GetSubnetToL1Conversion(subnetID ids.ID) (SubnetToL1Conversion, error)
	SetSubnetToL1Conversion(subnetID ids.ID, c SubnetToL1Conversion)

	GetSubnetTransformation(subnetID ids.ID) (*txs.Tx, error)
	AddSubnetTransformation(transformSubnetTx *txs.Tx)

	AddChain(createChainTx *txs.Tx)

	GetTx(txID ids.ID) (*txs.Tx, status.Status, error)
	AddTx(tx *txs.Tx, status status.Status)
}

// Prior to https://github.com/ava-labs/avalanchego/pull/1719, blocks were
// stored as a map from blkID to stateBlk. Nodes synced prior to this PR may
// still have blocks partially stored using this legacy format.
//
// TODO: Remove after v1.14.x is activated
type stateBlk struct {
	Bytes  []byte         `serialize:"true"`
	Status choices.Status `serialize:"true"`
}

/*
 * VMDB
 * |-. validators
 * | |-. current
 * | | |-. validator
 * | | | '-. list
 * | | |   '-- txID -> uptime + potential reward + potential delegatee reward
 * | | |-. delegator
 * | | | '-. list
 * | | |   '-- txID -> potential reward
 * | | |-. subnetValidator
 * | | | '-. list
 * | | |   '-- txID -> uptime + potential reward + potential delegatee reward
 * | | '-. subnetDelegator
 * | |   '-. list
 * | |     '-- txID -> potential reward
 * | |-. pending
 * | | |-. validator
 * | | | '-. list
 * | | |   '-- txID -> nil
 * | | |-. delegator
 * | | | '-. list
 * | | |   '-- txID -> nil
 * | | |-. subnetValidator
 * | | | '-. list
 * | | |   '-- txID -> nil
 * | | '-. subnetDelegator
 * | |   '-. list
 * | |     '-- txID -> nil
 * | |-. l1
 * | | |-. weights
 * | | | '-- subnetID -> weight
 * | | |-. subnetIDNodeID
 * | | | '-- subnetID+nodeID -> validationID
 * | | |-. active
 * | | | '-- validationID -> l1Validator
 * | | '-. inactive
 * | |   '-- validationID -> l1Validator
 * | |-. weight diffs by subnet ID
 * | | '-- subnet+height+nodeID -> weightChange
 * | |-. weight diffs by height
 * | | '-- height+subnet+nodeID -> weightChange
 * | '-. pub key diffs by subnet ID
 * |   '-- subnet+height+nodeID -> uncompressed public key or nil
 * | '-. pub key diffs by height
 * |   '-- height+subnet+nodeID -> uncompressed public key or nil
 * |-. blockIDs
 * | '-- height -> blockID
 * |-. blocks
 * | '-- blockID -> block bytes
 * |-. txs
 * | '-- txID -> tx bytes + tx status
 * |- rewardUTXOs
 * | '-. txID
 * |   '-. list
 * |     '-- utxoID -> utxo bytes
 * |- utxos
 * | '-- utxoDB
 * |-. subnets
 * | '-. list
 * |   '-- txID -> nil
 * |-. subnetOwners
 * | '-- subnetID -> owner
 * |-. subnetToL1Conversions
 * | '-- subnetID -> conversionID + chainID + addr
 * |-. chains
 * | '-. subnetID
 * |   '-. list
 * |     '-- txID -> nil
 * |-. expiryReplayProtection
 * | '-- timestamp + validationID -> nil
 * '-. singletons
 *   |-- initializedKey -> nil
 *   |-- blocksReindexedKey -> nil
 *   |-- timestampKey -> timestamp
 *   |-- feeStateKey -> feeState
 *   |-- l1ValidatorExcessKey -> l1ValidatorExcess
 *   |-- accruedFeesKey -> accruedFees
 *   |-- currentSupplyKey -> currentSupply
 *   |-- lastAcceptedKey -> lastAccepted
 *   '-- heightsIndexKey -> startIndexHeight + endIndexHeight
 */
// State is the complete persisted state of the Platform chain.
type State struct {
	validatorState *validatorState
	validators     validators.Manager
	ctx            *snow.Context
	upgrades       upgrade.Config
	metrics        metrics.Metrics
	rewards        reward.Calculator

	baseDB *versiondb.Database

	expiry     *btree.BTreeG[ExpiryEntry]
	expiryDiff *expiryDiff
	expiryDB   database.Database

	activeL1Validators  *activeL1Validators
	l1ValidatorsDiff    *l1ValidatorsDiff
	l1ValidatorsDB      database.Database
	weightsCache        cache.Cacher[ids.ID, uint64] // subnetID -> total L1 validator weight
	weightsDB           database.Database
	subnetIDNodeIDCache cache.Cacher[subnetIDNodeID, bool] // subnetID+nodeID -> is validator
	subnetIDNodeIDDB    database.Database
	activeDB            database.Database
	inactiveCache       cache.Cacher[ids.ID, maybe.Maybe[L1Validator]] // validationID -> L1Validator
	inactiveDB          database.Database

	currentStakers *baseStakers
	pendingStakers *baseStakers

	currentHeight uint64

	addedBlockIDs map[uint64]ids.ID            // map of height -> blockID
	blockIDCache  cache.Cacher[uint64, ids.ID] // cache of height -> blockID; if the entry is ids.Empty, it is not in the database
	blockIDDB     database.Database

	addedBlocks map[ids.ID]block.Block            // map of blockID -> Block
	blockCache  cache.Cacher[ids.ID, block.Block] // cache of blockID -> Block; if the entry is nil, it is not in the database
	blockDB     database.Database

	validatorsDB                 database.Database
	currentValidatorsDB          database.Database
	currentValidatorBaseDB       database.Database
	currentValidatorList         linkeddb.LinkedDB
	currentDelegatorBaseDB       database.Database
	currentDelegatorList         linkeddb.LinkedDB
	currentSubnetValidatorBaseDB database.Database
	currentSubnetValidatorList   linkeddb.LinkedDB
	currentSubnetDelegatorBaseDB database.Database
	currentSubnetDelegatorList   linkeddb.LinkedDB
	pendingValidatorsDB          database.Database
	pendingValidatorBaseDB       database.Database
	pendingValidatorList         linkeddb.LinkedDB
	pendingDelegatorBaseDB       database.Database
	pendingDelegatorList         linkeddb.LinkedDB
	pendingSubnetValidatorBaseDB database.Database
	pendingSubnetValidatorList   linkeddb.LinkedDB
	pendingSubnetDelegatorBaseDB database.Database
	pendingSubnetDelegatorList   linkeddb.LinkedDB

	validatorWeightDiffsBySubnetIDDB    database.Database
	validatorWeightDiffsByHeightDB      database.Database
	validatorPublicKeyDiffsBySubnetIDDB database.Database
	validatorPublicKeyDiffsByHeightDB   database.Database

	addedTxs map[ids.ID]*txAndStatus            // map of txID -> {*txs.Tx, Status}
	txCache  cache.Cacher[ids.ID, *txAndStatus] // txID -> {*txs.Tx, Status}; if the entry is nil, it is not in the database
	txDB     database.Database

	addedRewardUTXOs map[ids.ID][]*avax.UTXO            // map of txID -> []*UTXO
	rewardUTXOsCache cache.Cacher[ids.ID, []*avax.UTXO] // txID -> []*UTXO
	rewardUTXODB     database.Database

	modifiedUTXOs map[ids.ID]*avax.UTXO // map of modified UTXOID -> *UTXO; if the UTXO is nil, it has been removed
	utxoDB        database.Database
	utxoState     avax.UTXOState

	cachedSubnetIDs []ids.ID // nil if the subnets haven't been loaded
	addedSubnetIDs  []ids.ID
	subnetBaseDB    database.Database
	subnetDB        linkeddb.LinkedDB

	subnetOwners     map[ids.ID]fx.Owner                  // map of subnetID -> owner
	subnetOwnerCache cache.Cacher[ids.ID, fxOwnerAndSize] // cache of subnetID -> owner; if the entry is nil, it is not in the database
	subnetOwnerDB    database.Database

	subnetToL1Conversions     map[ids.ID]SubnetToL1Conversion            // map of subnetID -> conversion of the subnet
	subnetToL1ConversionCache cache.Cacher[ids.ID, SubnetToL1Conversion] // cache of subnetID -> conversion
	subnetToL1ConversionDB    database.Database

	transformedSubnets     map[ids.ID]*txs.Tx            // map of subnetID -> transformSubnetTx
	transformedSubnetCache cache.Cacher[ids.ID, *txs.Tx] // cache of subnetID -> transformSubnetTx; if the entry is nil, it is not in the database
	transformedSubnetDB    database.Database

	modifiedSupplies map[ids.ID]uint64             // map of subnetID -> current supply
	supplyCache      cache.Cacher[ids.ID, *uint64] // cache of subnetID -> current supply; if the entry is nil, it is not in the database
	supplyDB         database.Database

	addedChains  map[ids.ID][]*txs.Tx                    // maps subnetID -> the newly added chains to the subnet
	chainCache   cache.Cacher[ids.ID, []*txs.Tx]         // cache of subnetID -> the chains after all local modifications []*txs.Tx
	chainDBCache cache.Cacher[ids.ID, linkeddb.LinkedDB] // cache of subnetID -> linkedDB
	chainDB      database.Database

	// The persisted fields represent the current database value
	timestamp, persistedTimestamp                 time.Time
	feeState, persistedFeeState                   gas.State
	l1ValidatorExcess, persistedL1ValidatorExcess gas.Gas
	accruedFees, persistedAccruedFees             uint64
	currentSupply, persistedCurrentSupply         uint64
	// [lastAccepted] is the most recently accepted block.
	lastAccepted, persistedLastAccepted ids.ID
	// TODO: Remove indexedHeights once v1.11.3 has been released.
	indexedHeights *heightRange
	singletonDB    database.Database
}

// heightRange is used to track which heights are safe to use the native DB
// iterator for querying validator diffs.
//
// TODO: Remove once we are guaranteed nodes can not rollback to not support the
// new indexing mechanism.
type heightRange struct {
	LowerBound uint64 `serialize:"true"`
	UpperBound uint64 `serialize:"true"`
}

type ValidatorWeightDiff struct {
	Decrease bool   `serialize:"true"`
	Amount   uint64 `serialize:"true"`
}

func (v *ValidatorWeightDiff) Add(amount uint64) error { _ = "STUB: not implemented"; return nil }

func (v *ValidatorWeightDiff) Sub(amount uint64) error { _ = "STUB: not implemented"; return nil }

func (v *ValidatorWeightDiff) add(isNegative bool, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Adding to the same sign, so we can just add the amounts together.

// Adding to the opposite sign, so we need to subtract the smaller from the
// larger and keep the sign of the larger.

type txBytesAndStatus struct {
	Tx     []byte        `serialize:"true"`
	Status status.Status `serialize:"true"`
}

type txAndStatus struct {
	tx     *txs.Tx
	status status.Status
}

type fxOwnerAndSize struct {
	owner fx.Owner
	size  int
}

type SubnetToL1Conversion struct {
	ConversionID ids.ID `serialize:"true"`
	ChainID      ids.ID `serialize:"true"`
	Addr         []byte `serialize:"true"`
}

func txSize(_ ids.ID, tx *txs.Tx) int { _ = "STUB: not implemented"; return 0 }

func txAndStatusSize(_ ids.ID, t *txAndStatus) int { _ = "STUB: not implemented"; return 0 }

func blockSize(_ ids.ID, blk block.Block) int { _ = "STUB: not implemented"; return 0 }

func New(
	db database.Database,
	genesisBytes []byte,
	metricsReg prometheus.Registerer,
	validators validators.Manager,
	upgrades upgrade.Config,
	execCfg *config.Config,
	ctx *snow.Context,
	metrics metrics.Metrics,
	rewards reward.Calculator,
) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) GetStakingInfo(subnetID ids.ID, vdrID ids.NodeID) (StakingInfo, error) {
	_ = "STUB: not implemented"
	return *new(StakingInfo), nil
}

func (s *State) SetStakingInfo(subnetID ids.ID, vdrID ids.NodeID, stakingInfo StakingInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) GetExpiryIterator() (iterator.Iterator[ExpiryEntry], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HasExpiry allows for concurrent reads.
func (s *State) HasExpiry(entry ExpiryEntry) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *State) PutExpiry(entry ExpiryEntry) { _ = "STUB: not implemented"; return }

func (s *State) DeleteExpiry(entry ExpiryEntry) { _ = "STUB: not implemented"; return }

// GetCurrentValidators returns subnet and L1 validators for the given
// subnetID along with the current P-chain height.
// This method works for both subnets and L1s. Depending of the requested
// subnet/L1 validator schema, the return values can include only subnet
// validator, only L1 validators or both if there are initial stakers in the
// L1 conversion.
func (s *State) GetCurrentValidators(ctx context.Context, subnetID ids.ID) ([]*Staker, []L1Validator, uint64, error) {
	_ = "STUB: not implemented"
	// First add the current validators (non-L1)
	return nil, nil, 0, nil
}

// Then iterate over subnetIDNodeID DB and add the L1 validators

func (s *State) GetActiveL1ValidatorsIterator() (iterator.Iterator[L1Validator], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) NumActiveL1Validators() int { _ = "STUB: not implemented"; return 0 }

func (s *State) WeightOfL1Validators(subnetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetL1Validator allows for concurrent reads.
func (s *State) GetL1Validator(validationID ids.ID) (L1Validator, error) {
	_ = "STUB: not implemented"
	return *new(L1Validator), nil
}

// getPersistedL1Validator returns the currently persisted
// L1Validator with the given validationID. It is guaranteed that any
// returned validator is either active or inactive (not deleted).
func (s *State) getPersistedL1Validator(validationID ids.ID) (L1Validator, error) {
	_ = "STUB: not implemented"
	return *new(L1Validator), nil
}

func (s *State) HasL1Validator(subnetID ids.ID, nodeID ids.NodeID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *State) PutL1Validator(l1Validator L1Validator) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) GetCurrentValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) PutCurrentValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (s *State) DeleteCurrentValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

// verifyNoDelegators checks that the validator for the subnetID and nodeID pair does not have
// delegators associated with it.
func verifyNoDelegators(cs CurrentStakers, subnetID ids.ID, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) GetCurrentDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) PutCurrentDelegator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (s *State) DeleteCurrentDelegator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (s *State) GetCurrentStakerIterator() (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) GetPendingValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) PutPendingValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (s *State) DeletePendingValidator(staker *Staker) { _ = "STUB: not implemented"; return }

func (s *State) GetPendingDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) PutPendingDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (s *State) DeletePendingDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (s *State) GetPendingStakerIterator() (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) GetSubnetIDs() ([]ids.ID, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *State) AddSubnet(subnetID ids.ID) { _ = "STUB: not implemented"; return }

func (s *State) GetSubnetOwner(subnetID ids.ID) (fx.Owner, error) {
	_ = "STUB: not implemented"
	return *new(fx.Owner), nil
}

func (s *State) SetSubnetOwner(subnetID ids.ID, owner fx.Owner) { _ = "STUB: not implemented"; return }

// GetSubnetToL1Conversion allows for concurrent reads.
func (s *State) GetSubnetToL1Conversion(subnetID ids.ID) (SubnetToL1Conversion, error) {
	_ = "STUB: not implemented"
	return *new(SubnetToL1Conversion), nil
}

func (s *State) SetSubnetToL1Conversion(subnetID ids.ID, c SubnetToL1Conversion) {
	_ = "STUB: not implemented"
	return
}

func (s *State) GetSubnetTransformation(subnetID ids.ID) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) AddSubnetTransformation(transformSubnetTxIntf *txs.Tx) {
	_ = "STUB: not implemented"
	return
}

func (s *State) GetChains(subnetID ids.ID) ([]*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) AddChain(createChainTxIntf *txs.Tx) { _ = "STUB: not implemented"; return }

func (s *State) getChainDB(subnetID ids.ID) linkeddb.LinkedDB {
	_ = "STUB: not implemented"
	return *new(linkeddb.LinkedDB)
}

func (s *State) GetTx(txID ids.ID) (*txs.Tx, status.Status, error) {
	_ = "STUB: not implemented"
	return nil, *new(status.Status), nil
}

func (s *State) AddTx(tx *txs.Tx, status status.Status) { _ = "STUB: not implemented"; return }

func (s *State) GetRewardUTXOs(txID ids.ID) ([]*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) AddRewardUTXO(txID ids.ID, utxo *avax.UTXO) { _ = "STUB: not implemented"; return }

func (s *State) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) UTXOIDs(addr []byte, start ids.ID, limit int) ([]ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) AddUTXO(utxo *avax.UTXO) { _ = "STUB: not implemented"; return }

func (s *State) DeleteUTXO(utxoID ids.ID) { _ = "STUB: not implemented"; return }

func (s *State) GetStartTime(nodeID ids.NodeID) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// GetTimestamp allows for concurrent reads.
func (s *State) GetTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *State) SetTimestamp(tm time.Time) { _ = "STUB: not implemented"; return }

func (s *State) GetFeeState() gas.State { _ = "STUB: not implemented"; return *new(gas.State) }

func (s *State) SetFeeState(feeState gas.State) { _ = "STUB: not implemented"; return }

func (s *State) GetL1ValidatorExcess() gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

func (s *State) SetL1ValidatorExcess(e gas.Gas) { _ = "STUB: not implemented"; return }

func (s *State) GetAccruedFees() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *State) SetAccruedFees(accruedFees uint64) { _ = "STUB: not implemented"; return }

func (s *State) GetLastAccepted() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *State) SetLastAccepted(lastAccepted ids.ID) { _ = "STUB: not implemented"; return }

func (s *State) GetCurrentSupply(subnetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *State) SetCurrentSupply(subnetID ids.ID, cs uint64) { _ = "STUB: not implemented"; return }

// ApplyAllValidatorWeightDiffs iterates from `startHeight` towards the genesis
// block until it has applied all of the diffs up to and including
// `endHeight`. Applying the diffs modifies `validators`.
//
// Invariant: If attempting to generate the validator set for
// `endHeight - 1`, `validators` must initially contain the validator
// weights for `startHeight`.
//
// Note: Because this function iterates towards the genesis, `startHeight`
// will typically be greater than or equal to `endHeight`. If `startHeight`
// is less than `endHeight`, no diffs will be applied.
func (s *State) ApplyAllValidatorWeightDiffs(
	ctx context.Context,
	allValidators map[ids.ID]map[ids.NodeID]*validators.GetValidatorOutput,
	startHeight uint64,
	endHeight uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the parsedHeight is less than our target endHeight, then we have
// fully processed the diffs from startHeight through endHeight.

// If this subnet previously had no validators, add the map back

// If the subnet has no validators, delete from the map

// ApplyValidatorWeightDiffs iterates from `startHeight` towards the genesis
// block until it has applied all of the diffs up to and including
// `endHeight`. Applying the diffs modifies `validators`.
//
// Invariant: If attempting to generate the validator set for
// `endHeight - 1`, `validators` must initially contain the validator
// weights for `startHeight`.
//
// Note: Because this function iterates towards the genesis, `startHeight`
// will typically be greater than or equal to `endHeight`. If `startHeight`
// is less than `endHeight`, no diffs will be applied.
func (s *State) ApplyValidatorWeightDiffs(
	ctx context.Context,
	validators map[ids.NodeID]*validators.GetValidatorOutput,
	startHeight uint64,
	endHeight uint64,
	subnetID ids.ID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the parsedHeight is less than our target endHeight, then we have
// fully processed the diffs from startHeight through endHeight.

func applyWeightDiff(
	vdrs map[ids.NodeID]*validators.GetValidatorOutput,
	nodeID ids.NodeID,
	weightDiff *ValidatorWeightDiff,
) error {
	_ = "STUB: not implemented"
	return nil
}

// This node isn't in the current validator set.

// The weight of this node changed at this block.

// The validator's weight was decreased at this block, so in the
// prior block it was higher.

// The validator's weight was increased at this block, so in the
// prior block it was lower.

// The validator's weight was 0 before this block so they weren't in the
// validator set.

// ApplyAllValidatorPublicKeyDiffs iterates from `startHeight` towards the
// genesis block until it has applied all of the diffs up to and including
// `endHeight`. Applying the diffs modifies `validators`.
//
// Invariant: If attempting to generate the validator set for
// `endHeight - 1`, `validators` must initially contain the validator
// weights for `startHeight`.
//
// Note: Because this function iterates towards the genesis, `startHeight`
// will typically be greater than or equal to `endHeight`. If `startHeight`
// is less than `endHeight`, no diffs will be applied.
func (s *State) ApplyAllValidatorPublicKeyDiffs(
	ctx context.Context,
	allValidators map[ids.ID]map[ids.NodeID]*validators.GetValidatorOutput,
	startHeight uint64,
	endHeight uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the parsedHeight is less than our target endHeight, then we have
// fully processed the diffs from startHeight through endHeight.

// A validator that is eventually removed may have a key diff before it was removed

// Nodes may see inconsistent public keys for heights before the new public
// key index was populated.

// ApplyValidatorPublicKeyDiffs iterates from `startHeight` towards the
// genesis block until it has applied all of the diffs up to and including
// `endHeight`. Applying the diffs modifies `validators`.
//
// Invariant: If attempting to generate the validator set for
// `endHeight - 1`, `validators` must initially contain the validator
// weights for `startHeight`.
//
// Note: Because this function iterates towards the genesis, `startHeight`
// will typically be greater than or equal to `endHeight`. If `startHeight`
// is less than `endHeight`, no diffs will be applied.
func (s *State) ApplyValidatorPublicKeyDiffs(
	ctx context.Context,
	validators map[ids.NodeID]*validators.GetValidatorOutput,
	startHeight uint64,
	endHeight uint64,
	subnetID ids.ID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the parsedHeight is less than our target endHeight, then we have
// fully processed the diffs from startHeight through endHeight.

// Note: this does not fallback to the linkeddb index because the linkeddb
// index does not contain entries for when to remove the public key.
//
// Nodes may see inconsistent public keys for heights before the new public
// key index was populated.

func (s *State) syncGenesis(genesisBlk block.Block, genesis *genesis.Genesis) error {
	_ = "STUB: not implemented"
	return nil
}

// Persist UTXOs that exist at genesis

// Persist primary network validator set at genesis

// We expect genesis validator txs to be either AddValidatorTx or
// AddPermissionlessValidatorTx.
//
// TODO: Enforce stricter type check

// Note: We use [StartTime()] here because genesis transactions are
// guaranteed to be pre-Durango activation.

// Ensure all chains that the genesis bytes say to create have the right
// network ID

// updateValidators is set to false here to maintain the invariant that the
// primary network's validator set is empty before the validator sets are
// initialized.
/*=updateValidators*/

// Load pulls data previously stored on disk that is expected to be in memory.
func (s *State) load() error { _ = "STUB: not implemented"; return nil }

func (s *State) loadMetadata() error { _ = "STUB: not implemented"; return nil }

// Lookup the most recently indexed range on disk. If we haven't started
// indexing the weights, then we keep the indexed heights as nil.

// If the indexed range is not up to date, then we will act as if the range
// doesn't exist.

func (s *State) loadExpiry() error { _ = "STUB: not implemented"; return nil }

func (s *State) loadActiveL1Validators() error { _ = "STUB: not implemented"; return nil }

func (s *State) loadCurrentValidators() error { _ = "STUB: not implemented"; return nil }

// Populate [StakerStartTime] using the tx as a default in the event
// it was added pre-durango and is not stored in the database.
//
// Note: We do not populate [LastUpdated] since it is expected to
// always be present on disk.

// Populate [StakerStartTime] and [LastUpdated] using the tx as a
// default in the event they are not stored in the database.

// Populate [StakerStartTime] using the tx as a default in the
// event it was added pre-durango and is not stored in the
// database.

func (s *State) loadPendingValidators() error { _ = "STUB: not implemented"; return nil }

// Invariant: initValidatorSets requires loadActiveL1Validators and
// loadCurrentValidators to have already been called.
func (s *State) initValidatorSets() error { _ = "STUB: not implemented"; return nil }

// Enforce the invariant that the validator set is empty here.

// Load active ACP-77 validators

// Load inactive ACP-77 validator weights
//
// TODO: L1s with no active weight should not be held in memory.

// It is required for the L1 validators to be loaded first so that the total
// weight is equal to the active weights here.

// This should never happen, as the total weight should always be at
// least the sum of the active weights.

// Load primary network and non-ACP77 validators

// The subnet validator's Public Key is inherited from the
// corresponding primary network validator.

func (s *State) write(updateValidators bool, height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) Close() error { _ = "STUB: not implemented"; return nil }

func (s *State) sync(genesis []byte) error { _ = "STUB: not implemented"; return nil }

// If the database wasn't previously initialized, create the platform chain
// anew using the provided genesis state.

func (s *State) init(genesisBytes []byte) error {
	// Create the genesis block and save it as being accepted (We don't do
	// genesisBlock.Accept() because then it'd look for genesisBlock's
	// non-existent parent)
	genesisID := hashing.ComputeHash256Array(genesisBytes)
	genesisBlock, err := block.NewApricotCommitBlock(genesisID, 0 /*height*/)
	if err != nil {
		return err
	}

	genesis, err := genesis.Parse(genesisBytes)
	if err != nil {
		return err
	}
	if err := s.syncGenesis(genesisBlock, genesis); err != nil {
		return err
	}

	if err := markInitialized(s.singletonDB); err != nil {
		return err
	}

	return s.Commit()
}

// AddStatelessBlock stores block as an accepted block.
// Invariant: [block] is an accepted block.
func (s *State) AddStatelessBlock(block block.Block) { _ = "STUB: not implemented"; return }

func (s *State) SetHeight(height uint64) { _ = "STUB: not implemented"; return }

// If indexedHeights hasn't been created yet, then we are newly tracking
// the range. This means we should initialize the LowerBound to the
// current height.

// Commit commits changes to the base database.
func (s *State) Commit() error { _ = "STUB: not implemented"; return nil }

// Abort discards uncommitted changes to the database.
func (s *State) Abort() { _ = "STUB: not implemented"; return }

func (s *State) Checksum() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// CommitBatch returns a batch of unwritten changes that, when written, will
// commit all pending changes to the base database.
func (s *State) CommitBatch() (database.Batch, error) {
	_ = "STUB: not implemented"
	// updateValidators is set to true here so that the validator manager is
	// kept up to date with the last accepted state.
	return *new(database.Batch), nil
}

/*=updateValidators*/

func (s *State) writeBlocks() error { _ = "STUB: not implemented"; return nil }

// Note: Evict is used rather than Put here because blk may end up
// referencing additional data (because of shared byte slices) that
// would not be properly accounted for in the cache sizing.

func (s *State) GetStatelessBlock(blockID ids.ID) (block.Block, error) {
	_ = "STUB: not implemented"
	return *new(block.Block), nil
}

func (s *State) GetBlockIDAtHeight(height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (s *State) writeExpiry() error { _ = "STUB: not implemented"; return nil }

// publicKeyDiff holds the previous and current public keys before and after applying a validator diff, respectively.
type publicKeyDiff struct {
	// prev is the public key before the current diff was applied.
	prev *bls.PublicKey
	// new is the public key after the current diff was applied.
	new *bls.PublicKey
}

// getPublicKeyDiff computes the BLS public key change for the given nodeID.
// It returns the key both before (prev) and after (new) the diff is applied.
//
// Either key in the result may be nil: prev is nil when the validator did
// not exist before this diff, and new is nil when the validator was
// deleted (not replaced) in this diff.
func getPublicKeyDiff(
	nodeID ids.NodeID,
	current map[ids.NodeID]*baseStaker,
	diffs map[ids.NodeID]*diffValidator,
) publicKeyDiff {
	_ = "STUB: not implemented"
	// If the validator was deleted, there is no post-diff validator and new
	// stays nil.
	return *new(publicKeyDiff)
}

// If the validator was removed or replaced, the prev key is on the removed
// entry.
// If the validator was unmodified, the prev key is equal to the new key.
// Otherwise, the validator was added and prev stays nil.

// updateValidatorManager updates the validator manager with the pending
// validator set changes.
//
// This function must be called prior to writeCurrentStakers and
// writeL1Validators.
//
// TODO: L1s with no active weight should not be held in memory.
func (s *State) updateValidatorManager(updateValidators bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Record the change in weight and/or public key for each validator.

// If the validator was replaced, addedWeight and removedWeight
// reflect the new and old weights respectively so we use them.
// Otherwise, a single validator is being modified (or a delegator is added),
// so the net weight change below is used.

// If the validator was deleted, we would have already removed its weight above.

// We're just adding a delegator, so we only need to update the weight of the existing validator without
// adding a new staker to the validator manager.

// Remove all deleted L1 validators. This must be done before adding new
// L1 validators to support the case where a validator is removed and then
// immediately re-added with a different validationID.

// Deleting a non-existent validator is a noop. This can happen if
// the validator was added and then immediately removed.

// Now that the removed L1 validators have been deleted, perform additions
// and modifications.

// Modifying an existing validator

// This validator's active status isn't changing. This means
// the effectiveNodeIDs are equal.

// This validator's active status is changing.

// Adding a new validator

// Update the stake metrics

type validatorDiff struct {
	weightDiff    ValidatorWeightDiff
	prevPublicKey []byte
	newPublicKey  []byte
}

// calculateValidatorDiffs calculates the validator set diff contained by the
// pending validator set changes.
//
// This function must be called prior to writeCurrentStakers.
func (s *State) calculateValidatorDiffs() (map[subnetIDNodeID]*validatorDiff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate the changes to the pre-ACP-77 validator set

// Calculate the changes to the ACP-77 validator set

// Delete the prior validator

// If the validator is being removed, we shouldn't work to re-add it.

// Add the new validator

// writeValidatorDiffs writes the validator set diff contained by the pending
// validator set changes to disk.
//
// This function must be called prior to writeCurrentStakers.
func (s *State) writeValidatorDiffs(height uint64) error { _ = "STUB: not implemented"; return nil }

// Write the changes to the database

// getOrSetDefault returns the value at k in m if it exists. If it doesn't
// exist, it sets m[k] to a new value and returns that value.
func getOrSetDefault[K comparable, V any](m map[K]*V, k K) *V {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) writeCurrentStakers(codecVersion uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// removed and added are handled with separate ifs (not
// if-else) because during a validator replacement both are set.

// The validator is being added.
//
// Invariant: It's impossible for a delegator to have been rewarded
// in the same block that the validator was added.

func writeCurrentDelegatorDiff(
	currentDelegatorList linkeddb.LinkedDB,
	validatorDiff *diffValidator,
	codecVersion uint16,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) writePendingStakers() error { _ = "STUB: not implemented"; return nil }

func writePendingDiff(
	pendingValidatorList linkeddb.LinkedDB,
	pendingDelegatorList linkeddb.LinkedDB,
	validatorDiff *diffValidator,
) error {
	_ = "STUB: not implemented"
	// removed and added are handled with separate ifs (not if-else) because
	// during a validator replacement both are set.
	return nil
}

func (s *State) writeL1Validators() error {
	_ = "STUB: not implemented"
	// Write modified weights
	return nil
}

// The L1 validator diff application is split into two loops to ensure that all
// deletions to the subnetIDNodeIDDB happen prior to any additions.
// Otherwise replacing an L1 validator by deleting it and then re-adding it with a
// different validationID could result in an inconsistent state.

// Delete the prior validator if it exists

// Update the subnetIDNodeID mapping

// Add the new validator

func (s *State) writeTXs() error { _ = "STUB: not implemented"; return nil }

// Note that we're serializing a [txBytesAndStatus] here, not a
// *txs.Tx, so we don't use [txs.Codec].

// Note: Evict is used rather than Put here because stx may end up
// referencing additional data (because of shared byte slices) that
// would not be properly accounted for in the cache sizing.

func (s *State) writeRewardUTXOs() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeUTXOs() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeSubnets() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeSubnetOwners() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeSubnetToL1Conversions() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeTransformedSubnets() error { _ = "STUB: not implemented"; return nil }

// Note: Evict is used rather than Put here because tx may end up
// referencing additional data (because of shared byte slices) that
// would not be properly accounted for in the cache sizing.

func (s *State) writeSubnetSupplies() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeChains() error { _ = "STUB: not implemented"; return nil }

func (s *State) writeMetadata() error { _ = "STUB: not implemented"; return nil }

// Returns the block and whether it is a [stateBlk].
// Invariant: blkBytes is safe to parse with blocks.GenesisCodec
//
// TODO: Remove after v1.14.x is activated
func parseStoredBlock(blkBytes []byte) (block.Block, bool, error) {
	_ = "STUB: not implemented"
	// Attempt to parse as blocks.Block
	return *new(block.Block), false, nil
}

// Fallback to [stateBlk]

// ReindexBlocks converts any block indices using the legacy storage format to
// the new format. If this database has already updated the indices, this
// function will return immediately, without iterating over the database.
//
// TODO: Remove after v1.14.x is activated
func (s *State) ReindexBlocks(lock sync.Locker, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// It is possible that new blocks are added after grabbing this iterator.
// New blocks are guaranteed to be persisted in the new format, so we don't
// need to check them.

// Releasing is done using a closure to ensure that updating blockIterator
// will result in having the most recent iterator released when executing
// the deferred function.

// This block was previously stored using the legacy format, update the
// index to remove the usage of stateBlk.

// We must hold the lock during committing to make sure we don't
// attempt to commit to disk while a block is concurrently being
// accepted.

// We release the iterator here to allow the underlying database to
// clean up deleted state.

// We take the minimum here because it's possible that the node is
// currently bootstrapping. This would mean that grabbing the lock
// could take an extremely long period of time; which we should not
// delay processing for.

// Make sure not to include the sleep duration into the next index
// duration.

// Ensure we fully iterated over all blocks before writing that indexing has
// finished.
//
// Note: This is needed because a transient read error could cause the
// iterator to stop early.

// We must hold the lock during committing to make sure we don't attempt to
// commit to disk while a block is concurrently being accepted.

func (s *State) GetUptime(vdrID ids.NodeID) (time.Duration, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

func (s *State) SetUptime(vdrID ids.NodeID, upDuration time.Duration, lastUpdated time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func markInitialized(db database.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }

func isInitialized(db database.KeyValueReader) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func putFeeState(db database.KeyValueWriter, feeState gas.State) error {
	_ = "STUB: not implemented"
	return nil
}

func getFeeState(db database.KeyValueReader) (gas.State, error) {
	_ = "STUB: not implemented"
	return *new(gas.State), nil
}
