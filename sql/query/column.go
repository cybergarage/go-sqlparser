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

import (
	"github.com/cybergarage/go-safecast/safecast"
)

// Column represents a column.
type Column struct {
	name string
	*DataDef
	*Literal
	*BindParam
}

// ColumnOption represents a column option function.
type ColumnOption = func(*Column)

// NewColumn returns a column instance.
func NewColumnWithOptions(opts ...ColumnOption) *Column {
	col := &Column{
		name:    "",
		DataDef: nil,
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
func WithColumnData(data *DataDef) func(*Column) {
	return func(col *Column) {
		col.DataDef = data
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

// SetValue sets a value.
func (col *Column) SetValue(v any) error {
	col.Literal.SetValue(v)
	return col.SetDef(col.DataDef)
}

// SetDef sets the column definition to update the column value.
func (col *Column) SetDef(dataDef *DataDef) error {
	col.DataDef = dataDef
	if dataDef == nil {
		return nil
	}

	switch dataDef.DataType() {
	case TinyIntData:
		var v int8
		if err := safecast.ToInt8(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case SmallIntData:
		var v int16
		if err := safecast.ToInt16(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case IntData, IntegerData:
		var v int32
		if err := safecast.ToInt32(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case BigIntData:
		var v int64
		if err := safecast.ToInt64(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case FloatData:
		var v float32
		if err := safecast.ToFloat32(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case DoubleData:
		var v float64
		if err := safecast.ToFloat64(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case BooleanData:
		var v bool
		if err := safecast.ToBool(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	}
	return nil
}

// Copy returns a copy of the column.
func (col *Column) Copy() *Column {
	return NewColumnWithOptions(
		WithColumnName(col.name),
		WithColumnData(col.DataDef),
		WithColumnLiteral(col.Literal),
	)
}

// String returns the string representation.
func (col *Column) String() string {
	return col.name
}

// DefinitionString returns the definition string representation.
func (col *Column) DefinitionString() string {
	if col.DataDef == nil {
		return col.name
	}
	return col.name + " " + col.DataDef.String()
}

// SelectorString returns the selector string representation.
func (col *Column) SelectorString() string {
	return col.name
}
