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

package query

// AlterTable is a "ALTER TABLE" statement.
type AlterTable struct {
	*Schema
	*Table
}

// NewAlterTableWith returns a new AlterTable statement instance with the specified options.
func NewAlterTableWith(schemaName string, tblName string) *AlterTable {
	return &AlterTable{
		Schema: NewSchemaWith(schemaName, NewColumns(), NewIndexes()),
		Table:  NewTableWith(tblName),
	}
}

// StatementType returns the statement type.
func (stmt *AlterTable) StatementType() StatementType {
	return AlterTableStatement
}

// String returns the statement string representation.
func (stmt *AlterTable) String() string {
	return stmt.Schema.String()
}
