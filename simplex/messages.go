// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
)

func newBlockProposal(
	chainID ids.ID,
	msg *simplex.VerifiedBlockMessage,
) (*p2p.Simplex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newVote(
	chainID ids.ID,
	vote *simplex.Vote,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newEmptyVote(
	chainID ids.ID,
	emptyVote *simplex.EmptyVote,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newFinalizeVote(
	chainID ids.ID,
	finalizeVote *simplex.FinalizeVote,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newNotarization(
	chainID ids.ID,
	notarization *simplex.Notarization,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newEmptyNotarization(
	chainID ids.ID,
	emptyNotarization *simplex.EmptyNotarization,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newFinalization(
	chainID ids.ID,
	finalization *simplex.Finalization,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newReplicationRequest(
	chainID ids.ID,
	replicationRequest *simplex.ReplicationRequest,
) *p2p.Simplex {
	_ = "STUB: not implemented"
	return nil
}

func newReplicationResponse(
	chainID ids.ID,
	replicationResponse *simplex.VerifiedReplicationResponse,
) (*p2p.Simplex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blockHeaderToP2P(bh simplex.BlockHeader) *p2p.BlockHeader {
	_ = "STUB: not implemented"
	return nil
}

func protocolMetadataToP2P(md simplex.ProtocolMetadata) *p2p.ProtocolMetadata {
	_ = "STUB: not implemented"
	return nil
}

func quorumRoundToP2P(qr *simplex.VerifiedQuorumRound) (*p2p.QuorumRound, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This can only happen if the finalization of the genesis block is being sent

func emptyVoteMetadataToP2P(ev simplex.EmptyVoteMetadata) *p2p.EmptyVoteMetadata {
	_ = "STUB: not implemented"
	return nil
}
