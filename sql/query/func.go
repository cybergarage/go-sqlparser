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
	"strings"
)

// FunctionType represents a function type.
type FunctionType int

const (
	MathFunctionType FunctionType = iota
	AggregateFunctionType
	CastFunctionType
	ArithFunctionType
)

// FunctionExecutor represents a function executor interface.
type FunctionExecutor interface {
	// Name returns the name of the function.
	Name() string
	// Type returns the type of the function.
	Type() FunctionType
	// Execute returns the executed value with the specified arguments.
	Execute(...any) (any, error)
}

// Function represents a .
type Function interface {
	// Name returns the function name.
	Name() string
	// IsName returns true whether the function name is the specified one.
	IsName(string) bool
	// Arguments returns the argument list.
	Arguments() ArgumentList
	// IsAsterisk returns true if the argument list is "*".
	IsAsterisk() bool
	// Execute returns the executor of the function.
	Executor() (FunctionExecutor, error)
	// String returns a string representation of the function.
	String() string
}

// function represents a function interface.
type function struct {
	name string
	ArgumentList
}

// NewFunctionWith returns a function instance.
func NewFunctionWith(name string, args ...Argument) *function {
	fn := &function{
		name:         strings.ToUpper(name),
		ArgumentList: NewArgumentsWith(args...),
	}
	return fn
}

// Name returns the function name.
func (fn *function) Name() string {
	return fn.name
}

// IsName returns true whether the function name is the specified one.
func (fn *function) IsName(name string) bool {
	return fn.name == name
}

// IsAsterisk returns true if the argument list is "*".
func (fn *function) IsAsterisk() bool {
	l := len(fn.ArgumentList)
	switch {
	case l == 1:
		return fn.ArgumentList[0].IsAsterisk()
	case l == 0:
		return true
	}
	return false
}

// Executor returns the executor of the function.
func (fn *function) Executor() (FunctionExecutor, error) {
	return GetFunctionExecutor(fn.name)
}

// String returns a string representation of the function.
func (fn *function) String() string {
	return fn.name + "(" + fn.ArgumentList.String() + ")"
}
