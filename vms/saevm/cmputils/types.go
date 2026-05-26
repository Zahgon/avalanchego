// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build !prod && !nocmpopts

package cmputils

import (
	"github.com/ava-labs/libevm/core/types"
	"github.com/google/go-cmp/cmp"
)

// BigInts returns a [cmp.Comparer] for [big.Int] pointers. A nil pointer is not
// equal to zero.
func BigInts() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// HexutilBigs returns a [cmp.Comparer] for [hexutil.Big] pointers. A nil
// pointer is not equal to zero.
func HexutilBigs() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// BlocksByHash returns a [cmp.Comparer] for [types.Block] pointers, equating
// them by hash alone.
func BlocksByHash() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// TransactionsByHash returns a [cmp.Comparer] for [types.Transaction] pointers,
// equating them by hash alone.
func TransactionsByHash() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// Receipts returns a set of [cmp.Options] for comparing [types.Receipt] values.
func Receipts() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// ReceiptsByTxHash returns a [cmp.Comparer] for [types.Receipt] pointers,
// equating them by transaction hash alone.
func ReceiptsByTxHash() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// CmpByMerkleRoots returns a [cmp.Comparer] for [types.DerivableList] values,
// equating them by their Merkle roots.
func CmpByMerkleRoots[T types.DerivableList]() cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

// Blocks returns a set of [cmp.Options] for comparing [types.Block] values.
// The [Headers] option MUST be used alongside this but isn't included
// automatically, to avoid duplication.
func Blocks() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// Headers returns a set of [cmp.Options] for comparing [type.Headers] values.
func Headers() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// Without the [IfIn] filter, any other use of [BigInts] will result in
// ambiguous comparers as [cmp] can't deduplicate them.

// LoadAtomicPointers returns a set of [cmp.Transformer] options that convert
// [atomic.Pointer] instances of `T` into their underlying `*T`. If the atomic
// under test is not itself a pointer (i.e. not *atomic.Pointer) then the
// returned options are NOT safe for concurrent use with said atomic as its lock
// is copied when passed as an argument to the transformer.
func LoadAtomicPointers[T any]() cmp.Options {
	_ = "STUB: not implemented"

	// Although accepting an [atomic.Pointer] value copies a lock, this is
	// unavoidable but OK in tests given the non-concurrency documentation
	// above.
	return *new(cmp.Options)
}

//nolint:govet

// NilSlicesAreEmpty returns a [cmp.Transformer] that converts `S(nil)` values
// into `S{}`, for use when [cmpopts.EquateEmpty] is too general.
func NilSlicesAreEmpty[S ~[]E, E any]() cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

func typeName[T any]() string { _ = "STUB: not implemented"; return "" }

// StateDBs returns a [cmp.Transformer] that converts [state.StateDB] instances
// into [state.Dump] equivalents.
func StateDBs() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }
