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
	std_strings "strings"

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

// Len returns the length of the column array.
func (columns ColumnList) Len() int {
	return len(columns)
}

// Selectors returns a selector array.
func (columns ColumnList) Selectors() SelectorList {
	return NewSelectorsWithColums(columns...)
}

// ColumnAt returns a column by the specified index.
func (columns ColumnList) ColumnAt(n int) (*Column, error) {
	if len(columns) <= n {
		return nil, newErrColumnIndexOutOfRange(n)
	}
	return columns[n], nil
}

// ColumnByName returns a column by the specified name.
func (columns ColumnList) ColumnByName(name string) (*Column, error) {
	for _, column := range columns {
		if std_strings.EqualFold(column.Name(), name) {
			return column, nil
		}
	}
	return nil, newErrColumnNotFound(name)
}

// ColumnIndexByName returns a column index by the specified name.
func (columns ColumnList) ColumnIndexByName(name string) (int, error) {
	for n, column := range columns {
		if column.Name() == name {
			return n, nil
		}
	}
	return -1, newErrColumnNotFound(name)
}

// Names returns a column name array.
func (columns ColumnList) Names() []string {
	names := make([]string, len(columns))
	for n, col := range columns {
		names[n] = col.Name()
	}
	return names
}

// NameString returns a string representation of the the column names.
func (columns ColumnList) NameString() string {
	strs := make([]string, len(columns))
	for n, col := range columns {
		strs[n] = col.Name()
	}
	return strings.JoinWithComma(strs)
}

// ValueString returns a string representation of the the column values.
func (columns ColumnList) ValueString() string {
	strs := make([]string, len(columns))
	for n, col := range columns {
		strs[n] = col.ValueString()
	}
	return strings.JoinWithComma(strs)
}

// SetSchema sets a schema to update column values.
func (columns ColumnList) SetSchema(schema *Schema) error {
	for _, col := range columns {
		schemaCol, err := schema.ColumnByName(col.Name())
		if err != nil {
			return err
		}
		err = col.SetDef(schemaCol.DataDef)
		if err != nil {
			return err
		}
	}
	return nil
}

// DefinitionString returns a string representation of the the column definitions.
func (columns ColumnList) DefinitionString() string {
	strs := make([]string, len(columns))
	for n, col := range columns {
		strs[n] = col.DefinitionString()
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

// Copy returns a copy of the column list.
func (columns ColumnList) Copy() ColumnList {
	cpColums := make(ColumnList, len(columns))
	for n, col := range columns {
		cpColums[n] = col.Copy()
	}
	return cpColums
}
