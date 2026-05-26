// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

const (
	kubeRuntime     = "kube"
	kubeFlagsPrefix = kubeRuntime + "-"
	kubeDocPrefix   = "[kube runtime] "
)

var (
	errKubeNamespaceRequired     = errors.New("--kube-namespace is required")
	errKubeImageRequired         = errors.New("--kube-image is required")
	errKubeMinVolumeSizeRequired = fmt.Errorf("--kube-volume-size must be >= %d", tmpnet.MinimumVolumeSizeGB)
)

type kubeRuntimeVars struct {
	namespace              string
	image                  string
	volumeSizeGB           uint
	useExclusiveScheduling bool
	schedulingLabelKey     string
	schedulingLabelValue   string
	config                 *KubeconfigVars
}

func (v *kubeRuntimeVars) registerWithFlag() { _ = "STUB: not implemented"; return }

func (v *kubeRuntimeVars) registerWithFlagSet(flagSet *pflag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

func (v *kubeRuntimeVars) register(stringVar varFunc[string], uintVar varFunc[uint], boolVar varFunc[bool]) {
	_ = "STUB: not implemented"
	return
}

func (v *kubeRuntimeVars) getKubeRuntimeConfig() (*tmpnet.KubeRuntimeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
