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
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
)

// NewSchemaColumnsResultSetFromResultSet returns a new SchemaColumnsResultSet instance from the specified ResultSet.
func NewSchemaColumnsResultSetFromResultSet(rs ResultSet) (SchemaColumnsResultSet, error) {
	columns := []SchemaColumn{}
	for rs.Next() {
		row, err := rs.Row()
		if err != nil {
			return nil, err
		}
		catalog, err := row.ValueBy(SchemaColumnsCatalog)
		if err != nil {
			return nil, err
		}
		schema, err := row.ValueBy(SchemaColumnsSchema)
		if err != nil {
			return nil, err
		}
		table, err := row.ValueBy(SchemaColumnsTableName)
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
			WithColumnCatalog(catalog),
			WithColumnSchema(schema),
			WithColumnTable(table),
			WithColumnName(name),
			WithColumnDataType(typ),
		)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	return NewSchemaColumns(
		WithColumns(columns),
	), nil
}

// NewSchemaColumnsResultSetFromSchemas returns a new ResultSet instance from the specified schemas.
func NewSchemaColumnsResultSetFromSchemas(schemas []query.Schema) (ResultSet, error) {
	rsSchema, err := NewSchemaColumnsResultSetSchema()
	if err != nil {
		return nil, err
	}
	rows := []resultset.Row{}
	for _, schema := range schemas {
		for _, column := range schema.Columns() {
			obj := map[string]any{}
			for _, rsColumn := range rsSchema.Columns() {
				switch rsColumn.Name() {
				case SchemaColumnsCatalog:
					obj[SchemaColumnsCatalog] = DefaultSchemaColumnsCatalog
				case SchemaColumnsSchema:
					obj[SchemaColumnsSchema] = rsSchema.DatabaseName()
				case SchemaColumnsTableName:
					obj[SchemaColumnsTableName] = rsSchema.TableName()
				case SchemaColumnsColumnName:
					obj[SchemaColumnsColumnName] = column.Name()
				case SchemaColumnsDataType:
					obj[SchemaColumnsDataType] = column.DataType().String()
				default:
					return nil, newErrInvalidColumn(rsColumn.Name())
				}
			}
			row := resultset.NewRow(
				resultset.WithRowObject(obj),
				resultset.WithRowSchema(rsSchema),
			)
			rows = append(rows, row)
		}
	}
	rsOpts := []resultset.ResultSetOption{
		resultset.WithResultSetSchema(rsSchema),
		resultset.WithResultSetRowsAffected(0),
		resultset.WithResultSetRows(rows),
	}
	return resultset.NewResultSet(rsOpts...), nil
}
