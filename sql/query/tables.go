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

// Tables represens a column array.
type Tables []*Table

// NewTables returns a column array instance.
func NewTables() Tables {
	return make(Tables, 0)
}

// NewTablesWith returns a column array instance with the specified columns.
func NewTablesWith(columns ...*Table) Tables {
	c := make(Tables, len(columns))
	copy(c, columns)
	return c
}

// Table returns a column array.
func (colums Tables) Tables() Tables {
	return colums
}

// String returns a string representation.
func (colums Tables) String() string {
	strs := make([]string, len(colums))
	for n, tbl := range colums {
		strs[n] = tbl.String()
	}
	return strings.JoinWithComma(strs)
}
