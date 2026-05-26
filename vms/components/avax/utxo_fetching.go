// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

// GetBalance returns the current balance of [addrs]
func GetBalance(db UTXOReader, addrs set.Set[ids.ShortID]) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetAllUTXOs(db UTXOReader, addrs set.Set[ids.ShortID]) ([]*UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPaginatedUTXOs returns UTXOs such that at least one of the addresses in
// [addrs] is referenced.
//
// Returns at most [limit] UTXOs.
//
// Only returns UTXOs associated with addresses >= [startAddr].
//
// For address [startAddr], only returns UTXOs whose IDs are greater than
// [startUTXOID].
//
// Returns:
// * The fetched UTXOs
// * The address associated with the last UTXO fetched
// * The ID of the last UTXO fetched
func GetPaginatedUTXOs(
	db UTXOReader,
	addrs set.Set[ids.ShortID],
	lastAddr ids.ShortID,
	lastUTXOID ids.ID,
	limit int,
) ([]*UTXO, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}

// IDs of UTXOs already in the list
// the limit diminishes which can impact the expected return

// enforces the same ordering for pagination

// Skip addresses before [startAddr]

// The last address searched

// Get UTXOs associated with [addr]

// The last searched UTXO - not the last found

// Already have this UTXO in the list

// Found [limit] utxos; stop.

// Didn't reach the [limit] utxos; no more were found
