// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

const (
	minNameLen      = 1
	maxNameLen      = 128
	minSymbolLen    = 1
	maxSymbolLen    = 4
	maxDenomination = 32
)

var (
	_ txs.Visitor = (*SyntacticVerifier)(nil)

	errWrongNumberOfCredentials     = errors.New("wrong number of credentials")
	errInitialStatesNotSortedUnique = errors.New("initial states not sorted and unique")
	errNameTooShort                 = fmt.Errorf("name is too short, minimum size is %d", minNameLen)
	errNameTooLong                  = fmt.Errorf("name is too long, maximum size is %d", maxNameLen)
	errSymbolTooShort               = fmt.Errorf("symbol is too short, minimum size is %d", minSymbolLen)
	errSymbolTooLong                = fmt.Errorf("symbol is too long, maximum size is %d", maxSymbolLen)
	errNoFxs                        = errors.New("assets must support at least one Fx")
	errIllegalNameCharacter         = errors.New("asset's name must be made up of only letters and numbers")
	errIllegalSymbolCharacter       = errors.New("asset's symbol must be all upper case letters")
	errUnexpectedWhitespace         = errors.New("unexpected whitespace provided")
	errDenominationTooLarge         = errors.New("denomination is too large")
	errOperationsNotSortedUnique    = errors.New("operations not sorted and unique")
	errNoOperations                 = errors.New("an operationTx must have at least one operation")
	errDoubleSpend                  = errors.New("inputs attempt to double spend an input")
	errNoImportInputs               = errors.New("no import inputs")
	errNoExportOutputs              = errors.New("no export outputs")
)

type SyntacticVerifier struct {
	*Backend
	Tx *txs.Tx
}

func (v *SyntacticVerifier) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (v *SyntacticVerifier) CreateAssetTx(tx *txs.CreateAssetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SyntacticVerifier) OperationTx(tx *txs.OperationTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SyntacticVerifier) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (v *SyntacticVerifier) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }
