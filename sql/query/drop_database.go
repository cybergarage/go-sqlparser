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

// DropDatabase is a "DROP DATABASE" statement.
type DropDatabase struct {
	*Schema
	*IfExists
}

// NewDropDatabaseWith returns a new DropDatabase statement instance with the specified parameters.
func NewDropDatabaseWith(schema *Schema, ife *IfExists) *DropDatabase {
	return &DropDatabase{
		Schema:   schema,
		IfExists: ife,
	}
}

// StatementType returns the statement type.
func (stmt *DropDatabase) StatementType() StatementType {
	return DropDatabaseStatement
}

// String returns the statement string representation.
func (stmt *DropDatabase) String() string {
	return stmt.Schema.String()
}
