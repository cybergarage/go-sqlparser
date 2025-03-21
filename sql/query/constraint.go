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

// Constraint represents a column constraint.
type Constraint int

const (
	// ConstraintNone represents no constraint.
	ConstraintNone       Constraint = 0x00
	PrimaryKeyConstraint Constraint = 0x01
	NotNullConstraint    Constraint = 0x02
	UniqueConstraint     Constraint = 0x04
)

// IsPrimaryKey returns true whether the column is a primary key.
func (c Constraint) IsPrimaryKey() bool {
	return (c & PrimaryKeyConstraint) != 0
}

// IsNotNull returns true whether the column is not null.
func (c Constraint) IsNotNull() bool {
	return (c & NotNullConstraint) != 0
}

// IsUnique returns true whether the column is unique.
func (c Constraint) IsUnique() bool {
	return (c & UniqueConstraint) != 0
}

// String returns the string representation.
func (c Constraint) String() string {
	switch c {
	case PrimaryKeyConstraint:
		return "PRIMARY KEY"
	case NotNullConstraint:
		return "NOT NULL"
	case UniqueConstraint:
		return "UNIQUE"
	}
	return ""
}
