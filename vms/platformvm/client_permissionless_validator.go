// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
)

// ClientStaker is the representation of a staker sent via client.
type ClientStaker struct {
	// the txID of the transaction that added this staker.
	TxID ids.ID
	// the Unix time when they start staking
	StartTime uint64
	// the Unix time when they are done staking
	EndTime uint64
	// the validator weight when sampling validators
	Weight uint64
	// the node ID of the staker
	NodeID ids.NodeID
}

// ClientOwner is the repr. of a reward owner sent over client
type ClientOwner struct {
	Locktime  uint64
	Threshold uint32
	Addresses []ids.ShortID
}

type ClientL1Validator struct {
	ValidationID          *ids.ID
	RemainingBalanceOwner *ClientOwner
	DeactivationOwner     *ClientOwner
	MinNonce              *uint64
	Balance               *uint64
}

// ClientPermissionlessValidator is the repr. of a permissionless validator sent
// over client
type ClientPermissionlessValidator struct {
	ClientStaker
	ClientL1Validator
	ValidationRewardOwner  *ClientOwner
	DelegationRewardOwner  *ClientOwner
	PotentialReward        *uint64
	AccruedDelegateeReward *uint64
	DelegationFee          float32
	// Uptime is deprecated for Subnet Validators.
	// It will be available only for Primary Network Validators.
	Uptime *float32
	// Connected is deprecated for Subnet Validators.
	// It will be available only for Primary Network Validators.
	Connected *bool
	Signer    *signer.ProofOfPossession
	// The delegators delegating to this validator
	DelegatorCount  *uint64
	DelegatorWeight *uint64
	Delegators      []ClientDelegator
}

// ClientDelegator is the repr. of a delegator sent over client
type ClientDelegator struct {
	ClientStaker
	RewardOwner     *ClientOwner
	PotentialReward *uint64
}

func apiStakerToClientStaker(validator api.Staker) ClientStaker {
	_ = "STUB: not implemented"
	return *new(ClientStaker)
}

func apiOwnerToClientOwner(rewardOwner *api.Owner) (*ClientOwner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getClientPermissionlessValidators(validatorsSliceIntf []interface{}) ([]ClientPermissionlessValidator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the validator is a L1 validator, we need to set the L1 fields as well

func getClientL1Validator(apiValidator api.PermissionlessValidator) (ClientL1Validator, error) {
	_ = "STUB: not implemented"
	return *new(ClientL1Validator), nil
}

func getClientPrimaryOrSubnetValidator(apiValidator api.PermissionlessValidator) (ClientPermissionlessValidator, error) {
	_ = "STUB: not implemented"
	return *new(ClientPermissionlessValidator), nil
}
