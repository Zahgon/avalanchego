// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/crypto/bls"
)

var (
	_ Signer = (*ProofOfPossession)(nil)

	ErrInvalidProofOfPossession = errors.New("invalid proof of possession")
)

type ProofOfPossession struct {
	PublicKey [bls.PublicKeyLen]byte `serialize:"true" json:"publicKey"`
	// BLS signature proving ownership of [PublicKey]. The signed message is the
	// [PublicKey].
	ProofOfPossession [bls.SignatureLen]byte `serialize:"true" json:"proofOfPossession"`

	// publicKey is the parsed version of [PublicKey]. It is populated in
	// [Verify].
	publicKey *bls.PublicKey
}

func NewProofOfPossession(sk bls.Signer) (*ProofOfPossession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ProofOfPossession) Verify() error { _ = "STUB: not implemented"; return nil }

func (p *ProofOfPossession) Key() *bls.PublicKey { _ = "STUB: not implemented"; return nil }

type jsonProofOfPossession struct {
	PublicKey         string `json:"publicKey"`
	ProofOfPossession string `json:"proofOfPossession"`
}

func (p *ProofOfPossession) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ProofOfPossession) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
