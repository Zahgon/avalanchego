// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
)

var ErrInvalidParameters = errors.New("simplex parameters must be valid")

type ValidatorInfo struct {
	NodeID ids.NodeID `json:"nodeID" yaml:"nodeID"`

	// PublicKey is the public key of the validator.
	// It should be in the compressed public key format.
	PublicKey []byte `json:"publicKey" yaml:"publicKey"`
}

type Parameters struct {
	MaxNetworkDelay    time.Duration   `json:"maxNetworkDelay"    yaml:"maxNetworkDelay"`
	MaxRebroadcastWait time.Duration   `json:"maxRebroadcastWait" yaml:"maxRebroadcastWait"`
	InitialValidators  []ValidatorInfo `json:"initialValidators"  yaml:"initialValidators"`
}

var DefaultParameters = Parameters{
	MaxNetworkDelay:    5 * time.Second,
	MaxRebroadcastWait: 5 * time.Second,
}

func (p Parameters) Verify() error { _ = "STUB: not implemented"; return nil }

// TODO: we need to validate InitialValidators contains only unique nodes with valid keys.
// See: https://github.com/ava-labs/avalanchego/issues/5023
