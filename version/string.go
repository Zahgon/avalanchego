// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

// GitCommit is set in the build script at compile time
var GitCommit string

// Versions contains the versions relevant to a build of avalanchego. In
// addition to supporting construction of the string displayed by
// --version, it is used to produce the output of --version-json and can
// be used to unmarshal that output.
type Versions struct {
	Application string `json:"application"`
	Database    string `json:"database"`
	RPCChainVM  uint64 `json:"rpcchainvm"`
	// Commit may be empty if GitCommit was not set at compile time
	Commit string `json:"commit"`
	Go     string `json:"go"`
}

func GetVersions() *Versions { _ = "STUB: not implemented"; return nil }

func (v *Versions) String() string {
	_ = "STUB: not implemented"
	// This format maintains consistency with previous --version output
	return ""
}
