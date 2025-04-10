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

// [Information Schema - PostgreSQL
// https://www.postgresql.org/docs/current/information-schema.html
// The INFORMATION\_SCHEMA TABLES Table - MySQL]
// https://dev.mysql.com/doc/refman/8.0/en/information-schema.html
// Information Schema TABLES Table - MariaDB Knowledge Base
// https://mariadb.com/kb/en/information-schema-tables-table/
// System Information Schema Views (Transact-SQL) - SQL Server | Microsoft Learn
// https://learn.microsoft.com/en-us/sql/relational-databases/system-information-schema-views/system-information-schema-views-transact-sql?view=sql-server-ver16)
// Introduction to INFORMATION\_SCHEMA | BigQuery | Google Cloud
// https://cloud.google.com/bigquery/docs/information-schema-intro

// SchemaColumnsResultSet represents a schema columns resultset.
func NewSchemaColumnsResultSetSchema() (resultset.Schema, error) {
	columnDefs := []struct {
		name     string
		dataType resultset.DataType
	}{
		{SchemaColumnsCatalog, query.TextType},
		{SchemaColumnsSchema, query.TextType},
		{SchemaColumnsTableName, query.TextType},
		{SchemaColumnsColumnName, query.TextType},
		{SchemaColumnsDataType, query.TextType},
	}

	colums := []resultset.Column{}
	for _, columDef := range columnDefs {
		colums = append(colums, resultset.NewColumn(
			resultset.WithColumnName(columDef.name),
			resultset.WithColumnType(columDef.dataType),
		))
	}

	return resultset.NewSchema(
		resultset.WithSchemaDatabaseName(InformationSchema),
		resultset.WithSchemaTableName(InformationSchemaColumns),
		resultset.WithSchemaColumns(colums),
	), nil
}
