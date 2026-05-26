// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package staking

import (
	"errors"
)

var (
	ErrUnsupportedAlgorithm     = errors.New("staking: cannot verify signature: unsupported algorithm")
	ErrECDSAVerificationFailure = errors.New("staking: ECDSA verification failure")
)

// CheckSignature verifies that the signature is a valid signature over signed
// from the certificate.
//
// Ref: https://github.com/golang/go/blob/go1.19.12/src/crypto/x509/x509.go#L793-L797
// Ref: https://github.com/golang/go/blob/go1.19.12/src/crypto/x509/x509.go#L816-L879
func CheckSignature(cert *Certificate, msg []byte, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}
