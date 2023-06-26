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

// DropTable is a "DROP TABLE" statement.
type DropTable struct {
	*Schema
	*IfExists
}

// NewDropTableWith returns a new DropTable statement instance with the specified name.
func NewDropTableWith(schema *Schema, ife *IfExists) *DropTable {
	return &DropTable{
		Schema:   schema,
		IfExists: ife,
	}
}

// StatementType returns the statement type.
func (stmt *DropTable) StatementType() StatementType {
	return DropTableStatement
}

// String returns the statement string representation.
func (stmt *DropTable) String() string {
	return stmt.Schema.String()
}
