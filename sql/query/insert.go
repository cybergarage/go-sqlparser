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

import (
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// Insert is a "INSERT" statement.
type Insert struct {
	*Table
	ColumnList
}

// NewInsertWith returns a new Insert statement instance with the specified parameters.
func NewInsertWith(tbl *Table, columns ColumnList) *Insert {
	return &Insert{
		Table:      tbl,
		ColumnList: columns,
	}
}

// StatementType returns the statement type.
func (stmt *Insert) StatementType() StatementType {
	return InsertStatement
}

// SetSchema sets a schema to update column values.
func (stmt *Insert) SetSchema(schema *Schema) error {
	for _, col := range stmt.ColumnList {
		schemaCol, err := schema.ColumnByName(col.Name())
		if err != nil {
			return err
		}
		err = col.SetDef(schemaCol.DataDef)
		if err != nil {
			return err
		}
	}
	return nil
}

// String returns the statement string representation.
func (stmt *Insert) String() string {
	strs := []string{
		"INSERT",
		"INTO",
		stmt.Table.String(),
		"(" + stmt.ColumnList.NameString() + ")",
		"VALUES",
		"(" + stmt.ColumnList.ValueString() + ")",
	}
	return strings.JoinWithSpace(strs)
}
