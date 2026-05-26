// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ava-labs/libevm/ethclient"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/utils/logging"
)

type Test interface {
	Run(tc tests.TestContext, wallet *Wallet)
}

type Worker struct {
	PrivKey *ecdsa.PrivateKey
	Nonce   uint64
	Client  *ethclient.Client
}

type LoadGenerator struct {
	wallets []*Wallet
	test    Test
}

func NewLoadGenerator(
	workers []Worker,
	chainID *big.Int,
	metricsNamespace string,
	registry *prometheus.Registry,
	test Test,
) (LoadGenerator, error) {
	_ = "STUB: not implemented"
	return *new(LoadGenerator), nil
}

func (l LoadGenerator) Run(
	ctx context.Context,
	log logging.Logger,
	loadTimeout time.Duration,
	testTimeout time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

// execTestWithRecovery ensures assertion-related panics encountered during test execution are recovered
// and that deferred cleanups are always executed before returning.
func execTestWithRecovery(ctx context.Context, log logging.Logger, test Test, wallet *Wallet, testTimeout time.Duration) {
	_ = "STUB: not implemented"
	return
}
