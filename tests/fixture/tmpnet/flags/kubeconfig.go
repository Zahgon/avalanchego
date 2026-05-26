// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"github.com/spf13/pflag"
)

const KubeconfigPathEnvVar = "KUBECONFIG"

type KubeconfigVars struct {
	Path    string
	Context string
}

// NewKubeconfigFlagVars registers kubeconfig flag variables for stdlib flag
func NewKubeconfigFlagVars() *KubeconfigVars { _ = "STUB: not implemented"; return nil }

// internal method enabling configuration of the doc prefix
func newKubeconfigFlagVars(docPrefix string) *KubeconfigVars { _ = "STUB: not implemented"; return nil }

// NewKubeconfigFlagSetVars registers kubeconfig flag variables for pflag
func NewKubeconfigFlagSetVars(flagSet *pflag.FlagSet) *KubeconfigVars {
	_ = "STUB: not implemented"
	return nil
}

// internal method enabling configuration of the doc prefix
func newKubeconfigFlagSetVars(flagSet *pflag.FlagSet, docPrefix string) *KubeconfigVars {
	_ = "STUB: not implemented"
	return nil
}

func (v *KubeconfigVars) register(stringVar varFunc[string], docPrefix string) {
	_ = "STUB: not implemented"
	// the default kubeConfig path is set to empty to allow for the use of a projected
	// token when running in-cluster
	return
}
