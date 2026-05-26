// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/types"
)

type ancestry struct {
	parent, lastSettled *Block
}

var (
	errBlockResettled       = errors.New("block re-settled")
	errBlockAncestryChanged = errors.New("block ancestry changed during settlement")
)

// MarkSettled marks the block as having been settled. This function MUST NOT be
// called more than once. The atomic pointer to the last-settled block is
// updated before [Block.WaitUntilSettled] returns.
//
// After a call to MarkSettled, future calls to [Block.ParentBlock] and
// [Block.LastSettled] will return nil.
func (b *Block) MarkSettled(lastSettled *atomic.Pointer[Block]) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) markSettled(lastSettled *atomic.Pointer[Block]) error {
	_ = "STUB: not implemented"
	return nil
}

// almost certainly means concurrent calls to this method

// We have to return something to keen the compiler happy, even though we
// expect the Fatal to be, well, fatal.

// MarkSynchronous combines [Block.MarkExecuted] and [Block.MarkSettled], and is
// reserved for the last pre-SAE block, which MAY be the genesis block. These
// blocks are, by definition, self-settling so require special treatment as such
// behaviour is impossible under SAE rules.
//
// Arguments required by [Block.MarkExecuted] but not accepted by
// MarkSynchronous are derived from the block to maintain invariants. The
// `subSecondBlockTime` argument MUST follow the same constraints as the
// respective [hook.Points] method.
//
// MarkSynchronous and [Block.Synchronous] are not safe for concurrent use. This
// method MUST therefore be called *before* instantiating the SAE VM.
//
// Wherever MarkSynchronous results in different behaviour to
// [Block.MarkSettled], the respective methods are documented as such. They can
// otherwise be considered identical.
//
// Unlike [Block.MarkExecuted], MarkSynchronous does not call
// [Block.SetAsHeadBlock], which MUST be done by the caller, i.f.f. the chain
// has not yet commenced asynchronous execution.
//
// TODO(arr4n) refactor to avoid requiring DB writes.
func (b *Block) MarkSynchronous(hooks hook.Points, db ethdb.Database, xdb types.ExecutionResults, excessAfter gas.Gas) error {
	_ = "STUB: not implemented"
	return nil

	// Receipts of a synchronous block have already been "settled" by the block
	// itself. As the only reason to pass receipts here is for later settlement
	// in another block, there is no need to pass anything meaningful as it
	// would also require them to be received as an argument to MarkSynchronous.
}

// Target, excess, and config _after_ are a requirement of
// [Block.MarkExecuted].

// WaitUntilSettled blocks until either [Block.MarkSettled] is called or the
// [context.Context] is cancelled.
func (b *Block) WaitUntilSettled(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Settled reports whether either of [Block.MarkSettled] or
// [Block.MarkSynchronous] have been called without resulting in an error.
func (b *Block) Settled() bool { _ = "STUB: not implemented"; return false }

// Synchronous reports whether [Block.MarkSynchronous] has been called without
// resulting in an error.
func (b *Block) Synchronous() bool { _ = "STUB: not implemented"; return false }

func (b *Block) ancestor(ifSettledErrMsg string, get func(*ancestry) *Block) *Block {
	_ = "STUB: not implemented"
	return nil
}

const (
	getParentOfSettledErrMsg  = "Get parent of settled block"
	getSettledOfSettledErrMsg = "Get last-settled of settled block"
)

// ParentBlock returns the block's parent unless [Block.MarkSettled] has been
// called, in which case it returns nil and logs an error.
func (b *Block) ParentBlock() *Block { _ = "STUB: not implemented"; return nil }

// LastSettled returns the last-settled block at the time of b's acceptance,
// unless [Block.MarkSettled] has been called, in which case it returns nil and
// logs an error. If [Block.MarkSynchronous] was called instead, LastSettled
// always returns `b` itself, without logging. Note that this value might not be
// distinct between contiguous blocks.
func (b *Block) LastSettled() *Block { _ = "STUB: not implemented"; return nil }

// Settles returns the executed blocks that b settles if it is accepted by
// consensus. If `x` is the block height of the `b.ParentBlock().LastSettled()`
// and `y` is the height of the `b.LastSettled()`, then Settles returns the
// contiguous, half-open range (x,y] or an empty slice i.f.f. x==y. Every block
// therefore returns a disjoint (and possibly empty) set of historical blocks.
//
// It is not valid to call Settles after a call to [Block.MarkSettled] on either
// b or its parent. If [Block.MarkSynchronous] was called instead, Settles
// always returns a single-element slice of `b` itself.
func (b *Block) Settles() []*Block { _ = "STUB: not implemented"; return nil }

// Range returns the blocks in the continuous, half-open interval (start, end]
// in order of increasing height.
//
// The `start` block MAY be settled, but all other blocks in the range MUST NOT
// be settled. It is assumed that `start` can be reached by traversing up the
// chain from `end`.
//
// If the two arguments are the same block, Range returns an empty slice.
func Range(start, end *Block) []*Block { _ = "STUB: not implemented"; return nil }

var errIncompleteBlockHistory = errors.New("incomplete block history when determining last-settled block")

// LastToSettleAt returns (a) the last block to be settled at time `settleAt` if
// building on the specified parent block, and (b) a boolean to indicate if
// settlement is currently possible. If the returned boolean is false, the
// execution stream is lagging and LastToSettleAt MAY be called again after some
// indeterminate delay.
//
// It is not valid to call LastToSettleAt with a parent on which
// [Block.MarkSettled] was called directly (i.e. [Block.MarkSynchronous] does
// not preclude the parent from usage here).
//
// See the Example for [Block.WhenChildSettles] for one usage of the returned
// block.
func LastToSettleAt(hooks hook.Points, settleAt time.Time, parent *Block) (b *Block, ok bool, _ error) {
	_ = "STUB: not implemented"

	// Avoids having to perform this check at every return.
	return nil, false, nil
}

// A block can be the last to settle at some time i.f.f. two criteria are
// met:
//
// 1. The block has finished execution by said time and;
//
// 2. The block's child is known to have *not* finished execution or be
//    unable to finish by that time.
//
// The block currently being built can never finish in time, so we start
// with criterion (2) being met.

// The only way [Block.ParentBlock] can be nil is if `block` was already
// settled (see invariant in [Block]). If a block was already settled then
// only it or a later (i.e. unsettled) block can be returned by this loop,
// therefore we have a guarantee that the loop update will never result in
// `block==nil`.

// Although the below [Block.Settled] check (performed in the last
// loop iteration) precludes this from happening, that assumes no
// settlement concurrently with a call to [LastToSettleAt]. While
// that may be true now, the consequence of a race condition when
// omitting this check would be a panic for nil-pointer
// dereferencing.

// Guarantees that the loop will always exit as the last pre-SAE block
// (perhaps the genesis) is always settled, by definition.

// There may have been a race between this check and the
// interim-execution one above, so we have to check again.

// TODO(arr4n) more fine-grained checks are possible by computing the
// minimum possible gas consumption of blocks. For example,
// `block.BuildTime()+block.intrinsicGasSum()` can be compared against
// `equivSettleAt`, as can the sum of a chain of blocks.

// Note that a grandchild block having unknown execution completion time
// does not rule out knowing a child's completion time, so this could be
// set to true in a future loop iteration.
