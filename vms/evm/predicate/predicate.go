// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package predicate

import (
	"errors"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// delimiter separates the actual predicate bytes from the padded zero bytes.
//
// Predicates are encoded in the Access List of transactions by using in the
// access tuples. This means that the length must be a multiple of
// [common.HashLength].
//
// Even if the original predicate bytes is a multiple of [common.HashLength],
// the delimiter must be appended to support decoding.
const delimiter = 0xff

var (
	errMissingDelimiter = errors.New("no delimiter found")
	errExcessPadding    = errors.New("predicate included excess padding")
	errWrongDelimiter   = errors.New("wrong delimiter")
)

// Predicate is a message padded with the delimiter and zeros and chunked into
// 32-byte chunks.
type Predicate []common.Hash

// New constructs a predicate from raw predicate bytes.
//
// It chunks the predicate by appending [predicate.Delimiter] and zero-padding
// to a multiple of 32 bytes.
func New(b []byte) Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// Copy over chunks that don't require padding.

// Add the delimiter and required padding to the last chunk.

// Bytes converts the chunked predicate into the original message.
//
// Returns an error if it finds an incorrect encoding.
func (p Predicate) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type Predicates interface {
	HasPredicate(address common.Address) bool
}

// FromAccessList extracts predicates from a transaction's access list.
//
// If an address is specified multiple times in the access list, each set of
// storage keys for that address is considered an individual predicate.
func FromAccessList(rules Predicates, list types.AccessList) map[common.Address][]Predicate {
	_ = "STUB: not implemented"
	return nil
}
