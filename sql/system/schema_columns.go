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

// PostgreSQL: Documentation: 17: 35.17. columns
// https://www.postgresql.org/docs/current/infoschema-columns.html

const (
	DefaultSchemaColumnsCatalog = "def"

	SchemaColumnsCatalog    = "TABLE_CATALOG"
	SchemaColumnsSchema     = "TABLE_SCHEMA"
	SchemaColumnsTable      = "TABLE_NAME"
	SchemaColumnsColumnName = "COLUMN_NAME"
	SchemaColumnsDataType   = "DATA_TYPE"
	SchemaColumnsQuery      = "SELECT * information_schema.COLUMNS"
)

// SchemaColumnsResultSet represents a schema columns result set.
type SchemaColumnsResultSet interface {
	// Columns returns the columns.
	Columns() []SchemaColumn
	// LookupColumn returns the column by name.
	LookupColumn(name string) (SchemaColumn, error)
}

// NewSchemaColumnsQueryWithTableNames returns a new schema columns query with table names.
func NewSchemaColumnsQueryWithTableNames(tableNames []string) string {
	if len(tableNames) == 0 {
		return SchemaColumnsQuery
	}
	query := SchemaColumnsQuery + " WHERE "
	for n, tableName := range tableNames {
		if 0 < n {
			query += " OR "
		}
		query += "TABLE_NAME = '" + tableName + "'"
	}
	return query
}
