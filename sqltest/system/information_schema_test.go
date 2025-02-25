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
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/system"
)

func TestInfomationSchema(t *testing.T) {
	t.Run("Columns", func(t *testing.T) {
		tests := []struct {
			dbName   string
			tblName  string
			expected string
		}{
			{
				dbName:   "sqltest1740462150031924000",
				tblName:  "test",
				expected: "SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = 'sqltest1740462150031924000' AND TABLE_NAME = 'test'",
			},
		}
		for _, test := range tests {
			t.Run(test.dbName+"."+test.tblName, func(t *testing.T) {
				sysStmt, err := system.NewSchemaColumnsStatement(
					system.WithSchemaColumnsStatementDatabaseName(test.dbName),
					system.WithSchemaColumnsStatementTableNames([]string{test.tblName}),
				)
				if err != nil {
					t.Error(err)
					return
				}
				stmt := sysStmt.Statement()
				if stmt == nil {
					t.Error("stmt is nil")
					return
				}
				stmtStr := stmt.String()
				if !strings.EqualFold(stmtStr, test.expected) {
					t.Errorf("expected: %s, got: %s", test.expected, stmtStr)
				}
			})
		}

	})
}
