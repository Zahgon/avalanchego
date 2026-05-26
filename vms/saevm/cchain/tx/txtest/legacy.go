// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txtest

import (
	"errors"
	"testing"

	// Imported for [vm.VerifierBackend] comment resolution.
	_ "github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/vm"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
)

var errUnexpectedCredentialType = errors.New("unexpected credential type")

// ParseOld parses a transaction using coreth's old parsing logic while
// enforcing restrictions imposed by the new parsing logic.
//
// Coreth's parsing logic is overly permissive and depends on later verification
// in [vm.VerifierBackend].
func ParseOld(b []byte) (*atomic.Tx, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseOlds parses a slice of transaction using coreth's old parsing logic
// while enforcing restrictions imposed by the new parsing logic.
//
// Coreth's parsing logic is overly permissive and depends on later verification
// in [vm.VerifierBackend].
func ParseOlds(b []byte) ([]*atomic.Tx, error) { _ = "STUB: not implemented"; return nil, nil }

// ToOld converts a transaction from the new format into coreth's old format.
func ToOld(tb testing.TB, newTx *tx.Tx) *atomic.Tx { _ = "STUB: not implemented"; return nil }

// ToOlds converts a slice of transactions from the new format into coreth's old
// format.
func ToOlds(tb testing.TB, newTxs []*tx.Tx) []*atomic.Tx { _ = "STUB: not implemented"; return nil }

// ToNew converts a transaction from coreth's old format into the new format.
func ToNew(tb testing.TB, oldTx *atomic.Tx) *tx.Tx {
	_ = "STUB: not implemented"

	// Don't use [atomic.Tx.SignedBytes] in case it wasn't properly initialized.
	return nil
}

// ToNews converts a slice of transactions from coreth's old format into the new
// format.
func ToNews(tb testing.TB, oldTxs []*atomic.Tx) []*tx.Tx { _ = "STUB: not implemented"; return nil }
