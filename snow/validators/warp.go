// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"

	avajson "github.com/ava-labs/avalanchego/utils/json"
)

var _ utils.Sortable[*Warp] = (*Warp)(nil)

type WarpSet struct {
	// Slice, in canonical ordering, of the validators that have a public key.
	Validators []*Warp
	// The total weight of all the validators, including the ones that don't
	// have a public key.
	TotalWeight uint64
}

type jsonWarpSet struct {
	Validators  []*Warp        `json:"validators"`
	TotalWeight avajson.Uint64 `json:"totalWeight"`
}

func (w WarpSet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *WarpSet) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type Warp struct {
	PublicKey *bls.PublicKey
	// PublicKeyBytes is expected to be in the uncompressed form.
	PublicKeyBytes []byte
	Weight         uint64
	NodeIDs        []ids.NodeID
}

func (w *Warp) Compare(o *Warp) int { _ = "STUB: not implemented"; return 0 }

type jsonWarp struct {
	PublicKey string         `json:"publicKey"`
	Weight    avajson.Uint64 `json:"weight"`
	NodeIDs   []ids.NodeID   `json:"nodeIDs"`
}

func (w Warp) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Warp) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// FlattenValidatorSet converts the provided vdrSet into a canonical ordering.
func FlattenValidatorSet(vdrSet map[ids.NodeID]*GetValidatorOutput) (WarpSet, error) {
	_ = "STUB: not implemented"
	return *new(WarpSet), nil
}

// Impossible to overflow here

// Sort validators by public key
