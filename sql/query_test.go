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
		q reflect.Type
		t reflect.Type
	}{
		{
			reflect.TypeOf(query.NewCreateDatabaseWith("test", nil)),
			reflect.TypeOf((*CreateDatabase)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewCreateTableWith(nil, nil)),
			reflect.TypeOf((*CreateTable)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewAlterDatabaseWith("from", "to")),
			reflect.TypeOf((*AlterDatabase)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewAlterTableWith("test")),
			reflect.TypeOf((*AlterTable)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewDropDatabaseWith("test", nil)),
			reflect.TypeOf((*DropDatabase)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewDropTableWith(nil, nil)),
			reflect.TypeOf((*DropTable)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewInsertWith(nil, nil)),
			reflect.TypeOf((*Insert)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewSelectWith(nil, nil, nil)),
			reflect.TypeOf((*Select)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewUpdateWith(nil, nil, nil)),
			reflect.TypeOf((*Update)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewDeleteWith(nil)),
			reflect.TypeOf((*Delete)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewBegin()),
			reflect.TypeOf((*Begin)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewCommit()),
			reflect.TypeOf((*Commit)(nil)).Elem(),
		},
		{
			reflect.TypeOf(query.NewRollback()),
			reflect.TypeOf((*Rollback)(nil)).Elem(),
		},
	}

	for _, test := range tests {
		if ok := test.q.Implements(test.t); !ok {
			t.Errorf("%s is not implemented (%s)", test.q, test.t.Name())
		}
	}
}
