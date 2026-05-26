// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snow

import (
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var (
	_ Acceptor = acceptorWrapper{}

	_ AcceptorGroup = (*acceptorGroup)(nil)
)

// Acceptor is implemented when a struct is monitoring if a message is accepted
type Acceptor interface {
	// Accept must be called before [containerID] is committed to the VM as
	// accepted.
	//
	// If the returned error is non-nil, the chain associated with [ctx] should
	// shut down and not commit [container] or any other container to its
	// database as accepted.
	Accept(ctx *ConsensusContext, containerID ids.ID, container []byte) error
}

type acceptorWrapper struct {
	Acceptor

	// If true and Accept returns an error, the chain this callback corresponds
	// to will stop.
	dieOnError bool
}

type AcceptorGroup interface {
	// Calling Accept() calls all of the registered acceptors for the relevant
	// chain.
	Acceptor

	// RegisterAcceptor causes [acceptor] to be called every time an operation
	// is accepted on chain [chainID].
	// If [dieOnError], chain [chainID] stops if Accept returns a non-nil error.
	RegisterAcceptor(chainID ids.ID, acceptorName string, acceptor Acceptor, dieOnError bool) error

	// DeregisterAcceptor removes an acceptor from the group.
	DeregisterAcceptor(chainID ids.ID, acceptorName string) error
}

type acceptorGroup struct {
	log logging.Logger

	lock sync.RWMutex
	// Chain ID --> Acceptor Name --> Acceptor
	acceptors map[ids.ID]map[string]acceptorWrapper
}

func NewAcceptorGroup(log logging.Logger) AcceptorGroup {
	_ = "STUB: not implemented"
	return *new(AcceptorGroup)
}

func (a *acceptorGroup) Accept(ctx *ConsensusContext, containerID ids.ID, container []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptorGroup) RegisterAcceptor(chainID ids.ID, acceptorName string, acceptor Acceptor, dieOnError bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptorGroup) DeregisterAcceptor(chainID ids.ID, acceptorName string) error {
	_ = "STUB: not implemented"
	return nil
}
