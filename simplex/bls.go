// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"errors"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
)

var (
	errSignatureVerificationFailed = errors.New("signature verification failed")
	errSignerNotFound              = errors.New("signer not found in the membership set")
	errInvalidNodeID               = errors.New("unable to parse node ID")
	errFailedToParseSignature      = errors.New("failed to parse signature")
)

var _ simplex.Signer = (*BLSSigner)(nil)

type SignFunc func(msg []byte) (*bls.Signature, error)

// BLSSigner signs messages encoded with the provided ChainID and NetworkID.
// using the SignBLS function.
type BLSSigner struct {
	chainID   ids.ID
	networkID uint32
	// signBLS is passed in because we support both software and hardware BLS signing.
	signBLS SignFunc
}

type BLSVerifier struct {
	nodeID2PK map[ids.NodeID]*bls.PublicKey
	networkID uint32
	chainID   ids.ID

	canonicalNodeIDs       []ids.NodeID
	canonicalNodeIDIndices map[ids.NodeID]int
}

func NewBLSAuth(config *Config) (BLSSigner, BLSVerifier, error) {
	_ = "STUB: not implemented"
	return *new(BLSSigner), *new(BLSVerifier), nil
}

// Sign returns a signature on the given message using BLS signature scheme.
// It encodes the message to sign with the chain ID, and network ID,
func (s *BLSSigner) Sign(message []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type encodedSimplexSignedPayload struct {
	NetworkID uint32 `serialize:"true"`
	ChainID   ids.ID `serialize:"true"`
	Message   []byte `serialize:"true"`
}

func encodeMessageToSign(message []byte, chainID ids.ID, networkID uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v BLSVerifier) Verify(message []byte, signature []byte, signer simplex.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func createVerifier(config *Config) (BLSVerifier, error) {
	_ = "STUB: not implemented"
	return *new(BLSVerifier), nil
}
