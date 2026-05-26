// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomictest

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/wrappers"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

// TODO: Remove this and use actual codec and transactions (export, import)
var (
	_           atomic.UnsignedAtomicTx = (*TestUnsignedTx)(nil)
	TestTxCodec codec.Manager
)

func init() {
	TestTxCodec = codec.NewDefaultManager()
	c := linearcodec.NewDefault()

	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&TestUnsignedTx{}),
		TestTxCodec.RegisterCodec(0, c),
	)

	if errs.Errored() {
		panic(errs.Err)
	}
}

type TestUnsignedTx struct {
	GasUsedV                    uint64                    `serialize:"true"`
	AcceptRequestsBlockchainIDV ids.ID                    `serialize:"true"`
	AcceptRequestsV             *avalancheatomic.Requests `serialize:"true"`
	VerifyV                     error
	IDV                         ids.ID `serialize:"true" json:"id"`
	BurnedV                     uint64 `serialize:"true"`
	UnsignedBytesV              []byte
	SignedBytesV                []byte
	InputUTXOsV                 set.Set[ids.ID]
	VisitV                      error
	EVMStateTransferV           error
}

// GasUsed implements the UnsignedAtomicTx interface
func (t *TestUnsignedTx) GasUsed(bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Verify implements the UnsignedAtomicTx interface
		nil
}

func (t *TestUnsignedTx) Verify(*snow.Context, extras.Rules) error {
	_ = "STUB: not implemented"

	// AtomicOps implements the UnsignedAtomicTx interface
	return nil
}

func (t *TestUnsignedTx) AtomicOps() (ids.ID, *avalancheatomic.Requests, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil, nil
}

// Initialize implements the UnsignedAtomicTx interface
func (*TestUnsignedTx) Initialize(_, _ []byte) {
	_ = "STUB: not implemented"

	// ID implements the UnsignedAtomicTx interface
	return
}

func (t *TestUnsignedTx) ID() ids.ID {
	_ = "STUB: not implemented"

	// Burned implements the UnsignedAtomicTx interface
	return *new(ids.ID)
}

func (t *TestUnsignedTx) Burned(ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Bytes implements the UnsignedAtomicTx interface
		nil
}

func (t *TestUnsignedTx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// SignedBytes implements the UnsignedAtomicTx interface
func (t *TestUnsignedTx) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }

// InputUTXOs implements the UnsignedAtomicTx interface
func (t *TestUnsignedTx) InputUTXOs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// Visit implements the UnsignedAtomicTx interface
func (t *TestUnsignedTx) Visit(atomic.Visitor) error {
	_ = "STUB: not implemented"

	// EVMStateTransfer implements the UnsignedAtomicTx interface
	return nil
}

func (t *TestUnsignedTx) EVMStateTransfer(*snow.Context, atomic.StateDB) error {
	_ = "STUB: not implemented"
	return nil
}

var TestBlockchainID = ids.GenerateTestID()

func GenerateTestImportTxWithGas(gasUsed uint64, burned uint64) *atomic.Tx {
	_ = "STUB: not implemented"
	return nil
}

func GenerateTestImportTx() *atomic.Tx { _ = "STUB: not implemented"; return nil }

func GenerateTestExportTx() *atomic.Tx { _ = "STUB: not implemented"; return nil }

func NewTestTx() *atomic.Tx { _ = "STUB: not implemented"; return nil }

func NewTestTxs(numTxs int) []*atomic.Tx { _ = "STUB: not implemented"; return nil }
