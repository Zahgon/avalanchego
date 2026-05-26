// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"crypto/tls"
	"errors"
	"io"
)

var (
	ErrNoCertsSent        = errors.New("no certificates sent by peer")
	ErrEmptyCert          = errors.New("certificate sent by peer is empty")
	ErrEmptyPublicKey     = errors.New("no public key sent by peer")
	ErrCurveMismatch      = errors.New("only P256 is allowed for ECDSA")
	ErrUnsupportedKeyType = errors.New("key type is not supported")
)

// TLSConfig returns the TLS config that will allow secure connections to other
// peers.
//
// It is safe, and typically expected, for [keyLogWriter] to be [nil].
// [keyLogWriter] should only be enabled for debugging.
func TLSConfig(cert tls.Certificate, keyLogWriter io.Writer) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

// We do not use the TLS CA functionality to authenticate a
// hostname. We only require an authenticated channel based on the
// peer's public key. Therefore, we can safely skip CA verification.
//
// During our security audit by Quantstamp, this was investigated
// and confirmed to be safe and correct.
//#nosec G402

// ValidateCertificate validates TLS certificates according their public keys on the leaf certificate in the certification chain.
func ValidateCertificate(cs tls.ConnectionState) error { _ = "STUB: not implemented"; return nil }
