// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handler

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/networking/tracker"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"

	commontracker "github.com/ava-labs/avalanchego/snow/engine/common/tracker"
)

const (
	numDispatchersToClose = 3
	// If a consensus message takes longer than this to process, the handler
	// will log a warning.
	syncProcessingTimeWarnLimit = 30 * time.Second
)

var (
	_ Handler = (*handler)(nil)

	errMissingEngine  = errors.New("missing engine")
	errNoStartingGear = errors.New("failed to select starting gear")
	errorShuttingDown = errors.New("shutting down")
)

type Handler interface {
	health.Checker

	Context() *snow.ConsensusContext
	// ShouldHandle returns true if the node with the given ID is allowed to send
	// messages to this chain. If the node is not allowed to send messages to
	// this chain, the message should be dropped.
	ShouldHandle(nodeID ids.NodeID) bool

	SetEngineManager(engineManager *EngineManager)
	GetEngineManager() *EngineManager

	SetOnStopped(onStopped func())
	Start(ctx context.Context, recoverPanic bool)
	Push(ctx context.Context, msg Message)
	Len() int

	Stop(ctx context.Context)
	StopWithError(ctx context.Context, err error)
	// AwaitStopped returns an error if the call would block and [ctx] is done.
	// Even if [ctx] is done when passed into this function, this function will
	// return a nil error if it will not block.
	AwaitStopped(ctx context.Context) (time.Duration, error)
}

// handler passes incoming messages from the network to the consensus engine.
// (Actually, it receives the incoming messages from a ChainRouter, but same difference.)
type handler struct {
	haltBootstrapping func()

	metrics *metrics

	nf           *common.NotificationForwarder
	subscription common.Subscription
	cn           *block.ChangeNotifier

	// Useful for faking time in tests
	clock mockable.Clock

	ctx *snow.ConsensusContext
	// TODO: consider using peerTracker instead of validators
	// since peerTracker is already tracking validators
	validators validators.Manager
	// Receives messages from the VM
	msgFromVMChan chan common.Message
	// How often to fire an internal gossip request to the engine. A value of 0
	// disables the gossip ticker so [engine.Gossip] is never called.
	gossipFrequency time.Duration

	engineManager *EngineManager

	// onStopped is called in a goroutine when this handler finishes shutting
	// down. If it is nil then it is skipped.
	onStopped func()

	// Tracks cpu/disk usage caused by each peer.
	resourceTracker tracker.ResourceTracker

	// Holds messages that [engine] hasn't processed yet.
	// [unprocessedMsgsCond.L] must be held while accessing [syncMessageQueue].
	syncMessageQueue MessageQueue
	// Holds messages that [engine] hasn't processed yet.
	// [unprocessedAsyncMsgsCond.L] must be held while accessing [asyncMessageQueue].
	asyncMessageQueue MessageQueue
	// Worker pool for handling asynchronous consensus messages
	asyncMessagePool errgroup.Group

	closeOnce            sync.Once
	startClosingTime     time.Time
	totalClosingTime     time.Duration
	closingChan          chan struct{}
	numDispatchersClosed atomic.Uint32
	// Closed when this handler and [engine] are done shutting down
	closed chan struct{}

	subnet subnets.Subnet

	// Tracks the peers that are currently connected to this subnet
	peerTracker commontracker.Peers
	p2pTracker  *p2p.PeerTracker
}

// Initialize this consensus handler
// [engine] must be initialized before initializing this handler
func New(
	ctx *snow.ConsensusContext,
	cn *block.ChangeNotifier,
	subscription common.Subscription,
	validators validators.Manager,
	gossipFrequency time.Duration,
	threadPoolSize int,
	resourceTracker tracker.ResourceTracker,
	subnet subnets.Subnet,
	peerTracker commontracker.Peers,
	p2pTracker *p2p.PeerTracker,
	reg prometheus.Registerer,
	haltBootstrapping func(),
) (Handler, error) {
	_ = "STUB: not implemented"
	return *new(Handler), nil
}

func (h *handler) Context() *snow.ConsensusContext { _ = "STUB: not implemented"; return nil }

func (h *handler) ShouldHandle(nodeID ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (h *handler) SetEngineManager(engineManager *EngineManager) { _ = "STUB: not implemented"; return }

func (h *handler) GetEngineManager() *EngineManager { _ = "STUB: not implemented"; return nil }

func (h *handler) SetOnStopped(onStopped func()) { _ = "STUB: not implemented"; return }

func (h *handler) selectStartingGear(ctx context.Context) (common.Engine, error) {
	_ = "STUB: not implemented"
	return *new(common.Engine), nil
}

// drop bootstrap state from previous runs before starting state sync

func (h *handler) Start(ctx context.Context, recoverPanic bool) { _ = "STUB: not implemented"; return }

func (h *handler) Notify(ctx context.Context, msg common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Push the message onto the handler's queue
func (h *handler) Push(ctx context.Context, msg Message) { _ = "STUB: not implemented"; return }

func (h *handler) Len() int { _ = "STUB: not implemented"; return 0 }

// Note: It is possible for Stop to be called before/concurrently with Start.
//
// Invariant: Stop must never block.
func (h *handler) Stop(_ context.Context) { _ = "STUB: not implemented"; return }

// Must hold the locks here to ensure there's no race condition in where
// we check the value of [h.closing] after the call to [Signal].

func (h *handler) StopWithError(ctx context.Context, err error) { _ = "STUB: not implemented"; return }

func (h *handler) AwaitStopped(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (h *handler) dispatchSync(ctx context.Context) { _ = "STUB: not implemented"; return }

// Handle sync messages from the router

// Get the next message we should process. If the handler is shutting
// down, we may fail to pop a message.

// If there is an error handling the message, shut down the chain

func (h *handler) dispatchAsync(ctx context.Context) {
	_ = "STUB: not implemented"

	// We never return an error in any of our functions, so it is safe to
	// drop any error here.
	return
}

// Handle async messages from the router

// Get the next message we should process. If the handler is shutting
// down, we may fail to pop a message.

func (h *handler) dispatchChans(ctx context.Context) { _ = "STUB: not implemented"; return }

// Handle messages generated by the handler and the VM

// Any returned error is treated as fatal
func (h *handler) handleSyncMsg(ctx context.Context, msg Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the chain is in normal operation at the start of message
// execution (may change during execution)

// We will attempt to pass the message to the requested type for the state
// we are currently in.

// The peer is requesting an engine type that hasn't been initialized
// yet. This means we know that this isn't a response, so we can safely
// drop the message.

// The peer is requesting an engine type that has been initialized, so
// we should attempt to honor the request.

// Note: [msg.EngineType] may have been provided by the peer as an
// invalid option. I.E. not one of DAG, CHAIN, or UNSPECIFIED.
// In this case, we treat the value the same way as UNSPECIFIED.
//
// If the peer didn't request a specific engine type, we default to the
// current engine.

// This should only happen if the peer is not following the protocol.
// This can happen if the chain only has a Snowman engine and the peer
// requested an Avalanche engine handle the message.

// Invariant: Response messages can never be dropped here. This is because
//            the timeout has already been cleared. This means the engine
//            should be invoked with a failure message if parsing of the
//            response fails.

// State messages should always be sent to the snowman engine

// Bootstrapping messages may be forwarded to either avalanche or snowman
// engines, depending on the EngineType field

// Connection messages can be sent to the currently executing engine

func (h *handler) handleAsyncMsg(ctx context.Context, msg Message) {
	_ = "STUB: not implemented"
	return
}

// Any returned error is treated as fatal
func (h *handler) executeAsyncMsg(ctx context.Context, msg Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Any returned error is treated as fatal
func (h *handler) handleChanMsg(msg *message.InboundMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the chain is in normal operation at the start of message
// execution (may change during execution)

func (h *handler) popUnexpiredMsg(queue MessageQueue) (context.Context, Message, bool) {
	_ = "STUB: not implemented"

	// Get the next message we should process. If the handler is shutting
	// down, we may fail to pop a message.
	return *new(context.Context), *new(Message), false
}

// If this message's deadline has passed, don't process it.

// Invariant: if closeDispatcher is called, Stop has already been called.
func (h *handler) closeDispatcher(ctx context.Context) { _ = "STUB: not implemented"; return }

// Note: shutdown is only called after all message dispatchers have exited or if
// no message dispatchers ever started.
func (h *handler) shutdown(ctx context.Context, startClosingTime time.Time) {
	_ = "STUB: not implemented"
	// If we are shutting down but haven't properly started, we don't need to
	// close the notification forwarder.
	return
}
