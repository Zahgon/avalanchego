// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertextest

import (
	"testing"

	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
)

var _ vertex.Manager = (*Manager)(nil)

type Manager struct {
	Builder
	Parser
	Storage
}

func NewManager(t *testing.T) *Manager { _ = "STUB: not implemented"; return nil }

func (m *Manager) Default(cant bool) { _ = "STUB: not implemented"; return }
