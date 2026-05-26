// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package wallet

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/wallet/chain/p/builder"
	"github.com/ava-labs/avalanchego/wallet/chain/p/signer"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var _ Backend = (*backend)(nil)

// Backend defines the full interface required to support a P-chain wallet.
type Backend interface {
	builder.Backend
	signer.Backend

	AcceptTx(ctx context.Context, tx *txs.Tx) error
}

type backend struct {
	common.ChainUTXOs

	ownersLock sync.RWMutex
	owners     map[ids.ID]fx.Owner // subnetID or validationID -> owner
}

func NewBackend(utxos common.ChainUTXOs, owners map[ids.ID]fx.Owner) Backend {
	_ = "STUB: not implemented"
	return *new(Backend)
}

func (b *backend) AcceptTx(ctx context.Context, tx *txs.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backend) addUTXOs(ctx context.Context, destinationChainID ids.ID, utxos []*avax.UTXO) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backend) removeUTXOs(ctx context.Context, sourceChain ids.ID, utxoIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backend) GetOwner(_ context.Context, ownerID ids.ID) (fx.Owner, error) {
	_ = "STUB: not implemented"
	return *new(fx.Owner), nil
}

func (b *backend) setOwner(ownerID ids.ID, owner fx.Owner) { _ = "STUB: not implemented"; return }
