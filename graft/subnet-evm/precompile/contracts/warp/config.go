// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
)

const (
	WarpDefaultQuorumNumerator uint64 = 67
	WarpQuorumNumeratorMinimum uint64 = 33
	WarpQuorumDenominator      uint64 = 100
)

var (
	_ precompileconfig.Config     = (*Config)(nil)
	_ precompileconfig.Predicater = (*Config)(nil)
	_ precompileconfig.Accepter   = (*Config)(nil)
)

var (
	ErrInvalidQuorumRatio         = errors.New("invalid warp quorum ratio")
	errOverflowSignersGasCost     = errors.New("overflow calculating warp signers gas cost")
	errInvalidPredicateBytes      = errors.New("cannot unpack predicate bytes")
	errInvalidWarpMsg             = errors.New("cannot unpack warp message")
	errCannotParseWarpMsg         = errors.New("cannot parse warp message")
	errInvalidWarpMsgPayload      = errors.New("cannot unpack warp message payload")
	errInvalidAddressedPayload    = errors.New("cannot unpack addressed payload")
	errInvalidBlockHashPayload    = errors.New("cannot unpack block hash payload")
	errCannotGetNumSigners        = errors.New("cannot fetch num signers from warp message")
	errWarpCannotBeActivated      = errors.New("warp cannot be activated before Durango")
	errFailedVerification         = errors.New("cannot verify warp signature")
	errCannotRetrieveValidatorSet = errors.New("cannot retrieve validator set")
)

// Config implements the precompileconfig.Config interface and
// adds specific configuration for Warp.
type Config struct {
	precompileconfig.Upgrade
	QuorumNumerator              uint64 `json:"quorumNumerator"`
	RequirePrimaryNetworkSigners bool   `json:"requirePrimaryNetworkSigners"`
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// Warp with the given quorum numerator.
func NewConfig(blockTimestamp *uint64, quorumNumerator uint64, requirePrimaryNetworkSigners bool) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultConfig returns a config for a network upgrade at [blockTimestamp] that enables
// Warp with the default quorum numerator (0 denotes using the default).
func NewDefaultConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables Warp.
func NewDisableConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

// Key returns the key for the Warp precompileconfig.
// This should be the same key as used in the precompile module.
func (*Config) Key() string {
	_ = "STUB: not implemented"

	// Verify tries to verify Config and returns an error accordingly.
	return ""
}

func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	_ = "STUB: not implemented"
	return nil

	// If Warp attempts to activate before Durango, fail verification
}

// If a non-default quorum numerator is specified and it is less than the minimum, return an error

// Equal returns true if [s] is a [*Config] and it has been configured identical to [c].
func (c *Config) Equal(s precompileconfig.Config) bool {
	_ = "STUB: not implemented"
	// typecast before comparison
	return false
}

//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
func (*Config) Accept(acceptCtx *precompileconfig.AcceptContext, blockHash common.Hash, blockNumber uint64, txHash common.Hash, logIndex int, topics []common.Hash, logData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PredicateGas returns the amount of gas necessary to verify the predicate
// PredicateGas charges for:
// 1. Base cost of the message
// 2. Size of the message
// 3. Number of signers
// 4. TODO: Lookup of the validator set
//
// If the payload of the warp message fails parsing, return a non-nil error invalidating the transaction.
func (*Config) PredicateGas(pred predicate.Predicate, rules precompileconfig.Rules) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// VerifyPredicate returns whether the predicate described by [predicateBytes] passes verification.
func (c *Config) VerifyPredicate(predicateContext *precompileconfig.PredicateContext, pred predicate.Predicate) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: PredicateGas should be called before VerifyPredicate, so we should never reach an error case here.

// For the X-chain and the C-chain, chains can be configured not to
// require the primary network validators to have signed the warp
// message and to use the, likely smaller, local subnet's validator set.
//
// The primary network validator set is never required when verifying
// messages from the P-chain because the P-chain is always synced.
