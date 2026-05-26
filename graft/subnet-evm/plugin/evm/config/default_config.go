// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"time"
)

const defaultCommitInterval = 4096

func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// EIP-1820: https://eips.ethereum.org/EIPS/eip-1820

// Provides 2 minutes of buffer (2s block target) for a commit delay

// 50M Gas Limit
// 100 AVAX

// Default to no maximum API call duration

// Default to no maximum WS CPU usage

// Default to no maximum WS CPU usage

// Default to no maximum on the number of blocks per getLogs request

// Default size (MB) for the offline pruner to use

// MB
// blocks
// StateSyncMinBlocks is the minimum number of blocks the blockchain
// should be ahead of local last accepted to perform state sync.
// This constant is chosen so normal bootstrapping is preferred when it would
// be faster than state sync.
// time assumptions:
// - normal bootstrap processing time: ~14 blocks / second
// - state sync time: ~6 hrs.

// the number of key/values to ask peers for per request

// Estimated block count in 24 hours with 2s block accept period

// Mempool settings

// urgent + floating queue capacity with 4:1 ratio,

// RPC settings

// 25MB
// Subnet EVM API settings

// Database settings

// Additional settings with sensible defaults

func timeToDuration(t time.Duration) Duration { _ = "STUB: not implemented"; return *new(Duration) }
