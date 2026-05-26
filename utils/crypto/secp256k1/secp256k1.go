// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1

import (
	"errors"
	"fmt"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"

	stdecdsa "crypto/ecdsa"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

const (
	// SignatureLen is the number of bytes in a secp2561k recoverable signature
	SignatureLen = 65

	// PrivateKeyLen is the number of bytes in a secp2561k recoverable private
	// key
	PrivateKeyLen = 32

	// PublicKeyLen is the number of bytes in a secp2561k recoverable public key
	PublicKeyLen = 33

	// from the decred library:
	// compactSigMagicOffset is a value used when creating the compact signature
	// recovery code inherited from Bitcoin and has no meaning, but has been
	// retained for compatibility.  For historical purposes, it was originally
	// picked to avoid a binary representation that would allow compact
	// signatures to be mistaken for other components.
	compactSigMagicOffset = 27

	PrivateKeyPrefix = "PrivateKey-"
	nullStr          = "null"
)

var (
	ErrInvalidSig              = errors.New("invalid signature")
	errCompressed              = errors.New("wasn't expecting a compressed key")
	errMissingQuotes           = errors.New("first and last characters should be quotes")
	errMissingKeyPrefix        = fmt.Errorf("private key missing %s prefix", PrivateKeyPrefix)
	errInvalidPrivateKeyLength = fmt.Errorf("private key has unexpected length, expected %d", PrivateKeyLen)
	errInvalidPublicKeyLength  = fmt.Errorf("public key has unexpected length, expected %d", PublicKeyLen)
	errInvalidSigLen           = errors.New("invalid signature length")
	errMutatedSig              = errors.New("signature was mutated from its original format")
)

func NewPrivateKey() (*PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func ToPublicKey(b []byte) (*PublicKey, error) { _ = "STUB: not implemented"; return nil, nil }

func ToPrivateKey(b []byte) (*PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func RecoverPublicKey(msg, sig []byte) (*PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RecoverPublicKeyFromHash(hash, sig []byte) (*PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RecoverCache struct {
	cache cache.Cacher[ids.ID, *PublicKey]
}

func NewRecoverCache(size int) *RecoverCache { _ = "STUB: not implemented"; return nil }

func (r *RecoverCache) RecoverPublicKey(msg, sig []byte) (*PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RecoverCache) RecoverPublicKeyFromHash(hash, sig []byte) (*PublicKey, error) {
	_ = "STUB: not implemented"
	// TODO: This type should always be initialized by calling NewRecoverCache.
	return nil, nil
}

type PublicKey struct {
	pk    *secp256k1.PublicKey
	addr  ids.ShortID
	bytes []byte
}

func (k *PublicKey) Verify(msg, sig []byte) bool { _ = "STUB: not implemented"; return false }

func (k *PublicKey) VerifyHash(hash, sig []byte) bool { _ = "STUB: not implemented"; return false }

// ToECDSA returns the ecdsa representation of this public key
func (k *PublicKey) ToECDSA() *stdecdsa.PublicKey { _ = "STUB: not implemented"; return nil }

func (k *PublicKey) Address() ids.ShortID { _ = "STUB: not implemented"; return *new(ids.ShortID) }

func (k *PublicKey) EthAddress() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (k *PublicKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

type PrivateKey struct {
	sk    *secp256k1.PrivateKey
	pk    *PublicKey
	bytes []byte
}

func (k *PrivateKey) PublicKey() *PublicKey { _ = "STUB: not implemented"; return nil }

func (k *PrivateKey) Address() ids.ShortID { _ = "STUB: not implemented"; return *new(ids.ShortID) }

func (k *PrivateKey) EthAddress() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (k *PrivateKey) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PrivateKey) SignHash(hash []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// returns [v || r || s]

// ToECDSA returns the ecdsa representation of this private key
func (k *PrivateKey) ToECDSA() *stdecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (k *PrivateKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (k *PrivateKey) String() string {
	_ = "STUB: not implemented"
	// We assume that the maximum size of a byte slice that
	// can be stringified is at least the length of a SECP256K1 private key
	return ""
}

func (k *PrivateKey) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PrivateKey) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PrivateKey) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If "null", do nothing

func (k *PrivateKey) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (k *PrivateKey) unmarshalText(text string) error { _ = "STUB: not implemented"; return nil }

// raw sig has format [v || r || s] whereas the sig has format [r || s || v]
func rawSigToSig(sig []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// sig has format [r || s || v] whereas the raw sig has format [v || r || s]
func sigToRawSig(sig []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G602: length is validated above

// verifies the signature format in format [r || s || v]
func verifySECP256K1RSignatureFormat(sig []byte) error { _ = "STUB: not implemented"; return nil }
