// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package bind

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

const basefeeWiggleMultiplier = 2

var (
	errNoEventSignature       = errors.New("no event signature")
	errEventSignatureMismatch = errors.New("event signature mismatch")
)

// SignerFn is a signer function callback when a contract requires a method to
// sign the transaction before submission.
type SignerFn func(common.Address, *types.Transaction) (*types.Transaction, error)

// CallOpts is the collection of options to fine tune a contract call request.
type CallOpts struct {
	Accepted    bool            // Whether to operate on the accepted state or the last known one
	From        common.Address  // Optional the sender address, otherwise the first account is used
	BlockNumber *big.Int        // Optional the block number on which the call should be performed
	BlockHash   common.Hash     // Optional the block hash on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// TransactOpts is the collection of authorization data required to create a
// valid Ethereum transaction.
type TransactOpts struct {
	From   common.Address // Ethereum account to send the transaction from
	Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
	Signer SignerFn       // Method to use for signing the transaction (mandatory)

	Value     *big.Int // Funds to transfer along the transaction (nil = 0 = no funds)
	GasPrice  *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasFeeCap *big.Int // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	GasTipCap *big.Int // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	GasLimit  uint64   // Gas limit to set for the transaction execution (0 = estimate)

	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)

	NoSend bool // Do all transact steps but do not send the transaction
}

// FilterOpts is the collection of options to fine tune filtering for events
// within a bound contract.
type FilterOpts struct {
	Start uint64  // Start of the queried range
	End   *uint64 // End of the range (nil = latest)

	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// WatchOpts is the collection of options to fine tune subscribing for events
// within a bound contract.
type WatchOpts struct {
	Start   *uint64         // Start of the queried range (nil = latest)
	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// MetaData collects all metadata for a bound contract.
type MetaData struct {
	mu   sync.Mutex
	Sigs map[string]string
	Bin  string
	ABI  string
	ab   *abi.ABI
}

func (m *MetaData) GetAbi() (*abi.ABI, error) { _ = "STUB: not implemented"; return nil, nil }

// BoundContract is the base wrapper object that reflects a contract on the
// Ethereum network. It contains a collection of methods that are used by the
// higher level contract bindings to operate.
type BoundContract struct {
	address    common.Address     // Deployment address of the contract on the Ethereum blockchain
	abi        abi.ABI            // Reflect based ABI to access the correct Ethereum methods
	caller     ContractCaller     // Read interface to interact with the blockchain
	transactor ContractTransactor // Write interface to interact with the blockchain
	filterer   ContractFilterer   // Event filtering to interact with the blockchain
}

// NewBoundContract creates a low level contract interface through which calls
// and transactions may be made through.
func NewBoundContract(address common.Address, abi abi.ABI, caller ContractCaller, transactor ContractTransactor, filterer ContractFilterer) *BoundContract {
	_ = "STUB: not implemented"
	return nil
}

// DeployContract deploys a contract onto the Ethereum blockchain and binds the
// deployment address with a Go wrapper.
func DeployContract(opts *TransactOpts, abi abi.ABI, bytecode []byte, backend ContractBackend, params ...interface{}) (common.Address, *types.Transaction, *BoundContract, error) {
	_ = "STUB: not implemented"
	// Otherwise try to deploy the contract
	return *new(common.Address), nil, nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (c *BoundContract) Call(opts *CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	// Don't crash on a lazy user
	return nil
}

// Pack the input, call and unpack the results

// Make sure we have a contract to operate on, and bail out otherwise.

// Make sure we have a contract to operate on, and bail out otherwise.

// Make sure we have a contract to operate on, and bail out otherwise.

// Transact invokes the (paid) contract method with params as input values.
func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	// Otherwise pack up the parameters and invoke the contract
	return nil, nil
}

// todo(rjl493456442) check whether the method is payable or not,
// reject invalid transaction at the first place

// RawTransact initiates a transaction with the given raw calldata as the input.
// It's usually used to initiate transactions for invoking **Fallback** function.
func (c *BoundContract) RawTransact(opts *TransactOpts, calldata []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	// todo(rjl493456442) check whether the method is payable or not,
	// reject invalid transaction at the first place
	return nil, nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (c *BoundContract) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	// todo(rjl493456442) check the payable fallback or receive is defined
	// or not, reject invalid transaction at the first place
	return nil, nil
}

func (c *BoundContract) createDynamicTx(opts *TransactOpts, contract *common.Address, input []byte, head *types.Header) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	// Normalize value
	return nil, nil
}

// Estimate TipCap

// Estimate FeeCap

// Estimate GasLimit

// create the transaction

func (c *BoundContract) createLegacyTx(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Normalize value

// Estimate GasPrice

// Estimate GasLimit

// create the transaction

func (c *BoundContract) estimateGasLimit(opts *TransactOpts, contract *common.Address, input []byte, gasPrice, gasTipCap, gasFeeCap, value *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Gas estimation cannot succeed without code for method invocations.
		nil
}

func (c *BoundContract) getNonce(opts *TransactOpts) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// transact executes an actual transaction invocation, first deriving any missing
// authorization fields, and then scheduling the transaction for execution.
func (c *BoundContract) transact(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the transaction

// Only query for basefee if gasPrice not specified

// Chain is not London ready -> use legacy transaction

// Sign the transaction and schedule it for execution

// FilterLogs filters contract logs for past blocks, returning the necessary
// channels to construct a strongly typed bound iterator on top of them.
func (c *BoundContract) FilterLogs(opts *FilterOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error) {
	_ = "STUB: not implemented"
	// Don't crash on a lazy user
	return nil, *new(event.Subscription), nil
}

// Append the event selector to the query parameters and construct the topic set

// Start the background filtering

/* TODO(karalabe): Replace the rest of the method below with this when supported
sub, err := c.filterer.SubscribeFilterLogs(ensureContext(opts.Context), config, logs)
*/

// WatchLogs filters subscribes to contract logs for future blocks, returning a
// subscription object that can be used to tear down the watcher.
func (c *BoundContract) WatchLogs(opts *WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error) {
	_ = "STUB: not implemented"
	// Don't crash on a lazy user
	return nil, *new(event.Subscription), nil
}

// Append the event selector to the query parameters and construct the topic set

// Start the background filtering

// UnpackLog unpacks a retrieved log into the provided output structure.
func (c *BoundContract) UnpackLog(out interface{}, event string, log types.Log) error {
	_ = "STUB: not implemented"
	// Anonymous events are not supported.
	return nil
}

// UnpackLogIntoMap unpacks a retrieved log into the provided map.
func (c *BoundContract) UnpackLogIntoMap(out map[string]interface{}, event string, log types.Log) error {
	_ = "STUB: not implemented"
	// Anonymous events are not supported.
	return nil
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
