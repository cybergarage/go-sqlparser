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
	"fmt"
	"strings"

	"github.com/cybergarage/go-safecast/safecast"
	"github.com/cybergarage/go-sqlparser/sql/fn"
)

// ColumnOption represents a column option function.
type ColumnOption = func(*column)

// Column represents a column interface.
type Column interface {
	// Name returns the column name.
	Name() string
	// IsName returns true whether the column name is the specified one.
	IsName(string) bool
	// Definition returns the column definition.
	Definition() ColumnDef
	// Constraint returns the column constrains.
	Constraint() Constraint
	// DataType returns the column data type.
	DataType() DataType
	// DataTypeSize returns the column data type size.
	DataTypeSize() int
	// HasValue returns true whether the column has a value.
	HasValue() bool
	// SetValue sets a value.
	SetValue(any) error
	// Value returns the column value.
	Value() any
	// ValueType returns the column value type.
	ValueType() LiteralType
	// IsFunction returns true whether the column is a function.
	IsFunction() bool
	// Function returns the function if the column is a function.
	Function() (Function, bool)
	// Arguments returns the executor arguments.
	Arguments() Arguments
	// Copy returns a copy of the column.
	Copy() Column
	// DefinitionString returns the definition string representation.
	DefinitionString() string
	// String returns the string representation.
	String() string
	// ColumnHelper provides additional methods for columns in a query.
	ColumnHelper
}

// DefinitionStringer represents a definition stringer interface.
type DefinitionStringer interface {
	DefinitionString() string
}

// UpdatorStringer represents a updator stringer interface.
type UpdatorStringer interface {
	UpdatorString() string
}

// column represents a column.
type column struct {
	name string
	ColumnDef
	*Literal
	fn   Function
	args Arguments
}

// NewColumn returns a column instance.
func NewColumnWithOptions(opts ...ColumnOption) Column {
	col := &column{
		name:      "",
		ColumnDef: nil,
		Literal:   nil,
		fn:        nil,
		args:      fn.NewArguments(),
	}
	for _, opt := range opts {
		opt(col)
	}
	return col
}

// WithColumnName sets a column name.
func WithColumnName(name string) func(*column) {
	return func(col *column) {
		col.name = name
	}
}

// WithColumnData sets a column data.
func WithColumnData(data ColumnDef) func(*column) {
	return func(col *column) {
		col.ColumnDef = data
	}
}

// WithColumnLiteral sets a column data.
func WithColumnLiteral(l *Literal) func(*column) {
	return func(col *column) {
		col.Literal = l
	}
}

// WithColumnFunction sets a column function.
func WithColumnFunction(fn Function) func(*column) {
	return func(col *column) {
		col.fn = fn
	}
}

// WithColumnFunctionExecutor sets a column function.
func WithColumnFunctionExecutor(executor FunctionExecutor) func(*column) {
	return func(col *column) {
		col.fn = fn.NewFunctionWith(
			fn.WithFunctionExecutor(executor),
		)
	}
}

// WithColumnArguments sets column arguments.
func WithColumnArguments(args []string) func(*column) {
	return func(col *column) {
		col.args = fn.NewArgumentStrings(args...)
	}
}

// NewColumn returns a column instance.
func NewColumnWithName(name string) Column {
	return NewColumnWithOptions(WithColumnName(name))
}

// Name returns the column name.
func (col *column) Name() string {
	return col.name
}

// IsFunction returns true whether the column is a function.
func (col *column) IsFunction() bool {
	return col.fn != nil
}

// Executor returns the executor.
func (col *column) Function() (Function, bool) {
	if col.fn == nil {
		return nil, false
	}
	return col.fn, true
}

// Arguments returns the executor arguments.
func (col *column) Arguments() Arguments {
	return col.args
}

// IsName returns true whether the column name is the specified one.
func (col *column) IsName(name string) bool {
	return col.name == name
}

// SetValue sets a value.
func (col *column) SetValue(v any) error {
	col.Literal.SetValue(v)
	return col.SetDefinition(col.ColumnDef)
}

// SetDefinition sets the column definition to update the column value.
func (col *column) SetDefinition(dataDef ColumnDef) error {
	col.ColumnDef = dataDef

	if dataDef == nil || col.Literal == nil {
		return nil
	}

	switch dataDef.DataType() {
	case TinyIntType:
		var v int8
		if err := safecast.ToInt8(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case SmallIntType:
		var v int16
		if err := safecast.ToInt16(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case IntType, IntegerType:
		var v int32
		if err := safecast.ToInt32(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case BigIntType:
		var v int64
		if err := safecast.ToInt64(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case FloatType:
		var v float32
		if err := safecast.ToFloat32(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case DoubleType:
		var v float64
		if err := safecast.ToFloat64(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case BooleanType:
		var v bool
		if err := safecast.ToBool(col.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	}
	return nil
}

// Copy returns a copy of the column.
func (col *column) Copy() Column {
	return NewColumnWithOptions(
		WithColumnName(col.name),
		WithColumnData(col.ColumnDef),
		WithColumnLiteral(col.Literal),
	)
}

// String returns the string representation.
func (col *column) String() string {
	return col.name
}

// DefinitionString returns the definition string representation.
func (col *column) DefinitionString() string {
	if col.ColumnDef == nil {
		return col.name
	}
	return col.name + " " + col.ColumnDef.String()
}

// Definition returns the column definition.
func (col *column) Definition() ColumnDef {
	return col.ColumnDef
}

// UpdatorString returns the updator string representation.
func (col *column) UpdatorString() string {
	if fn, ok := col.Function(); ok {
		if executor, err := fn.Executor(); err == nil {
			strs := []string{}
			for n, arg := range col.args {
				strs = append(strs, fmt.Sprintf("%v", arg))
				if n == 0 {
					strs = append(strs, executor.Name())
				}
			}
			return strings.Join(strs, " ")
		}
	}
	if col.Literal != nil {
		return col.Literal.String()
	}
	return ""
}
