// Copyright (C) 2025 The go-sqlparser Authors. All rights reserved.
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

package system

import (
	"github.com/cybergarage/go-sqlparser/sql"
)

type schemaColumnsStatement struct {
	stmt     sql.Select
	dbName   string
	tblNames []string
}

// SchemaColumnsStatement represents a schema columns statement.
type SchemaColumnsStatementOption func(*schemaColumnsStatement)

// With SchemaColumnsStatementDatabaseName sets the database name.
func WithSchemaColumnsStatementDatabaseName(db string) SchemaColumnsStatementOption {
	return func(stmt *schemaColumnsStatement) {
		stmt.dbName = db
	}
}

// WithSchemaColumnsStatementTableName sets the table name.
func WithSchemaColumnsStatementTableNames(tables []string) SchemaColumnsStatementOption {
	return func(stmt *schemaColumnsStatement) {
		stmt.tblNames = tables
	}
}

// WithSchemaColumnsStatementSelect sets the select statement.
func WithSchemaColumnsStatementSelect(sel sql.Select) SchemaColumnsStatementOption {
	return func(stmt *schemaColumnsStatement) {
		stmt.stmt = sel
	}
}

// NewSchemaColumnsStatement returns a new SchemaColumnsStatement.
func NewSchemaColumnsStatement(opts ...SchemaColumnsStatementOption) (SchemaColumnsStatement, error) {
	stmt := &schemaColumnsStatement{
		stmt:     nil,
		dbName:   "",
		tblNames: []string{},
	}
	for _, opt := range opts {
		opt(stmt)
	}
	if stmt.stmt == nil {
		if err := stmt.generateSelectStatement(); err != nil {
			return nil, err
		}
	} else {
		if err := stmt.parseSelectStatement(); err != nil {
			return nil, err
		}
	}
	return stmt, nil
}

func (stmt *schemaColumnsStatement) generateSelectStatement() error {
	query := SchemaColumnsQuery
	if 0 < len(stmt.tblNames) {
		query += " WHERE "
		for n, tblName := range stmt.tblNames {
			if 0 < n {
				query += " OR "
			}
			query += "TABLE_NAME = '" + tblName + "'"
		}
	}
	parsedStmts, err := sql.NewParser().ParseString(query)
	if err != nil {
		return err
	}
	if len(parsedStmts) != 1 {
		return newErrInvalidQuery(query)
	}
	selectStmt, ok := parsedStmts[0].(sql.Select)
	if !ok {
		return newErrInvalidQuery(query)
	}
	stmt.stmt = selectStmt
	return nil
}

func (stmt *schemaColumnsStatement) parseSelectStatement() error {
	return nil
}

// Statement returns the statement.
func (stmt *schemaColumnsStatement) Statement() sql.Select {
	return stmt.stmt
}

// DatabaseName returns the database name.
func (stmt *schemaColumnsStatement) DatabaseName() string {
	return stmt.dbName
}

// TableNames returns the table names.
func (stmt *schemaColumnsStatement) TableNames() []string {
	return stmt.tblNames
}
