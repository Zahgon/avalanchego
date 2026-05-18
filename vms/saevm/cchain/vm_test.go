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
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
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
	goleak.VerifyTestMain(m, saetest.GoleakOptions()...)
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

// assertTxAccepted asserts that [Client.GetTx] returns the given tx at
// the given block height.
func (s *SUT) assertTxAccepted(tb testing.TB, want *tx.Tx, wantHeight uint64) {
	tb.Helper()

	got, gotHeight, err := s.GetTx(tb.Context(), want.ID())
	require.NoErrorf(tb, err, "%T.GetTx()", s.Client)
	if diff := cmp.Diff(want, got, txtest.CmpOpt()); diff != "" {
		tb.Errorf("%T.GetTx() (-want +got):\n%s", s.Client, diff)
	}
	require.Equalf(tb, wantHeight, gotHeight, "%T.GetTx() block height", s.Client)
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

// wallet builds and signs cross-chain transactions on behalf of a single key.
// It is the analog of [wallet/chain/c.Wallet] for SAE.
type wallet struct {
	sk      *secp256k1.PrivateKey
	snowCtx *snow.Context
	client  *Client
	nonce   uint64
}

// newWallet returns a [*wallet] backed by sk for the chain described by
// snowCtx. client is queried when building imports to discover spendable
// UTXOs.
func newWallet(sk *secp256k1.PrivateKey, snowCtx *snow.Context, client *Client) *wallet {
	return &wallet{
		sk:      sk,
		snowCtx: snowCtx,
		client:  client,
	}
}

// newExportTx builds and signs an [tx.Export] sending outputs to
// destinationChain. The wallet contributes a single AVAX input from its eth
// address with Amount = sum(outputs.Amt) + fee, using its next nonce.
func (w *wallet) newExportTx(
	tb testing.TB,
	destinationChain ids.ID,
	outputs []*secp256k1fx.TransferOutput,
	fee uint64,
) (*tx.Tx, *tx.Export) {
	tb.Helper()

	avaxAssetID := w.snowCtx.AVAXAssetID
	var exportedAmount uint64
	transferable := make([]*avax.TransferableOutput, len(outputs))
	for i, out := range outputs {
		transferable[i] = &avax.TransferableOutput{
			Asset: avax.Asset{ID: avaxAssetID},
			Out:   out,
		}
		exportedAmount += out.Amt
	}

	export := &tx.Export{
		NetworkID:        w.snowCtx.NetworkID,
		BlockchainID:     w.snowCtx.ChainID,
		DestinationChain: destinationChain,
		Ins: []tx.Input{{
			Address: w.sk.EthAddress(),
			Amount:  exportedAmount + fee,
			AssetID: avaxAssetID,
			Nonce:   w.nonce,
		}},
		ExportedOutputs: transferable,
	}
	w.nonce++

	return w.sign(tb, export, 1), export
}

// getUTXOs paginates [Client.GetUTXOs] and returns every UTXO controlled by
// any of addrs that has been exported to this chain from sourceChain.
func (w *wallet) getUTXOs(tb testing.TB, sourceChain ids.ID, addrs ...ids.ShortID) []*avax.UTXO {
	tb.Helper()

	var (
		sourceStr   = sourceChain.String()
		startAddr   ids.ShortID
		startUTXOID ids.ID
		utxos       []*avax.UTXO
	)
	for {
		const limit = 1024
		page, endAddr, endUTXOID, err := w.client.GetUTXOs(
			tb.Context(),
			addrs,
			sourceStr,
			limit,
			startAddr,
			startUTXOID,
		)
		require.NoErrorf(tb, err, "%T.GetUTXOs()", w.client)
		utxos = append(utxos, page...)
		if len(page) < limit {
			return utxos
		}
		startAddr, startUTXOID = endAddr, endUTXOID
	}
}

// newImportTx builds and signs an [tx.Import] consuming all spendable AVAX
// UTXOs that have been exported to this chain from sourceChain and are owned
// by the wallet, crediting the total imported (minus fee) to `to` on the
// C-Chain.
func (w *wallet) newImportTx(
	tb testing.TB,
	sourceChain ids.ID,
	to common.Address,
	fee uint64,
) (*tx.Tx, *tx.Import) {
	tb.Helper()

	utxos := w.getUTXOs(tb, sourceChain, w.sk.Address())

	var (
		avaxAssetID  = w.snowCtx.AVAXAssetID
		importedAVAX uint64
		inputs       = make([]*avax.TransferableInput, 0, len(utxos))
	)
	for _, utxo := range utxos {
		if utxo.Asset.ID != avaxAssetID {
			continue
		}

		out, ok := utxo.Out.(*secp256k1fx.TransferOutput)
		require.Truef(tb, ok, "unexpected UTXO output type %T", utxo.Out)

		importedAVAX += out.Amt
		inputs = append(inputs, &avax.TransferableInput{
			UTXOID: utxo.UTXOID,
			Asset:  utxo.Asset,
			In: &secp256k1fx.TransferInput{
				Amt: out.Amt,
				Input: secp256k1fx.Input{
					SigIndices: []uint32{0},
				},
			},
		})
	}
	require.Greaterf(tb, importedAVAX, fee, "imported AVAX insufficient to cover fee")

	imp := &tx.Import{
		NetworkID:      w.snowCtx.NetworkID,
		BlockchainID:   w.snowCtx.ChainID,
		SourceChain:    sourceChain,
		ImportedInputs: inputs,
		Outs: []tx.Output{{
			Address: to,
			Amount:  importedAVAX - fee,
			AssetID: avaxAssetID,
		}},
	}
	return w.sign(tb, imp, len(inputs)), imp
}

// sign wraps u in a [tx.Tx] with numCreds copies of a single-sig credential
// over u.
func (w *wallet) sign(tb testing.TB, u tx.Unsigned, numCreds int) *tx.Tx {
	tb.Helper()

	sig := txtest.Sign(tb, u, w.sk)
	creds := make([]tx.Credential, numCreds)
	for i := range creds {
		creds[i] = &secp256k1fx.Credential{Sigs: []txtest.Signature{sig}}
	}
	return &tx.Tx{
		Unsigned: u,
		Creds:    creds,
	}
}

// addNAVAX returns balance + nAVAXDelta nAVAX (scaled to aAVAX). The delta
// may be negative. It panics if the result does not fit in a uint256.
func addNAVAX(balance uint256.Int, nAVAXDelta int64) uint256.Int {
	delta := new(big.Int).Mul(big.NewInt(nAVAXDelta), big.NewInt(tx.X2CRate))
	sum := new(big.Int).Add(balance.ToBig(), delta)
	result, overflow := uint256.FromBig(sum)
	if overflow {
		panic("addNAVAX: result overflows uint256")
	}
	return *result
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

	w := newWallet(sk, sut.snowCtx, sut.Client)
	signedExport, export := w.newExportTx(t, sut.snowCtx.XChainID, []*secp256k1fx.TransferOutput{{
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
	sut.assertBalance(t, sender, addNAVAX(initialBalance, -amountBurned))
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
	w := newWallet(sk, sut.snowCtx, sut.Client)
	signedImport, _ := w.newImportTx(t, sut.snowCtx.XChainID, recipient, txFee)

	blk := sut.issueAndExecute(t, signedImport)
	sut.assertTxAccepted(t, signedImport, blk.NumberU64())
	const amountMinted = utxoAmount - txFee
	sut.assertBalance(t, recipient, tx.ScaleAVAX(amountMinted))
}

// TestIssueTxFailures asserts that [Client.IssueTx] surfaces a meaningful
// error for each category of pool-level rejection.
//
// The cases are ordered to mirror the validation pipeline in [txpool.Add]:
// sanity check → credentials → state nonce → state balance. Each tx is
// built fresh so it triggers exactly one failure mode.
func TestIssueTxFailures(t *testing.T) {
	sk := txtest.NewKey(t)
	sender := sk.EthAddress()
	sut := newSUT(t, options.Func[sutConfig](func(c *sutConfig) {
		c.genesis.Alloc = saetest.MaxAllocFor(sender)
	}))

	const (
		exportedAmount = 50
		txFee          = 50
	)
	outputs := []*secp256k1fx.TransferOutput{{
		Amt: exportedAmount,
		OutputOwners: secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{sk.Address()},
		},
	}}

	tests := []struct {
		name      string
		construct func(*testing.T, *wallet) *tx.Tx
		wantErr   string
	}{
		{
			name: "wrong network ID",
			construct: func(t *testing.T, w *wallet) *tx.Tx {
				_, export := w.newExportTx(t, sut.snowCtx.XChainID, outputs, txFee)
				export.NetworkID++
				return w.sign(t, export, 1)
			},
			wantErr: "wrong network ID",
		},
		{
			name: "credential signed by wrong key",
			construct: func(t *testing.T, w *wallet) *tx.Tx {
				signed, _ := w.newExportTx(t, sut.snowCtx.XChainID, outputs, txFee)
				wrong := txtest.NewKey(t)
				signed.Creds = []tx.Credential{&secp256k1fx.Credential{
					Sigs: []txtest.Signature{txtest.Sign(t, signed.Unsigned, wrong)},
				}}
				return signed
			},
			wantErr: "signature does not match address",
		},
		{
			name: "nonce mismatch",
			construct: func(t *testing.T, w *wallet) *tx.Tx {
				w.nonce = 99
				signed, _ := w.newExportTx(t, sut.snowCtx.XChainID, outputs, txFee)
				return signed
			},
			wantErr: "nonce mismatch",
		},
		{
			name: "insufficient balance",
			construct: func(t *testing.T, _ *wallet) *tx.Tx {
				// A wallet for a fresh key holds zero balance.
				pauper := newWallet(txtest.NewKey(t), sut.snowCtx, sut.Client)
				signed, _ := pauper.newExportTx(t, sut.snowCtx.XChainID, outputs, txFee)
				return signed
			},
			wantErr: "insufficient funds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := newWallet(sk, sut.snowCtx, sut.Client)
			err := sut.IssueTx(t.Context(), tt.construct(t, w))
			require.ErrorContainsf(t, err, tt.wantErr, "%T.IssueTx()", sut.Client)
		})
	}
}

// TestIssueTxAlreadyKnown asserts that resubmitting a tx already in the pool
// returns an error mentioning the duplicate, while leaving the original
// admission intact.
func TestIssueTxAlreadyKnown(t *testing.T) {
	sk := txtest.NewKey(t)
	sender := sk.EthAddress()
	sut := newSUT(t, options.Func[sutConfig](func(c *sutConfig) {
		c.genesis.Alloc = saetest.MaxAllocFor(sender)
	}))

	const (
		exportedAmount = 50
		txFee          = 50
	)
	w := newWallet(sk, sut.snowCtx, sut.Client)
	signed, _ := w.newExportTx(t, sut.snowCtx.XChainID, []*secp256k1fx.TransferOutput{{
		Amt: exportedAmount,
		OutputOwners: secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{sk.Address()},
		},
	}}, txFee)

	require.NoErrorf(t, sut.IssueTx(t.Context(), signed), "%T.IssueTx() first call", sut.Client)
	err := sut.IssueTx(t.Context(), signed)
	require.ErrorContainsf(t, err, "already in pool", "%T.IssueTx() resubmit", sut.Client)
}
