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

type schema struct {
	databaeName string
	tableName   string
	columns     []Column
	// baseSchema is a temporary variable used only with WithSchemaSelector()
	baseSchema Schema
}

// SchemaOption represents a functional option for resultsetSchema.
type SchemaOption func(*schema) error

// WithSchemaDatabaseName returns a functional option for resultsetSchema.
func WithSchemaDatabaseName(name string) SchemaOption {
	return func(schema *schema) error {
		schema.databaeName = name
		return nil
	}
}

// WithSchemaTableName returns a functional option for resultsetSchema.
func WithSchemaTableName(name string) SchemaOption {
	return func(schema *schema) error {
		schema.tableName = name
		return nil
	}
}

// WithSchemaColumns returns a functional option for resultsetSchema.
func WithSchemaColumns(columns []Column) SchemaOption {
	return func(schema *schema) error {
		schema.columns = columns
		return nil
	}
}

// NewSchema returns a new Schema with the specified options.
func NewSchema(opts ...SchemaOption) Schema {
	schema := newSchema()
	for _, opt := range opts {
		opt(schema)
	}
	return schema
}

// NewSchemaFrom returns a new Schema from the specified options.
func NewSchemaFrom(opts ...SchemaOption) (Schema, error) {
	schema := newSchema()
	for _, opt := range opts {
		if err := opt(schema); err != nil {
			return nil, err
		}
	}
	return schema, nil
}

func newSchema() *schema {
	return &schema{
		baseSchema:  nil,
		databaeName: "",
		tableName:   "",
		columns:     []Column{},
	}
}

// DatabaseName returns the database name.
func (schema *schema) DatabaseName() string {
	return schema.databaeName
}

// TableName returns the table name.
func (schema *schema) TableName() string {
	return schema.tableName
}

// Columns returns the columns.
func (schema *schema) Columns() []Column {
	return schema.columns
}

func newErrColumnNotFound(name string) error {
	return fmt.Errorf("column (%s) %w", name, query.ErrNotFound)
}

// LookupColumn returns the column by the specified name.
func (schema *schema) LookupColumn(name string) (Column, error) {
	for _, column := range schema.columns {
		if column.Name() == name {
			return column, nil
		}
	}
	return nil, newErrColumnNotFound(name)
}

// Selectors returns the selectors.
func (schema *schema) Selectors() query.Selectors {
	selectors := query.NewSelectors()
	for _, column := range schema.columns {
		selectors = append(selectors, column)
	}
	return selectors
}
