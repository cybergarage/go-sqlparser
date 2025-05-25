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

// FunctionType represents a function type.
type FunctionType int

const (
	UnknownFunctionType FunctionType = iota
	MathFunction
	AggregateFunction
	CastFunction
	ArithFunction
)

// Function represents a .
type Function interface {
	// Name returns the function name.
	Name() string
	// IsName returns true whether the function name is the specified one.
	IsName(string) bool
	// Type returns the function type.
	Type() FunctionType
	// IsType returns true whether the function type is the specified one.
	IsType(FunctionType) bool
	// Arguments returns the argument list.
	Arguments() Arguments
	// Executor returns the executor of the function.
	Executor(opts ...ExecutorOption) (Executor, error)
	// IsAggregator returns true if the function is an aggregator function.
	IsAggregator() bool
	// Aggregator returns the aggregator of the function.
	Aggregator(opts ...AggregatorOption) (Aggregator, error)
	// String returns a string representation of the function.
	String() string
	// SelectorHelper provides additional methods for selectors.
	SelectorHelper
}

// SelectorHelper provides additional methods for selectors.
type SelectorHelper interface {
	// IsAsterisk returns true if the selector is an asterisk.
	IsAsterisk() bool
}
