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
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	_ "embed"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/internal/flags"
	"github.com/ava-labs/libevm/log"
	"github.com/urfave/cli/v2"
)

//go:embed template-readme.md
var readme string

var (
	// Flags needed by abigen
	abiFlag = &cli.StringFlag{
		Name:  "abi",
		Usage: "Path to the contract ABI json to generate, - for STDIN",
	}
	typeFlag = &cli.StringFlag{
		Name:  "type",
		Usage: "Struct name for the precompile (default = {abi file name})",
	}
	pkgFlag = &cli.StringFlag{
		Name:  "pkg",
		Usage: "Go package name to generate the precompile into (default = {type})",
	}
	outFlag = &cli.StringFlag{
		Name:  "out",
		Usage: "Output folder for the generated precompile files, - for STDOUT (default = ./precompile/contracts/{pkg}). Test files won't be generated if STDOUT is used",
	}
)

var app = flags.NewApp("subnet-evm precompile generator tool")

func init() {
	app.Name = "precompilegen"
	app.Flags = []cli.Flag{
		abiFlag,
		outFlag,
		pkgFlag,
		typeFlag,
	}
	app.Action = precompilegen
}

func precompilegen(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// If the entire solidity code was specified, build and bind based on that

// Load up the ABI

// we should not generate the abi file if output is set to stdout

// get file name from the output path

// if output is set to stdout, we should not generate the test codes

// Generate the contract precompile

// Either flush it out to a file or display on the standard output
// Skip displaying test codes here.

// Create your file

// Write the ABI to the output folder

// Write the README to the output folder

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
