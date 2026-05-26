// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"errors"
	"time"
)

const (
	// MinPercentConnectedBuffer is the safety buffer for calculation of
	// MinPercentConnected. This increases the required percentage above
	// alpha/k. This value must be [0-1].
	// 0 means MinPercentConnected = alpha/k.
	// 1 means MinPercentConnected = 1 (fully connected).
	MinPercentConnectedBuffer = .2

	errMsg = `__________                    .___
\______   \____________     __| _/__.__.
 |    |  _/\_  __ \__  \   / __ <   |  |
 |    |   \ |  | \// __ \_/ /_/ |\___  |
 |______  / |__|  (____  /\____ |/ ____|
        \/             \/      \/\/

  🏆    🏆    🏆    🏆    🏆    🏆    🏆
  ________ ________      ________________
 /  _____/ \_____  \    /  _  \__    ___/
/   \  ___  /   |   \  /  /_\  \|    |
\    \_\  \/    |    \/    |    \    |
 \______  /\_______  /\____|__  /____|
        \/         \/         \/
`
)

var (
	DefaultParameters = Parameters{
		K:                     20,
		AlphaPreference:       15,
		AlphaConfidence:       15,
		Beta:                  20,
		ConcurrentRepolls:     4,
		OptimalProcessing:     10,
		MaxOutstandingItems:   256,
		MaxItemProcessingTime: 30 * time.Second,
	}

	ErrParametersInvalid = errors.New("parameters invalid")
)

// Parameters required for snowball consensus
type Parameters struct {
	// K is the number of nodes to query and sample in a round.
	K int `json:"k" yaml:"k"`
	// Alpha is used for backwards compatibility purposes and is only referenced
	// during json parsing.
	Alpha *int `json:"alpha,omitempty" yaml:"alpha,omitempty"`
	// AlphaPreference is the vote threshold to change your preference.
	AlphaPreference int `json:"alphaPreference" yaml:"alphaPreference"`
	// AlphaConfidence is the vote threshold to increase your confidence.
	AlphaConfidence int `json:"alphaConfidence" yaml:"alphaConfidence"`
	// Beta is the number of consecutive successful queries required for
	// finalization.
	Beta int `json:"beta" yaml:"beta"`
	// ConcurrentRepolls is the number of outstanding polls the engine will
	// target to have while there is something processing.
	ConcurrentRepolls int `json:"concurrentRepolls" yaml:"concurrentRepolls"`
	// OptimalProcessing is used to limit block creation when a large number of
	// blocks are processing.
	OptimalProcessing int `json:"optimalProcessing" yaml:"optimalProcessing"`

	// Reports unhealthy if more than this number of items are outstanding.
	MaxOutstandingItems int `json:"maxOutstandingItems" yaml:"maxOutstandingItems"`

	// Reports unhealthy if there is an item processing for longer than this
	// duration.
	MaxItemProcessingTime time.Duration `json:"maxItemProcessingTime" yaml:"maxItemProcessingTime"`
}

// Verify returns nil if the parameters describe a valid initialization.
//
// An initialization is valid if the following conditions are met:
//
// - K/2 < AlphaPreference <= AlphaConfidence <= K
// - 0 < ConcurrentRepolls <= Beta
// - 0 < OptimalProcessing
// - 0 < MaxOutstandingItems
// - 0 < MaxItemProcessingTime
//
// Note: K/2 < K implies that 0 <= K/2, so we don't need an explicit check that
// AlphaPreference is positive.
func (p Parameters) Verify() error { _ = "STUB: not implemented"; return nil }

func (p Parameters) MinPercentConnectedHealthy() float64 {
	_ = "STUB: not implemented"
	// AlphaConfidence is used here to ensure that the node can still feasibly
	// accept operations. If AlphaPreference were used, committing could be
	// extremely unlikely to happen, even while healthy.
	return 0
}

type terminationCondition struct {
	alphaConfidence int
	beta            int
}

func newSingleTerminationCondition(alphaConfidence int, beta int) []terminationCondition {
	_ = "STUB: not implemented"
	return nil
}
