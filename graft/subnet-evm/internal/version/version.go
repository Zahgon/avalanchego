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
// Copyright 2022 The go-ethereum Authors
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

// Package version implements reading of build version information.
package version

import (
	"runtime/debug"
)

const ourPath = "github.com/ava-labs/avalanchego/graft/subnet-evm" // Path to our module

// These variables are set at build-time by the linker when the build is
// done by build/ci.go.
var gitCommit, gitDate string

// VCSInfo represents the git repository state.
type VCSInfo struct {
	Commit string // head commit hash
	Date   string // commit time in YYYYMMDD format
	Dirty  bool
}

// VCS returns version control information of the current executable.
func VCS() (VCSInfo, bool) {
	_ = "STUB: not implemented"
	return *

	// Use information set by the build script if present.
	new(VCSInfo), false
}

// ClientName creates a software name/version identifier according to common
// conventions in the Ethereum p2p network.
func ClientName(clientIdentifier string) string { _ = "STUB: not implemented"; return "" }

// runtimeInfo returns build and platform information about the current binary.
//
// If the package that is currently executing is a prefixed by our go-ethereum
// module path, it will print out commit and date VCS information. Otherwise,
// it will assume it's imported by a third-party and will return the imported
// version and whether it was replaced by another module.
func Info() (version, vcs string) { _ = "STUB: not implemented"; return "", "" }

// versionInfo returns version information for the currently executing
// implementation.
//
// Depending on how the code is instantiated, it returns different amounts of
// information. If it is unable to determine which module is related to our
// package it falls back to the hardcoded values in the params package.
func versionInfo(info *debug.BuildInfo) string {
	_ = "STUB: not implemented"
	// If the main package is from our repo, prefix version with "geth".
	return ""
}

// Not our main package, so explicitly print out the module path and
// version.

// These can be empty when invoked with "go run".

// If our module path wasn't imported, it's unclear which
// version of our code they are running. Fallback to hardcoded
// version.

// Our package is a dependency for the main module. Return path and
// version data for both.

// If our package was replaced by something else, also note that.

// findModule returns the module at path.
func findModule(info *debug.BuildInfo, path string) *debug.Module {
	_ = "STUB: not implemented"
	return nil
}
