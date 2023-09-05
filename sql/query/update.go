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

// Update is a "UPDATE" statement.
type Update struct {
	table *Table
	ColumnList
	*Condition
}

// NewUpdateWith returns a new Update statement instance with the specified parameters.
func NewUpdateWith(tbl *Table, columns ColumnList, w *Condition) *Update {
	return &Update{
		table:      tbl,
		ColumnList: columns,
		Condition:  w,
	}
}

// StatementType returns the statement type.
func (stmt *Update) StatementType() StatementType {
	return UpdateStatement
}

// Table returns the table.
func (stmt *Update) Table() *Table {
	return stmt.table
}

// TableName returns the table name.
func (stmt *Update) TableName() string {
	return stmt.table.TableName()
}

// To returns the table.
func (stmt *Delete) To() *Table {
	return stmt.table
}

// String returns the statement string representation.
func (stmt *Update) String() string {
	strs := []string{
		"UPDATE",
		stmt.table.String(),
		"SET",
	}
	colums := []string{}
	for _, colum := range stmt.ColumnList {
		name := colum.Name()
		value := colum.ValueString()
		colum_set := fmt.Sprintf("%s = %s", name, value)
		colums = append(colums, colum_set)
	}
	strs = append(strs, strings.JoinWithComma(colums))
	if stmt.Condition != nil {
		strs = append(strs, "WHERE", stmt.Condition.String())
	}
	return strings.JoinWithSpace(strs)
}
