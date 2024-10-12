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

package sql

import (
	"reflect"
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/query"
)

func TestQueryInterface(t *testing.T) {
	tests := []struct {
		q   any
		qst query.StatementType
		qt  reflect.Type
	}{
		{
			query.NewCreateDatabaseWith("", nil),
			query.CreateDatabaseStatement,
			reflect.TypeOf((*CreateDatabase)(nil)).Elem(),
		},
		{
			query.NewCreateTableWith(nil, nil),
			query.CreateTableStatement,
			reflect.TypeOf((*CreateTable)(nil)).Elem(),
		},
		{
			query.NewAlterDatabaseWith("", ""),
			query.AlterDatabaseStatement,
			reflect.TypeOf((*AlterDatabase)(nil)).Elem(),
		},
		{
			query.NewAlterTableWith(""),
			query.AlterTableStatement,
			reflect.TypeOf((*AlterTable)(nil)).Elem(),
		},
		{
			query.NewDropDatabaseWith("", nil),
			query.DropDatabaseStatement,
			reflect.TypeOf((*DropDatabase)(nil)).Elem(),
		},
		{
			query.NewDropTableWith(nil, nil),
			query.DropTableStatement,
			reflect.TypeOf((*DropTable)(nil)).Elem(),
		},
		{
			query.NewInsertWith(nil, nil),
			query.InsertStatement,
			reflect.TypeOf((*Insert)(nil)).Elem(),
		},
		{
			query.NewSelectWith(nil, nil, nil),
			query.SelectStatement,
			reflect.TypeOf((*Select)(nil)).Elem(),
		},
		{
			query.NewUpdateWith(nil, nil, nil),
			query.UpdateStatement,
			reflect.TypeOf((*Update)(nil)).Elem(),
		},
		{
			query.NewDeleteWith(nil),
			query.DeleteStatement,
			reflect.TypeOf((*Delete)(nil)).Elem(),
		},
		{
			query.NewBegin(),
			query.BeginStatement,
			reflect.TypeOf((*Begin)(nil)).Elem(),
		},
		{
			query.NewCommit(),
			query.CommitStatement,
			reflect.TypeOf((*Commit)(nil)).Elem(),
		},
		{
			query.NewRollback(),
			query.RollbackStatement,
			reflect.TypeOf((*Rollback)(nil)).Elem(),
		},
		{
			query.NewCopyWith("", ""),
			query.CopyStatement,
			reflect.TypeOf((*Copy)(nil)).Elem(),
		},
		{
			query.NewVacuumWith(nil),
			query.VacuumStatement,
			reflect.TypeOf((*Vacuum)(nil)).Elem(),
		},
		{
			query.NewTruncateWith(nil),
			query.TruncateStatement,
			reflect.TypeOf((*Truncate)(nil)).Elem(),
		},
	}

	for _, test := range tests {
		q, ok := test.q.(Query)
		if !ok {
			t.Errorf("The query type is invalid (%s)", test.q)
		}
		if q.StatementType() != test.qst {
			t.Errorf("The statement type is invalid (%d != %d)", q.StatementType(), test.qst)
		}
		qt := reflect.TypeOf(q)
		if ok := qt.Implements(test.qt); !ok {
			t.Errorf("%s is not implemented (%s)", qt, test.qt.Name())
		}
	}
}
