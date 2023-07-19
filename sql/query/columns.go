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

	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// ColumnList represens a column array.
type ColumnList []*Column

// NewColumns returns a column array instance.
func NewColumns() ColumnList {
	return make(ColumnList, 0)
}

// NewColumnsWith returns a column array instance with the specified columns.
func NewColumnsWith(columns ...*Column) ColumnList {
	c := make(ColumnList, len(columns))
	copy(c, columns)
	return c
}

// Column returns a column array.
func (columns ColumnList) Columns() ColumnList {
	return columns
}

// ColumnByName returns a column by the specified name.
func (columns ColumnList) ColumnByName(name string) (*Column, error) {
	for _, colum := range columns {
		if colum.Name() == name {
			return colum, nil
		}
	}
	return nil, newErrColumnNotFound(name)
}

// Names returns a column name array.
func (columns ColumnList) Names() []string {
	names := make([]string, len(columns))
	for n, col := range columns {
		names[n] = col.Name()
	}
	return names
}

// NameString returns a string representation of the the colum names.
func (columns ColumnList) NameString() string {
	strs := make([]string, len(columns))
	for n, col := range columns {
		strs[n] = col.Name()
	}
	return strings.JoinWithComma(strs)
}

// ValueString returns a string representation of the the colum values.
func (columns ColumnList) ValueString() string {
	strs := make([]string, len(columns))
	for n, col := range columns {
		strs[n] = fmt.Sprintf("%v", col.Value())
	}
	return strings.JoinWithComma(strs)
}

// DefString returns a string representation of the the colum definitions.
func (columns ColumnList) DefString() string {
	strs := make([]string, len(columns))
	for n, col := range columns {
		strs[n] = col.DefString()
	}
	return strings.JoinWithComma(strs)
}

// IsSelectAll returns true if the column list is "*".
func (columns ColumnList) IsSelectAll() bool {
	l := len(columns)
	switch {
	case l == 1:
		return columns[0].Name() == Asterisk
	case l == 0:
		return true
	}
	return false
}
