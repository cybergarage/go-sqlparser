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

// CreateDatabase represents a "CREATE DATABASE" statement interface.
type CreateDatabase interface {
	Statement
	DatabaseName() string
	IfNotExists() bool
}

// createDatabaseStmt is a "CREATE DATABASE" statement.
type createDatabaseStmt struct {
	Database
	*IfNotExistsOpt
}

// NewCreateDatabaseWith returns a new createDatabase statement instance with the specified options.
func NewCreateDatabaseWith(name string, ifne *IfNotExistsOpt) CreateDatabase {
	return &createDatabaseStmt{
		Database:       NewDatabaseWith(name),
		IfNotExistsOpt: ifne,
	}
}

// StatementType returns the statement type.
func (stmt *createDatabaseStmt) StatementType() StatementType {
	return CreateDatabaseStatement
}

// String returns the statement string representation.
func (stmt *createDatabaseStmt) String() string {
	elems := []string{
		"CREATE DATABASE"}
	if stmt.IfNotExists() {
		elems = append(elems, stmt.IfNotExistsOpt.String())
	}
	elems = append(elems, stmt.DatabaseName())
	return strings.JoinWithSpace(elems)
}
