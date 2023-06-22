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

// Table represents a table.
type Table struct {
	name string
}

// NewTableWith returns a new Table instance with the specified name.
func NewTableWith(name string) *Table {
	return &Table{
		name: name,
	}
}

// Table returns the table.
func (db *Table) Table() *Table {
	return db
}

// Name returns the table name.
func (db *Table) Name() string {
	return db.name
}

// String returns the string representation.
func (db *Table) String() string {
	return db.name
}