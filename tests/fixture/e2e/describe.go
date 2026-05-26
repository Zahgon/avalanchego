// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e

const (
	// For label usage in ginkgo invocation, see: https://onsi.github.io/ginkgo/#spec-labels

	// Label for filtering a test that is not primarily a C-Chain test
	// but nonetheless uses the C-Chain. Intended to support
	// execution of all C-Chain tests by the coreth repo in an e2e job.
	UsesCChainLabel = "uses-c"
)

// DescribeXChain annotates the tests for X-Chain.
func DescribeXChain(text string, args ...interface{}) bool { _ = "STUB: not implemented"; return false }

// DescribeXChainSerial annotates serial tests for X-Chain.
func DescribeXChainSerial(text string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// DescribePChain annotates the tests for P-Chain.
func DescribePChain(text string, args ...interface{}) bool { _ = "STUB: not implemented"; return false }

// DescribeCChain annotates the tests for C-Chain.
func DescribeCChain(text string, args ...interface{}) bool { _ = "STUB: not implemented"; return false }
