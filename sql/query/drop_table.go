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

// DropTable represents a "DROP TABLE" statement interface.
type DropTable interface {
	Statement
	// Tables returns the table list.
	Tables() TableList
	// TableNames returns the table names.
	TableNames() []string
	// IfExists returns true if the "IF EXISTS" option is set.
	IfExists() bool
}

// dropTableStmt is a "DROP TABLE" statement.
type dropTableStmt struct {
	TableList
	*IfExistsOpt
}

// NewDropTableWith returns a new dropTable statement instance with the specified parameters.
func NewDropTableWith(tbls TableList, ife *IfExistsOpt) *dropTableStmt {
	return &dropTableStmt{
		TableList:   tbls,
		IfExistsOpt: ife,
	}
}

// StatementType returns the statement type.
func (stmt *dropTableStmt) StatementType() StatementType {
	return DropTableStatement
}

// String returns the statement string representation.
func (stmt *dropTableStmt) String() string {
	strs := []string{
		"DROP",
		"TABLE",
	}
	if stmt.IfExists() {
		strs = append(strs, stmt.IfExistsOpt.String())
	}
	strs = append(strs, stmt.TableList.String())
	return strings.JoinWithSpace(strs)
}
