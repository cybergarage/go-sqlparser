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

// Vacuum represents a "VACUUM" statement interface.
type Vacuum interface {
	Statement
	TableName() string
}

// vacuumStmt is a "VACUUM" statement.
type vacuumStmt struct {
	Table
}

// NewVacuum returns a new vacuum statement instance.
func NewVacuum() *vacuumStmt {
	return &vacuumStmt{
		Table: nil,
	}
}

// NewVacuumWith returns a new vacuum statement instance with the specified parameters.
func NewVacuumWith(tbl Table) *vacuumStmt {
	return &vacuumStmt{
		Table: tbl,
	}
}

// StatementType returns the statement type.
func (stmt *vacuumStmt) StatementType() StatementType {
	return VacuumStatement
}

// String returns the statement string representation.
func (stmt *vacuumStmt) String() string {
	strs := []string{
		"VACUUM",
	}
	if stmt.Table != nil {
		strs = append(strs, stmt.Table.FullTableName())
	}
	return strings.JoinWithSpace(strs)
}
