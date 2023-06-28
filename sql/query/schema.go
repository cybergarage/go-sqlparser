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

import "github.com/cybergarage/go-sqlparser/sql/util/strings"

// Schema represents a table schema.
type Schema struct {
	*Table
	Columns
	Indexes
}

// NewSchemaWith returns a new schema statement instance with the parameters.
func NewSchemaWith(name string, colums Columns, indexes Indexes) *Schema {
	return &Schema{
		Table:   NewTableWith(name),
		Columns: colums,
		Indexes: indexes,
	}
}

// Schema returns the schema.
func (schema *Schema) Schema() *Schema {
	return schema
}

// String returns the statement string representation.
func (schema *Schema) String() string {
	columsStr := "("
	columsStr += schema.Columns.DefString()
	if 0 < len(schema.Indexes) {
		columsStr += " ,"
		columsStr += schema.Indexes.DefString()
	}
	columsStr += ")"

	elems := []string{
		"CREATE TABLE",
		schema.Table.String(),
		columsStr,
	}
	return strings.Join(elems)
}
