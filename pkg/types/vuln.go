// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package types

import "github.com/cruise-org/cruise/pkg/enums"

type Vulnerability struct {
	VulnID     string
	Pkg        string // PkgName
	Severity   enums.Severity
	Title      string
	Published  string
	PrimaryURL string
}
