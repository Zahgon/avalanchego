// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"net/http"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm/client"
)

type ValidatorsAPI struct {
	vm *VM
}

func (api *ValidatorsAPI) GetCurrentValidators(httpReq *http.Request, req *client.GetCurrentValidatorsRequest, reply *client.GetCurrentValidatorsResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// Filter by requested nodeIDs if specified

// Skip if specific nodeIDs were requested and this isn't one of them

// Transform this to a percentage (0-100) to make it consistent
// with currentValidators in PlatformVM API
