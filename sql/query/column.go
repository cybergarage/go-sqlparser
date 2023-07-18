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

// Column represents a column.
type Column struct {
	name string
	*Data
	*Literal
	*BindParam
}

// ColumnOption represents a column option function.
type ColumnOption = func(*Column)

// NewColumn returns a column instance.
func NewColumnWithOptions(opts ...ColumnOption) *Column {
	col := &Column{
		name:    "",
		Data:    nil,
		Literal: nil,
	}
	for _, opt := range opts {
		opt(col)
	}
	return col
}

// WithColumnName sets a column name.
func WithColumnName(name string) func(*Column) {
	return func(col *Column) {
		col.name = name
	}
}

// WithColumnData sets a column data.
func WithColumnData(data *Data) func(*Column) {
	return func(col *Column) {
		col.Data = data
	}
}

// WithColumnLiteral sets a column data.
func WithColumnLiteral(l *Literal) func(*Column) {
	return func(col *Column) {
		col.Literal = l
	}
}

// NewColumn returns a column instance.
func NewColumnWithName(name string) *Column {
	return NewColumnWithOptions(WithColumnName(name))
}

// Name returns the column name.
func (col *Column) Name() string {
	return col.name
}

// IsName returns true whether the column name is the specified one.
func (col *Column) IsName(name string) bool {
	return col.name == name
}

// String returns the string representation.
func (col *Column) String() string {
	return col.name
}

// String returns the string representation.
func (col *Column) DefString() string {
	if col.Data == nil {
		return col.name
	}
	return col.name + " " + col.Data.String()
}
