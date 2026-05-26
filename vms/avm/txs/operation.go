// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var (
	ErrNilOperation              = errors.New("nil operation is not valid")
	ErrNilFxOperation            = errors.New("nil fx operation is not valid")
	ErrNotSortedAndUniqueUTXOIDs = errors.New("utxo IDs not sorted and unique")
)

type Operation struct {
	avax.Asset `serialize:"true"`
	UTXOIDs    []*avax.UTXOID  `serialize:"true"  json:"inputIDs"`
	FxID       ids.ID          `serialize:"false" json:"fxID"`
	Op         fxs.FxOperation `serialize:"true"  json:"operation"`
}

func (op *Operation) Verify() error { _ = "STUB: not implemented"; return nil }

type operationAndCodec struct {
	op    *Operation
	codec codec.Manager
}

func (o *operationAndCodec) Compare(other *operationAndCodec) int {
	_ = "STUB: not implemented"
	return 0
}

func SortOperations(ops []*Operation, c codec.Manager) { _ = "STUB: not implemented"; return }

func IsSortedAndUniqueOperations(ops []*Operation, c codec.Manager) bool {
	_ = "STUB: not implemented"
	return false
}

type innerSortOperationsWithSigners struct {
	ops     []*Operation
	signers [][]*secp256k1.PrivateKey
	codec   codec.Manager
}

func (ops *innerSortOperationsWithSigners) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}

func (ops *innerSortOperationsWithSigners) Len() int { _ = "STUB: not implemented"; return 0 }

func (ops *innerSortOperationsWithSigners) Swap(i, j int) { _ = "STUB: not implemented"; return }

func SortOperationsWithSigners(ops []*Operation, signers [][]*secp256k1.PrivateKey, codec codec.Manager) {
	_ = "STUB: not implemented"
	return
}
