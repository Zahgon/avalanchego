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
	"github.com/ava-labs/avalanchego/tests/fixture/subnet"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/cmd/issue/status"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/cmd/issue/transfer"
)

const (
	NumKeys         = 5
	PollingInterval = 50 * time.Millisecond
)

func main() {
	// TODO(marun) Support choosing the log format
	tc := antithesis.NewInstrumentedTestContext(tests.NewDefaultLogger(""))
	defer tc.RecoverAndExit()
	require := require.New(tc)

	c := antithesis.NewConfigWithSubnets(
		tc,
		&tmpnet.Network{
			Owner: "antithesis-xsvm",
		},
		func(nodes ...*tmpnet.Node) []*tmpnet.Subnet {
			return []*tmpnet.Subnet{
				subnet.NewXSVMOrPanic("xsvm", genesis.VMRQKey, nodes...),
			}
		},
	)
	ctx := tests.DefaultNotifyContext(c.Duration, tc.DeferCleanup)
	// Ensure contexts sourced from the test context use the notify context as their parent
	tc.SetDefaultContextParent(ctx)

	require.Len(c.ChainIDs, 1)
	tc.Log().Debug("raw chain ID",
		zap.String("chainID", c.ChainIDs[0]),
	)
	chainID, err := ids.FromString(c.ChainIDs[0])
	require.NoError(err, "failed to parse chainID")
	tc.Log().Info("node and chain configuration",
		zap.Stringer("chainID", chainID),
		zap.Strings("uris", c.URIs),
	)

	genesisWorkload := &workload{
		id:      0,
		log:     tests.NewDefaultLogger(fmt.Sprintf("worker %d", 0)),
		chainID: chainID,
		key:     genesis.VMRQKey,
		addrs:   set.Of(genesis.VMRQKey.Address()),
		uris:    c.URIs,
	}

	workloads := make([]*workload, NumKeys)
	workloads[0] = genesisWorkload

	initialAmount := 100 * units.KiloAvax
	for i := 1; i < NumKeys; i++ {
		key, err := secp256k1.NewPrivateKey()
		require.NoError(err, "failed to generate key")

		var (
			addr          = key.Address()
			baseStartTime = time.Now()
		)
		transferTxStatus, err := transfer.Transfer(
			ctx,
			&transfer.Config{
				URI:        c.URIs[0],
				ChainID:    chainID,
				AssetID:    chainID,
				Amount:     initialAmount,
				To:         addr,
				PrivateKey: genesisWorkload.key,
			},
		)
		require.NoError(err, "failed to issue initial funding transfer")
		tc.Log().Info("issued initial funding transfer",
			zap.Stringer("txID", transferTxStatus.TxID),
			zap.Duration("duration", time.Since(baseStartTime)),
		)

		genesisWorkload.confirmTransferTx(ctx, transferTxStatus)

		workloads[i] = &workload{
			id:      i,
			log:     tests.NewDefaultLogger(fmt.Sprintf("worker %d", i)),
			chainID: chainID,
			key:     key,
			addrs:   set.Of(addr),
			uris:    c.URIs,
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
	id      int
	log     logging.Logger
	chainID ids.ID
	key     *secp256k1.PrivateKey
	addrs   set.Set[ids.ShortID]
	uris    []string
}

func (w *workload) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Any assertion failure from this test context will result in process exit due to the
// panic being rethrown. This ensures that failures in test setup are fatal.

func (w *workload) confirmTransferTx(ctx context.Context, tx *status.TxIssuance) {
	_ = "STUB: not implemented"
	return
}
