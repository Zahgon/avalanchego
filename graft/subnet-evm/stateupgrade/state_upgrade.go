// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stateupgrade

import (
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params/extras"
)

// Configure applies the state upgrade to the state.
func Configure(stateUpgrade *extras.StateUpgrade, chainConfig ChainContext, state StateDB, blockContext BlockContext) error {
	_ = "STUB: not implemented"
	return nil
}

// upgradeAccount applies the state upgrade to the given account.
func upgradeAccount(account common.Address, upgrade extras.StateUpgradeAccount, state StateDB, isEIP158 bool) error {
	_ = "STUB: not implemented"
	// Create the account if it does not exist
	return nil
}

// if the nonce is 0, set the nonce to 1 as we would when deploying a contract at
// the address.
