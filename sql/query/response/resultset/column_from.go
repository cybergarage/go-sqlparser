// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
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

package resultset

import (
	"github.com/cybergarage/go-sqlparser/sql/query"
)

// NewColumnFrom returns a new resultset column from the specified column.
func NewColumnFrom(column query.Column) (Column, error) {
	return NewColumn(
		WithColumnName(column.Name()),
		WithColumnType(column.DataType()),
		WithColumnConstraint(column.Constraint()),
	), nil
}

// NewColumnsFrom returns a new resultset columns from the specified column list.
func NewColumnsFrom(column query.Columns) []Column {
	columns := make([]Column, len(column))
	for n, c := range column {
		columns[n] = c
	}
	return columns
}
