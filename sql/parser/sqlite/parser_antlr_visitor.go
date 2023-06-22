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

package sqlite

import (
	"github.com/cybergarage/go-sqlparser/sql/parser/sqlite/antlr"
	"github.com/cybergarage/go-sqlparser/sql/query"
)

type antlrVisitor struct {
	*antlr.BaseSQLiteParserVisitor
}

func newANTLRVisitor() *antlrVisitor {
	v := &antlrVisitor{
		BaseSQLiteParserVisitor: &antlr.BaseSQLiteParserVisitor{},
	}
	return v
}

func (v *antlrVisitor) VisitParse(ctx *antlr.ParseContext) interface{} {
	return newStatementListWith(ctx.AllSql_stmt_list())
}

func newStatementListWith(ctxList []antlr.ISql_stmt_listContext) query.StatementList {
	stmtList := query.NewStatementList()
	for _, ctx := range ctxList {
		for _, sqlStmt := range ctx.AllSql_stmt() {
			stmt := newStatementWith(sqlStmt)
			if stmt == nil {
				continue
			}
			stmtList = append(stmtList, stmt)
		}
	}
	return stmtList
}

func newStatementWith(ctx antlr.ISql_stmtContext) query.Statement {
	if stmt := ctx.Create_database_stmt(); stmt != nil {
		return newCreateDatabaseWith(stmt)
	}
	if stmt := ctx.Create_table_stmt(); stmt != nil {
		return newCreateTableWith(stmt)
	}
	return nil
}

func newCreateDatabaseWith(ctx antlr.ICreate_database_stmtContext) *query.CreateDatabase {
	dbName := ctx.Database_name().GetText()
	ifNotExists := query.NewIfNotExistsWith(false)
	if ctx.If_not_exists() != nil {
		ifNotExists = query.NewIfNotExistsWith(true)
	}
	return query.NewCreateDatabaseWith(dbName, ifNotExists)
}

func newCreateTableWith(ctx antlr.ICreate_table_stmtContext) *query.CreateTable {
	ifNotExists := query.NewIfNotExistsWith(false)
	if ctx.If_not_exists() != nil {
		ifNotExists = query.NewIfNotExistsWith(true)
	}
	return query.NewCreateTableWith(newSchemaWith(ctx), ifNotExists)
}

func newSchemaWith(ctx antlr.ICreate_table_stmtContext) *query.Schema {
	tblName := ctx.Table_name().GetText()
	colums := query.NewColumns()
	indexes := query.NewIndexes()
	for _, columDef := range ctx.AllColumn_def() {
		colum := newColumnWith(columDef)
		colums = append(colums, colum)
		for _, columnConst := range columDef.AllColumn_constraint() {
			if isPrimary := columnConst.PRIMARY_(); isPrimary != nil {
				indexes = append(indexes, query.NewIndexWith("", query.PrimaryIndex, query.NewColumnsWith(colum)))
			}
		}
	}
	return query.NewSchemaWith(tblName, colums, indexes)
}

func newColumnWith(ctx antlr.IColumn_defContext) *query.Column {
	name := ctx.Column_name().GetText()
	t, err := query.NewDataTypeFrom(ctx.Type_name().GetText())
	if err != nil {
		t = query.UnknownData
	}
	return query.NewColumnWith(name, t, nil)
}
