// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
)

// expectedBlockHeight is the block height that activates the proposerVM fork.
// We issue 2 txs (one per block) to reach block height 2.
const expectedBlockHeight = 2

// IssueTxsToActivateProposerVMFork issues transactions at the current
// timestamp, which should be after the ProposerVM activation time (aka
// ApricotPhase4). This should generate a PostForkBlock because its parent block
// (genesis) has a timestamp (0) that is greater than or equal to the fork
// activation time of 0. Therefore, subsequent blocks should be built with
// BuildBlockWithContext.
func IssueTxsToActivateProposerVMFork(
	ctx context.Context, chainID *big.Int, fundedKey *ecdsa.PrivateKey,
	client *ethclient.Client,
) error {
	_ = "STUB: not implemented"
	return nil
}

// should be pretty generous for c-chain and subnets

// Send exactly 2 transactions, waiting for each to be included in a block

// Wait for this transaction to be included in a block
