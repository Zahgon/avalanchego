// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"math/big"

	"github.com/ava-labs/libevm/accounts"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"

	ethereum "github.com/ava-labs/libevm"
)

func (b *backend) ChainDb() ethdb.Database {
	_ = "STUB: not implemented" //nolint:staticcheck // this name required by ethapi.Backend interface
	return *new(ethdb.Database)
}

func (b *backend) RPCTxFeeCap() float64 { _ = "STUB: not implemented"; return 0 }

func (*backend) UnprotectedAllowed() bool {
	_ = "STUB: not implemented"

	// ExtRPCEnabled reports that external RPC access is enabled. This adds an
	// additional security measure in case we add support for the personal API.
	return false
}

func (*backend) ExtRPCEnabled() bool {
	_ = "STUB: not implemented"

	// Total difficulty does not make sense in snowman consensus, as it is not PoW.
	// Ethereum, post merge (switch to PoS), sets the difficulty of each block to 0
	// (see: https://github.com/ethereum/go-ethereum/blob/be92f5487e67939b8dbbc9675d6c15be76ffd18d/consensus/beacon/consensus.go#L228-L231)
	// and no longer exposes the total difficulty of the chain at all via the API.
	//
	// TODO(JonathanOppenheimer): Once we update libevm, remove GetTd.
	return false
}

func (*backend) GetTd(ctx context.Context, _ common.Hash) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (*backend) SyncProgress() ethereum.SyncProgress {
	_ = "STUB: not implemented"
	// Avalanchego does not expose APIs until after the node has fully synced.
	return *new(ethereum.SyncProgress)
}

func (b *backend) AccountManager() *accounts.Manager { _ = "STUB: not implemented"; return nil }
