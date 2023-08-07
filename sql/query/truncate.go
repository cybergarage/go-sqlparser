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

// Truncate is a "TRUNCATE" statement.
type Truncate struct {
	TableList
}

// NewTruncateWith returns a new truncate statement instance with the specified parameters.
func NewTruncateWith(tbls TableList) *Truncate {
	return &Truncate{
		TableList: tbls,
	}
}

// StatementType returns the statement type.
func (stmt *Truncate) StatementType() StatementType {
	return TruncateStatement
}

// String returns the statement string representation.
func (stmt *Truncate) String() string {
	strs := []string{
		"TRUNCATE",
		"TABLE",
		stmt.TableList.String(),
	}
	return strings.JoinWithSpace(strs)
}
