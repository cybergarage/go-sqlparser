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

// NewColumnList returns a column array instance.
func NewColumnList() ColumnList {
	return make(ColumnList, 0)
}

// NewColumnListWith returns a column array instance with the specified columns.
func NewColumnListWith(columns ...*Column) ColumnList {
	c := make(ColumnList, len(columns))
	copy(c, columns)
	return c
}

// Column returns a column array.
func (colums ColumnList) Columns() ColumnList {
	return colums
}

// Names returns a column name array.
func (colums ColumnList) Names() []string {
	names := make([]string, len(colums))
	for n, col := range colums {
		names[n] = col.Name()
	}
	return names
}

// NameString returns a string representation of the the colum names.
func (colums ColumnList) NameString() string {
	strs := make([]string, len(colums))
	for n, col := range colums {
		strs[n] = col.Name()
	}
	return strings.JoinWithComma(strs)
}

// ValueString returns a string representation of the the colum values.
func (colums ColumnList) ValueString() string {
	strs := make([]string, len(colums))
	for n, col := range colums {
		strs[n] = fmt.Sprintf("%v", col.Value())
	}
	return strings.JoinWithComma(strs)
}

// DefString returns a string representation of the the colum definitions.
func (colums ColumnList) DefString() string {
	strs := make([]string, len(colums))
	for n, col := range colums {
		strs[n] = col.DefString()
	}
	return strings.JoinWithComma(strs)
}
