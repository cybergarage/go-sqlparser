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
	"strconv"

	"github.com/cybergarage/go-sqlparser/sql/parser/sqlite/antlr"
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
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
	if stmt := ctx.Alter_database_stmt(); stmt != nil {
		return newAlterDatabaseWith(stmt)
	}
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
	if stmt := ctx.Commit_stmt(); stmt != nil {
		return query.NewCommit()
	}
	// Transaction statements
	if stmt := ctx.Begin_stmt(); stmt != nil {
		return newBeginWith(stmt)
	}
	if stmt := ctx.Rollback_stmt(); stmt != nil {
		return newRollbackWith(stmt)
	}
	if stmt := ctx.Commit_stmt(); stmt != nil {
		return newCommitWith(stmt)
	}
	// Extra statements
	if stmt := ctx.Copy_stmt(); stmt != nil {
		return newCopyWith(stmt)
	}
	if stmt := ctx.Truncate_stmt(); stmt != nil {
		return newTruncateWith(stmt)
	}
	if stmt := ctx.Vacuum_stmt(); stmt != nil {
		return newVacuumWith(stmt)
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
	ifExists := query.NewIfExistsWith(false)
	if ctx.If_exists() != nil {
		ifExists = query.NewIfExistsWith(true)
	}
	tbls := query.NewTables()
	for _, tbl := range ctx.AllTable_name() {
		tbls = append(tbls, query.NewTableWith(tbl.GetText()))
	}
	return query.NewDropTableWith(tbls, ifExists)
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

func newAlterDatabaseWith(ctx antlr.IAlter_database_stmtContext) *query.AlterDatabase {
	dbName := ctx.Database_name().GetText()
	toDbName := ""
	if ctx := ctx.Rename_database_to(); ctx != nil {
		toDbName = ctx.GetNew_database_name().GetText()
	}
	return query.NewAlterDatabaseWith(dbName, toDbName)
}

func newAlterTableWith(ctx antlr.IAlter_table_stmtContext) *query.AlterTable {
	opts := []query.AlterTableOption{}
	if ctx := ctx.Schema_name(); ctx != nil {
		schemaName := ctx.GetText()
		opts = append(opts, query.WithAlterTableSchema(schemaName))
	}
	if ctx := ctx.Rename_table_to(); ctx != nil {
		tblName := ctx.Table_name().GetText()
		opts = append(opts, query.WithAlterTableRenameTo(query.NewTableWith(tblName)))
	}
	if ctx := ctx.Rename_table_colum(); ctx != nil {
		fromColumn := query.NewColumnWithName(ctx.GetOld_column_name().GetText())
		toColumn := query.NewColumnWithName(ctx.GetNew_column_name().GetText())
		opts = append(opts, query.WithAlterTableRenameColumn(fromColumn, toColumn))
	}
	if ctx := ctx.Add_column(); ctx != nil {
		column := newColumnWith(ctx.Column_def())
		if ctx := ctx.Column_constraint(); ctx != nil {
			if ctx.Primary_key_constraint() != nil {
				column.SetConstant(query.ColumnConstraintPrimaryKey)
			}
		}
		opts = append(opts, query.WithAlterTableAddColumn(column))
	}
	if ctx := ctx.Drop_column(); ctx != nil {
		column := query.NewColumnWithOptions(query.WithColumnName(ctx.GetText()))
		opts = append(opts, query.WithAlterTableDropColumn(column))
	}
	tblName := ctx.GetTarget_table_name().GetText()
	return query.NewAlterTableWith(tblName, opts...)
}

func newTableSchemaWith(ctx antlr.ICreate_table_stmtContext) *query.Schema {
	tblName := ctx.Table_name().GetText()
	columns := query.NewColumns()
	indexes := query.NewIndexes()
	for _, columDef := range ctx.AllColumn_def() {
		column := newColumnWith(columDef)
		columns = append(columns, column)
		for _, columnConst := range columDef.AllColumn_constraint() {
			if isPrimary := columnConst.Primary_key_constraint(); isPrimary != nil {
				indexes = append(indexes, query.NewPrimaryIndexWith(query.NewColumnsWith(column)))
			}
		}
	}
	for _, tblConst := range ctx.AllTable_constraint() {
		if primaryDef := tblConst.Primary_key_def(); primaryDef != nil {
			indexColums := query.NewColumns()
			for _, columDef := range primaryDef.AllIndexed_column() {
				column := newIndexedColumnWith(columDef)
				indexColums = append(indexColums, column)
			}
			indexes = append(indexes, query.NewPrimaryIndexWith(indexColums))
		}
	}
	return query.NewSchemaWith(tblName, query.WithSchemaColumns(columns), query.WithSchemaIndexes(indexes))
}

func newIndexSchemaWith(ctx antlr.ICreate_index_stmtContext) *query.Schema {
	tblName := ctx.Table_name().GetText()
	columns := query.NewColumns()
	indexes := query.NewIndexes()
	for _, columDef := range ctx.AllIndexed_column() {
		column := newIndexedColumnWith(columDef)
		columns = append(columns, column)
	}
	return query.NewSchemaWith(tblName, query.WithSchemaColumns(columns), query.WithSchemaIndexes(indexes))
}

func newColumnWith(ctx antlr.IColumn_defContext) *query.Column {
	opts := []query.ColumnOption{
		query.WithColumnName(ctx.Column_name().GetText()),
	}
	if typ := ctx.Type_name(); typ != nil {
		t, err := query.NewDataFrom(strings.SplitDataTypeString(typ.GetText()))
		if err != nil {
			t = &query.DataDef{Type: query.UnknownData, Length: -1}
		}
		opts = append(opts, query.WithColumnData(t))
	}
	return query.NewColumnWithOptions(opts...)
}

func newIndexedColumnWith(ctx antlr.IIndexed_columnContext) *query.Column {
	name := ctx.Column_name().GetText()
	t, err := query.NewDataFrom("", -1) // FIXME
	if err != nil {
		t = &query.DataDef{Type: query.UnknownData, Length: -1}
	}
	return query.NewColumnWithOptions(query.WithColumnName(name), query.WithColumnData(t))
}

func newInsertWith(ctx antlr.IInsert_stmtContext) *query.Insert {
	tbl := query.NewTableWith(ctx.Table_name().GetText())
	columns := query.NewColumns()
	names := []string{}
	for _, name := range ctx.AllColumn_name() {
		names = append(names, name.GetText())
	}
	valueIdx := 0
	for _, row := range ctx.Values_clause().AllValue_row() {
		for _, expr := range row.AllExpr() {
			if v := expr.Comparison_expr(); v != nil {
				columns = append(columns,
					query.NewColumnWithOptions(query.WithColumnName(v.Column_name().GetText()),
						query.WithColumnLiteral(newLiteralValueWith(v.Literal_value()))))
			} else if v := expr.Literal_value(); v != nil {
				if valueIdx < len(names) {
					columns = append(columns,
						query.NewColumnWithOptions(query.WithColumnName(names[valueIdx]),
							query.WithColumnLiteral(newLiteralValueWith(v))))
					valueIdx++
				}
			} else if v := expr.Bind_param(); v != nil {
				if valueIdx < len(names) {
					columns = append(columns,
						query.NewColumnWithOptions(query.WithColumnName(names[valueIdx]),
							query.WithColumnLiteral(newBindParamWith(v))))
					valueIdx++
				}
			}
		}
	}
	return query.NewInsertWith(tbl, columns)
}

func newUpdateWith(ctx antlr.IUpdate_stmtContext) *query.Update {
	tbl := query.NewTableWith(ctx.GetTable().GetText())
	columns := query.NewColumns()
	for _, set := range ctx.AllUpdate_column_set() {
		name := set.Column_name().GetText()
		opts := []query.ColumnOption{
			query.WithColumnName(name),
		}
		if v := set.Expr().Literal_value(); v != nil {
			opts = append(opts, query.WithColumnLiteral(newLiteralValueWith(v)))
		} else if v := set.Expr().Bind_param(); v != nil {
			opts = append(opts, query.WithColumnLiteral(newBindParamWith(v)))
		} else if v := set.Expr().Arithmetic_expr(); v != nil {
			var executor query.FunctionExecutor
			ope := v.GetOpe().GetText()
			switch ope {
			case "+":
				executor = query.NewAddFunction(ope)
			case "-":
				executor = query.NewSubFunction(ope)
			case "*":
				executor = query.NewMulFunction(ope)
			case "/":
				executor = query.NewDivFunction(ope)
			case "%":
				executor = query.NewModFunction(ope)
			}
			if executor != nil {
				opts = append(opts, query.WithColumnFunction(executor))
				args := []any{
					v.Column_name().GetText(),
					v.Expr().GetText(),
				}
				opts = append(opts, query.WithColumnArguments(args))
			}
		}
		columns = append(columns, query.NewColumnWithOptions(opts...))
	}
	var where *query.Condition
	if w := ctx.GetWhereExpr(); w != nil {
		where = query.NewConditionWith(newExprWith(w))
	}
	return query.NewUpdateWith(tbl, columns, where)
}

func newSelectWith(ctx antlr.ISelect_stmtContext) *query.Select {
	sels := query.NewSelectors()
	tbls := query.NewTables()
	opts := []query.SelectOption{}
	var topExpr query.Expr
	if parentQuery := ctx.GetParentQuery(); parentQuery != nil {
		for _, col := range parentQuery.AllResult_column() {
			sel, err := newSelectorFrom(col)
			if err != nil {
				continue
			}
			sels = append(sels, sel)
		}
		for _, from := range ctx.GetParentQuery().AllFrom() {
			if tbl := from.From_table(); tbl != nil {
				tbls = append(tbls, query.NewTableWith(tbl.GetText()))
			}
		}
		if w := parentQuery.GetWhereExpr(); w != nil {
			topExpr = newExprWith(w)
		}
		groupBy := parentQuery.GetGroupByExpr()
		if 0 < len(groupBy) {
			opts = append(opts, query.WithSelectGroupBy(groupBy[0].GetText()))
		}
	}
	var where *query.Condition
	if topExpr != nil {
		where = query.NewConditionWith(topExpr)
	}

	if ctx := ctx.Order_by_stmt(); ctx != nil {
		orderBy := ctx.AllOrdering_term()
		if 0 < len(orderBy) {
			orders := []*query.Order{}
			for _, order := range orderBy {
				orderColumn := order.Expr().GetText()
				orderType := query.OrderNone
				if orderSpec := order.Asc_desc(); orderSpec != nil {
					orderType = query.NewOrderTypeWith(orderSpec.GetText())
				}
				if orderType.IsNone() {
					continue
				}
				orders = append(orders, query.NewOrderWith(orderColumn, orderType))
			}
			if 0 < len(orders) {
				opts = append(opts, query.WithSelectOrderBy(query.NewOrderByWith(orders)))
			}
		}
	}

	if ctx := ctx.Limit_stmt(); ctx != nil {
		limits := ctx.AllExpr()
		if 0 < len(limits) {
			limit, err := strconv.Atoi(limits[0].GetText())
			if err != nil {
				limit = 0
			}
			offset := 0
			if 1 < len(limits) {
				offset, err = strconv.Atoi(limits[1].GetText())
				if err != nil {
					offset = 0
				}
			}
			opts = append(opts, query.WithSelectLimit(offset, limit))
		}
	}

	return query.NewSelectWith(sels, tbls, where, opts...)
}

func newSelectorFrom(ctx antlr.IResult_columnContext) (query.Selector, error) {
	expr := ctx.Expr()
	if expr != nil {
		if fn := expr.Function(); fn != nil {
			args := query.NewArguments()
			if star := fn.STAR(); star != nil {
				args = append(args, query.NewArgumentWith(star.GetText()))
			} else {
				for _, arg := range fn.AllExpr() {
					args = append(args, query.NewArgumentWith(arg.GetText()))
				}
			}
			return query.NewFunctionWith(fn.Function_name().GetText(), args...), nil
		}
	}
	return query.NewColumnWithOptions(query.WithColumnName(ctx.GetText())), nil
}

func newDeleteWith(ctx antlr.IDelete_stmtContext) *query.Delete {
	tbl := query.NewTableWith(ctx.GetTable().GetText())
	opts := []query.DeleteOption{}
	if w := ctx.GetWhereExpr(); w != nil {
		opts = append(opts, query.WithDeleteCondition(query.NewConditionWith(newExprWith(w))))
	}
	return query.NewDeleteWith(tbl, opts...)
}

func newCopyWith(ctx antlr.ICopy_stmtContext) *query.Copy {
	opts := []query.CopyOption{}
	if ctx := ctx.Copy_column_list(); ctx != nil {
		columns := query.NewColumns()
		for _, column := range ctx.AllColumn_name() {
			columns = append(columns, query.NewColumnWithName(strings.UnEscapeString(column.GetText())))
		}
		opts = append(opts, query.WithCopyColumns(columns...))
	}
	if ctx := ctx.Copy_format(); ctx != nil {
		fmt := ctx.GetFormat_type().GetText()
		opts = append(opts, query.WithCopyFormat(fmt))
	}
	return query.NewCopyWith(
		strings.UnEscapeString(ctx.GetTable().GetText()),
		strings.UnEscapeString(ctx.Source_name().GetText()),
		opts...,
	)
}

func newTruncateWith(ctx antlr.ITruncate_stmtContext) *query.Truncate {
	tbls := query.NewTables()
	for _, table := range ctx.AllTable_name() {
		tbls = append(tbls, query.NewTableWith(table.GetText()))
	}
	return query.NewTruncateWith(tbls)
}

func newVacuumWith(ctx antlr.IVacuum_stmtContext) *query.Vacuum {
	if schema, ok := ctx.Schema_name().(*antlr.Schema_nameContext); ok {
		return query.NewVacuumWith(query.NewTableWith(schema.GetText()))
	}
	return query.NewVacuum()
}

func newBeginWith(ctx antlr.IBegin_stmtContext) *query.Begin {
	return query.NewBegin()
}

func newRollbackWith(ctx antlr.IRollback_stmtContext) *query.Rollback {
	return query.NewRollback()
}

func newCommitWith(ctx antlr.ICommit_stmtContext) *query.Commit {
	return query.NewCommit()
}

func newLiteralValueWith(ctx antlr.ILiteral_valueContext) *query.Literal {
	if ctx == nil {
		return nil
	}
	if ctx := ctx.String_literal(); ctx != nil {
		v := strings.UnEscapeString(ctx.GetText())
		return query.NewLiteralWith(v, query.WithLiteralType(query.StringLiteral))
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
			return query.NewCmpExpr(query.EQ, c, l)
		}
		if cmpExpr.NOT_EQ1() != nil || cmpExpr.NOT_EQ2() != nil {
			return query.NewCmpExpr(query.NEQ, c, l)
		}
		if cmpExpr.LT() != nil {
			return query.NewCmpExpr(query.LT, c, l)
		}
		if cmpExpr.GT() != nil {
			return query.NewCmpExpr(query.GT, c, l)
		}
		if cmpExpr.LT_EQ() != nil {
			return query.NewCmpExpr(query.LE, c, l)
		}
		if cmpExpr.GT_EQ() != nil {
			return query.NewCmpExpr(query.GE, c, l)
		}
	}
	return nil
}
