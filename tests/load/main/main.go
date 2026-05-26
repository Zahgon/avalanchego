// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"flag"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/ava-labs/libevm/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/tests/load"
	"github.com/ava-labs/avalanchego/tests/load/contracts"
)

const (
	blockchainID     = "C"
	metricsNamespace = "load"
	pollFrequency    = time.Millisecond
	testTimeout      = time.Minute
	defaultNodeCount = 5
)

var (
	flagVars *e2e.FlagVars

	loadTimeoutArg     time.Duration
	firewoodEnabledArg bool
	numWorkersArg      int
)

func init() {
	flagVars = e2e.RegisterFlags(
		e2e.WithDefaultNodeCount(defaultNodeCount),
	)

	flag.DurationVar(
		&loadTimeoutArg,
		"load-timeout",
		0,
		"the duration that the load test should run for",
	)
	flag.BoolVar(
		&firewoodEnabledArg,
		"firewood",
		false,
		"whether to use Firewood in Coreth",
	)
	flag.IntVar(
		&numWorkersArg,
		"num-workers",
		defaultNodeCount,
		"the number of workers to use for the load test",
	)

	flag.Parse()
}

func main() {
	log := tests.NewDefaultLogger("")
	tc := tests.NewTestContext(log)
	defer tc.RecoverAndExit()

	require := require.New(tc)

	numNodes, err := flagVars.NodeCount()
	require.NoError(err, "failed to get node count")

	nodes := tmpnet.NewNodesOrPanic(numNodes)

	keys, err := tmpnet.NewPrivateKeys(numWorkersArg)
	require.NoError(err)

	primaryChainConfigs := tmpnet.DefaultChainConfigs()
	if firewoodEnabledArg {
		primaryChainConfigs = newPrimaryChainConfigsWithFirewood()
	}
	network := &tmpnet.Network{
		Nodes:               nodes,
		PreFundedKeys:       keys,
		PrimaryChainConfigs: primaryChainConfigs,
	}

	e2e.NewTestEnvironment(tc, flagVars, network)

	ctx := tests.DefaultNotifyContext(0, tc.DeferCleanup)
	wsURIs, err := tmpnet.GetNodeWebsocketURIs(network.Nodes, blockchainID)
	require.NoError(err)

	registry := prometheus.NewRegistry()
	metricsServer, err := tests.NewPrometheusServer(registry)
	require.NoError(err)
	tc.DeferCleanup(func() {
		require.NoError(metricsServer.Stop())
	})

	monitoringConfigFilePath, err := tmpnet.WritePrometheusSDConfig("load-test", tmpnet.SDConfig{
		Targets: []string{metricsServer.Address()},
		Labels:  network.GetMonitoringLabels(),
	}, false)
	require.NoError(err, "failed to generate monitoring config file")

	tc.DeferCleanup(func() {
		require.NoError(
			os.Remove(monitoringConfigFilePath),
			"failed †o remove monitoring config file",
		)
	})

	workers := make([]load.Worker, len(keys))
	for i := range len(keys) {
		wsURI := wsURIs[i%len(wsURIs)]
		client, err := ethclient.Dial(wsURI)
		require.NoError(err)

		workers[i] = load.Worker{
			PrivKey: keys[i].ToECDSA(),
			Client:  client,
		}
	}

	chainID, err := workers[0].Client.ChainID(ctx)
	require.NoError(err)

	tokenContract, err := newTokenContract(ctx, chainID, &workers[0], workers[1:])
	require.NoError(err)

	randomTest, err := load.NewRandomTest(
		ctx,
		chainID,
		&workers[0],
		rand.NewSource(time.Now().UnixMilli()),
		tokenContract,
	)
	require.NoError(err)

	generator, err := load.NewLoadGenerator(
		workers,
		chainID,
		metricsNamespace,
		registry,
		randomTest,
	)
	require.NoError(err)

	generator.Run(ctx, log, loadTimeoutArg, testTimeout)
}

// newTokenContract deploys an instance of an ERC20 token and distributes the
// token supply evenly between the deployer and the recipients
func newTokenContract(
	ctx context.Context,
	chainID *big.Int,
	deployer *load.Worker,
	recipients []load.Worker,
) (*contracts.ERC20, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assumes that token has 18 decimals

// newPrimaryChainConfigsWithFirewood extends the default primary chain configs
// by enabling Firewood on the C-Chain.
func newPrimaryChainConfigsWithFirewood() map[string]tmpnet.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}

// firewoodConfig represents the minimum configuration required to enable
// Firewood in Coreth.
//
// Ref: https://github.com/ava-labs/avalanchego/graft/coreth/issues/1180
