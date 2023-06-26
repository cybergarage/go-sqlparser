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
	// DDL (Data Definition Language)
	if stmt := ctx.Create_database_stmt(); stmt != nil {
		return newCreateDatabaseWith(stmt)
	}
	if stmt := ctx.Create_table_stmt(); stmt != nil {
		return newCreateTableWith(stmt)
	}
	if stmt := ctx.Create_index_stmt(); stmt != nil {
		return newCreateIndexWith(stmt)
	}
	// DMQL (Data Manipulation Language)
	if stmt := ctx.Insert_stmt(); stmt != nil {
		return newInsertWith(stmt)
	}
	if stmt := ctx.Update_stmt(); stmt != nil {
		return newUpdateWith(stmt)
	}
	if stmt := ctx.Select_stmt(); stmt != nil {
		return newSelectWith(stmt)
	}
	if stmt := ctx.Delete_stmt(); stmt != nil {
		return newDeleteWith(stmt)
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
	return query.NewCreateTableWith(newTableSchemaWith(ctx), ifNotExists)
}

func newCreateIndexWith(ctx antlr.ICreate_index_stmtContext) *query.CreateTable {
	ifNotExists := query.NewIfNotExistsWith(false)
	if ctx.If_not_exists() != nil {
		ifNotExists = query.NewIfNotExistsWith(true)
	}
	return query.NewCreateTableWith(newIndexSchemaWith(ctx), ifNotExists)
}

/*
func newDropDatabaseWith(ctx antlr.IDrop_database_stmtContext) *query.DropDatabase {
	dbName := ctx.Database_name().GetText()
	return query.NewDropDatabaseWith(dbName)
}

func newDropTableWith(ctx antlr.IDrop_table_stmtContext) *query.DropTable {
	return query.NewDropTableWith(newTableSchemaWith(ctx))
}

func newDropIndexWith(ctx antlr.IDrop_index_stmtContext) *query.DropTable {
	return query.NewDropTableWith(newIndexSchemaWith(ctx))
}

func newAlterDatabaseWith(ctx antlr.IAlter_database_stmtContext) *query.AlterDatabase {
	dbName := ctx.Database_name().GetText()
	return query.NewAlterDatabaseWith(dbName)
}

func newAlterTableWith(ctx antlr.IAlter_table_stmtContext) *query.AlterTable {
	return query.NewAlterTableWith(newTableSchemaWith(ctx))
}

func newAlterIndexWith(ctx antlr.IAlter_index_stmtContext) *query.AlterTable {
	return query.NewAlterTableWith(newIndexSchemaWith(ctx))
}
*/

func newTableSchemaWith(ctx antlr.ICreate_table_stmtContext) *query.Schema {
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

func newIndexSchemaWith(ctx antlr.ICreate_index_stmtContext) *query.Schema {
	tblName := ctx.Table_name().GetText()
	colums := query.NewColumns()
	indexes := query.NewIndexes()
	for _, columDef := range ctx.AllIndexed_column() {
		colum := newIndexedColumnWith(columDef)
		colums = append(colums, colum)
	}
	return query.NewSchemaWith(tblName, colums, indexes)
}

func newColumnWith(ctx antlr.IColumn_defContext) *query.Column {
	name := ctx.Column_name().GetText()
	t, err := query.NewDataTypeFrom(ctx.Type_name().GetText(), -1)
	if err != nil {
		t = &query.DataType{Type: query.UnknownData, Length: -1}
	}
	return query.NewColumnWith(name, t, nil)
}

func newIndexedColumnWith(ctx antlr.IIndexed_columnContext) *query.Column {
	name := ctx.Column_name().GetText()
	t, err := query.NewDataTypeFrom("", -1) // FIXME
	if err != nil {
		t = &query.DataType{Type: query.UnknownData, Length: -1}
	}
	return query.NewColumnWith(name, t, nil)
}

func newInsertWith(ctx antlr.IInsert_stmtContext) *query.Insert {
	return query.NewInsertWith()
}

func newUpdateWith(ctx antlr.IUpdate_stmtContext) *query.Update {
	return query.NewUpdateWith()
}

func newSelectWith(ctx antlr.ISelect_stmtContext) *query.Select {
	return query.NewSelectWith()
}

func newDeleteWith(ctx antlr.IDelete_stmtContext) *query.Delete {
	return query.NewDeleteWith()
}
