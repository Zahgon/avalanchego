// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import "github.com/ava-labs/libevm/metrics"

type verifierStats struct {
	messageParseFail metrics.Counter
	// AddressedCall metrics
	addressedCallValidationFail metrics.Counter
	// BlockRequest metrics
	blockValidationFail metrics.Counter
	// Uptime metrics
	uptimeValidationFail metrics.Counter
}

func newVerifierStats() *verifierStats { _ = "STUB: not implemented"; return nil }

func (h *verifierStats) IncAddressedCallValidationFail() { _ = "STUB: not implemented"; return }

func (h *verifierStats) IncBlockValidationFail() { _ = "STUB: not implemented"; return }

func (h *verifierStats) IncMessageParseFail() { _ = "STUB: not implemented"; return }

func (h *verifierStats) IncUptimeValidationFail() { _ = "STUB: not implemented"; return }
