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
	RoundFunctionName = "ROUND"
	SqrtFunctionName  = "SQRT"
	LogFunctionName   = "LOG"
	Log10FunctionName = "LOG10"
	ExpFunctionName   = "EXP"
	PowerFunctionName = "POWER"
	ModFunctionName   = "MOD"
	SinFunctionName   = "SIN"
	CosFunctionName   = "COS"
	TanFunctionName   = "TAN"
	RandFunctionName  = "RAND"
	PiFunctionName    = "PI"
	// String.
	UpperFunctionName = "UPPER"
	LowerFunctionName = "LOWER"
	TrimFunctionName  = "TRIM"
	// Time.
	CurrentTimestampFunctionName = "CURRENT_TIMESTAMP"
	NowFunctionName              = "NOW"
	// Aggregate.
	MaxFunctionName   = "MAX"
	MinFunctionName   = "MIN"
	SumFunctionName   = "SUM"
	AvgFunctionName   = "AVG"
	CountFunctionName = "COUNT"
	// Airthmetic.
	AddOperatorID = "+"
	SubOperatorID = "-"
	MulOperatorID = "*"
	DivOperatorID = "/"
	ModOperatorID = "%"
)

var registeredFunctionNames = []string{
	// Math.
	AbsFunctionName,
	CeilFunctionName,
	FloorFunctionName,
	RoundFunctionName,
	SqrtFunctionName,
	LogFunctionName,
	Log10FunctionName,
	ExpFunctionName,
	PowerFunctionName,
	ModFunctionName,
	SinFunctionName,
	CosFunctionName,
	TanFunctionName,
	RandFunctionName,
	PiFunctionName,
	// String.
	UpperFunctionName,
	LowerFunctionName,
	TrimFunctionName,
	// Time.
	CurrentTimestampFunctionName,
	NowFunctionName,
	// Aggregate.
	MaxFunctionName,
	MinFunctionName,
	SumFunctionName,
	AvgFunctionName,
	CountFunctionName,
	// Airthmetic.
	AddOperatorID,
	SubOperatorID,
	MulOperatorID,
	DivOperatorID,
	ModOperatorID,
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
