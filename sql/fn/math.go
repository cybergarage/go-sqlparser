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

package fn

import (
	"math"
	"strings"

	"github.com/cybergarage/go-safecast/safecast"
)

// MathFunc represents an math function.
type MathFunc func(any) (any, error)

// MathResultSet represents a result set of an math function.
type MathResultSet map[any]float64

// mathFunction represents a base math function.
type mathFunction struct {
	name string
	math MathFunc
}

// NewMathFunctionWith returns a new base math function with the specified name and math.
func NewMathFunctionWith(name string, fn MathFunc) Executor {
	return &mathFunction{
		name: strings.ToUpper(name),
		math: fn,
	}
}

// Name returns the name of the function.
func (fn *mathFunction) Name() string {
	return fn.name
}

// Type returns the type of the function.
func (fn *mathFunction) Type() FunctionType {
	return MathFunction
}

// Execute returns the executed value with the specified arguments.
func (fn *mathFunction) Execute(args ...any) (any, error) {
	if len(args) != 1 {
		return nil, newErrInvalidArguments(fn.name, args)
	}

	return fn.math(args[0])
}

// NewAbsFunction returns a new abs function.
func NewAbsFunction() Executor {
	return NewMathFunctionWith(
		AbsFunctionName,
		func(v any) (any, error) {
			var value float64
			err := safecast.ToFloat64(v, &value)
			if err != nil {
				return nil, err
			}
			return float64(math.Abs(value)), nil
		},
	)
}

// NewFloorFunction returns a new floor function.
func NewFloorFunction() Executor {
	return NewMathFunctionWith(
		FloorFunctionName,
		func(v any) (any, error) {
			var value float64
			err := safecast.ToFloat64(v, &value)
			if err != nil {
				return nil, err
			}
			return int64(math.Floor(value)), nil
		},
	)
}

// NewCeilFunction returns a new ceil function.
func NewCeilFunction() Executor {
	return NewMathFunctionWith(
		CeilFunctionName,
		func(v any) (any, error) {
			var value float64
			err := safecast.ToFloat64(v, &value)
			if err != nil {
				return nil, err
			}
			return int64(math.Ceil(value)), nil
		},
	)
}
