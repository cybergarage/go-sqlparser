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
	"strings"

	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-sqlparser/sql/query"
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
		if err := stmt.traverseSelectStatement(); err != nil {
			return nil, err
		}
	}
	return stmt, nil
}

func (stmt *schemaColumnsStatement) generateSelectStatement() error {
	query := SchemaColumnsQuery + " WHERE "
	if 0 < len(stmt.dbName) {
		query += SchemaColumnsSchema + " = '" + stmt.dbName + "'"
	}
	switch {
	case len(stmt.tblNames) == 1:
		query += " AND " + SchemaColumnsTableName + " = '" + stmt.tblNames[0] + "'"
	case 1 < len(stmt.tblNames):
		escapedTableNames := make([]string, len(stmt.tblNames))
		for n, tblName := range stmt.tblNames {
			escapedTableNames[n] = "'" + tblName + "'"
		}
		query += " AND " + SchemaColumnsTableName + " IN (" + strings.Join(escapedTableNames, ",") + ")"
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

func (stmt *schemaColumnsStatement) traverseSelectStatement() error {
	var traverseExpr func(expr query.Expr) error
	traverseExpr = func(expr query.Expr) error {
		switch e := expr.(type) {
		case *query.AndExpr:
			if err := traverseExpr(e.Left()); err != nil {
				return err
			}
			if err := traverseExpr(e.Right()); err != nil {
				return err
			}
			return nil
		case *query.CmpExpr:
			if e.Operator() != query.EQ {
				return newErrInvalidQuery(stmt.stmt.String())
			}
			var name string
			switch v := e.Left().Value().(type) {
			case string:
				name = v
			default:
				return newErrInvalidQuery(stmt.stmt.String())
			}
			var value string
			switch v := e.Right().Value().(type) {
			case string:
				value = v
			default:
				return newErrInvalidQuery(stmt.stmt.String())
			}
			switch {
			case strings.EqualFold(name, SchemaColumnsSchema):
				stmt.dbName = value
			case strings.EqualFold(name, SchemaColumnsTableName):
				stmt.tblNames = append(stmt.tblNames, value)
			}
			return nil
		}
		return newErrInvalidQuery(stmt.stmt.String())
	}

	stmt.tblNames = []string{}
	err := traverseExpr(stmt.stmt.Where().Expr())
	if err != nil {
		return err
	}

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
