// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"github.com/ava-labs/libevm/eth/filters"
	"github.com/ava-labs/libevm/rpc"
)

// Taken as the defaults from geth / libevm's `node.DefaultConfig`.
const (
	batchLimit           = 1000
	batchResponseMaxSize = 25 * 1000 * 1000 // 25 MB
)

// Server returns the Provider's [rpc.Server], with all configured JSON-RPC
// namespace handlers registered.
func (p *Provider) Server() *rpc.Server { _ = "STUB: not implemented"; return nil }

func (b *backend) server(filter *filters.FilterAPI) (*rpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Standard Ethereum APIs are documented at: https://ethereum.org/developers/docs/apis/json-rpc
// geth-specific APIs are documented at: https://geth.ethereum.org/docs/interacting-with-geth/rpc

// Standard Ethereum node APIs:
// - web3_clientVersion
// - web3_sha3

// Standard Ethereum node APIs:
// - net_listening
// - net_peerCount
// - net_version

// geth-specific APIs:
// - txpool_content
// - txpool_contentFrom
// - txpool_inspect
// - txpool_status

// Standard Ethereum node APIs:
// - eth_gasPrice
// - eth_maxPriorityFeePerGas
// - eth_feeHistory
// - eth_syncing

// Standard Ethereum node APIs:
// - eth_blockNumber
// - eth_chainId
// - eth_getBlockByHash
// - eth_getBlockByNumber
// - eth_getBlockReceipts
// - eth_getUncleByBlockHashAndIndex
// - eth_getUncleByBlockNumberAndIndex
// - eth_getUncleCountByBlockHash
// - eth_getUncleCountByBlockNumber
//
// geth-specific APIs:
// - eth_getHeaderByHash
// - eth_getHeaderByNumber

// Standard Ethereum node APIs:
// - eth_getBlockTransactionCountByHash
// - eth_getBlockTransactionCountByNumber
// - eth_getTransactionByBlockHashAndIndex
// - eth_getTransactionByBlockNumberAndIndex
// - eth_getTransactionByHash
// - eth_getTransactionCount
// - eth_getTransactionReceipt
// - eth_sendRawTransaction
// - eth_sendTransaction
// - eth_sign
// - eth_signTransaction
//
// Undocumented APIs:
// - eth_fillTransaction
// - eth_getRawTransactionByBlockHashAndIndex
// - eth_getRawTransactionByBlockNumberAndIndex
// - eth_getRawTransactionByHash
// - eth_pendingTransactions
// - eth_resend

// Standard Ethereum node APIS:
// - eth_getFilterChanges
// - eth_getFilterLogs
// - eth_getLogs
// - eth_newBlockFilter
// - eth_newFilter
// - eth_newPendingTransactionFilter
// - eth_uninstallFilter
//
// geth-specific APIs:
// - eth_subscribe
//  - newHeads
//  - newPendingTransactions
//  - logs

// Avalanche-custom eth extensions:

// geth-specific APIs:
// - debug_chaindbCompact
// - debug_chaindbProperty
// - debug_dbAncient
// - debug_dbAncients
// - debug_dbGet
// - debug_getRawBlock
// - debug_getRawHeader
// - debug_getRawReceipts
// - debug_getRawTransaction
// - debug_printBlock
// - debug_setHead          (no-op, logs info)

// geth-specific APIs:
// - debug_blockProfile
// - debug_cpuProfile
// - debug_freeOSMemory
// - debug_gcStats
// - debug_goTrace
// - debug_memStats
// - debug_mutexProfile
// - debug_setBlockProfileRate
// - debug_setGCPercent
// - debug_setMutexProfileFraction
// - debug_stacks
// - debug_startCPUProfile
// - debug_startGoTrace
// - debug_stopCPUProfile
// - debug_stopGoTrace
// - debug_verbosity
// - debug_vmodule
// - debug_writeBlockProfile
// - debug_writeMemProfile
// - debug_writeMutexProfile

// geth-specific APIs:
