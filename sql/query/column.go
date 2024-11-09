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
	Name() string
	Executor() FunctionExecutor
	Arguments() []any
	Constrains() ColumnConstraint
	IsName(string) bool
	HasLiteral() bool
	SetValue(any) error
	Value() any
	ValueType() LiteralType
	ValueString() string
	SetConstant(ColumnConstraint)
	SetDefinition(*DataDef) error
	Definition() *DataDef
	DefinitionString() string
	DataType() DataType
	SelectorString() string
	ExecuteUpdator(map[string]any) (any, error)
	UpdatorString() string
	Copy() Column
	String() string
}

// column represents a column.
type column struct {
	name string
	*DataDef
	*Literal
	FunctionExecutor
	consts ColumnConstraint
	args   []any
}

// NewColumn returns a column instance.
func NewColumnWithOptions(opts ...ColumnOption) Column {
	col := &column{
		name:             "",
		DataDef:          nil,
		Literal:          nil,
		FunctionExecutor: nil,
		args:             []any{},
		consts:           ConstraintNone,
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
func WithColumnData(data *DataDef) func(*column) {
	return func(col *column) {
		col.DataDef = data
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

func WithColumnConstant(c ColumnConstraint) func(*column) {
	return func(col *column) {
		col.consts |= c
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

// Executor returns the executor.
func (col *column) Executor() FunctionExecutor {
	return col.FunctionExecutor
}

// Arguments returns the executor arguments.
func (col *column) Arguments() []any {
	return col.args
}

// Constrains returns the column constrains.
func (col *column) Constrains() ColumnConstraint {
	return col.consts
}

// IsName returns true whether the column name is the specified one.
func (col *column) IsName(name string) bool {
	return col.name == name
}

// SetValue sets a value.
func (col *column) SetValue(v any) error {
	col.Literal.SetValue(v)
	return col.SetDefinition(col.DataDef)
}

// SetConstant sets a constant.
func (col *column) SetConstant(c ColumnConstraint) {
	col.consts |= c
}

// SetDefinition sets the column definition to update the column value.
func (col *column) SetDefinition(dataDef *DataDef) error {
	col.DataDef = dataDef

	if dataDef == nil || col.Literal == nil {
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
func (col *column) Copy() Column {
	return NewColumnWithOptions(
		WithColumnName(col.name),
		WithColumnData(col.DataDef),
		WithColumnLiteral(col.Literal),
	)
}

// String returns the string representation.
func (col *column) String() string {
	return col.name
}

// DefinitionString returns the definition string representation.
func (col *column) DefinitionString() string {
	if col.DataDef == nil {
		return col.name
	}
	return col.name + " " + col.DataDef.String()
}

// Definition returns the column definition.
func (col *column) Definition() *DataDef {
	return col.DataDef
}

// SelectorString returns the selector string representation.
func (col *column) SelectorString() string {
	return col.name
}

// ExecuteUpdator executes the executor with the specified row.
func (col *column) ExecuteUpdator(row map[string]any) (any, error) {
	newErrInvalidUpdateExecutor := func(col *column) error {
		return fmt.Errorf("%v is %w", col.UpdatorString(), ErrInvalid)
	}

	if col.FunctionExecutor == nil {
		return nil, newErrInvalidUpdateExecutor(col)
	}

	args := col.Arguments()
	if len(args) < 2 {
		return nil, newErrInvalidUpdateExecutor(col)
	}
	leftExprName, ok := args[0].(string)
	if !ok {
		return nil, newErrInvalidUpdateExecutor(col)
	}
	v, ok := row[leftExprName]
	if !ok {
		return nil, newErrInvalidUpdateExecutor(col)
	}
	args[0] = v
	rv, err := col.FunctionExecutor.Execute(args...)
	if err != nil {
		return nil, err
	}

	return rv, nil
}

// UpdatorString returns the updator string representation.
func (col *column) UpdatorString() string {
	if col.Executor() != nil {
		strs := []string{}
		for n, arg := range col.args {
			strs = append(strs, fmt.Sprintf("%v", arg))
			if n == 0 {
				strs = append(strs, col.Executor().Name())
			}
		}
		return strings.Join(strs, " ")
	}
	if col.Literal != nil {
		return col.Literal.String()
	}
	return ""
}
