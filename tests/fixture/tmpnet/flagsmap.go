// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

// Defines a mapping of flag keys to values intended to be supplied to
// an invocation of an AvalancheGo node.
type FlagsMap map[string]string

// Utility function simplifying construction of a FlagsMap from a file.
func ReadFlagsMap(path string, description string) (FlagsMap, error) {
	_ = "STUB: not implemented"
	return *new(FlagsMap), nil
}

// SetDefault ensures the effectiveness of a flag override by only
// setting a value supplied whose key is not already explicitly set.
func (f FlagsMap) SetDefault(key string, value string) { _ = "STUB: not implemented"; return }

// SetDefaults ensures the effectiveness of flag overrides by only
// setting values supplied in the defaults map that are not already
// explicitly set.
func (f FlagsMap) SetDefaults(defaults FlagsMap) { _ = "STUB: not implemented"; return }

// Write simplifies writing a FlagsMap to the provided path. The
// description is used in error messages.
func (f FlagsMap) Write(path string, description string) error {
	_ = "STUB: not implemented"
	return nil
}
