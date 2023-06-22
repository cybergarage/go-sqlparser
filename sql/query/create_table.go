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

// CreateTable is a "CREATE TABLE" statement.
type CreateTable struct {
	*Schema
}

// NewCreateTableWith returns a new CreateTable statement instance with the specified name.
func NewCreateTableWith(schema *Schema) *CreateTable {
	return &CreateTable{
		Schema: schema,
	}
}

// StatementType returns the statement type.
func (stmt *CreateTable) StatementType() StatementType {
	return CreateTableStatement
}

// String returns the statement string representation.
func (stmt *CreateTable) String() string {
	return stmt.Schema.String()
}
