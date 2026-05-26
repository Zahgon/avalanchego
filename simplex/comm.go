// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"errors"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/snow/networking/sender"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_               simplex.Communication = (*Comm)(nil)
	errNodeNotFound                       = errors.New("node not found in the validator list")
)

type Comm struct {
	logger   simplex.Logger
	subnetID ids.ID
	chainID  ids.ID
	// broadcastNodes are the nodes that should receive broadcast messages
	broadcastNodes set.Set[ids.NodeID]
	// allNodes are the IDs of all the nodes in the subnet
	allNodes []simplex.NodeID

	// sender is used to send messages to other nodes
	sender     sender.ExternalSender
	msgBuilder message.OutboundMsgBuilder
}

func NewComm(config *Config) (*Comm, error) { _ = "STUB: not implemented"; return nil, nil }

// grab all the nodes that are validators for the subnet

// skip our own node ID

func (c *Comm) Nodes() []simplex.NodeID { _ = "STUB: not implemented"; return nil }

func (c *Comm) Send(msg *simplex.Message, destination simplex.NodeID) {
	_ = "STUB: not implemented"
	return
}

func (c *Comm) Broadcast(msg *simplex.Message) { _ = "STUB: not implemented"; return }

func (c *Comm) simplexMessageToOutboundMessage(msg *simplex.Message) (*message.OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
