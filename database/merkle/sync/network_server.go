// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/utils/units"

	pb "github.com/ava-labs/avalanchego/proto/pb/sync"
)

const (
	// Maximum number of key-value pairs to return in a proof.
	// This overrides any other Limit specified in a RangeProofRequest
	// or ChangeProofRequest if the given Limit is greater.
	MaxKeyValuesLimit = 2048
	// Estimated max overhead, in bytes, of putting a proof into a message.
	// We use this to ensure that the proof we generate is not too large to fit in a message.
	// TODO: refine this estimate. This is almost certainly a large overestimate.
	estimatedMessageOverhead = 4 * units.KiB
	maxByteSizeLimit         = constants.DefaultMaxMessageSize - estimatedMessageOverhead
)

var (
	_ p2p.Handler = (*ProofHandler[any, any])(nil)

	errMinProofSizeIsTooLarge = errors.New("cannot generate any proof within the requested limit")

	errInvalidBytesLimit    = errors.New("bytes limit must be greater than 0")
	errInvalidKeyLimit      = errors.New("key limit must be greater than 0")
	errInvalidStartRootHash = fmt.Errorf("start root hash must have length %d", hashing.HashLen)
	errInvalidEndRootHash   = fmt.Errorf("end root hash must have length %d", hashing.HashLen)
	errInvalidBounds        = errors.New("start key is greater than end key")
	errInvalidRootHash      = fmt.Errorf("root hash must have length %d", hashing.HashLen)
	errEmptyProof           = errors.New("proof for empty trie requested")
)

func NewProofHandler[R any, C any](db DB[R, C], rangeProofMarshaler Marshaler[R], changeProofMarshaler Marshaler[C]) *ProofHandler[R, C] {
	_ = "STUB: not implemented"
	return nil
}

type ProofHandler[R any, C any] struct {
	db                   DB[R, C]
	rangeProofMarshaler  Marshaler[R]
	changeProofMarshaler Marshaler[C]
}

func (*ProofHandler[_, _]) AppGossip(context.Context, ids.NodeID, []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *ProofHandler[R, C]) AppRequest(ctx context.Context, _ ids.NodeID, _ time.Time, requestBytes []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *ProofHandler[R, C]) handleRangeProofRequest(ctx context.Context, req *pb.RangeProofRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// override limits if they exceed caps

// drop request

// The proof was too large. Try to shrink it.

func (h *ProofHandler[R, C]) handleChangeProofRequest(ctx context.Context, req *pb.ChangeProofRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// override limits if they exceed caps

// We should only fail to get a change proof if we have insufficient history.
// Other errors are unexpected.
// TODO define custom errors

// g.db doesn't have endRoot in its history.
// We can't generate a change or range proof.

// g.db doesn't have sufficient history to generate change proof.
// Generate a range proof for the end root ID instead.

// We generated a change proof. See if it's small enough.

// The proof was too large. Try to shrink it.

// Returns nil iff [req] is well-formed.
func validateChangeProofRequest(req *pb.ChangeProofRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Returns nil iff [req] is well-formed.
func validateRangeProofRequest(req *pb.RangeProofRequest) error {
	_ = "STUB: not implemented"
	return nil
}
