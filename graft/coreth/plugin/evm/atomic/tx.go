// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var _ gossip.Gossipable = (*Tx)(nil)

const (
	X2CRateUint64       uint64 = 1_000_000_000
	x2cRateMinus1Uint64 uint64 = X2CRateUint64 - 1
)

var (
	ErrWrongNetworkID = errors.New("tx was issued with a different network ID")
	ErrNilTx          = errors.New("tx is nil")
	ErrNoValueOutput  = errors.New("output has no value")
	ErrNoValueInput   = errors.New("input has no value")
	ErrNoGasUsed      = errors.New("no gas used")
	errNilOutput      = errors.New("nil output")
	errNilInput       = errors.New("nil input")
	errEmptyAssetID   = errors.New("empty asset ID is not valid")
	errNilBaseFee     = errors.New("cannot calculate dynamic fee with nil baseFee")
	errFeeOverflow    = errors.New("overflow occurred while calculating the fee")
)

// Constants for calculating the gas consumed by atomic transactions
var (
	TxBytesGas   uint64 = 1
	EVMOutputGas uint64 = (common.AddressLength + wrappers.LongLen + hashing.HashLen) * TxBytesGas
	EVMInputGas  uint64 = (common.AddressLength+wrappers.LongLen+hashing.HashLen+wrappers.LongLen)*TxBytesGas + secp256k1fx.CostPerSignature
	// X2CRate is the conversion rate between the smallest denomination on the X-Chain
	// 1 nAVAX and the smallest denomination on the C-Chain 1 wei. Where 1 nAVAX = 1 gWei.
	// This is only required for AVAX because the denomination of 1 AVAX is 9 decimal
	// places on the X and P chains, but is 18 decimal places within the EVM.
	X2CRate       = uint256.NewInt(X2CRateUint64)
	x2cRateMinus1 = uint256.NewInt(x2cRateMinus1Uint64)
)

// EVMOutput defines an output that is added to the EVM state created by import transactions
type EVMOutput struct {
	Address common.Address `serialize:"true" json:"address"`
	Amount  uint64         `serialize:"true" json:"amount"`
	AssetID ids.ID         `serialize:"true" json:"assetID"`
}

func (out EVMOutput) Compare(other EVMOutput) int { _ = "STUB: not implemented"; return 0 }

// EVMInput defines an input created from the EVM state to fund export transactions
type EVMInput struct {
	Address common.Address `serialize:"true" json:"address"`
	Amount  uint64         `serialize:"true" json:"amount"`
	AssetID ids.ID         `serialize:"true" json:"assetID"`
	Nonce   uint64         `serialize:"true" json:"nonce"`
}

func (in EVMInput) Compare(other EVMInput) int { _ = "STUB: not implemented"; return 0 }

// Verify ...
func (out *EVMOutput) Verify() error { _ = "STUB: not implemented"; return nil }

// Verify ...
func (in *EVMInput) Verify() error { _ = "STUB: not implemented"; return nil }

type AtomicBlockContext interface {
	AtomicTxs() []*Tx
}

// Visitor allows executing custom logic against the underlying transaction types.
type Visitor interface {
	ImportTx(*UnsignedImportTx) error
	ExportTx(*UnsignedExportTx) error
}

// UnsignedTx is an unsigned transaction
type UnsignedTx interface {
	Initialize(unsignedBytes, signedBytes []byte)
	ID() ids.ID
	GasUsed(fixedFee bool) (uint64, error)
	Burned(assetID ids.ID) (uint64, error)
	Bytes() []byte
	SignedBytes() []byte
}

type StateDB interface {
	AddBalance(common.Address, *uint256.Int)
	AddBalanceMultiCoin(common.Address, common.Hash, *big.Int)

	SubBalance(common.Address, *uint256.Int)
	SubBalanceMultiCoin(common.Address, common.Hash, *big.Int)

	GetBalance(common.Address) *uint256.Int
	GetBalanceMultiCoin(common.Address, common.Hash) *big.Int

	GetNonce(common.Address) uint64
	SetNonce(common.Address, uint64)
}

// UnsignedAtomicTx is an unsigned operation that can be atomically accepted
type UnsignedAtomicTx interface {
	UnsignedTx

	// InputUTXOs returns the UTXOs this tx consumes
	InputUTXOs() set.Set[ids.ID]
	// Verify attempts to verify that the transaction is well formed
	Verify(ctx *snow.Context, rules extras.Rules) error
	// Visit calls the corresponding method for the underlying transaction type
	// implementing [Visitor].
	// This is used in semantic verification of the tx.
	Visit(v Visitor) error
	// AtomicOps returns the blockchainID and set of atomic requests that
	// must be applied to shared memory for this transaction to be accepted.
	// The set of atomic requests must be returned in a consistent order.
	AtomicOps() (ids.ID, *atomic.Requests, error)

	EVMStateTransfer(ctx *snow.Context, state StateDB) error
}

// Tx is a signed transaction
type Tx struct {
	// The body of this transaction
	UnsignedAtomicTx `serialize:"true" json:"unsignedTx"`

	// The credentials of this transaction
	Creds []verify.Verifiable `serialize:"true" json:"credentials"`
}

func (tx *Tx) Compare(other *Tx) int { _ = "STUB: not implemented"; return 0 }

// Sign this transaction with the provided signers
func (tx *Tx) Sign(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Attach credentials

// Sign hash

// Attach credential

// BlockFeeContribution calculates how much AVAX towards the block fee contribution was paid
// for via this transaction denominated in [avaxAssetID] with [baseFee] used to calculate the
// cost of this transaction. This function also returns the [gasUsed] by the
// transaction for inclusion in the [baseFee] algorithm.
func (tx *Tx) BlockFeeContribution(fixedFee bool, avaxAssetID ids.ID, baseFee *big.Int) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Calculate the amount of AVAX that has been burned above the required fee denominated
// in C-Chain native 18 decimal places

func (tx *Tx) GossipID() ids.ID {
	_ = "STUB: not implemented"

	// innerSortInputsAndSigners implements sort.Interface for EVMInput
	return *new(ids.ID)
}

type innerSortInputsAndSigners struct {
	inputs  []EVMInput
	signers [][]*secp256k1.PrivateKey
}

func (ins *innerSortInputsAndSigners) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ins *innerSortInputsAndSigners) Len() int { _ = "STUB: not implemented"; return 0 }

func (ins *innerSortInputsAndSigners) Swap(i, j int) { _ = "STUB: not implemented"; return }

// SortEVMInputsAndSigners sorts the list of EVMInputs based on the addresses and assetIDs
func SortEVMInputsAndSigners(inputs []EVMInput, signers [][]*secp256k1.PrivateKey) {
	_ = "STUB: not implemented"
	return
}

// EffectiveGasPrice returns the price per gas that the transaction is paying
// denominated in aAVAX/gas.
//
// The result is rounded down to the nearest aAVAX/gas.
func EffectiveGasPrice(
	tx UnsignedTx,
	avaxAssetID ids.ID,
	isApricotPhase5 bool,
) (uint256.Int, error) {
	_ = "STUB: not implemented"
	return *new(uint256.Int), nil
}

// gasPrice = burned * x2cRate / gasUsed

// calculates the amount of AVAX that must be burned by an atomic transaction
// that consumes [cost] at [baseFee].
func CalculateDynamicFee(cost uint64, baseFee *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// fee = (cost * baseFee + [X2CRate] - 1) / [X2CRate]

// the fee is more than can fit in a uint64

func calcBytesCost(n int) uint64 { _ = "STUB: not implemented"; return 0 }
