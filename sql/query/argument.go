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

// Argument represents am argument in a function.
type Argument string

// NewArgument returns a new argument instance from the specified string.
func NewArgumentWith(arg string) Argument {
	return Argument(arg)
}

// Name returns the name of the argument.
func (arg Argument) Name() string {
	return string(arg)
}

// IsName returns true whether the argument name is the specified one.
func (arg Argument) IsName(name string) bool {
	return arg.Name() == name
}

// IsAsterisk returns true whether the argument name is the asterisk.
func (arg Argument) IsAsterisk() bool {
	return arg.IsName("*")
}

// String returns the string representation of the argument.
func (arg Argument) String() string {
	return string(arg)
}
