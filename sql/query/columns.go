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

// Columns represens a column array.
type Columns []*Column

// NewColumns returns a column array instance.
func NewColumns() Columns {
	return make(Columns, 0)
}

// NewColumnsWith returns a column array instance with the specified columns.
func NewColumnsWith(columns ...*Column) Columns {
	c := make(Columns, len(columns))
	copy(c, columns)
	return c
}

// Column returns a column array.
func (colums Columns) Columns() Columns {
	return colums
}
