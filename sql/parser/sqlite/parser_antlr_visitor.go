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
	// DDL (Data Definition Language) - CREATE
	if stmt := ctx.Create_database_stmt(); stmt != nil {
		return newCreateDatabaseWith(stmt)
	}
	if stmt := ctx.Create_table_stmt(); stmt != nil {
		return newCreateTableWith(stmt)
	}
	if stmt := ctx.Create_index_stmt(); stmt != nil {
		return newCreateIndexWith(stmt)
	}
	// DDL (Data Definition Language) - DROP
	if stmt := ctx.Drop_database_stmt(); stmt != nil {
		return newDropDatabaseWith(stmt)
	}
	if stmt := ctx.Drop_table_stmt(); stmt != nil {
		return newDropTableWith(stmt)
	}
	if stmt := ctx.Drop_index_stmt(); stmt != nil {
		return newDropIndexWith(stmt)
	}
	// DDL (Data Definition Language) - ALTER
	if stmt := ctx.Alter_table_stmt(); stmt != nil {
		return newAlterTableWith(stmt)
	}
	// DML (Data Manipulation Language)
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

func newDropDatabaseWith(ctx antlr.IDrop_database_stmtContext) *query.DropDatabase {
	dbName := ctx.Database_name().GetText()
	ifExists := query.NewIfExistsWith(false)
	if ctx.If_exists() != nil {
		ifExists = query.NewIfExistsWith(true)
	}
	return query.NewDropDatabaseWith(dbName, ifExists)
}

func newDropTableWith(ctx antlr.IDrop_table_stmtContext) *query.DropTable {
	schemaName := ""
	if ctx.Schema_name() != nil {
		schemaName = ctx.Schema_name().GetText()
	}
	tblName := ctx.Table_name().GetText()
	ifExists := query.NewIfExistsWith(false)
	if ctx.If_exists() != nil {
		ifExists = query.NewIfExistsWith(true)
	}
	return query.NewDropTableWith(schemaName, tblName, ifExists)
}

func newDropIndexWith(ctx antlr.IDrop_index_stmtContext) *query.DropIndex {
	schemaName := ""
	if ctx.Schema_name() != nil {
		schemaName = ctx.Schema_name().GetText()
	}
	idxName := ctx.Index_name().GetText()
	ifExists := query.NewIfExistsWith(false)
	if ctx.If_exists() != nil {
		ifExists = query.NewIfExistsWith(true)
	}
	return query.NewDropIndexWith(schemaName, idxName, ifExists)
}

func newAlterTableWith(ctx antlr.IAlter_table_stmtContext) *query.AlterTable {
	schemaName := ""
	if ctx.Schema_name() != nil {
		schemaName = ctx.Schema_name().GetText()
	}
	tblName := "" // ctx.Table_name().GetText()
	return query.NewAlterTableWith(schemaName, tblName)
}

func newTableSchemaWith(ctx antlr.ICreate_table_stmtContext) *query.Schema {
	tblName := ctx.Table_name().GetText()
	colums := query.NewColumnList()
	indexes := query.NewIndexes()
	for _, columDef := range ctx.AllColumn_def() {
		colum := newColumnWith(columDef)
		colums = append(colums, colum)
		for _, columnConst := range columDef.AllColumn_constraint() {
			if isPrimary := columnConst.Primary_key_constraint(); isPrimary != nil {
				indexes = append(indexes, query.NewPrimaryIndexWith(query.NewColumnListWith(colum)))
			}
		}
	}
	for _, tblConst := range ctx.AllTable_constraint() {
		if primaryDef := tblConst.Primary_key_def(); primaryDef != nil {
			indexColums := query.NewColumnList()
			for _, columDef := range primaryDef.AllIndexed_column() {
				colum := newIndexedColumnWith(columDef)
				indexColums = append(indexColums, colum)
			}
			indexes = append(indexes, query.NewPrimaryIndexWith(indexColums))
		}
	}
	return query.NewSchemaWith(tblName, query.WithSchemaColumns(colums), query.WithSchemaIndexes(indexes))
}

func newIndexSchemaWith(ctx antlr.ICreate_index_stmtContext) *query.Schema {
	tblName := ctx.Table_name().GetText()
	colums := query.NewColumnList()
	indexes := query.NewIndexes()
	for _, columDef := range ctx.AllIndexed_column() {
		colum := newIndexedColumnWith(columDef)
		colums = append(colums, colum)
	}
	return query.NewSchemaWith(tblName, query.WithSchemaColumns(colums), query.WithSchemaIndexes(indexes))
}

func newColumnWith(ctx antlr.IColumn_defContext) *query.Column {
	name := ctx.Column_name().GetText()
	t, err := query.NewDataFrom(ctx.Type_name().GetText(), -1)
	if err != nil {
		t = &query.Data{Type: query.UnknownData, Length: -1}
	}
	return query.NewColumnWithOptions(query.WithColumnName(name), query.WithColumnData(t))
}

func newIndexedColumnWith(ctx antlr.IIndexed_columnContext) *query.Column {
	name := ctx.Column_name().GetText()
	t, err := query.NewDataFrom("", -1) // FIXME
	if err != nil {
		t = &query.Data{Type: query.UnknownData, Length: -1}
	}
	return query.NewColumnWithOptions(query.WithColumnName(name), query.WithColumnData(t))
}

func newInsertWith(ctx antlr.IInsert_stmtContext) *query.Insert {
	tbl := query.NewTableWith(ctx.Table_name().GetText())
	names := []string{}
	for _, name := range ctx.AllColumn_name() {
		names = append(names, name.GetText())
	}
	values := []any{}
	for _, row := range ctx.Values_clause().AllValue_row() {
		for _, expr := range row.AllExpr() {
			if v := expr.Literal_value(); v != nil {
				values = append(values, newLiteralValueWith(v))
			} else if v := expr.Bind_param(); v != nil {
				values = append(values, newBindParamWith(v))
			}
		}
	}
	colums := query.NewColumnList()
	for n, name := range names {
		var v any
		if n < len(values) {
			v = values[n]
		}
		colums = append(colums, query.NewColumnWithOptions(query.WithColumnName(name), query.WithColumnLiteral(query.NewLiteralWith(v))))
	}
	return query.NewInsertWith(tbl, colums)
}

func newUpdateWith(ctx antlr.IUpdate_stmtContext) *query.Update {
	tbl := query.NewTableWith(ctx.GetTable().GetText())
	cols := query.NewColumnList()
	for _, set := range ctx.AllUpdate_column_set() {
		name := set.Column_name().GetText()
		opts := []query.ColumnOption{
			query.WithColumnName(name),
		}
		if v := set.Expr().Literal_value(); v != nil {
			opts = append(opts, query.WithColumnLiteral(newLiteralValueWith(v)))
		} else if v := set.Expr().Bind_param(); v != nil {
			opts = append(opts, query.WithColumnLiteral(newBindParamWith(v)))
		}
		cols = append(cols, query.NewColumnWithOptions(opts...))
	}
	var where *query.Where
	if w := ctx.GetWhereExpr(); w != nil {
		where = query.NewWhereWith(newExprWith(w))
	}
	return query.NewUpdateWith(tbl, cols, where)
}

func newSelectWith(ctx antlr.ISelect_stmtContext) *query.Select {
	cols := query.NewColumnList()
	tbls := query.NewTables()
	var topExpr query.Expr
	if parentQuery := ctx.GetParentQuery(); parentQuery != nil {
		for _, col := range parentQuery.AllResult_column() {
			cols = append(cols, query.NewColumnWithOptions(query.WithColumnName(col.GetText())))
		}
		for _, from := range ctx.GetParentQuery().AllFrom() {
			if tbl := from.From_table(); tbl != nil {
				tbls = append(tbls, query.NewTableWith(tbl.GetText()))
			}
		}
		if w := parentQuery.GetWhereExpr(); w != nil {
			topExpr = newExprWith(w)
		}
	}
	var where *query.Where
	if topExpr != nil {
		where = query.NewWhereWith(topExpr)
	}
	return query.NewSelectWith(cols, tbls, where)
}

func newDeleteWith(ctx antlr.IDelete_stmtContext) *query.Delete {
	tbl := query.NewTableWith(ctx.GetTable().GetText())
	var where *query.Where
	if w := ctx.GetWhereExpr(); w != nil {
		where = query.NewWhereWith(newExprWith(w))
	}
	return query.NewDeleteWith(tbl, where)
}

func newLiteralValueWith(ctx antlr.ILiteral_valueContext) *query.Literal {
	if ctx == nil {
		return nil
	}
	return query.NewLiteralWith(ctx.GetText())
}

func newBindParamWith(ctx antlr.IBind_paramContext) *query.Literal {
	if ctx == nil {
		return nil
	}
	return query.NewLiteralWith(query.NewBindParamWith(ctx.GetText()))
}

func newExprWith(ctx antlr.IExprContext) query.Expr {
	if cmpExpr := ctx.Comparison_expr(); cmpExpr != nil {
		c := query.NewColumnWithName(cmpExpr.Column_name().GetText())
		l := newLiteralValueWith(cmpExpr.Literal_value())
		if cmpExpr.ASSIGN() != nil {
			return query.NewCompExpr(query.EqualOp, c, l)
		}
		if cmpExpr.NOT_EQ1() != nil || cmpExpr.NOT_EQ2() != nil {
			return query.NewCompExpr(query.NotEqualOp, c, l)
		}
		if cmpExpr.LT() != nil {
			return query.NewCompExpr(query.LessThanOp, c, l)
		}
		if cmpExpr.GT() != nil {
			return query.NewCompExpr(query.GreaterThanOp, c, l)
		}
		if cmpExpr.LT_EQ() != nil {
			return query.NewCompExpr(query.LessEqualOp, c, l)
		}
		if cmpExpr.GT_EQ() != nil {
			return query.NewCompExpr(query.GreaterEqualOp, c, l)
		}
	}
	return nil
}
