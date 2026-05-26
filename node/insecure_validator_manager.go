// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/version"
)

type insecureValidatorManager struct {
	router.ExternalHandler
	log    logging.Logger
	vdrs   validators.Manager
	weight uint64
}

func (i *insecureValidatorManager) Connected(vdrID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// Sybil protection is disabled so we don't have a txID that added the
// peer as a validator. Because each validator needs a txID associated
// with it, we hack one together by padding the nodeID with zeroes.

func (i *insecureValidatorManager) Disconnected(vdrID ids.NodeID) {
	_ = "STUB: not implemented"
	// RemoveWeight will only error here if there was an error reported during
	// Add.
	return
}
