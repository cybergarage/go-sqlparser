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

// SelectorString returns a string representation of the selector array.
func (selectors Selectors) SelectorString() string {
	strs := make([]string, len(selectors))
	for n, col := range selectors {
		switch t := col.(type) {
		case columnSelectorStringer:
			strs[n] = t.SelectorString()
		case Function:
			strs[n] = t.String()
		default:
			strs[n] = col.Name()
		}
	}
	return strings.JoinWithComma(strs)
}
