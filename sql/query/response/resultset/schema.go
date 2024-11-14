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

// Schema represents a schema interface in a resultset.
type Schema interface {
	// DatabaseName returns the database name.
	DatabaseName() string
	// TableName returns the table name.
	TableName() string
	// Selectows returns the selectors.
	Selectors() query.Selectors
	// Columns returns the columns.
	Columns() []Column
	// LookupColumn returns the column by the specified name.
	LookupColumn(name string) (Column, error)
}
