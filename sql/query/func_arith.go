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
	"strings"

	"github.com/cybergarage/go-safecast/safecast"
)

// ArithFunc represents an arithmetic function.
type ArithFunc func(any, any) (any, error)

// ArithResultSet represents a result set of an arithmetic function.
type ArithResultSet map[any]float64

// ArithFunction represents a base arithmetic function.
type ArithFunction struct {
	column   string
	operator ArithFunc
}

// NewArithFunctionWith returns a new base arithmetic function with the specified name and arithmetic.
func NewArithFunctionWith(name string, fn ArithFunc) *ArithFunction {
	return &ArithFunction{
		column:   strings.ToUpper(name),
		operator: fn,
	}
}

// Name returns the name of the function.
func (fn *ArithFunction) Name() string {
	return fn.column
}

// ColumnName returns the name of the column.
func (fn *ArithFunction) ColumnName() string {
	return fn.column
}

// Type returns the type of the function.
func (fn *ArithFunction) Type() FunctionType {
	return ArithFunctionType
}

// Execute returns the executed value with the specified arguments.
func (fn *ArithFunction) Execute(args ...any) (any, error) {
	if len(args) != 2 {
		return nil, newErrInvalidArguments(fn.column, args)
	}
	return fn.operator(args[0], args[1])
}

// NewAbsFunction returns a new abs function.
func NewAddFunction() *ArithFunction {
	return NewArithFunctionWith(
		AbsFunctionName,
		func(v1, v2 any) (any, error) {
			var fv1 float64
			err := safecast.ToFloat64(v1, &fv1)
			if err != nil {
				return nil, err
			}
			var fv2 float64
			err = safecast.ToFloat64(v2, &fv2)
			if err != nil {
				return nil, err
			}
			return (fv1 + fv2), nil
		},
	)
}
