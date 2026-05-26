// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	errWrongNumberOfCredentials = errors.New("should have the same number of credentials as inputs")
	errIsImmutable              = errors.New("is immutable")
	errUnauthorizedModification = errors.New("unauthorized modification")
)

// verifyPoASubnetAuthorization carries out the validation for modifying a PoA
// subnet. This is an extension of [verifySubnetAuthorization] that additionally
// verifies that the subnet being modified is currently a PoA subnet.
func verifyPoASubnetAuthorization(
	fx fx.Fx,
	chainState state.Chain,
	sTx *txs.Tx,
	subnetID ids.ID,
	subnetAuth verify.Verifiable,
) ([]verify.Verifiable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifySubnetAuthorization carries out the validation for modifying a subnet.
// The last credential in [tx.Creds] is used as the subnet authorization.
// Returns the remaining tx credentials that should be used to authorize the
// other operations in the tx.
func verifySubnetAuthorization(
	fx fx.Fx,
	chainState state.Chain,
	tx *txs.Tx,
	subnetID ids.ID,
	subnetAuth verify.Verifiable,
) ([]verify.Verifiable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyAuthorization carries out the validation of an auth. The last
// credential in [tx.Creds] is used as the authorization.
// Returns the remaining tx credentials that should be used to authorize the
// other operations in the tx.
func verifyAuthorization(
	fx fx.Fx,
	tx *txs.Tx,
	owner fx.Owner,
	auth verify.Verifiable,
) ([]verify.Verifiable, error) {
	_ = "STUB: not implemented"
	return nil,

		// Ensure there is at least one credential for the subnet authorization
		nil
}
