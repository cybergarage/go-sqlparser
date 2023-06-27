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

import (
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// CreateDatabase is a "CREATE DATABASE" statement.
type CreateDatabase struct {
	*Database
	*IfNotExists
}

// NewCreateDatabaseWith returns a new CreateDatabase statement instance with the specified options.
func NewCreateDatabaseWith(name string, ifne *IfNotExists) *CreateDatabase {
	return &CreateDatabase{
		Database:    NewDatabaseWith(name),
		IfNotExists: ifne,
	}
}

// StatementType returns the statement type.
func (stmt *CreateDatabase) StatementType() StatementType {
	return CreateDatabaseStatement
}

// String returns the statement string representation.
func (stmt *CreateDatabase) String() string {
	elems := []string{
		"CREATE DATABASE"}
	if stmt.IfNotExists.Enabled() {
		elems = append(elems, stmt.IfNotExists.String())
	}
	elems = append(elems, stmt.name)
	return strings.Join(elems)
}
