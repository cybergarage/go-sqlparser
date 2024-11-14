// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
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

package resultset

import (
	"fmt"

	"github.com/cybergarage/go-sqlparser/sql/query"
)

type resultsetSchema struct {
	databaeName string
	tableName   string
	columns     []Column
}

// ResultSetSchemaOptions represents a functional option for resultsetSchema.
type ResultSetSchemaOptions func(*resultsetSchema)

// WithResultSetSchemaDatabaseName returns a functional option for resultsetSchema.
func WithResultSetSchemaDatabaseName(name string) ResultSetSchemaOptions {
	return func(schema *resultsetSchema) {
		schema.databaeName = name
	}
}

// WithResultSetSchemaTableName returns a functional option for resultsetSchema.
func WithResultSetSchemaTableName(name string) ResultSetSchemaOptions {
	return func(schema *resultsetSchema) {
		schema.tableName = name
	}
}

// WithResultSetSchemaResultSetColumns returns a functional option for resultsetSchema.
func WithResultSetSchemaResultSetColumns(columns []Column) ResultSetSchemaOptions {
	return func(schema *resultsetSchema) {
		schema.columns = columns
	}
}

// NewResultSetSchema returns a new resultsetSchema.
func NewResultSetSchema(opts ...ResultSetSchemaOptions) ResultSetSchema {
	schema := &resultsetSchema{
		databaeName: "",
		tableName:   "",
		columns:     []Column{},
	}
	for _, opt := range opts {
		opt(schema)
	}
	return schema
}

// DatabaseName returns the database name.
func (schema *resultsetSchema) DatabaseName() string {
	return schema.databaeName
}

// TableName returns the table name.
func (schema *resultsetSchema) TableName() string {
	return schema.tableName
}

// Columns returns the columns.
func (schema *resultsetSchema) Columns() []Column {
	return schema.columns
}

func newErrColumnNotFound(name string) error {
	return fmt.Errorf("column (%s) %w", name, query.ErrNotFound)
}

// LookupColumn returns the column by the specified name.
func (schema *resultsetSchema) LookupColumn(name string) (Column, error) {
	for _, column := range schema.columns {
		if column.Name() == name {
			return column, nil
		}
	}
	return nil, newErrColumnNotFound(name)
}

// Selectors returns the selectors.
func (schema *resultsetSchema) Selectors() query.SelectorList {
	selectors := query.NewSelectors()
	for _, column := range schema.columns {
		selectors = append(selectors, column)
	}
	return selectors
}