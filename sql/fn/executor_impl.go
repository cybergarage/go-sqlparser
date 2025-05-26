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
	"fmt"
	"strings"
)

// execFunc is a function type that takes any number of arguments and returns a value and an error.
type execFunc func(...any) (any, error)

// execOption is a function that modifies the execImpl.
type execOption func(*execImpl)

// execImpl represents a base math function.
type execImpl struct {
	name string
	args []string
	t    FunctionType
	fn   execFunc
}

// withExecName sets the name for the execImpl.
func withExecName(name string) execOption {
	return func(ex *execImpl) {
		ex.name = name
	}
}

// withExecArguments sets the arguments for the execImpl.
func withExecArguments(args []string) execOption {
	return func(ex *execImpl) {
		ex.args = args
	}
}

// NewMathFunctionWith returns a new base math function with the specified name and math.
func newExecWith(name string, t FunctionType, opts ...execOption) *execImpl {
	ex := &execImpl{
		name: strings.ToUpper(name),
		t:    t,
		args: make([]string, 0),
		fn: func(args ...any) (any, error) {
			return nil, fmt.Errorf("function %s not implemented", name)
		},
	}
	for _, opt := range opts {
		opt(ex)
	}
	return ex
}

// Name returns the name of the function.
func (ex *execImpl) Name() string {
	return ex.name
}

// Type returns the type of the function.
func (ex *execImpl) Type() FunctionType {
	return ex.t
}

// Arguments returns the arguments of the executor.
func (ex *execImpl) Arguments() []string {
	return ex.args
}

func (ex *execImpl) ExecuteArgs(args ...any) (any, error) {
	return ex.fn(args...)
}

func (ex *execImpl) ExecuteMap(m map[string]any) (any, error) {
	row := make([]any, 0, len(ex.args))
	for _, arg := range ex.args {
		value, ok := m[arg]
		if !ok {
			return nil, fmt.Errorf("%w column %s not found in map", ErrNotFound, arg)
		}
		row = append(row, value)
	}
	return ex.ExecuteArgs(row...)
}

// Execute executes the function with the provided value.
func (ex *execImpl) Execute(v any) (any, error) {
	switch v := v.(type) {
	case []any:
		return ex.ExecuteArgs(v)
	case map[string]any:
		return ex.ExecuteMap(v)
	case Row:
		return ex.ExecuteArgs(v)
	case Map:
		return ex.ExecuteMap(v)
	default:
		return ex.ExecuteArgs(v)
	}
}
