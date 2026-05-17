// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/libevm/options"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database/memdb"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx/txtest"
	"github.com/ava-labs/avalanchego/vms/saevm/saetest"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	snowcommon "github.com/ava-labs/avalanchego/snow/engine/common"
	saeparams "github.com/ava-labs/avalanchego/vms/saevm/params"
)

// nAVAXToAAVAX is the scaling factor between the smallest denomination on the
// X/P-Chain (1 nAVAX) and the C-Chain (1 aAVAX).
const nAVAXToAAVAX = 1_000_000_000

func TestMain(m *testing.M) {
	evm.RegisterAllLibEVMExtras()
	goleak.VerifyTestMain(m, goleak.IgnoreCurrent())
}

// SUT is the system under test for the cchain [VM]. It bundles the [VM]
// itself and an HTTP [Client] connected to an in-process [httptest.Server].
type SUT struct {
	*VM
	*Client

	snowCtx *snow.Context
	memory  *atomic.Memory
}

type (
	sutConfig struct {
		// genesis is the C-Chain genesis. Tests configure it via
		// [options.Func]; see [TestExport] for an example.
		genesis core.Genesis
	}
	sutOption = options.Option[sutConfig]
)

// newSUT initializes a cchain [VM] with the configured genesis, transitions it
// to [snow.NormalOp], and mounts its HTTP handlers behind a local
// [httptest.Server] at the paths [NewClient] expects.
func newSUT(tb testing.TB, opts ...sutOption) *SUT {
	tb.Helper()

	var (
		vm  = &VM{}
		ctx = tb.Context()
		db  = memdb.New()
		cfg = options.ApplyTo(&sutConfig{
			genesis: core.Genesis{
				Config:     saetest.ChainConfig(),
				Alloc:      types.GenesisAlloc{},
				Timestamp:  saeparams.TauSeconds,
				Difficulty: big.NewInt(0), // irrelevant but required to marshal
			},
		}, opts...)
	)

	// The VM and shared memory MUST share an underlying database so that
	// [atomic.SharedMemory.Apply] writes to the VM DB.
	memory := atomic.NewMemory(prefixdb.New([]byte("sharedmemory"), db))
	snowCtx := snowtest.Context(tb, snowtest.CChainID)
	snowCtx.SharedMemory = memory.NewSharedMemory(snowtest.CChainID)
	snowCtx.Log = saetest.NewTBLogger(tb, logging.Debug)

	chainDB := prefixdb.New([]byte("chain"), db)

	genesisBytes, err := json.Marshal(cfg.genesis)
	require.NoErrorf(tb, err, "json.Marshal(%T)", cfg.genesis)

	appSender := &enginetest.Sender{
		SendAppGossipF: func(context.Context, snowcommon.SendConfig, []byte) error {
			return nil
		},
	}

	require.NoErrorf(tb, vm.Initialize(
		ctx,
		snowCtx,
		chainDB,
		genesisBytes,
		nil, // upgradeBytes
		nil, // configBytes
		nil, // fxs
		appSender,
	), "%T.Initialize()", vm)
	tb.Cleanup(func() {
		require.NoErrorf(tb, vm.Shutdown(context.WithoutCancel(tb.Context())), "%T.Shutdown()", vm)
	})
	require.NoErrorf(tb, vm.SetState(ctx, snow.NormalOp), "%T.SetState(%s)", vm, snow.NormalOp)

	handlers, err := vm.CreateHandlers(ctx)
	require.NoErrorf(tb, err, "%T.CreateHandlers()", vm)

	mux := http.NewServeMux()
	for path, h := range handlers {
		mux.Handle(avaxHTTPPrefix+path, h)
	}
	server := httptest.NewServer(mux)
	tb.Cleanup(server.Close)

	return &SUT{
		VM:      vm,
		Client:  NewClient(server.URL),
		snowCtx: snowCtx,
		memory:  memory,
	}
}

// assertUTXOsExist fails tb unless shared memory between peerChainID and the
// C-Chain contains each of the expected UTXOs.
func (sut *SUT) assertUTXOsExist(tb testing.TB, peerChainID ids.ID, expected ...*avax.UTXO) {
	tb.Helper()

	keys := make([][]byte, len(expected))
	for i, u := range expected {
		inputID := u.InputID()
		keys[i] = inputID[:]
	}
	peerMemory := sut.memory.NewSharedMemory(peerChainID)
	raw, err := peerMemory.Get(snowtest.CChainID, keys)
	require.NoErrorf(tb, err, "%T.Get()", peerMemory)

	got := make([]*avax.UTXO, len(raw))
	for i, b := range raw {
		got[i], err = tx.ParseUTXO(b)
		require.NoErrorf(tb, err, "tx.ParseUTXO()")
	}
	if diff := cmp.Diff(expected, got, utxoCmpOpts); diff != "" {
		tb.Errorf("UTXOs in shared memory with %s (-want +got):\n%s", peerChainID, diff)
	}
}

// addUTXOs seeds shared memory between peerChainID and the C-Chain with the
// given UTXOs.
func (sut *SUT) addUTXOs(tb testing.TB, peerChainID ids.ID, utxos ...*avax.UTXO) {
	tb.Helper()

	elems := make([]*atomic.Element, len(utxos))
	for i, utxo := range utxos {
		inputID := utxo.InputID()
		e := &atomic.Element{
			Key:   inputID[:],
			Value: txtest.MustMarshalUTXO(tb, utxo),
		}
		if o, ok := utxo.Out.(avax.Addressable); ok {
			e.Traits = o.Addresses()
		}
		elems[i] = e
	}
	peerMemory := sut.memory.NewSharedMemory(peerChainID)
	err := peerMemory.Apply(map[ids.ID]*atomic.Requests{
		snowtest.CChainID: {PutRequests: elems},
	})
	require.NoErrorf(tb, err, "%T.Apply()", peerMemory)
}

// balance returns the balance of addr at the last-executed state.
func (sut *SUT) balance(tb testing.TB, addr common.Address) uint256.Int {
	tb.Helper()

	state, err := sut.LastExecutedState()
	require.NoErrorf(tb, err, "%T.LastExecutedState()", sut.VM)
	return *state.GetBalance(addr)
}

// assertBalanceChange asserts that addr's balance equals before plus
// nAVAXDelta nAVAX (scaled to aAVAX). The delta may be negative.
func (sut *SUT) assertBalanceChange(tb testing.TB, addr common.Address, before uint256.Int, nAVAXDelta int64) {
	tb.Helper()

	delta := new(big.Int).Mul(big.NewInt(nAVAXDelta), big.NewInt(nAVAXToAAVAX))
	want, overflow := uint256.FromBig(new(big.Int).Add(before.ToBig(), delta))
	require.Falsef(tb, overflow, "want balance overflows uint256")
	require.Equalf(tb, *want, sut.balance(tb, addr), "balance of %s after %+d nAVAX", addr, nAVAXDelta)
}

// issueAndExecute submits t through the HTTP [Client] and drives the consensus
// loop to produce, accept, and execute the next block, which is returned.
func (sut *SUT) issueAndExecute(tb testing.TB, t *tx.Tx) *blocks.Block {
	tb.Helper()

	require.NoErrorf(tb, sut.IssueTx(tb.Context(), t), "%T.IssueTx()", sut.Client)
	return sut.runConsensusLoop(tb)
}

// assertTxAccepted asserts that [Client.GetAtomicTx] returns the given tx at
// the given block height.
func (sut *SUT) assertTxAccepted(tb testing.TB, expected *tx.Tx, height uint64) {
	tb.Helper()

	got, gotHeight, err := sut.GetAtomicTx(tb.Context(), expected.ID())
	require.NoErrorf(tb, err, "%T.GetAtomicTx()", sut.Client)
	if diff := cmp.Diff(expected, got, txtest.CmpOpt()); diff != "" {
		tb.Errorf("%T.GetAtomicTx() (-want +got):\n%s", sut.Client, diff)
	}
	require.Equalf(tb, height, gotHeight, "%T.GetAtomicTx() block height", sut.Client)
}

// runConsensusLoop builds a block on top of the last-accepted block, drives it
// through verify+accept, and blocks until it has been executed.
func (sut *SUT) runConsensusLoop(tb testing.TB) *blocks.Block {
	tb.Helper()

	ctx := tb.Context()
	lastAcceptedID, err := sut.LastAccepted(ctx)
	require.NoErrorf(tb, err, "%T.LastAccepted()", sut.VM)

	// TODO(StephenButtolph): When implementing Warp, we will need to provide
	// meaningful block contexts.
	var blockCtx *block.Context
	require.NoErrorf(tb, sut.SetPreference(ctx, lastAcceptedID, blockCtx), "%T.SetPreference()", sut.VM)

	blk, err := sut.BuildBlock(ctx, blockCtx)
	require.NoErrorf(tb, err, "%T.BuildBlock()", sut.VM)
	require.NoErrorf(tb, sut.VerifyBlock(ctx, blockCtx, blk), "%T.VerifyBlock()", sut.VM)
	require.NoErrorf(tb, sut.AcceptBlock(ctx, blk), "%T.AcceptBlock()", sut.VM)
	require.NoErrorf(tb, blk.WaitUntilExecuted(ctx), "%T.WaitUntilExecuted()", blk)
	return blk
}

// TestExport exercises the cchain VM end-to-end with an Export tx: it builds
// and signs the tx, issues it through the HTTP [Client], drives a block
// through build/verify/accept, waits for execution to complete, and verifies
// both that the tx is reported as accepted and that the sender's C-Chain
// balance dropped by exactly the consumed amount.
func TestExport(t *testing.T) {
	sk := txtest.NewKey(t)
	sender := sk.EthAddress()
	sut := newSUT(t, options.Func[sutConfig](func(c *sutConfig) {
		c.genesis.Alloc = saetest.MaxAllocFor(sender)
	}))

	const (
		// inputAmount is the nAVAX consumed from the sender's C-Chain balance.
		// outputAmount is the nAVAX that ends up as a UTXO on the destination
		// chain. The difference (inputAmount - outputAmount) is burned as the
		// gas fee.
		inputAmount  = 100
		outputAmount = 50
	)
	avaxAssetID := sut.snowCtx.AVAXAssetID
	export := &tx.Export{
		NetworkID:        sut.snowCtx.NetworkID,
		BlockchainID:     sut.snowCtx.ChainID,
		DestinationChain: sut.snowCtx.XChainID,
		Ins: []tx.Input{{
			Address: sender,
			Amount:  inputAmount,
			AssetID: avaxAssetID,
			Nonce:   0,
		}},
		ExportedOutputs: []*avax.TransferableOutput{{
			Asset: avax.Asset{ID: avaxAssetID},
			Out: &secp256k1fx.TransferOutput{
				Amt: outputAmount,
				OutputOwners: secp256k1fx.OutputOwners{
					Threshold: 1,
					Addrs:     []ids.ShortID{sk.Address()},
				},
			},
		}},
	}
	crossTx := &tx.Tx{
		Unsigned: export,
		Creds: []tx.Credential{
			&secp256k1fx.Credential{
				Sigs: []txtest.Signature{txtest.Sign(t, export, sk)},
			},
		},
	}

	initialBalance := sut.balance(t, sender)
	blk := sut.issueAndExecute(t, crossTx)
	sut.assertTxAccepted(t, crossTx, blk.NumberU64())
	sut.assertBalanceChange(t, sender, initialBalance, -inputAmount)

	// The exported outputs must land in the destination chain's shared memory.
	wantUTXOs := txtest.ExportedUTXOs(crossTx.ID(), export)
	sut.assertUTXOsExist(t, sut.snowCtx.XChainID, wantUTXOs...)
}

// utxoCmpOpts compares [avax.UTXO] values, ignoring the unexported cached id
// on [avax.UTXOID] and equating nil and empty slices.
var utxoCmpOpts = cmp.Options{
	cmpopts.IgnoreUnexported(avax.UTXOID{}, secp256k1fx.OutputOwners{}),
	cmpopts.EquateEmpty(),
}

// TestImport exercises the cchain VM end-to-end with an Import tx: it seeds a
// UTXO into the X->C shared memory, builds and signs the Import tx, issues it
// through the HTTP [Client], drives a block through build/verify/accept,
// waits for execution to complete, and verifies both that the tx is reported
// as accepted and that the recipient's C-Chain balance increased by exactly
// the minted amount.
func TestImport(t *testing.T) {
	sk := txtest.NewKey(t)
	sut := newSUT(t)
	avaxAssetID := sut.snowCtx.AVAXAssetID

	const (
		// utxoAmount is the nAVAX in the seeded source UTXO.
		// outputAmount is the nAVAX credited to the recipient on the C-Chain.
		// The difference (utxoAmount - outputAmount) is burned as the gas fee.
		utxoAmount   = 100
		outputAmount = 50
	)

	// Seed an X-Chain UTXO controlled by sk and destined for the C-Chain.
	utxoID := avax.UTXOID{TxID: ids.GenerateTestID()}
	sut.addUTXOs(t, snowtest.XChainID, &avax.UTXO{
		UTXOID: utxoID,
		Asset:  avax.Asset{ID: avaxAssetID},
		Out: &secp256k1fx.TransferOutput{
			Amt: utxoAmount,
			OutputOwners: secp256k1fx.OutputOwners{
				Threshold: 1,
				Addrs:     []ids.ShortID{sk.Address()},
			},
		},
	})

	// Use a fresh recipient address whose balance starts at 0 so we can
	// trivially verify the mint.
	recipient := common.Address{0xde, 0xad, 0xbe, 0xef}
	imp := &tx.Import{
		NetworkID:    sut.snowCtx.NetworkID,
		BlockchainID: sut.snowCtx.ChainID,
		SourceChain:  sut.snowCtx.XChainID,
		ImportedInputs: []*avax.TransferableInput{{
			UTXOID: utxoID,
			Asset:  avax.Asset{ID: avaxAssetID},
			In: &secp256k1fx.TransferInput{
				Amt:   utxoAmount,
				Input: secp256k1fx.Input{SigIndices: []uint32{0}},
			},
		}},
		Outs: []tx.Output{{
			Address: recipient,
			Amount:  outputAmount,
			AssetID: avaxAssetID,
		}},
	}
	crossTx := &tx.Tx{
		Unsigned: imp,
		Creds: []tx.Credential{
			&secp256k1fx.Credential{
				Sigs: []txtest.Signature{txtest.Sign(t, imp, sk)},
			},
		},
	}

	initialBalance := sut.balance(t, recipient)
	blk := sut.issueAndExecute(t, crossTx)

	sut.assertTxAccepted(t, crossTx, blk.NumberU64())

	sut.assertBalanceChange(t, recipient, initialBalance, outputAmount)
}
