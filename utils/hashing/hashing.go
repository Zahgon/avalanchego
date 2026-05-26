// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package hashing

import (
	"crypto/sha256"
	"errors"

	// This file generates addresses from public keys with ripemd160. Though ripemd160 is not
	// generally recommended for use, the small size of the public key input is considered harder to
	// attack than larger payloads.
	//
	// Bitcoin similarly uses ripemd160 to generate addresses from public keys.
	//
	// Reference: https://online.tugraz.at/tug_online/voe_main2.getvolltext?pCurrPk=17675
	"golang.org/x/crypto/ripemd160" //nolint:gosec
)

const (
	HashLen = sha256.Size
	AddrLen = ripemd160.Size
)

var ErrInvalidHashLen = errors.New("invalid hash length")

// Hash256 A 256 bit long hash value.
type Hash256 = [HashLen]byte

// Hash160 A 160 bit long hash value.
type Hash160 = [ripemd160.Size]byte

// ComputeHash256Array computes a cryptographically strong 256 bit hash of the
// input byte slice.
func ComputeHash256Array(buf []byte) Hash256 {
	_ = "STUB: not implemented"
	return *

	// ComputeHash256 computes a cryptographically strong 256 bit hash of the input
	// byte slice.
	new(Hash256)
}

func ComputeHash256(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

// ComputeHash160Array computes a cryptographically strong 160 bit hash of the
// input byte slice.
func ComputeHash160Array(buf []byte) Hash160 { _ = "STUB: not implemented"; return *new(Hash160) }

// ComputeHash160 computes a cryptographically strong 160 bit hash of the input
// byte slice.
func ComputeHash160(buf []byte) []byte {
	_ = "STUB: not implemented"
	// See the comment on the ripemd160 import as to why the risk of use is
	// considered acceptable.
	return nil
}

//nolint:gosec

// Checksum creates a checksum of [length] bytes from the 256 bit hash of the
// byte slice.
//
// Returns: the lower [length] bytes of the hash
// Panics if length > 32.
func Checksum(bytes []byte, length int) []byte { _ = "STUB: not implemented"; return nil }

func ToHash256(bytes []byte) (Hash256, error) { _ = "STUB: not implemented"; return *new(Hash256), nil }

func ToHash160(bytes []byte) (Hash160, error) { _ = "STUB: not implemented"; return *new(Hash160), nil }

func PubkeyBytesToAddress(key []byte) []byte { _ = "STUB: not implemented"; return nil }
