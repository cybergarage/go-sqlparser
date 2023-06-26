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

// AlterDatabase is a "ALTER DATABASE" statement.
type AlterDatabase struct {
	*Database
}

// NewAlterDatabaseWith returns a new AlterDatabase statement instance with the specified options.
func NewAlterDatabaseWith(name string) *AlterDatabase {
	return &AlterDatabase{
		Database: NewDatabaseWith(name),
	}
}

// StatementType returns the statement type.
func (stmt *AlterDatabase) StatementType() StatementType {
	return AlterDatabaseStatement
}

// String returns the statement string representation.
func (stmt *AlterDatabase) String() string {
	return "ALTER DATABASE " + stmt.name
}
