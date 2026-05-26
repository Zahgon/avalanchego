// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"context"

	"github.com/ava-labs/libevm/common"
)

// API describes the set of methods offered over the RPC interface
type API struct {
	Namespace string      // namespace under which the rpc methods of Service are exposed
	Version   string      // deprecated - this field is no longer used, but retained for compatibility
	Service   interface{} // receiver instance which holds the methods
	Name      string      // Name of the API
}

// ServerCodec implements reading, parsing and writing RPC messages for the server side of
// a RPC session. Implementations must be go-routine safe since the codec can be called in
// multiple go-routines concurrently.
type ServerCodec interface {
	peerInfo() PeerInfo
	readBatch() (msgs []*jsonrpcMessage, isBatch bool, err error)
	close()

	jsonWriter
}

// jsonWriter can write JSON messages to its underlying connection.
// Implementations must be safe for concurrent use.
type jsonWriter interface {
	// writeJSON writes a message to the connection.
	writeJSON(ctx context.Context, msg interface{}, isError bool) error
	// writeJSONSkipDeadline writes a message to the connection with the option of skipping the deadline.
	writeJSONSkipDeadline(ctx context.Context, msg interface{}, isError bool, skip bool) error
	// Closed returns a channel which is closed when the connection is closed.
	closed() <-chan interface{}
	// RemoteAddr returns the peer address of the connection.
	remoteAddr() string
}

type BlockNumber int64

const (
	SafeBlockNumber      = BlockNumber(-4)
	FinalizedBlockNumber = BlockNumber(-3)
	LatestBlockNumber    = BlockNumber(-2)
	PendingBlockNumber   = BlockNumber(-1)
	EarliestBlockNumber  = BlockNumber(0)
)

// UnmarshalJSON parses the given JSON fragment into a BlockNumber. It supports:
// - "accepted", "safe", "finalized", "latest", "earliest" or "pending" as string arguments
// - the block number
// Returned errors:
// - an invalid block number error when the given argument isn't a known strings
// - an out of range error when the given block number is either too little or too large
func (bn *BlockNumber) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Include "finalized" as an option for compatibility with FinalizedBlockNumber from geth.

// Int64 returns the block number as int64.
func (bn BlockNumber) Int64() int64 {
	_ = "STUB: not implemented"

	// MarshalText implements encoding.TextMarshaler. It marshals:
	// - "accepted", "latest", "earliest" or "pending" as strings
	// - other numbers as hex
	return 0
}

func (bn BlockNumber) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (bn BlockNumber) String() string { _ = "STUB: not implemented"; return "" }

// IsAccepted returns true if this blockNumber should be treated as a request for the last accepted block
func (bn BlockNumber) IsAccepted() bool { _ = "STUB: not implemented"; return false }

func (bn BlockNumber) IsLatest() bool { _ = "STUB: not implemented"; return false }

type BlockNumberOrHash struct {
	BlockNumber      *BlockNumber `json:"blockNumber,omitempty"`
	BlockHash        *common.Hash `json:"blockHash,omitempty"`
	RequireCanonical bool         `json:"requireCanonical,omitempty"`
}

func (bnh *BlockNumberOrHash) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Include "finalized" as an option for compatibility with FinalizedBlockNumber from geth.

func (bnh *BlockNumberOrHash) Number() (BlockNumber, bool) {
	_ = "STUB: not implemented"
	return *new(BlockNumber), false
}

func (bnh *BlockNumberOrHash) String() string { _ = "STUB: not implemented"; return "" }

func (bnh *BlockNumberOrHash) Hash() (common.Hash, bool) {
	_ = "STUB: not implemented"
	return *new(common.Hash), false
}

func BlockNumberOrHashWithNumber(blockNr BlockNumber) BlockNumberOrHash {
	_ = "STUB: not implemented"
	return *new(BlockNumberOrHash)
}

func BlockNumberOrHashWithHash(hash common.Hash, canonical bool) BlockNumberOrHash {
	_ = "STUB: not implemented"
	return *new(BlockNumberOrHash)
}
