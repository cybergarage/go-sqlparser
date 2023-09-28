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

import "strings"

// Schema represents a table schema.
type Schema struct {
	*Table
	ColumnList
	IndexList
}

// SchemaOption represents a schema option function.
type SchemaOption = func(*Schema)

// NewSchemaWith returns a new schema statement instance with the parameters.
func NewSchemaWith(name string, opts ...SchemaOption) *Schema {
	s := &Schema{
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
func WithSchemaColumns(columns ColumnList) func(*Schema) {
	return func(schema *Schema) {
		schema.ColumnList = columns
	}
}

// WithSchemaIndexes returns a schema option to set the indexes.
func WithSchemaIndexes(idxes IndexList) func(*Schema) {
	return func(schema *Schema) {
		schema.IndexList = idxes
	}
}

// SchemaTable returns the table.
func (schema *Schema) SchemaTable() *Table {
	return schema.Table
}

// SchemaName returns the table name.
func (schema *Schema) SchemaName() string {
	return schema.TableName()
}

// AddColumn adds a column.
func (schema *Schema) AddColumn(column *Column) error {
	schema.ColumnList = append(schema.ColumnList, column)
	return nil
}

// AddIndex adds an index.
func (schema *Schema) AddIndex(index *Index) error {
	schema.IndexList = append(schema.IndexList, index)
	return nil
}

// DropColumn drops a column by the specified name.
func (schema *Schema) DropColumn(name string) error {
	for n, column := range schema.ColumnList {
		if strings.EqualFold(column.Name(), name) {
			continue
		}
		schema.ColumnList = append(schema.ColumnList[:n], schema.ColumnList[n+1:]...)
		return nil
	}
	return newErrColumnNotFound(name)
}
