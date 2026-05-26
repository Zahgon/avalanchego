// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"context"
	"io"
	"sync/atomic"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/txpool/legacypool"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/params"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/sae/rpc"
	"github.com/ava-labs/avalanchego/vms/saevm/saedb"
	"github.com/ava-labs/avalanchego/vms/saevm/saexec"
	"github.com/ava-labs/avalanchego/vms/saevm/txgossip"

	snowcommon "github.com/ava-labs/avalanchego/snow/engine/common"
	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// VM implements all of [adaptor.ChainVM] except for the `Initialize` method,
// which needs to be provided by a harness. In all cases, the harness MUST
// provide a last-synchronous block, which MAY be the genesis.
type VM struct {
	*p2p.Network
	Peers          *p2p.Peers
	ValidatorPeers *p2p.Validators

	hooks   hook.Points
	config  Config
	snowCtx *snow.Context
	metrics *prometheus.Registry

	db  ethdb.Database
	xdb saetypes.ExecutionResults

	consensusState utils.Atomic[snow.State]

	preference atomic.Pointer[blocks.Block]
	last       struct {
		accepted, settled atomic.Pointer[blocks.Block]
		synchronous       uint64
	}
	acceptedBlocks event.FeedOf[*blocks.Block]
	// Consensus-critical blocks are those either (a) undergoing a consensus
	// decision; or (b) informing consensus invariants (e.g. artefacts to
	// settle). The latter is defined as the history of accepted blocks up to,
	// and including, the last-settled block.
	consensusCritical *syncMap[common.Hash, *blocks.Block]

	exec         *saexec.Executor
	mempool      *txgossip.Set
	blockBuilder blockBuilder
	rpcProvider  *rpc.Provider
	newTxs       chan struct{}

	// toClose are closed in reverse order during [VM.Shutdown]. If a resource
	// depends on another resource, it MUST be added AFTER the resource it
	// depends on.
	toClose []io.Closer
}

// closerFunc adapts a func() error to [io.Closer].
type closerFunc func() error

var _ io.Closer = (*closerFunc)(nil)

func (f closerFunc) Close() error {
	_ = "STUB: not implemented"

	// A Config configures construction of a new [VM].
	return nil
}

type Config struct {
	MempoolConfig legacypool.Config
	DBConfig      saedb.Config
	RPCConfig     rpc.Config

	ExcessAfterLastSynchronous gas.Gas

	Now func() time.Time // defaults to [time.Now] if nil
}

// NewVM returns a new [VM] that is ready for use immediately upon return.
// [VM.Shutdown] MUST be called to release resources.
//
// The state root of the last synchronous block MUST be available when creating
// a [triedb.Database] from the provided [ethdb.Database] and [triedb.Config]
// (the latter provided via the [Config]).
func NewVM[T hook.Transaction](
	ctx context.Context,
	hooks hook.PointsG[T],
	cfg Config,
	snowCtx *snow.Context,
	chainConfig *params.ChainConfig,
	db ethdb.Database,
	lastSynchronous *types.Block,
	sender snowcommon.AppSender,
) (_ *VM, retErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ==========  Sync -> Async  ==========

// ==========  Block State  ==========

// ==========  Mempool  ==========

// ==========  Block Builder  ==========

// ==========  P2P Gossip  ==========

// ==========  RPC Provider  ==========

// canonicaliseLastSynchronous writes all necessary information to the database
// to have the block be considered accepted/canonical by SAE. If there are any
// canonical blocks at a height greater than the provided block then this
// function is a no-op, which makes it effectively idempotent with respect to
// the rest of SAE processing.
func canonicaliseLastSynchronous(db ethdb.Database, block *blocks.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// If any other block has been accepted then the last synchronous block
// must have been canonicalised in a previous initialisation.

// signalNewTxsToEngine subscribes to the [txpool.TxPool] to unblock
// [VM.WaitForEvent] when necessary. [VM.Shutdown] MUST be called to release a
// goroutine started by this method.
func (vm *VM) signalNewTxsToEngine() { _ = "STUB: not implemented"; return }

/*reorgs but ignored by legacypool*/

// guaranteed to be closed due to unsubscribing

// See [VM.WaitForEvent] for why this requires a buffer.

// coverage visualisation

// coverage visualization

// WaitForEvent returns immediately if there are already pending transactions in
// the mempool, otherwise it blocks until the mempool notifies it of new
// transactions. In both cases it returns [snowcommon.PendingTxs]. In the latter
// scenario it respects context cancellation.
func (vm *VM) WaitForEvent(ctx context.Context) (snowcommon.Message, error) {
	_ = "STUB: not implemented"
	return *new(snowcommon.Message), nil
}

// probably has something buffered

// Sends on the `newTxs` channel are performed on a best-effort basis, which
// could race here if it weren't for the channel buffer.

func (vm *VM) numPendingTxs() int { _ = "STUB: not implemented"; return 0 }

// SetState notifies the VM of a transition in the state lifecycle.
func (vm *VM) SetState(ctx context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown gracefully closes the VM.
func (vm *VM) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (vm *VM) close() error { _ = "STUB: not implemented"; return nil }

// Version reports the VM's version.
func (*VM) Version(context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (vm *VM) log() logging.Logger { _ = "STUB: not implemented"; return *new(logging.Logger) }
