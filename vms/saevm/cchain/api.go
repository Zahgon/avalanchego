// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/json"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/state"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/txpool"
)

type service struct {
	ctx    *snow.Context
	txpool *txpool.Txpool
	state  *state.State

	chainAlias string
	hrp        string
}

func newService(
	ctx *snow.Context,
	txpool *txpool.Txpool,
	state *state.State,
) (*service, error) {
	chainAlias, err := ctx.BCLookup.PrimaryAlias(ctx.ChainID)
	if err != nil {
		return nil, err
	}
	return &service{
		ctx:    ctx,
		txpool: txpool,
		state:  state,

		chainAlias: chainAlias,
		hrp:        constants.GetHRP(ctx.NetworkID),
	}, nil
}

func (s *service) GetUTXOs(_ *http.Request, a *api.GetUTXOsArgs, r *api.GetUTXOsReply) error {
	s.ctx.Log.Debug("API called",
		zap.String("service", "avax"),
		zap.String("method", "getUTXOs"),
		logging.UserStrings("addresses", a.Addresses),
		zap.Stringer("encoding", a.Encoding),
	)

	sourceChainID, err := s.ctx.BCLookup.Lookup(a.SourceChain)
	if err != nil {
		return fmt.Errorf("parsing source chainID %q: %w", a.SourceChain, err)
	}

	const maxAddrs = 1024
	if len(a.Addresses) > maxAddrs {
		return fmt.Errorf("too many addresses: %d exceeds %d", len(a.Addresses), maxAddrs)
	}

	addrs := make([][]byte, len(a.Addresses))
	for i, str := range a.Addresses {
		addr, err := s.parseAddress(str)
		if err != nil {
			return fmt.Errorf("parsing address %q: %w", str, err)
		}
		addrs[i] = addr[:]
	}

	var (
		startAddr ids.ShortID
		startUTXO ids.ID
	)
	if a.StartIndex != (api.Index{}) {
		startAddr, err = s.parseAddress(a.StartIndex.Address)
		if err != nil {
			return fmt.Errorf("parsing start address %q: %w", a.StartIndex.Address, err)
		}
		startUTXO, err = ids.FromString(a.StartIndex.UTXO)
		if err != nil {
			return fmt.Errorf("parsing start utxoID %q: %w", a.StartIndex.UTXO, err)
		}
	}

	const maxLimit = 1024
	utxos, lastAddr, lastUTXO, err := s.ctx.SharedMemory.Indexed(
		sourceChainID,
		addrs,
		startAddr[:],
		startUTXO[:],
		int(min(a.Limit, maxLimit)),
	)
	if err != nil {
		return fmt.Errorf("retrieving UTXOs: %w", err)
	}

	r.UTXOs = make([]string, len(utxos))
	for i, utxo := range utxos {
		r.UTXOs[i], err = formatting.Encode(a.Encoding, utxo)
		if err != nil {
			return fmt.Errorf("encoding utxo: %w", err)
		}
	}

	endAddr, err := ids.ToShortID(lastAddr)
	if err != nil {
		endAddr = ids.ShortEmpty
	}
	r.EndIndex.Address, err = address.Format(s.chainAlias, s.hrp, endAddr[:])
	if err != nil {
		return fmt.Errorf("formatting address: %w", err)
	}

	endUTXO, err := ids.ToID(lastUTXO)
	if err != nil {
		endUTXO = ids.Empty
	}
	r.EndIndex.UTXO = endUTXO.String()

	r.NumFetched = json.Uint64(len(utxos))
	r.Encoding = a.Encoding
	return nil
}

func (s *service) parseAddress(str string) (ids.ShortID, error) {
	if a, err := ids.ShortFromString(str); err == nil {
		return a, nil
	}

	chainAlias, hrp, addrBytes, err := address.Parse(str)
	if err != nil {
		return ids.ShortID{}, err
	}
	if hrp != s.hrp {
		return ids.ShortID{}, fmt.Errorf("expected hrp %q but got %q", s.hrp, hrp)
	}
	chainID, err := s.ctx.BCLookup.Lookup(chainAlias)
	if err != nil {
		return ids.ShortID{}, err
	}
	if chainID != s.ctx.ChainID {
		return ids.ShortID{}, fmt.Errorf("expected chainID to be %q but was %q", s.ctx.ChainID, chainID)
	}
	return ids.ToShortID(addrBytes)
}

func (s *service) IssueTx(_ *http.Request, a *api.FormattedTx, r *api.JSONTxID) error {
	s.ctx.Log.Debug("API called",
		zap.String("service", "avax"),
		zap.String("method", "issueTx"),
		logging.UserString("tx", a.Tx),
		zap.Stringer("encoding", a.Encoding),
	)

	txBytes, err := formatting.Decode(a.Encoding, a.Tx)
	if err != nil {
		return fmt.Errorf("decoding transaction: %w", err)
	}
	t, err := tx.Parse(txBytes)
	if err != nil {
		return fmt.Errorf("parsing transaction: %w", err)
	}

	// TODO(StephenButtolph): Push gossip the tx.

	r.TxID = t.ID()
	return s.txpool.Add(t)
}

type GetAtomicTxReply struct {
	api.FormattedTx
	Height json.Uint64 `json:"blockHeight"`
}

func (s *service) GetAtomicTx(_ *http.Request, a *api.GetTxArgs, r *GetAtomicTxReply) error {
	s.ctx.Log.Debug("API called",
		zap.String("service", "avax"),
		zap.String("method", "getAtomicTx"),
		zap.Stringer("txID", a.TxID),
		zap.Stringer("encoding", a.Encoding),
	)

	t, height, err := s.state.GetTx(a.TxID)
	if err != nil {
		return fmt.Errorf("fetching tx: %w", err)
	}
	txBytes, err := t.Bytes()
	if err != nil {
		return fmt.Errorf("marshalling tx: %w", err)
	}
	r.Tx, err = formatting.Encode(a.Encoding, txBytes)
	if err != nil {
		return fmt.Errorf("encoding tx: %w", err)
	}
	r.Encoding = a.Encoding
	r.Height = json.Uint64(height)
	return nil
}

// Client interacts with the avax API served by the C-Chain.
type Client struct {
	r rpc.EndpointRequester
}

const (
	avaxHTTPPrefix = "/ext/" + constants.ChainAliasPrefix + "/C"
	avaxHTTPPath   = avaxHTTPPrefix + avaxHTTPExtensionPath
)

// NewClient returns a [Client] that targets the C-Chain reachable at uri.
func NewClient(uri string) *Client {
	return &Client{
		r: rpc.NewEndpointRequester(uri + avaxHTTPPath),
	}
}

// IssueTx submits t to the txpool.
func (c *Client) IssueTx(ctx context.Context, t *tx.Tx, options ...rpc.Option) error {
	txBytes, err := t.Bytes()
	if err != nil {
		return fmt.Errorf("marshalling tx: %w", err)
	}
	txStr, err := formatting.Encode(formatting.Hex, txBytes)
	if err != nil {
		return fmt.Errorf("encoding tx: %w", err)
	}

	err = c.r.SendRequest(ctx, "avax.issueTx", &api.FormattedTx{
		Tx:       txStr,
		Encoding: formatting.Hex,
	}, &api.JSONTxID{}, options...)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	return nil
}

// GetAtomicTx returns an accepted cross-chain transaction along with the block
// height at which it was accepted.
func (c *Client) GetAtomicTx(ctx context.Context, txID ids.ID, options ...rpc.Option) (*tx.Tx, uint64, error) {
	res := &GetAtomicTxReply{}
	err := c.r.SendRequest(ctx, "avax.getAtomicTx", &api.GetTxArgs{
		TxID:     txID,
		Encoding: formatting.Hex,
	}, res, options...)
	if err != nil {
		return nil, 0, fmt.Errorf("sending request: %w", err)
	}

	txBytes, err := formatting.Decode(res.Encoding, res.Tx)
	if err != nil {
		return nil, 0, fmt.Errorf("decoding tx: %w", err)
	}
	t, err := tx.Parse(txBytes)
	if err != nil {
		return nil, 0, fmt.Errorf("parsing tx: %w", err)
	}
	return t, uint64(res.Height), nil
}

// GetAtomicUTXOs returns the UTXOs controlled by addrs that have been exported
// to this chain from sourceChain.
//
// Paginates via startAddr and startUTXOID; pass the zero values on the first
// call and the returned (endAddr, endUTXOID) on each subsequent call until
// fewer than limit results are returned.
func (c *Client) GetAtomicUTXOs(
	ctx context.Context,
	addrs []ids.ShortID,
	sourceChain string,
	limit uint32,
	startAddr ids.ShortID,
	startUTXOID ids.ID,
	options ...rpc.Option,
) ([]*avax.UTXO, ids.ShortID, ids.ID, error) {
	res := &api.GetUTXOsReply{}
	err := c.r.SendRequest(ctx, "avax.getUTXOs", &api.GetUTXOsArgs{
		Addresses:   ids.ShortIDsToStrings(addrs),
		SourceChain: sourceChain,
		Limit:       json.Uint32(limit),
		StartIndex: api.Index{
			Address: startAddr.String(),
			UTXO:    startUTXOID.String(),
		},
		Encoding: formatting.Hex,
	}, res, options...)
	if err != nil {
		return nil, ids.ShortID{}, ids.Empty, fmt.Errorf("sending request: %w", err)
	}

	utxos := make([]*avax.UTXO, len(res.UTXOs))
	for i, raw := range res.UTXOs {
		utxoBytes, err := formatting.Decode(res.Encoding, raw)
		if err != nil {
			return nil, ids.ShortID{}, ids.Empty, fmt.Errorf("decoding utxo %d: %w", i, err)
		}
		utxos[i], err = tx.ParseUTXO(utxoBytes)
		if err != nil {
			return nil, ids.ShortID{}, ids.Empty, fmt.Errorf("parsing utxo %d: %w", i, err)
		}
	}
	endAddr, err := address.ParseToID(res.EndIndex.Address)
	if err != nil {
		return nil, ids.ShortID{}, ids.Empty, fmt.Errorf("parsing end address: %w", err)
	}
	endUTXOID, err := ids.FromString(res.EndIndex.UTXO)
	if err != nil {
		return nil, ids.ShortID{}, ids.Empty, fmt.Errorf("parsing end utxoID: %w", err)
	}
	return utxos, endAddr, endUTXOID, nil
}
