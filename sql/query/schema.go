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

// Schema represents a table schema interface.
type Schema interface {
	// FullTableName returns the full table name.
	FullTableName() string
	// SchemaName returns the schema name.
	SchemaName() string
	// TableName returns the table name.
	TableName() string
	// Columns returns the all columns.
	Columns() Columns
	// LookupColumn returns a column by the specified name.
	LookupColumn(name string) (Column, error)
	// AddColumn adds a column.
	AddColumn(column Column) error
	// DropColumn drops a column by the specified name.
	DropColumn(name string) error
	// Indexes returns the all indexes.
	Indexes() Indexes
	// LookupIndex returns an index by the specified name.
	LookupIndex(name string) (Index, error)
	// AddIndex adds an index.
	AddIndex(index Index) error
	// DropIndex drops an index by the specified name.
	DropIndex(name string) error
	// Selectors returns the all selectors from the columns.
	Selectors() Selectors
}

// schema represents a table schema.
type schema struct {
	Table
	columns Columns
	indexes Indexes
}

// SchemaOption represents a schema option function.
type SchemaOption = func(*schema)

// NewSchemaWith returns a new schema statement instance with the parameters.
func NewSchemaWith(name string, opts ...SchemaOption) Schema {
	s := &schema{
		Table:   NewTableWith(name),
		columns: NewColumns(),
		indexes: NewIndexes(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithSchemaColumns returns a schema option to set the columns.
func WithSchemaColumns(columns Columns) func(*schema) {
	return func(schema *schema) {
		schema.columns = columns
	}
}

// WithSchemaIndexes returns a schema option to set the indexes.
func WithSchemaIndexes(idxes Indexes) func(*schema) {
	return func(schema *schema) {
		schema.indexes = idxes
	}
}

// SchemaTable returns the table.
func (schema *schema) SchemaTable() Table {
	return schema.Table
}

// SchemaName returns the table name.
func (schema *schema) SchemaName() string {
	return schema.TableName()
}

// SchemaName returns the table name.
func (schema *schema) Name() string {
	return schema.TableName()
}

// Selectors returns the all selectors from the columns.
func (schema *schema) Selectors() Selectors {
	return schema.columns.Selectors()
}

// Columns returns the all columns.
func (schema *schema) Columns() Columns {
	return schema.columns
}

// LookupColumn returns a column by the specified name.
func (schema *schema) LookupColumn(name string) (Column, error) {
	return schema.columns.LookupColumn(name)
}

// AddColumn adds a column.
func (schema *schema) AddColumn(column Column) error {
	schema.columns = append(schema.columns, column)
	return nil
}

// DropColumn drops a column by the specified name.
func (schema *schema) DropColumn(name string) error {
	for n, column := range schema.columns {
		if strings.EqualFold(column.Name(), name) {
			continue
		}
		schema.columns = append(schema.columns[:n], schema.columns[n+1:]...)
		return nil
	}
	return newErrColumnNotFound(name)
}

// AddIndex adds an index.
func (schema *schema) AddIndex(index Index) error {
	schema.indexes = append(schema.indexes, index)
	return nil
}

// Indexes returns the all indexes.
func (schema *schema) Indexes() Indexes {
	return schema.indexes
}

// LookupIndex returns an index by the specified name.
func (schema *schema) LookupIndex(name string) (Index, error) {
	return schema.indexes.LookupIndex(name)
}

// DropIndex drops an index by the specified name.
func (schema *schema) DropIndex(name string) error {
	for n, index := range schema.indexes {
		if strings.EqualFold(index.Name(), name) {
			continue
		}
		schema.indexes = append(schema.indexes[:n], schema.indexes[n+1:]...)
		return nil
	}
	return newErrIndexNotFound(name)
}
