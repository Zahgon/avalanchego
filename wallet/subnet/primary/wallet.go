// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package primary

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/wallet/chain/c"
	"github.com/ava-labs/avalanchego/wallet/chain/x"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"

	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
)

// Wallet provides chain wallets for the primary network.
type Wallet struct {
	p pwallet.Wallet
	x x.Wallet
	c c.Wallet
}

func (w *Wallet) P() pwallet.Wallet { _ = "STUB: not implemented"; return *new(pwallet.Wallet) }

func (w *Wallet) X() x.Wallet { _ = "STUB: not implemented"; return *new(x.Wallet) }

func (w *Wallet) C() c.Wallet {
	_ = "STUB: not implemented"

	// Creates a new default wallet
	return *new(c.Wallet)
}

func NewWallet(p pwallet.Wallet, x x.Wallet, c c.Wallet) *Wallet {
	_ = "STUB: not implemented"
	return nil
}

// Creates a Wallet with the given set of options
func NewWalletWithOptions(w *Wallet, options ...common.Option) *Wallet {
	_ = "STUB: not implemented"
	return nil
}

type WalletConfig struct {
	// Subnet IDs that the wallet should know about to be able to generate
	// transactions.
	SubnetIDs []ids.ID // optional
	// Validation IDs that the wallet should know about to be able to generate
	// transactions.
	ValidationIDs []ids.ID // optional
}

// MakeWallet returns a wallet that supports issuing transactions to the chains
// living in the primary network.
//
// On creation, the wallet attaches to the provided uri and fetches all UTXOs
// that reference any of the provided keys. If the UTXOs are modified through an
// external issuance process, such as another instance of the wallet, the UTXOs
// may become out of sync. The wallet will also fetch all requested P-chain
// owners.
//
// The wallet manages all state locally, and performs all tx signing locally.
func MakeWallet(
	ctx context.Context,
	uri string,
	avaxKeychain keychain.Keychain,
	ethKeychain c.EthKeychain,
	config WalletConfig,
) (*Wallet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakePWallet returns a P-chain wallet that supports issuing transactions.
//
// On creation, the wallet attaches to the provided uri and fetches all UTXOs
// that reference any of the provided keys. If the UTXOs are modified through an
// external issuance process, such as another instance of the wallet, the UTXOs
// may become out of sync. The wallet will also fetch all requested P-chain
// owners.
//
// The wallet manages all state locally, and performs all tx signing locally.
func MakePWallet(
	ctx context.Context,
	uri string,
	keychain keychain.Keychain,
	config WalletConfig,
) (pwallet.Wallet, error) {
	_ = "STUB: not implemented"
	return *new(pwallet.Wallet), nil
}
