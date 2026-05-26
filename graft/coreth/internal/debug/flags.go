// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package debug

import (
	"io"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/ava-labs/avalanchego/graft/coreth/internal/flags"
	"github.com/ava-labs/libevm/log"
	"github.com/urfave/cli/v2"
)

var (
	verbosityFlag = &cli.IntFlag{
		Name:     "verbosity",
		Usage:    "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
		Value:    3,
		Category: flags.LoggingCategory,
	}
	logVmoduleFlag = &cli.StringFlag{
		Name:     "log.vmodule",
		Usage:    "Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. eth/*=5,p2p=4)",
		Value:    "",
		Category: flags.LoggingCategory,
	}
	vmoduleFlag = &cli.StringFlag{
		Name:     "vmodule",
		Usage:    "Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. eth/*=5,p2p=4)",
		Value:    "",
		Hidden:   true,
		Category: flags.LoggingCategory,
	}
	logjsonFlag = &cli.BoolFlag{
		Name:     "log.json",
		Usage:    "Format logs with JSON",
		Hidden:   true,
		Category: flags.LoggingCategory,
	}
	logFormatFlag = &cli.StringFlag{
		Name:     "log.format",
		Usage:    "Log format to use (json|logfmt|terminal)",
		Category: flags.LoggingCategory,
	}
	logFileFlag = &cli.StringFlag{
		Name:     "log.file",
		Usage:    "Write logs to a file",
		Category: flags.LoggingCategory,
	}
	logRotateFlag = &cli.BoolFlag{
		Name:     "log.rotate",
		Usage:    "Enables log file rotation",
		Category: flags.LoggingCategory,
	}
	logMaxSizeMBsFlag = &cli.IntFlag{
		Name:     "log.maxsize",
		Usage:    "Maximum size in MBs of a single log file",
		Value:    100,
		Category: flags.LoggingCategory,
	}
	logMaxBackupsFlag = &cli.IntFlag{
		Name:     "log.maxbackups",
		Usage:    "Maximum number of log files to retain",
		Value:    10,
		Category: flags.LoggingCategory,
	}
	logMaxAgeFlag = &cli.IntFlag{
		Name:     "log.maxage",
		Usage:    "Maximum number of days to retain a log file",
		Value:    30,
		Category: flags.LoggingCategory,
	}
	logCompressFlag = &cli.BoolFlag{
		Name:     "log.compress",
		Usage:    "Compress the log files",
		Value:    false,
		Category: flags.LoggingCategory,
	}
	pprofFlag = &cli.BoolFlag{
		Name:     "pprof",
		Usage:    "Enable the pprof HTTP server",
		Category: flags.LoggingCategory,
	}
	pprofPortFlag = &cli.IntFlag{
		Name:     "pprof.port",
		Usage:    "pprof HTTP server listening port",
		Value:    6060,
		Category: flags.LoggingCategory,
	}
	pprofAddrFlag = &cli.StringFlag{
		Name:     "pprof.addr",
		Usage:    "pprof HTTP server listening interface",
		Value:    "127.0.0.1",
		Category: flags.LoggingCategory,
	}
	memprofilerateFlag = &cli.IntFlag{
		Name:     "pprof.memprofilerate",
		Usage:    "Turn on memory profiling with the given rate",
		Value:    runtime.MemProfileRate,
		Category: flags.LoggingCategory,
	}
	blockprofilerateFlag = &cli.IntFlag{
		Name:     "pprof.blockprofilerate",
		Usage:    "Turn on block profiling with the given rate",
		Category: flags.LoggingCategory,
	}
	cpuprofileFlag = &cli.StringFlag{
		Name:     "pprof.cpuprofile",
		Usage:    "Write CPU profile to the given file",
		Category: flags.LoggingCategory,
	}
	traceFlag = &cli.StringFlag{
		Name:     "trace",
		Usage:    "Write execution trace to the given file",
		Category: flags.LoggingCategory,
	}
)

// Flags holds all command-line flags required for debugging.
var Flags = []cli.Flag{
	verbosityFlag,
	logVmoduleFlag,
	vmoduleFlag,
	logjsonFlag,
	logFormatFlag,
	logFileFlag,
	logRotateFlag,
	logMaxSizeMBsFlag,
	logMaxBackupsFlag,
	logMaxAgeFlag,
	logCompressFlag,
	pprofFlag,
	pprofAddrFlag,
	pprofPortFlag,
	memprofilerateFlag,
	blockprofilerateFlag,
	cpuprofileFlag,
	traceFlag,
}

var (
	glogger       *log.GlogHandler
	logOutputFile io.WriteCloser
)

func init() {
	glogger = log.NewGlogHandler(log.NewTerminalHandler(os.Stderr, false))
}

// Setup initializes profiling and logging based on the CLI flags.
// It should be called as early as possible in the program.
func Setup(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

// Lumberjack uses <processname>-lumberjack.log in is.TempDir() if empty.
// so typically /tmp/geth-lumberjack.log on linux

// Retain backwards compatibility with `--log.json` flag if `--log.format` not set

// Unknown log format specified

// logging

// Retain backwards compatibility with `--vmodule` flag if `--log.vmodule` not set

// profiling, tracing

// pprof server

func StartPProf(address string) { _ = "STUB: not implemented"; return }

// Exit stops all running profiles, flushing their output to the
// respective file.
func Exit() { _ = "STUB: not implemented"; return }

func validateLogLocation(path string) error { _ = "STUB: not implemented"; return nil }

// Check if the path is writable by trying to create a temporary file
