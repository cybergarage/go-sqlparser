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
)

// ExecutorFunc is a function type that takes any number of arguments and returns a value and an error.
type ExecutorFunc func(...any) (any, error)

// ExecutorOption is a function that modifies the execImpl.
type ExecutorOption func(*execImpl)

// execImpl represents a base math function.
type execImpl struct {
	name string
	args []string
	t    FunctionType
	fn   ExecutorFunc
}

// WithExecutorName sets the name for the executor.
func WithExecutorName(name string) ExecutorOption {
	return func(ex *execImpl) {
		ex.name = name
	}
}

// WithExecutorArguments sets the arguments for the executor.
func WithExecutorArguments(args []string) ExecutorOption {
	return func(ex *execImpl) {
		ex.args = args
	}
}

// WithExecutorType sets the type for the executor.
func WithExecutorType(t FunctionType) ExecutorOption {
	return func(ex *execImpl) {
		ex.t = t
	}
}

// WithExecutorFunc sets the function for the executor.
func WithExecutorFunc(fn ExecutorFunc) ExecutorOption {
	return func(ex *execImpl) {
		ex.fn = fn
	}
}

// NewExecutor returns a new executor with the specified name and type.
func NewExecutor(opts ...ExecutorOption) *execImpl {
	ex := &execImpl{
		name: "",
		t:    UnknownFunctionType,
		args: make([]string, 0),
		fn: func(args ...any) (any, error) {
			return nil, fmt.Errorf("executor function not implemented")
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

func (ex *execImpl) ExecuteArgs(args []any) (any, error) {
	return ex.fn(args...)
}

func (ex *execImpl) ExecuteMap(m map[string]any) (any, error) {
	row := make([]any, 0, len(ex.args))
	for _, arg := range ex.args {
		v, ok := m[arg]
		if !ok {
			v = arg
		}
		row = append(row, v)
	}
	return ex.ExecuteArgs(row)
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
	case nil:
		return ex.fn()
	default:
		return ex.fn(v)
	}
}
