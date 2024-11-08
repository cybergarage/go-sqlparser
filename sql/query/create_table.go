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

// CreateTable represents a "CREATE TABLE" statement interface.
type CreateTable interface {
	Statement
	TableName() string
	Schema() Schema
	IfNotExists() bool
}

// createTableStmt is a "CREATE TABLE" statement.
type createTableStmt struct {
	schema Schema
	*IfNotExistsOpt
}

// NewCreateTableWith returns a new createTable statement instance with the specified options.
func NewCreateTableWith(schema Schema, ifne *IfNotExistsOpt) CreateTable {
	return &createTableStmt{
		schema:         schema,
		IfNotExistsOpt: ifne,
	}
}

// Schema returns the schema.
func (stmt *createTableStmt) Schema() Schema {
	return stmt.schema
}

// TableName returns the table name.
func (stmt *createTableStmt) TableName() string {
	return stmt.schema.TableName()
}

// StatementType returns the statement type.
func (stmt *createTableStmt) StatementType() StatementType {
	return CreateTableStatement
}

// String returns the statement string representation.
func (stmt *createTableStmt) String() string {
	columnsStr := "("
	columnsStr += stmt.schema.Columns().DefinitionString()
	if 0 < len(stmt.schema.Indexes()) {
		columnsStr += ", "
		columnsStr += stmt.schema.Indexes().DefinitionString()
	}
	columnsStr += ")"

	elems := []string{
		"CREATE TABLE",
		stmt.IfNotExistsOpt.String(),
		stmt.TableName(),
		columnsStr,
	}
	return strings.JoinWithSpace(elems)
}
