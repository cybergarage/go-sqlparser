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

// CreateDatabase is a "CREATE DATABASE" statement.
type CreateDatabase struct {
	name string
}

// NewCreateDatabaseWith returns a new CreateDatabase statement instance with the specified name.
func NewCreateDatabaseWith(name string) *CreateDatabase {
	return &CreateDatabase{
		name: name,
	}
}

// StatementType returns the statement type.
func (stmt *CreateDatabase) StatementType() StatementType {
	return CreateDatabaseStatement
}

// String returns the statement string representation.
func (stmt *CreateDatabase) String() string {
	return "CREATE DATABASE " + stmt.name
}
