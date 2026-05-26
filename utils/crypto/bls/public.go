// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bls

import (
	"errors"

	blst "github.com/supranational/blst/bindings/go"
)

const PublicKeyLen = blst.BLST_P1_COMPRESS_BYTES

var (
	ErrNoPublicKeys               = errors.New("no public keys")
	ErrFailedPublicKeyDecompress  = errors.New("couldn't decompress public key")
	errInvalidPublicKey           = errors.New("invalid public key")
	errFailedPublicKeyAggregation = errors.New("couldn't aggregate public keys")
)

type (
	PublicKey          = blst.P1Affine
	AggregatePublicKey = blst.P1Aggregate
)

// PublicKeyToCompressedBytes returns the compressed big-endian format of the
// public key.
func PublicKeyToCompressedBytes(pk *PublicKey) []byte { _ = "STUB: not implemented"; return nil }

// PublicKeyFromCompressedBytes parses the compressed big-endian format of the
// public key into a public key.
func PublicKeyFromCompressedBytes(pkBytes []byte) (*PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublicKeyToUncompressedBytes returns the uncompressed big-endian format of
// the public key.
func PublicKeyToUncompressedBytes(key *PublicKey) []byte { _ = "STUB: not implemented"; return nil }

// PublicKeyFromValidUncompressedBytes parses the uncompressed big-endian format
// of the public key into a public key. It is assumed that the provided bytes
// are valid.
func PublicKeyFromValidUncompressedBytes(pkBytes []byte) *PublicKey {
	_ = "STUB: not implemented"
	return nil
}

// AggregatePublicKeys aggregates a non-zero number of public keys into a single
// aggregated public key.
// Invariant: all [pks] have been validated.
func AggregatePublicKeys(pks []*PublicKey) (*PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify the [sig] of [msg] against the [pk].
// The [sig] and [pk] may have been an aggregation of other signatures and keys.
// Invariant: [pk] and [sig] have both been validated.
func Verify(pk *PublicKey, sig *Signature, msg []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify the possession of the secret pre-image of [sk] by verifying a [sig] of
// [msg] against the [pk].
// The [sig] and [pk] may have been an aggregation of other signatures and keys.
// Invariant: [pk] and [sig] have both been validated.
func VerifyProofOfPossession(pk *PublicKey, sig *Signature, msg []byte) bool {
	_ = "STUB: not implemented"
	return false
}
