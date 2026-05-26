// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saetest

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// A KeyChain manages a set of private keys (suitable only for tests) to sign
// transactions.
type KeyChain struct {
	keys  []*ecdsa.PrivateKey
	addrs []common.Address
}

// NewUNSAFEKeyChain returns a new key chain with the specified number of
// accounts. Private keys are generated deterministically.
func NewUNSAFEKeyChain(tb testing.TB, accounts uint) *KeyChain {
	_ = "STUB: not implemented"
	return nil
}

// SignTx returns [types.SignNewTx], called with `data` and the respective
// `account` key.
func (kc *KeyChain) SignTx(tb testing.TB, signer types.Signer, account int, data types.TxData) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// A Wallet manages a set of private keys (suitable only for tests) and nonces
// to sign transactions.
type Wallet struct {
	*KeyChain
	nonces []uint64 // MUST have same length as `kc.keys`
	signer types.Signer
}

// NewUNSAFEWallet returns a new wallet with the specified number of accounts.
// Private keys are generated deterministically.
func NewUNSAFEWallet(tb testing.TB, accounts uint, signer types.Signer) *Wallet {
	_ = "STUB: not implemented"
	return nil
}

// NewWalletWithKeyChain returns a new wallet, backed by the provided key chain.
func NewWalletWithKeyChain(kc *KeyChain, signer types.Signer) *Wallet {
	_ = "STUB: not implemented"
	return nil
}

// Addresses returns all addresses managed by the key chain.
func (kc *KeyChain) Addresses() []common.Address { _ = "STUB: not implemented"; return nil }

// SetNonceAndSign overrides the nonce in the `data` with the next one for the
// account, then signs and returns the transaction. The wallet's record of the
// account nonce begins at zero and increments after every successful call to
// this method.
func (w *Wallet) SetNonceAndSign(tb testing.TB, account int, data types.TxData) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// DecrementNonce decrements the nonce of the specified account. This is useful
// for retrying transactions with updated parameters.
func (w *Wallet) DecrementNonce(tb testing.TB, account int) { _ = "STUB: not implemented"; return }

// MaxAllocFor returns a genesis allocation with [MaxUint256] as the balance for
// all addresses provided.
func MaxAllocFor(addrs ...common.Address) types.GenesisAlloc {
	_ = "STUB: not implemented"
	return *new(types.GenesisAlloc)
}
