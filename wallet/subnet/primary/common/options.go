// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	ethcommon "github.com/ava-labs/libevm/common"
)

const defaultPollFrequency = 100 * time.Millisecond

// IssuanceReceipt is the information known after issuing a transaction.
type IssuanceReceipt struct {
	// Identifies the primary chain ("P", "X" or "C")
	ChainAlias string
	// ID of the issued transaction
	TxID ids.ID
	// The time from initiation to issuance
	Duration time.Duration
}

// ConfirmationReceipt is the information known after issuing and confirming a
// transaction.
type ConfirmationReceipt struct {
	// Identifies the primary chain ("P", "X" or "C")
	ChainAlias string
	// ID of the issued transaction
	TxID ids.ID
	// The time from initiation to confirmation
	TotalDuration time.Duration
	// The time from issuance to confirmation. It does not include the duration
	// of issuance.
	ConfirmationDuration time.Duration
}

type Option func(*Options)

type Options struct {
	ctx context.Context

	customAddressesSet bool
	customAddresses    set.Set[ids.ShortID]

	customEthAddressesSet bool
	customEthAddresses    set.Set[ethcommon.Address]

	baseFee *big.Int

	minIssuanceTimeSet bool
	minIssuanceTime    uint64

	allowStakeableLocked bool

	changeOwner *secp256k1fx.OutputOwners

	memo []byte

	assumeDecided bool

	pollFrequencySet bool
	pollFrequency    time.Duration

	issuanceHandler     func(IssuanceReceipt)
	confirmationHandler func(ConfirmationReceipt)
}

func NewOptions(ops []Option) *Options { _ = "STUB: not implemented"; return nil }

func UnionOptions(first, second []Option) []Option { _ = "STUB: not implemented"; return nil }

func (o *Options) applyOptions(ops []Option) { _ = "STUB: not implemented"; return }

func (o *Options) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (o *Options) Addresses(defaultAddresses set.Set[ids.ShortID]) set.Set[ids.ShortID] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) EthAddresses(defaultAddresses set.Set[ethcommon.Address]) set.Set[ethcommon.Address] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) BaseFee(defaultBaseFee *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (o *Options) MinIssuanceTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (o *Options) AllowStakeableLocked() bool { _ = "STUB: not implemented"; return false }

func (o *Options) ChangeOwner(defaultOwner *secp256k1fx.OutputOwners) *secp256k1fx.OutputOwners {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) Memo() []byte { _ = "STUB: not implemented"; return nil }

func (o *Options) AssumeDecided() bool { _ = "STUB: not implemented"; return false }

func (o *Options) PollFrequency() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (o *Options) IssuanceHandler() func(IssuanceReceipt) { _ = "STUB: not implemented"; return nil }

func (o *Options) ConfirmationHandler() func(ConfirmationReceipt) {
	_ = "STUB: not implemented"
	return nil
}

func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCustomAddresses(addrs set.Set[ids.ShortID]) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCustomEthAddresses(addrs set.Set[ethcommon.Address]) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBaseFee(baseFee *big.Int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMinIssuanceTime(minIssuanceTime uint64) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithStakeableLocked() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithChangeOwner(changeOwner *secp256k1fx.OutputOwners) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMemo(memo []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAssumeDecided() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPollFrequency(pollFrequency time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithIssuanceHandler(f func(IssuanceReceipt)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithConfirmationHandler(f func(ConfirmationReceipt)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
