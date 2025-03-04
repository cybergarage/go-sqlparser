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

// UnEscapeString returns an unescaped string.
func UnEscapeString(s string) string {
	return strings.Trim(s, "'\"")
}

// EscapeString returns an escaped string.
func EscapeString(s string) string {
	return fmt.Sprintf("'%v'", s)
}

// Equal returns true whether the specified strings are equal.
func Equal(s, t string) bool {
	return strings.Compare(s, t) == 0
}

// EqualFold returns true whether the specified strings are equal.
func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// UnEscapeNameString returns an unescaped column name string.
func UnEscapeNameString(s string) string {
	return strings.Trim(s, "'`()")
}
