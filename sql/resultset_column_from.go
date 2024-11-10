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

package sql

import (
	"github.com/cybergarage/go-sqlparser/sql/query"
)

// NewResultSetColumnFrom returns a new resultset column from the specified column.
func NewResultSetColumnFrom(column query.Column) (ResultSetColumn, error) {
	return NewResultSetColumn(
		WithResultSetColumnName(column.Name()),
		WithResultSetColumnType(column.DataType()),
		WithResultSetColumnConstraint(column.Constraint()),
	), nil
}

// NewResultSetColumnsFrom returns a new resultset columns from the specified column list.
func NewResultSetColumnsFrom(column query.ColumnList) []ResultSetColumn {
	columns := make([]ResultSetColumn, len(column))
	for n, c := range column {
		columns[n] = c
	}
	return columns
}
