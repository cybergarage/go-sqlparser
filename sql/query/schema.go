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
	Columns() ColumnList
	// LookupColumn returns a column by the specified name.
	LookupColumn(name string) (Column, error)
	// AddColumn adds a column.
	AddColumn(column Column) error
	// DropColumn drops a column by the specified name.
	DropColumn(name string) error
	// Indexes returns the all indexes.
	Indexes() IndexList
	// LookupIndex returns an index by the specified name.
	LookupIndex(name string) (Index, error)
	// AddIndex adds an index.
	AddIndex(index Index) error
	// DropIndex drops an index by the specified name.
	DropIndex(name string) error
	// Selectors returns the all selectors from the columns.
	Selectors() SelectorList
}

// schema represents a table schema.
type schema struct {
	Table
	ColumnList
	IndexList
}

// SchemaOption represents a schema option function.
type SchemaOption = func(*schema)

// NewSchemaWith returns a new schema statement instance with the parameters.
func NewSchemaWith(name string, opts ...SchemaOption) Schema {
	s := &schema{
		Table:      NewTableWith(name),
		ColumnList: NewColumns(),
		IndexList:  NewIndexes(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithSchemaColumns returns a schema option to set the columns.
func WithSchemaColumns(columns ColumnList) func(*schema) {
	return func(schema *schema) {
		schema.ColumnList = columns
	}
}

// WithSchemaIndexes returns a schema option to set the indexes.
func WithSchemaIndexes(idxes IndexList) func(*schema) {
	return func(schema *schema) {
		schema.IndexList = idxes
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

// AddColumn adds a column.
func (schema *schema) AddColumn(column Column) error {
	schema.ColumnList = append(schema.ColumnList, column)
	return nil
}

// AddIndex adds an index.
func (schema *schema) AddIndex(index Index) error {
	schema.IndexList = append(schema.IndexList, index)
	return nil
}

// DropColumn drops a column by the specified name.
func (schema *schema) DropColumn(name string) error {
	for n, column := range schema.ColumnList {
		if strings.EqualFold(column.Name(), name) {
			continue
		}
		schema.ColumnList = append(schema.ColumnList[:n], schema.ColumnList[n+1:]...)
		return nil
	}
	return newErrColumnNotFound(name)
}

// DropIndex drops an index by the specified name.
func (schema *schema) DropIndex(name string) error {
	for n, index := range schema.IndexList {
		if strings.EqualFold(index.Name(), name) {
			continue
		}
		schema.IndexList = append(schema.IndexList[:n], schema.IndexList[n+1:]...)
		return nil
	}
	return newErrIndexNotFound(name)
}
