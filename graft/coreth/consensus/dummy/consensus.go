// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dummy

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/evm/acp226"
)

var (
	errUnclesUnsupported      = errors.New("uncles unsupported")
	errExtDataGasUsedNil      = errors.New("extDataGasUsed is nil")
	errExtDataGasUsedTooLarge = errors.New("extDataGasUsed is not uint64")
	ErrInvalidBlockGasCost    = errors.New("invalid blockGasCost")
	errInvalidExtDataGasUsed  = errors.New("invalid extDataGasUsed")
)

type Mode struct {
	ModeSkipHeader   bool
	ModeSkipBlockFee bool
	ModeSkipCoinbase bool
}

type (
	OnFinalizeAndAssembleCallbackType = func(
		header *types.Header,
		parent *types.Header,
		state *state.StateDB,
		txs []*types.Transaction,
	) (
		extraData []byte,
		blockFeeContribution *big.Int,
		extDataGasUsed *big.Int,
		err error,
	)

	OnExtraStateChangeType = func(
		block *types.Block,
		parent *types.Header,
		statedb *state.StateDB,
	) (
		blockFeeContribution *big.Int,
		extDataGasUsed *big.Int,
		err error,
	)

	ConsensusCallbacks struct {
		OnFinalizeAndAssemble OnFinalizeAndAssembleCallbackType
		OnExtraStateChange    OnExtraStateChangeType
	}

	DummyEngine struct {
		cb                  ConsensusCallbacks
		consensusMode       Mode
		desiredTargetExcess *gas.Gas
		desiredDelayExcess  *acp226.DelayExcess
	}
)

func NewDummyEngine(
	cb ConsensusCallbacks,
	mode Mode,
	desiredTargetExcess *gas.Gas, // Guides the target gas excess (ACP-176) toward the desired value
	desiredDelayExcess *acp226.DelayExcess, // Guides the min delay excess (ACP-226) toward the desired value
) *DummyEngine {
	_ = "STUB: not implemented"
	return nil
}

func NewETHFaker() *DummyEngine { _ = "STUB: not implemented"; return nil }

func NewFaker() *DummyEngine { _ = "STUB: not implemented"; return nil }

func NewFakerWithCallbacks(cb ConsensusCallbacks) *DummyEngine {
	_ = "STUB: not implemented"
	return nil
}

func NewFakerWithMode(cb ConsensusCallbacks, mode Mode) *DummyEngine {
	_ = "STUB: not implemented"
	return nil
}

func NewCoinbaseFaker() *DummyEngine { _ = "STUB: not implemented"; return nil }

func NewFullFaker() *DummyEngine { _ = "STUB: not implemented"; return nil }

func verifyHeaderGasFields(config *extras.ChainConfig, header *types.Header, parent *types.Header) error {
	_ = "STUB: not implemented"
	// Verifying the gas used occurs earlier in the block validation process in verifyIntrinsicGas, so
	// customheader.VerifyGasUsed is not called here.
	return nil
}

// Verify header.BaseFee matches the expected value.

// Enforce BlockGasCost constraints

// Verify ExtDataGasUsed not present before AP4

// ExtDataGasUsed correctness is checked during block validation
// (when the validator has access to the block contents)

// modified from consensus.go
func verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parent *types.Header, uncle bool) error {
	_ = "STUB: not implemented"
	// Ensure that we do not verify an uncle
	return nil
}

// Verify the extra data is well-formed.

// Ensure gas-related header fields are correct

// Verify that the block number is parent's +1

func (*DummyEngine) Author(header *types.Header) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (eng *DummyEngine) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header) error {
	_ = "STUB: not implemented"
	// If we're running a full engine faking, accept any input as valid
	return nil
}

// Short circuit if the header is known, or it's parent not

// Sanity checks passed, do a proper verification

func (*DummyEngine) VerifyUncles(_ consensus.ChainReader, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (*DummyEngine) Prepare(_ consensus.ChainHeaderReader, header *types.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (eng *DummyEngine) Finalize(chain consensus.ChainHeaderReader, block *types.Block, parent *types.Header, state *state.StateDB, receipts []*types.Receipt) error {
	_ = "STUB: not implemented"
	// Perform extra state change while finalizing the block
	return nil
}

// Verify the BlockGasCost set in the header matches the expected value.

// Validate extDataGasUsed and BlockGasCost match expectations
//
// NOTE: This is a duplicate check of what is already performed in
// blockValidator but is done here for symmetry with FinalizeAndAssemble.

// Verify the block fee was paid.

func (eng *DummyEngine) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, parent *types.Header, state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt,
) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate the required block gas cost for this block.

// Verify that this block covers the block fee.

// finalize the header.Extra

// Set the min delay excess

// commit the final state root

// Header seems complete, assemble into a block and return

func (*DummyEngine) CalcDifficulty(_ consensus.ChainHeaderReader, _ uint64, _ *types.Header) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (*DummyEngine) Close() error { _ = "STUB: not implemented"; return nil }
