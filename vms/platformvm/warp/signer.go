// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
)

var (
	_ Signer = (*signer)(nil)

	ErrWrongSourceChainID = errors.New("wrong SourceChainID")
	ErrWrongNetworkID     = errors.New("wrong networkID")
)

type Signer interface {
	// Returns this node's BLS signature over an unsigned message. If the caller
	// does not have the authority to sign the message, an error will be
	// returned.
	//
	// Assumes the unsigned message is correctly initialized.
	Sign(msg *UnsignedMessage) ([]byte, error)
}

func NewSigner(sk bls.Signer, networkID uint32, chainID ids.ID) Signer {
	_ = "STUB: not implemented"
	return *new(Signer)
}

type signer struct {
	sk        bls.Signer
	networkID uint32
	chainID   ids.ID
}

func (s *signer) Sign(msg *UnsignedMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
