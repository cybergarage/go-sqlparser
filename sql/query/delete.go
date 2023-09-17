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
type DeleteOption = func(*Delete)

// Delete is a "DELETE" statement.
type Delete struct {
	table *Table
	*Condition
}

// NewDeleteWith returns a new Delete statement instance with the specified parameters.
func NewDeleteWith(tbl *Table, opts ...DeleteOption) *Delete {
	stmt := &Delete{
		table:     tbl,
		Condition: nil,
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

func WithDeleteCondition(cond *Condition) func(*Delete) {
	return func(stmt *Delete) {
		stmt.Condition = cond
	}
}

// StatementType returns the statement type.
func (stmt *Delete) StatementType() StatementType {
	return DeleteStatement
}

// Table returns the table.
func (stmt *Delete) Table() *Table {
	return stmt.table
}

// TableName returns the table name.
func (stmt *Delete) TableName() string {
	return stmt.table.TableName()
}

// From returns the table.
func (stmt *Delete) From() *Table {
	return stmt.table
}

// String returns the statement string representation.
func (stmt *Delete) String() string {
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
