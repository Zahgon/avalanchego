// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"context"
	"errors"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/proto/pb/p2p"
)

var (
	errNilField            = errors.New("nil field")
	errInvalidDigestLength = errors.New("invalid digest length")
	errInvalidSigner       = errors.New("invalid signer")
)

func emptyNotarizationMessageFromP2P(emptyNotarization *p2p.EmptyNotarization, qcDeserializer *QCDeserializer) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func notarizationMessageFromP2P(notarization *p2p.QuorumCertificate, qcDeserializer *QCDeserializer) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalizationMessageFromP2P(finalization *p2p.QuorumCertificate, qcDeserializer *QCDeserializer) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blockProposalFromP2P(ctx context.Context, blockProposal *p2p.BlockProposal, deserializer *blockDeserializer) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func voteFromP2P(vote *p2p.Vote) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func emptyVoteFromP2P(emptyVote *p2p.EmptyVote) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalizeVoteFromP2P(finalizeVote *p2p.Vote) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replicationRequestFromP2P(replicationRequest *p2p.ReplicationRequest) *simplex.Message {
	_ = "STUB: not implemented"
	return nil
}

func replicationResponseFromP2P(ctx context.Context, replicationResponse *p2p.ReplicationResponse, blockDeserializer *blockDeserializer, qcDeserializer *QCDeserializer) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HELPERS -----------------
func p2pVoteToSimplexVote(p2pVote *p2p.Vote) (simplex.Vote, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Vote), nil
}

func p2pSignatureToSimplexSignature(p2pSig *p2p.Signature) (simplex.Signature, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Signature), nil
}

func p2pBlockHeaderToSimplexBlockHeader(p2pHeader *p2p.BlockHeader) (simplex.BlockHeader, error) {
	_ = "STUB: not implemented"
	return *new(simplex.BlockHeader), nil
}

func p2pMetadataToSimplexMetadata(p2pMetadata *p2p.ProtocolMetadata) (simplex.ProtocolMetadata, error) {
	_ = "STUB: not implemented"
	return *new(simplex.ProtocolMetadata), nil
}

func emptyVoteMetadataFromP2P(emptyVote *p2p.EmptyVoteMetadata) (simplex.EmptyVoteMetadata, error) {
	_ = "STUB: not implemented"
	return *new(simplex.EmptyVoteMetadata), nil
}

func digestFromP2P(p2pDigest []byte) (simplex.Digest, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Digest), nil
}

func quorumCertificateFromP2P(qcBytes []byte, qcDeserializer *QCDeserializer) (simplex.QuorumCertificate, error) {
	_ = "STUB: not implemented"
	return *new(simplex.QuorumCertificate), nil
}

func notarizationFromP2P(notarization *p2p.QuorumCertificate, qcDeserializer *QCDeserializer) (*simplex.Notarization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func emptyNotarizationFromP2P(emptyNotarization *p2p.EmptyNotarization, qcDeserializer *QCDeserializer) (*simplex.EmptyNotarization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalizationFromP2P(finalization *p2p.QuorumCertificate, qcDeserializer *QCDeserializer) (*simplex.Finalization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func quorumRoundFromP2P(ctx context.Context, qr *p2p.QuorumRound, blockDeserializer *blockDeserializer, qcDeserializer *QCDeserializer) (*simplex.QuorumRound, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
