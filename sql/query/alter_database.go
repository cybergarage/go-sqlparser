// Copyright (C) 2019 Satoshi Konno. All rights reserved.
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

// AlterDatabase is a "ALTER DATABASE" statement.
type AlterDatabase struct {
	*Database
	to *Database
}

// NewAlterDatabaseWith returns a new AlterDatabase statement instance with the specified options.
func NewAlterDatabaseWith(name string, to string) *AlterDatabase {
	return &AlterDatabase{
		Database: NewDatabaseWith(name),
		to:       NewDatabaseWith(to),
	}
}

// StatementType returns the statement type.
func (stmt *AlterDatabase) StatementType() StatementType {
	return AlterDatabaseStatement
}

// RenameTo returns the "TO" database.
func (stmt *AlterDatabase) RenameTo() *Database {
	return stmt.to
}

// String returns the statement string representation.
func (stmt *AlterDatabase) String() string {
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
