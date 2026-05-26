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

package ethapi

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/libevm/accounts"
	"github.com/ava-labs/libevm/accounts/keystore"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/common/math"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
)

// estimateGasErrorRatio is the amount of overestimation eth_estimateGas is
// allowed to produce in order to speed up calculations.
const estimateGasErrorRatio = 0.015

var errBlobTxNotSupported = errors.New("signing blob transactions not supported")

// EthereumAPI provides an API to access Ethereum related information.
type EthereumAPI struct {
	b Backend
}

// NewEthereumAPI creates a new Ethereum protocol API.
func NewEthereumAPI(b Backend) *EthereumAPI { _ = "STUB: not implemented"; return nil }

// GasPrice returns a suggestion for a gas price for legacy transactions.
func (s *EthereumAPI) GasPrice(ctx context.Context) (*hexutil.Big, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BaseFee returns an estimate for what the base fee will be on the next block if
// it is produced now.
func (s *EthereumAPI) BaseFee(ctx context.Context) (*hexutil.Big, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MaxPriorityFeePerGas returns a suggestion for a gas tip cap for dynamic fee transactions.
func (s *EthereumAPI) MaxPriorityFeePerGas(ctx context.Context) (*hexutil.Big, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type feeHistoryResult struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

// FeeHistory returns the fee market history.
func (s *EthereumAPI) FeeHistory(ctx context.Context, blockCount math.HexOrDecimal64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*feeHistoryResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Syncing allows the caller to determine whether the chain is syncing or not.
// In geth, the response is either a map representing an ethereum.SyncProgress
// struct or "false" (indicating the chain is not syncing).
// In coreth, avalanchego prevents API calls unless bootstrapping is complete,
// so we always return false here for API compatibility.
func (s *EthereumAPI) Syncing() (interface{}, error) {
	_ = "STUB: not implemented"

	// TxPoolAPI offers and API for the transaction pool. It only operates on data that is non-confidential.
	return nil, nil
}

type TxPoolAPI struct {
	b Backend
}

// NewTxPoolAPI creates a new tx pool service that gives information about the transaction pool.
func NewTxPoolAPI(b Backend) *TxPoolAPI { _ = "STUB: not implemented"; return nil }

// Content returns the transactions contained within the transaction pool.
func (s *TxPoolAPI) Content() map[string]map[string]map[string]*RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// Flatten the pending transactions

// Flatten the queued transactions

// ContentFrom returns the transactions contained within the transaction pool.
func (s *TxPoolAPI) ContentFrom(addr common.Address) map[string]map[string]*RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// Build the pending transactions

// Build the queued transactions

// Status returns the number of pending and queued transaction in the pool.
func (s *TxPoolAPI) Status() map[string]hexutil.Uint { _ = "STUB: not implemented"; return nil }

// Inspect retrieves the content of the transaction pool and flattens it into an
// easily inspectable list.
func (s *TxPoolAPI) Inspect() map[string]map[string]map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Define a formatter to flatten a transaction into a string

// Flatten the pending transactions

// Flatten the queued transactions

// EthereumAccountAPI provides an API to access accounts managed by this node.
// It offers only methods that can retrieve accounts.
type EthereumAccountAPI struct {
	am *accounts.Manager
}

// NewEthereumAccountAPI creates a new EthereumAccountAPI.
func NewEthereumAccountAPI(am *accounts.Manager) *EthereumAccountAPI {
	_ = "STUB: not implemented"
	return nil
}

// Accounts returns the collection of accounts this node manages.
func (s *EthereumAccountAPI) Accounts() []common.Address { _ = "STUB: not implemented"; return nil }

// PersonalAccountAPI provides an API to access accounts managed by this node.
// It offers methods to create, (un)lock en list accounts. Some methods accept
// passwords and are therefore considered private by default.
type PersonalAccountAPI struct {
	am        *accounts.Manager
	nonceLock *AddrLocker
	b         Backend
}

// NewPersonalAccountAPI creates a new PersonalAccountAPI.
func NewPersonalAccountAPI(b Backend, nonceLock *AddrLocker) *PersonalAccountAPI {
	_ = "STUB: not implemented"
	return nil
}

// ListAccounts will return a list of addresses for accounts this node manages.
func (s *PersonalAccountAPI) ListAccounts() []common.Address { _ = "STUB: not implemented"; return nil }

// rawWallet is a JSON representation of an accounts.Wallet interface, with its
// data contents extracted into plain fields.
type rawWallet struct {
	URL      string             `json:"url"`
	Status   string             `json:"status"`
	Failure  string             `json:"failure,omitempty"`
	Accounts []accounts.Account `json:"accounts,omitempty"`
}

// ListWallets will return a list of wallets this node manages.
func (s *PersonalAccountAPI) ListWallets() []rawWallet { _ = "STUB: not implemented"; return nil }

// return [] instead of nil if empty

// OpenWallet initiates a hardware wallet opening procedure, establishing a USB
// connection and attempting to authenticate via the provided passphrase. Note,
// the method may return an extra challenge requiring a second open (e.g. the
// Trezor PIN matrix challenge).
func (s *PersonalAccountAPI) OpenWallet(url string, passphrase *string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeriveAccount requests an HD wallet to derive a new account, optionally pinning
// it for later reuse.
func (s *PersonalAccountAPI) DeriveAccount(url string, path string, pin *bool) (accounts.Account, error) {
	_ = "STUB: not implemented"
	return *new(accounts.Account), nil
}

// NewAccount will create a new account and returns the address for the new account.
func (s *PersonalAccountAPI) NewAccount(password string) (common.AddressEIP55, error) {
	_ = "STUB: not implemented"
	return *new(common.AddressEIP55), nil
}

// fetchKeystore retrieves the encrypted keystore from the account manager.
func fetchKeystore(am *accounts.Manager) (*keystore.KeyStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ImportRawKey stores the given hex encoded ECDSA key into the key directory,
// encrypting it with the passphrase.
func (s *PersonalAccountAPI) ImportRawKey(privkey string, password string) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// UnlockAccount will unlock the account associated with the given address with
// the given password for duration seconds. If duration is nil it will use a
// default of 300 seconds. It returns an indication if the account was unlocked.
func (s *PersonalAccountAPI) UnlockAccount(ctx context.Context, addr common.Address, password string, duration *uint64) (bool, error) {
	_ = "STUB: not implemented"
	// When the API is exposed by external RPC(http, ws etc), unless the user
	// explicitly specifies to allow the insecure account unlocking, otherwise
	// it is disabled.
	return false, nil
}

// LockAccount will lock the account associated with the given address when it's unlocked.
func (s *PersonalAccountAPI) LockAccount(addr common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// signTransaction sets defaults and signs the given transaction
// NOTE: the caller needs to ensure that the nonceLock is held, if applicable,
// and release it after the transaction has been submitted to the tx pool
func (s *PersonalAccountAPI) signTransaction(ctx context.Context, args *TransactionArgs, passwd string) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	// Look up the wallet containing the requested signer
	return nil, nil
}

// Set some sanity defaults and terminate on failure

// Assemble the transaction and sign with the wallet

// SendTransaction will create a transaction from the given arguments and
// tries to sign it with the key associated with args.From. If the given
// passwd isn't able to decrypt the key it fails.
func (s *PersonalAccountAPI) SendTransaction(ctx context.Context, args TransactionArgs, passwd string) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// Hold the mutex around signing to prevent concurrent assignment of
	// the same nonce to multiple accounts.
	new(common.Hash), nil
}

// SignTransaction will create a transaction from the given arguments and
// tries to sign it with the key associated with args.From. If the given passwd isn't
// able to decrypt the key it fails. The transaction is returned in RLP-form, not broadcast
// to other nodes
func (s *PersonalAccountAPI) SignTransaction(ctx context.Context, args TransactionArgs, passwd string) (*SignTransactionResult, error) {
	_ = "STUB: not implemented"
	// No need to obtain the noncelock mutex, since we won't be sending this
	// tx into the transaction pool, but right back to the user
	return nil, nil
}

// Before actually signing the transaction, ensure the transaction fee is reasonable.

// Sign calculates an Ethereum ECDSA signature for:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The key used to calculate the signature is decrypted with the given password.
//
// https://geth.ethereum.org/docs/interacting-with-geth/rpc/ns-personal#personal-sign
func (s *PersonalAccountAPI) Sign(ctx context.Context, data hexutil.Bytes, addr common.Address, passwd string) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	// Look up the wallet containing the requested signer
	return *new(hexutil.Bytes), nil
}

// Assemble sign the data with the wallet

// Transform V from 0/1 to 27/28 according to the yellow paper

// EcRecover returns the address for the account that was used to create the signature.
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be 27 or 28 for legacy reasons.
//
// https://geth.ethereum.org/docs/interacting-with-geth/rpc/ns-personal#personal-ecrecover
func (s *PersonalAccountAPI) EcRecover(ctx context.Context, data, sig hexutil.Bytes) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// Transform yellow paper V from 27/28 to 0/1

// InitializeWallet initializes a new wallet at the provided URL, by generating and returning a new private key.
func (s *PersonalAccountAPI) InitializeWallet(ctx context.Context, url string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Unpair deletes a pairing between wallet and geth.
func (s *PersonalAccountAPI) Unpair(ctx context.Context, url string, pin string) error {
	_ = "STUB: not implemented"
	return nil
}

// BlockChainAPI provides an API to access Ethereum blockchain data.
type BlockChainAPI struct {
	b Backend
}

// NewBlockChainAPI creates a new Ethereum blockchain API.
func NewBlockChainAPI(b Backend) *BlockChainAPI { _ = "STUB: not implemented"; return nil }

// ChainId is the EIP-155 replay-protection chain id for the current Ethereum chain config.
//
// Note, this method does not conform to EIP-695 because the configured chain ID is always
// returned, regardless of the current head block. We used to return an error when the chain
// wasn't synced up to a block where EIP-155 is enabled, but this behavior caused issues
// in CL clients.
func (api *BlockChainAPI) ChainId() *hexutil.Big { _ = "STUB: not implemented"; return nil }

// BlockNumber returns the block number of the chain head.
func (s *BlockChainAPI) BlockNumber() hexutil.Uint64 {
	_ = "STUB: not implemented"
	return *new(hexutil.Uint64)
}

// latest header should always be available

// GetBalance returns the amount of wei for the given address in the state of the
// given block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta
// block numbers are also allowed.
func (s *BlockChainAPI) GetBalance(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountResult structs for GetProof
type AccountResult struct {
	Address      common.Address  `json:"address"`
	AccountProof []string        `json:"accountProof"`
	Balance      *hexutil.Big    `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        hexutil.Uint64  `json:"nonce"`
	StorageHash  common.Hash     `json:"storageHash"`
	StorageProof []StorageResult `json:"storageProof"`
}

type StorageResult struct {
	Key   string       `json:"key"`
	Value *hexutil.Big `json:"value"`
	Proof []string     `json:"proof"`
}

// proofList implements ethdb.KeyValueWriter and collects the proofs as
// hex-strings for delivery to rpc-caller.
type proofList []string

func (n *proofList) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

func (n *proofList) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// GetProof returns the Merkle-proof for a given account and optionally some storage keys.
// If the requested block is part of historical blocks and the node does not accept
// getting proofs for historical blocks, an error is returned.
func (s *BlockChainAPI) GetProof(ctx context.Context, address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (*AccountResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deserialize all keys. This prevents state access on invalid input.

// Create the proofs for the storageKeys.

// Output key encoding is a bit special: if the input was a 32-byte hash, it is
// returned as such. Otherwise, we apply the QUANTITY encoding mandated by the
// JSON-RPC spec for getProof. This behavior exists to preserve backwards
// compatibility with older client versions.

// Create the accountProof.

// decodeHash parses a hex-encoded 32-byte hash. The input may optionally
// be prefixed by 0x and can have a byte length up to 32.
func decodeHash(s string) (h common.Hash, inputLength int, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), 0, nil
}

// GetHeaderByNumber returns the requested canonical block header.
//   - When blockNr is -1 the chain pending header is returned.
//   - When blockNr is -2 the chain latest header is returned.
//   - When blockNr is -3 the chain finalized header is returned.
//   - When blockNr is -4 the chain safe header is returned.
func (s *BlockChainAPI) GetHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// coreth has no notion of a pending block
// if number == rpc.PendingBlockNumber {
// 	// Pending header need to nil out a few fields
// 	for _, field := range []string{"hash", "nonce", "miner"} {
// 		response[field] = nil
// 	}
// }

// GetHeaderByHash returns the requested header by hash.
func (s *BlockChainAPI) GetHeaderByHash(ctx context.Context, hash common.Hash) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockByNumber returns the requested canonical block.
//   - When blockNr is -1 the chain pending block is returned.
//   - When blockNr is -2 the chain latest block is returned.
//   - When blockNr is -3 the chain finalized block is returned.
//   - When blockNr is -4 the chain safe block is returned.
//   - When fullTx is true all transactions in the block are returned, otherwise
//     only the transaction hash is returned.
func (s *BlockChainAPI) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// coreth has no notion of a pending block
// if err == nil && number == rpc.PendingBlockNumber {
// 	// Pending blocks need to nil out a few fields
// 	for _, field := range []string{"hash", "nonce", "miner"} {
// 		response[field] = nil
// 	}
// }

// GetBlockByHash returns the requested block. When fullTx is true all transactions in the block are returned in full
// detail, otherwise only the transaction hash is returned.
func (s *BlockChainAPI) GetBlockByHash(ctx context.Context, hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUncleByBlockNumberAndIndex returns the uncle block for the given block number and index.
func (s *BlockChainAPI) GetUncleByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUncleByBlockHashAndIndex returns the uncle block for the given block hash and index.
func (s *BlockChainAPI) GetUncleByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUncleCountByBlockNumber returns number of uncles in the block for the given block number
func (s *BlockChainAPI) GetUncleCountByBlockNumber(ctx context.Context, blockNr rpc.BlockNumber) *hexutil.Uint {
	_ = "STUB: not implemented"
	return nil
}

// GetUncleCountByBlockHash returns number of uncles in the block for the given block hash
func (s *BlockChainAPI) GetUncleCountByBlockHash(ctx context.Context, blockHash common.Hash) *hexutil.Uint {
	_ = "STUB: not implemented"
	return nil
}

// GetCode returns the code stored at the given address in the state for the given block number.
func (s *BlockChainAPI) GetCode(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetStorageAt returns the storage from the state at the given address, key and
// block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta block
// numbers are also allowed.
func (s *BlockChainAPI) GetStorageAt(ctx context.Context, address common.Address, hexKey string, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetBlockReceipts returns the block receipts for the given block hash or number or tag.
func (s *BlockChainAPI) GetBlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When the block doesn't exist, the RPC method should return JSON null
// as per specification.

// Derive the sender.

// OverrideAccount indicates the overriding fields of account during the execution
// of a message call.
// Note, state and stateDiff can't be specified at the same time. If state is
// set, message execution will only use the data in the given state. Otherwise
// if statDiff is set, all diff will be applied first and then execute the call
// message.
type OverrideAccount struct {
	Nonce     *hexutil.Uint64              `json:"nonce"`
	Code      *hexutil.Bytes               `json:"code"`
	Balance   **hexutil.Big                `json:"balance"`
	State     *map[common.Hash]common.Hash `json:"state"`
	StateDiff *map[common.Hash]common.Hash `json:"stateDiff"`
}

// StateOverride is the collection of overridden accounts.
type StateOverride map[common.Address]OverrideAccount

// Apply overrides the fields of specified accounts into the given state.
func (diff *StateOverride) Apply(state *state.StateDB) error { _ = "STUB: not implemented"; return nil }

// Override account nonce.

// Override account(contract) code.

// Override account balance.

// Replace entire state if caller requires.

// Apply state diff into specified accounts.

// Now finalize the changes. Finalize is normally performed between transactions.
// By using finalize, the overrides are semantically behaving as
// if they were created in a transaction just before the tracing occur.

// BlockOverrides is a set of header fields to override.
type BlockOverrides struct {
	Number      *hexutil.Big
	Difficulty  *hexutil.Big
	Time        *hexutil.Uint64
	GasLimit    *hexutil.Uint64
	Coinbase    *common.Address
	BaseFee     *hexutil.Big
	BlobBaseFee *hexutil.Big
}

// Apply overrides the given header fields into the given block context.
func (diff *BlockOverrides) Apply(blockCtx *vm.BlockContext) { _ = "STUB: not implemented"; return }

// ChainContextBackend provides methods required to implement ChainContext.
type ChainContextBackend interface {
	Engine() consensus.Engine
	HeaderByNumber(context.Context, rpc.BlockNumber) (*types.Header, error)
}

// ChainContext is an implementation of core.ChainContext. It's main use-case
// is instantiating a vm.BlockContext without having access to the BlockChain object.
type ChainContext struct {
	b   ChainContextBackend
	ctx context.Context
}

// NewChainContext creates a new ChainContext object.
func NewChainContext(ctx context.Context, backend ChainContextBackend) *ChainContext {
	_ = "STUB: not implemented"
	return nil
}

func (context *ChainContext) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	return *new(consensus.Engine)
}

func (context *ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	_ = "STUB: not implemented"
	// This method is called to get the hash for a block number when executing the BLOCKHASH
	// opcode. Hence no need to search for non-canonical blocks.
	return nil
}

func doCall(ctx context.Context, b Backend, args TransactionArgs, state *state.StateDB, header *types.Header, overrides *StateOverride, blockOverrides *BlockOverrides, timeout time.Duration, globalGasCap uint64) (*core.ExecutionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup context so it may be cancelled the call has completed
// or, in case of unmetered gas, setup a context with a timeout.

// Make sure the context is cancelled when the call has completed
// this makes sure resources are cleaned up.

// Get a new instance of the EVM.

// Wait for the context to be done and cancel the evm. Even if the
// EVM has finished, cancelling may be done (repeatedly)

// Execute the message.

// If the timer caused an abort, return an appropriate error message

func DoCall(ctx context.Context, b Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *StateOverride, blockOverrides *BlockOverrides, timeout time.Duration, globalGasCap uint64) (*core.ExecutionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the request is for the pending block, override the block timestamp, number, and estimated
// base fee, so that the check runs as if it were run on a newly generated block.

// Override header with a copy to ensure the original header is not modified

// Grab the hash of the unmodified header, so that the modified header can point to the
// prior block as its parent.

// Call executes the given transaction on the state for the given block number.
//
// Additionally, the caller can specify a batch of contract for fields overriding.
//
// Note, this function doesn't make and changes in the state/blockchain and is
// useful to execute and retrieve values.
func (s *BlockChainAPI) Call(ctx context.Context, args TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *StateOverride, blockOverrides *BlockOverrides) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// If the result contains a revert reason, try to unpack and return it.

// DoEstimateGas returns the lowest possible gas limit that allows the transaction to run
// successfully at block `blockNrOrHash`. It returns error if the transaction would revert, or if
// there are unexpected failures. The gas limit is capped by both `args.Gas` (if non-nil &
// non-zero) and `gasCap` (if non-zero).
func DoEstimateGas(ctx context.Context, b Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *StateOverride, gasCap uint64) (hexutil.Uint64, error) {
	_ = "STUB: not implemented"
	// Retrieve the base state and mutate it with any overrides
	return *new(hexutil.Uint64), nil
}

// Construct the gas estimator option from the user input

// If the user has not specified a gas limit, use the block gas limit

// Run the gas estimation andwrap any revertals into a custom return

// EstimateGas returns the lowest possible gas limit that allows the transaction to run
// successfully at block `blockNrOrHash`, or the latest block if `blockNrOrHash` is unspecified. It
// returns error if the transaction would revert or if there are unexpected failures. The returned
// value is capped by both `args.Gas` (if non-nil & non-zero) and the backend's RPCGasCap
// configuration (if non-zero).
// Note: Required blob gas is not computed in this method.
func (s *BlockChainAPI) EstimateGas(ctx context.Context, args TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *StateOverride) (hexutil.Uint64, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Uint64), nil
}

// RPCMarshalHeader converts the given header to the RPC output .
func RPCMarshalHeader(head *types.Header) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// RPCMarshalBlock converts the given block to the RPC output which depends on fullTx. If inclTx is true transactions are
// returned. When fullTx is true the returned block contains full transaction details, otherwise it will only contain
// transaction hashes.
func RPCMarshalBlock(block *types.Block, inclTx bool, fullTx bool, config *params.ChainConfig) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// rpcMarshalHeader uses the generalized output filler, then adds the total difficulty field, which requires
// a `BlockchainAPI`.
func (s *BlockChainAPI) rpcMarshalHeader(ctx context.Context, header *types.Header) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Note: Coreth enforces that the difficulty of a block is always 1, such that the total difficulty of a block
// will be equivalent to its height.

// rpcMarshalBlock uses the generalized output filler, then adds the total difficulty field, which requires
// a `BlockchainAPI`.
func (s *BlockChainAPI) rpcMarshalBlock(ctx context.Context, b *types.Block, inclTx bool, fullTx bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: Coreth enforces that the difficulty of a block is always 1, such that the total difficulty of a block
// will be equivalent to its height.

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash           *common.Hash      `json:"blockHash"`
	BlockNumber         *hexutil.Big      `json:"blockNumber"`
	From                common.Address    `json:"from"`
	Gas                 hexutil.Uint64    `json:"gas"`
	GasPrice            *hexutil.Big      `json:"gasPrice"`
	GasFeeCap           *hexutil.Big      `json:"maxFeePerGas,omitempty"`
	GasTipCap           *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerBlobGas    *hexutil.Big      `json:"maxFeePerBlobGas,omitempty"`
	Hash                common.Hash       `json:"hash"`
	Input               hexutil.Bytes     `json:"input"`
	Nonce               hexutil.Uint64    `json:"nonce"`
	To                  *common.Address   `json:"to"`
	TransactionIndex    *hexutil.Uint64   `json:"transactionIndex"`
	Value               *hexutil.Big      `json:"value"`
	Type                hexutil.Uint64    `json:"type"`
	Accesses            *types.AccessList `json:"accessList,omitempty"`
	ChainID             *hexutil.Big      `json:"chainId,omitempty"`
	BlobVersionedHashes []common.Hash     `json:"blobVersionedHashes,omitempty"`
	V                   *hexutil.Big      `json:"v"`
	R                   *hexutil.Big      `json:"r"`
	S                   *hexutil.Big      `json:"s"`
	YParity             *hexutil.Uint64   `json:"yParity,omitempty"`
}

// newRPCTransaction returns a transaction that will serialize to the RPC
// representation, with the given location metadata set (if available).
func newRPCTransaction(tx *types.Transaction, blockHash common.Hash, blockNumber uint64, blockTime uint64, index uint64, baseFee *big.Int, config *params.ChainConfig) *RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// if a legacy transaction has an EIP-155 chain id, include it explicitly

// if the transaction has been mined, compute the effective gas price

// price = min(gasTipCap + baseFee, gasFeeCap)

// if the transaction has been mined, compute the effective gas price

// effectiveGasPrice computes the transaction gas fee, based on the given basefee value.
//
//	price = min(gasTipCap + baseFee, gasFeeCap)
func effectiveGasPrice(tx *types.Transaction, baseFee *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// NewRPCTransaction returns a pending transaction that will serialize to the RPC representation
// Note: in go-ethereum this function is called NewRPCPendingTransaction.
// In coreth, we have renamed it to NewRPCTransaction as it is used for accepted transactions as well.
func NewRPCTransaction(tx *types.Transaction, current *types.Header, baseFee *big.Int, config *params.ChainConfig) *RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// newRPCTransactionFromBlockIndex returns a transaction that will serialize to the RPC representation.
func newRPCTransactionFromBlockIndex(b *types.Block, index uint64, config *params.ChainConfig) *RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// newRPCRawTransactionFromBlockIndex returns the bytes of a transaction given a block and a transaction index.
func newRPCRawTransactionFromBlockIndex(b *types.Block, index uint64) hexutil.Bytes {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes)
}

// accessListResult returns an optional accesslist
// It's the result of the `debug_createAccessList` RPC call.
// It contains an error if the transaction itself failed.
type accessListResult struct {
	Accesslist *types.AccessList `json:"accessList"`
	Error      string            `json:"error,omitempty"`
	GasUsed    hexutil.Uint64    `json:"gasUsed"`
}

// CreateAccessList creates an EIP-2930 type AccessList for the given transaction.
// Reexec and BlockNrOrHash can be specified to create the accessList on top of a certain state.
func (s *BlockChainAPI) CreateAccessList(ctx context.Context, args TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash) (*accessListResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccessList creates an access list for the given transaction.
// If the accesslist creation fails an error is returned.
// If the transaction itself fails, an vmErr is returned.
func AccessList(ctx context.Context, b Backend, blockNrOrHash rpc.BlockNumberOrHash, args TransactionArgs) (acl types.AccessList, gasUsed uint64, vmErr error, err error) {
	_ = "STUB: not implemented"
	// Retrieve the execution context
	return *new(types.AccessList), 0, nil, nil
}

// Ensure any missing fields are filled, extract the recipient and input data

// Retrieve the precompiles since they don't need to be added to the access list

// Create an initial tracer

// Retrieve the current access list to expand

// Copy the original db so we don't modify it

// Set the accesslist to the last al

// Apply the transaction with the access list tracer

// TransactionAPI exposes methods for reading and creating transaction data.
type TransactionAPI struct {
	b         Backend
	nonceLock *AddrLocker
	signer    types.Signer
}

// NewTransactionAPI creates a new RPC service with methods for interacting with transactions.
func NewTransactionAPI(b Backend, nonceLock *AddrLocker) *TransactionAPI {
	_ = "STUB: not implemented"
	// The signer used by the API should always be the 'latest' known one because we expect
	// signers to be backwards-compatible with old transactions.
	return nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block with the given block number.
func (s *TransactionAPI) GetBlockTransactionCountByNumber(ctx context.Context, blockNr rpc.BlockNumber) *hexutil.Uint {
	_ = "STUB: not implemented"
	return nil
}

// GetBlockTransactionCountByHash returns the number of transactions in the block with the given hash.
func (s *TransactionAPI) GetBlockTransactionCountByHash(ctx context.Context, blockHash common.Hash) *hexutil.Uint {
	_ = "STUB: not implemented"
	return nil
}

// GetTransactionByBlockNumberAndIndex returns the transaction for the given block number and index.
func (s *TransactionAPI) GetTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) *RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// GetTransactionByBlockHashAndIndex returns the transaction for the given block hash and index.
func (s *TransactionAPI) GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) *RPCTransaction {
	_ = "STUB: not implemented"
	return nil
}

// GetRawTransactionByBlockNumberAndIndex returns the bytes of the transaction for the given block number and index.
func (s *TransactionAPI) GetRawTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) hexutil.Bytes {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes)
}

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
func (s *TransactionAPI) GetRawTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) hexutil.Bytes {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes)
}

// GetTransactionCount returns the number of transactions the given address has sent for the given block number
func (s *TransactionAPI) GetTransactionCount(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	_ = "STUB: not implemented"
	// Ask transaction pool for the nonce which includes pending transactions
	return nil, nil
}

// Resolve block number and use its state to ask for the nonce

// GetTransactionByHash returns the transaction for the given hash
func (s *TransactionAPI) GetTransactionByHash(ctx context.Context, hash common.Hash) (*RPCTransaction, error) {
	_ = "STUB: not implemented"
	// Try to return an already finalized transaction
	return nil, nil
}

// No finalized transaction, try to retrieve it from the pool

// GetRawTransactionByHash returns the bytes of the transaction for the given hash.
func (s *TransactionAPI) GetRawTransactionByHash(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	// Retrieve a finalized transaction, or a pooled otherwise
	return *new(hexutil.Bytes), nil
}

// GetTransactionReceipt returns the transaction receipt for the given transaction hash.
func (s *TransactionAPI) GetTransactionReceipt(ctx context.Context, hash common.Hash) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// transaction is not fully indexed

// transaction is not existent or reachable

// Derive the sender.

// marshalReceipt marshals a transaction receipt into a JSON object.
func marshalReceipt(receipt *types.Receipt, blockHash common.Hash, blockNumber uint64, signer types.Signer, tx *types.Transaction, txIndex int) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Assign receipt status or post state.

// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation

// sign is a helper function that signs a transaction with the private key of the given address.
func (s *TransactionAPI) sign(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	// Look up the wallet containing the requested signer
	return nil, nil
}

// Request the wallet to sign the transaction

// SubmitTransaction is a helper function that submits tx to txPool and logs a message.
func SubmitTransaction(ctx context.Context, b Backend, tx *types.Transaction) (common.Hash, error) {
	_ = "STUB: not implemented"
	// If the transaction fee cap is already specified, ensure the
	// fee of the given transaction is _reasonable_.
	return *new(common.Hash), nil
}

// Ensure only eip155 signed transactions are submitted if EIP155Required is set.

// Print a log with full tx details for manual investigations and interventions

// SendTransaction creates a transaction for the given argument, sign it and submit it to the
// transaction pool.
func (s *TransactionAPI) SendTransaction(ctx context.Context, args TransactionArgs) (common.Hash, error) {
	_ = "STUB: not implemented"
	// Look up the wallet containing the requested signer
	return *new(common.Hash), nil
}

// Hold the mutex around signing to prevent concurrent assignment of
// the same nonce to multiple accounts.

// Set some sanity defaults and terminate on failure

// Assemble the transaction and sign with the wallet

// FillTransaction fills the defaults (nonce, gas, gasPrice or 1559 fields)
// on a given unsigned transaction, and returns it to the caller for further
// processing (signing + broadcast).
func (s *TransactionAPI) FillTransaction(ctx context.Context, args TransactionArgs) (*SignTransactionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Set some sanity defaults and terminate on failure
}

// Assemble the transaction and obtain rlp

// SendRawTransaction will add the signed transaction to the transaction pool.
// The sender is responsible for signing the transaction and using the correct nonce.
func (s *TransactionAPI) SendRawTransaction(ctx context.Context, input hexutil.Bytes) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Sign calculates an ECDSA signature for:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message).
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The account associated with addr must be unlocked.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sign
func (s *TransactionAPI) Sign(addr common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	// Look up the wallet containing the requested signer
	return *new(hexutil.Bytes), nil
}

// Sign the requested hash with the wallet

// Transform V from 0/1 to 27/28 according to the yellow paper

// SignTransactionResult represents a RLP encoded signed transaction.
type SignTransactionResult struct {
	Raw hexutil.Bytes      `json:"raw"`
	Tx  *types.Transaction `json:"tx"`
}

// SignTransaction will sign the given transaction with the from account.
// The node needs to have the private key of the account corresponding with
// the given from address and it needs to be unlocked.
func (s *TransactionAPI) SignTransaction(ctx context.Context, args TransactionArgs) (*SignTransactionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Before actually sign the transaction, ensure the transaction fee is reasonable.

// PendingTransactions returns the transactions that are in the transaction pool
// and have a from address that is one of the accounts this node manages.
func (s *TransactionAPI) PendingTransactions() ([]*RPCTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resend accepts an existing transaction and a new gas price and limit. It will remove
// the given transaction from the pool and reinsert it with the new gas price and limit.
func (s *TransactionAPI) Resend(ctx context.Context, sendArgs TransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Before replacing the old transaction, ensure the _new_ transaction fee is reasonable.

// Iterate the pending list for replacement

// Match. Re-sign and send the transaction.

// DebugAPI is the collection of Ethereum APIs exposed over the debugging
// namespace.
type DebugAPI struct {
	b Backend
}

// NewDebugAPI creates a new instance of DebugAPI.
func NewDebugAPI(b Backend) *DebugAPI { _ = "STUB: not implemented"; return nil }

// GetRawHeader retrieves the RLP encoding for a single header.
func (api *DebugAPI) GetRawHeader(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetRawBlock retrieves the RLP encoded for a single block.
func (api *DebugAPI) GetRawBlock(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetRawReceipts retrieves the binary-encoded receipts of a single block.
func (api *DebugAPI) GetRawReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRawTransaction returns the bytes of the transaction for the given hash.
func (s *DebugAPI) GetRawTransaction(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	// Retrieve a finalized transaction, or a pooled otherwise
	return *new(hexutil.Bytes), nil
}

// PrintBlock retrieves a block and returns its pretty printed form.
func (api *DebugAPI) PrintBlock(ctx context.Context, number uint64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NetAPI offers network related RPC methods
type NetAPI struct {
	networkVersion uint64
}

// NewNetAPI creates a new net API instance.
func NewNetAPI(networkVersion uint64) *NetAPI { _ = "STUB: not implemented"; return nil }

// Listening returns an indication if the node is listening for network connections.
func (s *NetAPI) Listening() bool {
	_ = "STUB: not implemented"
	// always listening
	return false
}

// PeerCount returns the number of connected peers
func (s *NetAPI) PeerCount() hexutil.Uint {
	_ = "STUB: not implemented"
	return *

	// Version returns the current ethereum protocol version.
	new(hexutil.Uint)
}

func (s *NetAPI) Version() string { _ = "STUB: not implemented"; return "" }

// checkTxFee is an internal function used to check whether the fee of
// the given transaction is _reasonable_(under the cap).
func checkTxFee(gasPrice *big.Int, gas uint64, cap float64) error {
	_ = "STUB: not implemented"
	// Short circuit if there is no cap for transaction fee at all.
	return nil
}
