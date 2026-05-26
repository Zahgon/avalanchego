// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/antithesishq/antithesis-sdk-go/lifecycle"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/ethclient"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/tests"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/tests/utils"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/antithesis"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/logging"

	ago_tests "github.com/ava-labs/avalanchego/tests"
)

const NumKeys = 5

func main() {
	logger := ago_tests.NewDefaultLogger("")
	tc := antithesis.NewInstrumentedTestContext(logger)
	defer tc.RecoverAndExit()
	require := require.New(tc)

	c := antithesis.NewConfigWithSubnets(
		tc,
		// TODO(marun) Centralize network configuration for all test types
		utils.NewTmpnetNetwork(
			"antithesis-subnet-evm",
			nil,
			tmpnet.FlagsMap{},
		),
		func(nodes ...*tmpnet.Node) []*tmpnet.Subnet {
			return []*tmpnet.Subnet{
				utils.NewTmpnetSubnet("subnet-evm", tests.Genesis, utils.DefaultChainConfig, nodes...),
			}
		},
	)
	ctx := ago_tests.DefaultNotifyContext(c.Duration, tc.DeferCleanup)

	// Ensure contexts sourced from the test context use the notify context as their parent
	tc.SetDefaultContextParent(ctx)

	require.Len(c.ChainIDs, 1)
	logger.Info("Starting testing",
		zap.Strings("chainIDs", c.ChainIDs),
	)
	chainID, err := ids.FromString(c.ChainIDs[0])
	require.NoError(err, "failed to parse chainID")

	genesisClient, err := ethclient.Dial(getChainURI(c.URIs[0], chainID.String()))
	require.NoError(err, "failed to dial chain")
	genesisKey := tmpnet.HardhatKey.ToECDSA()
	genesisWorkload := &workload{
		id:     0,
		log:    ago_tests.NewDefaultLogger(fmt.Sprintf("worker %d", 0)),
		client: genesisClient,
		key:    genesisKey,
	}

	workloads := make([]*workload, NumKeys)
	workloads[0] = genesisWorkload

	initialAmount := uint64(1_000_000_000_000_000)
	for i := 1; i < NumKeys; i++ {
		key, err := crypto.ToECDSA(crypto.Keccak256([]byte{uint8(i)}))
		require.NoError(err, "failed to generate key")

		require.NoError(transferFunds(ctx, genesisClient, genesisKey, crypto.PubkeyToAddress(key.PublicKey), initialAmount, logger))

		client, err := ethclient.Dial(getChainURI(c.URIs[i%len(c.URIs)], chainID.String()))
		require.NoError(err, "failed to dial chain")

		workloads[i] = &workload{
			id:     i,
			log:    ago_tests.NewDefaultLogger(fmt.Sprintf("worker %d", i)),
			client: client,
			key:    key,
		}
	}

	lifecycle.SetupComplete(map[string]any{
		"msg":        "initialized workers",
		"numWorkers": NumKeys,
	})

	for _, w := range workloads[1:] {
		go w.run(ctx)
	}
	genesisWorkload.run(ctx)
}

type workload struct {
	id     int
	client ethclient.Client
	log    logging.Logger
	key    *ecdsa.PrivateKey
}

// newTestContext returns a test context that ensures that log output and assertions are
// associated with this worker.
func (w *workload) newTestContext(ctx context.Context) *ago_tests.SimpleTestContext {
	_ = "STUB: not implemented"
	return nil
}

func (w *workload) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Any assertion failure from this test context will result in process exit due to the
// panic being rethrown. This ensures that failures in test setup are fatal.

func (w *workload) executeTest(ctx context.Context) {
	_ = "STUB: not implemented"
	// TODO(marun) What should this value be?
	return
}

// TODO(marun) Exercise a wider variety of transactions

// Log the error and continue since the problem may be
// transient. require.NoError is only for errors that should stop
// execution.

func getChainURI(nodeURI string, blockchainID string) string { _ = "STUB: not implemented"; return "" }

func transferFunds(ctx context.Context, client ethclient.Client, key *ecdsa.PrivateKey, recipientAddress common.Address, txAmount uint64, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
