// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/json"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
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

	var addrs set.Set[ids.ShortID]
	for _, str := range a.Addresses {
		addr, err := s.parseAddress(str)
		if err != nil {
			return fmt.Errorf("parsing address %q: %w", str, err)
		}
		addrs.Add(addr)
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
	limit := int(min(a.Limit, maxLimit))
	utxos, endAddr, endUTXOID, err := GetAtomicUTXOs(
		s.ctx.SharedMemory,
		sourceChainID,
		addrs,
		startAddr,
		startUTXO,
		limit,
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

	r.EndIndex.Address, err = address.Format(s.chainAlias, s.hrp, endAddr.Bytes())
	if err != nil {
		return fmt.Errorf("formatting address: %w", err)
	}

	r.EndIndex.UTXO = endUTXOID.String()
	r.NumFetched = json.Uint64(len(utxos))
	r.Encoding = a.Encoding
	return nil
}

// GetAtomicUTXOs returns exported UTXOs such that at least one of the
// addresses in [addrs] is referenced.
//
// Returns at most [limit] UTXOs.
//
// Returns:
// * The fetched UTXOs
// * The address associated with the last UTXO fetched
// * The ID of the last UTXO fetched
// * Any error that may have occurred upstream.
func GetAtomicUTXOs(
	sharedMemory atomic.SharedMemory,
	chainID ids.ID,
	addrs set.Set[ids.ShortID],
	startAddr ids.ShortID,
	startUTXOID ids.ID,
	limit int,
) ([][]byte, ids.ShortID, ids.ID, error) {
	addrsList := make([][]byte, 0, addrs.Len())
	for addr := range addrs {
		addrsList = append(addrsList, addr.Bytes())
	}

	allUTXOBytes, lastAddr, lastUTXO, err := sharedMemory.Indexed(
		chainID,
		addrsList,
		startAddr.Bytes(),
		startUTXOID[:],
		limit,
	)
	if err != nil {
		return nil, ids.ShortID{}, ids.Empty, fmt.Errorf("fetching UTXOs: %w", err)
	}

	lastAddrID, err := ids.ToShortID(lastAddr)
	if err != nil {
		lastAddrID = ids.ShortEmpty
	}
	lastUTXOID, err := ids.ToID(lastUTXO)
	if err != nil {
		lastUTXOID = ids.Empty
	}
	return allUTXOBytes, lastAddrID, lastUTXOID, nil
}

func (s *service) parseAddress(str string) (ids.ShortID, error) {
	if a, err := ids.ShortFromString(str); err == nil {
		return a, nil
	}

	chainAlias, hrp, bytes, err := address.Parse(str)
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

	return ids.ToShortID(bytes)
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
	tx, err := tx.Parse(txBytes)
	if err != nil {
		return fmt.Errorf("parsing transaction: %w", err)
	}

	// TODO(StephenButtolph): Push gossip the tx.

	r.TxID = tx.ID()
	return s.txpool.Add(tx)
}

type apiTx struct {
	api.FormattedTx
	Height json.Uint64 `json:"blockHeight"`
}

func (s *service) GetAtomicTx(_ *http.Request, a *api.GetTxArgs, r *apiTx) error {
	s.ctx.Log.Debug("API called",
		zap.String("service", "avax"),
		zap.String("method", "getAtomicTx"),
		zap.Stringer("txID", a.TxID),
		zap.Stringer("encoding", a.Encoding),
	)

	tx, height, err := s.state.GetTx(a.TxID)
	if err != nil {
		return fmt.Errorf("fetching tx: %w", err)
	}
	txBytes, err := tx.Bytes()
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
