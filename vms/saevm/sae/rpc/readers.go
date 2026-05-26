// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
)

// neverErrs is a convenience wrapper, intended for use with [rawdb] Read*()
// functions, to avoid having to cast them as [blocks.DBReader] before calling
// [blocks.DBReader.WithNilErr]. It makes call sites cleaner.
func neverErrs[T any](r blocks.DBReader[T]) blocks.DBReaderWithErr[T] {
	_ = "STUB: not implemented"
	return nil
}

func notFoundIsNil[T any](x *T, err error) (*T, error) {
	_ = "STUB: not implemented"
	// [blocks.ErrNonCanonicalBlock] wraps [blocks.ErrNotFound], which
	// would be a misleading error to return.
	return nil, nil
}

// Note that these readers will only work for canonical blocks (blocks that are guaranteed
// to be executed) to ensure that every block will eventually have post-execution artefacts.
// Non-canonical blocks are rejected with [blocks.ErrNonCanonicalBlock].

func readByNumber[T any](c Chain, n rpc.BlockNumber, read blocks.DBReader[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readByHash[T any](c Chain, hash common.Hash, fromMem blocks.Extractor[T], fromDB blocks.DBReader[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readByNumberOrHash[T any](c Chain, blockNrOrHash rpc.BlockNumberOrHash, fromMem blocks.Extractor[T], fromDB blocks.DBReaderWithErr[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readByNumberAndHash[T any](c Chain, h common.Hash, num rpc.BlockNumber, fromMem blocks.Extractor[T], fromDB blocks.DBReader[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
