// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ SetCallbackListener = (*logger)(nil)

type logger struct {
	log      logging.Logger
	subnetID ids.ID
	nodeIDs  set.Set[ids.NodeID]
}

// NewLogger returns a callback listener that will log validator set changes for
// the specified validators
func NewLogger(
	log logging.Logger,
	subnetID ids.ID,
	nodeIDs ...ids.NodeID,
) SetCallbackListener {
	_ = "STUB: not implemented"
	return *new(SetCallbackListener)
}

func (l *logger) OnValidatorAdded(
	nodeID ids.NodeID,
	pk *bls.PublicKey,
	txID ids.ID,
	weight uint64,
) {
	_ = "STUB: not implemented"
	return
}

func (l *logger) OnValidatorRemoved(
	nodeID ids.NodeID,
	weight uint64,
) {
	_ = "STUB: not implemented"
	return
}

func (l *logger) OnValidatorWeightChanged(
	nodeID ids.NodeID,
	oldWeight uint64,
	newWeight uint64,
) {
	_ = "STUB: not implemented"
	return
}
