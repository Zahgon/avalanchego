// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package upgradetest

import (
	"time"

	"github.com/ava-labs/avalanchego/upgrade"
)

// GetConfig returns an upgrade config with the provided fork scheduled to have
// been initially activated and all other forks to be unscheduled.
func GetConfig(fork Fork) upgrade.Config { _ = "STUB: not implemented"; return *new(upgrade.Config) }

// GetConfigWithUpgradeTime returns an upgrade config with the provided fork
// scheduled to be activated at the provided upgradeTime and all other forks to
// be unscheduled.
func GetConfigWithUpgradeTime(fork Fork, upgradeTime time.Time) upgrade.Config {
	_ = "STUB: not implemented"
	return *new(upgrade.Config)
}

// Initialize all forks to be unscheduled

// Schedule the requested forks at the provided upgrade time

// SetTimesTo sets the upgrade time of the provided fork, and all prior forks,
// to the provided upgradeTime.
func SetTimesTo(c *upgrade.Config, fork Fork, upgradeTime time.Time) {
	_ = "STUB: not implemented"
	return
}
