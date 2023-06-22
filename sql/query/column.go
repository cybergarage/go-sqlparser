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
	name  string
	t     DataType
	value any
}

// NewColumn returns a column instance.
func NewColumnWith(name string, t DataType, v any) *Column {
	col := &Column{
		name:  name,
		t:     t,
		value: v,
	}
	return col
}

// Name returns the column name.
func (col *Column) Name() string {
	return col.name
}

// Type returns the column type.
func (col *Column) Type() DataType {
	return col.t
}

// Value returns the column value.
func (col *Column) Value() interface{} {
	return col.value
}

// String returns the string representation.
func (col *Column) String() string {
	return col.name
}
