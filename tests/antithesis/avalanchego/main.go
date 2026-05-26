// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/antithesishq/antithesis-sdk-go/lifecycle"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/antithesis"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"

	xtxs "github.com/ava-labs/avalanchego/vms/avm/txs"
	ptxs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

const NumKeys = 5

// TODO(marun) Extract the common elements of test execution for reuse across test setups

func main() {
	// TODO(marun) Support choosing the log format
	tc := antithesis.NewInstrumentedTestContext(tests.NewDefaultLogger(""))
	defer tc.RecoverAndExit()
	require := require.New(tc)

	c := antithesis.NewConfig(
		tc,
		&tmpnet.Network{
			Owner: "antithesis-avalanchego",
		},
	)
	ctx := tests.DefaultNotifyContext(c.Duration, tc.DeferCleanup)
	// Ensure contexts sourced from the test context use the notify context as their parent
	tc.SetDefaultContextParent(ctx)

	kc := secp256k1fx.NewKeychain(genesis.EWOQKey)
	walletSyncStartTime := time.Now()
	wallet := e2e.NewWallet(tc, kc, tmpnet.NodeURI{URI: c.URIs[0]})
	tc.Log().Info("synced wallet",
		zap.Duration("duration", time.Since(walletSyncStartTime)),
	)

	genesisWorkload := &workload{
		id:     0,
		log:    tests.NewDefaultLogger(fmt.Sprintf("worker %d", 0)),
		wallet: wallet,
		addrs:  set.Of(genesis.EWOQKey.Address()),
		uris:   c.URIs,
	}

	workloads := make([]*workload, NumKeys)
	workloads[0] = genesisWorkload

	var (
		genesisXWallet  = wallet.X()
		genesisXBuilder = genesisXWallet.Builder()
		genesisXContext = genesisXBuilder.Context()
		avaxAssetID     = genesisXContext.AVAXAssetID
	)
	for i := 1; i < NumKeys; i++ {
		key, err := secp256k1.NewPrivateKey()
		require.NoError(err, "failed to generate key")

		var (
			addr          = key.Address()
			baseStartTime = time.Now()
		)
		baseTx, err := genesisXWallet.IssueBaseTx([]*avax.TransferableOutput{{
			Asset: avax.Asset{
				ID: avaxAssetID,
			},
			Out: &secp256k1fx.TransferOutput{
				Amt: 100 * units.KiloAvax,
				OutputOwners: secp256k1fx.OutputOwners{
					Threshold: 1,
					Addrs: []ids.ShortID{
						addr,
					},
				},
			},
		}})
		require.NoError(err, "failed to issue initial funding X-chain baseTx")
		tc.Log().Info("issued initial funding X-chain baseTx",
			zap.Stringer("txID", baseTx.ID()),
			zap.Duration("duration", time.Since(baseStartTime)),
		)

		require.NoError(genesisWorkload.confirmXChainTx(ctx, baseTx), "failed to confirm initial funding X-chain baseTx")

		uri := c.URIs[i%len(c.URIs)]
		kc := secp256k1fx.NewKeychain(key)
		walletSyncStartTime := time.Now()
		wallet := e2e.NewWallet(tc, kc, tmpnet.NodeURI{URI: uri})
		tc.Log().Info("synced wallet",
			zap.Duration("duration", time.Since(walletSyncStartTime)),
		)

		workloads[i] = &workload{
			id:     i,
			log:    tests.NewDefaultLogger(fmt.Sprintf("worker %d", i)),
			wallet: wallet,
			addrs:  set.Of(addr),
			uris:   c.URIs,
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
	log    logging.Logger
	wallet *primary.Wallet
	addrs  set.Set[ids.ShortID]
	uris   []string
}

// newTestContext returns a test context that ensures that log output and assertions are
// associated with this worker.
func (w *workload) newTestContext(ctx context.Context) *tests.SimpleTestContext {
	_ = "STUB: not implemented"
	return nil
}

func (w *workload) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Any assertion failure from this test context will result in process exit due to the
// panic being rethrown. This ensures that failures in test setup are fatal.

// executeTest executes a test at random.
func (w *workload) executeTest(ctx context.Context) { _ = "STUB: not implemented"; return }

// Panics will be recovered without being rethrown, ensuring that test failures are not fatal.

// Ensure this value matches the number of tests + 1 to offset
// 0-based + 1 for sleep case in the switch statement for flowID

// TODO(marun) Create abstraction for a test that supports a name e.g. `aTest{name: "foo", mytestfunc}`

// TODO(marun) Enable execution of the banff e2e test as part of https://github.com/ava-labs/avalanchego/issues/4049
// w.log.Info("executing banff.TestCustomAssetTransfer")
// addr, _ := w.addrs.Peek()
// banff.TestCustomAssetTransfer(tc, *w.wallet, addr)

func (w *workload) issueXChainBaseTx(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *workload) issueXChainCreateAssetTx(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *workload) issueXChainOperationTx(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *workload) issueXToPTransfer(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *workload) issuePToXTransfer(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *workload) makeOwner() secp256k1fx.OutputOwners {
	_ = "STUB: not implemented"
	return *new(secp256k1fx.OutputOwners)
}

func (w *workload) confirmXChainTx(ctx context.Context, tx *xtxs.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workload) confirmPChainTx(ctx context.Context, tx *ptxs.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workload) verifyXChainTxConsumedUTXOs(ctx context.Context, tx *xtxs.Tx) {
	_ = "STUB: not implemented"
	return
}

func (w *workload) verifyPChainTxConsumedUTXOs(ctx context.Context, tx *ptxs.Tx) {
	_ = "STUB: not implemented"
	return
}
