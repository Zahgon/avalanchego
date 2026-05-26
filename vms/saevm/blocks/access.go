// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"errors"
	"fmt"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/vms/saevm/types"
)

type (
	// A Chain provides access to the full life cycle of a [Block].
	Chain interface {
		ConsensusCritical
		Frontier
		DB() ethdb.Database
		XDB() types.ExecutionResults
	}

	// ConsensusCritical blocks are currently in use by a consensus mechanism,
	// at any point in the life cycle, including recently rejected. There is no
	// guarantee of access to every [Block] being used by consensus nor is there
	// a guarantee that a returned [Block] is still being used, but every
	// returned [Block] MUST be treated as consensus-critical.
	//
	// A ConsensusCritical source can be thought of as a cache of ready-made
	// [Block] instances, and consumers SHOULD treat return values as read-only.
	// Sources MUST be thread-safe and every returned [Block] MUST uphold all
	// life-cycle invariants.
	ConsensusCritical interface {
		ConsensusCriticalBlock(common.Hash) (*Block, bool)
	}

	// The Frontier is a thread-safe view of all life-cycle stages of a [Block].
	Frontier interface {
		AcceptanceFrontier
		ExecutionFrontier
		SettlementFrontier
	}
	// The AcceptanceFrontier is a thread-safe view of the last accepted
	// [Block].
	AcceptanceFrontier interface {
		LastAccepted() *Block
	}
	// The ExecutionFrontier is a thread-safe view of the last executed [Block].
	ExecutionFrontier interface {
		LastExecuted() *Block
	}
	// The SettlementFrontier is a thread-safe view of the last settled [Block].
	SettlementFrontier interface {
		LastSettled() *Block
	}
)

// ErrNotFound is returned by the Resolve*() and From*() functions when the
// requested [Block] could not be discerned or loaded, respectively.
//
// Note that the From*() functions DO NOT intercept return arguments. If a
// [DBReaderWithErr] returns `nil, nil`—as is common in geth when the `T` is not
// found—then so too will the From*() function return a nil error even though
// ErrNotFound is more appropriate. This pattern MUST NOT be considered as
// precedence as it is a foot gun and MUST be limited to, and contained within,
// packages in which it is idiomatic behaviour because of reliance on equivalent
// geth functionality.
var ErrNotFound = errors.New("block not found")

// ErrFutureBlockNotResolved is a specific case of [ErrNotFound], which it
// wraps, returned when attempting to resolve an [rpc.BlockNumber] not yet
// accepted by consensus. Such numbers are ambiguous as there may be more than
// one matching [Block].
var ErrFutureBlockNotResolved = fmt.Errorf("%w: not accepted yet", ErrNotFound)

// ErrNonCanonicalBlock is a specific case of [ErrFutureBlockNotResolved], which
// it wraps, returned when an unambiguous [Block] could be resolved (by its
// hash), but (a) it has not yet been accepted by consensus; and (b) the
// [rpc.BlockNumberOrHash.RequireCanonical] field was set to `true`.
var ErrNonCanonicalBlock = fmt.Errorf("%w: canonical block required", ErrFutureBlockNotResolved)

// ResolveRPCNumber converts an [rpc.BlockNumber] into the specific number of
// the corresponding block, treating named blocks as relative to execution:
//
//   - [rpc.PendingBlockNumber] is that returned by the [AcceptanceFrontier],
//     its execution status being unknown but eventually guaranteed.
//   - [rpc.LatestBlockNumber] is that returned by the [ExecutionFrontier].
//   - [rpc.SafeBlockNumber] and [rpc.FinalizedBlockNumber] are both that
//     returned by the [SettlementFrontier].
//
// Explicit block numbers are returned unchanged, as long as a corresponding
// [Block] has been accepted, otherwise [ErrFutureBlockNotResolved] is returned
// instead.
//
// # Important note re finality
//
// "Safe" and "finalized", as being separate to "last", are both Ethereum
// concepts with imperfect translation to SAE. Finality is the property of a
// canonical block being permanently so, without the possibility of re-org,
// which occurs immediately at acceptance under SAE. In Ethereum, "safety" is an
// interim stage towards finality, with greater but still imperfect guarantees.
//
// In keeping with "pending" and "latest", the point of reference for safety is
// relative to execution results, and the only non-determinism that can arise is
// due to (very rare) disk corruption. Settlement demonstrates consensus of
// execution results, such that they can be considered "safe" in the lay sense
// of the word. The "finalized" block is defined identically because we wish to
// maintain monotonicity of the labels, and no further guarantees are possible
// after settlement.
func ResolveRPCNumber(f Frontier, bn rpc.BlockNumber) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//#nosec G115 -- Non-negative check performed above

// Errors returned when resolving an invalid [rpc.BlockNumberOrHash].
var (
	ErrNeitherNumberNorHash = fmt.Errorf("%T carrying neither number nor hash", rpc.BlockNumberOrHash{})
	ErrBothNumberAndHash    = fmt.Errorf("%T carrying both number and hash", rpc.BlockNumberOrHash{})
)

// ResolveRPCNumberOrHash converts an [rpc.BlockNumberOrHash] into the specific
// number and hash of the corresponding [Block]. See [ResolveRPCNumber] for
// treatment of named block numbers.
//
// The [Block] with the returned [common.Hash] is guaranteed to have the
// returned number but is only guaranteed to be the canonical block of said
// number if (a) it was specified by [rpc.BlockNumberOrHash.Number], or (b) the
// [rpc.BlockNumberOrHash.RequireCanonical] field is `true`.
func ResolveRPCNumberOrHash(c Chain, numOrHash rpc.BlockNumberOrHash) (uint64, common.Hash, error) {
	_ = "STUB: not implemented"
	return 0, *new(common.Hash), nil
}

// [ResolveRPCNumber] is documented as only returning canonical blocks,
// so we don't need to check for a zero hash.

// TODO(JonathanOppenheimer): avoid the DB read to confirm if canonical

// We only write canonical blocks to the database so there's no need to
// perform a check.

type (
	// A DBReader returns any block-related artefact from the database. It is
	// typically one of the [rawdb] `Read*()` functions.
	DBReader[T any] func(db ethdb.Reader, hash common.Hash, num uint64) *T
	// A DBReaderWithErr is equivalent to a [DBReader] except that it MAY return
	// an error.
	DBReaderWithErr[T any] func(db ethdb.Reader, hash common.Hash, num uint64) (*T, error)
	// An Extractor returns any artefact from a [Block]. It is typically one of
	// the type's nulladic methods.
	Extractor[T any] func(*Block) *T
)

// WithNilErr converts the [DBReader] into a [DBReaderWithErr] that always
// returns a nil error.
func (r DBReader[T]) WithNilErr() DBReaderWithErr[T] { _ = "STUB: not implemented"; return nil }

// FromNumber resolves the canonical [Block] for the given [rpc.BlockNumber]
// and returns the result of calling `fromDB` with its number and hash.
func FromNumber[T any](c Chain, n rpc.BlockNumber, fromDB DBReaderWithErr[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromHash returns `fromConsensus()` if a [Block] with the specified hash is
// returned by the [ConsensusCritical] method of the [Chain], otherwise it returns
// `fromDB()` i.f.f. the block was previously accepted. If `fromDB()` is called
// then the block is guaranteed to exist if read with [rawdb] functions.
func FromHash[T any](c Chain, hash common.Hash, requireCanonical bool, fromConsensus Extractor[T], fromDB DBReaderWithErr[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(JonathanOppenheimer): avoid the DB read to confirm if canonical

// FromNumberOrHash resolves the [Block] for the given [rpc.BlockNumberOrHash]
// and returns the result of `fromConsensus` or `fromDB`, preferring the former.
// See [ResolveRPCNumberOrHash] for canonicality guarantees.
func FromNumberOrHash[T any](c Chain, blockNrOrHash rpc.BlockNumberOrHash, fromConsensus Extractor[T], fromDB DBReaderWithErr[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromNumberAndHash behaves like [FromNumberOrHash] except that it accepts both
// the number and hash as separate, required arguments. It verifies that any
// [Block] found in consensus has the expected number, and disallows named block
// numbers such as [rpc.LatestBlockNumber].
//
// Unlike [FromNumberOrHash], there are no canonicality guarantees as the
// provision of a hash resolves the ambiguity described in the comment on
// [ErrFutureBlockNotResolved].
func FromNumberAndHash[T any](c Chain, hash common.Hash, rpcNum rpc.BlockNumber, fromConsensus Extractor[T], fromDB DBReaderWithErr[T]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//#nosec G115 -- Non-negative check performed above
