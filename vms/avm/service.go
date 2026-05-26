// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"errors"
	"net/http"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/vms/components/avax"

	avajson "github.com/ava-labs/avalanchego/utils/json"
)

const (
	// Max number of addresses that can be passed in as argument to GetUTXOs
	maxGetUTXOsAddrs = 1024

	// Max number of items allowed in a page
	maxPageSize uint64 = 1024
)

var (
	errTxNotCreateAsset = errors.New("transaction doesn't create an asset")
	errNilTxID          = errors.New("nil transaction ID")
	errNoAddresses      = errors.New("no addresses provided")
	errNotLinearized    = errors.New("chain is not linearized")
)

// FormattedAssetID defines a JSON formatted struct containing an assetID as a string
type FormattedAssetID struct {
	AssetID ids.ID `json:"assetID"`
}

// Service defines the base service for the asset vm
type Service struct{ vm *VM }

// GetBlock returns the requested block.
func (s *Service) GetBlock(_ *http.Request, args *api.GetBlockArgs, reply *api.GetBlockResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockByHeight returns the block at the given height.
func (s *Service) GetBlockByHeight(_ *http.Request, args *api.GetBlockByHeightArgs, reply *api.GetBlockResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHeight returns the height of the last accepted block.
func (s *Service) GetHeight(_ *http.Request, _ *struct{}, reply *api.GetHeightResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// IssueTx attempts to issue a transaction into consensus
func (s *Service) IssueTx(_ *http.Request, args *api.FormattedTx, reply *api.JSONTxID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTxStatusReply defines the GetTxStatus replies returned from the API
type GetTxStatusReply struct {
	Status choices.Status `json:"status"`
}

// GetTxStatus returns the status of the specified transaction
//
// Deprecated: GetTxStatus only returns Accepted or Unknown, GetTx should be
// used instead to determine if the tx was accepted.
func (s *Service) GetTxStatus(_ *http.Request, args *api.JSONTxID, reply *GetTxStatusReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTx returns the specified transaction
func (s *Service) GetTx(_ *http.Request, args *api.GetTxArgs, reply *api.GetTxReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUTXOs gets all utxos for passed in addresses
func (s *Service) GetUTXOs(_ *http.Request, args *api.GetUTXOsArgs, reply *api.GetUTXOsReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAssetDescriptionArgs are arguments for passing into GetAssetDescription requests
type GetAssetDescriptionArgs struct {
	AssetID string `json:"assetID"`
}

// GetAssetDescriptionReply defines the GetAssetDescription replies returned from the API
type GetAssetDescriptionReply struct {
	FormattedAssetID
	Name         string        `json:"name"`
	Symbol       string        `json:"symbol"`
	Denomination avajson.Uint8 `json:"denomination"`
}

// GetAssetDescription creates an empty account with the name passed in
func (s *Service) GetAssetDescription(_ *http.Request, args *GetAssetDescriptionArgs, reply *GetAssetDescriptionReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBalanceArgs are arguments for passing into GetBalance requests
type GetBalanceArgs struct {
	Address        string `json:"address"`
	AssetID        string `json:"assetID"`
	IncludePartial bool   `json:"includePartial"`
}

// GetBalanceReply defines the GetBalance replies returned from the API
type GetBalanceReply struct {
	Balance avajson.Uint64 `json:"balance"`
	UTXOIDs []avax.UTXOID  `json:"utxoIDs"`
}

// GetBalance returns the balance of an asset held by an address.
// If ![args.IncludePartial], returns only the balance held solely
// (1 out of 1 multisig) by the address and with a locktime in the past.
// Otherwise, returned balance includes assets held only partially by the
// address, and includes balances with locktime in the future.
func (s *Service) GetBalance(_ *http.Request, args *GetBalanceArgs, reply *GetBalanceReply) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO make this not specific to *secp256k1fx.TransferOutput

type Balance struct {
	AssetID string         `json:"asset"`
	Balance avajson.Uint64 `json:"balance"`
}

type GetAllBalancesArgs struct {
	api.JSONAddress
	IncludePartial bool `json:"includePartial"`
}

// GetAllBalancesReply is the response from a call to GetAllBalances
type GetAllBalancesReply struct {
	Balances []Balance `json:"balances"`
}

// GetAllBalances returns a map where:
//
// Key: ID of an asset such that [args.Address] has a non-zero balance of the asset
// Value: The balance of the asset held by the address
//
// If ![args.IncludePartial], returns only unlocked balance/UTXOs with a 1-out-of-1 multisig.
// Otherwise, returned balance/UTXOs includes assets held only partially by the
// address, and includes balances with locktime in the future.
func (s *Service) GetAllBalances(_ *http.Request, args *GetAllBalancesArgs, reply *GetAllBalancesReply) error {
	_ = "STUB: not implemented"
	return nil
}

// IDs of assets the address has a non-zero balance of
// key: ID (as bytes). value: balance of that asset

// TODO make this not specific to *secp256k1fx.TransferOutput

// 0 if key doesn't exist

type GetTxFeeReply struct {
	TxFee            avajson.Uint64 `json:"txFee"`
	CreateAssetTxFee avajson.Uint64 `json:"createAssetTxFee"`
}

func (s *Service) GetTxFee(_ *http.Request, _ *struct{}, reply *GetTxFeeReply) error {
	_ = "STUB: not implemented"
	return nil
}
