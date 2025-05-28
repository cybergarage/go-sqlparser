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
	"math/rand"
)

// NewAbs returns a new abs function.
func NewAbs(opts ...ExecutorOption) Executor {
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

// NewFloor returns a new floor function.
func NewFloor(opts ...ExecutorOption) Executor {
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

// NewCeil returns a new ceil function.
func NewCeil(opts ...ExecutorOption) Executor {
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

// NewRound returns a new round function.
func NewRound(opts ...ExecutorOption) Executor {
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

// NewSqrt returns a new sqrt function.
func NewSqrt(opts ...ExecutorOption) Executor {
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

// NewLog returns a new log function.
func NewLog(opts ...ExecutorOption) Executor {
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

// NewLog10 returns a new log10 function.
func NewLog10(opts ...ExecutorOption) Executor {
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

// NewExp returns a new exp function.
func NewExp(opts ...ExecutorOption) Executor {
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

// NewPower returns a new power function.
func NewPower(opts ...ExecutorOption) Executor {
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

// NewMod returns a new mod function.
func NewMod(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		ModFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 2 {
				return nil, newErrInvalidArguments(ModFunctionName, args)
			}
			a := args[0]
			b := args[1]
			if b == 0 {
				return nil, newErrInvalidArguments(ModFunctionName, args)
			}
			return math.Mod(a, b), nil
		},
		opts...,
	)
}

// NewSin returns a new sin function.
func NewSin(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		SinFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(SinFunctionName, args)
			}
			return math.Sin(args[0]), nil
		},
		opts...,
	)
}

// NewCos returns a new cos function.
func NewCos(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		CosFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(CosFunctionName, args)
			}
			return math.Cos(args[0]), nil
		},
		opts...,
	)
}

// NewTan returns a new tan function.
func NewTan(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		TanFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(TanFunctionName, args)
			}
			return math.Tan(args[0]), nil
		},
		opts...,
	)
}

// NewRand returns a new rand function that generates a random float64 in [0.0, 1.0).
func NewRand(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		RandFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 0 {
				return nil, newErrInvalidArguments(RandFunctionName, args)
			}
			// Returns a random float64 in [0.0, 1.0).
			return math.Float64frombits(uint64(rand.Int63() & 0x7FFFFFFFFFFFFFFF)), nil
		},
		opts...,
	)
}

// NewPI returns a new pi function that returns the value of π (pi).
func NewPI(opts ...ExecutorOption) Executor {
	return NewMathFunctionWith(
		PiFunctionName,
		func(args []float64) (any, error) {
			if len(args) != 0 {
				return nil, newErrInvalidArguments(PiFunctionName, args)
			}
			return math.Pi, nil
		},
		opts...,
	)
}
