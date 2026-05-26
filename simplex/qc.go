// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

//go:generate go tool canoto $GOFILE

import (
	"errors"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ simplex.QuorumCertificate   = (*QC)(nil)
	_ simplex.QCDeserializer      = (*QCDeserializer)(nil)
	_ simplex.SignatureAggregator = (*SignatureAggregator)(nil)

	errFailedToParseQC       = errors.New("failed to parse quorum certificate")
	errUnexpectedSigners     = errors.New("unexpected number of signers in quorum certificate")
	errSignatureAggregation  = errors.New("signature aggregation failed")
	errEncodingMessageToSign = errors.New("failed to encode message to sign")
	errDuplicateSigner       = errors.New("duplicate signer in quorum certificate")
	errInvalidBitSet         = errors.New("bitset is invalid")
	errFailedToFilterSigners = errors.New("failed to filter signers")
)

// QC represents a quorum certificate in the Simplex consensus protocol.
type QC struct {
	verifier *BLSVerifier
	sig      *bls.Signature
	signers  []ids.NodeID
}

// CanotoQC is the Canoto representation of a quorum certificate
type canotoQC struct {
	Sig     [bls.SignatureLen]byte `canoto:"fixed bytes,1"`
	Signers []byte                 `canoto:"bytes,2"`

	canotoData canotoData_canotoQC
}

// Signers returns the list of signers for the quorum certificate.
func (qc *QC) Signers() []simplex.NodeID { _ = "STUB: not implemented"; return nil }

// Verify checks if the quorum certificate is valid by verifying the aggregated signature against the signers' public keys.
func (qc *QC) Verify(msg []byte) error { _ = "STUB: not implemented"; return nil }

// ensure signers are not double counted and are in the membership set

// aggregate the public keys

// Bytes serializes the quorum certificate into bytes.
func (qc *QC) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (qc *QC) createSignersBitSet() []byte { _ = "STUB: not implemented"; return nil }

// index should always exist, since we deserialized the signers from the same verifier

type QCDeserializer struct {
	verifier *BLSVerifier
}

// DeserializeQuorumCertificate deserializes a quorum certificate from bytes.
func (d *QCDeserializer) DeserializeQuorumCertificate(bytes []byte) (simplex.QuorumCertificate, error) {
	_ = "STUB: not implemented"
	return *new(simplex.QuorumCertificate), nil
}

// SignatureAggregator aggregates signatures into a quorum certificate.
type SignatureAggregator struct {
	verifier *BLSVerifier
}

// Aggregate aggregates the provided signatures into a quorum certificate.
// It requires at least a quorum of signatures to succeed.
// If any signature is from a signer not in the membership set, it returns an error.
func (a *SignatureAggregator) Aggregate(signatures []simplex.Signature) (simplex.QuorumCertificate, error) {
	_ = "STUB: not implemented"
	return *new(simplex.QuorumCertificate), nil
}

// IsQuorum checks if the provided nodes are a quorum of the membership set.
// For now, this is calculated using one node = one vote, but in the future we can adjust
// this calculation to cross reference validator weights if we want to support PoS.
func (a *SignatureAggregator) IsQuorum(nodes []simplex.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *QCDeserializer) signersFromBytes(signerBytes []byte) ([]ids.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filterNodes returns the nodeIDs in nodeIDs whose
// bit is set to 1 in indices.
//
// Returns an error if indices references an unknown node.
func filterNodes(
	indices set.Bits,
	nodeIDs []ids.NodeID,
) ([]ids.NodeID, error) {
	_ = "STUB: not implemented"
	// Verify that all alleged signers exist
	return nil, nil
}

// -1 to convert from length to index
