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

// WithSchemaTableSchema returns a functional option for WithSchemaSelector().
func WithSchemaTableSchema(tblSchema query.Schema) SchemaOption {
	return func(schema *schema) error {
		schema.tableName = tblSchema.TableName()
		schema.tableSchema = tblSchema
		return nil
	}
}

// WithSchemaSelector returns a functional option for resultsetSchema.
func WithSchemaSelectors(selectors query.Selectors) SchemaOption {
	return func(schema *schema) error {
		if schema.tableSchema == nil {
			return fmt.Errorf("base schema is nil; set it first using WithSchemaQuerySchema")
		}

		rsSchemaColumns := []Column{}

		for _, selector := range selectors {
			var rsSchemaColumn Column
			fx, ok := selector.Function()
			if !ok {
				selectorName := selector.Name()
				schemaColumn, err := schema.tableSchema.LookupColumn(selectorName)
				if err != nil {
					return err
				}
				rsSchemaColumn = schemaColumn
			} else {
				dataType, err := query.NewDataTypeForFunction(fx)
				if err != nil {
					return err
				}
				rsSchemaColumn = NewColumn(
					WithColumnName(selector.String()),
					WithColumnType(dataType),
					WithColumnFunction(fx),
				)
			}
			rsSchemaColumns = append(rsSchemaColumns, rsSchemaColumn)
		}

		schema.columns = rsSchemaColumns

		return nil
	}
}
