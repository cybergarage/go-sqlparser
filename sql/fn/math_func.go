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
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(AbsFunctionName, args)
			}
			return math.Abs(args[0]), nil
		},
		opts...,
	)
	return ex
}

// NewFloorFunction returns a new floor function.
func NewFloorFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		FloorFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(FloorFunctionName, args)
			}
			return int(math.Floor(args[0])), nil
		},
		opts...,
	)
}

// NewCeilFunction returns a new ceil function.
func NewCeilFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		CeilFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(CeilFunctionName, args)
			}
			return int(math.Ceil(args[0])), nil
		},
		opts...,
	)
}

func NewRoundFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		RoundFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(RoundFunctionName, args)
			}
			return int(math.Round(args[0])), nil
		},
		opts...,
	)
}

// NewSqrtFunction returns a new sqrt function.
func NewSqrtFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		SqrtFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(SqrtFunctionName, args)
			}
			v := args[0]
			if v <= 0 {
				return nil, newErrNegativeValue(args)
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
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(LogFunctionName, args)
			}
			v := args[0]
			if v <= 0 {
				return nil, newErrNegativeValue(args)
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
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(Log10FunctionName, args)
			}
			v := args[0]
			if v <= 0 {
				return nil, newErrNegativeValue(args)
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
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(ExpFunctionName, args)
			}
			return math.Exp(args[0]), nil
		},
		opts...,
	)
}

// NewPowerFunction returns a new power function.
func NewPowerFunction(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		PowerFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 2 {
				return nil, newErrInvalidArguments(PowerFunctionName, args)
			}
			base := args[0]
			exp := args[1]
			if base < 0 && exp != float64(int(exp)) {
				return nil, newErrNegativeValue(args)
			}
			return math.Pow(base, exp), nil
		},
		opts...,
	)
}
