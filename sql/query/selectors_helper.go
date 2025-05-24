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
	"github.com/cybergarage/go-sqlparser/sql/fn"
)

// IsAsterisk returns true if the selector list is "*".
func (selectors Selectors) IsAsterisk() bool {
	l := len(selectors)
	switch {
	case l == 1:
		return selectors[0].Name() == Asterisk
	case l == 0:
		return true
	}
	return false
}

// HasFunction returns true if the selector list has a function.
func (selectors Selectors) HasFunction() bool {
	for _, selector := range selectors {
		_, ok := selector.(Function)
		if ok {
			return true
		}
	}
	return false
}

// HasFunctionWithType returns true if the selector list has a function with the specified type.
func (selectors Selectors) HasFunctionWithType(t FunctionType) bool {
	for _, selector := range selectors {
		fx, ok := selector.(Function)
		if !ok {
			continue
		}
		if fx.Type() == t {
			return true
		}
	}
	return false
}

// HasAggregator returns true if the selector list has an aggregate function.
func (selectors Selectors) HasAggregator() bool {
	return selectors.HasFunctionWithType(fn.AggregateFunction)
}

// Functions returns a function array.
func (selectors Selectors) Functions() []Function {
	fns := make([]Function, 0)
	for _, selector := range selectors {
		fn, ok := selector.(Function)
		if !ok {
			continue
		}
		fns = append(fns, fn)
	}
	return fns
}

// LookupFunction returns a function with the specified name.
func (selectors Selectors) LookupFunction(name string) (Function, error) {
	for _, selector := range selectors {
		fn, ok := selector.(Function)
		if !ok {
			continue
		}
		if fn.IsName(name) {
			return fn, nil
		}
	}
	return nil, newErrNotFoundFunction(name)
}

// Executors returns a function executor array.
func (selectors Selectors) Executors() ([]FunctionExecutor, error) {
	executors := make([]FunctionExecutor, 0)
	for _, selector := range selectors {
		fn, ok := selector.(Function)
		if !ok {
			continue
		}
		executor, err := fn.Executor()
		if err != nil {
			return nil, err
		}
		executors = append(executors, executor)
	}
	return executors, nil
}

// LookupExecutor returns a function executor with the specified name.
func (selectors Selectors) LookupExecutor(name string) (FunctionExecutor, error) {
	fn, err := selectors.LookupFunction(name)
	if err != nil {
		return nil, err
	}
	return fn.Executor()
}

// ExecutorsForType returns a function executor array with the specified type.
func (selectors Selectors) ExecutorsForType(t FunctionType) ([]FunctionExecutor, error) {
	executors := make([]FunctionExecutor, 0)
	for _, selector := range selectors {
		fn, ok := selector.(Function)
		if !ok {
			continue
		}
		executor, err := fn.Executor()
		if err != nil {
			return nil, err
		}
		if executor.Type() != t {
			continue
		}
		executors = append(executors, executor)
	}
	return executors, nil
}

// AggregateFunctions returns an aggregate function array.
func (selectors Selectors) Aggregators() (AggregatorSet, error) {
	aggregators := make([]fn.Aggregator, 0)
	for _, selector := range selectors {
		fn, ok := selector.(Function)
		if !ok {
			continue
		}
		aggregator, err := fn.Aggregator()
		if err != nil {
			return nil, err
		}
		aggregators = append(aggregators, aggregator)
	}
	return fn.NewAggregatorSetWith(aggregators), nil
}
