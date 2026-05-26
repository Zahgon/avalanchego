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
// Copyright 2015 The go-ethereum Authors
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

package flags

import (
	"encoding"
	"flag"
	"math/big"

	"github.com/urfave/cli/v2"
)

// DirectoryString is custom type which is registered in the flags library which cli uses for
// argument parsing. This allows us to expand Value to an absolute path when
// the argument is parsed
type DirectoryString string

func (s *DirectoryString) String() string { _ = "STUB: not implemented"; return "" }

func (s *DirectoryString) Set(value string) error { _ = "STUB: not implemented"; return nil }

var (
	_ cli.Flag              = (*DirectoryFlag)(nil)
	_ cli.RequiredFlag      = (*DirectoryFlag)(nil)
	_ cli.VisibleFlag       = (*DirectoryFlag)(nil)
	_ cli.DocGenerationFlag = (*DirectoryFlag)(nil)
	_ cli.CategorizableFlag = (*DirectoryFlag)(nil)
)

// DirectoryFlag is custom cli.Flag type which expand the received string to an absolute path.
// e.g. ~/.ethereum -> /home/username/.ethereum
type DirectoryFlag struct {
	Name string

	Category    string
	DefaultText string
	Usage       string

	Required   bool
	Hidden     bool
	HasBeenSet bool

	Value DirectoryString

	Aliases []string
	EnvVars []string
}

// For cli.Flag:

func (f *DirectoryFlag) Names() []string { _ = "STUB: not implemented"; return nil }
func (f *DirectoryFlag) IsSet() bool     { _ = "STUB: not implemented"; return false }
func (f *DirectoryFlag) String() string  { _ = "STUB: not implemented"; return "" }

// Apply called by cli library, grabs variable from environment (if in env)
// and adds variable to flag set for parsing.
func (f *DirectoryFlag) Apply(set *flag.FlagSet) error { _ = "STUB: not implemented"; return nil }

// For cli.RequiredFlag:

func (f *DirectoryFlag) IsRequired() bool {
	_ = "STUB: not implemented"

	// For cli.VisibleFlag:
	return false
}

func (f *DirectoryFlag) IsVisible() bool {
	_ = "STUB: not implemented"

	// For cli.CategorizableFlag:
	return false
}

func (f *DirectoryFlag) GetCategory() string {
	_ = "STUB: not implemented"

	// For cli.DocGenerationFlag:
	return ""
}

func (f *DirectoryFlag) TakesValue() bool     { _ = "STUB: not implemented"; return false }
func (f *DirectoryFlag) GetUsage() string     { _ = "STUB: not implemented"; return "" }
func (f *DirectoryFlag) GetValue() string     { _ = "STUB: not implemented"; return "" }
func (f *DirectoryFlag) GetEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (f *DirectoryFlag) GetDefaultText() string { _ = "STUB: not implemented"; return "" }

type TextMarshaler interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// textMarshalerVal turns a TextMarshaler into a flag.Value
type textMarshalerVal struct {
	v TextMarshaler
}

func (v textMarshalerVal) String() string { _ = "STUB: not implemented"; return "" }

func (v textMarshalerVal) Set(s string) error { _ = "STUB: not implemented"; return nil }

var (
	_ cli.Flag              = (*TextMarshalerFlag)(nil)
	_ cli.RequiredFlag      = (*TextMarshalerFlag)(nil)
	_ cli.VisibleFlag       = (*TextMarshalerFlag)(nil)
	_ cli.DocGenerationFlag = (*TextMarshalerFlag)(nil)
	_ cli.CategorizableFlag = (*TextMarshalerFlag)(nil)
)

// TextMarshalerFlag wraps a TextMarshaler value.
type TextMarshalerFlag struct {
	Name string

	Category    string
	DefaultText string
	Usage       string

	Required   bool
	Hidden     bool
	HasBeenSet bool

	Value TextMarshaler

	Aliases []string
	EnvVars []string
}

// For cli.Flag:

func (f *TextMarshalerFlag) Names() []string { _ = "STUB: not implemented"; return nil }
func (f *TextMarshalerFlag) IsSet() bool     { _ = "STUB: not implemented"; return false }
func (f *TextMarshalerFlag) String() string  { _ = "STUB: not implemented"; return "" }

func (f *TextMarshalerFlag) Apply(set *flag.FlagSet) error { _ = "STUB: not implemented"; return nil }

// For cli.RequiredFlag:

func (f *TextMarshalerFlag) IsRequired() bool {
	_ = "STUB: not implemented"

	// For cli.VisibleFlag:
	return false
}

func (f *TextMarshalerFlag) IsVisible() bool {
	_ = "STUB: not implemented"

	// For cli.CategorizableFlag:
	return false
}

func (f *TextMarshalerFlag) GetCategory() string {
	_ = "STUB: not implemented"

	// For cli.DocGenerationFlag:
	return ""
}

func (f *TextMarshalerFlag) TakesValue() bool     { _ = "STUB: not implemented"; return false }
func (f *TextMarshalerFlag) GetUsage() string     { _ = "STUB: not implemented"; return "" }
func (f *TextMarshalerFlag) GetEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (f *TextMarshalerFlag) GetValue() string { _ = "STUB: not implemented"; return "" }

func (f *TextMarshalerFlag) GetDefaultText() string { _ = "STUB: not implemented"; return "" }

// GlobalTextMarshaler returns the value of a TextMarshalerFlag from the global flag set.
func GlobalTextMarshaler(ctx *cli.Context, name string) TextMarshaler {
	_ = "STUB: not implemented"
	return *new(TextMarshaler)
}

var (
	_ cli.Flag              = (*BigFlag)(nil)
	_ cli.RequiredFlag      = (*BigFlag)(nil)
	_ cli.VisibleFlag       = (*BigFlag)(nil)
	_ cli.DocGenerationFlag = (*BigFlag)(nil)
	_ cli.CategorizableFlag = (*BigFlag)(nil)
)

// BigFlag is a command line flag that accepts 256 bit big integers in decimal or
// hexadecimal syntax.
type BigFlag struct {
	Name string

	Category    string
	DefaultText string
	Usage       string

	Required   bool
	Hidden     bool
	HasBeenSet bool

	Value        *big.Int
	defaultValue *big.Int

	Aliases []string
	EnvVars []string
}

// For cli.Flag:

func (f *BigFlag) Names() []string { _ = "STUB: not implemented"; return nil }
func (f *BigFlag) IsSet() bool     { _ = "STUB: not implemented"; return false }
func (f *BigFlag) String() string  { _ = "STUB: not implemented"; return "" }

func (f *BigFlag) Apply(set *flag.FlagSet) error {
	_ = "STUB: not implemented"
	// Set default value so that environment wont be able to overwrite it
	return nil
}

// For cli.RequiredFlag:

func (f *BigFlag) IsRequired() bool {
	_ = "STUB: not implemented"

	// For cli.VisibleFlag:
	return false
}

func (f *BigFlag) IsVisible() bool {
	_ = "STUB: not implemented"

	// For cli.CategorizableFlag:
	return false
}

func (f *BigFlag) GetCategory() string {
	_ = "STUB: not implemented"

	// For cli.DocGenerationFlag:
	return ""
}

func (f *BigFlag) TakesValue() bool     { _ = "STUB: not implemented"; return false }
func (f *BigFlag) GetUsage() string     { _ = "STUB: not implemented"; return "" }
func (f *BigFlag) GetValue() string     { _ = "STUB: not implemented"; return "" }
func (f *BigFlag) GetEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (f *BigFlag) GetDefaultText() string { _ = "STUB: not implemented"; return "" }

// bigValue turns *big.Int into a flag.Value
type bigValue big.Int

func (b *bigValue) String() string { _ = "STUB: not implemented"; return "" }

func (b *bigValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

// GlobalBig returns the value of a BigFlag from the global flag set.
func GlobalBig(ctx *cli.Context, name string) *big.Int { _ = "STUB: not implemented"; return nil }

// Expands a file path
// 1. replace tilde with users home dir
// 2. expands embedded environment variables
// 3. cleans the path, e.g. /a/b/../c -> /a/c
// Note, it has limitations, e.g. ~someuser/tmp will not be expanded
func expandPath(p string) string {
	_ = "STUB: not implemented"
	// Named pipes are not file paths on windows, ignore
	return ""
}

func HomeDir() string { _ = "STUB: not implemented"; return "" }

func eachName(f cli.Flag, fn func(string)) { _ = "STUB: not implemented"; return }
