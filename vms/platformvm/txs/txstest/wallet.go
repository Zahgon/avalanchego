// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txstest

import (
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

func NewWallet(
	t testing.TB,
	ctx *snow.Context,
	config *config.Internal,
	state *state.State,
	kc *secp256k1fx.Keychain,
	subnetIDs []ids.ID,
	validationIDs []ids.ID,
	chainIDs []ids.ID,
) wallet.Wallet {
	_ = "STUB: not implemented"
	return *new(wallet.Wallet)
}

type client struct {
	backend wallet.Backend
}

func (c *client) IssueTx(
	tx *txs.Tx,
	options ...common.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}
