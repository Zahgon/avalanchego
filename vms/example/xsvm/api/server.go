// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"net/http"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/builder"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/chain"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

// Server defines the xsvm API server.
type Server interface {
	Network(r *http.Request, args *struct{}, reply *NetworkReply) error
	Genesis(r *http.Request, args *struct{}, reply *GenesisReply) error
	Nonce(r *http.Request, args *NonceArgs, reply *NonceReply) error
	Balance(r *http.Request, args *BalanceArgs, reply *BalanceReply) error
	Loan(r *http.Request, args *LoanArgs, reply *LoanReply) error
	IssueTx(r *http.Request, args *IssueTxArgs, reply *IssueTxReply) error
	LastAccepted(r *http.Request, args *struct{}, reply *LastAcceptedReply) error
	Block(r *http.Request, args *BlockArgs, reply *BlockReply) error
	Message(r *http.Request, args *MessageArgs, reply *MessageReply) error
}

func NewServer(
	ctx *snow.Context,
	genesis *genesis.Genesis,
	state database.KeyValueReader,
	chain chain.Chain,
	builder builder.Builder,
) Server {
	_ = "STUB: not implemented"
	return *new(Server)
}

type server struct {
	ctx     *snow.Context
	genesis *genesis.Genesis
	state   database.KeyValueReader
	chain   chain.Chain
	builder builder.Builder
}

type NetworkReply struct {
	NetworkID uint32 `json:"networkID"`
	SubnetID  ids.ID `json:"subnetID"`
	ChainID   ids.ID `json:"chainID"`
}

func (s *server) Network(_ *http.Request, _ *struct{}, reply *NetworkReply) error {
	_ = "STUB: not implemented"
	return nil
}

type GenesisReply struct {
	Genesis *genesis.Genesis `json:"genesis"`
}

func (s *server) Genesis(_ *http.Request, _ *struct{}, reply *GenesisReply) error {
	_ = "STUB: not implemented"
	return nil
}

type NonceArgs struct {
	Address ids.ShortID `json:"address"`
}

type NonceReply struct {
	Nonce uint64 `json:"nonce"`
}

func (s *server) Nonce(_ *http.Request, args *NonceArgs, reply *NonceReply) error {
	_ = "STUB: not implemented"
	return nil
}

type BalanceArgs struct {
	Address ids.ShortID `json:"address"`
	AssetID ids.ID      `json:"assetID"`
}

type BalanceReply struct {
	Balance uint64 `json:"balance"`
}

func (s *server) Balance(_ *http.Request, args *BalanceArgs, reply *BalanceReply) error {
	_ = "STUB: not implemented"
	return nil
}

type LoanArgs struct {
	ChainID ids.ID `json:"chainID"`
}

type LoanReply struct {
	Amount uint64 `json:"amount"`
}

func (s *server) Loan(_ *http.Request, args *LoanArgs, reply *LoanReply) error {
	_ = "STUB: not implemented"
	return nil
}

type IssueTxArgs struct {
	Tx []byte `json:"tx"`
}

type IssueTxReply struct {
	TxID ids.ID `json:"txID"`
}

func (s *server) IssueTx(r *http.Request, args *IssueTxArgs, reply *IssueTxReply) error {
	_ = "STUB: not implemented"
	return nil
}

type LastAcceptedReply struct {
	BlockID    ids.ID `json:"blockID"`
	BlockBytes []byte `json:"blockBytes"`
}

func (s *server) LastAccepted(_ *http.Request, _ *struct{}, reply *LastAcceptedReply) error {
	_ = "STUB: not implemented"
	return nil
}

type BlockArgs struct {
	BlockID ids.ID `json:"blockID"`
}

type BlockReply struct {
	BlockBytes []byte `json:"blockBytes"`
}

func (s *server) Block(_ *http.Request, args *BlockArgs, reply *BlockReply) error {
	_ = "STUB: not implemented"
	return nil
}

type MessageArgs struct {
	TxID ids.ID `json:"txID"`
}

type MessageReply struct {
	Message   *warp.UnsignedMessage `json:"message"`
	Signature []byte                `json:"signature"`
}

func (s *server) Message(_ *http.Request, args *MessageArgs, reply *MessageReply) error {
	_ = "STUB: not implemented"
	return nil
}
