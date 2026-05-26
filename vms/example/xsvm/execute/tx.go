// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package execute

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/tx"
)

const (
	QuorumNumerator   = 2
	QuorumDenominator = 3
)

var (
	_ tx.Visitor = (*Tx)(nil)

	errFeeTooHigh          = errors.New("fee too high")
	errWrongChainID        = errors.New("wrong chainID")
	errMissingBlockContext = errors.New("missing block context")
	errDuplicateImport     = errors.New("duplicate import")
)

type Tx struct {
	Context      context.Context
	ChainContext *snow.Context
	Database     database.KeyValueReaderWriterDeleter

	SkipVerify   bool
	BlockContext *block.Context

	TxID        ids.ID
	Sender      ids.ShortID
	TransferFee uint64
	ExportFee   uint64
	ImportFee   uint64
}

func (t *Tx) Transfer(tf *tx.Transfer) error { _ = "STUB: not implemented"; return nil }

func (t *Tx) Export(e *tx.Export) error { _ = "STUB: not implemented"; return nil }

func (t *Tx) Import(i *tx.Import) error { _ = "STUB: not implemented"; return nil }
