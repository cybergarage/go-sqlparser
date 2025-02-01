// Copyright (C) 2025 The go-sqlparser Authors. All rights reserved.
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

package system

import (
	"github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
)

// ResultSet represents a response resultset interface.
type ResultSet = resultset.ResultSet

type schemaColumns struct {
	tableCatalog string
	tableSchema  string
	tableName    string
	columns      []Column
}

// NewSchemaColumns returns a new SchemaColumns instance.
type SchemaColumnOption func(*schemaColumns)

// WithTableCatalog returns a SchemaColumnOption that sets the table catalog.
func WithTableCatalog(catalog string) SchemaColumnOption {
	return func(s *schemaColumns) {
		s.tableCatalog = catalog
	}
}

// WithTableSchema returns a SchemaColumnOption that sets the table schema.
func WithTableSchema(schema string) SchemaColumnOption {
	return func(s *schemaColumns) {
		s.tableSchema = schema
	}
}

// WithTableName returns a SchemaColumnOption that sets the table name.
func WithTableName(name string) SchemaColumnOption {
	return func(s *schemaColumns) {
		s.tableName = name
	}
}

// WithColumns returns a SchemaColumnOption that sets the columns.
func WithColumns(columns []Column) SchemaColumnOption {
	return func(s *schemaColumns) {
		s.columns = columns
	}
}

// NewSchemaColumns returns a new SchemaColumns instance.
func NewSchemaColumns(opts ...SchemaColumnOption) SchemaColumns {
	s := &schemaColumns{
		tableCatalog: DefaultSchemaColumnsCatalog,
		tableSchema:  "",
		tableName:    "",
		columns:      []Column{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// NewSchemaColumnsFromResultSet returns a new SchemaColumns instance from a ResultSet.
func NewSchemaColumnsFromResultSet(rs ResultSet) (SchemaColumns, error) {
	var tableSchema string
	var tableName string
	schema := rs.Schema()
	if schema != nil {
		tableSchema = schema.DatabaseName()
		tableName = schema.TableName()
	}

	columns := []Column{}
	for rs.Next() {
		row, err := rs.Row()
		if err != nil {
			return nil, err
		}
		name, err := row.ValueBy(SchemaColumnsColumnName)
		if err != nil {
			return nil, err
		}
		typ, err := row.ValueBy(SchemaColumnsDataType)
		if err != nil {
			return nil, err
		}
		column, err := NewColumn(
			WithColumnName(name),
			WithColumnDataType(typ),
		)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	return NewSchemaColumns(
		WithTableSchema(tableSchema),
		WithTableName(tableName),
		WithColumns(columns),
	), nil
}

// TableCatalog returns the table catalog.
func (s *schemaColumns) TableCatalog() string {
	return s.tableCatalog
}

// TableSchema returns the table schema.
func (s *schemaColumns) TableSchema() string {
	return s.tableSchema
}

// TableName returns the table name.
func (s *schemaColumns) TableName() string {
	return s.tableName
}

// Columns returns the columns.
func (s *schemaColumns) Columns() []Column {
	return s.columns
}
