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

package fn_test

import (
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/fn"
)

func TestIsRegisteredFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Exact matches
		{"ABS uppercase", "ABS", true},
		{"abs lowercase", "abs", true},
		{"Ceil mixed case", "CeIl", true},
		{"MAX", "MAX", true},
		{"min", "min", true},
		{"SUM", "SUM", true},
		{"avg", "avg", true},
		{"COUNT", "COUNT", true},
		// Operator matches
		{"Add operator", "+", true},
		{"Sub operator", "-", true},
		{"Mul operator", "*", true},
		{"Div operator", "/", true},
		{"Mod operator", "%", true},
		// Function call patterns
		{"ABS function call", "ABS(", true},
		{"SUM function call", "SUM(", true},
		{"COUNT function call", "COUNT(", true},
		{"MIN function call", "MIN(", true},
		{"MAX function call", "MAX(", true},
		{"AVG function call", "AVG(", true},
		{"CEIL function call", "CEIL(", true},
		{"FLOOR function call", "FLOOR(", true},
		// Suffix with parenthesis
		{"ABS with closing paren", "ABS)", true},
		{"SUM with closing paren", "SUM)", true},
		// Invalid/unknown
		{"Unknown function", "FOO", false},
		{"Empty string", "", false},
		{"Whitespace", " ", false},
		{"Partial match", "AB", false},
		{"Operator with space", "+ ", false},
		{"Function with extra chars", "SUMX", false},
		{"Function with prefix", "XSUM", false},
		{"Operator with parenthesis", "+(", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fn.IsRegisteredFunction(tt.input)
			if result != tt.expected {
				t.Errorf("IsRegisteredFunction(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
