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

// Selectors represens a selector array.
type Selectors []Selector

// NewSelectors returns a selector array instance.
func NewSelectors() Selectors {
	return make(Selectors, 0)
}

// NewSelectorsWith returns a selector array instance with the specified selectors.
func NewSelectorsWith(selectors ...Selector) Selectors {
	s := make(Selectors, len(selectors))
	copy(s, selectors)
	return s
}

// NewSelectorsWithColums returns a selector array instance with the specified selectors.
func NewSelectorsWithColums(selectors ...Column) Selectors {
	s := make(Selectors, len(selectors))
	for n, selector := range selectors {
		s[n] = selector
	}
	return s
}

// Selector returns a selector array.
func (selectors Selectors) Selectors() Selectors {
	return selectors
}

// Column returns a column array.
func (selectors Selectors) Columns() []Column {
	cols := make([]Column, 0)
	for _, selector := range selectors {
		col, ok := selector.(Column)
		if !ok {
			continue
		}
		cols = append(cols, col)
	}
	return cols
}

// SelectorNames returns a name array.
func (selectors Selectors) SelectorNames() []string {
	names := make([]string, 0)
	for _, selector := range selectors {
		names = append(names, selector.Name())
	}
	return names
}

// Len returns the length of the selector array.
func (selectors Selectors) Len() int {
	return len(selectors)
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

// FunctionExecutors returns a function executor array.
func (selectors Selectors) FunctionExecutors() ([]FunctionExecutor, error) {
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

// LookupFunctionExecutor returns a function executor with the specified name.
func (selectors Selectors) LookupFunctionExecutor(name string) (FunctionExecutor, error) {
	fn, err := selectors.LookupFunction(name)
	if err != nil {
		return nil, err
	}
	return fn.Executor()
}

// FunctionExecutorsForType returns a function executor array with the specified type.
func (selectors Selectors) FunctionExecutorsForType(t FunctionType) ([]FunctionExecutor, error) {
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
func (selectors Selectors) AggregateFunctions() ([]FunctionExecutor, error) {
	return selectors.FunctionExecutorsForType(AggregateFunctionType)
}

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
		fn, ok := selector.(Function)
		if !ok {
			continue
		}
		executor, err := GetFunctionExecutor(fn.Name())
		if err != nil {
			continue
		}
		if executor.Type() == t {
			return true
		}
	}
	return false
}

// HasAggregateFunction returns true if the selector list has an aggregate function.
func (selectors Selectors) HasAggregateFunction() bool {
	return selectors.HasFunctionWithType(AggregateFunctionType)
}

// SelectorString returns a string representation of the selector array.
func (selectors Selectors) SelectorString() string {
	strs := make([]string, len(selectors))
	for n, col := range selectors {
		sel, ok := col.(columnSelectorStringer)
		if ok {
			strs[n] = sel.SelectorString()
		} else {
			strs[n] = col.Name()
		}
	}
	return strings.JoinWithComma(strs)
}
