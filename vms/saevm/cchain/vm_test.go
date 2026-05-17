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
func (s *SUT) assertUTXOsExist(tb testing.TB, peerChainID ids.ID, want ...*avax.UTXO) {
	tb.Helper()

	keys := make([][]byte, len(want))
	for i, u := range want {
		inputID := u.InputID()
		keys[i] = inputID[:]
	}
	peerMemory := s.memory.NewSharedMemory(peerChainID)
	raw, err := peerMemory.Get(snowtest.CChainID, keys)
	require.NoErrorf(tb, err, "%T.Get()", peerMemory)

	got := make([]*avax.UTXO, len(raw))
	for i, b := range raw {
		got[i] = txtest.MustParseUTXO(tb, b)
	}
	opts := cmp.Options{
		cmpopts.IgnoreUnexported(avax.UTXOID{}, secp256k1fx.OutputOwners{}),
		cmpopts.EquateEmpty(),
	}
	if diff := cmp.Diff(want, got, opts); diff != "" {
		tb.Errorf("UTXOs in shared memory with %s (-want +got):\n%s", peerChainID, diff)
	}
}

// addUTXOs seeds shared memory between peerChainID and the C-Chain with the
// given UTXOs.
func (s *SUT) addUTXOs(tb testing.TB, peerChainID ids.ID, utxos ...*avax.UTXO) {
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
	peerMemory := s.memory.NewSharedMemory(peerChainID)
	err := peerMemory.Apply(map[ids.ID]*atomic.Requests{
		snowtest.CChainID: {PutRequests: elems},
	})
	require.NoErrorf(tb, err, "%T.Apply()", peerMemory)
}

// balance returns the balance of addr at the last-executed state.
func (s *SUT) balance(tb testing.TB, addr common.Address) uint256.Int {
	tb.Helper()

	state, err := s.LastExecutedState()
	require.NoErrorf(tb, err, "%T.LastExecutedState()", s.VM)
	return *state.GetBalance(addr)
}

// assertBalance asserts that addr's balance at the last-executed state equals
// want.
func (s *SUT) assertBalance(tb testing.TB, addr common.Address, want uint256.Int) {
	tb.Helper()
	require.Equalf(tb, want, s.balance(tb, addr), "balance of %s", addr)
}

// issueAndExecute submits t through the HTTP [Client] and drives the consensus
// loop to produce, accept, and execute the next block, which is returned.
func (s *SUT) issueAndExecute(tb testing.TB, t *tx.Tx) *blocks.Block {
	tb.Helper()

	require.NoErrorf(tb, s.IssueTx(tb.Context(), t), "%T.IssueTx()", s.Client)
	return s.runConsensusLoop(tb)
}

// assertTxAccepted asserts that [Client.GetAtomicTx] returns the given tx at
// the given block height.
func (s *SUT) assertTxAccepted(tb testing.TB, want *tx.Tx, wantHeight uint64) {
	tb.Helper()

	got, gotHeight, err := s.GetAtomicTx(tb.Context(), want.ID())
	require.NoErrorf(tb, err, "%T.GetAtomicTx()", s.Client)
	if diff := cmp.Diff(want, got, txtest.CmpOpt()); diff != "" {
		tb.Errorf("%T.GetAtomicTx() (-want +got):\n%s", s.Client, diff)
	}
	require.Equalf(tb, wantHeight, gotHeight, "%T.GetAtomicTx() block height", s.Client)
}

// runConsensusLoop builds a block on top of the last-accepted block, drives it
// through verify+accept, and blocks until it has been executed.
func (s *SUT) runConsensusLoop(tb testing.TB) *blocks.Block {
	tb.Helper()

	ctx := tb.Context()
	lastAcceptedID, err := s.LastAccepted(ctx)
	require.NoErrorf(tb, err, "%T.LastAccepted()", s.VM)

	// TODO(StephenButtolph): When implementing Warp, we will need to provide
	// meaningful block contexts.
	var blockCtx *block.Context
	require.NoErrorf(tb, s.SetPreference(ctx, lastAcceptedID, blockCtx), "%T.SetPreference()", s.VM)

	e, err := s.WaitForEvent(ctx)
	require.NoErrorf(tb, err, "%T.WaitForEvent()", s.VM)
	require.Equalf(tb, snowcommon.PendingTxs, e, "%T.WaitForEvent() event", s.VM)

	blk, err := s.BuildBlock(ctx, blockCtx)
	require.NoErrorf(tb, err, "%T.BuildBlock()", s.VM)
	require.NoErrorf(tb, s.VerifyBlock(ctx, blockCtx, blk), "%T.VerifyBlock()", s.VM)
	require.NoErrorf(tb, s.AcceptBlock(ctx, blk), "%T.AcceptBlock()", s.VM)
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
		exportedAmount = 50
		txFee          = 50
	)

	wallet := txtest.NewWallet(sk, sut.snowCtx, sut.memory)
	signedExport, export := wallet.NewExportTx(t, sut.snowCtx.XChainID, []*secp256k1fx.TransferOutput{{
		Amt: exportedAmount,
		OutputOwners: secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{sk.Address()},
		},
	}}, txFee)

	initialBalance := sut.balance(t, sender)
	blk := sut.issueAndExecute(t, signedExport)
	sut.assertTxAccepted(t, signedExport, blk.NumberU64())
	const amountBurned = exportedAmount + txFee
	sut.assertBalance(t, sender, txtest.AddNAVAX(initialBalance, -amountBurned))
	sut.assertUTXOsExist(t, sut.snowCtx.XChainID, txtest.ExportedUTXOs(signedExport.ID(), export)...)
}

// TestImport exercises the cchain VM end-to-end with an Import tx: it seeds a
// UTXO into the X->C shared memory, builds and signs the Import tx, issues it
// through the HTTP [Client], drives a block through build/verify/accept,
// waits for execution to complete, and verifies both that the tx is reported
// as accepted and that the recipient's C-Chain balance increased by exactly
// the minted amount.
func TestImport(t *testing.T) {
	sut := newSUT(t)

	const utxoAmount = 100
	sk := txtest.NewKey(t)

	// Seed an X-Chain UTXO controlled by sk and destined for the C-Chain.
	sut.addUTXOs(t, snowtest.XChainID, &avax.UTXO{
		UTXOID: avax.UTXOID{TxID: ids.GenerateTestID()},
		Asset:  avax.Asset{ID: sut.snowCtx.AVAXAssetID},
		Out: &secp256k1fx.TransferOutput{
			Amt: utxoAmount,
			OutputOwners: secp256k1fx.OutputOwners{
				Threshold: 1,
				Addrs:     []ids.ShortID{sk.Address()},
			},
		},
	})

	const txFee = 50

	// Use a fresh recipient address whose balance starts at 0 so we can
	// trivially verify the mint.
	recipient := common.Address{0xde, 0xad, 0xbe, 0xef}
	wallet := txtest.NewWallet(sk, sut.snowCtx, sut.memory)
	signedImport, _ := wallet.NewImportTx(t, sut.snowCtx.XChainID, recipient, txFee)

	blk := sut.issueAndExecute(t, signedImport)
	sut.assertTxAccepted(t, signedImport, blk.NumberU64())
	const amountMinted = utxoAmount - txFee
	sut.assertBalance(t, recipient, tx.ScaleAVAX(amountMinted))
}
