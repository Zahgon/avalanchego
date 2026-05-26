// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"errors"
	"net/http"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/txpool"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/client"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/json"

	atomicstate "github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/state"
	avalanchegossip "github.com/ava-labs/avalanchego/network/p2p/gossip"
)

const (
	// Max number of addresses that can be passed in as argument to GetUTXOs
	maxGetUTXOsAddrs = 1024
	maxUTXOsToFetch  = 1024
)

var (
	errNoAddresses   = errors.New("no addresses provided")
	errNoSourceChain = errors.New("no source chain provided")
	errNilTxID       = errors.New("nil transaction ID")
)

// AvaxAPI offers Avalanche network related API methods
type AvaxAPI struct {
	// If non-nil, bc is used to prevent returning any transactions as accepted
	// prior to having advanced the chain state to the block containing the
	// transaction.
	bc *core.BlockChain

	Context      *snow.Context
	Mempool      *txpool.Mempool
	PushGossiper *avalanchegossip.PushGossiper[*atomic.Tx]
	AcceptedTxs  *atomicstate.AtomicRepository
}

// GetUTXOs gets all utxos for passed in addresses
func (service *AvaxAPI) GetUTXOs(_ *http.Request, args *api.GetUTXOsArgs, reply *api.GetUTXOsReply) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *AvaxAPI) IssueTx(_ *http.Request, args *api.FormattedTx, response *api.JSONTxID) error {
	_ = "STUB: not implemented"
	return nil
}

// If the tx was either already in the mempool or was added to the mempool,
// we push it to the network for inclusion. If the tx was previously added
// to the mempool through p2p gossip, this will ensure this node also pushes
// it to the network.

// GetAtomicTxStatus returns the status of the specified transaction
func (service *AvaxAPI) GetAtomicTxStatus(_ *http.Request, args *api.JSONTxID, reply *client.GetAtomicTxStatusReply) error {
	_ = "STUB: not implemented"
	return nil
}

type FormattedTx struct {
	api.FormattedTx
	BlockHeight *json.Uint64 `json:"blockHeight,omitempty"`
}

// GetAtomicTx returns the specified transaction
func (service *AvaxAPI) GetAtomicTx(_ *http.Request, args *api.GetTxArgs, reply *FormattedTx) error {
	_ = "STUB: not implemented"
	return nil
}

// getAtomicTx returns the requested transaction, status, and height.
// If the status is [atomic.Unknown], then the returned transaction will be nil.
func (service *AvaxAPI) getAtomicTx(txID ids.ID) (*atomic.Tx, atomic.Status, *json.Uint64, error) {
	_ = "STUB: not implemented"
	return nil, *new(atomic.Status), nil, nil
}

// Since chain state updates run asynchronously with VM block
// acceptance, avoid returning [atomic.Accepted] until the chain state
// reaches the block containing the atomic tx.
