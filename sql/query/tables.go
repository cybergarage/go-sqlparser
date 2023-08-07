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

// TableList represens a column array.
type TableList []*Table

// NewTables returns a column array instance.
func NewTables() TableList {
	return make(TableList, 0)
}

// NewTablesWith returns a column array instance with the specified columns.
func NewTablesWith(tables ...*Table) TableList {
	c := make(TableList, len(tables))
	copy(c, tables)
	return c
}

// Table returns a column array.
func (tbls TableList) Tables() TableList {
	return tbls
}

// String returns a string representation.
func (tbls TableList) String() string {
	strs := make([]string, len(tbls))
	for n, tbl := range tbls {
		strs[n] = tbl.String()
	}
	return strings.JoinWithComma(strs)
}
