// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package syncer

import (
	"github.com/ava-labs/firewood-go-ethhash/ffi"

	"github.com/ava-labs/avalanchego/database/merkle/sync"
	"github.com/ava-labs/avalanchego/ids"
)

var (
	_ sync.Marshaler[*RangeProof] = rangeProofMarshaler{}
	_ sync.Marshaler[struct{}]    = changeProofMarshaler{}
)

type rangeProofMarshaler struct{}

func (rangeProofMarshaler) Marshal(r *RangeProof) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rangeProofMarshaler) Unmarshal(data []byte) (*RangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RangeProof struct {
	rp        *ffi.RangeProof
	root      ids.ID
	maxLength int
}

// TODO: implement an actual ChangeProof marshaler.
type changeProofMarshaler struct{}

func (changeProofMarshaler) Marshal(struct{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (changeProofMarshaler) Unmarshal([]byte) (struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
