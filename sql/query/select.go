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

import "github.com/cybergarage/go-sqlparser/sql/util/strings"

// Select is a "SELECT" statement.
type Select struct {
	Tables
	Columns
	*Where
}

// NewSelectWith returns a new Select statement instance with the specified parameters.
func NewSelectWith(colums Columns, tbls Tables, wehre *Where) *Select {
	return &Select{
		Columns: colums,
		Tables:  tbls,
		Where:   wehre,
	}
}

// StatementType returns the statement type.
func (stmt *Select) StatementType() StatementType {
	return SelectStatement
}

// String returns the statement string representation.
func (stmt *Select) String() string {
	columnsStr := "*"
	if 0 < len(stmt.Columns) {
		columnsStr = stmt.Columns.NameString()
	}
	strs := []string{
		"SELECT",
		columnsStr,
		"FROM",
		stmt.Tables.String(),
		"WHERE",
		stmt.Where.String(),
	}
	return strings.JoinWithSpace(strs)
}
