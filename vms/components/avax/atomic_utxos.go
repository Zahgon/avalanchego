// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

// GetAtomicUTXOs returns exported UTXOs such that at least one of the
// addresses in [addrs] is referenced.
//
// Returns at most [limit] UTXOs.
//
// Returns:
// * The fetched UTXOs
// * The address associated with the last UTXO fetched
// * The ID of the last UTXO fetched
// * Any error that may have occurred upstream.
func GetAtomicUTXOs(
	sharedMemory atomic.SharedMemory,
	codec codec.Manager,
	chainID ids.ID,
	addrs set.Set[ids.ShortID],
	startAddr ids.ShortID,
	startUTXOID ids.ID,
	limit int,
) ([]*UTXO, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}
