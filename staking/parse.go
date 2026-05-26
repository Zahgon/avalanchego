// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package staking

import (
	"crypto"
	"crypto/rsa"
	"encoding/asn1"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	MaxCertificateLen = 2 * units.KiB

	allowedRSASmallModulusLen     = 2048
	allowedRSALargeModulusLen     = 4096
	allowedRSAPublicExponentValue = 65537
)

var (
	ErrCertificateTooLarge                   = fmt.Errorf("staking: certificate length is greater than %d", MaxCertificateLen)
	ErrMalformedCertificate                  = errors.New("staking: malformed certificate")
	ErrMalformedTBSCertificate               = errors.New("staking: malformed tbs certificate")
	ErrMalformedVersion                      = errors.New("staking: malformed version")
	ErrMalformedSerialNumber                 = errors.New("staking: malformed serial number")
	ErrMalformedSignatureAlgorithmIdentifier = errors.New("staking: malformed signature algorithm identifier")
	ErrMalformedIssuer                       = errors.New("staking: malformed issuer")
	ErrMalformedValidity                     = errors.New("staking: malformed validity")
	ErrMalformedSPKI                         = errors.New("staking: malformed spki")
	ErrMalformedPublicKeyAlgorithmIdentifier = errors.New("staking: malformed public key algorithm identifier")
	ErrMalformedSubjectPublicKey             = errors.New("staking: malformed subject public key")
	ErrMalformedOID                          = errors.New("staking: malformed oid")
	ErrInvalidRSAPublicKey                   = errors.New("staking: invalid RSA public key")
	ErrInvalidRSAModulus                     = errors.New("staking: invalid RSA modulus")
	ErrInvalidRSAPublicExponent              = errors.New("staking: invalid RSA public exponent")
	ErrRSAModulusNotPositive                 = errors.New("staking: RSA modulus is not a positive number")
	ErrUnsupportedRSAModulusBitLen           = errors.New("staking: unsupported RSA modulus bitlen")
	ErrRSAModulusIsEven                      = errors.New("staking: RSA modulus is an even number")
	ErrUnsupportedRSAPublicExponent          = errors.New("staking: unsupported RSA public exponent")
	ErrFailedUnmarshallingEllipticCurvePoint = errors.New("staking: failed to unmarshal elliptic curve point")
	ErrUnknownPublicKeyAlgorithm             = errors.New("staking: unknown public key algorithm")
)

// ParseCertificate parses a single certificate from the given ASN.1.
//
// This function does not validate that the certificate is valid to be used
// against normal TLS implementations.
//
// Ref: https://github.com/golang/go/blob/go1.19.12/src/crypto/x509/parser.go#L789-L968
func ParseCertificate(bytes []byte) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume the length and tag bytes.

// Read the "to be signed" certificate into input.

// Read the "subject public key info" into input.

// Read the public key algorithm identifier.

// Note: Unlike the x509 package, we require parsing the public key.

// Ref: https://github.com/golang/go/blob/go1.19.12/src/crypto/x509/parser.go#L215-L306
func parsePublicKey(oid asn1.ObjectIdentifier, publicKey asn1.BitString) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

// ValidateRSAPublicKeyIsWellFormed validates the given RSA public key
func ValidateRSAPublicKeyIsWellFormed(pub *rsa.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}
