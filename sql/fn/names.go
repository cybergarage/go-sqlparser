// Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fn

import (
	"strings"
)

const (
	// Math.
	AbsFunctionName   = "ABS"
	CeilFunctionName  = "CEIL"
	FloorFunctionName = "FLOOR"
	// Aggregate.
	MaxFunctionName   = "MAX"
	MinFunctionName   = "MIN"
	SumFunctionName   = "SUM"
	AvgFunctionName   = "AVG"
	CountFunctionName = "COUNT"
	// Airthmetic.
	AddFunctionID = "+"
	SubFunctionID = "-"
	MulFunctionID = "*"
	DivFunctionID = "/"
	ModFunctionID = "%"
)

var registeredFunctionNames = []string{
	AbsFunctionName,
	CeilFunctionName,
	FloorFunctionName,
	MaxFunctionName,
	MinFunctionName,
	SumFunctionName,
	AvgFunctionName,
	CountFunctionName,
	AddFunctionID,
	SubFunctionID,
	MulFunctionID,
	DivFunctionID,
	ModFunctionID,
}

// RegisteredFunctionNames returns a list of all registered function names.
func RegisteredFunctionNames() []string {
	return registeredFunctionNames
}

// IsRegisteredFunctionNamePattern checks if the given name matches any registered function name pattern.
func IsRegisteredFunction(name string) bool {
	if len(name) == 0 {
		return false
	}
	name = strings.ToUpper(name)
	for _, fn := range registeredFunctionNames {
		switch len(fn) {
		case 1: // Single character functions like '+', '-', etc.
			if name == fn {
				return true
			}
		default: // Functions with multiple characters
			if len(name) == len(fn) {
				if name == fn {
					return true
				}
			}
			if strings.HasPrefix(name, fn+"(") || strings.HasSuffix(name, ")") {
				return true
			}
		}
	}
	return false
}
