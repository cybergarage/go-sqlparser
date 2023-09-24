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
	"strings"
)

// Table represents a table.
type Table struct {
	schema string
	name   string
}

// TableOption represents a table option function.
type TableOption = func(*Table)

// NewTableWith returns a new Table instance with the specified name.
func NewTableWith(name string, opts ...TableOption) *Table {
	tbl := &Table{
		name: name,
	}
	for _, opt := range opts {
		opt(tbl)
	}
	return tbl
}

// WithTableSchema sets a table schema.
func WithTableSchema(schema string) func(*Table) {
	return func(tbl *Table) {
		tbl.schema = schema
	}
}

// Name returns the schema and table name.
func (tbl *Table) Name() string {
	names := []string{}
	if 0 < len(tbl.schema) {
		names = append(names, tbl.schema)
	}
	if 0 < len(tbl.name) {
		names = append(names, tbl.name)
	}
	return strings.Join(names, ".")
}

// IsName returns true whether the table is named.
func (tbl *Table) IsName(name string) bool {
	return tbl.Name() == name
}

// SchemaName returns the table schema name.
func (tbl *Table) SchemaName() string {
	return tbl.schema
}

// IsSchemaName returns true whether the table is named.
func (tbl *Table) IsSchemaName(name string) bool {
	return tbl.schema == name
}

// TableName returns the table name for embedded use.
func (tbl *Table) TableName() string {
	return tbl.name
}

// IsTableName returns true whether the table is named.
func (tbl *Table) IsTableName(name string) bool {
	return tbl.name == name
}

// String returns the string representation.
func (tbl *Table) String() string {
	return tbl.Name()
}
