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

// SchemaName returns the table name.
func (schema *Schema) SchemaName() string {
	return schema.TableName()
}
