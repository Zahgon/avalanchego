// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	_ StakerTx        = (*AddSubnetValidatorTx)(nil)
	_ ScheduledStaker = (*AddSubnetValidatorTx)(nil)

	errAddPrimaryNetworkValidator = errors.New("can't add primary network validator with AddSubnetValidatorTx")
)

// AddSubnetValidatorTx is an unsigned addSubnetValidatorTx
type AddSubnetValidatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// The validator
	SubnetValidator `serialize:"true" json:"validator"`
	// Auth that will be allowing this validator into the network
	SubnetAuth verify.Verifiable `serialize:"true" json:"subnetAuthorization"`
}

func (tx *AddSubnetValidatorTx) NodeID() ids.NodeID {
	_ = "STUB: not implemented"
	return *new(ids.NodeID)
}

func (*AddSubnetValidatorTx) PublicKey() (*bls.PublicKey, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (*AddSubnetValidatorTx) PendingPriority() Priority {
	_ = "STUB: not implemented"
	return *new(Priority)
}

func (*AddSubnetValidatorTx) CurrentPriority() Priority {
	_ = "STUB: not implemented"
	return *new(Priority)
}

// SyntacticVerify returns nil iff [tx] is valid
func (tx *AddSubnetValidatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

// cache that this is valid

func (tx *AddSubnetValidatorTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
