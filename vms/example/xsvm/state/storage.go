// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

var (
	errWrongNonce          = errors.New("wrong nonce")
	errInsufficientBalance = errors.New("insufficient balance")
)

/*
 * VMDB
 * |-- initializedKey -> nil
 * |-. blocks
 * | |-- lastAcceptedKey -> blockID
 * | |-- height -> blockID
 * | '-- blockID -> block bytes
 * |-. addresses
 * | '-- addressID -> nonce
 * | '-- addressID + chainID -> balance
 * |-. chains
 * | |-- chainID -> balance
 * | '-- chainID + loanID -> nil
 * '-. message
 *   '-- txID -> message bytes
 */

// Chain state

func IsInitialized(db database.KeyValueReader) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func SetInitialized(db database.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }

// Block state

func GetLastAccepted(db database.KeyValueReader) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func SetLastAccepted(db database.KeyValueWriter, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func GetBlockIDByHeight(db database.KeyValueReader, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func GetBlock(db database.KeyValueReader, blkID ids.ID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddBlock(db database.KeyValueWriter, height uint64, blkID ids.ID, blk []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Address state

func GetNonce(db database.KeyValueReader, address ids.ShortID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func SetNonce(db database.KeyValueWriter, address ids.ShortID, nonce uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func IncrementNonce(db database.KeyValueReaderWriter, address ids.ShortID, nonce uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func GetBalance(db database.KeyValueReader, address ids.ShortID, chainID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func SetBalance(db database.KeyValueWriterDeleter, address ids.ShortID, chainID ids.ID, balance uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func DecreaseBalance(db database.KeyValueReaderWriterDeleter, address ids.ShortID, chainID ids.ID, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func IncreaseBalance(db database.KeyValueReaderWriterDeleter, address ids.ShortID, chainID ids.ID, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Chain state

func HasLoanID(db database.KeyValueReader, chainID ids.ID, loanID ids.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AddLoanID(db database.KeyValueWriter, chainID ids.ID, loanID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func GetLoan(db database.KeyValueReader, chainID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func SetLoan(db database.KeyValueWriterDeleter, chainID ids.ID, balance uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func DecreaseLoan(db database.KeyValueReaderWriterDeleter, chainID ids.ID, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func IncreaseLoan(db database.KeyValueReaderWriterDeleter, chainID ids.ID, amount uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Message state

func GetMessage(db database.KeyValueReader, txID ids.ID) (*warp.UnsignedMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetMessage(db database.KeyValueWriter, txID ids.ID, message *warp.UnsignedMessage) error {
	_ = "STUB: not implemented"
	return nil
}
