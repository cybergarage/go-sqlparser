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

// AlterDatabase represents a "ALTER DATABASE" statement interface.
type AlterDatabase interface {
	Statement
	// DatabaseName returns the database name.
	DatabaseName() string
	// RenameTo returns the "TO" database.
	RenameTo() Database
}

// alterDatabaseStmt is a "ALTER DATABASE" statement.
type alterDatabaseStmt struct {
	Database

	to Database
}

// NewAlterDatabaseWith returns a new alterDatabase statement instance with the specified options.
func NewAlterDatabaseWith(name string, to string) *alterDatabaseStmt {
	return &alterDatabaseStmt{
		Database: NewDatabaseWith(name),
		to:       NewDatabaseWith(to),
	}
}

// StatementType returns the statement type.
func (stmt *alterDatabaseStmt) StatementType() StatementType {
	return AlterDatabaseStatement
}

// RenameTo returns the "TO" database.
func (stmt *alterDatabaseStmt) RenameTo() Database {
	return stmt.to
}

// String returns the statement string representation.
func (stmt *alterDatabaseStmt) String() string {
	elems := []string{
		"ALTER",
		"DATABASE",
		stmt.DatabaseName(),
		"RENAME",
		"TO",
		stmt.to.DatabaseName(),
	}
	return strings.JoinWithSpace(elems)
}
