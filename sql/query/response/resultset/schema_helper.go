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

// WithSchemaQuerySchema returns a functional option for WithSchemaSelector().
func WithSchemaQuerySchema(querySchema query.Schema) SchemaOption {
	return func(schema *schema) error {
		schema.querySchema = querySchema
		return nil
	}
}

// WithSchemaSelector returns a functional option for resultsetSchema.
func WithSchemaSelectors(selectors query.Selectors) SchemaOption {
	return func(schema *schema) error {
		if schema.querySchema == nil {
			return fmt.Errorf("query schema is nil; set it first using WithSchemaQuerySchema")
		}

		rsSchemaColums := []Column{}

		for _, selector := range selectors {
			var rsSchemaColumn Column
			fx, ok := selector.Function()
			if !ok {
				selectorName := selector.Name()
				shemaColumn, err := schema.querySchema.LookupColumn(selectorName)
				if err != nil {
					return err
				}
				rsSchemaColumn, err = NewColumnFrom(shemaColumn)
				if err != nil {
					return err
				}
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
			rsSchemaColums = append(rsSchemaColums, rsSchemaColumn)
		}

		schema.columns = rsSchemaColums

		return nil
	}
}
