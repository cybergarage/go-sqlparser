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
	"strings"

	"github.com/cybergarage/go-sqlparser/sql"
)

// PostgreSQL: Documentation: 17: 35.17. columns
// https://www.postgresql.org/docs/current/infoschema-columns.html

const (
	DefaultSchemaColumnsCatalog = "def"

	InformationSchema       = "information_schema"
	SchemaColumns           = InformationSchema + ".COLUMNS"
	SchemaColumnsCatalog    = "TABLE_CATALOG"
	SchemaColumnsSchema     = "TABLE_SCHEMA"
	SchemaColumnsTable      = "TABLE_NAME"
	SchemaColumnsColumnName = "COLUMN_NAME"
	SchemaColumnsDataType   = "DATA_TYPE"
	SchemaColumnsQuery      = "SELECT * FROM " + SchemaColumns
)

// IsSchemaColumsQuery returns true if the statement is a schema columns query.
func IsSchemaColumsQuery(stmt sql.Select) bool {
	tables := stmt.From()
	if len(tables) != 1 {
		return false
	}
	if !strings.EqualFold(tables[0].FullTableName(), SchemaColumns) {
		return false
	}
	return true
}
