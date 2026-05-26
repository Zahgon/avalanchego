// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package localsigner

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/crypto/bls"

	blst "github.com/supranational/blst/bindings/go"
)

var (
	ErrFailedSecretKeyDeserialize            = errors.New("couldn't deserialize secret key")
	_                             bls.Signer = (*LocalSigner)(nil)
)

type secretKey = blst.SecretKey

type LocalSigner struct {
	sk *secretKey
	pk *bls.PublicKey
}

// NewSecretKey generates a new secret key from the local source of
// cryptographically secure randomness.
func New() (*LocalSigner, error) { _ = "STUB: not implemented"; return nil, nil }

// zero out the ikm

// ToBytes returns the big-endian format of the secret key.
func (s *LocalSigner) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

// FromBytes parses the big-endian format of the secret key into a
// secret key.
func FromBytes(skBytes []byte) (*LocalSigner, error) { _ = "STUB: not implemented"; return nil, nil }

func FromFile(keyPath string) (bls.Signer, error) {
	_ = "STUB: not implemented"
	return *new(bls.Signer), nil
}

func (s *LocalSigner) ToFile(keyPath string) error { _ = "STUB: not implemented"; return nil }

func FromFileOrPersistNew(keyPath string) (bls.Signer, error) {
	_ = "STUB: not implemented"
	return *new(bls.Signer), nil
}

// PublicKey returns the public key that corresponds to this secret
// key.
func (s *LocalSigner) PublicKey() *bls.PublicKey {
	_ = "STUB: not implemented"

	// Sign [msg] to authorize this message
	return nil
}

func (s *LocalSigner) Sign(msg []byte) (*bls.Signature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sign [msg] to prove the ownership
func (s *LocalSigner) SignProofOfPossession(msg []byte) (*bls.Signature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*LocalSigner) Shutdown() error { _ = "STUB: not implemented"; return nil }
