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

package sql

// Column represents a column.
type Column struct {
	name  string
	value interface{}
}

// NewColumn returns a column instance.
func NewColumn() *Column {
	col := &Column{
		name:  "",
		value: nil,
	}
	return col
}

// NewColumnWithName returns a column instance with the specified name.
func NewColumnWithName(name string) *Column {
	col := &Column{
		name: name,
	}
	return col
}

// SetName sets a name into the column.
func (col *Column) SetName(name string) {
	col.name = name
}

// Name returns the column name.
func (col *Column) Name() string {
	return col.name
}

// SetValue sets a value into the column.
func (col *Column) SetValue(value interface{}) {
	col.value = value
}

// Value returns the column value.
func (col *Column) Value() interface{} {
	return col.value
}

// String returns the string representation.
func (col *Column) String() string {
	return col.name
}
