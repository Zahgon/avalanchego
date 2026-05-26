// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const EnvPrefix = "avago"

var DashesToUnderscores = strings.NewReplacer("-", "_")

func EnvVarName(prefix string, key string) string {
	_ = "STUB: not implemented"
	// e.g. MY_PREFIX, network-id -> MY_PREFIX_NETWORK_ID
	return ""
}

// BuildViper returns the viper environment from parsing config file from
// default search paths and any parsed command line flags
func BuildViper(fs *pflag.FlagSet, args []string) (*viper.Viper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// load node configs from flags or file, depending on which flags are set

// Config deprecations must be after v.ReadInConfig

func deprecateConfigs(v *viper.Viper, output io.Writer) { _ = "STUB: not implemented"; return }
