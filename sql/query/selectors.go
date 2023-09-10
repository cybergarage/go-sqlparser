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

// SelectorList represens a selector array.
type SelectorList []Selector

// NewSelectors returns a selector array instance.
func NewSelectors() SelectorList {
	return make(SelectorList, 0)
}

// NewSelectorsWith returns a selector array instance with the specified selectors.
func NewSelectorsWith(selectors ...Selector) SelectorList {
	s := make(SelectorList, len(selectors))
	copy(s, selectors)
	return s
}

// NewSelectorsWithColums returns a selector array instance with the specified selectors.
func NewSelectorsWithColums(selectors ...*Column) SelectorList {
	s := make(SelectorList, len(selectors))
	for n, selector := range selectors {
		s[n] = selector
	}
	return s
}

// Selector returns a selector array.
func (selectors SelectorList) Selectors() SelectorList {
	return selectors
}

// Functions returns a function array.
func (selectors SelectorList) Functions() []*Function {
	fns := make([]*Function, 0)
	for _, selector := range selectors {
		fn, ok := selector.(*Function)
		if !ok {
			continue
		}
		fns = append(fns, fn)
	}
	return fns
}

// IsSelectAll returns true if the selector list is "*".
func (selectors SelectorList) IsSelectAll() bool {
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
func (selectors SelectorList) HasFunction() bool {
	for _, selector := range selectors {
		_, ok := selector.(*Function)
		if ok {
			return true
		}
	}
	return false
}

// HasFunctionWithType returns true if the selector list has a function with the specified type.
func (selectors SelectorList) HasFunctionWithType(t FunctionType) bool {
	for _, selector := range selectors {
		fn, ok := selector.(*Function)
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
func (selectors SelectorList) HasAggregateFunction() bool {
	return selectors.HasFunctionWithType(AggregateFunctionType)
}

// SelectorString returns a string representation of the selector array.
func (selectors SelectorList) SelectorString() string {
	strs := make([]string, len(selectors))
	for n, col := range selectors {
		strs[n] = col.SelectorString()
	}
	return strings.JoinWithComma(strs)
}
