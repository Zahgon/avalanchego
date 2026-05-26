// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"
	"sync"
)

var _ fmt.Stringer = (*Application)(nil)

type Application struct {
	Name  string `json:"name"  yaml:"name"`
	Major int    `json:"major" yaml:"major"`
	Minor int    `json:"minor" yaml:"minor"`
	Patch int    `json:"patch" yaml:"patch"`

	makeStrOnce sync.Once
	str         string
}

func (a *Application) String() string { _ = "STUB: not implemented"; return "" }

func (a *Application) initString() { _ = "STUB: not implemented"; return }

// Semantic returns the semantic version string (e.g., "v1.14.1")
func (a *Application) Semantic() string { _ = "STUB: not implemented"; return "" }

// SemanticWithCommit returns the semantic version string with an optional git commit suffix
func (a *Application) SemanticWithCommit(gitCommit string) string {
	_ = "STUB: not implemented"
	return ""
}

// Compare returns
//
//	-1 if a is less than o,
//	 0 if a equals o,
//	+1 if a is greater than o.
func (a *Application) Compare(o *Application) int { _ = "STUB: not implemented"; return 0 }
