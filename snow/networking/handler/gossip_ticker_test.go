// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGossipTickerDisabled(t *testing.T) {
	g := newGossipTicker(0)
	defer g.Stop()

	// A nil channel is never ready in a select, so the gossip arm of the
	// dispatcher is effectively disabled if g.C == nil.
	require.Nil(t, g.C)
}

func TestGossipTickerEnabled(t *testing.T) {
	g := newGossipTicker(time.Millisecond)
	defer g.Stop()

	require.NotNil(t, g.C)
	select {
	case <-g.C:
	case <-time.After(time.Second):
		t.Fatal("ticker did not fire")
	}
}
