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
)

// NewAbsFunction returns a new abs function.
func NewAbsFunction(opts ...ExecutorOption) Executor {
	ex := NewMathFunctionWith(
		AbsFunctionName,
		func(v float64) (any, error) {
			return math.Abs(v), nil
		},
		opts...,
	)
	return ex
}

// NewFloorFunction returns a new floor function.
func NewFloorFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		FloorFunctionName,
		func(v float64) (any, error) {
			return int(math.Floor(v)), nil
		},
		opts...,
	)
}

// NewCeilFunction returns a new ceil function.
func NewCeilFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		CeilFunctionName,
		func(v float64) (any, error) {
			return int(math.Ceil(v)), nil
		},
		opts...,
	)
}

func NewRoundFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		RoundFunctionName,
		func(v float64) (any, error) {
			return int(math.Round(v)), nil
		},
		opts...,
	)
}

// NewSqrtFunction returns a new sqrt function.
func NewSqrtFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		SqrtFunctionName,
		func(v float64) (any, error) {
			if v < 0 {
				return nil, newErrNegativeValue(v)
			}
			return math.Sqrt(v), nil
		},
		opts...,
	)
}

// NewLogFunction returns a new log function.
func NewLogFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		LogFunctionName,
		func(v float64) (any, error) {
			if v <= 0 {
				return nil, newErrNegativeValue(v)
			}
			return math.Log(v), nil
		},
		opts...,
	)
}

// NewLog10Function returns a new log10 function.
func NewLog10Function(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		Log10FunctionName,
		func(v float64) (any, error) {
			if v <= 0 {
				return nil, newErrNegativeValue(v)
			}
			return math.Log10(v), nil
		},
		opts...,
	)
}

// NewExpFunction returns a new exp function.
func NewExpFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		ExpFunctionName,
		func(v float64) (any, error) {
			return math.Exp(v), nil
		},
		opts...,
	)
}

// NewPowerFunction returns a new power function.
// func NewPowerFunction(opts ...ExecutorOption) Executor {
// 	return NewMathFunctionWith(
// 		PowerFunctionName,
// 		func(v float64, p float64) (any, error) {
// 			if v < 0 && p != math.Trunc(p) {
// 				return nil, ErrNegativePower
// 			}
// 			return math.Pow(v, p), nil
// 		},
// 		opts...,
// 	)
// }
