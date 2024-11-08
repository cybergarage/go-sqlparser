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
	Tables() TableList
	IfExists() bool
}

// dropTable is a "DROP TABLE" statement.
type dropTable struct {
	TableList
	*IfExistsOpt
}

// NewDropTableWith returns a new dropTable statement instance with the specified parameters.
func NewDropTableWith(tbls TableList, ife *IfExistsOpt) *dropTable {
	return &dropTable{
		TableList:   tbls,
		IfExistsOpt: ife,
	}
}

// StatementType returns the statement type.
func (stmt *dropTable) StatementType() StatementType {
	return DropTableStatement
}

// String returns the statement string representation.
func (stmt *dropTable) String() string {
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
