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

package query

import (
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// ArgumentList represens an argument array in a function.
type ArgumentList []Argument

// NewArgumentsWith returns a argument array instance with the specified arguments.
func NewArgumentsWith(args ...Argument) ArgumentList {
	list := make(ArgumentList, len(args))
	copy(list, args)
	return list
}

// Argument returns an argument array.
func (arguments ArgumentList) Arguments() ArgumentList {
	return arguments
}

// String returns a string representation of the the argument list.
func (arguments ArgumentList) String() string {
	strs := make([]string, len(arguments))
	for n, str := range arguments {
		strs[n] = string(str)
	}
	return strings.JoinWithComma(strs)
}
