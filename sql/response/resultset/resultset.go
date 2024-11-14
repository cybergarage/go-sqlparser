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
	"github.com/cybergarage/go-sqlparser/sql/query"
)

// ResultSetSchema represents a schema interface in a resultset.
type ResultSetSchema interface {
	// DatabaseName returns the database name.
	DatabaseName() string
	// TableName returns the table name.
	TableName() string
	// Selectows returns the selectors.
	Selectors() query.SelectorList
	// Columns returns the columns.
	Columns() []Column
	// LookupColumn returns the column by the specified name.
	LookupColumn(name string) (Column, error)
}

// ResultSetRow represents a row interface.
type ResultSetRow interface {
	// Values returns the all values.
	Values() []any
	// ValueAt returns the value at the specified index.
	ValueAt(int) (any, error)
	// Scan scans the values.
	Scan(...any) error
	// ScanAt scans the value at the specified index.
	ScanAt(int, any) error
}

// ResultSet represents a response resultset interface.
type ResultSet interface {
	// Schema returns the schema.
	Schema() ResultSetSchema
	// Next returns the next row.
	Next() bool
	// Row returns the current row.
	Row() ResultSetRow
	// RowsAffected returns the number of rows affected.
	RowsAffected() uint64
	// Close closes the resultset.
	Close() error
}