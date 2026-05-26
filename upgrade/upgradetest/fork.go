// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package upgradetest

const (
	NoUpgrades Fork = iota
	ApricotPhase1
	ApricotPhase2
	ApricotPhase3
	ApricotPhase4
	ApricotPhase5
	ApricotPhasePre6
	ApricotPhase6
	ApricotPhasePost6
	Banff
	Cortina
	Durango
	Etna
	Fortuna
	Granite
	Helicon

	Latest Fork = iota - 1
)

// Fork is an enum of all the major network upgrades.
type Fork int

func (f Fork) String() string { _ = "STUB: not implemented"; return "" }
