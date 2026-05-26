// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var (
	_ UTXOs      = (*utxos)(nil)
	_ ChainUTXOs = (*chainUTXOs)(nil)
)

type UTXOs interface {
	AddUTXO(ctx context.Context, sourceChainID, destinationChainID ids.ID, utxo *avax.UTXO) error
	RemoveUTXO(ctx context.Context, sourceChainID, destinationChainID, utxoID ids.ID) error

	UTXOs(ctx context.Context, sourceChainID, destinationChainID ids.ID) ([]*avax.UTXO, error)
	GetUTXO(ctx context.Context, sourceChainID, destinationChainID, utxoID ids.ID) (*avax.UTXO, error)
}

type ChainUTXOs interface {
	AddUTXO(ctx context.Context, destinationChainID ids.ID, utxo *avax.UTXO) error
	RemoveUTXO(ctx context.Context, sourceChainID, utxoID ids.ID) error

	UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error)
	GetUTXO(ctx context.Context, sourceChainID, utxoID ids.ID) (*avax.UTXO, error)
}

func NewUTXOs() UTXOs { _ = "STUB: not implemented"; return *new(UTXOs) }

func NewChainUTXOs(chainID ids.ID, utxos UTXOs) ChainUTXOs {
	_ = "STUB: not implemented"
	return *new(ChainUTXOs)
}

type utxos struct {
	lock sync.RWMutex
	// sourceChainID -> destinationChainID -> utxoID -> utxo
	sourceToDestToUTXOIDToUTXO map[ids.ID]map[ids.ID]map[ids.ID]*avax.UTXO
}

func (u *utxos) AddUTXO(_ context.Context, sourceChainID, destinationChainID ids.ID, utxo *avax.UTXO) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *utxos) RemoveUTXO(_ context.Context, sourceChainID, destinationChainID, utxoID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *utxos) UTXOs(_ context.Context, sourceChainID, destinationChainID ids.ID) ([]*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *utxos) GetUTXO(_ context.Context, sourceChainID, destinationChainID, utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type chainUTXOs struct {
	utxos   UTXOs
	chainID ids.ID
}

func (c *chainUTXOs) AddUTXO(ctx context.Context, destinationChainID ids.ID, utxo *avax.UTXO) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *chainUTXOs) RemoveUTXO(ctx context.Context, sourceChainID, utxoID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *chainUTXOs) UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *chainUTXOs) GetUTXO(ctx context.Context, sourceChainID, utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
