// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"github.com/spf13/pflag"
)

type CollectorVars struct {
	StartMetricsCollector bool
	StartLogsCollector    bool
}

// NewCollectorFlagVars registers collector flag variables for stdlib flag
func NewCollectorFlagVars() *CollectorVars { _ = "STUB: not implemented"; return nil }

// NewCollectorFlagSetVars registers collector flag variables for pflag
func NewCollectorFlagSetVars(flagSet *pflag.FlagSet) *CollectorVars {
	_ = "STUB: not implemented"
	return nil
}

func (v *CollectorVars) register(boolVar varFunc[bool]) { _ = "STUB: not implemented"; return }
