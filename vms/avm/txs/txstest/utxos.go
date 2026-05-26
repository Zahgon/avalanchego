// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txstest

import (
	"context"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/avm/state"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/wallet/chain/x/builder"
	"github.com/ava-labs/avalanchego/wallet/chain/x/signer"
)

const maxPageSize uint64 = 1024

var (
	_ builder.Backend = (*walletUTXOsAdapter)(nil)
	_ signer.Backend  = (*walletUTXOsAdapter)(nil)
)

func newUTXOs(
	ctx *snow.Context,
	state state.State,
	sharedMemory atomic.SharedMemory,
	codec codec.Manager,
) *utxos {
	_ = "STUB: not implemented"
	return nil
}

type utxos struct {
	xchainID     ids.ID
	state        state.State
	sharedMemory atomic.SharedMemory
	codec        codec.Manager
}

func (u *utxos) UTXOs(addrs set.Set[ids.ShortID], sourceChainID ids.ID) ([]*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *utxos) GetUTXO(addrs set.Set[ids.ShortID], chainID, utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type walletUTXOsAdapter struct {
	utxos *utxos
	addrs set.Set[ids.ShortID]
}

func (w *walletUTXOsAdapter) UTXOs(_ context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletUTXOsAdapter) GetUTXO(_ context.Context, chainID, utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
