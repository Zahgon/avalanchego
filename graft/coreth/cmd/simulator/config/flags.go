// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"errors"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const Version = "v0.1.1"

const (
	ConfigFilePathKey = "config-file"
	LogLevelKey       = "log-level"
	EndpointsKey      = "endpoints"
	MaxFeeCapKey      = "max-fee-cap"
	MaxTipCapKey      = "max-tip-cap"
	WorkersKey        = "workers"
	TxsPerWorkerKey   = "txs-per-worker"
	KeyDirKey         = "key-dir"
	VersionKey        = "version"
	TimeoutKey        = "timeout"
	BatchSizeKey      = "batch-size"
	MetricsPortKey    = "metrics-port"
	MetricsOutputKey  = "metrics-output"
)

var (
	ErrNoEndpoints = errors.New("must specify at least one endpoint")
	ErrNoWorkers   = errors.New("must specify non-zero number of workers")
	ErrNoTxs       = errors.New("must specify non-zero number of txs-per-worker")
)

type Config struct {
	Endpoints     []string      `json:"endpoints"`
	MaxFeeCap     int64         `json:"max-fee-cap"`
	MaxTipCap     int64         `json:"max-tip-cap"`
	Workers       int           `json:"workers"`
	TxsPerWorker  uint64        `json:"txs-per-worker"`
	KeyDir        string        `json:"key-dir"`
	Timeout       time.Duration `json:"timeout"`
	BatchSize     uint64        `json:"batch-size"`
	MetricsPort   uint64        `json:"metrics-port"`
	MetricsOutput string        `json:"metrics-output"`
}

func BuildConfig(v *viper.Viper) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// Note: it's technically valid for the fee/tip cap to be 0, but cannot
// be less than 0.

func BuildViper(fs *pflag.FlagSet, args []string) (*viper.Viper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildFlagSet returns a complete set of flags for simulator
func BuildFlagSet() *pflag.FlagSet { _ = "STUB: not implemented"; return nil }

func addSimulatorFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }
