// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"crypto"
	"errors"
	"net/netip"
	"time"

	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
)

var (
	errTimestampTooFarInFuture = errors.New("timestamp too far in the future")
	errInvalidTLSSignature     = errors.New("invalid TLS signature")
)

// UnsignedIP is used for a validator to claim an IP. The [Timestamp] is used to
// ensure that the most updated IP claim is tracked by peers for a given
// validator.
type UnsignedIP struct {
	AddrPort  netip.AddrPort
	Timestamp uint64
}

// Sign this IP with the provided signer and return the signed IP.
func (ip *UnsignedIP) Sign(tlsSigner crypto.Signer, blsSigner bls.Signer) (*SignedIP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ip *UnsignedIP) bytes() []byte { _ = "STUB: not implemented"; return nil }

// SignedIP is a wrapper of an UnsignedIP with the signature from a signer.
type SignedIP struct {
	UnsignedIP
	TLSSignature      []byte
	BLSSignature      *bls.Signature
	BLSSignatureBytes []byte
}

// Returns nil if:
// * [ip.Timestamp] is not after [maxTimestamp].
// * [ip.TLSSignature] is a valid signature over [ip.UnsignedIP] from [cert].
func (ip *SignedIP) Verify(
	cert *staking.Certificate,
	maxTimestamp time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}
