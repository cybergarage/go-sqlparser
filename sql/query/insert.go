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

// Insert represents a "INSERT" statement interface.
type Insert interface {
	Statement
	// TableName returns the table name.
	TableName() string
	// Columns returns the columns.
	Columns() Columns
}

// insertStmt is a "INSERT" statement.
type insertStmt struct {
	Table
	columns Columns
}

// NewInsertWith returns a new insert statement instance with the specified parameters.
func NewInsertWith(tbl Table, columns Columns) Insert {
	return &insertStmt{
		Table:   tbl,
		columns: columns,
	}
}

// StatementType returns the statement type.
func (stmt *insertStmt) StatementType() StatementType {
	return InsertStatement
}

// Columns returns the columns.
func (stmt *insertStmt) Columns() Columns {
	return stmt.columns
}

// String returns the statement string representation.
func (stmt *insertStmt) String() string {
	strs := []string{
		"INSERT",
		"INTO",
		stmt.Table.FullTableName(),
		"(" + stmt.Columns().NameString() + ")",
		"VALUES",
		"(" + stmt.Columns().ValueString() + ")",
	}
	return strings.JoinWithSpace(strs)
}
