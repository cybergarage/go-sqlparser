// Copyright (C) 2024 The go-sqlparser Authors. All rights reserved.
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

// Use represents a "USE" statement interface.
type Use interface {
	Statement
	DatabaseName() string
}

type useStmt struct {
	*Database
}

func NewUseWith(name string) *useStmt {
	return &useStmt{
		Database: NewDatabaseWith(name),
	}
}

func (stmt *useStmt) StatementType() StatementType {
	return UseStatement
}

// String returns the statement string representation.
func (stmt *useStmt) String() string {
	strs := []string{"USE", stmt.Database.String()}
	return strings.JoinWithSpace(strs)
}
