// Copyright (C) 2022 The go-sqlparser Authors All rights reserved.
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
	"fmt"
	"strconv"
	"strings"
)

// JoinWithSpace returns a string by joining the specified elements with a space.
func JoinWithSpace(elems []string) string {
	return strings.Join(elems, " ")
}

// JoinWithComma returns a string by joining the specified elements with a comma.
func JoinWithComma(elems []string) string {
	return strings.Join(elems, ", ")
}

// stringLiteralTrimSet represents a string literal trim set.
const stringLiteralTrimSet = "'\""

// UnEscapeString returns an unescaped string.
func UnEscapeString(s string) string {
	return strings.Trim(s, stringLiteralTrimSet)
}

// EscapeString returns an escaped string.
func EscapeString(s string) string {
	return fmt.Sprintf("'%v'", s)
}

// SplitDataTypeString returns a data type and a size from the specified string.
func SplitDataTypeString(s string) (string, int) {
	sep := " "
	reps := []string{"(", ")"}
	for _, rep := range reps {
		s = strings.ReplaceAll(s, rep, sep)
	}
	elems := strings.Split(s, sep)
	if len(elems) == 0 {
		return "", 0
	}
	dataType := elems[0]
	dataSize := -1
	if 2 <= len(elems) {
		i, err := strconv.Atoi(elems[1])
		if err == nil {
			dataSize = i
		}
	}
	return dataType, dataSize
}
