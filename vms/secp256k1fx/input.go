// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"
)

const (
	CostPerSignature uint64 = 1000
)

var (
	ErrNilInput                    = errors.New("nil input")
	ErrInputIndicesNotSortedUnique = errors.New("address indices not sorted and unique")
)

type Input struct {
	// This input consumes an output, which has an owner list.
	// This input will be spent with a list of signatures.
	// SignatureList[i] is the signature of OwnerList[i]
	SigIndices []uint32 `serialize:"true" json:"signatureIndices"`
}

func (in *Input) Cost() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Verify this input is syntactically valid
func (in *Input) Verify() error { _ = "STUB: not implemented"; return nil }
