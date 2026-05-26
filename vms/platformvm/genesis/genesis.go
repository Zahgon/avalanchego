// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

// Note that since an Avalanche network has exactly one Platform Chain,
// and the Platform Chain defines the genesis state of the network
// (who is staking, which chains exist, etc.), defining the genesis
// state of the Platform Chain is the same as defining the genesis
// state of the network.

var (
	errUTXOHasNoValue         = errors.New("genesis UTXO has no value")
	errValidatorHasNoWeight   = errors.New("validator has not weight")
	errValidatorAlreadyExited = errors.New("validator would have already unstaked")
	errStakeOverflow          = errors.New("validator stake exceeds limit")

	_ utils.Sortable[Allocation] = Allocation{}
)

// UTXO adds messages to UTXOs
type UTXO struct {
	avax.UTXO `serialize:"true"`
	Message   []byte `serialize:"true" json:"message"`
}

// Genesis represents a genesis state of the platform chain
type Genesis struct {
	UTXOs         []*UTXO   `serialize:"true"`
	Validators    []*txs.Tx `serialize:"true"`
	Chains        []*txs.Tx `serialize:"true"`
	Timestamp     uint64    `serialize:"true"`
	InitialSupply uint64    `serialize:"true"`
	Message       string    `serialize:"true"`
}

func Parse(genesisBytes []byte) (*Genesis, error) { _ = "STUB: not implemented"; return nil, nil }

// Allocation is a UTXO on the Platform Chain that exists at the chain's genesis
type Allocation struct {
	Locktime uint64
	Amount   uint64
	Address  string
	Message  []byte
}

// Compare compares two allocations
func (a Allocation) Compare(other Allocation) int { _ = "STUB: not implemented"; return 0 }

// Validator represents a validator at genesis
type Validator struct {
	TxID      ids.ID
	StartTime uint64
	EndTime   uint64
	Weight    uint64
	NodeID    ids.NodeID
}

// Owner is the repr. of a reward owner at genesis
type Owner struct {
	Locktime  uint64
	Threshold uint32
	Addresses []string
}

// GenesisPermissionlessValidator represents a permissionless validator at genesis
type PermissionlessValidator struct {
	Validator
	RewardOwner        *Owner
	DelegationFee      float32
	ExactDelegationFee uint32
	Staked             []Allocation
	Signer             *signer.ProofOfPossession
}

// Chain defines a chain that exists at the network's genesis
// [GenesisData] is the initial state of the chain.
// [VMID] is the ID of the VM this chain runs.
// [FxIDs] are the IDs of the Fxs the chain supports.
// [Name] is a human-readable, non-unique name for the chain.
// [SubnetID] is the ID of the subnet that validates the chain
type Chain struct {
	GenesisData []byte
	VMID        ids.ID
	FxIDs       []ids.ID
	Name        string
	SubnetID    ids.ID
}

// bech32ToID takes bech32 address and produces a shortID
func bech32ToID(addrStr string) (ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ShortID), nil
}

// New builds the genesis state of the P-Chain (and thereby the Avalanche network.)
// [avaxAssetID] is the ID of the AVAX asset
// [networkID] is the ID of the network
// [allocations] are the UTXOs on the Platform Chain that exist at genesis.
// [validators] are the validators of the primary network at genesis.
// [chains] are the chains that exist at genesis.
// [time] is the Platform Chain's time at network genesis.
// [initialSupply] is the initial supply of the AVAX asset.
// [message] is the message to be sent to the genesis UTXOs.
func New(
	avaxAssetID ids.ID,
	networkID uint32,
	allocations []Allocation,
	validators []PermissionlessValidator,
	chains []Chain,
	time uint64,
	initialSupply uint64,
	message string,
) (*Genesis, error) {
	_ = "STUB: not implemented"
	// Specify the UTXOs on the Platform chain that exist at genesis
	return nil, nil
}

// Specify the validators that are validating the primary network at genesis

// Specify the chains that exist at genesis

// Bytes serializes the Genesis to bytes using the PlatformVM genesis codec
func (g *Genesis) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
