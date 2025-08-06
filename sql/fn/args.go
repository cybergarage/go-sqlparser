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
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// Arguments represens an argument array in a function.
type Arguments []Argument

// NewArguments returns an empty argument array instance.
func NewArguments() Arguments {
	return make(Arguments, 0)
}

// NewArgumentsWith returns an argument array instance with the specified arguments.
func NewArgumentsWith(args ...Argument) Arguments {
	list := make(Arguments, len(args))
	copy(list, args)
	return list
}

// NewArgumentStrings returns an argument array instance with the specified string arguments.
func NewArgumentStrings(strs ...string) Arguments {
	args := make(Arguments, len(strs))
	for n, arg := range strs {
		args[n] = NewArgumentWith(arg)
	}
	return args
}

// IsAsterisk returns true if the argument list is an asterisk.
func (args Arguments) IsAsterisk() bool {
	l := len(args)
	switch l {
	case 1:
		return args[0].IsAsterisk()
	case 0:
		return true
	}
	return false
}

// Name returns the name of the first argument in the list.
func (args Arguments) Names() []string {
	names := make([]string, len(args))
	for n, arg := range args {
		names[n] = arg.Name()
	}
	return names
}

// String returns a string representation of the the argument list.
func (args Arguments) String() string {
	strs := make([]string, len(args))
	for n, str := range args {
		strs[n] = string(str)
	}
	return strings.JoinWithComma(strs)
}
