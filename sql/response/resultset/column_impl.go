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

type column struct {
	name string
	t    DataType
	c    Constraint
}

// ResultSetColumnOptions represents a functional option for resultsetColumn.
type ColumnOption func(*column)

// WithColumnName returns a functional option for resultsetColumn.
func WithColumnName(name string) ColumnOption {
	return func(col *column) {
		col.name = name
	}
}

// WithColumnType returns a functional option for resultsetColumn.
func WithColumnType(t DataType) ColumnOption {
	return func(col *column) {
		col.t = t
	}
}

// WithColumnConstraint returns a functional option for resultsetColumn.
func WithColumnConstraint(c Constraint) ColumnOption {
	return func(col *column) {
		col.c = c
	}
}

// ResultSetColumn represents a resultset column interface.
func NewColumn(opts ...ColumnOption) Column {
	col := &column{
		name: "",
		t:    0,
		c:    0,
	}
	for _, opt := range opts {
		opt(col)
	}
	return col
}

// Name returns the column name.
func (col *column) Name() string {
	return col.name
}

// DataType returns the column type.
func (col *column) DataType() DataType {
	return col.t
}

// Constraint returns the column constraint.
func (col *column) Constraint() Constraint {
	return col.c
}
