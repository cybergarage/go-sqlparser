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

// DataType represents a data type.
type DataType = query.DataType

// Constraint represents a column constraint.
type Constraint = query.Constraint

// Column represents a column interface in a resultset.
type Column interface {
	// Name returns the column name.
	Name() string
	// DataType returns the data type.
	DataType() DataType
	// Constraint returns the column constraint.
	Constraint() Constraint
	// String returns the string representation of the column.
	String() string
}
