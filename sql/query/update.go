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

// Update represents a "UPDATE" statement interface.
type Update interface {
	Statement
	// Table returns the table.
	TableName() string
	// Columns returns the columns.
	Columns() ColumnList
	// Where returns the condition.
	Where() Condition
}

// updateStmt is a "UPDATE" statement.
type updateStmt struct {
	table Table
	ColumnList
	Condition
}

// NewUpdateWith returns a new updateStmt statement instance with the specified parameters.
func NewUpdateWith(tbl Table, columns ColumnList, w Condition) Update {
	return &updateStmt{
		table:      tbl,
		ColumnList: columns,
		Condition:  w,
	}
}

// StatementType returns the statement type.
func (stmt *updateStmt) StatementType() StatementType {
	return UpdateStatement
}

// Table returns the table.
func (stmt *updateStmt) Table() Table {
	return stmt.table
}

// TableName returns the table name.
func (stmt *updateStmt) TableName() string {
	return stmt.table.TableName()
}

// Where returns the condition.
func (stmt *updateStmt) Where() Condition {
	return stmt.Condition
}

// String returns the statement string representation.
func (stmt *updateStmt) String() string {
	strs := []string{
		"UPDATE",
		stmt.table.FullTableName(),
		"SET",
	}
	columns := []string{}
	for _, column := range stmt.ColumnList {
		name := column.Name()
		value := column.String()
		updator, ok := column.(columnUpdateStringer)
		if ok {
			value = updator.UpdatorString()
		}
		columSet := fmt.Sprintf("%s = %s", name, value)
		columns = append(columns, columSet)
	}
	strs = append(strs, strings.JoinWithComma(columns))
	if !stmt.Condition.IsEmpty() {
		strs = append(strs, "WHERE", stmt.Condition.String())
	}
	return strings.JoinWithSpace(strs)
}
