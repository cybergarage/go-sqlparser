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

func TestJoinWithSpaceStrings(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"abc"}, "abc"},
		{[]string{"abc", "def"}, "abc def"},
	}

	for _, testCase := range testCases {
		actual := JoinWithSpace(testCase.input)
		if actual != testCase.expected {
			t.Errorf("actual (%v) != expected (%v)", actual, testCase.expected)
		}
	}
}

func TestJoinWithCommaStrings(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"abc"}, "abc"},
		{[]string{"abc", "def"}, "abc, def"},
	}

	for _, testCase := range testCases {
		actual := JoinWithComma(testCase.input)
		if actual != testCase.expected {
			t.Errorf("actual (%v) != expected (%v)", actual, testCase.expected)
		}
	}
}

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

func TestSplitDataTypeString(t *testing.T) {
	testCases := []struct {
		inputDataTypeString string
		expectedDataType    string
		expectedSize        int
	}{
		{"", "", -1},
		{"abc", "abc", -1},
		{"abc(10)", "abc", 10},
	}

	for _, testCase := range testCases {
		actualDataType, actualSize := SplitDataTypeString(testCase.inputDataTypeString)
		if actualDataType != testCase.expectedDataType {
			t.Errorf("actual (%v) != expected (%v)", actualDataType, testCase.expectedDataType)
		}
		if actualSize != testCase.expectedSize {
			t.Errorf("actual (%v) != expected (%v)", actualSize, testCase.expectedSize)
		}
	}
}
