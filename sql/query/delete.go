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

// DeleteOption represents a delete option function.
type DeleteOption = func(*deleteStmt)

// Delete represents a "DELETE" statement interface.
type Delete interface {
	Statement
	TableName() string
	Where() *Condition
}

// deleteStmt is a "DELETE" statement.
type deleteStmt struct {
	table *Table
	*Condition
}

// NewDeleteWith returns a new deleteStmt statement instance with the specified parameters.
func NewDeleteWith(tbl *Table, opts ...DeleteOption) *deleteStmt {
	stmt := &deleteStmt{
		table:     tbl,
		Condition: nil,
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

func WithDeleteCondition(cond *Condition) func(*deleteStmt) {
	return func(stmt *deleteStmt) {
		stmt.Condition = cond
	}
}

// StatementType returns the statement type.
func (stmt *deleteStmt) StatementType() StatementType {
	return DeleteStatement
}

// Table returns the table.
func (stmt *deleteStmt) Table() *Table {
	return stmt.table
}

// TableName returns the table name.
func (stmt *deleteStmt) TableName() string {
	return stmt.table.TableName()
}

// From returns the table.
func (stmt *deleteStmt) From() *Table {
	return stmt.table
}

// String returns the statement string representation.
func (stmt *deleteStmt) String() string {
	strs := []string{
		"DELETE",
		"FROM",
		stmt.table.String(),
	}
	if stmt.Condition != nil {
		strs = append(strs, "WHERE", stmt.Condition.String())
	}
	return strings.JoinWithSpace(strs)
}
