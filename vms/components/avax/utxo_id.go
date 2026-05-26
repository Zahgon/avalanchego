// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	errNilUTXOID                 = errors.New("nil utxo ID is not valid")
	errMalformedUTXOIDString     = errors.New("unexpected number of tokens in string")
	errFailedDecodingUTXOIDTxID  = errors.New("failed decoding UTXOID TxID")
	errFailedDecodingUTXOIDIndex = errors.New("failed decoding UTXOID index")

	_ verify.Verifiable       = (*UTXOID)(nil)
	_ utils.Sortable[*UTXOID] = (*UTXOID)(nil)
)

type UTXOID struct {
	// Serialized:
	TxID        ids.ID `serialize:"true" json:"txID"`
	OutputIndex uint32 `serialize:"true" json:"outputIndex"`

	// Symbol is false if the UTXO should be part of the DB
	Symbol bool `json:"-"`
	// id is the unique ID of a UTXO, it is calculated from TxID and OutputIndex
	id ids.ID
}

// InputSource returns the source of the UTXO that this input is spending
func (utxo *UTXOID) InputSource() (ids.ID, uint32) {
	_ = "STUB: not implemented"
	return *new(ids.ID), 0
}

// InputID returns a unique ID of the UTXO that this input is spending
func (utxo *UTXOID) InputID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// Symbolic returns if this is the ID of a UTXO in the DB, or if it is a
// symbolic input
func (utxo *UTXOID) Symbolic() bool { _ = "STUB: not implemented"; return false }

func (utxo *UTXOID) String() string { _ = "STUB: not implemented"; return "" }

// UTXOIDFromString attempts to parse a string into a UTXOID
func UTXOIDFromString(s string) (*UTXOID, error) { _ = "STUB: not implemented"; return nil, nil }

func (utxo *UTXOID) Verify() error { _ = "STUB: not implemented"; return nil }

func (utxo *UTXOID) Compare(other *UTXOID) int { _ = "STUB: not implemented"; return 0 }
