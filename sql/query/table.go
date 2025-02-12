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

// TableNameSep represents a table name separator.
const TableNameSep = "."

// Table represents a table interface.
type Table interface {
	// FullTableName returns the full name of the table including schema and table name.
	FullTableName() string
	// IsFullTableName returns true if the provided name matches the full name of the table,
	IsFullTableName(name string) bool
	// SchemaName returns the table schema name.
	SchemaName() string
	// IsSchemaName returns true whether the table is named.
	IsSchemaName(name string) bool
	// TableName returns the table name for embedded use.
	TableName() string
	// IsTableName returns true whether the table is named.
	IsTableName(name string) bool
}

// table represents a table.
type table struct {
	schema string
	name   string
}

// TableOption represents a table option function.
type TableOption = func(*table)

// NewTableWith returns a new table instance with the specified name.
func NewTableWith(name string, opts ...TableOption) Table {
	tbl := &table{
		name:   name,
		schema: "",
	}
	for _, opt := range opts {
		opt(tbl)
	}
	return tbl
}

// WithTableSchema sets a table schema.
func WithTableSchema(schema string) func(*table) {
	return func(tbl *table) {
		tbl.schema = schema
	}
}

// FullTableName returns the full name of the table including schema and table name.
func (tbl *table) FullTableName() string {
	names := []string{}
	if 0 < len(tbl.name) {
		if !strings.Contains(tbl.name, TableNameSep) {
			if 0 < len(tbl.schema) {
				names = append(names, tbl.schema)
			}
		}
		names = append(names, tbl.name)
	}
	return strings.Join(names, TableNameSep)
}

// IsFullTableName returns true if the provided name matches the full name of the table,.
func (tbl *table) IsFullTableName(name string) bool {
	return strings.EqualFold(tbl.FullTableName(), name)
}

// SchemaName returns the table schema name.
func (tbl *table) SchemaName() string {
	return tbl.schema
}

// IsSchemaName returns true whether the table is named.
func (tbl *table) IsSchemaName(name string) bool {
	return strings.EqualFold(tbl.schema, name)
}

// TableName returns the table name for embedded use.
func (tbl *table) TableName() string {
	return tbl.name
}

// IsTableName returns true whether the table is named.
func (tbl *table) IsTableName(name string) bool {
	return strings.EqualFold(tbl.name, name)
}

// String returns the string representation.
func (tbl *table) String() string {
	return tbl.FullTableName()
}
