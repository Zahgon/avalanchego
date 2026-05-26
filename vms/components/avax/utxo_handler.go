// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import "github.com/ava-labs/avalanchego/ids"

// Removes the UTXOs consumed by [ins] from the UTXO set
func Consume(utxoDB UTXODeleter, ins []*TransferableInput) { _ = "STUB: not implemented"; return }

// Adds the UTXOs created by [outs] to the UTXO set.
// [txID] is the ID of the tx that created [outs].
func Produce(
	utxoDB UTXOAdder,
	txID ids.ID,
	outs []*TransferableOutput,
) {
	_ = "STUB: not implemented"
	return
}
