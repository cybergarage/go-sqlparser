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

// Vacuum is a "VACUUM" statement.
type Vacuum struct {
	*Table
}

// NewVacuumWith returns a new vacuum statement instance with the specified parameters.
func NewVacuumWith(tbl *Table) *Vacuum {
	return &Vacuum{
		Table: tbl,
	}
}

// StatementType returns the statement type.
func (stmt *Vacuum) StatementType() StatementType {
	return VacuumStatement
}

// String returns the statement string representation.
func (stmt *Vacuum) String() string {
	strs := []string{
		"VACUUM",
		stmt.Table.String(),
	}
	return strings.JoinWithSpace(strs)
}
