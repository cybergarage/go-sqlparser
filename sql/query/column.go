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
	// ValueString returns the column value string.
	ValueString() string
	// IsPlaceHolder returns true whether the column is a place holder.
	IsPlaceHolder() bool
	// IsFunction returns true whether the column is a function.
	IsFunction() (Function, bool)
	// Arguments returns the executor arguments.
	Arguments() []any
	// Copy returns a copy of the column.
	Copy() Column
	// String returns the string representation.
	String() string
}

type columnSelectorStringer interface {
	SelectorString() string
}

type columnDefStringer interface {
	DefinitionString() string
}

type columnUpdatorStringer interface {
	UpdatorString() string
}

// column represents a column.
type column struct {
	name string
	ColumnDef
	*Literal
	FunctionExecutor
	args []any
}

// NewColumn returns a column instance.
func NewColumnWithOptions(opts ...ColumnOption) Column {
	col := &column{
		name:             "",
		ColumnDef:        nil,
		Literal:          nil,
		FunctionExecutor: nil,
		args:             []any{},
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
func WithColumnFunction(fn FunctionExecutor) func(*column) {
	return func(col *column) {
		col.FunctionExecutor = fn
	}
}

// WithColumnArguments sets column arguments.
func WithColumnArguments(args []any) func(*column) {
	return func(col *column) {
		col.args = args
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
func (col *column) IsFunction() (Function, bool) {
	if col.FunctionExecutor == nil {
		return nil, false
	}
	return NewFunctionWith(
		WithFunctionExecutor(col.FunctionExecutor),
	), true
}

// Executor returns the executor.
func (col *column) Executor() (FunctionExecutor, bool) {
	if col.FunctionExecutor == nil {
		return nil, false
	}
	return col.FunctionExecutor, true
}

// Arguments returns the executor arguments.
func (col *column) Arguments() []any {
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
		if err := safecast.ToInt8(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case SmallIntType:
		var v int16
		if err := safecast.ToInt16(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case IntType, IntegerType:
		var v int32
		if err := safecast.ToInt32(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case BigIntType:
		var v int64
		if err := safecast.ToInt64(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case FloatType:
		var v float32
		if err := safecast.ToFloat32(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case DoubleType:
		var v float64
		if err := safecast.ToFloat64(col.Literal.v, &v); err != nil {
			return err
		}
		col.Literal.SetValue(v)
	case BooleanType:
		var v bool
		if err := safecast.ToBool(col.Literal.v, &v); err != nil {
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

// SelectorString returns the selector string representation.
func (col *column) SelectorString() string {
	return col.name
}

// UpdatorString returns the updator string representation.
func (col *column) UpdatorString() string {
	if executor, ok := col.Executor(); ok {
		strs := []string{}
		for n, arg := range col.args {
			strs = append(strs, fmt.Sprintf("%v", arg))
			if n == 0 {
				strs = append(strs, executor.Name())
			}
		}
		return strings.Join(strs, " ")
	}
	if col.Literal != nil {
		return col.Literal.String()
	}
	return ""
}
