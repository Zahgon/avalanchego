// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/acp118"
	"github.com/ava-labs/avalanchego/proto/pb/platformvm"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
)

const (
	ErrFailedToParseWarpAddressedCall = iota + 1
	ErrWarpAddressedCallHasSourceAddress
	ErrFailedToParseWarpAddressedCallPayload
	ErrUnsupportedWarpAddressedCallPayloadType

	ErrFailedToParseJustification
	ErrConversionDoesNotExist
	ErrMismatchedConversionID

	ErrInvalidJustificationType
	ErrFailedToParseSubnetID
	ErrMismatchedValidationID
	ErrValidationDoesNotExist
	ErrValidationExists
	ErrFailedToParseRegisterL1Validator
	ErrValidationCouldBeRegistered

	ErrImpossibleNonce
	ErrWrongNonce
	ErrWrongWeight
)

var _ acp118.Verifier = (*signatureRequestVerifier)(nil)

type signatureRequestVerifier struct {
	stateLock sync.Locker
	state     state.Chain
}

func (s signatureRequestVerifier) Verify(
	_ context.Context,
	unsignedMessage *warp.UnsignedMessage,
	justification []byte,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

func (s signatureRequestVerifier) verifySubnetToL1Conversion(
	msg *message.SubnetToL1Conversion,
	justification []byte,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

func (s signatureRequestVerifier) verifyL1ValidatorRegistration(
	msg *message.L1ValidatorRegistration,
	justificationBytes []byte,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

// verifyL1ValidatorRegistered verifies that the validationID is currently a
// validator.
func (s signatureRequestVerifier) verifyL1ValidatorRegistered(
	validationID ids.ID,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

// Verify that the validator exists

// verifySubnetValidatorNotCurrentlyRegistered verifies that the validationID
// could only correspond to a validator from a ConvertSubnetToL1Tx and that it
// is not currently a validator.
func (s signatureRequestVerifier) verifySubnetValidatorNotCurrentlyRegistered(
	validationID ids.ID,
	justification *platformvm.SubnetIDIndex,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

// Verify that the provided subnetID has been converted.

// Verify that the validator is not in the current state

// Either the validator was removed or it was never registered as part of
// the subnet conversion.

// verifySubnetValidatorCanNotValidate verifies that the validationID is not
// currently and can never become a validator.
func (s signatureRequestVerifier) verifySubnetValidatorCanNotValidate(
	validationID ids.ID,
	justificationBytes []byte,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

// Verify that the validator does not currently exist

// The expiry time has passed

// If the validation ID was successfully registered and then removed, it can
// never be re-used again even if its expiry has not yet passed.

// The validator has been removed

func (s signatureRequestVerifier) verifyL1ValidatorWeight(
	msg *message.L1ValidatorWeight,
) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

// If the peer is attempting to verify that the weight of the validator
// is 0, they should be requesting a [message.L1ValidatorRegistration]
// with Registered set to false.

// The nonce and weight are correct
