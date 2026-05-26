// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
)

var (
	_ Signature = (*BitSetSignature)(nil)

	ErrInvalidBitSet      = errors.New("bitset is invalid")
	ErrInsufficientWeight = errors.New("signature weight is insufficient")
	ErrInvalidSignature   = errors.New("signature is invalid")
	ErrParseSignature     = errors.New("failed to parse signature")
)

type Signature interface {
	fmt.Stringer

	// NumSigners is the number of [bls.PublicKeys] that participated in the
	// [Signature]. This is exposed because users of these signatures typically
	// impose a verification fee that is a function of the number of
	// signers.
	NumSigners() (int, error)

	// Verify that this signature was signed by at least [quorumNum]/[quorumDen]
	// of the validators of [msg.SourceChainID] at [pChainHeight].
	//
	// Invariant: [msg] is correctly initialized.
	Verify(
		msg *UnsignedMessage,
		networkID uint32,
		validators validators.WarpSet,
		quorumNum uint64,
		quorumDen uint64,
	) error
}

type BitSetSignature struct {
	// Signers is a big-endian byte slice encoding which validators signed this
	// message.
	Signers   []byte                 `serialize:"true"`
	Signature [bls.SignatureLen]byte `serialize:"true"`
}

func (s *BitSetSignature) NumSigners() (int, error) {
	_ = "STUB: not implemented"
	// Parse signer bit vector
	//
	// We assert that the length of [signerIndices.Bytes()] is equal
	// to [len(s.Signers)] to ensure that [s.Signers] does not have
	// any unnecessary zero-padding to represent the [set.Bits].
	return 0, nil
}

func (s *BitSetSignature) Verify(
	msg *UnsignedMessage,
	networkID uint32,
	validators validators.WarpSet,
	quorumNum uint64,
	quorumDen uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse signer bit vector
//
// We assert that the length of [signerIndices.Bytes()] is equal
// to [len(s.Signers)] to ensure that [s.Signers] does not have
// any unnecessary zero-padding to represent the [set.Bits].

// Get the validators that (allegedly) signed the message.

// Because [signers] is a subset of [validators.Validators], this can never error.

// Make sure the signature's weight is sufficient.

// Parse the aggregate signature

// Create the aggregate public key

// Verify the signature

func (s *BitSetSignature) String() string { _ = "STUB: not implemented"; return "" }

// VerifyWeight returns [nil] if [sigWeight] is at least [quorumNum]/[quorumDen]
// of [totalWeight].
// If [sigWeight >= totalWeight * quorumNum / quorumDen] then return [nil]
func VerifyWeight(
	sigWeight uint64,
	totalWeight uint64,
	quorumNum uint64,
	quorumDen uint64,
) error {
	_ = "STUB: not implemented"
	// Verifies that quorumNum * totalWeight <= quorumDen * sigWeight
	return nil
}
