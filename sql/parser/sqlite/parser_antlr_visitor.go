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
	return v.visitStatementList(ctx.AllSql_stmt_list())
}

func (v *antlrVisitor) visitStatementList(ctxList []antlr.ISql_stmt_listContext) query.StatementList {
	stmtList := query.NewStatementList()
	for _, ctx := range ctxList {
		for _, sqlStmt := range ctx.AllSql_stmt() {
			stmt := v.visitSqlStatement(sqlStmt)
			if stmt == nil {
				continue
			}
			stmtList = append(stmtList, stmt)
		}
	}
	return stmtList
}

func (v *antlrVisitor) visitSqlStatement(ctx antlr.ISql_stmtContext) query.Statement {
	if createStmt := ctx.Create_database_stmt(); createStmt != nil {
		dbName := createStmt.Database_name().GetText()
		return query.NewCreateDatabaseWith(dbName)
	}
	return nil
}
