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
	"fmt"
	"strings"
)

// function represents a function interface.
type function struct {
	typ      FunctionType
	name     string
	executor FunctionExecutor
	ArgumentList
}

// FunctionOption represents a function option function.
type FunctionOption = func(*function)

// NewFunctionWith returns a function instance.
func NewFunctionWith(opts ...FunctionOption) Function {
	fn := &function{
		typ:          UnknownFunctionType,
		name:         "",
		executor:     nil,
		ArgumentList: NewArguments(),
	}
	for _, opt := range opts {
		opt(fn)
	}
	return fn
}

// WithFunctionName sets the function name.
func WithFunctionName(name string) FunctionOption {
	return func(fn *function) {
		fn.SetName(name)
		executor, err := GetFunctionExecutor(name)
		if err == nil {
			fn.executor = executor
			fn.typ = executor.Type()
		}
	}
}

// WithFunctionType sets the function type.
func WithFunctionType(t FunctionType) FunctionOption {
	return func(fn *function) {
		fn.typ = t
	}
}

// WithFunctionExecutor sets the function executor.
func WithFunctionExecutor(executor FunctionExecutor) FunctionOption {
	return func(fn *function) {
		fn.executor = executor
		fn.name = executor.Name()
		fn.typ = executor.Type()
	}
}

// WithFunctionArguments sets the function arguments.
func WithFunctionArguments(args ...Argument) FunctionOption {
	return func(fn *function) {
		fn.ArgumentList = NewArgumentsWith(args...)
	}
}

// SetName sets the function name.
func (fn *function) SetName(name string) {
	fn.name = strings.ToUpper(name)
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

// Type returns the function type.
func (fn *function) Type() FunctionType {
	return fn.typ
}

// IsType returns true whether the function type is the specified one.
func (fn *function) IsType(t FunctionType) bool {
	return fn.typ == t
}

// Executor returns the executor of the function.
func (fn *function) Executor() (FunctionExecutor, error) {
	if fn.executor != nil {
		return fn.executor, nil
	}
	return GetFunctionExecutor(fn.name)
}

// ExecuteUpdator executes the executor with the specified row.
func (fn *function) Execute(col Column, row map[string]any) (any, error) {
	newErrInvalidArgs := func(fn Function, col Column, row map[string]any) error {
		return fmt.Errorf("%v is %w arguments (%s:%s)", fn.String(), ErrInvalid, col.String(), row)
	}

	executor, err := fn.Executor()
	if err != nil {
		return nil, err
	}

	args := col.Arguments()
	if len(args) < 2 {
		return nil, newErrInvalidArgs(fn, col, row)
	}
	leftExprName, ok := args[0].(string)
	if !ok {
		return nil, newErrInvalidArgs(fn, col, row)
	}
	v, ok := row[leftExprName]
	if !ok {
		return nil, newErrInvalidArgs(fn, col, row)
	}
	args[0] = v
	rv, err := executor.Execute(args...)
	if err != nil {
		return nil, err
	}

	return rv, nil
}

// String returns a string representation of the function.
func (fn *function) String() string {
	return fn.name + "(" + fn.ArgumentList.String() + ")"
}
