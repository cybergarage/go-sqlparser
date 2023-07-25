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

package strings

import (
	"testing"
)

func TestUnEscapeStrings(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"abc", "abc"},
		{"'abc'", "abc"},
	}

	for _, testCase := range testCases {
		actual := UnEscapeString(testCase.input)
		if actual != testCase.expected {
			t.Errorf("actual (%v) != expected (%v)", actual, testCase.expected)
		}
	}
}

func TestEscapeStrings(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "''"},
		{"abc", "'abc'"},
	}

	for _, testCase := range testCases {
		actual := EscapeString(testCase.input)
		if actual != testCase.expected {
			t.Errorf("actual (%v) != expected (%v)", actual, testCase.expected)
		}
	}
}
