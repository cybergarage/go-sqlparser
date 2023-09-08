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

// ColumnConstraint represents a column constraint.
type ColumnConstraint int

const (
	// ColumnConstraintNone represents no constraint.
	ColumnConstraintNone       ColumnConstraint = 0x00
	ColumnConstraintPrimaryKey ColumnConstraint = 0x01
	ColumnConstraintNotNull    ColumnConstraint = 0x02
	ColumnConstraintUnique     ColumnConstraint = 0x04
)

// Column represents a column.
type Column struct {
	name string
	*DataDef
	*Literal
	*BindParam
	consts ColumnConstraint
}

// IsPrimaryKey returns true whether the column is a primary key.
func (c ColumnConstraint) IsPrimaryKey() bool {
	return (c & ColumnConstraintPrimaryKey) != 0
}

// String returns the string representation.
func (c ColumnConstraint) String() string {
	switch c {
	case ColumnConstraintPrimaryKey:
		return "PRIMARY KEY"
	case ColumnConstraintNotNull:
		return "NOT NULL"
	case ColumnConstraintUnique:
		return "UNIQUE"
	}
	return ""
}
