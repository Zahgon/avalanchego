// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package staking

import (
	"crypto/tls"
)

// InitNodeStakingKeyPair generates a self-signed TLS key/cert pair to use in
// staking. The key and files will be placed at [keyPath] and [certPath],
// respectively. If there is already a file at [keyPath], returns nil.
func InitNodeStakingKeyPair(keyPath, certPath string) error {
	_ = "STUB: not implemented"
	// If there is already a file at [keyPath], do nothing
	return nil
}

// Ensure directory where key/cert will live exist

// Write cert to disk

// Make cert read-only

// Write key to disk

// Make key read-only

func LoadTLSCertFromBytes(keyBytes, certBytes []byte) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadTLSCertFromFiles(keyPath, certPath string) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTLSCert() (*tls.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }

// Creates a new staking private key / staking certificate pair.
// Returns the PEM byte representations of both.
func NewCertAndKeyBytes() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	// Create key to sign cert with
	return nil, nil, nil
}

// Create self-signed staking cert
