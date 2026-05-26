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

// Package bind generates Ethereum contract Go bindings.
//
// Detailed usage document and tutorial available on the go-ethereum Wiki page:
// https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts

package precompilebind

import (
	"errors"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/accounts/abi/bind"
)

var errNoAnonymousEvent = errors.New("event type must not be anonymous")

const (
	ContractFileName     = "contract.go"
	ConfigFileName       = "config.go"
	ModuleFileName       = "module.go"
	EventFileName        = "event.go"
	ContractTestFileName = "contract_test.go"
	ConfigTestFileName   = "config_test.go"
)

type PrecompileBindFile struct {
	// FileName is the name of the file to be generated.
	FileName string
	// Content is the content of the file to be generated.
	Content string
	// IsTest indicates whether the file is a test file.
	IsTest bool
}

func NewPrecompileBindFile(fileName string, content string, isTest bool) PrecompileBindFile {
	_ = "STUB: not implemented"
	return *new(PrecompileBindFile)
}

// PrecompileBind generates a Go binding for a precompiled contract. It returns a slice of
// PrecompileBindFile structs containing the file name and its contents.
func PrecompileBind(types []string, abiData string, bytecodes []string, fsigs []map[string]string, pkg string, lang bind.Lang, libs map[string]string, aliases map[string]string, abifilename string, generateTests bool) ([]PrecompileBindFile, error) {
	_ = "STUB: not implemented"
	// create hooks
	return nil, nil
}

// createPrecompileHook creates a bind hook for precompiled contracts.
func createPrecompileHook(abifilename string, template string) bind.BindHook {
	_ = "STUB: not implemented"
	return *new(bind.BindHook)
}

// verify first

// these functions are not needed for binded contract.
// AllowList struct can provide the same functionality,
// so we don't need to generate them.

func allowListEnabled(funcs map[string]*bind.TmplMethod) bool {
	_ = "STUB: not implemented"
	return false
}

func verifyABI(abiData string) error {
	_ = "STUB: not implemented"
	// check abi first
	return nil
}
