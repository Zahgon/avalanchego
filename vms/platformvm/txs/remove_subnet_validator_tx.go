// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	_ UnsignedTx = (*RemoveSubnetValidatorTx)(nil)

	ErrRemovePrimaryNetworkValidator = errors.New("can't remove primary network validator with RemoveSubnetValidatorTx")
)

// Removes a validator from a subnet.
type RemoveSubnetValidatorTx struct {
	BaseTx `serialize:"true"`
	// The node to remove from the subnet.
	NodeID ids.NodeID `serialize:"true" json:"nodeID"`
	// The subnet to remove the node from.
	Subnet ids.ID `serialize:"true" json:"subnetID"`
	// Proves that the issuer has the right to remove the node from the subnet.
	SubnetAuth verify.Verifiable `serialize:"true" json:"subnetAuthorization"`
}

func (tx *RemoveSubnetValidatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

func (tx *RemoveSubnetValidatorTx) Visit(visitor Visitor) error {
	_ = "STUB: not implemented"
	return nil
}
