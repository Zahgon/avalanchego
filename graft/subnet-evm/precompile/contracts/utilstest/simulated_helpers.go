// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utilstest

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/eth/ethconfig"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/node"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"

	sim "github.com/ava-labs/avalanchego/graft/subnet-evm/ethclient/simulated"
)

// NewAuth creates a new transactor with the given private key and chain ID.
func NewAuth(t *testing.T, key *ecdsa.PrivateKey, chainID *big.Int) *bind.TransactOpts {
	_ = "STUB: not implemented"
	return nil
}

// NewBackendWithPrecompile creates a simulated backend with the given precompile enabled
// at genesis and funds the specified addresses with 1 ETH each. Additional options can be passed
// to configure the backend.
func NewBackendWithPrecompile(
	t *testing.T,
	precompileCfg precompileconfig.Config,
	fundedAddrs []common.Address,
	opts ...func(*node.Config, *ethconfig.Config),
) *sim.Backend {
	_ = "STUB: not implemented"
	return nil
}

// WaitReceipt commits the simulated backend and waits for the transaction receipt.
func WaitReceipt(t *testing.T, b *sim.Backend, tx *types.Transaction) *types.Receipt {
	_ = "STUB: not implemented"
	return nil
}

// WaitReceiptSuccessful commits the backend, waits for the receipt, and asserts success.
func WaitReceiptSuccessful(t *testing.T, b *sim.Backend, tx *types.Transaction) *types.Receipt {
	_ = "STUB: not implemented"
	return nil
}
