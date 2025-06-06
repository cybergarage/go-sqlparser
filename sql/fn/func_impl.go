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
	"strings"
)

// function represents a function interface.
type function struct {
	typ        FunctionType
	name       string
	executor   Executor
	aggregator Aggregator
	args       Arguments
}

// FunctionOption represents a function option function.
type FunctionOption = func(*function)

// NewFunctionWith returns a function instance.
func NewFunctionWith(opts ...FunctionOption) Function {
	fn := &function{
		typ:        UnknownFunctionType,
		name:       "",
		executor:   nil,
		aggregator: nil,
		args:       NewArguments(),
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
		t, err := NewFunctionTypeForName(name)
		if err == nil {
			fn.typ = t
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
func WithFunctionExecutor(executor Executor) FunctionOption {
	return func(fn *function) {
		fn.executor = executor
		fn.name = executor.Name()
		fn.typ = executor.Type()
	}
}

// WithFunctionAggregator sets the function aggregator.
func WithFunctionAggregator(aggregator Aggregator) FunctionOption {
	return func(fn *function) {
		fn.name = aggregator.Name()
		fn.typ = AggregateFunction
	}
}

// WithFunctionArguments sets the function arguments.
func WithFunctionArguments(args ...Argument) FunctionOption {
	return func(fn *function) {
		fn.args = NewArgumentsWith(args...)
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

// Type returns the function type.
func (fn *function) Type() FunctionType {
	return fn.typ
}

// IsType returns true whether the function type is the specified one.
func (fn *function) IsType(t FunctionType) bool {
	return fn.typ == t
}

// IsAsterisk returns true whether the function is an asterisk.
func (fn *function) IsAsterisk() bool {
	return false
}

// IsFunction returns true whether the function is a function.
func (fn *function) IsFunction() bool {
	return true
}

// Function returns the function if the function is a function.
func (fn *function) Function() (Function, bool) {
	return fn, true
}

// Arguments returns the argument list.
func (fn *function) Arguments() Arguments {
	return fn.args
}

// IsAggregator returns true if the function is an aggregator function.
func (fn *function) IsAggregator() bool {
	return fn.typ == AggregateFunction
}

// Executor returns the executor of the function.
func (fn *function) Executor(opts ...ExecutorOption) (Executor, error) {
	if fn.executor != nil {
		return fn.executor, nil
	}
	opts = append(opts, WithExecutorArguments(fn.args.Names()))
	executor, err := NewExecutorForName(fn.name, opts...)
	if err != nil {
		return nil, err
	}
	fn.executor = executor
	fn.typ = executor.Type()
	return fn.executor, nil
}

// Aggregator returns the aggregator of the function.
func (fn *function) Aggregator(opts ...AggregatorOption) (Aggregator, error) {
	if fn.aggregator != nil {
		return fn.aggregator, nil
	}
	opts = append(opts, WithAggregatorArguments(fn.args.Names()))
	aggregator, err := NewAggregatorForName(fn.name, opts...)
	if err != nil {
		return nil, err
	}
	fn.aggregator = aggregator
	fn.typ = AggregateFunction
	return fn.aggregator, nil
}

// String returns a string representation of the function.
func (fn *function) String() string {
	return fn.name + "(" + fn.args.String() + ")"
}
